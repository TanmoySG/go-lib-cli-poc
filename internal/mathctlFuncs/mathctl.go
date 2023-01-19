package mathctlFuncs

import (
	"fmt"
	"github.com/TanmoySG/go-lib-cli-poc/internal/op"
	"github.com/urfave/cli/v2"
	"strconv"
)

func Add(ctx *cli.Context) error {
	num1, num2 := "1", "2"
	if ctx.NArg() > 1 {
		num1 = ctx.Args().Get(0)
		num2 = ctx.Args().Get(1)
	}

	n1, err := strconv.Atoi(num1)
	if err != nil {
		fmt.Printf("not a num %s: %s", num1, err)
		return err
	}

	n2, err := strconv.Atoi(num2)
	if err != nil {
		fmt.Printf("not a num %s: %s", num2, err)
		return err
	}

	n := op.Add(n1, n2)
	fmt.Printf("%v", n)

	return nil
}
