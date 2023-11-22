package main

func add() {}

func main() {
	go add()
	go func() {}()
}
