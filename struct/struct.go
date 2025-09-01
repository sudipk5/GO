// package main

// import "fmt"

// type ab struct{
// 	id  string
// 	name string
// 	roll int
// 	hight float32
// }

// func (a *ab) getid(b string){
//   a.name = b
// }
// func main(){
//  ad:=ab{
// 	id:"hlw",
// 	name : "sudip",
// 	roll: 125,
// 	hight: 5.8,
//  }

//  ad.getid("kundu")
//  b:=ab{

// 	id :"lol",
// 	name:"sourav",
// 	roll:97,
// 	hight: 5.6,
//  }
//   fmt.Println(ad)
//   fmt.Println(b)
// }

// package main

// import "fmt"

// type order2 struct {
// 	book string
// }
// type order struct {
// 	name string
// 	roll int
// 	age  int
// 	order2
// }

// func main() {

// 	b := order2{
// 		book: "harray",
// 	}

// 	a := order{
// 		name:   "sudip",
// 		roll:   125,
// 		age:    23,
// 		order2: b,
// 	}

// 	fmt.Println(a)

// }



// package main
// import "fmt"

// type order struct{
// 	id string
// 	amount float32
// 	status string
	
// }

// func (o *order) changeStatus(s string){
// 		o.status=s
// }

//  func (o *order) GetAmount() float32{

// 	return o.amount

//  }

//  func(o *order) Rp(id string) string{
// 	o.id=id
// 	return id

//  }


// type Adamas struct{
// 	Name string
// 	Roll int
// 	Age int
// 	Program string
// }


// func main(){

// 	sudip_order:=order{
// 		id:"hlw09",
// 		amount: 900,
// 		status: "okk",
// 	}
// 	fmt.Println(sudip_order.Rp("kundu90"))

//  fmt.Println(sudip_order)



//      sudip:=Adamas{
// 		Name: "Sudip Kundu",
// 		Roll: 125,
// 		Age:23,
// 		Program:"Bca",
// 	 }
// 	 sudip.Age=24

// 	 soubhik :=Adamas{
// 		Name: "Soubhik Kundu",
// 		Roll: 01,
// 		Age: 25,
// 		Program: "Btec",
// 	 }
// 	 soubhik.Age=24
// 	 fmt.Println(sudip)
// 	 fmt.Println(soubhik)


// 	 book:=order{
// 		id: "kolkata89",
// 		amount: 900,
// 		status: "pikup",

// 	 }

	 
// 	 book.changeStatus("ship")
//       fmt.Println(book)
// 	  fmt.Println(book.GetAmount())


// }



package main
import "fmt"


type famaly struct{
	mother string

	father string

	anyBrather string
}

type new struct{
	name string

	age int

	Dob bool

	famaly
}

func main(){


 newjh:=famaly{
	mother: "soma",
	father: "shibu",
	anyBrather: "yes",

 }
 new_01:=new{
	name: "sudip",
	
	age: 23,

	Dob:true,

	famaly:newjh,
	

}

fmt.Println(new_01)





}