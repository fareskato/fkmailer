package fkmailer

// TemplateData data to be send to template
type TemplateData struct {
	Title          string
	Body           string
	AdditionalInfo string
}

type fKMail struct {
	Domain      string
	Host        string
	Port        int
	Username    string
	Password    string
	Encryption  string
	FromAddress string
	FromName    string
}

type FKMessage struct {
	From        string
	FromName    string
	To          string
	Subject     string
	Attachments []string
	Data        TemplateData
}
