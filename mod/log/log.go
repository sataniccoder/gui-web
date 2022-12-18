package log

import (
	"fmt"
	"os"
	"time"
)

func Log(text string, file string) {
	// genreate the log
	lo := "[" + time.Now().String() + "] " + text // thats it lol

	err := os.WriteFile("/tmp/dat1", []byte(lo), 0644)
	if err != nil {
		fmt.Println("[!!] ERROR: ", err.Error())
	}
}
