package handlers

import (
  "net/http"
  "encoding/json"
  "io"
  "io/ioutil"

  "github.com/svasandani/pos-app/menu"
  "github.com/svasandani/pos-app/server"
  // "github.com/svasandani/pos-app/transaction"
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
