package main

import (
	"fmt"
	"time"
)
func GetCurrentQuarter(now func() time.Time)(int,string){
  current := now()
  q := ((int(current.Month()) - 1)/3 + 1)
  qs := fmt.Sprintf("Q%d-%d" ,q,current.Year())
  return q, qs
}
func main() {
	q, qs := GetCurrentQuarter(time.Now)
  fmt.Println("Current quarter", q , "Formatted : " , qs)
}

