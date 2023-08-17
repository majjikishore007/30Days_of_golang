package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


type Response struct{
	Quote Quote `json:"quote"`
}

type Quote struct{
	Body string `json:"body"`
}


func main(){
	response, err := http.Get("https://favqs.com/api/qotd")
	if err != nil {
		fmt.Println("We can't retrieve your quote. You have internet connection?")
	} else {
		fmt.Println("res body ::",response.Body)
		rawQuotes, _ := io.ReadAll(response.Body)
		fmt.Println("raw quotes ::",rawQuotes)
		var quote Response
		json.Unmarshal(rawQuotes, &quote)
		fmt.Println(quote.Quote.Body)
	}
}