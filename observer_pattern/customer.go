package observer

import "fmt"

type Customer struct {
	Id string
}

func (c *Customer) update(itemName string) {
	fmt.Printf("Sending Email to Customer %s for item  %s\n", c.Id, itemName)
}

func (c *Customer) getID() string {
	return c.Id
}
