package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"unicode"
)

type Product struct {
	name  string
	price float64 // TL
}

type CartItem struct {
	product  Product
	quantity int
}

func (ci CartItem) Total() float64 {
	return ci.product.price * float64(ci.quantity)
}

type Cart struct {
	items map[string]*CartItem
}

func NewCart() *Cart {
	return &Cart{items: make(map[string]*CartItem)}
}

var ErrInvalidQuantity = errors.New("quantity must be > 0")

type ErrProductNotFound struct {
	Name string
}

func (e ErrProductNotFound) Error() string {
	return fmt.Sprintf("product %s not found in cart", e.Name)
}

func validateName(name string) bool {
	for _, r := range name {
		if !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == ' ' || r == '-' || r == '_') {
			return false
		}
	}
	return len(strings.TrimSpace(name)) > 0
}

func (c *Cart) AddItem(product Product, quantity int) error {
	if !validateName(product.name) {
		return fmt.Errorf("invalid product name: %q", product.name)
	}
	if quantity <= 0 {
		return ErrInvalidQuantity
	}
	if item, exists := c.items[product.name]; exists {
		item.quantity += quantity
		return nil
	}
	c.items[product.name] = &CartItem{product: product, quantity: quantity}
	return nil
}

func (c *Cart) RemoveItem(name string) error {
	if _, exists := c.items[name]; exists {
		delete(c.items, name)
		return nil
	}
	return ErrProductNotFound{Name: name}
}

func (c *Cart) UpdateQuantity(name string, newQty int) error {
	item, exists := c.items[name]
	if !exists {
		return ErrProductNotFound{Name: name}
	}
	if newQty < 0 {
		return ErrInvalidQuantity
	}
	if newQty == 0 {
		delete(c.items, name)
		return nil
	}
	item.quantity = newQty
	return nil
}

func (c *Cart) Clear() {
	c.items = make(map[string]*CartItem)
}

func (c *Cart) Discount(percent float64) {
	if percent <= 0 {
		return
	}
	for _, it := range c.items {
		it.product.price = it.product.price * (1.0 - percent/100.0)
	}
}

// Total: sepet toplamını hesaplar
func (c *Cart) Total() float64 {
	var sum float64
	for _, it := range c.items {
		sum += it.Total()
	}
	return sum
}

func (c Cart) String() string {
	if len(c.items) == 0 {
		return "Your cart is empty"
	}

	names := make([]string, 0, len(c.items))
	for name := range c.items {
		names = append(names, name)
	}
	sort.Strings(names)

	var b strings.Builder
	b.WriteString("Cart Items:\n")
	for _, name := range names {
		it := c.items[name]
		b.WriteString(fmt.Sprintf("%s: %d x %.2f TL = %.2f TL\n",
			it.product.name, it.quantity, it.product.price, it.Total()))
	}
	b.WriteString("-------------------------\n")
	b.WriteString(fmt.Sprintf("Total: %.2f TL\n", c.Total()))
	return b.String()
}

func main() {
	cart := NewCart()

	apple := Product{name: "apple", price: 15}
	coffee := Product{name: "coffee", price: 25}
	laptop := Product{name: "laptop", price: 15000}
	invalid := Product{name: "test", price: 1000}

	must(cart.AddItem(apple, 2))
	must(cart.AddItem(coffee, 1))
	must(cart.AddItem(laptop, 1))

	if err := cart.AddItem(invalid, 1); err != nil {
		fmt.Println("AddItem error:", err)
	}

	fmt.Println(cart)

	fmt.Println("\nAfter updates:\n" + cart.String())

	cart.Discount(10)
	fmt.Println("\nAfter 10% discount:\n" + cart.String())

	cart.Clear()
	fmt.Println("\nAfter clear:\n" + cart.String())
}

func must(err error) {
	if err != nil {
		fmt.Println("error:", err)
	}
}
