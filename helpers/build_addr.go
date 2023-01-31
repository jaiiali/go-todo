package helpers

import (
	"fmt"
	"os"
)

func BuildAddr() string {
	host := os.Getenv("HTTP_HOST")
	port := os.Getenv("HTTP_PORT")

	return fmt.Sprintf("%s:%s", host, port)
}
