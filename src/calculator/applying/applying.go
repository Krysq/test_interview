package applying

import (
	"../models"
	"context"
	"time"
)

type CalculatorUse struct {
	contextTimeout time.Duration
}

func CalculatorUseNew(timeout time.Duration) *CalculatorUse {
	return &CalculatorUse{contextTimeout: timeout}
}

func (cu *CalculatorUse) Calculate(c context.Context, params models.InputStruct) models.ResultsStruct {
	_, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	results := models.ResultsStruct{
		Success: true,
		ErrCode: "",
		Value:   0,
	}
	switch params.Operation {
	case "+":
		results.Value = params.A + params.B
	case "-":
		results.Value = params.A - params.B
	case "*":
		results.Value = params.A * params.B
	case "/":
		if params.B == 0 {
			return models.ResultsStruct{
				Success: false,
				ErrCode: "by 0",
				Value:   0,
			}
		}
		results.Value = params.A / params.B

	}

	return results
}

