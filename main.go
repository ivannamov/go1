package main

import "fmt"

var (
	applePrice float64 = 5.99
	pearPrice  float64 = 7
	myMoney    float64 = 23
)

func main() {
	fmt.Println("дано", applePrice, pearPrice)
	sum := applePrice*9 + pearPrice*8
	fmt.Println("Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?", sum)

	fmt.Println("дано", pearPrice, myMoney)
	pearQuantity := int(myMoney / pearPrice)
	fmt.Println("Скільки груш ми можемо купити?", pearQuantity)

	fmt.Println("дано", applePrice, myMoney)
	appleQuantity := int(myMoney / applePrice)
	fmt.Println("Скільки яблук ми можемо купити?", appleQuantity)

	fmt.Println("дано", applePrice, pearPrice, myMoney)
	canWeBye := applePrice*2+pearPrice*2 < myMoney
	fmt.Println("Чи ми можемо купити 2 груші та 2 яблука?", canWeBye)
}
