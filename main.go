package main

func main() {
	Test()
	// core.Run()
}

func Test() {
	var a = make(map[string]string)
	_, ok := a["hello"]
	a["a"] = "b"
	println(ok, a["a"])
}
