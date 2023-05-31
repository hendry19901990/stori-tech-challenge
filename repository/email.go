package repository

import (
	"bytes"
	"html/template"
	"net/smtp"
)

//Request struct
type Request struct {
	from    string
	to      []string
	subject string
	body    string
	auth *Auth
}

type Auth struct {
	Email string
	Password string
}

func NewRequest(to []string, subject, body string, auth *Auth) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
		auth: auth,
	}
}

func(r *Request) GetBody()string{
	return r.body
}


func (r *Request) SendEmail() error {
    auth := smtp.PlainAuth("", r.auth.Email, r.auth.Password, "smtp.gmail.com") 

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := "smtp.gmail.com:587"

	if err := smtp.SendMail(addr, auth, r.auth.Email, r.to, msg); err != nil {
	 return err
	}
	return  nil
}

func (r *Request) ParseTemplate(templateFile string, data interface{}) error {
	t, err := template.ParseFiles(templateFile) 
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}