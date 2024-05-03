# CLI Messaging App 
## Assignment by MyMandap - Submitted by PritpalSingh(aka ppcodes)

This is a simple command-line interface (CLI) messaging application written in Go. It allows users to send messages between each other, broadcast messages to all users, and view message logs.

## Key Components

### User Struct
The `User` struct represents a user with a unique ID and their message log.

```go
type User struct {
    ID  int
    Log []string
}
```

### Message Struct
The `Message` struct represents a message with sender ID, receiver ID, and content.

```go
type Message struct {
    SenderID   int
    ReceiverID int
    Content    string
}
```

### Sending a Message
The `sendMessage` function allows users to send messages between two users. It validates user inputs for sender ID, receiver ID, and message content. If the message content is empty, it fetches a random fact from an API.

```go
func sendMessage(reader *bufio.Reader) {
    // Input validation and error handling
}
```

### Broadcasting a Message
The `broadcastMessage` function allows users to broadcast a message to all users except the sender. It validates the sender ID and message content.

```go
func broadcastMessage(reader *bufio.Reader) {
    // Input validation and error handling
}
```

### Viewing Message Logs
The `viewMessageLog` function allows users to view the message log of a specific user. It validates the user ID and prints the message log.

```go
func viewMessageLog(reader *bufio.Reader) {
    // Input validation and error handling
}
```

### Fetching Random Facts
The `getRandomFact` function fetches a random fact from an API (`https://catfact.ninja/fact`). It handles errors during the API call and decoding of the response. Implementing this was the favourite part of the assignment.

```go
func getRandomFact() string {
    url := "https://catfact.ninja/fact"
	resp, err := http.Get(url)
	if err != nil {
		return "Failed to fetch random fact"
	}
	defer resp.Body.Close()

	var factResp map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&factResp); err != nil {
		return "Failed to decode random fact response"
	}

	fact, ok := factResp["fact"]
	if !ok {
		return "Failed to get random fact"
	}

	return fact
}
```

## Error Handling
The application performs error handling and input validation in various parts:
- Validating user inputs for sender ID, receiver ID, and user ID.
- Handling errors during conversion of user inputs to integers.
- Handling errors during HTTP requests to fetch random facts from the API.
- Handling errors during decoding of the JSON response from the API.
  

## Demo



If you liked this assignment, you can also check out my other assignments on [my GitHub](https://github.com/pritpal-singh) or on my [website](https://pritpalsingh.in).