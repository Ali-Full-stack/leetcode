// Fruits into baskets
package main

import (
	"fmt"
	"math"
)


func main(){
	trees :=[]int{0,1,6,6,4,4,6}
	max := FruitsIntoBaskets(trees)
	fmt.Println(max)
}

func FruitsIntoBaskets(trees []int)(max int){
	TreeTypesMap :=make(map[int]bool)
	var start int

	for end, treetype :=range trees {
		if len(TreeTypesMap) < 2 && !TreeTypesMap[treetype] {
			TreeTypesMap[treetype] = true
			max = int(math.Max(float64(max), float64(end - start + 1)))
		}else if TreeTypesMap[treetype]{
			max = int(math.Max(float64(max), float64(end - start + 1)))
		}else {
			TreeTypesMap =make(map[int]bool)
			TreeTypesMap[trees[end - 1]] = true
			TreeTypesMap[treetype] = true
			start = end - 1
		}

		for start > 0 && trees[start] == trees[start - 1]{
			start--
		}
		max = int(math.Max(float64(max), float64(end - start + 1)))

	}
	return max

}