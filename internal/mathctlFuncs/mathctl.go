package mathctlFuncs

import (
	"fmt"

	"github.com/TanmoySG/go-lib-cli-poc/internal/op"
	"github.com/urfave/cli/v2"
)

func Add(c *cli.Context) error {
	num1, num2 := c.Int("n1"), c.Int("n2")

	n := op.Add(num1, num2)
	fmt.Println(n)

	return nil
}


func Sub(c *cli.Context) error {
	num1, num2 := c.Int("n1"), c.Int("n2")

	n := op.Sub(num1, num2)
	fmt.Println(n)

	return nil
}