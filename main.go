package main

func main() {
	r := InitServer()
	r.Run(":3030")
}
