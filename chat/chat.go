package chat

import (
    "log"
    "net/http"
    "time"
    "github.com/googollee/go-socket.io"
    "sharit-backend/models"
)

func Run() {
    server,
    err := socketio.NewServer(nil)
    if err != nil {
        log.Fatal(err)
    }
    server.On("connection", func(so socketio.Socket) {
        log.Println("on connection")

        so.On("connect", func(data map[string] interface {}) {
					log.Println("connect")
            //userId, _ := data["userId"].(string)
            roomId, _ := data["roomId"].(string)
            so.Join("roomId")

            room, err := models.FindRoom(roomId)
            if err == nil { //continue

                for message, _ := range room.MessagesRoom {
                    so.Emit("newMessage", message)
                    so.BroadcastTo(roomId, "newMessage", message)
                }
            }
        })

        so.On("newMessage", func(data map[string] interface {}) {
						log.Println("new message")
            userId, _ := data["userId"].(string)
            roomId, _ := data["roomId"].(string)
            message,_ := data["message"].(string)

            var msg models.Message
            msg.UserId = userId
            msg.Text = message
            msg.Date = time.Now().UTC().Format(time.RFC3339Nano)
            room, err := models.FindRoom(roomId)
            if err == nil {
                err = room.PutMessage(msg)
            }

            log.Println("emit:", so.Emit("newMessage", msg))
            so.BroadcastTo(roomId, "newMessage", msg)
        })

        so.On("disconnection", func() {
            log.Println("on disconnect")
        })
    })
    server.On("error", func(so socketio.Socket, err error) {
        log.Println("error:", err)
    })

    http.Handle("/socket.io/", server)
    http.Handle("/", http.FileServer(http.Dir("./asset")))
    log.Println("Serving at localhost:5000...")
    log.Fatal(http.ListenAndServe(":5000", nil))
}


/*var lmMutex sync.Mutex

  sio, err := socketio.NewServer(nil)
	if err != nil {
	   log.Fatal(err)
	}

  sio.On("connection", func(socket socketio.Socket) {
    //var userId string
		socket.On("open", func (data map[string]interface{}){
      //userId, _ := data["userId"].(string)
      roomId, _ := data["roomId"].(string)
			fmt.Println("connect to: " + roomId)
			lmMutex.Lock()
      connect(socket, roomId)
			lmMutex.Unlock()
    })

    socket.On("newMessage", func (data map[string]interface{}) {
			fmt.Println("new message")
			userId, _ := data["userId"].(string)
      roomId, _ := data["roomId"].(string)
			message, _ := data["message"].(string)
			fmt.Println("new message " + userId + " " + roomId + " " + message)
      newMessage(socket, userId, roomId, message);
    })
  })

	sio.On("error", func(so socketio.Socket, err error) {
			log.Println("error:", err)
	})

	// Sets up the handlers and listen on port 80
	http.Handle("/socket.io/", sio)

	// Default to :8080 if not defined via environmental variable.

		var listen = ":8080"

	http.ListenAndServe(listen, nil)
	fmt.Println("hi")
}*/

/*func connect(socket socketio.Socket, roomId string) {
  socket.Join("/"+ roomId)

  room, err := models.FindRoom(roomId)
  if err == nil { //continue

    for i, _ := range room.MessagesRoom {
      socket.Emit("message", i)
			//so.BroadcastTo(websocketRoom, "message", string(jsonRes))
      //socket.Emit("message", room.MessagesRoom[i])
    }
  }
}

func newMessage(socket socketio.Socket, userId string, roomId string, message string) {
  var msg models.Message
	msg.UserId = userId
	msg.Text = message
  msg.Date = time.Now().UTC().Format(time.RFC3339Nano)
	room, err := models.FindRoom(roomId)
  if err == nil { //continue
		err = room.PutMessage(msg)
  }
	so.Emit("message", string(jsonRes))
	so.BroadcastTo(websocketRoom, "message", string(jsonRes))
}*/
