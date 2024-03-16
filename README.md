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

## Example

```go

      Domain:   "mailtrap",       // os.Getenv("MAIL_DOMAIN"),
      Host:     "3.209.246.195",  // os.Getenv("MAIL_HOST"),
      Port:     25,               // port,
      Username: "c2c7de5dc360ec", // os.Getenv("MAIL_USERNAME"),
      Password: "cc99f40ff334eb", // os.Getenv("MAIL_PASSWORD"),
      // Encryption:  // os.Getenv("MAIL_ENCRYPTION"),
      FromName:    "Fares Kato",      // os.Getenv("FROM_NAME"),
      FromAddress: "fares@gmail.com", 
    // os.Getenv("FROM_ADDRESS"),
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
