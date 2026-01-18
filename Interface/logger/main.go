package main

import "fmt"

type Logger interface {
	Log(msg string)
}

type ConsoleLogger struct{}

type FileLogger struct{}

func (c ConsoleLogger) Log(msg string) {
	fmt.Println("console:", msg)
}

func (f FileLogger) Log(msg string) {
	fmt.Println("file:", msg)
}

func doWork(l Logger) {
	l.Log("starting work")
	l.Log("working...")
	l.Log("done")
}

func main() {
	c := ConsoleLogger{}
	f := FileLogger{}

	doWork(c)
	doWork(f)
}
