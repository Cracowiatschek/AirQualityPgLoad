package request


import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"io"
	"strconv"
	"time"
)

type Request struct {
	Pages   float64
	Request string
	Body    []map[string]interface{}
}


type requestMismatch struct {
	arg 	int
	message string
}

func (e *requestMismatch) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func (r *Request) SetRequest (req string) (bool, error) {
	pattern, _ := regexp.Compile("https://[a-z, .]*.pl/[a-z, -]*/v1/rest/[a-z, A-Z/]*")

	match := pattern.MatchString(req)

	if match {
		r.Request = req
		return true, nil
	} else {
		return false, &requestMismatch{404, "Wrong structure of request" }
	}
}

func (r *Request) SetPages () (error) {
	var response map[string]interface{}
	if r.Request == "" {
		panic("Empty request")
	}
	req, err := http.Get(r.Request)

	if err != nil {
		return err
	}
	if req.StatusCode > 299 {
		panic(req.StatusCode)
	}

	body, _ := io.ReadAll(req.Body)
	req.Body.Close()

	if err := json.Unmarshal(body, &response); err != nil {
		panic(err)
	}
	r.Pages = response["totalPages"].(float64)
	return nil
}

func (r *Request) GetData (allPages bool, pauseTime int, item string) (error) {
	var responseDecoded map[string]interface{}

	if r.Request == ""{
		panic("Empty request")
	}

	if allPages == false {
		response, err := http.Get(r.Request)
		if err != nil {
			return err
		}

		body, _ := io.ReadAll(response.Body)

		if err := json.Unmarshal(body, &responseDecoded); err != nil {
			return err
		}

		values := responseDecoded[item].([]interface{})
		r.extractData(values)
	} else {
		for i := range int64(r.Pages) {
			fmt.Println(i,"Start", time.Now())
			request := r.Request+"?page="+strconv.FormatInt(i, 10)
			response, err := http.Get(request)
			if err != nil {
				return err
			}

			body, _ := io.ReadAll(response.Body)

			if err := json.Unmarshal(body, &responseDecoded); err != nil {
				return err
			}

			values := responseDecoded[item].([]interface{})
			r.extractData(values)
			fmt.Println(i,"Stop", time.Now())
			time.Sleep(time.Duration(pauseTime) * time.Millisecond)
		}
	}
	return nil
}

func (r *Request) extractData (data []interface{}) {
	for i:=0; i < len(data); i++ {
		r.Body = append(r.Body, data[i].(map[string]interface{}))
	}
}
