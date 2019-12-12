package calculator

type Cal struct {
	result int
}

func (c *Cal) Add(x int) {
	c.result += x
}

func (c *Cal) Sub(x int) {
	c.result -= x
}

func (c *Cal) MultipleBy(x int) {
	if x == 0 {
		c.result = 0
	} else {
		c.result *= x
	}
}

func (c *Cal) Press(x int) {
	c.result = x
}

func (c *Cal) Clear(x int) {
	c.result = 0
}

func (c *Cal) Result() int {
	return c.result
}
