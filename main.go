package main

import(
  "net/http"
  "os"
  "fmt"
  "gopkg.in/gomail.v2"
)

import _ "github.com/joho/godotenv/autoload"

func main() {
  http.HandleFunc("/", hello)
  http.HandleFunc("/contact", contact)

  port := os.Getenv("PORT")

  if port == "" {
    panic("PORT must be set")
  }

  http.ListenAndServe(":" + port, nil)
}

func hello(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("hello!"))
}

func contact(w http.ResponseWriter, r *http.Request){
    r.ParseForm()
    toEmail := os.Getenv("EMAIL")

    // c := fmt.Sprintf("Hi %s\n We've recieved a message from %s:\n %s \nThanks!", r.FormValue("name"), r.FormValue("email"), r.FormValue("message"))

    // // set details to recipient to say we got their submission
    // confirmation := new(mailDetails)
    // confirmation.To = r.FormValue("email")
    // confirmation.From = email
    // confirmation.Subject = "Thanks for your message!"
    // confirmation.Body = c

    // confirmation.sendMail()

    // send me the submission
    fromEmail := r.FormValue("email")
    name := r.FormValue("name")
    subject := fmt.Sprintf("New contact form submission from %s (%s)", name, fromEmail)
    body := fmt.Sprintf("You've recieved a message from %s(%s):\n %s", name, fromEmail, r.FormValue("message"))

    submission := new(mailDetails)
    submission.To = toEmail
    submission.From = fromEmail
    submission.Subject = subject
    submission.Body = body

    submission.sendMail()
}

type mailDetails struct {
  To string
  From string
  Subject string
  Body string
}

func (m mailDetails) sendMail() int {
  g := gomail.NewMessage()
  g.SetHeader("From", m.From)
  g.SetHeader("To", m.To)
  g.SetHeader("Subject", m.Subject)
  g.SetBody("text/html", m.Body)

  d := gomail.NewDialer(os.Getenv("MAILGUN_SMTP_SERVER"), 587, os.Getenv("MAILGUN_SMTP_LOGIN"), os.Getenv("MAILGUN_SMTP_PASSWORD"))

  if err := d.DialAndSend(g); err != nil {
    panic(err)
  }

  return 0
}
