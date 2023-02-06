package main

import "fmt"

//new custom type 'person'

// blueprint of person
type person struct {
	firstName string
	lastName  string
	contact   contactInfo //can also use a syntax 'contactInfo' in this the case the fielname will be same as type name
}

// blueprint of contactInfo
type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// *****************create an acual instance of person type*******************

	/*
		// method 1: relies on order of definition
		megha := person{"Megha", "Dandapat"}
		fmt.Println(megha)
		// method 2: independent of order
		alex := person{firstName: "Alex", lastName: "Anderson"}
		fmt.Println(alex)
		//method 3:
		var harry person //by default assigns zero value (for string:"", int,float:0, bool:false)

		//Update values
		harry.firstName = "Harry"
		harry.lastName = "Styles"
		fmt.Printf("%+v", harry) //print with property names
	*/

	//***********Nested structs**************
	//Note: put comma at end of each property even after last one
	//Note: no comma between properties while defining th type only while assigning values as shown below

	jim := person{
		firstName: "Jim",
		lastName:  "Halpert",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 76622,
		},
	}
	jimPointer := &jim
	jimPointer.updateName("jimmy")

	//IMP NOTE: The above two statements are equivalent to jim.updateName("jimmy")
	//There is no need to explicitly make an address type variable when our receiver is of type pointer as it will automatically behave as pass by reference
	jim.print()

}

// receiver func -> this basically mean we can call this func with any person type
func (p person) print() {
	fmt.Printf("%+v", p)
}

// ************Pass by Value***********
/*
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName //update did not actually take place since p is only a copy of jim
	//Since go is a pass by value thus it is only updating the value of copy not of the actual var
}
*/

// ************Pass by Reference***********

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName //must use brackets to enclose pointertype value
}

/*
Note:
1. When there is * in front of a pointer it gives the value of pointer
2. When there is * in front of a type then it is a description of the type. It means we require a pointer to that type
3. If we define a receiver with a type of pointer then Go allows us to call the function either with a pointer or the root type. In this case pointerToPerson and Person type both will work with the above function and will actuallly update original value in both case
*/

//*************Reference v/s Value Types***************
/*
Reference Types: Slice, Maps, Channels, pointers, functions -> still makes a copy but pointers are not required to update the original value since they are implemented such that it carry reference to the original vale

For eg. Slices conatin a pointer to head of an underlying array thus even after making a new copy of slice still it is pointing to the same head of array

Value Types: int, float, string, bool, structs -> requires pointers while updating values while passed inside function
*/
