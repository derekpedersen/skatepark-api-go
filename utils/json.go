package utils

import (
	"fmt"
	"io/ioutil"
)

func ReadJsonFile(file string) (string, error) {
	raw, err := ioutil.ReadFile(file)

	str := fmt.Sprintf("%s", raw)

	return str, err
}
