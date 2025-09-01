package main
import "fmt"


 func sum(num...int)int{
	total:=0

	for i :=0;i<len(num);i++{

		total+=num[i]

	}

	return total

 }

func main(){
    
	num:=[]int{1,4,2,3}

	fmt.Println("hlw world ")
	result := sum(10,15,20)
	result2 := sum(num ...)

	fmt.Println("Sum:", result , result2)



}