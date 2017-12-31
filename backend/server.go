package main

type Server struct {
	PrivateKey []byte
	Handlers   map[string]*EventHandler
	Users      map[string]*User
	Sessions   map[string]*Session
	Chats      map[string]*Chat
}

func NewServer() (*Server, error) {
	key, err := loadPrivateKey("./private.key")
	if err != nil {
		return nil, err
	}
	return &Server{
		key,
		make(map[string]*EventHandler),
		make(map[string]*User),
		make(map[string]*Session),
		make(map[string]*Chat),
	}, nil
}

func (s *Server) Handle(event string, eh EventHandler) {
	s.Handlers[event] = &eh
}

func (s *Server) HandleEvent(ss *Session, e *Event) {

}

func (s *Server) EndSession(ss *Session) {

}

func (s *Server) NewChat(chat *Chat) {
	s.Chats[chat.Name] = chat
}
