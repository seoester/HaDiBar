# iron/go is the alpine image with only ca-certificates added
FROM scratch

WORKDIR /app

# Now just add the binary
COPY HaDiBar HaDiBar
COPY settings.json settings.json
COPY users.json users.json
COPY accounts.json accounts.json
COPY webapp webapp

EXPOSE 8080/tcp

ENTRYPOINT ["./HaDiBar"]