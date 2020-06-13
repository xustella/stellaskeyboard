package stellaskeyboard

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadInt() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	//fmt.Println(len(input))

	// Atoi converts a string into an integer
	guess, err := strconv.Atoi(input)
	return guess, err
}

func ReadFloat64() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	//fmt.Println(len(input))

	//ParseFloat converts a string into a float64
	guess, err := strconv.ParseFloat(input, 64)
	return guess, err
}
