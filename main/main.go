package main

import (
	"fmt"
)

type card struct {
	value string
	suit  string
}

func main() {
	array_all_card := []card{}
	h1 := input_card()
	h2 := input_card()
	array_all_card = append(array_all_card, h1, h2)
	same_card_check(array_all_card)

	for i := 0; i < 10; i-- {
		TC_1 := input_card()
		TC_2 := input_card()
		TC_3 := input_card()
		array_all_card = append(array_all_card, TC_1, TC_2, TC_3)
		t := same_card_check(array_all_card)
		if t == 1 {
			i = (i * 0) + 20
			fmt.Println(array_all_card)
		} else if t == 0 {
			continue
		}

	}

}

func same_card_check(array_all_card []card) int {
	for i, v := range array_all_card {
		for _, vv := range array_all_card[i+1:] {
			if v == vv {
				return 0
			}
		}
	}
	return 1
}

func input_card() (hand card) {
	var card_, card_type int
	fmt.Print("what card do u have? ")
	fmt.Scanln(&card_)
	fmt.Print("what type of card? ")
	fmt.Scanln(&card_type)
	hand = pocket_hand_check(card_, card_type)
	return hand
}

func pocket_hand_check(value int, suit int) card {
	array_suit := [4]string{"spades", "hearts", "clubs", "diamonds"}
	array_value := [13]string{"ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "jack", "queen", "king"}
	array_hand := []card{}

	for _, v := range array_suit {
		for _, vv := range array_value {
			array_hand = append(array_hand, card{value: vv, suit: v})
		}
	}
	// fmt.Println(array_hand)
	var pocket card
	var a, b int
	if suit == 1 {
		a = 0
	} else if suit == 2 {
		a = 13
	} else if suit == 3 {
		a = 26
	} else if suit == 4 {
		a = 39
	} else {
		panic("error")
	}

	b = a + (value - 1)

	pocket = array_hand[b]

	return pocket
}
