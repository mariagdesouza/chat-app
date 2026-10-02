package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	c := NewChatHandler()

	http.HandleFunc("/send", c.ReceiveMessage)
	http.HandleFunc("/check", c.CheckMessages)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

type MessageCore struct {
	Text     string `json:"text"`
	FromUser string `json:"from"`
}

type MessageDigest struct {
	Message MessageCore `json:"message"`
	ToUser  string      `json:"to"`
}

type ChatHandler struct {
	UserMessages map[string][]MessageCore
}

func NewChatHandler() *ChatHandler {
	chatHandler := new(ChatHandler)
	chatHandler.UserMessages = make(map[string][]MessageCore)
	return chatHandler
}

func (c *ChatHandler) CheckMessages(w http.ResponseWriter, r *http.Request) {
	userVal := r.URL.Query()["user"]
	if len(userVal) > 0 {
		user := userVal[0]
		if user != "" {
			messages, ok := c.UserMessages[user]
			fmt.Println(messages)
			if ok {
				jsontext, err := json.Marshal(messages)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Fprintf(w, string(jsontext))
			}
		}
	}
}

func (c *ChatHandler) ReceiveMessage(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		fmt.Fprintf(w, "No message")
		return
	}
	var msg MessageDigest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	err = json.Unmarshal(body, &msg)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	if msg.ToUser != "" && msg.Message.Text != "" {
		message := MessageCore{Text: msg.Message.Text, FromUser: msg.Message.FromUser}
		c.UserMessages[msg.ToUser] = append(c.UserMessages[msg.ToUser], message)
	}
	//	fmt.Println(c.UserMessages[msg.ToUser])
	//fmt.Println("Message from", msg.Message.FromUser, "for", msg.ToUser, ":", msg.Message.Text)
}
