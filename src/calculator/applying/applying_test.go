package applying

import (
	"../models"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculation(t *testing.T) {
	type readInputTestCase struct {
		input  models.InputStruct
		output models.ResultsStruct
	}
	var applying = CalculatorUseNew(10)
	ctx := context.Background()
	testApply := []readInputTestCase{
		{
			input: models.InputStruct{
				A:         4,
				B:         5,
				Operation: "+",
			},
			output: models.ResultsStruct{
				Success: true,
				ErrCode: "",
				Value:   9,
			},
		},
		{
			input: models.InputStruct{
				A:         2.3,
				B:         5.0,
				Operation: "*",
			},
			output: models.ResultsStruct{
				Success: true,
				ErrCode: "",
				Value:   11.5,
			},
		},
		{
			input: models.InputStruct{
				A:         5,
				B:         2,
				Operation: "-",
			},
			output: models.ResultsStruct{
				Success: true,
				ErrCode: "",
				Value:   3,
			},
		},
		{
			input: models.InputStruct{
				A:         10,
				B:         0,
				Operation: "/",
			},
			output: models.ResultsStruct{
				Success: false,
				ErrCode: "by 0",
				Value:   0,
			},
		},
	}

	for _, testCase := range testApply {
		output := applying.Calculate(ctx, testCase.input)
		assert.Equal(t, output, testCase.output)
	}
}
