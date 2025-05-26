# Contact Directory Manager

A simple command-line contact directory manager written in Go.

## Features

- List all contacts
- Add new contacts
- Delete contacts
- Edit existing contacts
- Check if a contact exists
- Data persistence using JSON files
- Web interface with full CRUD operations

## Building the Application

```bash
go build
```

## Usage

### List all contacts
```bash
./go-tp -list
```

### Add a new contact
```bash
./go-tp -add -name "John Doe" -phone "0612345678"
```

### Delete a contact
```bash
./go-tp -delete -name "John Doe"
```

### Edit a contact
```bash
./go-tp -edit -name "John Doe" -new-phone "0612345679"
```

You can also change the name:
```bash
./go-tp -edit -name "John Doe" -new-name "Jonathan Doeson"
```

### Check if a contact exists
```bash
./go-tp -exists -name "John Doe"
```

### Specify a different data file
```bash
./go-tp -file "contacts.json" -list
```

### Start Web Server
Start the web server on default port (8080):
```bash
./go-tp -serve
```

Start the web server on a custom port:
```bash
./go-tp -serve -port 3000
```

Once the server is running, you can:
- View contacts in your browser at `http://localhost:3000` (or your custom port)
- Access the JSON API at `http://localhost:3000/api/contacts`
- Use the web interface to:
  - Add new contacts with the "Add New Contact" button
  - Edit existing contacts with the "Edit" button
  - Delete contacts with the "Delete" button
  - View all contacts in a clean, modern interface

## Running Tests

```bash
go test ./contacts
``` 