// +build linux

package main

import "syscall"

// Signal names to integers
var traps = map[string]syscall.Signal{
	"abrt":   syscall.Signal(0x06),
	"alrm":   syscall.Signal(0x0e),
	"bus":    syscall.Signal(0x07),
	"chld":   syscall.Signal(0x11),
	"cld":    syscall.Signal(0x11),
	"cont":   syscall.Signal(0x12),
	"fpe":    syscall.Signal(0x08),
	"hup":    syscall.Signal(0x01),
	"ill":    syscall.Signal(0x04),
	"int":    syscall.Signal(0x02),
	"io":     syscall.Signal(0x1d),
	"iot":    syscall.Signal(0x06),
	"kill":   syscall.Signal(0x09),
	"pipe":   syscall.Signal(0x0d),
	"poll":   syscall.Signal(0x1d),
	"prof":   syscall.Signal(0x1b),
	"pwr":    syscall.Signal(0x1e),
	"quit":   syscall.Signal(0x03),
	"segv":   syscall.Signal(0x0b),
	"stkflt": syscall.Signal(0x10),
	"stop":   syscall.Signal(0x13),
	"sys":    syscall.Signal(0x1f),
	"term":   syscall.Signal(0x0f),
	"trap":   syscall.Signal(0x05),
	"tstp":   syscall.Signal(0x14),
	"ttin":   syscall.Signal(0x15),
	"ttou":   syscall.Signal(0x16),
	"unused": syscall.Signal(0x1f),
	"urg":    syscall.Signal(0x17),
	"usr1":   syscall.Signal(0x0a),
	"usr2":   syscall.Signal(0x0c),
	"vtalrm": syscall.Signal(0x1a),
	"winch":  syscall.Signal(0x1c),
	"xcpu":   syscall.Signal(0x18),
	"xfsz":   syscall.Signal(0x19),
}
