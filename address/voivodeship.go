package address

import (
	"regexp"
	"strconv"
	"fmt"
)

type Voivodeship struct {
	Country				string
	Voivodeship			string
	VoivodeshipShort	string
}

type ValidationError struct {
	message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s", e.message)
}


func (v *Voivodeship) Validate() error {
	var message string

	countryLen := len(v.Country) > 20 && len(v.Country) > 0
	voivodeshipLen := len(v.Voivodeship) > 20 && len(v.Voivodeship) > 0
	voivodeshipShortLen := len(v.VoivodeshipShort) == 3

	countryChars, _ := regexp.MatchString("^[A-Z][a-ząćęłńóśźż]*", v.Country)
	voivodeshipChar, _ := regexp.MatchString("^[A-Z][a-ząćęłńóśźż]*-[A-Z][a-ząćęłńóśźż]*|[A-Z][a-ząćęłńóśźż]*", v.Voivodeship)
	voivodeshipShortChar, _ := regexp.MatchString("^[A-Z]*", v.VoivodeshipShort)

	if countryLen != true || countryChars != true {
		message += "Something went wrong with Country: value "+ v.Country + "(available characters: A-z with polish special chars), length "+ strconv.FormatInt(int64(len(v.Country)), 10) + " (range 1-20)."
	}

	if voivodeshipLen != true || voivodeshipChar != true {
		if len(message) > 0{
			message += "\n"
		}
		message += "Something went wrong with Voivodeship: value "+ v.Voivodeship + "(available characters: A-z with polish special chars, available two words separated '-'), length "+ strconv.FormatInt(int64(len(v.Voivodeship)), 10) + " (range 1-20)."
	}

	if voivodeshipShortLen != true || voivodeshipShortChar != true {
		if len(message) > 0{
			message += "\n"
		}
		message += "Something went wrong with Voivodeship: value "+ v.VoivodeshipShort + "(available characters: A-z with polish special chars), length "+ strconv.FormatInt(int64(len(v.VoivodeshipShort)), 10) + " (value must have 3 chars)."
	}

	if len(message) > 0 {
		return &ValidationError{message: message}
	}

	return nil
}