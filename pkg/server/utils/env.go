package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func setEnv() {
	readFile, err := os.Open("./.env")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("The following environmental variables have been set")
	log.Print()
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)
	i := 0
	for scanner.Scan() {
		strArr := strings.Split(scanner.Text(), "=")
		if len(strArr) != 2 {
			continue
		}
		fmt.Printf("%s ", strArr[0])
		if i == 2 {
			fmt.Println()
		}
        i++
		i %= 3
		os.Setenv(strArr[0], strArr[1])
	}
	fmt.Println()
	log.Println("Environmental variables have been set.")
}
