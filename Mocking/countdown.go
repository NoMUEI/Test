package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWorld = "Go"
const countStaart = 3

type Sleeper interface {
	Sleep()
}
type SpySleeper struct {
	Calls int
}

type ConfigurableSleeper struct {
	duration time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countStaart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}
	for i := countStaart; i > 0; i-- {

	}
	time.Sleep(1 * time.Second)
	fmt.Fprintln(out, finalWorld)
}
