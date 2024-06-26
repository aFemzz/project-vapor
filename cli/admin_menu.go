package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"vapor/entity"
	"vapor/handler"
	"vapor/utility"
)

func AdminMenu(admin entity.User, hd *handler.Handler) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Welcome Admin,", admin.Username)
		fmt.Println()
		fmt.Println("Choose menu:")
		fmt.Println("1. Add New Game")
		fmt.Println("2. Update Game")
		fmt.Println("3. Delete Game")
		fmt.Println("4. Report Order")
		fmt.Println("5. Top Selling Publisher")
		fmt.Println("6. Add Admin")
		fmt.Println("7. Log out as Admin")
		fmt.Println("0. Exit program")
		fmt.Println()

		fmt.Print("Input the number: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			fmt.Println("======================================")
			fmt.Println("             ADD NEW GAME ")
			fmt.Println("======================================")
			var game entity.Games
			var err error
			fmt.Print("Insert Game Title: ")
			game.Title, _ = reader.ReadString('\n')
			game.Title = strings.TrimSpace(game.Title)

			fmt.Print("Insert Description: ")
			game.Description, _ = reader.ReadString('\n')
			game.Description = strings.TrimSpace(game.Description)

			fmt.Print("Insert Game Price : ")
			price, _ := reader.ReadString('\n')
			price = strings.TrimSpace(price)
			game.Price, err = strconv.ParseFloat(price, 64)
			if err != nil {
				fmt.Println("Invalid input!!")
				return
			}

			fmt.Print("Insert Game Developer : ")
			game.Developer, _ = reader.ReadString('\n')
			game.Developer = strings.TrimSpace(game.Developer)

			fmt.Print("Insert Game Publisher : ")
			game.Publisher, _ = reader.ReadString('\n')
			game.Publisher = strings.TrimSpace(game.Publisher)

			fmt.Print("Insert Game Rating : ")
			rating, _ := reader.ReadString('\n')
			rating = strings.TrimSpace(rating)
			game.Rating, err = strconv.ParseFloat(rating, 64)
			if err != nil {
				fmt.Println("Invalid input!!")
				return
			}
			err = hd.AddGame(game)
			if err != nil {
				fmt.Println(err)
			}
		case "2":
			fmt.Println("======================================")
			fmt.Println("             UPDATE GAME ")
			fmt.Println("======================================")
			fmt.Print("Input Game Id to edit:")
			inputGameId, _ := reader.ReadString('\n')
			inputGameId = strings.TrimSpace(inputGameId)
			if inputGameId == "" {
				fmt.Println("You have to enter a game id")
				break
			}
			gameId, err := strconv.Atoi(inputGameId)
			if err != nil {
				fmt.Println(err)
				break
			}
			gameOldData, err := hd.GetGameById(gameId)
			if err != nil {
				fmt.Println(err)
				break
			}
			var gameNewData entity.Games
			gameNewData.GameID = gameOldData.GameID

			fmt.Print("Input New Game Title (Press enter to skip):")
			inputTitle, _ := reader.ReadString('\n')
			gameNewData.Title = strings.TrimSpace(inputTitle)
			if gameNewData.Title == "" {
				gameNewData.Title = gameOldData.Title
			}

			fmt.Print("Input New Game Price (Press enter to skip):")
			inputPrice, _ := reader.ReadString('\n')
			inputPrice = strings.TrimSpace(inputPrice)
			if inputPrice == "" {
				gameNewData.Price = gameOldData.Price
			} else {
				gameNewData.Price, err = strconv.ParseFloat(inputPrice, 64)
				if err != nil {
					fmt.Println(err)
					break
				}
			}

			fmt.Print("Input New Game Rating (Press enter to skip):")
			inputRating, _ := reader.ReadString('\n')
			inputRating = strings.TrimSpace(inputRating)
			if inputRating == "" {
				gameNewData.Rating = gameOldData.Rating
			} else {
				gameNewData.Rating, err = strconv.ParseFloat(inputRating, 64)
				if err != nil {
					fmt.Println(err)
					break
				}
			}

			err = hd.UpdateGame(gameNewData)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("Game successfully updated")
			utility.EnterToContinue()
		case "3":
			fmt.Println("======================================")
			fmt.Println("              DELETE GAME ")
			fmt.Println("======================================")
			fmt.Print("Input Game Id to delete:")
			inputGameId, _ := reader.ReadString('\n')
			inputGameId = strings.TrimSpace(inputGameId)
			if inputGameId == "" {
				fmt.Println("You have to enter a game id")
				break
			}
			gameId, err := strconv.Atoi(inputGameId)
			if err != nil {
				fmt.Println(err)
				return
			}
			err = hd.DeleteGame(gameId)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("Game successfully deleted")
			utility.EnterToContinue()
		case "4":
			fmt.Println("Please input range of date report you want")
			fmt.Print("Input Start Date (YYYY-MM-DD):")
			inputStart, _ := reader.ReadString('\n')
			inputStart = strings.TrimSpace(inputStart)
			if inputStart == "" {
				fmt.Println("Please input start date")
				return
			}

			fmt.Print("Input End Date (YYYY-MM-DD):")
			inputEnd, _ := reader.ReadString('\n')
			inputEnd = strings.TrimSpace(inputEnd)
			if inputEnd == "" {
				fmt.Println("Please input end date")
				return
			}
			fmt.Println()
			fmt.Println("=================================================================")
			fmt.Println("                          ORDER REPORT")
			fmt.Println("=================================================================")
			fmt.Println("GAME TITLE           | USERNAME             | DATE             |")
			reportData, total, err := hd.ReportOrder(inputStart, inputEnd)
			if err != nil {
				fmt.Println(err)
				break
			}
			for _, data := range reportData {
				utility.PrintSpace(data.GameTitle, len("GAME TITLE           "))
				utility.PrintSpace(data.Username, len(" USERNAME             "))
				utility.PrintSpace(data.Date, len(" DATE             "))
				fmt.Println()
			}
			fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
			fmt.Printf("Total revenue from %s to %s: %.2f\n", inputStart, inputEnd, total)
			utility.EnterToContinue()
		case "5":
			publisher, err := hd.TopSellingPublisher()
			if err != nil {
				fmt.Println(err.Error())
				break
			}

			fmt.Println("=================================================")
			fmt.Println("                 TOP 5 GAME PUBLISHER")
			fmt.Println("=================================================")
			fmt.Println("GAME PUBLISHER           | TOTAL BUY             |")
			for _, val := range publisher {
				utility.PrintSpace(val.Name, len("Game PUBLISHER           "))
				utility.PrintSpace(val.TotalBuy, len(" Total Buy             "))
				fmt.Println()
			}
			fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		case "6":
			fmt.Println("=============================================")
			fmt.Println("               ADD NEW ADMIN")
			fmt.Println("=============================================")
			fmt.Print("Insert admin name: ")
			username, _ := reader.ReadString('\n')
			username = strings.TrimSpace(username)

			fmt.Print("Insert admin email: ")
			email, _ := reader.ReadString('\n')
			email = strings.TrimSpace(email)

			fmt.Print("Insert admin password: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			err := hd.AddAdmin(username, password, email)
			if err != nil {
				fmt.Println(err)
			}

		case "7":
			fmt.Println("Logged out as admin ... ")
			return
		case "0":
			fmt.Println("Exit program ... ")
			os.Exit(1)
		}
	}
}
