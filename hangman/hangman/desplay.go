package hangman

import (
	"fmt"
)

func DrawWelcome() {
	fmt.Println(`
.__                                                 
|  |__ _____    ____    ____   _____ _____    ____  
|  |  \\__  \  /    \  / ___\ /     \\__  \  /    \ 
|   Y  \/ __ \|   |  \/ /_/  >  Y Y  \/ __ \|   |  \
|___|  (____  /___|  /\___  /|__|_|  (____  /___|  /
     \/     \/     \//_____/       \/     \/     \/ 

	`)
}

func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
	____
	|    |      
	|    o      
	|   /|\     
	|    |
	|   / \
	_|_
	|   |______
	|          |
	|__________|
		`
	case 1:
		draw = `
	____
	|    |      
	|    o      
	|   /|\     
	|    |
	|    
	_|_
	|   |______
	|          |
	|__________|
		`
	case 2:
		draw = `
	____
	|    |      
	|    o      
	|    |
	|    |
	|     
	_|_
	|   |______
	|          |
	|__________|
		`
	case 3:
		draw = `
	____
	|    |      
	|    o      
	|        
	|   
	|   
	_|_
	|   |______
	|          |
	|__________|
		`
	case 4:
		draw = `
	____
	|    |      
	|      
	|      
	|  
	|  
	_|_
	|   |______
	|          |
	|__________|
		`
	case 5:
		draw = `
	____
	|        
	|        
	|        
	|   
	|   
	_|_
	|   |______
	|          |
	|__________|
		`
	case 6:
		draw = `
	
	|     
	|     
	|     
	|
	|
	_|_
	|   |______
	|          |
	|__________|
		`
	case 7:
		draw = `
	_ _
	|   |______
	|          |
	|__________|
		`
	case 8:
		draw = `
	
		`
	}
	fmt.Println(draw)
}

func drawState(g *Game, guess string) {
	fmt.Print("Guessed: ")
	drawLetters(g.FoundLetters)

	fmt.Print("Used: ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Println("Good guess!")
	case "alreadyGuessed":
		fmt.Printf("Letter '%s' was alread used\n", guess)
	case "badGuess":
		fmt.Printf("Bad guess, '%s' is not in the word", guess)
	case "lost":
		fmt.Print("You lost: (! The word was:")
		drawLetters(g.Letters)
	case "won":
		fmt.Println("You Won! The word was: ")
		drawLetters(g.Letters)
	}
}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v ", c)
	}
	fmt.Println()
}
