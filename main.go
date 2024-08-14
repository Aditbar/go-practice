package miniChallange3

import (
	"fmt"
	"go-practice/miniChallange3"
	"os"
)

func main() {

	var args = os.Args[1:]

	if len(args) == 0 {
			fmt.Println("Tolong masukan nama atau nomer absen \nContoh: 'go run main.go Fitri' atau 'go run main.go 2'")
	} else {
			miniChallange3.CheckFriend(args)
	}
}