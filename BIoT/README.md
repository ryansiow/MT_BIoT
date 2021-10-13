### Initiate Go API
1. `go mod init ev3-iot.go`
2. `go mod tidy`

### Compile contracts and generate contract file
- ` abigen --sol={ev3SpeedMonitor}.sol --pkg=main --out={ev3SpeedMonitor}.go`


### Compile for EV3
- `GOOS=linux GOARCH=arm GOARM=5 go build ev3-iot.go ev3SpeedMonitor.go`

### Execute binary program
- `./ev3-iot`
