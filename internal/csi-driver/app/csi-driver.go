package app

import "time"

func Start() {
	for {
		print("Hello world\n")
    time.Sleep(15 * time.Second)
	}
}
