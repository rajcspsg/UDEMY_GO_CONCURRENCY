package main

import (
	"fmt"
	"sync"
	"time"
)



var resource1 bool
var resource2 bool
var resource3 bool


func serviceA(wg *sync.WaitGroup) {
	for range time.Tick(time.Second * 2) {
		fmt.Println("I am service A")
		
		if(resource1 == true) { resource2 = true; wg.Done() } /* else {
			resource1 = true;
		} //the fix */
	}	
}


func serviceB(wg *sync.WaitGroup) {
	for range time.Tick(time.Second * 2) {
		fmt.Println("I am service B")
		
		if(resource2 == true) { resource3 = true; wg.Done() }
		
	}
}


func serviceC(wg *sync.WaitGroup) {
	for range time.Tick(time.Second * 2) {
		fmt.Println("I am service C")
		if(resource3 == true) { resource1 = true; wg.Done() }
	}
}

var round int64
func watchdog(wg *sync.WaitGroup) {
	//Kill The Go Funcs reset the syste
	//if we end up in the same situation were in a livelock....		
}


var everythingDone bool
func mainProcess(wg *sync.WaitGroup) {	
	resource1 = false;	
	resource2 = false;
	resource3 = false;
	
	wg.Add(3)
	go serviceA(wg)
	go serviceB(wg)
	go serviceC(wg)
	
	wg.Wait()
	if resource1 && resource2 && resource3 {
		everythingDone = true
	}	
}

func main() {
	round = 1;
	var wg sync.WaitGroup
	
	go watchdog(&wg)
	go mainProcess(&wg)	
	
	for !everythingDone {
		time.Sleep(5 * time.Second) 
		fmt.Println("waiting...")
	}
	fmt.Println("Finally")
}	
