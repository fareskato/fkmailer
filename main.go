package fkmailer

import (
	"os"
	"strconv"

	"github.com/gookit/color"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		color.Errorln("Error loading .env file")
	}
	os.Exit(1)
}

/*
func main() {
	tData := fkmailer.TemplateData{
		Title:          "Welcome",
		Body:           "this is the body",
		AdditionalInfo: "Thank U",
	}
	message := fkmailer.Message{
		From:    "fares@gmail.com",
		To:      "ashamaz@gmail.com",
		Subject: "My first go mailer package",
		Data:    tData,
	}
	mailer := CreateMail()
	err := mailer.SendSMTPMessage(message, nil)

	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("email sended")
}
*/

func CreateMail() fKMail {
	portStr := os.Getenv("MAIL_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		color.Errorln(err.Error())
	}

	return fKMail{
		Domain:      os.Getenv("MAIL_DOMAIN"),
		Host:        os.Getenv("MAIL_HOST"),
		Port:        port,
		Username:    os.Getenv("MAIL_USERNAME"),
		Password:    os.Getenv("MAIL_PASSWORD"),
		Encryption:  os.Getenv("MAIL_ENCRYPTION"),
		FromName:    os.Getenv("MAIL_FROM_NAME"),
		FromAddress: os.Getenv("MAIL_FROM_ADDRESS"),
	}
}
