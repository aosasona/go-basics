package main

import "fmt"

func main() {
	states := map[string]string{}
	states["CA"] = "California"
	states["NY"] = "New York"
	states["TX"] = "Texas"
	states["FL"] = "Florida"
	states["IL"] = "Illinois"
	states["GA"] = "Georgia"
	states["WA"] = "Washington"
	states["OR"] = "Oregon"

	fmt.Println("Initial length - ", len(states))
	for _, element := range states {
		fmt.Println(element)
	}

	delete(states, "CA")
	delete(states, "NY")
	fmt.Println("After deleting CA and NY - ", len(states))

	GA := states["GA"]
	fmt.Println("GA - ", GA)
}
