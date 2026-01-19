package main



func Identity[T any](a T) T{
	return a
}

func main() {
   Identity(10)
   Identity("Hello")
   Identity(2.67) 
}



   