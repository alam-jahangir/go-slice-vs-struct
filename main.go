package main

import "fmt"

func main() {

	// Test variable
	// testVar := "Hello World"
	// fmt.Println(&testVar)  // Show Memory Address
	// fmt.Println(testVar)   // Show Value
	// fmt.Println(*&testVar) // Show Value

	test := person{
		firstName: "Test",
		lastName:  "Test1",
		contact: contactInfo{
			email:   "test@test.com",
			zipCode: 11200,
		},
		/*contactInfo: contactInfo{
			email:   "test@test.com",
			zipCode: 11200,
		},*/
	}

	contact := contactInfo{
		email:   "test1@test.com",
		zipCode: 9900,
	}
	contact.print()

	test.print()

	// Actual way compare to other language
	// test1 := &test
	// test1.updateName("Hello")

	// Go has special feature to identify automatically
	test.updateName("Hello")

	test.print()

	// Difference between slice and struct
	// GO is a copy of Value ( means pass by value )
	// Slice create 2 different data structure
	// DATA_STRUCTUTE_1: *) Pointer to Head *) Capacity *) Length . And this DATA_STRUCTURE_1 point to Array
	// DATA_STRUCTURE_2: Array
	// So For Slice, when we pass slice variable as value, both argument variable & original variable point to same array.
	// That's why, when we change any value in array, both variable will share same array data
	// Same thing happing for map, channel
	mySlice := []string{"A", "B", "C", "D"}
	updateMySlice(mySlice)
	fmt.Println(mySlice)

}

func updateMySlice(s []string) {
	s[0] = "Z"
}
