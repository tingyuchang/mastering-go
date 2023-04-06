package main

import (
	"net"
	"time"
)

var timeout = time.Duration(time.Second)

func timeoutHandler(network, host string) (net.Conn, error) {
	conn, err := net.DialTimeout(network, host, timeout)

	if err != nil {
		return nil, err
	}

	conn.SetDeadline(time.Now().Add(timeout))
	return conn, nil
}

func main() {
	//t := http.Transport{
	//	Dial: timeoutHandler,
	//}
	//c := &http.Client{
	//	Transport: &t,
	//}
}
