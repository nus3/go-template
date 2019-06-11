package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// TODO: SoupFactoryをSoupに変更する
type AbstractFactory interface {
	GetSoup() SoupFactory
	GetMain() MainFactory
	GetVegetables() VegetablesFactory
	GetOtherIngredients() OtherIngredientsFactory
}

type SoupFactory struct {
	name string
}

type MainFactory struct {
	name string
}

type VegetablesFactory struct {
	vegetables []string
}

type OtherIngredientsFactory struct {
	otheringredients []string
}

type MizutakiFactory struct{}

func (MizutakiFactory) GetSoup() SoupFactory {
	return SoupFactory{"昆布だし"}
}

func (MizutakiFactory) GetMain() MainFactory {
	return MainFactory{"鶏肉"}
}

func (MizutakiFactory) GetVegetables() VegetablesFactory {
	return VegetablesFactory{[]string{"ねぎ", "白菜", "大根"}}
}

func (MizutakiFactory) GetOtherIngredients() OtherIngredientsFactory {
	return OtherIngredientsFactory{[]string{"豆腐", "こんにゃく"}}
}

type HotPotFactroy struct{}

// Create returns HotPot instance.
func (*HotPotFactroy) Create(kind string) AbstractFactory {
	switch kind {
	case "水炊き":
		return MizutakiFactory{}
	default:
		return MizutakiFactory{}
	}
}

type HotPot struct {
	soup             SoupFactory
	main             MainFactory
	vegetables       VegetablesFactory
	otherIngredients OtherIngredientsFactory
}

// TODO: addに渡す引数はSoup構造体のみにする
func (h *HotPot) addSoup(a AbstractFactory) {
	h.soup = a.GetSoup()
}

func (h *HotPot) addMain(a AbstractFactory) {
	h.main = a.GetMain()
}

func (h *HotPot) addVegetables(a AbstractFactory) {
	h.vegetables = a.GetVegetables()
}

func (h *HotPot) addOtherIngredients(a AbstractFactory) {
	h.otherIngredients = a.GetOtherIngredients()
}

func main() {
	factory := &HotPotFactroy{}
	hotpotFactory := factory.Create("水炊き")
	hotpot := &HotPot{}
	hotpot.addSoup(hotpotFactory)
	hotpot.addMain(hotpotFactory)
	hotpot.addVegetables(hotpotFactory)
	hotpot.addOtherIngredients(hotpotFactory)

	fmt.Println(*hotpot)
}

// TODO: これはこれで保存しとく
// // Animal is an interface of cry of animal
// type Animal interface {
// 	Sing() string
// }

// // Cat is one kind of animal.
// type Cat struct{}

// // Sing returns the voice of animal.
// func (Cat) Sing() string {
// 	return "にゃんにゃん"
// }

// // Dog is one kind of animal.
// type Dog struct{}

// // Sing returns the voice of animal.
// func (Dog) Sing() string {
// 	return "バウバウ"
// }

// // AnimalFactory is factory to create instance.
// type AnimalFactory struct{}

// // Create returns Animal instance.
// func (*AnimalFactory) Create(kind string) Animal {
// 	switch kind {
// 	case "dog":
// 		return Dog{}
// 	case "cat":
// 		return Cat{}
// 	default:
// 		return Dog{}
// 	}
// }

// helper

func readLine(sc *bufio.Scanner) string {
	sc.Scan()
	return strings.TrimSpace(sc.Text())
}

func readLineToInt(sc *bufio.Scanner) int {
	intValue, err := strconv.Atoi(readLine(sc))
	if err != nil {
		panic(err)
	}

	return intValue
}

func explodeString(delimiter string, inputValue string) []string {
	splittedValue := strings.Split(inputValue, delimiter)

	var trimmedValues []string

	for _, value := range splittedValue {
		if value != "" {
			trimmedValues = append(trimmedValues, value)
		}
	}

	return trimmedValues
}

func explodeInt(delimiter string, inputValue string) []int {
	inputStrings := explodeString(" ", inputValue)

	var splittedInts []int

	for _, inputString := range inputStrings {
		var (
			intValue int
			err      error
		)
		intValue, err = strconv.Atoi(inputString)

		if err != nil {
			panic(err)
		}
		splittedInts = append(splittedInts, intValue)
	}

	return splittedInts
}
