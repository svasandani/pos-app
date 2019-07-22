package menu

import (
  "fmt"
  // "io"
  "io/ioutil"
  "log"
  // "os"
  "strings"
  "strconv"
)

var menu [][]string
var sku int = 0;
var filename string = "./menu_output.txt"

func getMenuFromFile() {
  sku = 1
  menu = nil

  content, err := ioutil.ReadFile(filename)

  if err != nil {
    log.Fatalf(err.Error())
  }

  result := strings.Split(string(content), "\n")

  for _, item := range(result) {
    menuitem := strings.Split(item, ",")
    if len(menuitem) < 1 {continue;}
    for i, string := range(menuitem) {
      menuitem[i] = strings.TrimSpace(string)
    }
    menuitem = append([]string{intToString(getSKU())}, menuitem...)
    if len(menuitem) > 2 {
      menu = append(menu, menuitem)
    }
  }
}

func writeMenuToFile() {
  output := ""

  for _, item := range(menu) {
    output += fmt.Sprintf("%v, %v, %v\n", item[1], item[2], item[3])
  }

  bytes := []byte(output)

  ioutil.WriteFile(filename, bytes, 0644)
}

func intToString(int int) string {
  return strconv.FormatInt(int64(int), 10)
}

func getSKU() int {
  sku++
  return sku
}

func GetMenu() [][]string {
  getMenuFromFile()
  return menu
}

func AddMenuItem(contents string, cooking_style string, price string) {
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

  menuitem := []string{intToString(getSKU()), contents, cooking_style, price}

  menu = append(menu, menuitem)

  writeMenuToFile()
}

func DeleteMenuItem(sku int) {
  getMenuFromFile()

  menu = append(menu[:sku-1], menu[sku:]...)

  writeMenuToFile()
}

func EditMenuItem(sku int, contents string, cooking_style string, price string) {
  getMenuFromFile()

  menuitem := menu[sku-1]

  if contents != "" {
    menuitem[1] = contents
  }
  if cooking_style != "" {
    menuitem[2] = cooking_style
  }
  if price != "" {
    menuitem[3] = price
  }

  writeMenuToFile()
}
