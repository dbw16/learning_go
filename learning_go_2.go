//Today’s exercise is somebody’s homework:
//
//Write a program that displays the digits from 1 to n then back down to 1; for instance,
//if n = 5, the program should display 123454321. You are permitted to use only a single for loop.
//
//The questioner did not specify what should happen when n reaches 10, so we will specify 0 < n < 10.
//
//Your task is to write the requested program; if you like, think of other ways to write that program.
//When you are finished, you are welcome to read or run a suggested solution,
//or to post your own solution or discuss the exercise in the comments below.
package main

import "fmt"

func main() {
	//first we will hard code in the number we want to use
	num := 9
	//we will create a array of length (num * 2) - 1
	var slice = make([]int, num*2-1)
	for i := 0; i < num; i++ {
		slice[i] = i + 1
		slice[((num*2)-1)-i-1] = i + 1
	}
	fmt.Print(slice)
}
