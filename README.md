# Contact Directory Manager

A simple command-line contact directory manager written in Go.

## Features

- List all contacts
- Add new contacts
- Delete contacts
- Edit existing contacts
- Check if a contact exists
- Data persistence using JSON files

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

## Running Tests

```bash
go test ./contacts
``` 