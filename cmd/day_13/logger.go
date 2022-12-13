package main

import "fmt"

type Logger struct {
	Level   int
	Enabled bool
}

func (l *Logger) AddLevel() {
	l.Level++
}

func (l *Logger) SubLevel() {
	l.Level--
}

func (l *Logger) Println(args ...interface{}) {
	if !l.Enabled {
		return
	}

	for i := 0; i < l.Level; i++ {
		fmt.Print("  ")
	}
	fmt.Println(args...)
}
