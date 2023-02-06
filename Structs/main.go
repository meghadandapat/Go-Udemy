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
	jim.updateName("jimmy")
	jim.print()

}

// reference func -> this basically mean we can call this func with any person type
func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
