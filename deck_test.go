package main

import (
	"os"
	"testing"
)

func TestNewDeck (t *testing.T){
	d:= newDeck()
	if len(d)!= 16{
		t.Errorf("Expected deck of cards is 16,but we get %v", len(d))
	}
	if d[0] != "AceofSpades"{
		t.Errorf("Expected Ace but we are getting the  %v", d[0])
	}
	if d[len(d)-1] !="FourofClubs"{
		t.Errorf("expected  clubsnbut wee are getting %v", d[len(d)-1])
	}
}

func SaveToDeckNewDeckFromFile (t *testing.T){
	os.Remove("_decktesting")
	deck:=newDeck()
	deck.saveToFile("_decktesting")
	loadDeck:=newDeckFromFile("_decktesting")
	if len(loadDeck) !=16{
		t.Errorf("Expeccted 16 but getting as %v", len(loadDeck))
	}
	os.Remove("_decktesting")


}