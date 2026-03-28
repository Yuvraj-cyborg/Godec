package main 

import ( 
	"github.com/Yuvraj-cyborg/godec/internal/codec"

	"fmt"
)


func main(){
	var path = "./assets/chris.jpeg"
	err := codec.Encode(path)

	fmt.Println(err)
}


