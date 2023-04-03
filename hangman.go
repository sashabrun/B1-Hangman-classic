package classic

import (
	"bufio"
	asciiArt "classic/ascii-art"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type HangManData struct {
	WordBase   string
	ShowWord   []string
	LetterFind string
	Attempts   int
	Position   int
	Ascii      string
}

func Hangman(path []string) {
	//------------------------ launch of program --------------------------------//
	var data HangManData
	if len(Checking(os.Args[1], "--startWith")) > 0 {
		content, _ := os.ReadFile(path[0] + path[1] + os.Args[2])
		err1 := json.Unmarshal(content, &data)
		if err1 != nil {
			print(err1)
			return
		}
	} else {
		word := RandomWord(path[0] + path[1] + path[2] + os.Args[1])
		data = HangManData{word, WordChoice(word), "", 10, -1, ""}
	}
	if len(os.Args[1:]) > 1 && os.Args[2] == "--letterFile" {
		data.Ascii = path[0] + path[1] + path[3] + os.Args[3]
	}
	var letter rune
	var choice string
	fail := false

	println("Good Luck, you have ", data.Attempts, " attempts.")

	// ------------------------ body of program --------------------------------//
	for data.Attempts > 0 {
		ShowText(data.Ascii, data.ShowWord)
		print("\n" + "\n" + "Choose: ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			return
		}
		choice = Lower(choice)

		if len(choice) == 1 {
			letter = rune(choice[0])
			if len(Checking(data.LetterFind, choice)) > 0 {
				data.Attempts--
				data.Position = Gallows(1, data.Position, path)
				println("\nalready present in the word,", data.Attempts, "attempts remaining\n")
				fail = true
			}
			data.LetterFind += choice
			if len(Checking(data.WordBase, choice)) >= 1 {
				index := Checking(data.WordBase, choice)
				for i := 0; i < len(index); i++ {
					data.ShowWord[index[i]] = string(letter - 32)
				}
			} else {
				if !fail {
					data.Attempts--
					data.Position = Gallows(1, data.Position, path)
					println("\nNot present in the word,", data.Attempts, "attempts remaining\n")
					fail = false
				}
			}
		} else {
			if choice == data.WordBase {
				println("\nCongrats !")
				return
			} else if choice == "stop" {
				b, _ := json.Marshal(data)
				save, _ := os.Create("save.txt")
				_, err := save.Write(b)
				if err != nil {
					return
				}
				return
			} else {
				data.Attempts -= 2
				data.Position = Gallows(2, data.Position, path)
				println("\nlie! is not the real word,", data.Attempts, "attempts remaining\n")
			}
		}
		if len(Checking(ListToString(data.ShowWord), "_")) == 0 {
			ShowText(data.Ascii, data.ShowWord)
			println("\n\nCongrats !")
			return
		}
	}
	println("You are bad, it's was :", data.WordBase)
}

func RandomWord(word string) string {
	rand.Seed(time.Now().UnixNano())
	file, _ := os.Open(word)
	var wordList []string
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		wordList = append(wordList, fileScanner.Text())
	}
	return wordList[rand.Intn(len(wordList))]
}

func Checking(word, choice string) []int {
	var ListInd []int
	for i := 0; i < len(word); i++ {
		if choice[0] == word[i] {
			ListInd = append(ListInd, i)
		}
	}
	return ListInd
}

func WordChoice(mot string) []string {
	rand.Seed(time.Now().UnixNano())
	var showWord []string
	nbrLetter := len(mot)/2 - 1
	for i := 0; i < len(mot); i++ {
		showWord = append(showWord, "_")
	}
	for x := 0; x < nbrLetter; x++ {
		ind := rand.Intn(len(mot))
		showWord[ind] = string(mot[ind] - 32)
	}
	return showWord
}

func Gallows(nbr, position int, path []string) int {
	jose, _ := os.ReadFile(path[0] + path[1] + "hangman.txt")
	position += 79 * nbr
	if position >= 710 {
		position = 709
	}
	fmt.Print(string(jose[position-78 : position]))
	return position
}

func ListToString(list []string) string {
	char := ""
	for i := 0; i < len(list); i++ {
		char += list[i]
	}
	return char
}

func StringToList(char string) []string {
	var list []string
	for i := 0; i < len(char); i++ {
		list = append(list, string(char[i]))
	}
	return list
}

func Lower(choice string) string {
	choice3 := ""
	for h := 0; h < len(choice); h++ {
		if choice[h] >= 65 && choice[h] <= 90 {
			choice3 += string(choice[h] + 32)
		} else {
			choice3 += string(choice[h])
		}
	}
	return choice3
}

func Upper(choice string) string {
	choice3 := ""
	for h := 0; h < len(choice); h++ {
		if choice[h] >= 97 && choice[h] <= 122 {
			choice3 += string(choice[h] - 32)
		} else {
			choice3 += string(choice[h])
		}
	}
	return choice3
}

func ShowText(maj string, word []string) {
	if maj != "" {
		asciiArt.Aff(ListToString(word), maj)
	} else {
		for i := 0; i < len(word); i++ {
			print(word[i] + " ")
		}
	}
}
