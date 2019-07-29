package main

import (
  "github.com/svasandani/pos-app/handlers"

  "github.com/codegangsta/negroni"
  gmux "github.com/gorilla/mux"
)

func main() {
  // var err error
  mux := gmux.NewRouter()

  mux.HandleFunc("/login", handlers.LoginHandler)

  mux.HandleFunc("/menu", handlers.MenuGetHandler).Methods("GET")
  mux.HandleFunc("/menu", handlers.MenuAddHandler).Methods("POST")
  mux.HandleFunc("/menu", handlers.MenuDeleteHandler).Methods("DELETE")
  mux.HandleFunc("/menu", handlers.MenuEditHandler).Methods("PUT")

  mux.HandleFunc("/transaction", handlers.TransactionGetAllHandler).Methods("GET")
  mux.HandleFunc("/transaction/new", handlers.TransactionNewHandler).Methods("GET")
  mux.HandleFunc("/transaction", handlers.TransactionAddDishHandler).Methods("POST")
  mux.HandleFunc("/transaction", handlers.TransactionDeleteDishHandler).Methods("DELETE")
  mux.HandleFunc("/transaction", handlers.TransactionServerHandler).Methods("PUT")
  mux.HandleFunc("/transaction/pay", handlers.TransactionPayHandler).Methods("POST")
  //
  mux.HandleFunc("/server", handlers.ServerAddHandler).Methods("POST")
  mux.HandleFunc("/server", handlers.ServerDeleteHandler).Methods("DELETE")

  n := negroni.Classic()
  n.UseHandler(mux)
  n.Run(":2000")
}

// func main() {
//   menufile := menu.GetMenu()
//   fmt.Println(menufile)
//
//   server := server.Server {ID: "iohsdofhowiegoi", Name: "Aileen", Date: "",}
//
//   menu.AddMenuItem(server, "chicken masala", "250")
//   // menu.DeleteMenuItem(server, 76)
//   menufile = menu.GetMenu()
//   // fmt.Println(menufile)
//   //
//   menu.EditMenuItem(server, 76, "", "350")
//   // // menu.AddMenuItem("chicken", "masala", "250")
//   // menufile = menu.GetMenu()
//   // fmt.Println(menufile)
//   // fmt.Println(menu.Menu)
//
//   // test := transaction.NewTransaction()
//   //
//   // transaction.AddDish(&test, 23)
//   // transaction.AddDish(&test, 51)
//   // transaction.AddDish(&test, 37)
//   // transaction.AddDish(&test, 57)
//   // transaction.AddDish(&test, 22)
//   // transaction.AddDish(&test, 13)
//   // transaction.AddDish(&test, 45)
//   // transaction.AddDish(&test, 57)
//   //
//   // transaction.ServeDish(&test, 0)
//   // transaction.ServeDish(&test, 1)
//   // transaction.ServeDish(&test, 2)
//   // transaction.ServeDish(&test, 3)
//   // transaction.ServeDish(&test, 4)
//   // transaction.ServeDish(&test, 5)
//   // transaction.ServeDish(&test, 6)
//   // transaction.ServeDish(&test, 7)
//   //
//   // fmt.Println(test)
//   // fmt.Println(test.Total)
//   //
//   // transaction.Pay(&test, "cash")
//
//   // transaction.GetSalesReport()
//
// }
