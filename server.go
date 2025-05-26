package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/Chahine-tech/go-tp/contacts"
)

type Server struct {
	dir *contacts.Directory
}

func NewServer(dataFile string) (*Server, error) {
	absDataFile, err := filepath.Abs(dataFile)
	if err != nil {
		return nil, err
	}
	return &Server{
		dir: contacts.NewDirectory(absDataFile),
	}, nil
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <title>Contact Directory</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            max-width: 800px;
            margin: 0 auto;
            padding: 20px;
        }
        .contact-list {
            list-style: none;
            padding: 0;
        }
        .contact-item {
            background: #f5f5f5;
            margin: 10px 0;
            padding: 15px;
            border-radius: 5px;
            display: flex;
            justify-content: space-between;
            align-items: center;
        }
        .contact-info {
            flex-grow: 1;
        }
        .contact-name {
            font-weight: bold;
            font-size: 1.2em;
        }
        .contact-phone {
            color: #666;
        }
    </style>
</head>
<body>
    <h1>Contact Directory</h1>
    <div id="contacts">
        {{range .}}
        <div class="contact-item">
            <div class="contact-info">
                <div class="contact-name">{{.Name}}</div>
                <div class="contact-phone">{{.Phone}}</div>
            </div>
        </div>
        {{end}}
    </div>
</body>
</html>`

	t, err := template.New("index").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	contacts := s.dir.List()
	if err := t.Execute(w, contacts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) handleAPI(w http.ResponseWriter, r *http.Request) {
	contacts := s.dir.List()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}
