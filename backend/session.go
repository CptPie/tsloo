package backend

import (
	"log"
	"math/rand"
	"net/http"

	database "github.com/cptpie/tsloo/data"
	"github.com/cptpie/tsloo/logging"
	"github.com/cptpie/tsloo/models"
	"github.com/gorilla/websocket"
)

type Session struct {
	Owner     *models.User
	Clients   map[*websocket.Conn]bool
	Broadcast chan PlaybackState
	Id        string
	upgrader  *websocket.Upgrader
	State     PlaybackState
	Current   database.Entry
	List      []database.Entry
	Log       *logging.Logger
}

type PlaybackState struct {
	Time  float64 `json:"time"`
	State string  `json:"state"` // "playing" or "paused"
	Mode  string  `json:"mode"`  // "youtube" or "mp3"
}

var initialState = PlaybackState{
	Time:  0,
	State: "paused",
	Mode:  "youtube",
}

var lastMessage *PlaybackState

func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func NewSession(owner *models.User, log *logging.Logger, list []database.Entry) (*Session, error) {
	return &Session{
		Owner:     owner,
		Clients:   make(map[*websocket.Conn]bool),
		Broadcast: make(chan PlaybackState),
		upgrader: &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
		Id:    randomString(8),
		State: initialState,
		List:  list,
		// Current: list[0],
	}, nil
}

func (s *Session) playEntry(entry database.Entry) error {
	s.Current = entry
	return nil
}

func (s *Session) handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.Log.Error("Error %v", err)
	}
	defer ws.Close()

	s.Clients[ws] = true

	// Send initial playback mode (YouTube or MP3) to the client
	initialState := PlaybackState{
		Time:  0,
		State: "paused",
		Mode:  "youtube", // Default mode
	}

	if lastMessage != nil {
		ws.WriteJSON(lastMessage)
	} else {
		ws.WriteJSON(initialState)
	}

	for {
		var state PlaybackState
		err := ws.ReadJSON(&state)
		if err != nil {
			log.Printf("Error: %v", err)
			delete(s.Clients, ws)
			break
		}
		s.Broadcast <- state
	}
}

func (s *Session) handleMessages() {
	for {
		state := <-s.Broadcast
		for client := range s.Clients {
			err := client.WriteJSON(state)
			if err != nil {
				log.Printf("Error: %v", err)
				client.Close()
				delete(s.Clients, client)
			}
		}
	}
}
