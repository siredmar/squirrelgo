package main

import "fmt"

var idCounter int

type Entity interface {
	nextStep()
	move(newx, newy int)
	updateEnergy(ne int)
	getEnergy() int
	getX() int
	getY() int
}

type EntityType struct {
	energy int
	name   string
	id     int
	x      int
	y      int
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

func (e None) getY() int {
	return 0
}

func (e None) getX() int {
	return e.x
}

func (e GoodBeast) getY() int {
	return e.y
}

func (e GoodBeast) getX() int {
	return e.x
}

func (e BadBeast) getY() int {
	return e.y
}

func (e BadBeast) getX() int {
	return e.x
}

func (e GoodPlant) getY() int {
	return e.y
}

func (e GoodPlant) getX() int {
	return e.x
}

func (e BadPlant) getY() int {
	return e.y
}

func (e BadPlant) getX() int {
	return e.x
}

func (e Wall) getY() int {
	return e.y
}

func (e Wall) getX() int {
	return e.x
}

func (e MasterSquirrel) getY() int {
	return e.y
}

func (e MasterSquirrel) getX() int {
	return e.x
}

func (e None) getEnergy() int {
	return e.energy
}

func (e *None) updateEnergy(ne int) {
	e.energy += ne
}

func (e GoodPlant) getEnergy() int {
	return e.energy
}

func (e *GoodPlant) updateEnergy(ne int) {
	e.energy += ne
}

func (e BadPlant) getEnergy() int {
	return e.energy
}

func (e *BadPlant) updateEnergy(ne int) {
	e.energy += ne
}

func (e MasterSquirrel) getEnergy() int {
	return e.energy
}

func (e *MasterSquirrel) updateEnergy(ne int) {
	e.energy += ne
}

func (e GoodBeast) getEnergy() int {
	return e.energy
}

func (e *GoodBeast) updateEnergy(ne int) {
	e.energy += ne
}

func (e BadBeast) getEnergy() int {
	return e.energy
}

func (e *BadBeast) updateEnergy(ne int) {
	e.energy += ne
}

func (e Wall) getEnergy() int {
	return e.energy
}

func (e *Wall) updateEnergy(ne int) {
	e.energy += ne
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

func (e GoodBeast) nextStep() {
	fmt.Println("nextStee from GoodBeast")
}

func (e BadBeast) nextStep() {
	fmt.Println("nextStee from BadBeast")
}

func (e Wall) nextStep() {
	fmt.Println("nextStee from Wall")
}

func (e *GoodPlant) move(newx, newy int) {
	e.x = newx
	e.y = newy
}

func (e *BadPlant) move(newx, newy int) {
	e.x = newx
	e.y = newy
}

func (e *MasterSquirrel) move(newx, newy int) {
	e.x = newx
	e.y = newy
}

func (e *GoodBeast) move(newx, newy int) {
	e.x = newx
	e.y = newy
}

func (e *BadBeast) move(newx, newy int) {
	e.x = newx
	e.y = newy
}

func (e *Wall) move(newx, newy int) {
}

func (e *None) move(newx, newy int) {
}

func createNone(x, y int) *None {
	e := new(None)
	e.energy = -999
	e.name = "None"
	e.x = x
	e.y = y
	return e

}

func createGoodPlant(x, y int) *GoodPlant {
	e := new(GoodPlant)
	e.id = idCounter
	idCounter++
	e.energy = 100
	e.name = "GoodPlant"
	e.x = x
	e.y = y
	return e
}

func createBadPlant(x, y int) *BadPlant {
	e := new(BadPlant)
	e.id = idCounter
	idCounter++
	e.energy = -100
	e.name = "BadPlant"
	e.x = x
	e.y = y
	return e
}

func createGoodBeast(x, y int) *GoodBeast {
	e := new(GoodBeast)
	e.id = idCounter
	idCounter++
	e.energy = 200
	e.name = "GoodBeast"
	e.x = x
	e.y = y
	return e
}

func createBadBeast(x, y int) *BadBeast {
	e := new(BadBeast)
	e.id = idCounter
	idCounter++
	e.energy = -150
	e.name = "BadBeast"
	e.x = x
	e.y = y
	return e
}

func createWall(x, y int) *Wall {
	e := new(Wall)
	e.id = idCounter
	idCounter++
	e.energy = -10
	e.name = "Wall"
	e.x = x
	e.y = y
	return e
}

func createMasterSquirrel(x, y int) *MasterSquirrel {
	e := new(MasterSquirrel)
	e.id = idCounter
	idCounter++
	e.energy = 1000
	e.name = "MasterSquirrel"
	e.x = x
	e.y = y
	return e
}
