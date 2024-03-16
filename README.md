# Mailer package

- **1:** Set mailer server configs in .env file like so:

```.env
    
      MAIL_DOMAIN="mailtrap"
      MAIL_HOST="mailtrap.io or ip address"
      MAIL_PORT=25
      MAIL_USERNAME="user name"
      MAIL_PASSWORD="secret password"
      MAIL_ENCRYPTION:"encryption"
      MAIL_FROM_NAME="Me"
      MAIL_FROM_ADDRESS="myaddress@mail.com"
      
```

- **2:** Create/Init the mailer

```go
  mailer := CreateMail()

```

- **3:** Create the message body of type TemplateData

```go

  data := fkmail.TemplateDate{
  Title:          "Welcome",
  Body:           "this is the body",
  AdditionalInfo: "Thank U",
  }

```

- **4:** Create the message(to be sent)

```go
 message := fkmailer.Message{
  From:    "me@gmail.com",
  To:      "yo@gmail.com",
  Subject: "My first go mailer package",
  Data:    data,
 }

```

- In case U want to send cc create them like so:

```go
  ccs := []string{"cc1@mail.com", "cc1@mail.com", ...etc}

```

- **5:** Send email:

```go
 // in case no cc
 err := mailer.SendSMTPMessage(message, nil)
//  err := mailer.SendSMTPMessage(message, ccs)

 if err != nil {
  log.Fatal(err.Error())
 }

```
