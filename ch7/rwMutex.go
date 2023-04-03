package main

import (
	"fmt"
	"sync"
	"time"
)

type Secret struct {
	RWM      sync.RWMutex
	Password string
}

var Password *Secret

func changePass(pass string) {
	fmt.Println("Change()")
	Password.RWM.Lock()
	fmt.Println("Locked!")
	Password.Password = pass
	Password.RWM.Unlock()
	fmt.Println("unlocked")
}

func showPass() {
	defer wg.Done()
	Password.RWM.RLock()
	fmt.Println("Read Lock")
	time.Sleep(2 * time.Second)
	fmt.Println("Pass: ", Password.Password)
	defer Password.RWM.RUnlock()
	fmt.Println("Read Unlock")
}

func main() {
	Password = &Secret{Password: "myPass"}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go showPass()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		changePass("123456")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		changePass("54321")
	}()

	wg.Wait()

	// Direct access to Password.password
	fmt.Println("Current password value:", Password.Password)
}
