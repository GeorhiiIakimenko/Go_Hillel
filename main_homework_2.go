package main

import "fmt"

// Creating the Game structure
type Game struct {
	Name         string
	Genre        string
	TotalPlayers int
	Rating       float64
	ReleaseYear  int
}

// Creating the TableGame structure
type TableGame struct {
	Game       Game
	Additional string
}

func main() {
	// Creating an object and filling all fields
	game1 := TableGame{
		Game: Game{
			Name:         "Chess",
			Genre:        "Board Game",
			TotalPlayers: 2,
			Rating:       4.5,
			ReleaseYear:  1000,
		},
		Additional: "Wooden pieces",
	}

	// Creating a slice and adding objects
	tableGames := make([]TableGame, 0, 3)
	tableGames = append(tableGames, game1)

	// Adding two more objects to the slice
	game2 := TableGame{
		Game: Game{
			Name:         "Monopoly",
			Genre:        "Board Game",
			TotalPlayers: 4,
			Rating:       4.0,
			ReleaseYear:  1935,
		},
		Additional: "Money",
	}
	tableGames = append(tableGames, game2)

	game3 := TableGame{
		Game: Game{
			Name:         "Scrabble",
			Genre:        "Board Game",
			TotalPlayers: 4,
			Rating:       4.2,
			ReleaseYear:  1938,
		},
		Additional: "Letter tiles",
	}
	tableGames = append(tableGames, game3)

	// Displaying the slice
	fmt.Println("Slice tableGames:")
	for _, game := range tableGames {
		fmt.Printf("%+v\n", game)
	}

	// Creating a map and adding elements
	gameMap := make(map[int]TableGame)
	for i, game := range tableGames {
		gameMap[i] = game
	}

	// Displaying the map
	fmt.Println("\nMap gameMap:")
	for _, game := range gameMap {
		fmt.Printf("%+v\n", game)
	}

	// Finding the sum of TotalPlayers in the map
	totalPlayersSum := 0
	for _, game := range gameMap {
		totalPlayersSum += game.Game.TotalPlayers
	}

	// Printing the sum
	fmt.Println("\nSum of TotalPlayers in the map:", totalPlayersSum)
}
