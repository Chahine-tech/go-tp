package contacts

import "fmt"

type Contact struct {
	Name string `json:"name"`
	Phone string `json:"phone"`
}

func (c Contact) FullName() string {
	return c.Name
}

func (c Contact) String() string {
	return fmt.Sprintf("%s: %s", c.Name, c.Phone)
}
