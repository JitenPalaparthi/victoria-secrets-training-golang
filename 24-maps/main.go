package main

import "fmt"

func main() {

	var mymap map[string]any // declaration but not instantiation
	mymap = make(map[string]any)
	mymap["522001"] = "Guntur-1"
	mymap["522002"] = "Guntur-2"
	mymap["560086"] = "Yeshvantpur Bangalore"
	mymap["560096"] = "Mahalakshmi Bangalore"
	mymap["690054"] = "Trivandrum-1"

	for key, value := range mymap {
		fmt.Println("Key:", key, "Value", value)
	}

	_, ok := mymap["522003"]
	if ok {
		fmt.Println("Key exists")
	} else {
		fmt.Println("Key does not")
	}

	delete(mymap, "522001")
	delete(mymap, "5234234")
	for key, value := range mymap {
		fmt.Println("Key:", key, "Value", value)
	}
}
