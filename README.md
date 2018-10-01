# minimal-mail-sender

Intended for use in a cronjob/job type infrastructure, entrypoint is the go binary to send mail

Works off environment variables, required:

    SMTP_HOST
    SMTP_RCPT
    SMTP_SENDER

Optional:

    SMTP_IGNORE_TLS
    SMTP_USER
    SMTP_PASSWORD
    SMTP_SUBJECT
    SMTP_BODY
    SMTP_PORT

See https://hub.docker.com/r/craftypenguins/minimal-mail-sender/
