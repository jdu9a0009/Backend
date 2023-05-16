package main

import "fmt"

func main() {
	//SELL TICKET APP

	countTicket := 100
	isContinue := "y"
	for isContinue == "y" || isContinue == "Y" {
		fmt.Printf("Bilet sotib olishni istayszmi?[y,N]?: ")
		fmt.Scanf("%v", &isContinue)
		if isContinue != "y" || isContinue == "Y" {
			fmt.Printf("Dastur tugatildi!!!\n")
			break
		}

		ticketNum := 0
		fmt.Printf("Nechta bilet sotib olishni istaysiz?: ")
		fmt.Scanf("%d", &ticketNum)
		if ticketNum <= countTicket {
			countTicket -= ticketNum
			fmt.Printf("%d ta bilet sotib olindi\n", ticketNum)
			fmt.Printf("%d dona bilet qoldi sotuvda\n", countTicket)
		} else if countTicket == 0 {
			fmt.Printf("Biletlar sotilib tugadi.Do'kon yopildi!!\n")
			break
		} else {
			fmt.Printf("Buncha bilet mavjud emas!\n")
			fmt.Printf("%d dona bilet qoldi\n", countTicket)
		}
	}
}
