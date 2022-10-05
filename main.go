package main

import "fmt"

func main() {
	baseURL := "https://recsys.acm.org/recsys22/session-%s/"
	one := fmt.Sprintf(baseURL, "1")
	fmt.Println(one)
}