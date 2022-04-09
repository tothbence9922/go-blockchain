package content

import "fmt"

type Content struct {
	Value int
}

func (c Content) getValue() int {
	return c.Value
}

func (c Content) String() string {
	return fmt.Sprint("%d", c.Value)
}
