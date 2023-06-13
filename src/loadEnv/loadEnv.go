package env

import (
	"errors"
	"fmt"
	"net/netip"
	"os"
	"strconv"

	"github.com/NoobforAl/BusinessActor/src/logger"
	"github.com/joho/godotenv"
)

var (
	dsn         string
	listen_IP   string
	listen_PORT int
)

var errNotValidIP = errors.New("ip4 is not valid, ")
var errNotPort = errors.New("Port is not valid, ")

func init() {
	if err := godotenv.Load("./.env"); err != nil {
		logger.Log.Println("No .env file found")
	}

	dsn = os.Getenv("DSN")
	listen_IP = os.Getenv("LISTEN_IP")

	ip, err := netip.ParseAddr(listen_IP)
	if err != nil && !ip.Is4() {
		panic(errors.Join(errNotValidIP, err))
	}

	port := os.Getenv("LISTEN_PORT")
	listen_PORT, err = strconv.Atoi(port)
	if err != nil || listen_PORT < 1 {
		panic(errors.Join(errNotPort, err))
	}
}

func GetDsn() string {
	return dsn
}

func GetAddrListen() string {
	return fmt.Sprintf("%s:%d", listen_IP, listen_PORT)
}
