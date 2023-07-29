package main

import (
	"fmt"

	"github.com/escape381/traveler/domain/service"
)

func main() {
	s := service.AnyService{}
	fmt.Println(s.AnyFunc())
}
