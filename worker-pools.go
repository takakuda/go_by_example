package main

import(
  "fmt"
  "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
  for j := range jobs {
    fmt.Println("worker", id, "started job", j)
    time.Sleep(tine.Second)
    fmt.Println("worker", id, "finished job", j)
    results <- j * 2
  }
}
