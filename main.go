package main

import "github.com/echocat/golang-kata-1/v1/presentation/cmd"

func main() {
	cmdapp := cmd.ConsoleApp{}
	cmdapp.InitInMemoryConsoleApp()
	cmdapp.Start()
}
