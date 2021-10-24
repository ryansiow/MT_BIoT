### Initiate Go API
1. `go mod init main.go`
2. `go mod tidy`

### Compile contracts and generate contract file
- ` abigen --sol={randomNumber}.sol --pkg=main --out={randomNumber}.go`


### Compile for EV3
- `GOOS=linux GOARCH=arm GOARM=5 go build main.go randomNumber.go`

### Execute binary program
- `./main`

### Run the web app
1. `cd randomNumber`
2. `npm i`
3. `node app.js`
