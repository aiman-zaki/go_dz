package services

import (
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//CreateFile : CREATEFILE TO LOCAL DIR
func CreateFile(now time.Time) *os.File {
	formated := now.Format("2006-01-02T15:04:05")
	f, err := os.Create("./logs/scrapper/" + formated)
	check(err)
	return f
}
