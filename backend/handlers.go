package main

func registerHandlers(r *Router) {
	r.Handle("new_chat", newChat)
	r.Handle("new_message", newMessage)
	r.Handle("set_user", setUser)
	r.Handle("subscribe_chat", subscribeChat)
}

func newChat(r *Router, u *UserConn, p *Packet) {}

func newMessage(r *Router, u *UserConn, p *Packet) {}

func setUser(r *Router, u *UserConn, p *Packet) {}

func subscribeChat(r *Router, u *UserConn, p *Packet) {}
