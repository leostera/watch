package main

import (
  "os"
  "time"

  tm "github.com/buger/goterm"
)

func main() {
  command := os.Args[1:]
  tm.Clear()
  loop(1*time.Second, func() {
    render(command)
  })
}

func render(command []string) {
  tm.MoveCursor(0,0)
  tm.Println(command)
  tm.Flush()
}

func loop(d time.Duration, fn func()) {
  time.After(d)
  fn()
  loop(d, fn)
}
