package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Chahine-tech/go-tp/contacts"
)

func main() {
	listCmd := flag.Bool("list", false, "List all contacts")
	addCmd := flag.Bool("add", false, "Add a new contact")
	deleteCmd := flag.Bool("delete", false, "Delete a contact")
	editCmd := flag.Bool("edit", false, "Edit an existing contact")
	existsCmd := flag.Bool("exists", false, "Check if a contact exists")
	serveCmd := flag.Bool("serve", false, "Start web server")

	name := flag.String("name", "", "Name")
	phone := flag.String("phone", "", "Phone number")

	newName := flag.String("new-name", "", "New name (for edit)")
	newPhone := flag.String("new-phone", "", "New phone number (for edit)")

	dataFile := flag.String("file", "contacts.json", "Data file to store contacts")
	port := flag.String("port", "8080", "Port for web server")

	flag.Parse()

	absDataFile, err := filepath.Abs(*dataFile)
	if err != nil {
		fmt.Printf("Error resolving data file path: %v\n", err)
		os.Exit(1)
	}

	dir := contacts.NewDirectory(absDataFile)

	switch {
	case *serveCmd:
		startServer(absDataFile, *port)
	case *listCmd:
		listContacts(dir)
	case *addCmd:
		addContact(dir, *name, *phone)
	case *deleteCmd:
		deleteContact(dir, *name)
	case *editCmd:
		editContact(dir, *name, *newName, *newPhone)
	case *existsCmd:
		checkExists(dir, *name)
	default:
		flag.Usage()
	}
}

func startServer(dataFile string, port string) {
	server, err := NewServer(dataFile)
	if err != nil {
		log.Fatalf("Error creating server: %v", err)
	}

	http.HandleFunc("/", server.handleIndex)
	http.HandleFunc("/api/contacts", server.handleAPI)
	http.HandleFunc("/api/contacts/", server.handleAPI)

	addr := ":" + port
	fmt.Printf("Starting server on http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func listContacts(dir *contacts.Directory) {
	contacts := dir.List()
	if len(contacts) == 0 {
		fmt.Println("No contacts found.")
		return
	}

	fmt.Println("Contacts:")
	for i, c := range contacts {
		fmt.Printf("%d. %s: %s\n", i+1, c.Name, c.Phone)
	}
}

func addContact(dir *contacts.Directory, name, phone string) {
	if name == "" {
		fmt.Println("Error: Name is required.")
		return
	}

	if phone == "" {
		fmt.Println("Error: Phone number is required.")
		return
	}

	contact := contacts.Contact{
		Name:  name,
		Phone: phone,
	}

	err := dir.Add(contact)
	if err != nil {
		fmt.Printf("Error adding contact: %v\n", err)
		return
	}

	fmt.Printf("Contact %s added successfully.\n", name)
}

func deleteContact(dir *contacts.Directory, name string) {
	if name == "" {
		fmt.Println("Error: Name is required.")
		return
	}

	err := dir.Delete(name)
	if err != nil {
		fmt.Printf("Error deleting contact: %v\n", err)
		return
	}

	fmt.Printf("Contact %s deleted successfully.\n", name)
}

func editContact(dir *contacts.Directory, name, newName, newPhone string) {
	if name == "" {
		fmt.Println("Error: Name is required.")
		return
	}

	// Find the contact first
	matches := dir.FindByName(name)
	if len(matches) == 0 {
		fmt.Printf("Error: Contact %s not found.\n", name)
		return
	}

	contact := matches[0]

	if newName != "" {
		contact.Name = newName
	}
	if newPhone != "" {
		contact.Phone = newPhone
	}

	err := dir.Edit(name, contact)
	if err != nil {
		fmt.Printf("Error updating contact: %v\n", err)
		return
	}

	fmt.Printf("Contact updated successfully.\n")
}

func checkExists(dir *contacts.Directory, name string) {
	if name == "" {
		fmt.Println("Error: Name is required.")
		return
	}

	contact := contacts.Contact{
		Name: name,
	}

	exists := dir.ContactExists(contact)
	if exists {
		fmt.Printf("Contact %s exists in the directory.\n", name)
		return
	}

	fmt.Printf("Contact %s does not exist in the directory.\n", name)
}
