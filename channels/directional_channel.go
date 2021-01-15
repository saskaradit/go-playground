package main

func main() {

	ch := make(chan int)
	cr := make(<-chan int) //receive
	cs := make(chan<- int) //send

}
