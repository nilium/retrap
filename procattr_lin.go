//+build darwin dragonfly freebsd netbsd openbsd linux solaris

package main

import "syscall"

var procattr = &syscall.SysProcAttr{
	Setpgid: true,
}
