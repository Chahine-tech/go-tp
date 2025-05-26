package contacts

import (
	"testing"
)

func TestAddAndList(t *testing.T) {
	dir := NewDirectory("")
	
	contact := Contact{
		Name: "John Doe",
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
	
	if contacts[0].Name != "John Doe" || contacts[0].Phone != "0612345678" {
		t.Fatalf("Contact details don't match. Got: %+v", contacts[0])
	}
}

func TestContactExists(t *testing.T) {
	dir := NewDirectory("")
	
	contact := Contact{
		Name: "Jane Smith",
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
		Name: "Bob Johnson",
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
		Name: "Alice Brown",
		Phone: "0612345678",
	}
	
	if err := dir.Add(contact); err != nil {
		t.Fatalf("Failed to add contact: %v", err)
	}
	
	err := dir.Delete("Alice Brown")
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
		Name: "Mark Wilson",
		Phone: "0612345678",
	}
	
	if err := dir.Add(contact); err != nil {
		t.Fatalf("Failed to add contact: %v", err)
	}
	
	updatedContact := Contact{
		Name: "Mark Wilson",
		Phone: "0612345678",
	}
	
	err := dir.Edit("Mark Wilson", updatedContact)
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
	dir := NewDirectory("")
	
	contact := Contact{
		Name: "Test User",
		Phone: "0612345678",
	}
	
	if err := dir.Add(contact); err != nil {
		t.Fatalf("Failed to add contact: %v", err)
	}
	
	contacts := dir.List()
	if len(contacts) != 1 {
		t.Fatalf("Expected 1 contact, got %d", len(contacts))
	}
	
	if contacts[0].Name != "Test User" {
		t.Fatalf("Contact details don't match. Got: %+v", contacts[0])
	}
} 