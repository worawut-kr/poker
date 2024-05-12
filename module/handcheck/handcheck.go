package handcheck

import (
	"fmt"
)

func hand_check() {

	fmt.Println("hello")
}

type card struct {
	value string
	suit  string
}

func pocket_hand_check(value int, suit int) {
	array_suit := [4]string{"spades", "hearts", "clubs", "diamonds"}
	array_value := [13]string{"ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "jack", "queen", "king"}
	array_hand := []card{}

	for _, v := range array_suit {
		for _, vv := range array_value {
			array_hand = append(array_hand, card{value: vv, suit: v})
		}
	}

	for i := 0; i < suit; i++ {
		for ii := 0; ii < value; ii++ {

		}
	}
}
