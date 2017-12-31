package main

type Packet struct {
	Type string
	Data interface{}
}

type Router struct {
	handlers map[string]EventHandler
	chats    map[string]*Chat
	users    map[int]*UserConn
}

type EventHandler func(*Router, *UserConn, *Packet)

func NewRouter() *Router {
	r := &Router{
		make(map[string]EventHandler),
		make(map[string]*Chat),
		make(map[int]*UserConn),
	}
	registerHandlers(r)
	return r
}

func (r *Router) NewChat(name string) {
	c := NewChat(name)
	r.chats[name] = c
}

func (r *Router) Handle(e string, eh EventHandler) {
	r.handlers[e] = eh
}

func (r *Router) RoutePacket(uc *UserConn, p *Packet) {
	if h, f := r.handlers[p.Type]; f {
		h(r, uc, p)
	}
}

func (r *Router) DisconnectUser(uid int) {
	if uc, f := r.users[uid]; f {
		uc.Conn.Close()
	}
}
