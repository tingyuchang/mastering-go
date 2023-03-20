package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	//for i := 0; i < 100; i++ {
	//	fmt.Printf("%d ", random(1, 10))
	//}
	//fmt.Println(getString(16))
	fmt.Println(generatePass(8))
}

//func random(min, max int) int {
//	// generate [min, max)
//	return rand.Intn(max-min) + min
//}
//
//func getString(len int64) string {
//	temp := ""
//	startChar := "!"
//	var i int64 = 1
//
//	for i < len {
//		myRand := random(MIN, MAX)
//		newChar := string(startChar[0] + byte(myRand))
//		temp += newChar
//		i++
//	}
//
//	return temp
//}

func generateBytes(n int64) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generatePass(s int64) (string, error) {
	b, err := generateBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
