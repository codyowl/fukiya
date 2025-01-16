package main

import (
	"fmt"
	"fukiya/utilities"
)



func main(){
	if utilities.IsKubePresent(){
		fmt.Println("Yes kubectl is present")
	}else{
		fmt.Println("No it is not here")
	}
}