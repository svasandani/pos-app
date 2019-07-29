package menu

import (
  "fmt"
  // "io"
  "io/ioutil"
  "log"
  // "os"
  "strings"
  "strconv"
  "sync"

  poslog "github.com/svasandani/pos-app/log"
)

type Dish struct {
  SKU int `json:"sku"`
  Name string `json:"name"`
  Price string `json:"price"`
}

var Menu []Dish
var sku int = 0;
var filename string = "./menu/menu.csv"
var mtx sync.RWMutex

func init() {
  getMenuFromFile()
}

func getMenuFromFile() {
  // mtx.Lock()
  // defer mtx.Unlock()

  sku = 0;
  Menu = nil

  content, err := ioutil.ReadFile(filename)

  if err != nil {
    log.Fatalf(err.Error())
  }

  result := strings.Split(string(content), "\n")

  for _, item := range(result) {
    menuitem := strings.Split(item, ",")

    if len(menuitem) < 2 {continue;}

    for i, string := range(menuitem) {
      menuitem[i] = strings.TrimSpace(string)
    }

    dish := Dish {SKU: getSKU(), Name: menuitem[0], Price: menuitem[1],}

    Menu = append(Menu, dish)
  }
}

func writeMenuToFile() {
  // mtx.Lock()
  // defer mtx.Unlock()

  output := ""

  for _, item := range(Menu) {
    output += fmt.Sprintf("%v, %v\n", item.Name, item.Price)
  }

  bytes := []byte(output)

  ioutil.WriteFile(filename, bytes, 0644)
}

func DishFromSKU(sku int) Dish {
  return Menu[sku-1]
}

func PriceFromSKU(sku int) int {
  price, _ := strconv.ParseInt(Menu[sku-1].Price, 10, 0)
  return int(price)
}

// func intToString(int int) string {
//   return strconv.FormatInt(int64(int), 10)
// }

func getSKU() int {
  sku++
  return sku
}

func GetMenu() []Dish {
  getMenuFromFile()
  return Menu
}

func AddMenuItem(server string, name string, price string) string {
  // mtx.Lock()
  // defer mtx.Unlock()

  // f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  // defer f.Close()
  //
  // if err != nil {
  //   log.Fatalf(err.Error())
  // }
  //
  // menuitem := fmt.Sprintf("%v, %v, %v\n", contents, cooking_style, price)
  //
  // if _, err := f.WriteString(menuitem); err != nil {
  //   log.Fatalf(err.Error())
  // }
  getMenuFromFile()

  menuitem := Dish {SKU: getSKU(), Name: name, Price: price}

  Menu = append(Menu, menuitem)

  writeMenuToFile()

  action := fmt.Sprintf("Added dish %v", menuitem)

  poslog.Log(server, action)

  return action
}

func DeleteMenuItem(server string, sku int) string {
  // mtx.Lock()
  // defer mtx.Unlock()

  getMenuFromFile()

  menuitem := DishFromSKU(sku)

  Menu = append(Menu[:sku-1], Menu[sku:]...)

  writeMenuToFile()

  action := fmt.Sprintf("Deleted dish %v", menuitem)

  poslog.Log(server, action)

  return action
}

func EditMenuItem(server string, sku int, name string, price string) string {
  // mtx.Lock()
  // defer mtx.Unlock()

  getMenuFromFile()

  menuitem := DishFromSKU(sku)

  dish := &Menu[sku-1]

  if name != "" {
    dish.Name = name
  }
  if price != "" {
    dish.Price = price
  }

  editedmenuitem := DishFromSKU(sku)

  writeMenuToFile()

  action := fmt.Sprintf("Edited dish %v to be %v", menuitem, editedmenuitem)

  poslog.Log(server, action)

  return action
}
