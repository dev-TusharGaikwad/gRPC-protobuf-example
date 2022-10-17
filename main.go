package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	quiz "gRPC-protobuf-example/datastruct"
)

func main() {
	jsonData, err := os.Open("filedata/quizdata.json")

	if err != nil {
		log.Fatal(err)
	}
	jsonByteData, err := ioutil.ReadAll(jsonData)

	if err != nil {
		log.Fatal(err)
	}

	var dataObj quiz.JsonData

	// j := string(jsonByteData)

	// fmt.Println(j)

	err = json.Unmarshal(jsonByteData, &dataObj)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(dataObj, "---------------------")

	for k, v := range dataObj.Quiz {
		fmt.Println(k, v, "----------------------")
	}
}
