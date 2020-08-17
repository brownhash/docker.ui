package main

import (
	"./internals"
	"fmt"
)

func main() {
	imageList, err := internals.Images()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(imageList)
}
