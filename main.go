package main

import (
  "fmt"
  // "io/ioutil"
  // "strings"
  // "strconv"

  "github.com/svasandani/pos-app/menu"
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
  // menu.EditMenuItem(76, "", "", "350")
  // // menu.AddMenuItem("chicken", "masala", "250")
  // menufile = menu.GetMenu()
  // fmt.Println(menufile)
  // fmt.Println(menu.Menu)

  test := transaction.NewTransaction()

  transaction.AddDish(&test, 72)
  transaction.AddDish(&test, 23)
  transaction.AddDish(&test, 45)
  transaction.AddDish(&test, 32)
  transaction.AddDish(&test, 65)
  transaction.AddDish(&test, 71)
  transaction.AddDish(&test, 12)
  transaction.AddDish(&test, 34)

  transaction.ServeDish(&test, 0)
  transaction.ServeDish(&test, 1)
  transaction.ServeDish(&test, 2)
  transaction.ServeDish(&test, 3)
  transaction.ServeDish(&test, 4)
  transaction.ServeDish(&test, 5)
  transaction.ServeDish(&test, 6)
  transaction.ServeDish(&test, 7)

  fmt.Println(test)
  fmt.Println(test.Total)
  fmt.Println(menu.Menu)

  transaction.Pay(&test, "cash")

}
