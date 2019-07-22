package transaction

import (
  // "fmt"
  "time"
  "strconv"

  "github.com/rs/xid"
  "github.com/svasandani/pos-app/menu"
)

type Transaction struct {
  ID string `json:"id"`
  Date string `json:"date"`
  Payment_method string `json:"payment_method"`
  SKUList []int `json:"skulist"`
  CompletedList []bool `json:"completedlist"`
}

func NewTransaction() Transaction {
  return Transaction {ID: xid.New().String(), Date: time.Now().Format("Jan-02-2006 3:04:05 PM"), Payment_method: "", SKUList: []int{}, CompletedList: []bool{}}
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
}

func UnserveDish(transaction *Transaction, index int) {
  setCompletedStatus(transaction, index, false)
}

func GetTotal(transaction *Transaction) int {
  price := 0

  for _, sku := range(transaction.SKUList) {
    dishprice, err := strconv.ParseInt(menu.Menu[sku-1].Price, 10, 0)

    if err != nil {
      panic(err)
    }

    price += int(dishprice)
  }

  return price
}
