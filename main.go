package main

import (
	"bufio"
	"encoding/json"
	"fmt"

	// "math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	// "time"
)

// User struct represents a user with a unique ID and their message log
type User struct {
	ID  int
	Log []string
}

// Message struct represents a message with sender ID, receiver ID, and content
type Message struct {
	SenderID   int
	ReceiverID int
	Content    string
}

// Global variables to store users and messages
var users map[int]*User
var messages []Message

func main() {
	users = make(map[int]*User)
	initializeUsers()

	defer func() {
		fmt.Println("\nMessage Logs:")
		for _, user := range users {
			fmt.Printf("User %d:\n", user.ID)
			for _, msg := range user.Log {
				fmt.Println(msg)
			}
			fmt.Println()
		}
	}()

	fmt.Println("Welcome to CLI Messaging App")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nOptions:")
		fmt.Println("1. Send Message between two users")
		fmt.Println("2. Broadcast Message to all users")
		fmt.Println("3. View Message Log of a User")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			sendMessage(reader)
		case "2":
			broadcastMessage(reader)
		case "3":
			viewMessageLog(reader)
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

// Function to initialize some sample users
func initializeUsers() {
	users[1] = &User{ID: 1}
	users[2] = &User{ID: 2}
	users[3] = &User{ID: 3}
}

// Function to send a message between two users
func sendMessage(reader *bufio.Reader) {
	fmt.Print("Enter sender ID: ")
	senderIDStr, _ := reader.ReadString('\n')
	senderIDStr = strings.TrimSpace(senderIDStr)
	senderID, _ := strconv.Atoi(senderIDStr)

	fmt.Print("Enter receiver ID: ")
	receiverIDStr, _ := reader.ReadString('\n')
	receiverIDStr = strings.TrimSpace(receiverIDStr)
	receiverID, _ := strconv.Atoi(receiverIDStr)

	fmt.Print("Enter message content: ")
	content, _ := reader.ReadString('\n')
	content = strings.TrimSpace(content)

	if content == "" {
		content = getRandomFact()
	}

	if _, ok := users[receiverID]; ok {
		messages = append(messages, Message{SenderID: senderID, ReceiverID: receiverID, Content: content})
		users[receiverID].Log = append(users[receiverID].Log, fmt.Sprintf("User %d received message from User %d: %s", receiverID, senderID, content))
		fmt.Println("Message sent successfully!")
	} else {
		fmt.Println("Receiver ID does not exist.")
	}
}

// Function to broadcast a message to all users
func broadcastMessage(reader *bufio.Reader) {
	fmt.Print("Enter sender ID: ")
	senderIDStr, _ := reader.ReadString('\n')
	senderIDStr = strings.TrimSpace(senderIDStr)
	senderID, _ := strconv.Atoi(senderIDStr)

	fmt.Print("Enter message content: ")
	content, _ := reader.ReadString('\n')
	content = strings.TrimSpace(content)

	if content == "" {
		content = getRandomFact()
	}

	for receiverID := range users {
		if receiverID != senderID {
			messages = append(messages, Message{SenderID: senderID, ReceiverID: receiverID, Content: content})
			users[receiverID].Log = append(users[receiverID].Log, fmt.Sprintf("User %d received broadcast message from User %d: %s", receiverID, senderID, content))
		}
	}

	fmt.Println("Broadcast message sent successfully!")
}

// Function to view message log of a user
func viewMessageLog(reader *bufio.Reader) {
	fmt.Print("Enter user ID: ")
	userIDStr, _ := reader.ReadString('\n')
	userIDStr = strings.TrimSpace(userIDStr)
	userID, _ := strconv.Atoi(userIDStr)

	if user, ok := users[userID]; ok {
		fmt.Printf("Message Log of User %d:\n", userID)
		for _, msg := range user.Log {
			fmt.Println(msg)
		}
	} else {
		fmt.Println("User ID does not exist.")
	}
}

// Function to fetch a random fact from an API
func getRandomFact() string {
	url := "https://catfact.ninja/fact"
	resp, err := http.Get(url)
	if err != nil {
		return "Failed to fetch random fact"
	}
	defer resp.Body.Close()

	var factResp struct {
		Fact string `json:"fact"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&factResp); err != nil {
		return "Failed to decode random fact response"
	}

	return factResp.Fact
}
