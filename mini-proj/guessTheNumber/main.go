package main

import (
	"fmt"
	"math/rand"
)


func Game(maxRange , chances int)  {
	var randomNumber int = rand.Intn(maxRange);
	var flag bool = false;
	println("The random number is ",randomNumber)
	for i:=0;i<chances;i++{
		var input int;
		println("Enter the number between 1 to 10");
		println("Enter the number");

		_, err := fmt.Scanf("%d", &input)
		if err != nil  {
			if input == randomNumber{
				println("You won the game");
				flag = true;
				break;
			}
		}else if input > randomNumber{
			println("The number is greater");	
			
		}else if input < randomNumber{		
			println("The number is smaller");
			
		}
		println("You have %d chances left",chances-i-1)
	}
	if !flag {
		println("You lost the game");
	}
}


func main(){	
	// give the use a range to guess 
	const maxRange int = 10;
	// make the chances fixed to guess the game 
	const chances int = 3;
	
	Game(maxRange,chances);
	
	// generate the random number between the range
	// check the input with the random number
	// provide the hint to the user by saying the number is greater or smaller
	// do this  while the chances are left
	// if the user guess the number then print the message that you won the game
	// if the user not guess the number then print the message that you lost the game
	// ask the user to play again


}