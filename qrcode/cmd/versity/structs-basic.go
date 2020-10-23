package main

import "fmt"

type people []struct {
	name      string
	age       int
	aadharnum string
	address
}

type address struct {
	name string
}

type friends []struct {
	groupName string
	count     int
	people
}

func main() {
	var count int
	count = 1
	var i []int
	i = []int{3, 4, 6}
	fmt.Println("i is", i)

	var p people
	p = people{
		{
			name: "ram", age: 43, aadharnum: "12321-12312-1231",
		},
		{
			name: "ram", age: 43, aadharnum: "12321-12312-1231",
		},
	}

	fmt.Printf("person %v\n", p)

	fmt.Printf("count %#v\n", count)

	var f friends
	f = friends{
		{
			"kokorona",
			4,
			people{
				{
					"Suresh",
					36,
					"1212-123-1123",
					address{name: "Cisco"},
				},
				{
					"Umang",
					32,
					"1212-123-1132",
					address{name: "Pari"},
				},
				{
					"Rekha",
					32,
					"1212-132-1132",
					address{name: "Wipro"},
				},
			},
		},
	}

	fmt.Println("Friends", f)
	for i2, fr := range f {
		fmt.Printf("Members of friends group %v: %v %v %v %v %v %v\n",
			fr.groupName,
			fr.people[i2].name,
			fr.people[i2+1].address.name,
			fr.people[i2+2].name,
			fr.people[i2].address.name,
			fr.people[i2+1].name,
			fr.people[i2+2].address.name,
		)
	}

}
