package main

import "testing"

func TestApp_getPriceText(t *testing.T) {
	open, _, _ := testApp.getPriceText()
	if open.Text != "Open: $1809.6200 USD" {
		t.Error("Wrong price returned", open.Text)
	}
}
