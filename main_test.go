package main

import (
	"testing"
	"time"
)

func mockTime(year int , month time.Month , day int) time.Time{
  return time.Date(year,month,day,0,0,0,0,time.UTC)
}

func TestGetCurrentQuarter(t *testing.T){
  tests :=[]struct{
    name string
    mockDate func() time.Time
    expectedQ int
    expectedQS string
  }{
    {"Q1 of 2023" , func() time.Time {return mockTime(2023,1,5)},1,"Q1-2023"},
    {"Q2 of 2023" , func() time.Time {return mockTime(2023,5,1)},2,"Q2-2023"},
    {"Q3 of 2023" , func() time.Time {return mockTime(2023,7,20)},3,"Q3-2023"},
    {"Q4 of 2023" , func() time.Time {return mockTime(2023,11,3)},4,"Q4-2023"},
  }

  // Loop through each test case
  for _,tc := range tests{
    t.Run(tc.name , func(t *testing.T){
      actualQ , actualQS := GetCurrentQuarter(tc.mockDate)

      if actualQ != tc.expectedQ {
        t.Errorf("Expected quarter %d, got %d" , tc.expectedQ , actualQ)
      }
      if actualQS != tc.expectedQS{
        t.Errorf("Expected quarter %s, got %s" , tc.expectedQS , actualQS)
      }  
      
    })
  }
}