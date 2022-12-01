# Master Go (Golang) Programming:The Complete Go Bootcamp 2023 -  Udemy course

Snippets of code and notes

## TIPS

```
go env
go run .
go run main.go
go run -x main.go
go build
go build -o app
gofmt -w main.go
gofmt -w -l dir/
go fmt
```

### Cross compiling
```
# On windows to Linux
GOOS=linux GOARCH=amd64 go build -o linuxapp
# On Linux to Windows
GOOS=windows GOARCH=amd64 go build -o linuxapp.exe
```

## Docs
https://go.dev/ref/spec  
https://go.dev/doc/effective_go  
https://go.dev/doc/  
https://pkg.go.dev/std  

## Useful sites
https://go.dev/play/  
https://gobyexample.com/  
https://pkg.go.dev/  
https://goplay.tools/


## Name conventions (idiomatic go)
* Use camel case
* Does not use undescore
* Avoid long names in shrt scopes
* Long names in wider/global scope is ok
* RULE: First letter in upper case exports the name
* Interface names uses "er" sufix 
* Acronyms should be all caps
  * writeToDB - OK
  * writeToDb - Db is acronym to Data Base, should be DB (all caps)