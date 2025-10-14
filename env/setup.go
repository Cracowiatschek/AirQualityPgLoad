package env

import (
	"errors"
	"os"
	"regexp"
	"strings"
)

type Environment struct {
	Path      string
	Variables map[string]string
}

func (e *Environment) checkPath() (bool, error) {
	var ErrWrongFileType = errors.New("Wrong environment file type. Need '.env' file")

	lastChars := e.Path[-4:]
	if lastChars != ".env" {
		return false, ErrWrongFileType
	}
	_, err := os.Stat(e.Path)

	if os.IsNotExist(err) {
		return false, err
	}

	return true, err
}

func (e *Environment) ReadVariables() error {
	var ErrWrongVariableStruct = errors.New("Wrong variable structure in .env file")
	corrPath, err := e.checkPath()
	variablePattern, _ := regexp.Compile("[A-Z_]*=\\S*")

	if corrPath {
		file, err := os.ReadFile(e.Path)
		if err != nil {
			return err
		}

		fileContent := string(file)
		rawVariables := strings.Split(fileContent, " ")

		for i := 0; i < len(rawVariables); i++ {

			if variablePattern.MatchString(rawVariables[i]) {
				name := strings.Split(rawVariables[i], "=")[0]
				value := strings.Split(rawVariables[i], "=")[1]
				e.Variables[name] = value
			} else {
				return ErrWrongVariableStruct
			}
		}
		return nil
	}

	return err
}
