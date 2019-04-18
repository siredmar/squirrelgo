package main

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
