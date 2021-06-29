package calculator

import (
	"../calculator/models"
	"context"
)

type Applying interface {
	Calculate(c context.Context, inputStruct models.InputStruct) models.ResultsStruct
}
