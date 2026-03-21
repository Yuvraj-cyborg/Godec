package main 

import ( 
	"github.com/Yuvraj-cyborg/godec/internal/codec"
	"fmt"
)


func main(){
	var path = "./assets/red.png"
	err := codec.Encode(path)	
	fmt.Println(err)
}


