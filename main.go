package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

func main() {
  var menu [][]string

  filename := "./menu_output.txt"

  content, err := ioutil.ReadFile(filename)
  if err != nil {
    panic(err)
  }

  result := strings.Split(string(content), "\n")

  for sku, i := range(result) {
    menuitem := strings.Split(i, ",")
    menuitem = append([]string{strconv.FormatInt(int64(sku), 10)}, menuitem...)
    menu = append(menu, menuitem)
  }

  fmt.Println(menu)

}
