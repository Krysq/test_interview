package utills

import (
	"../calculator/models"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
)

type ModelError struct {
	Message string `json:"message,omitempty"`
}

func SendServerError(errorMessage string, code int, w http.ResponseWriter) {
	log.Error().Msgf(errorMessage)
	w.WriteHeader(code)
	mes, _ := json.Marshal(models.ResultsStruct{ErrCode: errorMessage})
	w.Write(mes)
}

func SendOKAnswer(data interface{}, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	serializedData, err := json.Marshal(data)
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}

	_, err = w.Write(serializedData)
	if err != nil {
		message := fmt.Sprintf("HttpResponse is socket: %s", err.Error())
		log.Error().Msgf(message)
		return
	}
	log.Info().Msgf("OK")
}
