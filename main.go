package main

import(
  "net/http"
  // "log"
  // "os"
  "fmt"
  "gopkg.in/gomail.v2"
)
func main() {
  http.HandleFunc("/", hello)
  http.HandleFunc("/contact", contact)
  // do we serve off an enviroment variable here?
  http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("hello!"))
}
func contact(w http.ResponseWriter, r *http.Request){
    // f, err := os.OpenFile("log.txt", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
    // if err != nil {
    //   log.Fatal(err)
    // }

    // defer f.Close()
    // log.SetOutput(f)

    r.ParseForm()
    c := fmt.Sprintf("Hi %s\n We've recieved a message from %s:\n %s \nThanks!", r.FormValue("name"), r.FormValue("email"), r.FormValue("message"))

    // set details to recipient to say we got their submission
    confirmation := new(mailDetails)
    confirmation.To = r.FormValue("email")
    confirmation.From = "info@dankleiman.com"
    confirmation.Subject = "Thanks for your message!"
    confirmation.Body = c

    // log.Println(confirmation.Body)
    // confirmation.sendMail()

    // // send me the submission
    s := fmt.Sprintf("You've recieved a message from %s(%s):\n %s", r.FormValue("name"), r.FormValue("email"), r.FormValue("message"))

    submission := new(mailDetails)
    submission.To = "info@dankleiman.com"
    submission.From = r.FormValue("email")
    submission.Subject = "New contact form submission"
    submission.Body = s

    // log.Println(submission.Body)
    // submission.sendMail()
  }

  type mailDetails struct {
    To string
    From string
    Subject string
    Body string
  }

  func (m mailDetails) sendMail() int {
    // Set up authentication information.
    g := gomail.NewMessage()
    g.SetHeader("From", m.From)
    g.SetHeader("To", m.To)
    g.SetHeader("Subject", m.Subject)
    g.SetBody("text/html", m.Body)

    d := gomail.NewDialer("smtp.example.com", 587, "user", "123456")

    if err := d.DialAndSend(g); err != nil {
      panic(err)
    }

    return 0
  }
