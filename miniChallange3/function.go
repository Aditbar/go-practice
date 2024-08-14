package miniChallange3

import (
	"fmt"
	"strconv"
)

func CheckFriend(inputs []string) {
	for _, input := range inputs {
			found := false
			for _, friend := range FriendList {
					if strconv.Itoa(friend.ID) == input || friend.Nama == input {
							fmt.Printf("ID: %d\n", friend.ID)
							fmt.Printf("Nama: %s\n", friend.Nama)
							fmt.Printf("Alamat: %s\n", friend.Alamat)
							fmt.Printf("Pekerjaan: %s\n", friend.Pekerjaan)
							fmt.Printf("Alasan: %s\n", friend.Alasan)
							fmt.Println("---------")
							found = true
					}
			}
			if !found {
				fmt.Println("Data dengan nama/absen", input, "tsb tidak tersedia")
					fmt.Println("---------")
			}
	}
}