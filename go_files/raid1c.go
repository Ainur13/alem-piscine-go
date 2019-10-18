package main

//import "fmt"
import "github.com/01-edu/z01"

func Raid1a(x,y int) {
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 && j == 0 {
				z01.PrintRune('A')
			} else if i == 0 && j == x - 1 {
				z01.PrintRune('A')
			} else if i == y - 1 && j == 0{
				z01.PrintRune('C')
			} else if i == y - 1 && j == x - 1 {
				z01.PrintRune('C')
			} else if j == 0 || j == x - 1 {
				z01.PrintRune('B')
			} else if (j > 0 && j < x - 1) && (i == 0 || i == y - 1){
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}

			if j == x - 1 {
				z01.PrintRune(10)
			}
		}
	}
}

func main() {
	//Raid1a(5,3)
	//Raid1a(5,1)
	//Raid1a(1,1)
	Raid1a(1,5)
}
