package main
import "fmt"

func main() {
	applePrice:=5.99
	pearPrice:=7.
	bank:=23.


	fmt.Printf("1)Скільки грошей треба витратити на 9 яблук та 8 груш ? - %.2f грн\n", (9*applePrice) + (8*pearPrice))
	fmt.Printf("2)Скільки груш ми можемо придбати ? - %v шт\n", int(bank/pearPrice))
	fmt.Printf("3)Скільки яблук ми можемо придбати ? - %v шт\n", int(bank/applePrice))
	fmt.Println("4)Чи можемо придбати 2 груші та 2 яблука ? -",bank>=(2*applePrice) + (2*pearPrice))
}
