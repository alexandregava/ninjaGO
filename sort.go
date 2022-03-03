package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type SortByAge []user

func (a SortByAge) Len() int           { return len(a) }
func (a SortByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type SortByLast []user

func (a SortByLast) Len() int           { return len(a) }
func (a SortByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

type SortBySay []user

func (a SortBySay) Len() int           { return len(a) }
func (a SortBySay) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBySay) Less(i, j int) bool { return len(a[i].Sayings) < len(a[j].Sayings) }

func main() {
	u1 := user{
		First:   "James",
		Last:    "Bond",
		Age:     32,
		Sayings: []string{"gg", "hh", "ii", "ff"},
	}

	u2 := user{
		First:   "Miss",
		Last:    "Moneypenny",
		Age:     27,
		Sayings: []string{"d", "e", "f", "tt", "ss"},
	}

	u3 := user{
		First:   "M",
		Last:    "Hmmmm",
		Age:     54,
		Sayings: []string{"a", "b", "c"},
	}

	users := []user{u1, u2, u3}

	sort.Sort(SortByAge(users))
	fmt.Println(users)

	sort.Sort(SortByLast(users))
	fmt.Println(users)

	sort.Sort(SortBySay(users))
	fmt.Println(users)
}
