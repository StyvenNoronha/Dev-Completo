package main
import "fmt"
func main(){
	var Player1,Player2 string
	Player2="rock"
	Player1="scissors"

	if Player1 == Player2{
		fmt.Println("Draw!") 
	  } else if Player1 == "rock" && Player2 == "scissors"{
		fmt.Println("Player 1 won!") 
	  }else if Player1 == "scissors" && Player2 == "paper"{
		fmt.Println("Player 1 won!") 
	  }else if Player1 == "paper" && Player2 == "rock"{
		fmt.Println( "Player 1 won!")
	  } else {
		fmt.Println("Player 2 won!") 
	  }
	
}