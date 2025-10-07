package main

import (
	"fmt"
//	"io"
//	"net/http"
////	"bufio"
//	"encoding/json"
	"air-quality-loader/request"
)

func main() {
	pauses := request.RequestPauses{ // time in milliseconds
		Station: 30_000,
		Sensor: 50,
	}

	r := request.Request{}

	p, err := r.SetRequest("https://api.gios.gov.pl/pjp-api/v1/rest/station/findAll")
	if err != nil {
		panic(err)
	}
	if p != true {
		panic("Wrong request")
	}

	if err := r.SetPages(); err != nil {
		panic(err)
	}

	if err := r.GetData(true, pauses.Station, "Lista stacji pomiarowych"); err != nil {
		panic(err)
	}
	for i := range len(r.Body) {
		fmt.Println(r.Body[i])
	}
	fmt.Print(r.Body)
}