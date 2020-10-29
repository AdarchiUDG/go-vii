package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"sort"
)

func readLine(in *bufio.Reader) string {
	line, _ := in.ReadString('\n')
	return strings.TrimSuffix(line, "\n")
}

func main() {
	var numberOfInputs int
	inputs := []string{}
	in := bufio.NewReader(os.Stdin)

	ascending, err := os.Create("ascendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer ascending.Close()

	descending, err := os.Create("descendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer descending.Close()
	
	fmt.Print("Numero de strings: ")
	fmt.Scanf("%d", &numberOfInputs)
	fmt.Scanln()
	for i := 0; i < numberOfInputs; i++ {
		fmt.Print(i + 1, "> ")
		inputs = append(inputs, readLine(in))
	}

	sort.Strings(inputs)
	ascending.WriteString(strings.Join(inputs, "\n"))

	sort.Sort(sort.Reverse(sort.StringSlice(inputs)))
	descending.WriteString(strings.Join(inputs, "\n"))
}