package main

import (
  "fmt"
  // "io/ioutil"
  // "strings"
  // "strconv"

  // "github.com/svasandani/pos-app/menu"
  "github.com/svasandani/pos-app/transaction"
)

func main() {
  // menufile := menu.GetMenu()
  // fmt.Println(menufile)
  //
  // menu.AddMenuItem("chicken", "masala", "250")
  // // menu.DeleteMenuItem(77)
  // menufile = menu.GetMenu()
  // fmt.Println(menufile)
  //
  // menu.EditMenuItem(77, "", "", "350")
  // // menu.AddMenuItem("chicken", "masala", "250")
  // menufile = menu.GetMenu()
  // fmt.Println(menufile)

  fmt.Println(transaction.NewTransaction())

}
