package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	countdownStart = 3
	finalWord      = "Go!"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {
}

func (d DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		_, err := fmt.Fprintln(out, i)
		if err != nil {
			panic(err)
		}
		sleeper.Sleep()
	}
	_, err := fmt.Fprint(out, finalWord)
	if err != nil {
		fmt.Println("encountered an error")
	}
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
