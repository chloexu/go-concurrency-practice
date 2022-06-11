package main

import (
	"fmt"
	"math/rand"
	"time"
)

var woodTypes []string

type Shakeable interface {
	Shake()
}

type Choppable interface {
	Chop()
}

type Diggable interface {
	Dig()
}

type PlantType string

const (
	Tree   PlantType = "Tree"
	Shrub  PlantType = "Shrub"
	Flower PlantType = "Flower"
)

type PlantSize string

const (
	Seed    PlantSize = "Seed"
	Sprout  PlantSize = "Sprout"
	Sapling PlantSize = "Sapling"
	Mature  PlantSize = "Mature"
)

type Plant struct {
	plantType PlantType
	name      string
	hasFruit  bool
	fruitName string
	color     string
	size      PlantSize
}

func startsWithVowel(a string) bool {
	list := []string{"a", "e", "i", "o", "u"}
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func GetRandomWood() string {
	rand.Seed(time.Now().UnixNano())
	draw := rand.Intn(3)
	woodTypes = []string{"soft", "normal", "hard"}
	return woodTypes[draw]
}

func (p *Plant) Shake() {
	fmt.Printf("\n======= Shaking a %v tree...\n", p.name)
	if p.plantType == Tree && p.size == Mature && p.hasFruit {
		if startsWithVowel(p.fruitName) {
			fmt.Printf("Ah, an %v just fell off.\n", p.fruitName)
		} else {
			fmt.Printf("Ah, a %v just fell off.\n", p.fruitName)
		}
	} else {
		fmt.Println("You can't shake this plant.")
	}
}

func (p *Plant) Chop() {
	fmt.Printf("\n======= Chopping a %v tree...\n", p.name)
	if p.plantType == Tree && p.size == Mature {
		fmt.Printf("Woohoo, %v wood!\n", GetRandomWood())
	} else {
		fmt.Println("Nope, you can't chop this plant.")
	}
}

func (p *Plant) Dig() {
	fmt.Printf("\n======= Digging a %v tree...\n", p.name)
	canDig := false
	switch p.plantType {
	case Tree:
		if p.size != Mature {
			canDig = true
		}
	case Flower:
		canDig = true
	case Shrub:
		canDig = true
	}
	if canDig {
		fmt.Printf("You just digged out the %v %v.\n", p.name, p.plantType)
	} else {
		fmt.Println("You can't dig out this plant.")
	}
}

func main() {
	// example 1
	plantA := Plant{
		plantType: Shrub,
		name:      "hydrangea",
		hasFruit:  false,
		color:     "blue",
		size:      Mature,
	}

	fmt.Println("\n******** Hey there's a hydrangea! ********")
	fmt.Println(plantA)

	var i0 Diggable
	i0 = &plantA
	i0.Dig()

	// example 2
	plantB := Plant{
		plantType: Tree,
		name:      "coconut",
		fruitName: "coconut",
		hasFruit:  true,
		size:      Mature,
	}

	fmt.Println("\n******** Hey there's a coconut tree! ********")
	fmt.Println(plantB)
	var i1 Shakeable
	i1 = &plantB
	i1.Shake()

	var i2 Choppable
	i2 = &plantB
	i2.Chop()

	var i3 Diggable
	i3 = &plantB
	i3.Dig()

	// example 3
	plantC := Plant{
		plantType: Tree,
		name:      "apple",
		fruitName: "apple",
		hasFruit:  true,
		size:      Sapling,
	}

	fmt.Println("\n******** Hey there's a apple tree! ********")
	fmt.Println(plantC)
	var i4 Shakeable
	i4 = &plantC
	i4.Shake()

	var i5 Choppable
	i5 = &plantC
	i5.Chop()

	var i6 Diggable
	i6 = &plantC
	i6.Dig()
}
