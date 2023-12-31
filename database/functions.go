package database

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type MessageData struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Room      string    `json:"room"`
}

func (db *Database) AddMessage(message []byte) {
	var data MessageData

	if err := json.Unmarshal(message, &data); err != nil {
		fmt.Println("Error unmarshaling message to struct: ", err)
	}

	var statement string = "INSERT INTO messages (sender, message, timestamp, room) VALUES ($1, $2, $3, $4);"
	result, err := db.database.Exec(statement, data.Username, data.Message, data.Timestamp, data.Room)
	if err != nil {
		fmt.Println("Error inserting message into database: ", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error converting result to rows: ", err)
	}

	if rows == 0 {
		fmt.Println("No rows were added to the database: ", err)
	}
}

func (db *Database) GetMessages(room string) []MessageData {
	var messages []MessageData

	var statement string = "SELECT * FROM messages WHERE room = $1 ORDER BY timestamp DESC LIMIT 25;"
	rows, err := db.database.Query(statement, room)
	if err != nil {
		fmt.Println("Error getting messages: ", err)
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id        uuid.UUID
			username  string
			message   string
			timestamp time.Time
			room      string
		)

		if err := rows.Scan(&id, &username, &message, &timestamp, &room); err != nil {
			fmt.Println("Error scanning row to variables: ", err)
		}

		messages = append(messages, MessageData{Id: id, Username: username, Message: message, Timestamp: timestamp, Room: room})
	}

	reverse(messages)

	return messages
}

func reverse(arr []MessageData) {
	var length int = len(arr)

	for i := 0; i < length/2; i++ {
		arr[i], arr[length-i-1] = arr[length-i-1], arr[i]
	}
}
