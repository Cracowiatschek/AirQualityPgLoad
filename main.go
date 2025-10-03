package main

import (
	"fmt"
	"net/http"
	"bufio"
)

func main() {
	resp, err := http.Get("https://api.gios.gov.pl/pjp-api/v1/rest/station/findAll?size=500")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Ready", resp.Status)
	scanner := bufio.NewScanner(resp.Body)

	for i:=0; scanner.Scan(); i++ {
		fmt.Println(i, scanner.Text(), "\n")
	}

}