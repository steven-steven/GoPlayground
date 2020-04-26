package sorting

import (
	"fmt"
	"sort"
)

func Sort(){
	unsorted := []int{1,3,2,6,3,4}
	sort.Ints(unsorted)	//inplace sorting
	fmt.Println(unsorted)
}

//---------------- Custom Sorting ----------------

type Programmer struct {
	Age int
}

//define array type that implements Len(), Swap(), Less()
type byAge []Programmer

func (p byAge) Len() int{
	return len(p)
}
func (p byAge) Swap(i, j int) {	//swap given index
	p[i], p[j] = p[j], p[i]
}
func (p byAge) Less(i, j int) bool{	//is item at i less than (appears before) j 
	return p[i].Age < p[j].Age
}

func CustomSort(){
	programmers := []Programmer{
		Programmer{Age:30},
		Programmer{Age:20},
		Programmer{Age:50},
		Programmer{Age:1000},
	}
	sort.Sort(byAge(programmers))	//cast array to our type, and apply .Sort
	fmt.Println(programmers)
}