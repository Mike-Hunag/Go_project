package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/line/line-bot-sdk-go/linebot"
)

var (
	client *linebot.Client
)

type Message struct {
	Epoch    int     `json:"Epoch"`
	Loss     float64 `json:"Loss"`
	Accuracy float64 `json:"Accuracy"`
	CER      float64 `json:"CER"`
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	var messageData Message
	if err := json.Unmarshal(data, &messageData); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	client, err = linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_ACCESS_TOKEN"))

	if err != nil {
		log.Println(err.Error())
	}

	messageText := fmt.Sprintf("Epoch: %d\nLoss: %.4f\nAccuracy: %.4f\nCER: %.4f", messageData.Epoch, messageData.Loss, messageData.Accuracy, messageData.CER)

	message := linebot.NewTextMessage(messageText)

	if _, err := client.PushMessage(os.Getenv("USERID"), message).Do(); err != nil {
		log.Print(err)
	} else {
		fmt.Println("Push message sent to LINE User")
	}
}
