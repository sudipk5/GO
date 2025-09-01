// package main

// import (
// 	"fmt"
// 	"time"
// )

// func processNum(numchan chan int) {
// 	fmt.Println("Processing number ", <- numchan)

// }

// func main() {

// 	numchan := make(chan int)

// 	go processNum(numchan)

// 	numchan <- 5

// 	time.Sleep(time.Second * 2)

// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func sum (v chan int) {

// 	fmt.Println( <- v)
// }

// func main() {

// 	result := make(chan int)

// 	go sum(result)

// 	result <- 10

// 	time.Sleep(time.Second*2)

// }

// package main

// import "fmt"

// func sum(result chan int, num int, num2 int) {

// 	resul := num + num2

// 	result <- resul

// }

// func main() {

// 	summ := make(chan int)
// 	go sum(summ, 10, 20)
// 	res:=<-summ
// 	fmt.Println(res)

// }

// package main

// import "fmt"

// func task(done chan bool) {

// 	defer func () { done <- true }()

// 	fmt.Println("processimg...")

// }

// func main() {

// 	donr:=make(chan bool)

// 	go task(donr)

// 	 <-donr

// }


package main

import (
	"fmt"
	"time"
)


func Emailsend(emailchan chan string , done chan bool){
	defer func ()  {
		done <- true
	}()

	for email:= range emailchan{

		fmt.Println("sending email to", email)

		time.Sleep(time.Second)
	}

}

func main() {

	emailchan := make(chan string, 100)

	done := make(chan bool)

	go Emailsend(emailchan,done)


	for i :=0;i<5;i++{
		emailchan <- fmt.Sprintf("%d@gmail.com",i)
	}

	fmt.Println("done sending. ")
	close(emailchan)
	<- done

	// emailchan <- "sudip@89"
	// emailchan <- "sudip@ku89"

	// fmt.Println(<-emailchan)

	// fmt.Println(<-emailchan)

	

}
