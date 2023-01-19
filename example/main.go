package main

import (
	"fmt"

	"github.com/TanmoySG/go-lib-cli-poc/client"
)

func main() {
	c := client.NewClient(5)
	a := c.Add(6)
	fmt.Println(a)
}
