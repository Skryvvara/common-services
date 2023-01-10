package controllers

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Skryvvara/common-services/config"
	"gopkg.in/gomail.v2"
)

type JSON map[string]interface{}

type MailRequest struct {
	Subject string `json:"subject"`
	Message string `json:"message"`
	Data    JSON   `json:"data"`
}

func SendMail(w http.ResponseWriter, r *http.Request) {
	var mailRequest MailRequest
	if err := json.NewDecoder(r.Body).Decode(&mailRequest); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println(err)
		return
	}

	if mailRequest.Message == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println(fmt.Errorf("provided empty email body"))
		return
	}

	if ok := mailRequest.Data["honey"]; ok != nil {
		w.WriteHeader(http.StatusAccepted)
		return
	}

	var from string
	if ok := mailRequest.Data["from"]; ok != nil {
		from = ok.(string)
	} else {
		from = "skryvvara.com"
	}

	host := config.App.Mail.Host
	user := config.App.Mail.User
	password := config.App.Mail.Password
	port := config.App.Mail.Port
	secure := config.App.Mail.Secure
	to := config.App.Mail.To

	m := gomail.NewMessage()
	m.SetAddressHeader("From", user, from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", mailRequest.Subject)
	m.SetBody("text/html", mailRequest.Message)

	d := gomail.NewDialer(host, port, user, password)
	d.SSL = secure
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
