package ascii_art

import (
	"bufio"
	"os"
)

func Aff(word, name string) {

	var alphabet [][]string
	var letter []string
	nbr := 1

	file, _ := os.Open(name)

	fileScanner := bufio.NewScanner(file)

	// read line by line
	for fileScanner.Scan() {
		letter = append(letter, fileScanner.Text())
		if nbr%10 == 0 {
			alphabet = append(alphabet, letter)
			letter = letter[len(letter):]
			nbr = 1
		}
		nbr++
	}
	err := file.Close()
	if err != nil {
		return
	}

	//print message
	message := word
	for i := 0; i < 9; i++ {

		for z := 0; z < len(message); z++ {
			if message[z]-32 == 0 {
				print("      ")
			} else {
				print(alphabet[message[z]-32][i])
			}
		}
		print("\n")
	}
}
