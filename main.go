package main

import "fmt"

var (
	applePrice float64 = 5.99
	pearPrice  float64 = 7
	myMoney    float64 = 23
)

func main() {
	sum := applePrice*9 + pearPrice*8
	fmt.Println("Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?", sum)

	pearQuantity := int(myMoney / pearPrice)
	fmt.Println("Скільки груш ми можемо купити?", pearQuantity)

	appleQuantity := int(myMoney / applePrice)
	fmt.Println("Скільки яблук ми можемо купити?", appleQuantity)

	canWeBye := applePrice*2+pearPrice*2 < myMoney
	fmt.Println("Чи ми можемо купити 2 груші та 2 яблука?", canWeBye)
}
