package prayer

import "fmt"

func PrintCities() {
	for i, city := range Cities {
		fmt.Printf("%d - %s\n", i+1, city.Title)
	}
}

func ReceiveCity() {
	PrintCities()

	var i int
	fmt.Scan(&i)

	Location = Cities[i-1]
}
