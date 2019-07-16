package main

import "fmt"

// Icecream is
type Icecream interface {
	GetName() string
	HowSweet() string
}

// VanillaIcecream is
type VanillaIcecream struct{}

// GetName returns
func (v *VanillaIcecream) GetName() string {
	return "バニラアイスクリーム"
}

// HowSweet returns
func (v *VanillaIcecream) HowSweet() string {
	return "バニラ味"
}

// GreenTeaIcecream is
type GreenTeaIcecream struct{}

// GetName returns
func (g *GreenTeaIcecream) GetName() string {
	return "抹茶アイスクリーム"
}

// HowSweet returns
func (g *GreenTeaIcecream) HowSweet() string {
	return "抹茶味"
}

// CashewNutsToppingIcecream is
type CashewNutsToppingIcecream struct {
	ice Icecream
}

// GetName returns
func (c *CashewNutsToppingIcecream) GetName() string {
	name := "カシューナッツ" + c.ice.GetName()
	return name
}

// HowSweet returns
func (c *CashewNutsToppingIcecream) HowSweet() string {
	return c.ice.HowSweet()
}

func main() {
	vanilla := &VanillaIcecream{}
	cachVani := &CashewNutsToppingIcecream{ice: vanilla}
	fmt.Println(cachVani.GetName())
	fmt.Println(cachVani.HowSweet())

	green := &GreenTeaIcecream{}
	cachGree := &CashewNutsToppingIcecream{ice: green}
	fmt.Println(cachGree.GetName())
	fmt.Println(cachGree.HowSweet())
}
