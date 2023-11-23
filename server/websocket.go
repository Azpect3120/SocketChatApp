package server

import (
	"fmt"
	"sync"

	"github.com/Azpect3120/ChatApp/database"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	// Upgrader for converting HTTP to WS connections
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// Map of connections with nil values
	connections      = make(map[*websocket.Conn]struct{})

	// Mutex ownership controller 
	connectionsMutex sync.Mutex
)

// Open a socket connection on the back-end
func OpenSocket (ctx *gin.Context, db *database.Database) {
	// Convert HTTP connection to WS connection
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println("Unable to convert HTTP connection to WS connection: ", err)
		return
	}

	// Add the connection to the map
	connectionsMutex.Lock()
	connections[conn] = struct{}{}
	connectionsMutex.Unlock()

	// Remove the connection from the map when the client disconnects
	defer func() {
		connectionsMutex.Lock()
		delete(connections, conn)
		connectionsMutex.Unlock()
		conn.Close()
	}()

	// Infinite loop to listen for messages and send them back
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error sending message: ", err)
			break
		}

		broadCastMessage(messageType, message)
		db.AddMessage(message)

	}

}

// Send a message back to each connection
func broadCastMessage(messageType int, message []byte) {
	// Grant mut access to this thread
	connectionsMutex.Lock()
	defer connectionsMutex.Unlock()

	// Broadcast message to each connection
	for conn := range connections {
		if err := conn.WriteMessage(messageType, message); err != nil {
			fmt.Println(err)
		}
	}
}
