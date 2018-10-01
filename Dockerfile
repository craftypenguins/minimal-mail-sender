FROM scratch
ADD minimal-mail-sender /
ENTRYPOINT ["/minimal-mail-sender"]
