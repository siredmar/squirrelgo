package main

type distanceSortUp []distance

func (a distanceSortUp) Len() int {
	return len(a)
}
func (a distanceSortUp) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a distanceSortUp) Less(i, j int) bool {
	if a[i].count == a[j].count {
		return a[i].dist < a[j].dist
	} else {
		return a[i].count < a[j].count
	}
}

type distanceSortDown []distance

func (a distanceSortDown) Len() int {
	return len(a)
}
func (a distanceSortDown) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a distanceSortDown) Less(i, j int) bool {
	if a[i].count == a[j].count {
		return a[i].dist > a[j].dist
	} else {
		return a[i].count > a[j].count
	}
}

type distance struct {
	dist  float64
	count int
	index int
}
