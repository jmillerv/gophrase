
# Builds all binarys
build-all: build-linux build-osx build-windows

# Builds Linux executable
build-linux:
    cd ./src && GOOS=linux GOARCH=amd64 go build -o bin/gophrase-amd64-linux main.go

# Builds Windows executable
build-windows:
    cd ./src && GOOS=windows GOARCH=amd64 go build -o bin/gophrase-amd64.exe main.go

# Builds MacOS executable
build-osx:
    cd ./src && GOOS=darwin GOARCH=amd64 go build -o bin/gophrase-386-darwin main.go