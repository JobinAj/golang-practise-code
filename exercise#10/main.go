package main

import (
     "fmt"
     "time"
    )

func main(){
	Current_time:=time.Now()
	fmt.Printf("hello world the time is %v\n Thanks",Current_time.Format("2006/01/02 15:04:05"))
}

