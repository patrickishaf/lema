package common

import "testing"

func TestCreateErrorResponse(t *testing.T) {
	message := "unable to find source"
	res := CreateErrorResponse(message)
	status, ok := res["status"]
	if !ok {
		t.Errorf("failed to convert message to an error response. status field missing")
	}
	if status != "error" {
		t.Errorf("invalid value for status in CreateErrorResponse")
	}

	msg, ok2 := res["message"]
	if !ok2 {
		t.Errorf("failed to convert message to an error response. message field missing")
	}
	if msg != message {
		t.Errorf("incorrect conversion. invalid value in `message` key")
	}
}

func TestGenerateID(t *testing.T) {
	id, err := GenerateID()
	if err != nil {
		t.Errorf("failed to generate id. error: %v", err)
	}
	if len(id) != 32 {
		t.Errorf("generated id does not have a length of 32")
	}
}
