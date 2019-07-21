package menu

import (
  "fmt"
  // "io"
  "io/ioutil"
  "log"
  "os"
  "strings"
  "strconv"
)

var menu [][]string
var sku int = 0;
var filename string = "./menu_output.txt"

func getMenuFromFile() {
  sku = 0

  content, err := ioutil.ReadFile(filename)

  if err != nil {
    log.Fatalf(err.Error())
  }

  result := strings.Split(string(content), "\n")

  for _, item := range(result) {
    menuitem := strings.Split(item, ",")
    menuitem = append([]string{intToString(getSKU())}, menuitem...)
    if len(menuitem) > 2 {
      menu = append(menu, menuitem)
    }
  }
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
  f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  defer f.Close()

  if err != nil {
    log.Fatalf(err.Error())
  }

  menuitem := fmt.Sprintf("%v, %v, %v\n", contents, cooking_style, price)

  if _, err := f.WriteString(menuitem); err != nil {
    log.Fatalf(err.Error())
  }
}
