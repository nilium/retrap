package main

import "syscall"

var procattr = &syscall.SysProcAttr{
	Rfork: syscall.RFNOTEG,
}
