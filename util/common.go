package util

import "log"

func Attempt(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
