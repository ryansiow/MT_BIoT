### Initiate Go API
1. `go mod init main.go`
2. `go mod tidy`

### Compile contracts and generate contract file
- ` abigen --sol={coolNumber}.sol --pkg=main --out={coolNumber}.go`


### Compile for EV3
- `GOOS=linux GOARCH=arm GOARM=5 go build main.go coolNumber.go`

### Execute binary program
- `./main`

### Run the web app
1. `cd coolNumber`
2. `npm i`
3. `node app.js`
