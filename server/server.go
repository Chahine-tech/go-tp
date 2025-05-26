package server

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"

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
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	
	tmpl, err := template.ParseFiles("server/templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	contacts := s.dir.List()
	if err := tmpl.Execute(w, contacts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) handleAPI(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/contacts")

	switch {
	case path == "":
		// Handle /api/contacts
		switch r.Method {
		case "GET":
			contacts := s.dir.List()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(contacts)
		case "POST":
			var contact contacts.Contact
			if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if err := s.dir.Add(contact); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	case strings.HasPrefix(path, "/"):
		// Handle /api/contacts/:name
		name := strings.TrimPrefix(path, "/")
		switch r.Method {
		case "PUT":
			var contact contacts.Contact
			if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if err := s.dir.Edit(name, contact); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
		case "DELETE":
			if err := s.dir.Delete(name); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	default:
		http.NotFound(w, r)
	}
}
