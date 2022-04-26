package jsonfilegenerator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func GenerateFileByStruct(fileName string, _struct interface{}) {
	file, err := json.MarshalIndent(_struct, "", " ")
	if err != nil {
		fmt.Println("Error on MarshalIdent to create file", err)
	}

	err = ioutil.WriteFile(fileName+".json", file, 0644)
	if err != nil {
		fmt.Println("Error creatin file", err)
	}
}
