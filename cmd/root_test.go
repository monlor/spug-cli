package cmd

import "testing"

func TestSendMessage(t *testing.T) {
	readUser()
	err := SendMessage("hello")
	if err != nil {
		t.Error(err)
	}
}
