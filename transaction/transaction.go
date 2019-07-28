package transaction

import (
  "fmt"
  "time"
  // "strconv"
  "io/ioutil"
  "os"
  "path/filepath"
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

var transactionList []Transaction

func transactionFromID(id string) Transaction {
  for _, item := range(transactionList) {
    if item.ID == id {
      return item
    }
  }

  return Transaction {ID: "", Date: "", Server: "", Payment_method: "", SKUList: []int{}, CompletedList: []bool{}, Total: 0}
}

func writeTransactionToDisk(transactionid string) {
  transaction := transactionFromID(transactionid)

  datestring := time.Now().Format("060201")
  newpath := filepath.Join("./transaction_history/", datestring)
  os.MkdirAll(newpath, os.ModePerm)
  filename := "./transaction_history/" + datestring + "/" + transaction.ID + ".transaction"

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

// func NewTransaction() Transaction {
//   return Transaction {ID: xid.New().String(), Date: time.Now().Format("Jan-02-2006 3:04:05 PM"), Server: "", Payment_method: "", SKUList: []int{}, CompletedList: []bool{}, Total: 0}
// }

func NewTransaction() string {
  transactionList = append(transactionList, Transaction {ID: xid.New().String(), Date: time.Now().Format("Jan-02-2006 3:04:05 PM"), Server: servername, Payment_method: "", SKUList: []int{}, CompletedList: []bool{}, Total: 0})
}

func AddDish(transactionid string, sku int) {
  transaction := transactionFromID(transactionid)

  transaction.SKUList = append(transaction.SKUList, sku)
  transaction.CompletedList = append(transaction.CompletedList, false)
}

func DeleteDish(transactionid string, index int) {
  transaction := transactionFromID(transactionid)

  transaction.SKUList = append(transaction.SKUList[:index], transaction.SKUList[index-1:]...)
  transaction.CompletedList = append(transaction.CompletedList[:index], transaction.CompletedList[index-1:]...)
}

func setCompletedStatus(transactionid string, index int, value bool) {
  transaction := transactionFromID(transactionid)

  transaction.CompletedList[index] = value;
}

func ToggleServe(transactionid string, index int) {
  transaction := transactionFromID(transactionid)

  if transaction.CompletedList[index] {
    UnserveDish(transactionid, index)
  } else {
    ServeDish(transactionid, index)
  }
}

func ServeDish(transactionid string, index int) {
  transaction := transactionFromID(transactionid)

  setCompletedStatus(transaction, index, true)
  dishprice := menu.PriceFromSKU(transaction.SKUList[index])

  transaction.Total += dishprice
}

func UnserveDish(transactionid string, index int) {
  transaction := transactionFromID(transactionid)

  setCompletedStatus(transaction, index, false)
  dishprice := menu.PriceFromSKU(transaction.SKUList[index])

  transaction.Total -= dishprice
}

func Pay(transactionid string, method string) Transaction {
  transaction := transactionFromID(transactionid)

  if strings.ToLower(method) != "card" && strings.ToLower(method) != "cash" {
    return
  }

  transaction.Payment_method = method

  writeTransactionToDisk(transaction)

  return transaction
}

func GetSalesReport() {
  datestring := time.Now().Format("060201")
  root, _ := filepath.Abs("./transaction_history/" + datestring + "/")

  transactionmap := make(map[string]int)

  isroot := true
  numtransactions := 0

  err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
    if isroot {
      isroot = false
      return nil
    }

    numtransactions++

    content, _ := ioutil.ReadFile(path)
    blocks := strings.Split(string(content), "\n\n\n")
    results := strings.Split(blocks[0], "\n")

    for _, string := range(results) {
      string = strings.TrimSpace(string)
    }
    results = results[4:]

    for _, string := range(results) {
      transactionmap[strings.Split(string, " ")[0]]++
    }

    // var transactions []string
    // for _, string := range(results) {
    //   transactions = append(transactions, strings.Split(string, " ")[0])
    // }
    //
    // for _, sku := range(transactions) {
    //   transactionmap[sku]++
    // }

    return nil
  })

  fmt.Println(transactionmap)

  if err != nil {
    panic(err)
  }

}
