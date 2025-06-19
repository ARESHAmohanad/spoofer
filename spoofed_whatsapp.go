package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Rhymen/go-whatsapp"
)

func main() {
	// Initialize WhatsApp connection
	wac, err := whatsapp.NewConn(5 * time.Second)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating connection: %v\n", err)
		return
	}

	// Login to WhatsApp (replace with your QR code or session)
	qrChan := make(chan string)
	go func() {
		qr := <-qrChan
		fmt.Printf("Scan the QR code: %s\n", qr)
	}()

	session, err := wac.Login(qrChan)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error logging in: %v\n", err)
		return
	}
	fmt.Printf("Logged in: %v\n", session)

	// Send a text message
	chatID := "5511999999999@s.whatsapp.net" // Replace with the target chat ID
	msgID, err := wac.Send(whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: chatID,
		},
		Text: "Hello, this is a spoofed message!", // Replace with your message text
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error sending message: %v\n", err)
		return
	}
	fmt.Printf("Message sent successfully! Message ID: %s\n", msgID)
}
