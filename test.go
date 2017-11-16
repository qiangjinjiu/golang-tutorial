package main

import (
  "log"

  "github.com/djherbis/atime"
)

func main() {
  at, err := atime.Stat("d:\\ndpslides\\数字切片扫描\\绍兴市立医院\\05110432889\\01.1_20171108_05110432889_21.ndpi")
  if err != nil {
    log.Fatal(err.Error())
  }
  log.Println(at)
}