package server

type Server interface {
	Run()
}

type server struct {
	CORSHeader string
}

func New(CORSHeader string) Server {
	return &server{CORSHeader: CORSHeader}
}

func (*server) Run() {
	r := router()
	r.Run()
}
