package main

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
        .form-group {
            margin-bottom: 15px;
        }
        .form-group label {
            display: block;
            margin-bottom: 5px;
        }
        .form-group input {
            width: 100%;
            padding: 8px;
            border: 1px solid #ddd;
            border-radius: 4px;
        }
        .btn {
            padding: 8px 16px;
            border: none;
            border-radius: 4px;
            cursor: pointer;
            font-size: 14px;
        }
        .btn-primary {
            background: #007bff;
            color: white;
        }
        .btn-danger {
            background: #dc3545;
            color: white;
        }
        .btn-warning {
            background: #ffc107;
            color: black;
        }
        .actions {
            display: flex;
            gap: 10px;
        }
        .modal {
            display: none;
            position: fixed;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            background: rgba(0,0,0,0.5);
        }
        .modal-content {
            background: white;
            margin: 15% auto;
            padding: 20px;
            width: 50%;
            border-radius: 5px;
        }
        .close {
            float: right;
            cursor: pointer;
            font-size: 20px;
        }
    </style>
</head>
<body>
    <h1>Contact Directory</h1>
    
    <button class="btn btn-primary" onclick="showAddModal()">Add New Contact</button>
    
    <div id="contacts">
        {{range .}}
        <div class="contact-item" id="contact-{{.Name}}">
            <div class="contact-info">
                <div class="contact-name">{{.Name}}</div>
                <div class="contact-phone">{{.Phone}}</div>
            </div>
            <div class="actions">
                <button class="btn btn-warning" onclick="showEditModal('{{.Name}}', '{{.Phone}}')">Edit</button>
                <button class="btn btn-danger" onclick="deleteContact('{{.Name}}')">Delete</button>
            </div>
        </div>
        {{end}}
    </div>

    <!-- Add Contact Modal -->
    <div id="addModal" class="modal">
        <div class="modal-content">
            <span class="close" onclick="hideAddModal()">&times;</span>
            <h2>Add New Contact</h2>
            <form id="addForm" onsubmit="addContact(event)">
                <div class="form-group">
                    <label for="addName">Name:</label>
                    <input type="text" id="addName" required>
                </div>
                <div class="form-group">
                    <label for="addPhone">Phone:</label>
                    <input type="tel" id="addPhone" required>
                </div>
                <button type="submit" class="btn btn-primary">Add Contact</button>
            </form>
        </div>
    </div>

    <!-- Edit Contact Modal -->
    <div id="editModal" class="modal">
        <div class="modal-content">
            <span class="close" onclick="hideEditModal()">&times;</span>
            <h2>Edit Contact</h2>
            <form id="editForm" onsubmit="updateContact(event)">
                <input type="hidden" id="editOldName">
                <div class="form-group">
                    <label for="editName">Name:</label>
                    <input type="text" id="editName" required>
                </div>
                <div class="form-group">
                    <label for="editPhone">Phone:</label>
                    <input type="tel" id="editPhone" required>
                </div>
                <button type="submit" class="btn btn-primary">Update Contact</button>
            </form>
        </div>
    </div>

    <script>
        function showAddModal() {
            document.getElementById('addModal').style.display = 'block';
        }

        function hideAddModal() {
            document.getElementById('addModal').style.display = 'none';
        }

        function showEditModal(name, phone) {
            document.getElementById('editOldName').value = name;
            document.getElementById('editName').value = name;
            document.getElementById('editPhone').value = phone;
            document.getElementById('editModal').style.display = 'block';
        }

        function hideEditModal() {
            document.getElementById('editModal').style.display = 'none';
        }

        async function addContact(event) {
            event.preventDefault();
            const name = document.getElementById('addName').value;
            const phone = document.getElementById('addPhone').value;

            const response = await fetch('/api/contacts', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ name, phone }),
            });

            if (response.ok) {
                location.reload();
            } else {
                alert('Error adding contact');
            }
        }

        async function updateContact(event) {
            event.preventDefault();
            const oldName = document.getElementById('editOldName').value;
            const newName = document.getElementById('editName').value;
            const newPhone = document.getElementById('editPhone').value;

            const response = await fetch('/api/contacts/' + encodeURIComponent(oldName), {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ name: newName, phone: newPhone }),
            });

            if (response.ok) {
                location.reload();
            } else {
                alert('Error updating contact');
            }
        }

        async function deleteContact(name) {
            if (!confirm('Are you sure you want to delete this contact?')) {
                return;
            }

            const response = await fetch('/api/contacts/' + encodeURIComponent(name), {
                method: 'DELETE',
            });

            if (response.ok) {
                location.reload();
            } else {
                alert('Error deleting contact');
            }
        }
    </script>
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
