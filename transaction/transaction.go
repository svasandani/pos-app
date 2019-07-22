package transaction

import (
  "fmt"
  "time"
  // "strconv"
  "os"
  "io"
  "strings"

  "github.com/rs/xid"
  "github.com/svasandani/pos-app/menu"
)

type Transaction struct {
  ID string `json:"id"`
  Date string `json:"date"`
  Server string `json:"server"`
  Payment_method string `json:"payment_method"`
  SKUList []int `json:"skulist"`
  CompletedList []bool `json:"completedlist"`
  Total int `json:"total"`
  // Table int `json:"table"`
}

func writeTransactionToDisk(transaction *Transaction) {
  filename := "./transaction_history/" + transaction.ID + ".transaction"

  file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)

  if err != nil {
    panic(err)
  }

  output := fmt.Sprintf("Transaction ID: %v\nDate: %v\nServed by: %v\nPayment Method: %v\n", transaction.ID, transaction.Date, transaction.Server, transaction.Payment_method)

  for index, truth := range(transaction.CompletedList) {
    if truth {
      output += fmt.Sprintf("%-3v%-24v%8v\n", transaction.SKUList[index], menu.DishFromSKU(transaction.SKUList[index]).Name, menu.PriceFromSKU(transaction.SKUList[index]))
    }
  }

  output += fmt.Sprintf("\n\n\nTotal:%29v", transaction.Total)

  _, err = io.WriteString(file, output)

  if err != nil {
    panic(err)
  }
}

func NewTransaction() Transaction {
  return Transaction {ID: xid.New().String(), Date: time.Now().Format("Jan-02-2006 3:04:05 PM"), Server: "", Payment_method: "", SKUList: []int{}, CompletedList: []bool{}, Total: 0}
}

func AddDish(transaction *Transaction, sku int) {
  transaction.SKUList = append(transaction.SKUList, sku)
  transaction.CompletedList = append(transaction.CompletedList, false)
}

func DeleteDish(transaction *Transaction, index int) {
  transaction.SKUList = append(transaction.SKUList[:index], transaction.SKUList[index-1:]...)
  transaction.CompletedList = append(transaction.CompletedList[:index], transaction.CompletedList[index-1:]...)
}

func setCompletedStatus(transaction *Transaction, index int, value bool) {
  transaction.CompletedList[index] = value;
}

func ServeDish(transaction *Transaction, index int) {
  setCompletedStatus(transaction, index, true)
  dishprice := menu.PriceFromSKU(transaction.SKUList[index])

  transaction.Total += dishprice
}

func UnserveDish(transaction *Transaction, index int) {
  setCompletedStatus(transaction, index, false)
  dishprice := menu.PriceFromSKU(transaction.SKUList[index])

  transaction.Total -= dishprice
}

func Pay(transaction *Transaction, method string) {
  if strings.ToLower(method) != "card" && strings.ToLower(method) != "cash" {
    return
  }

  transaction.Payment_method = method

  writeTransactionToDisk(transaction)
}
