package client

import "github.com/TanmoySG/go-lib-cli-poc/internal/op"

type client struct {
	base int
}

type Client interface {
	Add(a int) int
}

func NewClient(base int) Client {
	return client{
		base: base,
	}
}

func (c client) Add(a int) int {
	return op.Add(c.base, a)
}
