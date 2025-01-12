package main

func main() {
	server := ServeAPI(":3001")
	server.Run()
}
