package main

import (
	"math"
	"sort"
)

type distanceSortUp []distance

func (a distanceSortUp) Len() int           { return len(a) }
func (a distanceSortUp) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a distanceSortUp) Less(i, j int) bool { return a[i].dist < a[j].dist }

type distanceSortDown []distance

func (a distanceSortDown) Len() int           { return len(a) }
func (a distanceSortDown) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a distanceSortDown) Less(i, j int) bool { return a[i].dist > a[j].dist }

type distance struct {
	dist  float64
	index int
}

func getEntityByAirDistance(e []Entity, x, y int, nearest bool) Entity {
	var s []distance
	if len(e) <= 0 {
		return nil
	}

	for i, entity := range e {
		d := math.Sqrt(float64(x-entity.getX())*float64(x-entity.getX()) + float64(y-entity.getY())*float64(y-entity.getY()))
		s = append(s, distance{d, i})
	}
	if nearest == true {
		sort.Sort(distanceSortUp(s))
	} else {
		sort.Sort(distanceSortDown(s))
	}
	return e[s[0].index]
}
