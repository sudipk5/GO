package main

import "fmt"

type paymenter interface{
	pay(amount float32) 
}
 

func (p payment) makepayment(amount float32) {

	p.gateway.pay(amount)

}

type payment struct{
	gateway paymenter
	
}

type spay struct{}
func (s spay) pay(amount float32) {
	
	fmt.Print("making payment spay", amount)
}


func main() {
	spay:=spay{}
	newpayment:=payment{
		gateway: spay,
		
	}
	newpayment.makepayment(100)
}
