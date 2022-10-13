# Go image uploading
## Prerequisites
- **[Go](https://go.dev/)** in your machine
- A **MongoDB** instance

## Setup
```sh
# Clone this repository
git clone https://github.com/Drylozu/GoImages.git

# Build the project
go build main.go
```

## Run it!
```sh
./main
## or (Windows)
main.exe
```

### Options
- `--host` (default: 127.0.0.1:3000). Set the interface to listen.
- `--prefork` (default: false). Enable prefork mode.