package main

import "fmt"

func main() {
	const USD float32 = 15.40
	const EUR float32 = 150
	const RUB float32 = 25548.5
	fmt.Println("USD in EUR", USD*0.86)
	fmt.Println("USD in RUB", USD*80.50)
	fmt.Println("EUR in RUB", USD*EUR)
}
