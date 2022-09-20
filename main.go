package main 
import "fmt"

type contactInfo struct {
	email string  
	zipCode int 
}

type person struct {
	firstName string
	lastName string 
	contact contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contact: contactInfo{
			email:"jim@gmail.com",
			zipCode: 94000,
		},

	}
	
	// jimPointer := &jim 
	// &variable: give me the memory address of the value this variable(jim) is pointing at, assigning address to new variable(jimpointer) before :=
	
	jim.updateName("Jimmy") // shortcuts of pointer 
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstname string) { // *pointer: Give me the value this memory address is point at 
	(*pointerToPerson).firstName = newFirstname // <- or you can call instance of the type and its func directly(line 28) but not need to assign a var of pointer to the type
												// to get the address of instance and change its value; for example, jim.updateName("Jimmy") will still work if this func of 
}												// type's receiver is pointer to this type;

func (p person) print() {
	fmt.Printf("%+v", p)
}

// address and value:
// if you have address and you want to get value of it, use *address(pointer to value)
// if you have value and you want to get address of the value, use &value
// * before a type is just description: we are working on a pointer to this type 
// * before a pointer is operator: we can manipulate the value that the pointer references 

// PS: value types(need pointer to change original value):int, float, string, bool, struct;
// Reference types(do not need pointer to change original value): slice, maps, channels, pointers, functions; 