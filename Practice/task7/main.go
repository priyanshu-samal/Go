package main

type IntHeap []int


func (h IntHeap) len()int{
        return len(h)
}

func (h *IntHeap) add(x any){
	*h=append(*h,x.(int))
}





func main(){

}