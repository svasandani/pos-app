package transaction

import (
  // "fmt"
  "time"

  "github.com/rs/xid"
)

type Transaction struct {
  ID string `json:"id"`
  Date string `json:"date"`
  Payment_method string `json:"payment_method"`
  SKUList []int `json:"skulist"`
}

func NewTransaction() Transaction {
  return Transaction {ID: xid.New().String(), Date: time.Now().Format("Jan-02-2006 3:04:05 PM"), Payment_method: "", SKUList: []int{}}
}
