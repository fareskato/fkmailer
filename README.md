# Mailer package

- **1:** Set mailer server configs(all fileds are required) in .env file like so:

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
  mailer := fkmailer.CreateMail()

```

- **3:** Create the message body of type TemplateData

```go

  data := fkmailer.TemplateData{
  Title:          "Welcome",
  Body:           "this is the body",
  AdditionalInfo: "Thank U",
  }

```

- **4:** Create the message(to be sent)

```go
/*

 these message fileds if not provided will take the fields
 of mailer(so they are optionals when set the message) 
 From        string
 FromName    string
*/
 message := fkmailer.FKMessage{
  // From:    "me@gmail.com",
  To:      "yo@gmail.com",
  Subject: "My first go mailer package",
  Data:    data,
 }

```

- In case U want to send cc create them like so:

```go
  ccs := []string{"cc1@mail.com", "cc1@mail.com", ...etc}

```

- **Notice:** You can use custom template by adding this to .env file:
CUSTOM_TPL= true
then U need to create template folder at the root with mail.gohtml file
like so:

```go
{{define "body"}}
<!doctype html>

<html lang="en">
    <head>
        <meta name="viewport" content="width=device-width" />
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
        <title>Message from System</title>
    </head>

    <body>
    <!-- additional stuff -->
      {{.Title}}
      {{.Body}}
      {{.AdditionalInfo}}
      <!-- additional stuff -->
    </body>
</html>
{{end}}

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
