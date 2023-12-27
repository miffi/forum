package server

func Run() {
	r := router()
	r.Run()
}
