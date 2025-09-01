package main

import (
	"fmt"
	"time"
)

type post struct {
	views int
}

func (p *post) inc() {
	p.views += 1
}

func main() {

	mypost := post{
		views: 0,
	}

	for i := 1; i <= 100; i++ {

		go mypost.inc()

	}
	time.Sleep(time.Second * 2)

	fmt.Println(mypost.views)

}
