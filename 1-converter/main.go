package main

import "fmt"

func main() {
	const UsdToEuro float32 = 0.86
	const UsdToRub float32 = 80.50
	const EuroToRub float32 = UsdToRub / UsdToEuro
	fmt.Println("USD in EUR", UsdToEuro*100)
	fmt.Println("USD in RUB", UsdToRub*100)
	fmt.Println("EURO in RUB", EuroToRub*100)
}
