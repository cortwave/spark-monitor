package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func sendNotification(text string, url string) {
	m := map[string]string{"text": text}
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	http.Post(url, "application/json", bytes.NewBuffer(b))
}
