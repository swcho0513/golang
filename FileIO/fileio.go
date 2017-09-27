package main

import (
  "io"
  "os"
)

func main() {
  fi, err := os.Open("input.txt")
  if err != nil {
    panic(err)
  }
  defer fi.Close()

  fo, err := os.Create("output.txt")
  if err != nil {
    panic(err)
  }
  defer fo.Close()

  buf := make([]byte, 1024)

  for {
    cnt, err := fi.Read(buf)
    if err != nil && err != io.EOF {
      panic(err)
    }

    if cnt == 0 {
      break
    }

    _, err = fo.Write(buf[:cnt])
    if err != nil {
      panic(err)
    }
  }
}
