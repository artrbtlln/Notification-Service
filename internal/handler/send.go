package handler

import (
	"encoding/json"
	"fmt"
	"learn/internal/model"
	"log"
	"net/http"
)

func (h *Handler) SendMail(w http.ResponseWriter, r *http.Request) {
	var mama []byte
	if _, err := r.Body.Read(mama); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	m := &model.Mail{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&m)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(m)

	h.service.Send(m)

}
