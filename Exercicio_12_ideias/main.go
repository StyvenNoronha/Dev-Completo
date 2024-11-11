package main

import "fmt"

func main() {
	//ideas := []string{"good", "bad", "good", "good", "bad", "good", "bad", "bad", "good", "bad", "bad"}
	//ideas1 := []string{"bad", "bad", "bad", "bad", "good", "good", "bad", "bad", "bad"}
	 ideas2:= []string{"bad", "bad", "bad", "bad", "bad"}

	var numberOffGoodIdeas int
	for _, idea := range ideas2 {
		if idea == "good" {
			numberOffGoodIdeas += 1
		}
	}
	if numberOffGoodIdeas == 0 {
		fmt.Println("Fail!")
	} else if numberOffGoodIdeas <= 2 {
		fmt.Println("Publish!")
	} else {
		fmt.Println("I smell a series!")
	}

}