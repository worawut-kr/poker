package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// สร้างแผนที่สำหรับค่าไพ่
	cardValues := map[int]string{
		1:  "Ace",
		2:  "2",
		3:  "3",
		4:  "4",
		5:  "5",
		6:  "6",
		7:  "7",
		8:  "8",
		9:  "9",
		10: "10",
		11: "Jack",
		12: "Queen",
		13: "King",
	}

	// สร้างแผนที่สำหรับชุดไพ่
	suits := map[int]string{
		1: "Spades",
		2: "Clubs",
		3: "Diamonds",
		4: "Hearts",
	}

	fmt.Println("กรุณาใส่หมายเลขสำหรับค่าไพ่ (1 ถึง 13):")
	cardValueInput, _ := reader.ReadString('\n')
	cardValueInput = cardValueInput[:len(cardValueInput)-1]
	cardValue, err := strconv.Atoi(cardValueInput)
	if err != nil || cardValue < 1 || cardValue > 13 {
		fmt.Println("ค่าไม่ถูกต้อง")
		return
	}

	fmt.Println("กรุณาใส่หมายเลขสำหรับชุดไพ่ (1: Spades, 2: Clubs, 3: Diamonds, 4: Hearts):")
	suitInput, _ := reader.ReadString('\n')
	suitInput = suitInput[:len(suitInput)-1]
	suit, err := strconv.Atoi(suitInput)
	if err != nil || suit < 1 || suit > 4 {
		fmt.Println("ค่าไม่ถูกต้อง")
		return
	}

	fmt.Printf("คุณเลือก: %s of %s\n", cardValues[cardValue], suits[suit])
}
