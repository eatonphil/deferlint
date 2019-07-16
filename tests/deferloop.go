package main

func f() {}

func main() {
	for {
		defer f()
	}
}
