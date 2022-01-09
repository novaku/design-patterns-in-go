package framework

type CheeseTopping struct {
	Pizza Pizza
}

func (c *CheeseTopping) GetPrice() int {
	pizzaPrice := c.Pizza.GetPrice()
	return pizzaPrice + 10
}