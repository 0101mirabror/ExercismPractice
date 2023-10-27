package main


import  "fmt"

func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool{
	return ((!knightIsAwake) && (archerIsAwake) && (prisonerIsAwake) && (!petDogIsPresent) )
}

func main(){

	var knightIsAwake bool = false
	var archerIsAwake bool = true
	var prisonerIsAwake = false
	var petDogIsPresent = false

	fmt.Println(CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent))
	// Output: false
}