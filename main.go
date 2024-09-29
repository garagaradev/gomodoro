package main
import (
  "fmt"
  "os"
  "strconv"
  "github.com/garagaradev/gomodoro/gomodoro")

func main(){
  if len(os.Args) < 2 {
    fmt.Println("Please provide work duration in minutes.")
    return
  }
  workDurationInput := os.Args[1]
  workDurationMinutes, err := strconv.Atoi(workDurationInput)
  if err != nil {
    fmt.Println("Invalid duration:", err)
    return
  }
  g := gomodoro.NewGomodoro(workDurationMinutes)
  g.StartGomodoro()
}
