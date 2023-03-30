package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
)

type Data struct {
	Key string `json:"key"`
	Val int    `json:"value"`
}

var DataRecords []Data

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

var MIN = 0
var MAX = 26

func getString(l int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == l {
			break
		}
		i++
	}
	return temp
}

func DeSerialized(e *json.Decoder, slice interface{}) error {
	return e.Decode(slice)
}

func Serialized(e *json.Encoder, slice interface{}) error {
	return e.Encode(slice)
}

func PrettyPrint(v interface{}) (err error) {

	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}
	fmt.Println(string(b))

	return nil
}

func JSONStream(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")
	err := encoder.Encode(data)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func main() {
	// Create random records
	var i int
	var t Data
	for i = 0; i < 2; i++ {
		t = Data{
			Key: getString(5),
			Val: random(1, 100),
		}
		DataRecords = append(DataRecords, t)
	}

	fmt.Println("Last record:", t)
	_ = PrettyPrint(t)

	val, _ := JSONStream(DataRecords)
	fmt.Println(val)

	// Create sample data
	//var i int
	//var t Data
	//for i = 0; i < 2; i++ {
	//	t = Data{
	//		Key: getString(5),
	//		Val: random(1, 100),
	//	}
	//	DataRecords = append(DataRecords, t)
	//}
	//
	//// bytes.Buffer is both an io.Reader and io.Writer
	//buf := new(bytes.Buffer)
	//
	//encoder := json.NewEncoder(buf)
	//err := Serialized(encoder, DataRecords)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Print("After Serialize:", buf)
	//
	//decoder := json.NewDecoder(buf)
	//var temp []Data
	//err = DeSerialized(decoder, &temp)
	//fmt.Println("After DeSerialize:")
	//for index, value := range temp {
	//	fmt.Println(index, value)
	//}
}
