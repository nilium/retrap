// +build freebsd

package main

import "syscall"

// Signal names to integers
var traps = map[string]syscall.Signal{
	"abrt":   syscall.Signal(0x06),
	"alrm":   syscall.Signal(0x0e),
	"bus":    syscall.Signal(0x0a),
	"chld":   syscall.Signal(0x14),
	"cont":   syscall.Signal(0x13),
	"emt":    syscall.Signal(0x07),
	"fpe":    syscall.Signal(0x08),
	"hup":    syscall.Signal(0x01),
	"ill":    syscall.Signal(0x04),
	"info":   syscall.Signal(0x1d),
	"int":    syscall.Signal(0x02),
	"io":     syscall.Signal(0x17),
	"iot":    syscall.Signal(0x06),
	"kill":   syscall.Signal(0x09),
	"librt":  syscall.Signal(0x21),
	"lwp":    syscall.Signal(0x20),
	"pipe":   syscall.Signal(0x0d),
	"prof":   syscall.Signal(0x1b),
	"quit":   syscall.Signal(0x03),
	"segv":   syscall.Signal(0x0b),
	"stop":   syscall.Signal(0x11),
	"sys":    syscall.Signal(0x0c),
	"term":   syscall.Signal(0x0f),
	"thr":    syscall.Signal(0x20),
	"trap":   syscall.Signal(0x05),
	"tstp":   syscall.Signal(0x12),
	"ttin":   syscall.Signal(0x15),
	"ttou":   syscall.Signal(0x16),
	"urg":    syscall.Signal(0x10),
	"usr1":   syscall.Signal(0x1e),
	"usr2":   syscall.Signal(0x1f),
	"vtalrm": syscall.Signal(0x1a),
	"winch":  syscall.Signal(0x1c),
	"xcpu":   syscall.Signal(0x18),
	"xfsz":   syscall.Signal(0x19),
}
