package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	x := "This is my name Md.Nihad Hossain.\nWhere I live in?I live in Bangladesh,Dhaka.\nI Love to learn new things."
	scanner := bufio.NewScanner(strings.NewReader(x))
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
