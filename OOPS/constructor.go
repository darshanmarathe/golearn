package main

import (
	"fmt"
)

func main() {
	newSession := newSession()
	newSession.values["username"] = "darshan"
	fmt.Println(newSession)
}

type Session struct {
	values map[string]string
}

func newSession() *Session {
	result := Session{}
	result.values = map[string]string{}
	return &result
}
