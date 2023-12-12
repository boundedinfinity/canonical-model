package main

import "fmt"

func main() {
	fmt.Println("Generating ISO 3166 : Country Codes")
	checkErr(processIso3166())

	fmt.Println("Generating ISO 4217 : Currency Codes")
	checkErr(processIso4217())
}
