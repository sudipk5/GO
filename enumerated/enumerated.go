package main

import "fmt"

type orderStatus string

const (
	Received orderStatus = "received"

	confirmed = "confirmed"

	prepared = "prepared"

	delivered = "delivered"
)

func chaneOrderStatus(status orderStatus) {

	fmt.Println("changing order status to", status)

}

// func chaneOrderStatu(status string) {

// 	fmt.Println("changing order status to", status)

// }

func main() {

	chaneOrderStatus(delivered)
	chaneOrderStatus(prepared)
	chaneOrderStatus(confirmed)
	// chaneOrderStatu("deliverd")
	// chaneOrderStatu("prepared")
	// chaneOrderStatu("confirmed")

}
