package fkmailer

import (
	"bytes"
	"errors"
	"html/template"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/gookit/color"
	"github.com/vanng822/go-premailer/premailer"
	mail "github.com/xhit/go-simple-mail/v2"
)

func (m *fKMail) SendSMTPMessage(msg FKMessage, ccs []string) error {
	if msg.From == "" {
		msg.From = m.FromAddress
	}

	if msg.FromName == "" {
		msg.FromName = m.FromName
	}
	formattedMessage, err := m.buildHTMLMessage(msg)
	if err != nil {
		return err
	}

	server := mail.NewSMTPClient()
	server.Host = m.Host
	server.Port = m.Port
	server.Username = m.Username
	server.Password = m.Password
	server.Encryption = m.getEncryption(m.Encryption)
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	// connect to smtp server
	smtpClient, err := server.Connect()
	if err != nil {
		return err
	}

	email := mail.NewMSG()
	email.SetFrom(msg.From).
		AddTo(msg.To).
		SetSubject(msg.Subject)

		// email body
	email.SetBody(mail.TextHTML, formattedMessage)

	// attachments
	if len(msg.Attachments) > 0 {
		for _, x := range msg.Attachments {
			email.AddAttachment(x)
		}
	}
	// cc
	for _, cc := range ccs {
		email.AddCc(cc)
	}
	// send the eamil
	err = email.Send(smtpClient)
	if err != nil {
		color.Errorln("SMTP connect error: ", err.Error())
		return err
	}
	// all good
	return nil
}

func (m *fKMail) buildHTMLMessage(msg FKMessage) (string, error) {
	var dir string
	customTplStr := os.Getenv("CUSTOM_TPL")
	isCustomTpl, err := strconv.ParseBool(customTplStr)
	if err != nil {
		return "", err
	}
	if isCustomTpl {
		wd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		dir = wd
	} else {
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			return "", errors.New("failed to determine mailer template file path")
		}
		dir = filepath.Dir(filename)
	}

	templateToRender := filepath.Join(dir, "templates", "mail.gohtml")
	t, err := template.New("email-html").ParseFiles(templateToRender)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err = t.ExecuteTemplate(&tpl, "body", msg.Data); err != nil {
		return "", err
	}

	formattedMessage := tpl.String()
	formattedMessage, err = m.inlineCSS(formattedMessage)
	if err != nil {
		return "", err
	}
	return formattedMessage, nil
}

func (m *fKMail) inlineCSS(s string) (string, error) {
	options := premailer.Options{
		RemoveClasses:     false,
		CssToAttributes:   false,
		KeepBangImportant: true,
	}

	prem, err := premailer.NewPremailerFromString(s, &options)
	if err != nil {
		return "", err
	}

	html, err := prem.Transform()
	if err != nil {
		return "", err
	}

	return html, nil
}

func (m *fKMail) getEncryption(s string) mail.Encryption {
	switch s {
	case "tls":
		return mail.EncryptionSTARTTLS
	case "ssl":
		return mail.EncryptionSSLTLS
	case "none", "":
		return mail.EncryptionNone
	default:
		return mail.EncryptionSTARTTLS
	}
}
