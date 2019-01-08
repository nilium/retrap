// Command retrap is a small program that spawns a subprocess and forwards remapped
// signals to it.
//
// Usage: retrap [SRC:DST] [SRC:-] -- CMD [ARG...]
//
// Signals are mapped by passing SRC:DST flags before the "--" separator,
// where SRC is the signal to trap in the parent process and DST is the signal
// to send to the subprocess. So, "INT:TERM" will trap an interrupt signal to
// the parent process and send a terminate signal to the subprocess (the CMD).
//
// In addition, a signal can be discarded by specifying "SRC:-", where DST is
// a hyphen. For example, if "INT:-" were passed, the parent process would trap
// and discard interrupt signals.
//
// By default, all signals are sent to the subprocess unless mapped.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"golang.org/x/sys/unix"
)

func usage(ec int, args ...interface{}) {
	if len(args) > 0 {
		fmt.Fprint(os.Stderr, fmt.Sprint(args...), "\n")
	}
	fmt.Fprintln(os.Stderr, "USAGE: retrap [SRC:DST] [SRC:-] -- CMD [ARG...]")
	os.Exit(ec)
}

func fatal(args ...interface{}) {
	fmt.Fprint(os.Stderr, fmt.Sprint(args...), "\n")
	os.Exit(1)
}

func tosignal(name string) (s syscall.Signal, ok bool) {
	s, ok = traps[strings.ToLower(name)]
	if !ok {
		i, err := strconv.Atoi(name)
		if err != nil {
			return 0, false
		}
		s, ok = syscall.Signal(i), true
	}
	return s, ok
}

func main() {
	var err error
	argv := os.Args[1:]
	if len(argv) == 0 {
		usage(2)
	}

	cwd, err := os.Getwd()
	if err != nil {
		fatal("cannot get working directory: ", err)
	}

	partition := -1
	for i, v := range argv {
		if v == "-help" || v == "--help" || v == "-h" {
			usage(2)
		}
		if v == "--" {
			partition = i
			break
		}
	}
	if partition == -1 {
		usage(1, "missing command")
	}

	toremap, cmd := argv[:partition], argv[partition+1:]
	if len(cmd) == 0 {
		usage(1, "missing command")
	}

	// Trap all known signals
	trap := make(chan os.Signal, 1)
	for _, signame := range traps {
		signal.Notify(trap, signame)
	}

	remap := map[os.Signal]syscall.Signal{}
	swallows := map[os.Signal]struct{}{}
	for _, t := range toremap {
		p := strings.SplitN(t, ":", 2)
		if len(p) != 2 {
			usage(1, "invalid signal trap (must be SIG:SIG or SIG:-)")
		}

		signame, ok := tosignal(p[0])
		if !ok {
			usage(1, "invalid input signal name ", strconv.Quote(p[0]), " in ", strconv.Quote(t))
		}

		if p[1] == "-" {
			swallows[signame] = struct{}{}
		} else {
			to, ok := tosignal(p[1])
			if !ok {
				usage(1, "invalid target signal name ", strconv.Quote(p[0]), " in ", strconv.Quote(t))
			}
			remap[signame] = to
		}
		signal.Notify(trap, signame)
	}

	p, err := exec.LookPath(cmd[0])
	if err != nil {
		// Let this fail through on forkexec
		p = cmd[0]
	}

	// TODO: make child process believe input/output is a tty when appropriate
	attr := syscall.ProcAttr{
		Files: []uintptr{
			os.Stdin.Fd(),
			os.Stdout.Fd(),
			os.Stderr.Fd(),
		},
		Sys: procattr,
		Env: os.Environ(),
		Dir: cwd,
	}
	var pid int
	if pid, err = syscall.ForkExec(p, cmd, &attr); err != nil {
		fatal("forkexec: ", err)
	}

	// Forward signals after remapping
	go func() {
		for s := range trap {
			if _, ok := swallows[s]; ok {
				continue
			} else if s2, ok := remap[s]; ok {
				s = s2
			}
			if sig, ok := s.(syscall.Signal); ok {
				unix.Kill(pid, sig)
			}
		}
	}()

	// Wait for children to exit
	var status unix.WaitStatus
	for err = nil; err != unix.ECHILD; _, err = unix.Wait4(-1, &status, 0, nil) {
	}
	os.Exit(status.ExitStatus())
}
