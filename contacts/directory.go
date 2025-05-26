package contacts

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
)

type Directory struct {
	Contacts []Contact `json:"contacts"`
	filename string
	mutex    sync.RWMutex
}

func NewDirectory(filename string) *Directory {
	dir := &Directory{
		Contacts: []Contact{},
		filename: filename,
	}

	if filename != "" {
		dir.Load()
	}
	
	return dir
}

func (d *Directory) Load() error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	_, err := os.Stat(d.filename)
	if os.IsNotExist(err) {
		return nil 
	}

	data, err := os.ReadFile(d.filename)
	if err != nil {
		return err
	}

	if len(data) > 0 {
		return json.Unmarshal(data, &d.Contacts)
	}
	return nil
}

func (d *Directory) Save() error {
	if d.filename == "" {
		return nil
	}

	d.mutex.RLock()
	defer d.mutex.RUnlock()

	data, err := json.MarshalIndent(d.Contacts, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(d.filename, data, 0644)
}

func (d *Directory) List() []Contact {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	
	contacts := make([]Contact, len(d.Contacts))
	copy(contacts, d.Contacts)
	return contacts
}

func (d *Directory) Add(contact Contact) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.contactExistsLocked(contact) {
		return errors.New("contact with this name already exists")
	}

	d.Contacts = append(d.Contacts, contact)
	return d.Save()
}

func (d *Directory) Delete(name string) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	for i, c := range d.Contacts {
		if c.Name == name {
			d.Contacts = append(d.Contacts[:i], d.Contacts[i+1:]...)
			return d.Save()
		}
	}

	return fmt.Errorf("contact '%s' not found", name)
}

func (d *Directory) Edit(oldName string, newContact Contact) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if (oldName != newContact.Name) && 
	   d.contactExistsLocked(newContact) {
		return errors.New("cannot update: a contact with the new name already exists")
	}

	for i, c := range d.Contacts {
		if c.Name == oldName {
			d.Contacts[i] = newContact
			return d.Save()
		}
	}

	return fmt.Errorf("contact '%s' not found", oldName)
}

func (d *Directory) ContactExists(contact Contact) bool {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	
	return d.contactExistsLocked(contact)
}

func (d *Directory) contactExistsLocked(contact Contact) bool {
	for _, c := range d.Contacts {
		if c.Name == contact.Name {
			return true
		}
	}
	return false
}

func (d *Directory) FindByName(name string) []Contact {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	var results []Contact
	for _, c := range d.Contacts {
		if (name == "" || c.Name == name) {
			results = append(results, c)
		}
	}
	return results
} 