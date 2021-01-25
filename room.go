package main

type Room struct {
	commBroadcast chan []byte
	addUser       chan *User
	removeUser    chan *User
	usersIn       map[*User]bool
}

func (r *Room) Run() {
	for {
		select {
		case user := <-r.addUser:
			{
				r.usersIn[user] = true
			}

		case user := <-r.removeUser:
			{
				delete(r.usersIn, user)
				close(user.commSend)
			}

		case msg := <-r.commBroadcast:
			{
				for u := range r.usersIn {
					u.commSend <- msg
				}
			}
		}
	}
}
