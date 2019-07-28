package log

import (
  "fmt"
  "time"
  // "strconv"
  "os"
  "io"
)

func Log(servername string, action string) {
  datestring := time.Now().Format("060201")
  filename := "./log_history/" + datestring + ".log"

  file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

  if err != nil {
    panic(err)
  }

  output := fmt.Sprintf("Server: %v, Date: %v, Action: %v\n", servername, datestring, action)

  _, err = io.WriteString(file, output)

  if err != nil {
    panic(err)
  }
}
