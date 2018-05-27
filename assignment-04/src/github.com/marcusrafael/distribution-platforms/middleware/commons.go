package middleware

import (
	"log"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panic(err, msg)
	}
}
