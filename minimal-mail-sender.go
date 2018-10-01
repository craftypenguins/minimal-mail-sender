package main

import (
  "net/smtp"
  "log"
  "os"
  "crypto/tls"
)

func getenv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}

func main() {
   host := os.Getenv("SMTP_HOST")
   rcpt := os.Getenv("SMTP_RCPT")
   sender := os.Getenv("SMTP_SENDER")

   ignoretls := false
   if len(os.Getenv("SMTP_IGNORE_TLS")) != 0 {
      ignoretls = true
   }

   user := os.Getenv("SMTP_USER")
   pwd := os.Getenv("SMTP_PASSWORD")

   data := os.Getenv("SMTP_BODY")

   subject := getenv("SMTP_SUBJECT", "Hey")

   if len(host) == 0 {
      log.Fatal("SMTP_HOST not set, exiting")
   } else if len(rcpt) == 0 {
      log.Fatal("SMTP_RCPT not set, exiting")
   } else if len(sender) == 0 {
      log.Fatal("SMTP_SENDER not set, exiting")
   } else if ( len(user) != 0 && len(pwd) == 0 ) {
      log.Fatal("SMTP_USER set but not SMTP_PASSWORD, exiting")
   }

   host_url := host + ":" + getenv("SMTP_PORT", "25")

   message := "Subject: " + subject + "\r\n" +
              "To: " + rcpt + "\r\n" + data;

   tlsconfig := &tls.Config{
      InsecureSkipVerify: ignoretls,
      ServerName: host,
   }

   c, err := smtp.Dial(host_url)
   if err != nil { log.Fatal(err) }

   tls_enabled, _ := c.Extension("STARTTLS")

   if tls_enabled {
      err = c.StartTLS(tlsconfig)
      if err != nil { log.Fatal(err) }
   }

   if len(user) != 0 {
      auth := smtp.PlainAuth(
         "",
         user,
         pwd,
         host,
      )
      err = c.Auth(auth)
      if err != nil { log.Fatal(err) }
   }

   err = c.Mail(sender)
   if err != nil { log.Fatal(err) }

   err = c.Rcpt(rcpt)
   if err != nil { log.Fatal(err) }

   w, err := c.Data()
   if err != nil { log.Fatal(err) }

   _, err = w.Write([]byte(message))
   if err != nil { log.Fatal(err) }

   err = w.Close()
   if err != nil { log.Fatal(err) }

   c.Quit()

   log.Println("Successfully sent mail to " + host_url)
}
