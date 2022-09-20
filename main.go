package main 
import "fmt"

type contactInfo struct {
	email string  
	zip int 
}

type person struct {
	firstName string
	lastName string 
}


func main() {
	//alex := person{firstName: "Alex", 
	//			   lastName: "Anderson"}
	//fmt.Println(alex)
	var alex person 

	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	
	fmt.Printf("%+v", alex)

	} 

