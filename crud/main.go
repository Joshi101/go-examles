package main

func main() {
	app := App{}
	app.init()
	app.handleRoutes()
	app.run()
}
