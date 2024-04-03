package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
)

type MyMessage struct {
	Text   string `json:"text"`
	ChatID string `json:"chat_id"`
}

func main() {
	e := echo.New()
	e.POST("/", messageHandler)
	e.Logger.Fatal(e.Start(":8787"))
}

func messageHandler(c echo.Context) error {
	msg := new(MyMessage)
	if err := c.Bind(msg); err != nil {
		return err
	}
	err := msg.SendMessage()
	if err != nil {
		log.Println("Failed to push the logs")
		return err
	}
	return c.JSON(http.StatusCreated, "The Text was successfuly sent to the wanted group")
}

func (mes *MyMessage) SendMessage() error {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	if mes.ChatID == "" {
		mes.ChatID = os.Getenv("DEFAULT_CHAT_ID")
	}
	message := map[string]interface{}{
		"chat_id": mes.ChatID,
		"text":    mes.Text,
	}
	fmt.Println(message)
	messageBytes, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error encoding message:", err)
		return err
	}
	time.Sleep(1 * time.Second)
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(messageBytes))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error response from Telegram API:", resp.Status)
	} else {
		fmt.Println("Message sent successfully!")
	}
	return nil
}
