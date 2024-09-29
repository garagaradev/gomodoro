package gomodoro
import (
  "fmt"
  "time"
  "github.com/gen2brain/beeep"
)

type Gomodoro struct {
  WorkDuration time.Duration
  ShortBreak time.Duration
  LongBreak time.Duration
  CyclesBeforeLongBreak int
  Cycles int
}

func NewGomodoro() *Gomodoro {
  return &Gomodoro {
    WorkDuration: 25 * time.Minute,
    ShortBreak: 5 * time.Minute,
    LongBreak: 15 * time.Minute,
    CyclesBeforeLongBreak: 4,
    Cycles: 0,
  }
}

func (g *Gomodoro) StartGomodoro(){
  fmt.Println("Starting a new session...")

  for g.Cycles < g.CyclesBeforeLongBreak {
    fmt.Println("Working Session is running...")
    g.timer(g.WorkDuration,"Time to take a rest...")

    g.Cycles++

    if g.Cycles == g.CyclesBeforeLongBreak {
      fmt.Println("Time to take a long long rest...Yeehaaaw!")
      g.timer(g.LongBreak,"Long rest session ends.")
      g.Cycles = 0 //reset
    }else {
      fmt.Println("Short rest")
      g.timer(g.ShortBreak,"Short rest session ends.")
    }
  }
}

func (g *Gomodoro) timer(duration time.Duration, message string){
  fmt.Printf("Start timer for %v...\n",duration)
  time.Sleep(duration)
  
  //alarm 
  beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
  //notification
  err := beeep.Notify("Gomodoro Timer:",message, "")
  if err != nil {
    fmt.Println("Error displaying notifications:",err)
  }
}
