package commons_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jeffersonto/feira-api/util/commons"
	"github.com/jeffersonto/feira-api/util/exceptions"
	"reflect"
	"testing"
)

type typeWrong struct {
}

func TestConvertToInt(t *testing.T) {
	tests := []struct {
		name          string
		input         interface{}
		intExpected   int64
		errorExpected error
	}{
		{
			name:          "Should return int When receive int",
			input:         123,
			intExpected:   123,
			errorExpected: nil,
		},
		{
			name:          "Should return int When receive float32",
			input:         float32(123),
			intExpected:   123,
			errorExpected: nil,
		},
		{
			name:          "Should return int When receive float64",
			input:         float64(123),
			intExpected:   123,
			errorExpected: nil,
		},
		{
			name:          "Should return int When receive string",
			input:         "123",
			intExpected:   123,
			errorExpected: nil,
		},
		{
			name:          "Should return 0 When receive string empty",
			input:         "",
			intExpected:   0,
			errorExpected: nil,
		},
		{
			name:          "Should return int When receive a json.Number",
			input:         json.Number("123"),
			intExpected:   123,
			errorExpected: nil,
		},
		{
			name:          "Should return error int When receive a invalid json.Number",
			input:         json.Number("123l"),
			intExpected:   0,
			errorExpected: exceptions.NewBadRequest(errors.New("could not convert value 123l to integer")),
		},
		{
			name:          "Should return error When receive invalid string",
			input:         "123l",
			intExpected:   0,
			errorExpected: exceptions.NewBadRequest(errors.New("could not convert value 123l to integer")),
		},
		{
			name:          "Should return error When receive invalid string",
			input:         &typeWrong{},
			intExpected:   0,
			errorExpected: exceptions.NewBadRequest(errors.New(fmt.Sprintf("could not convert value %v to integer", &typeWrong{}))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intReturn, err := commons.ConvertToInt(tt.input)

			if intReturn != tt.intExpected {
				t.Errorf("ConvertToInt() intReturn = %v, intExpected %v", intReturn, tt.intExpected)
			}
			if !reflect.DeepEqual(err, tt.errorExpected) {
				t.Errorf("ConvertToInt() apiErrorReturn = %v, apiErrorExpected %v", err, tt.errorExpected)
			}
		})
	}
}
