package main

import (
  "bytes"
  "fmt"
  "io"
  "log"
  "os"
  "os/exec"
  "time"

  tm "github.com/buger/goterm"
)

func main() {
  command := os.Args[1:]

  tm.Clear()
  tm.MoveCursor(0,0)

  loop(1*time.Second, func() {

    output, err := run(command)
    safe(err)

    render(output)

  })
}

func run(command []string) (bytes.Buffer, error) {
  name := command[0]
  args := command[1:]
  cmd  := exec.Command(name, args...)

  cmdPipe, err := cmd.StdoutPipe()
  if err != nil {
    return bytes.Buffer{}, err;
  }

  if err := cmd.Start(); err != nil {
    return bytes.Buffer{}, err;
  }

  pipeReader, pipeWriter := io.Pipe()

  go func() {
    _, err := io.Copy(pipeWriter, cmdPipe)
    // fixme: return error through a channel
    safe(err)
    pipeWriter.Close()
  } ()

  var buf bytes.Buffer
  _, err2 := io.Copy(&buf, pipeReader)
  safe(err2)

  return buf, nil
}

func l (s string) { fmt.Println(s) }

func safe(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func render(output bytes.Buffer) {
  tm.Println(output.String())
  tm.Flush()
}

func loop(d time.Duration, fn func()) {
  time.After(d)
  fn()
  loop(d, fn)
}
