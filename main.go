package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Игра для угадывания рандомного числа. У меня вышло с первого раза!

func main() {
	second := time.Now().Unix()
	rand.NewSource(second)
	target := rand.Intn(100) + 1
	var quesses int
	var success = false

	fmt.Println(target)
	for trying := 0; trying <= 10; trying++ {
		fmt.Println("You have:", trying, "guesses left")
		fmt.Println("Write a number:")
		quesses = reader()
		if quesses < target {
			fmt.Println("you quess was LOW")
		} else if quesses > target {
			fmt.Println("you quess was HIGH")
		} else {
			fmt.Println("you win")
			success = true
			break
		}
	}

	if !success {
		fmt.Println("Sorry, you didn't quess. It was:", target)
	}
}

func reader() int {

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	grade, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return grade
}
