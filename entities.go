package main

import "fmt"

var idCounter int

type Entity interface {
	nextStep()
}

type EntityType struct {
	energy int
	name   string
	id     int
}

type None struct {
	EntityType
}

type GoodBeast struct {
	EntityType
}
type BadBeast struct {
	EntityType
}

type GoodPlant struct {
	EntityType
}
type BadPlant struct {
	EntityType
}

type Wall struct {
	EntityType
}

type MasterSquirrel struct {
	EntityType
}

type MiniSquirrel struct {
	EntityType
}

func (e None) nextStep() {
}

func (e GoodPlant) nextStep() {
	fmt.Println("nextStee from GoodPlant")
}

func (e BadPlant) nextStep() {
	fmt.Println("nextStee from BadPlant")
}

func (e MasterSquirrel) nextStep() {
	fmt.Println("nextStee from MasterSquirrel")
}

func (e MiniSquirrel) nextStep() {
	fmt.Println("nextStee from MiniSquirrel")
}

func (e GoodBeast) nextStep() {
	fmt.Println("nextStee from GoodBeast")
}

func (e BadBeast) nextStep() {
	fmt.Println("nextStee from BadBeast")
}

func (e Wall) nextStep() {
	fmt.Println("nextStee from Wall")
}

func createNone() *None {
	e := new(None)
	e.energy = -999
	e.name = "None"
	return e
}

func createGoodPlant() *GoodPlant {
	e := new(GoodPlant)
	e.id = idCounter
	idCounter++
	e.energy = 100
	e.name = "GoodPlant"
	return e
}

func createBadPlant() *BadPlant {
	e := new(BadPlant)
	e.id = idCounter
	idCounter++
	e.energy = -100
	e.name = "BadPlant"
	return e
}

func createGoodBeast() *GoodBeast {
	e := new(GoodBeast)
	e.id = idCounter
	idCounter++
	e.energy = 200
	e.name = "GoodBeast"
	return e
}

func createBadBeast() *BadBeast {
	e := new(BadBeast)
	e.id = idCounter
	idCounter++
	e.energy = -150
	e.name = "BadBeast"
	return e
}

func createWall() *Wall {
	e := new(Wall)
	e.id = idCounter
	idCounter++
	e.energy = -10
	e.name = "Wall"
	return e
}

func createMasterSquirrel() *MasterSquirrel {
	e := new(MasterSquirrel)
	e.id = idCounter
	idCounter++
	e.energy = 1000
	e.name = "MasterSquirrel"
	return e
}

func createMiniSquirrel() *MiniSquirrel {
	e := new(MiniSquirrel)
	e.id = idCounter
	idCounter++
	e.energy = 0
	e.name = "MiniSquirrel"
	return e
}
