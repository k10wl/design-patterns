package main

import "fmt"

type ChatRoom struct {
	users map[string]*User
}

func (room ChatRoom) join(user *User) {
	fmt.Printf("$ %s has joined the room\n", user.name)
	for _, u := range room.users {
		u.receive("< "+user.name+" (chat)", "joined the room")
	}
	room.users[user.name] = user
}

func (room ChatRoom) broadcast(user *User, message string) {
	for _, val := range room.users {
		if val == user {
			val.receive("> "+user.name+" (glob)", message)
			continue
		}
		val.receive("< "+user.name+" (glob)", message)
	}
}

func (room ChatRoom) private(user *User, to string, message string) {
	if u, ok := room.users[to]; ok {
		u.receive("< "+user.name+" (whisper)", message)
		user.receive("> "+user.name+" to "+to, message)
	}
}

func (room *ChatRoom) showLogs() {
	fmt.Printf("\n> Logs\n")
	for _, u := range room.users {
		fmt.Printf("========\nUser: %s\n", u.name)
		for _, log := range u.logs {
			fmt.Printf("%s: %s\n", log.from, log.message)
		}
	}
}
