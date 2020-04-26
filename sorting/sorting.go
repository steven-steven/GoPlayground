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