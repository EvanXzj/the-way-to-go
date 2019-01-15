package main

var a string

func main() {
	a = "G"
	print(a) // G
	f1()
}

func f1() {
	a := "O"
	print(a) // O

	f2() // G   函数调用， a取的是全局变量
	func() {
		print(a) // O
	}()
}

func f2() {
	print(a)
}
