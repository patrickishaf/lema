package common

import (
	"testing"
)

func TestValidateStruct(t *testing.T) {
	type Rand struct {
		Name string `validate:"required"`
		Age  uint32 `validate:"required"`
	}

	input := Rand{
		Name: "Rand",
	}
	errs := ValidateStruct(input)
	if errs == nil {
		t.Errorf("failed to flag errors in invalid struct")
	}

	input2 := Rand{
		Age: 12,
	}
	errs2 := ValidateStruct(input2)
	if errs2 == nil {
		t.Errorf("failed to flag errors in invalid struct")
	}

	input3 := Rand{}
	errs3 := ValidateStruct(input3)
	if errs3 == nil {
		t.Errorf("failed to flag errors in invalid struct")
	}

	input4 := Rand{
		Name: "TikTok",
		Age:  12,
	}
	errs4 := ValidateStruct(input4)
	if errs4 != nil {
		t.Errorf("failed to validate valid struct")
	}
}

func TestValidateVariable(t *testing.T) {
	email := "joeybloggs.gmail.com"

	err := ValidateVariable(email, SchemaEmail)
	if err == nil {
		t.Errorf("failed to vaidate valid field")
	}
}
