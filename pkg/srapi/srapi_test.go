package srapi

import (
	"fmt"
	"testing"
)

func TestGetGameByID(t *testing.T) {
	game, err := GetGameByID("j1n8z91p")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("game = %+v\n", game)
}

func TestGetGameByName(t *testing.T) {
	games, err := GetGameByName("Space Quest 6")
	if err != nil {
		t.Fatal(err)
	}
	for _, game := range games {
		fmt.Printf("game = %+v\n", game)
	}
}
