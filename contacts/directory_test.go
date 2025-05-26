package contacts

import (
	"os"
	"testing"
)

func TestAddAndList(t *testing.T) {
	dir := NewDirectory("")
	
	contact := Contact{
		Name: "John",
		Phone: "0612345678",
	}
	
	err := dir.Add(contact)
	if err != nil {
		t.Fatalf("Failed to add contact: %v", err)
	}
	
	contacts := dir.List()
	if len(contacts) != 1 {
		t.Fatalf("Expected 1 contact, got %d", len(contacts))
	}
	
	if contacts[0].Name != "John" || contacts[0].Phone != "0612345678" {
		t.Fatalf("Contact details don't match. Got: %+v", contacts[0])
	}
}

func TestContactExists(t *testing.T) {
	dir := NewDirectory("")
	
	contact := Contact{
		Name: "Jane",
		Phone: "0612345678",
	}
	
	if err := dir.Add(contact); err != nil {
		t.Fatalf("Failed to add contact: %v", err)
	}
	
	exists := dir.ContactExists(contact)
	if !exists {
		t.Fatalf("Expected contact to exist, but it doesn't")
	}
	
	nonExistingContact := Contact{
		Name: "Bob",
		Phone: "0612345678",
	}
	
	exists = dir.ContactExists(nonExistingContact)
	if exists {
		t.Fatalf("Expected contact not to exist, but it does")
	}
}

func TestDeleteContact(t *testing.T) {
	dir := NewDirectory("")
	
	contact := Contact{
		Name: "Alice",
		Phone: "0612345678",
	}
	
	if err := dir.Add(contact); err != nil {
		t.Fatalf("Failed to add contact: %v", err)
	}
	
	err := dir.Delete("Alice")
	if err != nil {
		t.Fatalf("Failed to delete contact: %v", err)
	}
	
	contacts := dir.List()
	if len(contacts) != 0 {
		t.Fatalf("Expected 0 contacts after deletion, got %d", len(contacts))
	}
}

func TestEditContact(t *testing.T) {
	dir := NewDirectory("")
	
	contact := Contact{
		Name: "Mark",
		Phone: "0612345678",
	}
	
	if err := dir.Add(contact); err != nil {
		t.Fatalf("Failed to add contact: %v", err)
	}
	
	updatedContact := Contact{
		Name: "Mark",
		Phone: "0612345678",
	}
	
	err := dir.Edit("Mark", updatedContact)
	if err != nil {
		t.Fatalf("Failed to edit contact: %v", err)
	}
	
	contacts := dir.List()
	if len(contacts) != 1 {
		t.Fatalf("Expected 1 contact, got %d", len(contacts))
	}
	
	if contacts[0].Phone != "0612345678" {
		t.Fatalf("Phone number was not updated. Expected '0612345678', got '%s'", contacts[0].Phone)
	}
}

func TestPersistence(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "contacts-test-*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())
	tmpfile.Close()
	
	dir := NewDirectory(tmpfile.Name())
	
	contact := Contact{
		Name: "Test",
		Phone: "0612345678",
	}
	
	if err := dir.Add(contact); err != nil {
		t.Fatalf("Failed to add contact: %v", err)
	}
	
	dir2 := NewDirectory(tmpfile.Name())
	
	contacts := dir2.List()
	if len(contacts) != 1 {
		t.Fatalf("Expected 1 contact to be loaded from file, got %d", len(contacts))
	}
	
	if contacts[0].Name != "Test" {
		t.Fatalf("Loaded contact details don't match. Got: %+v", contacts[0])
	}
} 