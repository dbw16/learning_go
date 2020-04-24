//[ I’ve been busy at work this week with virus stuff. I’m working at home, which is less productive
//han at the office. And we are making accommodations regarding the virus for our students,
//which requires a lot of urgent work. So, a quick little exercise today. ]
//
//Given a length-n list like (a b c d e), the rotations of the list are the n
//lists (a b c d e), (b c d e a), (c d e a b), (d e a b c), and (e a b c d), in any order.
//
//Your task is to write a program that takes a list and returns its rotations. When you are finished,
//you are welcome to read or run a suggested solution, or to post your own solution or discuss the
//exercise in the comments below.

package main

import (
	"fmt"
)

//this will return a left rotation of a list
func rotate_left(my_list []int) []int {
	new_array := my_list[1:]
	new_array = append(new_array, my_list[0])
	return new_array

}

//////this will return a list of all possible rotations
func Rotations(my_list []int) [][]int {
	rotations := make([][]int, len(my_list))
	rotations[0] = my_list[:]

	for i := 1; i < len(my_list); i++ {
		rotations[i] = rotate_left(rotations[i-1])
	}
	return rotations
}

func main() {
	//first we will hard code in the array we want to rotate
	my_list := [...]int{1, 2, 3}

	//print all the roations of this list
	fmt.Print(Rotations(my_list[:]))
}
