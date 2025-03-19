package main

import "fmt"

func main() {
	login := promptData("Enter Login")
	password := promptData("Enter Password")
	url := promptData("Enter URL")

	outputPassword(login, password, url)
}

func promptData(prompt string) string {
	fmt.Println(prompt)
	var res string
	fmt.Scanln(&res)

	return res
}

func outputPassword(login, password, url string) {
	fmt.Println(login)
	fmt.Println(password)
	fmt.Println(url)
}
