package handlers

import (
  "net/http"
  "encoding/json"
  "io"
  "io/ioutil"
  "strconv"

  "github.com/svasandani/pos-app/menu"
  "github.com/svasandani/pos-app/server"
  "github.com/svasandani/pos-app/transaction"
)

func setupResponse(w *http.ResponseWriter, r *http.Request) {
	  (*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
  setupResponse(&w, r)

  servername := r.FormValue("user")
  serverpass := r.FormValue("id")

  if server.CheckServer(serverpass) != servername {
    http.Error(w, "Login failed!", http.StatusUnauthorized)
  }

  encoder := json.NewEncoder(w)

  if err := encoder.Encode(serverpass); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func MenuGetHandler(w http.ResponseWriter, r *http.Request) {
  setupResponse(&w, r)

  encoder := json.NewEncoder(w)

  if err := encoder.Encode(menu.GetMenu()); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func MenuAddHandler(w http.ResponseWriter, r *http.Request) {
  setupResponse(&w, r)
  defer r.Body.Close()

  serverid := r.FormValue("server")

  servername := server.CheckServer(serverid)

  if servername == "" {
    http.Error(w, "Server not found", http.StatusUnauthorized)
    return
  }

  dish, err := convertHTTPToDish(r.Body)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  encoder := json.NewEncoder(w)

  if err := encoder.Encode(menu.AddMenuItem(servername, dish.Name, dish.Price)); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func MenuDeleteHandler(w http.ResponseWriter, r *http.Request) {
  setupResponse(&w, r)
  defer r.Body.Close()

  serverid := r.FormValue("server")

  servername := server.CheckServer(serverid)

  if servername == "" {
    http.Error(w, "Server not found", http.StatusUnauthorized)
    return
  }

  dish, err := convertHTTPToDish(r.Body)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  encoder := json.NewEncoder(w)

  if err := encoder.Encode(menu.DeleteMenuItem(servername, dish.SKU)); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func MenuEditHandler(w http.ResponseWriter, r *http.Request) {
  setupResponse(&w, r)
  defer r.Body.Close()

  serverid := r.FormValue("server")

  servername := server.CheckServer(serverid)

  if servername == "" {
    http.Error(w, "Server not found", http.StatusUnauthorized)
    return
  }

  dish, err := convertHTTPToDish(r.Body)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  encoder := json.NewEncoder(w)

  if err := encoder.Encode(menu.EditMenuItem(servername, dish.SKU, dish.Name, dish.Price)); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func convertHTTPToDish(httpBody io.ReadCloser) (menu.Dish, error) {
  body, err := ioutil.ReadAll(httpBody)

  if err != nil {
    return menu.Dish{}, err
  }

  defer httpBody.Close()
  return convertJSONToDish(body)
}

func convertJSONToDish(jsonBody []byte) (menu.Dish, error) {
  var dish menu.Dish
  err := json.Unmarshal(jsonBody, &dish)

  if err != nil {
    return menu.Dish{}, err
  }

  return dish, nil
}

func ServerAddHandler(w http.ResponseWriter, r *http.Request) {
  setupResponse(&w, r)
  defer r.Body.Close()

  serverid := r.FormValue("server")

  servername := server.CheckServer(serverid)

  if servername == "" {
    http.Error(w, "Server not found", http.StatusUnauthorized)
    return
  }

  serverobject, err := convertHTTPToServer(r.Body)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  encoder := json.NewEncoder(w)

  if err := encoder.Encode(server.AddNewServer(servername, serverobject.Name)); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func ServerDeleteHandler(w http.ResponseWriter, r *http.Request) {
  setupResponse(&w, r)
  defer r.Body.Close()

  serverid := r.FormValue("server")

  servername := server.CheckServer(serverid)

  if servername == "" {
    http.Error(w, "Server not found", http.StatusUnauthorized)
    return
  }

  serverobject, err := convertHTTPToServer(r.Body)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  encoder := json.NewEncoder(w)

  if err := encoder.Encode(server.RemoveServer(servername, serverobject.ID)); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func convertHTTPToServer(httpBody io.ReadCloser) (server.Server, error) {
  body, err := ioutil.ReadAll(httpBody)

  if err != nil {
    return server.Server{}, err
  }

  defer httpBody.Close()
  return convertJSONToServer(body)
}

func convertJSONToServer(jsonBody []byte) (server.Server, error) {
  var serveritem server.Server
  err := json.Unmarshal(jsonBody, &serveritem)

  if err != nil {
    return server.Server{}, err
  }

  return serveritem, nil
}

func TransactionGetAllHandler(w http.ResponseWriter, r *http.Request) {
  setupResponse(&w, r)

  encoder := json.NewEncoder(w)

  if err := encoder.Encode(transaction.GetAll()); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func TransactionNewHandler(w http.ResponseWriter, r *http.Request) {
  setupResponse(&w, r)

  serverid := r.FormValue("server")

  servername := server.CheckServer(serverid)

  if servername == "" {
    http.Error(w, "Server not found", http.StatusUnauthorized)
    return
  }

  encoder := json.NewEncoder(w)

  if err := encoder.Encode(transaction.NewTransaction(servername)); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func TransactionAddDishHandler(w http.ResponseWriter, r *http.Request) {
  setupResponse(&w, r)

  id := r.FormValue("id")
  s, _ := strconv.ParseInt(r.FormValue("sku"), 10, 0)
  sku := int(s)

  encoder := json.NewEncoder(w)

  if err := encoder.Encode(transaction.AddDish(id, sku)); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func TransactionDeleteDishHandler(w http.ResponseWriter, r *http.Request) {
  setupResponse(&w, r)

  id := r.FormValue("id")
  i, _ := strconv.ParseInt(r.FormValue("index"), 10, 0)
  index := int(i)

  encoder := json.NewEncoder(w)

  if err := encoder.Encode(transaction.DeleteDish(id, index)); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func TransactionServerHandler(w http.ResponseWriter, r *http.Request) {
  setupResponse(&w, r)

  id := r.FormValue("id")
  i, _ := strconv.ParseInt(r.FormValue("index"), 10, 0)
  index := int(i)

  encoder := json.NewEncoder(w)

  if err := encoder.Encode(transaction.ToggleServe(id, index)); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func TransactionPayHandler(w http.ResponseWriter, r *http.Request) {
  setupResponse(&w, r)

  id := r.FormValue("id")
  method := r.FormValue("method")

  encoder := json.NewEncoder(w)

  if err := encoder.Encode(transaction.Pay(id, method)); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func convertHTTPToTransaction(httpBody io.ReadCloser) (transaction.Transaction, error) {
  body, err := ioutil.ReadAll(httpBody)

  if err != nil {
    return transaction.Transaction{}, err
  }

  defer httpBody.Close()
  return convertJSONToTransaction(body)
}

func convertJSONToTransaction(jsonBody []byte) (transaction.Transaction, error) {
  var transactionitem transaction.Transaction
  err := json.Unmarshal(jsonBody, &transactionitem)

  if err != nil {
    return transaction.Transaction{}, err
  }

  return transactionitem, nil
}
