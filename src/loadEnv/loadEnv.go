package env

import (
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

func init() {
	if err := godotenv.Load("./.env"); err != nil {
		logger.Log.Println("No .env file found")
	}

	dsn = os.Getenv("DSN")
	listen_IP = os.Getenv("LISTEN_IP")

	ip, err := netip.ParseAddr(listen_IP)
	if err != nil && !ip.Is4() {
		logger.Log.Printf("get error set default value Err: %s", err.Error())
		listen_IP = "0.0.0.0"
	}

	port := os.Getenv("LISTEN_PORT")
	listen_PORT, err = strconv.Atoi(port)
	if err != nil || listen_PORT < 1 {
		logger.Log.Printf("get error set default value Err: %s", err.Error())
		listen_PORT = 8080
	}
}

func GetDsn() string {
	return dsn
}

func GetAddrListen() string {
	return fmt.Sprintf("%s:%d", listen_IP, listen_PORT)
}
