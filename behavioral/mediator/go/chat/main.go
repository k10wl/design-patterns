package main

import "fmt"

type ChatMediator interface {
	join(*User)
	broadcast(user *User, message string)
	private(user *User, name string, message string)
}

type User struct {
	name string
	room ChatMediator
	logs []struct {
		from    string
		message string
	}
}

func (u *User) shout(m string) {
	u.room.broadcast(u, m)
}

func (u *User) whisper(to string, m string) {
	u.room.private(u, to, m)
}

func (u *User) join() {
	u.room.join(u)
}

func (u *User) receive(from string, m string) {
	u.logs = append(u.logs, struct {
		from    string
		message string
	}{from, m})
}

func (u *User) changeRoom(room ChatMediator) {
	// other rooms can be hot swapped without component logic changes
	u.room = room
}

func main() {
	fmt.Printf(">>> Mediator demo\n\n")
	chatRoom := ChatRoom{users: map[string]*User{}}
	kyle := User{name: "Kyle", room: chatRoom}
	stan := User{name: "Stan", room: chatRoom}
	kenny := User{name: "Kenny", room: chatRoom}
	kyle.join()
	stan.join()
	kyle.shout("Hi, Stan!")
	stan.shout("Oh, hi, Kyle")
	kenny.join()
	kenny.shout("jaklsdfjalksdfjkl")
	kenny.whisper("Kyle", "jaklsdfjalksdfjkl")
	kyle.shout("wtf")
	stan.whisper("Cartman", "Are you here?")
	chatRoom.showLogs()
}
