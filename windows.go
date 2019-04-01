// +build windows

package main

// BINARY defines the binary to be executed
const BINARY = "powershell.exe"

// ARGS defines additional arguments to pass
var ARGS = []string{"-nop", "-w", "hidden"}
