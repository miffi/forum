package main

// Application specific global state. This is used to shuttle around any
// information that is needed in multiple parts of the running program.
type application struct{}

func main() {
	app := &application{}
	router := app.router()
	router.Run()
}
