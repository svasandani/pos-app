package server

import (
  "fmt"
  "time"
  // "strconv"
  "io/ioutil"
  "log"
  "strings"
  "sync"

  "github.com/rs/xid"
  poslog "github.com/svasandani/pos-app/log"
)

type Server struct {
  ID string `json:"id"`
  Name string `json:"name"`
  Date string `json:"date"`
}

var serverList []Server
var filename string = "./server/servers.csv"
var mtx sync.RWMutex

func getServersFromFile() {
  // mtx.Lock()
  // defer mtx.Unlock()

  serverList = nil

  content, err := ioutil.ReadFile(filename)

  if err != nil {
    log.Fatalf(err.Error())
  }

  result := strings.Split(string(content), "\n")

  for _, server := range(result) {
    item := strings.Split(server, ", ")
      if len(item) < 2 {continue;}
    serverList = append(serverList, Server {ID: item[0], Name: item[1], Date: item[2]})
  }
}

func writeServersToFile() {
  // mtx.Lock()
  // defer mtx.Unlock()

  output := ""

  for _, item := range(serverList) {
    output += fmt.Sprintf("%v, %v, %v\n", item.ID, item.Name, item.Date)
  }

  bytes := []byte(output)

  ioutil.WriteFile(filename, bytes, 0644)
}

func findServerByID(id string) int {
  for index, item := range(serverList) {
    if item.ID == id {
      return index
    }
  }

  return -1
}

func FindServerByName(name string) Server {
  for _, item := range(serverList) {
    if item.Name == name {
      return item
    }
  }

  return Server {ID: "", Name: "", Date: ""}
}

func AddNewServer(name string) string {
  // mtx.Lock()
  // defer mtx.Unlock()

  datestring := time.Now().Format("060201")
  id := xid.New().String()

  getServersFromFile()

  serverList = append(serverList, Server {ID: id, Name: name, Date: datestring})

  writeServersToFile()

  action := fmt.Sprintf("Added new server %v", name)
  poslog.Log("nil", action)

  return id
}

func RemoveServer(id string) string {
  // mtx.Lock()
  // defer mtx.Unlock()

  getServersFromFile()

  index := findServerByID(id)

  name := serverList[index].Name

  serverList = append(serverList[:index], serverList[index+1:]...)

  writeServersToFile()

  action := fmt.Sprintf("Deleted server %v", name)
  poslog.Log("nil", action)

  return name
}

func CheckServer(id string) string {
  getServersFromFile()

  for _, item := range(serverList) {
    if item.ID == id {
      return item.Name
    }
  }

  return ""
}
