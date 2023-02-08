package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
		"http://golang.org",
	}
	c := make(chan string)

	//Case 1: Making request to each link one time

	for _, link := range links {
		// checkLink(link)
		go checkLink(link, c) //go routine ->always placed in front of func
		//The above line will not execute suceessfully because the main routine would exit before child routine completes their work
		//routines are used to handle blocking code
		//channels are used to coomunicate between running routines
		//The information that we pass into a channel or the data that we attempt to share between these different routines must all be of the same type.

	}
	//receiving a msg over a channel is a blocking operation
	// fmt.Println(<-c) if we place only a single print channel then only one child routine will run

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	//Case 2: Making request to each link multiple times

	//Note: Variable cannot be directly shared between routines. They must be passed as parameters
	for l := range c {
		//*********func literal ->  equivalent to anonymous func in JS
		//time.sleep should be written inside func literal otherwise throttling will take place
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

// use channel of  string as second para
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down !")
		// c <- "Might be down I think"
		c <- link
		return
	}
	fmt.Println(link, "is up !")
	c <- link
	// c <- "is up"
}

//While printing we observe a slight delay between each statements
//The rder of printing is same as the order in which our code is written
//We make request to first link from slice and wait for its response then make the next request
//This follows a serial order but it will consume lot of time when number og links increases

//We can take a parallel approach where we attempt to make request to all the links right away and then whenever we get any response print that status
//We will use Go routines and channels for this

//when we launch a go program, like when we compile it and execute it, we automatically create one go routine. This main routine executes our code line by line
//We can create additional go routine to launch new function that we want to be executed inside of its individual go routine

//In case of single CPU, Scheduler runs one routine until it finishes or makes a blocking call like HTTP reuest
//By default Go uses 1 CPU Core
//In case of multiple CPUs cores each one can run one go routine

//Concurrency: we don't have to wait. Even in case of single CPUs when one routine gets blocked another routine can run immediately
//Parallelism: Can be achieved only through multiple CPUs
