package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type TestArray struct {
	Unsorted []int `json:"unsorted"`
	Sorted   []int `json:"sorted"`
}

func GetArray() ([]TestArray, error) {
	var tests []TestArray

	pwd, _ := os.Getwd()
	jsonFile, err := os.Open(pwd + "/../utils/arrays.json")

	if err != nil {
		return nil, errors.New("cannot open json file")
	}

	defer jsonFile.Close()

	byteArray, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteArray, &tests)

	return tests, nil
}
