package commons

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/jeffersonto/feira-api/internal/entity/exceptions"
)

const (
	ErrConversion = "could not convert value %v to integer"
)

func ConvertToInt(input interface{}) (int64, error) {
	switch dataToConvert := input.(type) {
	case int:
		return int64(dataToConvert), nil
	case float32:
		return int64(dataToConvert), nil
	case float64:
		return int64(dataToConvert), nil
	case string:
		if dataToConvert == "" {
			dataToConvert = "0"
		}
		result, err := strconv.ParseInt(dataToConvert, 10, 0)
		if err != nil {
			return 0, exceptions.NewBadRequest(fmt.Errorf(ErrConversion, dataToConvert))
		}
		return result, nil
	case json.Number:
		floatResult, err := dataToConvert.Float64()
		if err != nil {
			return 0, exceptions.NewBadRequest(fmt.Errorf(ErrConversion, dataToConvert))
		}
		return int64(floatResult), nil
	default:
		return 0, exceptions.NewBadRequest(fmt.Errorf(ErrConversion, dataToConvert))
	}
}
