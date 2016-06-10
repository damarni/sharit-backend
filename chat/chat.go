package chat

import (
	"encoding/json"
	"log"
	"net/http"
	"sharit-backend/models"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/googollee/go-socket.io"
)

func Run() {
	server,
		err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		so.Emit("connection")

		so.On("setRoom", func(data string) {
			log.Println("setRoom")
			d, err := simplejson.NewJson([]byte(data))
			if err != nil {
				log.Println(err)
			} else {
				userId := d.Get("userId").MustString()
				roomId := d.Get("roomId").MustString()

				log.Println("userId:", userId)
				log.Println("roomId:", roomId)
				so.Join(roomId)

			}
		})

		so.On("newMessage", func(data string) {
			log.Println("newMessage")
			d, err := simplejson.NewJson([]byte(data))
			if err != nil {
				log.Println(err)
			} else {
				userId := d.Get("userId").MustString()
				roomId := d.Get("roomId").MustString()
				message := d.Get("message").MustString()

				time := time.Now().UTC().Format(time.RFC3339Nano)
				var msg models.Message
				msg.UserId = userId
				msg.Text = message
				msg.Date = time
				room, err := models.FindRoom(roomId)
				if err != nil {
					log.Println(err)
				} else {
					log.Println(userId, message, time)
					err = room.PutMessage(msg)
					if err != nil {
						log.Println(err)
					} else {
						newData := map[string]interface{}{
							"userId":  userId,
							"message": message,
							"date":    time,
						}
						json, err := json.Marshal(newData)
						if err != nil {
							log.Println(err)
						} else {
							log.Println(string(json))
							so.Emit("newMessage", string(json))
							so.BroadcastTo(roomId, "newMessage", string(json))
							log.Println("newMessageFinished")
						}
					}
				}
			}
		})

		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.HandleFunc("/socket.io/", func(w http.ResponseWriter, r *http.Request) {
		/*w.Header().Set("Access-Control-Allow-Origin", "*")//"http://localhost:8100")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		server.ServeHTTP(w, r)*/
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers",
				"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		}
		// Stop here if its Preflighted OPTIONS request
		if r.Method == "OPTIONS" {
			return
		}
		server.ServeHTTP(w, r)
	})

	//http.Handle("/socket.io/", server)
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
