package main

import "fmt"

type point struct {
	x int
	y int
}

var idCounter int

type Entity interface {
	nextStep(b [][]Entity)
	move(newx, newy int)
	updateEnergy(ne int)
	getEnergy() int
	getX() int
	getY() int
	getPath() []point
	setPath(p []point)
	setDistance(d float32)
	getDistance() float32
	getCosts() map[int]float64
}

type EntityType struct {
	energy   int
	name     string
	id       int
	x        int
	y        int
	path     []point
	distance float32
	costs    map[int]float64
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

func (e None) nextStep(b [][]Entity) {
}

func (e GoodPlant) nextStep(b [][]Entity) {
	fmt.Println("nextStee from GoodPlant")
}

func (e BadPlant) nextStep(b [][]Entity) {
	fmt.Println("nextStee from BadPlant")
}

func (e MasterSquirrel) nextStep(b [][]Entity) {
	fmt.Println("nextStee from MasterSquirrel")
}

func (e GoodBeast) nextStep(b [][]Entity) {
	fmt.Println("nextStee from GoodBeast")
}

func (e BadBeast) nextStep(b [][]Entity) {
	fmt.Println("nextStee from BadBeast")
}

func (e Wall) nextStep(b [][]Entity) {
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

func createNone(c xy) *None {
	e := new(None)
	e.energy = -999
	e.name = "None"
	e.x = c.x
	e.y = c.y
	e.distance = 100000000
	e.costs = map[int]float64{
		KindNone:           1.0,
		KindMasterSquirrel: 1.0,
		KindWall:           1.0,
		KindGoodBeast:      1.0,
		KindBadBeast:       1.0,
		KindGoodPlant:      1.0,
		KindBadPlant:       1.0,
	}
	return e
}

func createGoodPlant(c xy) *GoodPlant {
	e := new(GoodPlant)
	e.id = idCounter
	idCounter++
	e.energy = 100
	e.name = "GoodPlant"
	e.x = c.x
	e.y = c.y
	e.distance = 100000000
	e.costs = map[int]float64{
		KindNone:           1.0,
		KindMasterSquirrel: 1.0,
		KindWall:           1.0,
		KindGoodBeast:      1.0,
		KindBadBeast:       1.0,
		KindGoodPlant:      1.0,
		KindBadPlant:       1.0,
	}
	return e
}

func createBadPlant(c xy) *BadPlant {
	e := new(BadPlant)
	e.id = idCounter
	idCounter++
	e.energy = -100
	e.name = "BadPlant"
	e.x = c.x
	e.y = c.y
	e.distance = 100000000
	e.costs = map[int]float64{
		KindNone:           1.0,
		KindMasterSquirrel: 1.0,
		KindWall:           1.0,
		KindGoodBeast:      1.0,
		KindBadBeast:       1.0,
		KindGoodPlant:      1.0,
		KindBadPlant:       1.0,
	}
	return e
}

func createGoodBeast(c xy) *GoodBeast {
	e := new(GoodBeast)
	e.id = idCounter
	idCounter++
	e.energy = 200
	e.name = "GoodBeast"
	e.x = c.x
	e.y = c.y
	e.distance = 100000000
	e.costs = map[int]float64{
		KindNone:           200.0,
		KindMasterSquirrel: 20000000.0,
		KindWall:           100000.0,
		KindGoodBeast:      5000.0,
		KindBadBeast:       5000.0,
		KindGoodPlant:      5000.0,
		KindBadPlant:       5000.0,
	}
	return e
}

func createBadBeast(c xy) *BadBeast {
	e := new(BadBeast)
	e.id = idCounter
	idCounter++
	e.energy = -150
	e.name = "BadBeast"
	e.x = c.x
	e.y = c.y
	e.distance = 100000000
	e.costs = map[int]float64{
		KindNone:           200.0,
		KindMasterSquirrel: 1.0,
		KindWall:           100000.0,
		KindGoodBeast:      5000.0,
		KindBadBeast:       25000.0,
		KindGoodPlant:      5000.0,
		KindBadPlant:       5000.0,
	}
	return e
}

func createWall(c xy) *Wall {
	e := new(Wall)
	e.id = idCounter
	idCounter++
	e.energy = -10
	e.name = "Wall"
	e.x = c.x
	e.y = c.y
	e.distance = 100000000
	e.costs = map[int]float64{
		KindNone:           1.0,
		KindMasterSquirrel: 1.0,
		KindWall:           1.0,
		KindGoodBeast:      1.0,
		KindBadBeast:       1.0,
		KindGoodPlant:      1.0,
		KindBadPlant:       1.0,
	}
	return e
}

func createMasterSquirrel(c xy) *MasterSquirrel {
	e := new(MasterSquirrel)
	e.id = idCounter
	idCounter++
	e.energy = 1000
	e.name = "MasterSquirrel"
	e.x = c.x
	e.y = c.y
	e.distance = 100000000
	e.costs = map[int]float64{
		KindNone:           2000.0,
		KindMasterSquirrel: 20000.0,
		KindWall:           1000000.0,
		KindGoodBeast:      10000.0,
		KindBadBeast:       20000.0,
		KindGoodPlant:      1.0,
		KindBadPlant:       50000.0,
	}
	return e
}

func (e *GoodPlant) setPath(p []point) {
}

func (e *BadPlant) setPath(p []point) {
}

func (e *MasterSquirrel) setPath(p []point) {
	e.path = p
}

func (e *GoodBeast) setPath(p []point) {
	e.path = p
}

func (e *BadBeast) setPath(p []point) {
	e.path = p
}

func (e *Wall) setPath(p []point) {
}

func (e *None) setPath(p []point) {
}

func (e GoodPlant) getPath() []point {
	return e.path
}

func (e BadPlant) getPath() []point {
	return e.path
}

func (e MasterSquirrel) getPath() []point {
	return e.path
}

func (e GoodBeast) getPath() []point {
	return e.path
}

func (e BadBeast) getPath() []point {
	return e.path
}

func (e Wall) getPath() []point {
	return e.path
}

func (e None) getPath() []point {
	return e.path
}

func (e GoodPlant) setDistance(d float32) {
	e.distance = d
}

func (e BadPlant) setDistance(d float32) {
	e.distance = d
}

func (e MasterSquirrel) setDistance(d float32) {
	e.distance = d
}

func (e GoodBeast) setDistance(d float32) {
	e.distance = d
}

func (e BadBeast) setDistance(d float32) {
	e.distance = d
}

func (e Wall) setDistance(d float32) {
	e.distance = d
}

func (e None) setDistance(d float32) {
	e.distance = d
}

func (e GoodPlant) getDistance() float32 {
	return e.distance
}

func (e BadPlant) getDistance() float32 {
	return e.distance
}

func (e MasterSquirrel) getDistance() float32 {
	return e.distance
}

func (e GoodBeast) getDistance() float32 {
	return e.distance
}

func (e BadBeast) getDistance() float32 {
	return e.distance
}

func (e Wall) getDistance() float32 {
	return e.distance
}

func (e None) getDistance() float32 {
	return e.distance
}

func (e GoodPlant) getCosts() map[int]float64 {
	return e.costs
}

func (e BadPlant) getCosts() map[int]float64 {
	return e.costs
}

func (e MasterSquirrel) getCosts() map[int]float64 {
	return e.costs
}

func (e GoodBeast) getCosts() map[int]float64 {
	return e.costs
}

func (e BadBeast) getCosts() map[int]float64 {
	return e.costs
}

func (e Wall) getCosts() map[int]float64 {
	return e.costs
}

func (e None) getCosts() map[int]float64 {
	return e.costs
}
