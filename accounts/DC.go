package main

import (
	"fmt"
	"net/http"
)

func main() {
	// url := "https://discord.com/invite/BNcynmAxvp"

	// 
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the GET request:", err)
		return
	}
	defer resp.Body.Close()

	// 
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Headers:", resp.Header)
}
