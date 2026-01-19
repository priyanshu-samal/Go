package main
import(
	"fmt"
)

type Counter struct{
	s int
}

func Inc(c *Counter){
	c.s++
}

func main(){
     v:=Counter{10}
     //Inc(v)
	 fmt.Println(v)

	 Inc(&v)
	 fmt.Println(v)

}