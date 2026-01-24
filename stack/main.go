package main


type Stack struct{
	data[]int
	size int
	top int
}


func NewStack(size int) *Stack{
	return &Stack {
		data: make([]int,size),
		top:-1,
		size:size,
	}
}


func(s *Stack)Push( x int)bool{
	if s.top==s.size-1{
		return false
	}
	s.top++
	s.data[s.top]=x
	return true
}


func( s *Stack) Pop()(int,bool){
	if s.top==-1{
		return 0,false
	}
	x:=s.data[s.top]
	s.top--
	return x,true
}



func main(){

     s:=NewStack(10)
}