# Hangman-Classic

***

## Rules ##

You will have 10 attempts to complete the game.

   - First, the program will randomly choose a word in the file.
    
   - The program will reveal n random letters in the word, where n is the len(word) / 2 - 1
    
   - The program will read the standard input to suggest a letter.
    
   - If the letter is not present, it will print an error message and the number of attempts decreases (10->9->…0)
    
   - If the letter is present, it will reveal all the letters corresponding in the word.
    
   - The program continues until the word is either found, or the numbers of attempts is 0.

Let’s name José the poor man that will be hanging to this rope if you lose.

You will be given a file named hangman.txt that contains all the position of José. This file contains 10 positions corresponding to the 10 attempts needed to complete the game.

You will need to parse this file and display the appropriate position of José as the count of attemps decreases.



## bonus advanced-feature ##

The player can suggest either a word or a letter. (A word is a string of at least two character long)

If the word is found the game stops

If not, the counter of attempts decrease by 2

Stock the letters that are suggested by the player so the player can’t propose the same letter twice.

You can display error message when this case happens.

## bonus  ##

This project should display letters in ASCII-Art from those following files

First you need to pass the file in argument containing the ascii letters as argument to your program (see example below).

You will then need to parse this file.

## bonus : start-and-stop ##


Implement a keyword STOP in the standard input.

It will stop and exit the game.

It will save the status of the game in a file save.txt. The data in the file must be encoded with json.Marshal

Handle a flag --startWith save.txt in the command line, that allow you to launch the game with the file you saved with STOP command. The file will be decoded with json.UnMarshal

***

## Installation ##

download the repo : [here](https://ytrack.learn.ynov.com/git/fleo/hangman_classic_base/archive/master.zip)

or clone the repo :

```bash
git clone https://ytrack.learn.ynov.com/git/fleo/hangman_classic_base
```

and enjoy with the commands :

- level easy
```bash
go run . words.txt
```

- level medium
```bash
go run . words2.txt
```

- level hard
```bash
go run . words3.txt
```

- level easy + asci-art standard font
```bash
go run . words.txt --letterFile standard.txt
```

- level easy + asci-art shadow font
```bash
go run . words.txt --letterFile shadow.txt
```

- level easy + asci-art thinkertoy font
```bash
go run . words.txt --letterFile thinkertoy.txt
```

- play an old save
```bash
go run . --startWith save.txt
```

***

## Credit ##

Fauré Léo

Sasha Brun