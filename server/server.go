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
)

type Server struct {
  ID string
  Name string
  Date string
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
    output += fmt.Sprintf("%v, %v, %v", item.ID, item.Name, item.Date)
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

func AddNewServer(name string) {
  // mtx.Lock()
  // defer mtx.Unlock()

  datestring := time.Now().Format("060201")

  getServersFromFile()

  serverList = append(serverList, Server {ID: xid.New().String(), Name: name, Date: datestring})

  writeServersToFile()
}

func RemoveServer(id string) {
  // mtx.Lock()
  // defer mtx.Unlock()

  getServersFromFile()

  index := findServerByID(id)

  serverList = append(serverList[:index], serverList[index+1:]...)

  writeServersToFile()
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
