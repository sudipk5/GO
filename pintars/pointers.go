package main
import "fmt"
func sum(num *int){
	*num =5
	fmt.Println(*num)
}
func main(){
   num := 1
   sum(&num )
   fmt.Println(num) 
}