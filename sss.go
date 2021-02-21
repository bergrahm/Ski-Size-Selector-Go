package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const noIntValue = -1
const noStringValue = "No flag provided"

type person struct {
	name   string
	age    int
	style  string
	length int
}

func verifyConfig(p person) person {

	if p.age == noIntValue {
		fmt.Print("\nPlease enter age of skier in years\n-> ")
		p.age = promptAge()
		fmt.Printf("Skier is %d years old\n", p.age)
	}

	if p.length == noIntValue {
		fmt.Print("\nLength hasn't been specified.\nPlease enter length of skier in cm:\n-> ")
		p.length = promtLength()
		fmt.Printf("%d cm Selected.\n", p.length)

	}

	if p.style == noStringValue && p.age > 8 {
		fmt.Print("\nPlease enter your prefered style (f for freestyle. c for classic)\n-> ")
		p.style = promptStyle()
		fmt.Println(p.style + " Selected")

	}
	return p
}

func promptAge() int {
	i := parseIntFromStdin()
	if i < 0 || 120 < i {
		fmt.Print("Incorrect input. Please enter an age between 0 and 120\n-> ")
		i = parseIntFromStdin()
	}

	return i
}

func promptStyle() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	for input[:1] != "c" && input[:1] != "f" {
		fmt.Print("Incorrect input. \nPlease type freestyle or classic (f for freestyle or c for classic)\n-> ")
		input, _ = reader.ReadString('\n')
	}

	if strings.ToLower(input[:1]) == "c" {
		return "Classic"
	}
	return "Freestyle"
}

func promtLength() int {
	i := parseIntFromStdin()
	for i < 50 || 220 < i {
		fmt.Print("Incorrect input. Please enter a length between 50 and 220 cm\n-> ")
		i = parseIntFromStdin()
	}

	return i
}

func parseIntFromStdin() int {
	i := noIntValue
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	i, err := strconv.Atoi(input)
	if err != nil {
		i = parseIntFromStdin()
	}
	return i
}

func readConfig() person {
	namePtr := flag.String("n", "noname", "Name of skier")
	stylePtr := flag.String("s", noStringValue, "Prefered style of skier")
	agePtr := flag.Int("a", noIntValue, "Age of skier")
	lengthPtr := flag.Int("l", noIntValue, "Length of skier")

	flag.Parse()

	p := initPerson()
	fmt.Println(p)
	p.name = *namePtr
	p.age = *agePtr
	p.style = *stylePtr
	p.length = *lengthPtr
	fmt.Println(p)
	return *p

}

func initPerson() *person {
	return &person{"name_here", 0, "style", 0}
}
func selectSkiSize(p person) {
	skiLength := 0
	if p.age < 9 {
		skiLength = childSki(p.age) + p.length
		fmt.Printf("Suggested ski-size for a child is %d - %d cm\n", skiLength, skiLength+10)
	}
	if p.style == "Classic" {
		skiLength = classicSki(p.length)
		fmt.Printf("Suggested ski-size for a Classic skier of length %d is %d cm\n", p.length, skiLength)
	}
	if p.style == "Freestyle" {
		skiLength = freestyleSki(p.length)
		fmt.Printf("Suggested ski-size for a freestyle skier is %d - %d cm\n", skiLength, skiLength+5)
	}

}

func childSki(age int) int {
	if age < 5 {
		return 0
	}
	return 10
}

func classicSki(length int) int {
	skiLength := 20 + length
	if 207 < skiLength {
		return 207
	}

	return skiLength
}

func freestyleSki(length int) int {
	return 10 + length
}

func main() {
	p := readConfig()
	p = verifyConfig(p)
	selectSkiSize(p)

}
