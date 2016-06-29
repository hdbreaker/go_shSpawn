package main

import "syscall"

func main(){
	syscall.Exec("/bin/sh", nil, nil)
}
