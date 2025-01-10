package backend

import (
	"net/http"
	"os"
	"text/template"

	database "github.com/cptpie/tsloo/data"
	"github.com/cptpie/tsloo/logging"
)

type Backend struct {
	l  *logging.Logger
	db *database.Database
}

type TemplateData struct {
	Youtube   string
	AudioFile string
	Id        string
}

func New(db *database.Database, log *logging.Logger) (*Backend, error) {
	backend := &Backend{
		l:  log,
		db: db,
	}

	http.Handle("/", http.FileServer(http.Dir("./frontend/home.html")))

	session, err := NewSession(nil, backend.l, nil)
	if err != nil {
		return nil, err
	}

	backend.l.Info("New room at: http://localhost:8080/%s", session.Id)

	entry, err := backend.db.GetEntry(1)
	if err != nil {
		return nil, err
	}
	tmpl := template.Must(template.ParseFiles("./frontend/room.html"))

	http.HandleFunc("/"+session.Id, func(w http.ResponseWriter, r *http.Request) {
		data := TemplateData{
			Youtube:   entry.Youtube,
			AudioFile: entry.Mp3,
			Id:        session.Id,
		}

		err := tmpl.Execute(w, data)
		if err != nil {
			backend.l.Error("Template Error %v", err.Error())
		}
	})
	http.HandleFunc("/"+session.Id+"/ws", session.handleConnections)

	go session.handleMessages()

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Error("Failed to start server: %v", err.Error())
		os.Exit(1)
	} else {
		log.Info("Server started at :8080")
	}
	return backend, nil
}
