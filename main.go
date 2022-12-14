package main

func main() {
	r := InitServer()
	r.Run("127.0.0.1:3030")
}
