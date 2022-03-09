package gomodule

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Add(x int, y int) int {
	return x + y
}

func RetTwo() (int, int) {
	return 1, 2
}

func Multi(x int, y int) (result int) {
	result = x * y
	return
}

func Arithmetics() {
	var x uint64 = 1
	x = x << 63
	fmt.Println(x)
}

func SplitString(x string, y int) string {
	var args string = x
	var ret_value = strings.Split(args, " ")
	return ret_value[y]
}

func Check(x int, y int) {
	if x == y {
		fmt.Println("True")
	} else if x < y {
		fmt.Println("Smaller")
	} else {
		fmt.Println("Bigger")
	}
}

func Compare1(name string) {
	switch name {
	case "LOL":
		fmt.Println("LOOOOOOOOOOOOOL")
	case "Sigma":
		fmt.Println("Was das für ein Name")
	default:
		fmt.Println("Nichts")
	}
}

func Compare2() {
	switch {
	case 1 > 0:
		fmt.Println("Hätte ich gar nicht gedacht")
	}
}

func ForLoop1() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func ForLoop2(cake int) {
	for cake < 27 {
		cake++
		fmt.Println(cake)
	}
}

func DeferStatement() {
	defer fmt.Println("Defer 1")
	fmt.Println("LOL")
	defer fmt.Println("Defer 2")
}

func Zeiger() {
	/*i := 2
	fmt.Println(i)
	//pointer := &i
	changeICopy(i)
	fmt.Println(i)
	editNumberPointer(&i)
	fmt.Println(i)*/

	value1 := 27
	value2 := 44
	p1 := &value1
	p2 := &value2

	fmt.Println(*p1, *p2)
	swapPointer(&p1, &p2)
	fmt.Println(*p1, *p2)

}

func swapPointer(p1 **int, p2 **int) {
	tmp := *p1
	*p1 = *p2
	*p2 = tmp
}

func GetConsoleInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	return input
}

func Arrays() {
	var buffer [256]int
	for i := 0; i < len(buffer); i++ {
		buffer[i] = i
	}
	for _, element := range buffer {
		fmt.Println(element)
	}
}

func Slices() {
	var buffer [256]int
	for i := 0; i < len(buffer); i++ {
		buffer[i] = i
	}
	slice := buffer[27:72]

	for _, element := range slice {
		fmt.Println(element)
	}
	fmt.Println(cap(slice))
}

func SlicesTwo() {
	var array [256]int
	for i := 0; i < len(array); i++ {
		array[i] = i
	}
	slice := array[0:]
	slice = append(slice, 42)
	for _, element := range slice {
		fmt.Println(element)
	}
}

func CreateIntSlice() []int {
	var retArray [42]int
	for i := 0; i < len(retArray); i++ {
		retArray[i] = 0
	}
	return_value := retArray[0:]
	return return_value
}

func PrintSlice(slice []int) {
	for _, element := range slice {
		if element != 0 {
			fmt.Println(element)
		}
	}
}

/*func editNumberPointer(j *int) {
	*j = 27
}

func changeICopy(i int) {
	i = 27
}*/
