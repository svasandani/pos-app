package main

import (
  "fmt"
  // "io/ioutil"
  // "strings"
  // "strconv"

  "github.com/svasandani/pos-app/menu"
)

func main() {
  menufile := menu.GetMenu()
  fmt.Println(menufile)

  menu.AddMenuItem("chicken", "masala", "250")
  menufile = menu.GetMenu()
  fmt.Println(menufile)

}
