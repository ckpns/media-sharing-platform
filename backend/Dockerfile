FROM golang:1.18.1

ENV DB_CONNECTION_STRING=postgres://postgres:postgrespw@postgresDB:5432
ENV DB_AUTO_MIGRATE=true
ENV AUTH_HMAC_SECRET=topsecretkey
ENV AUTH_EXPIRATION_MINUTES=10080
ENV MAX_MULTIPART_MEMORY=20971520
ENV SUPPORTED_FILE_EXTENSIONS=.png|.jpg|.jpeg|.mp4
ENV HOST=0.0.0.0
ENV PORT=3000

EXPOSE 3000

COPY src/ /app
WORKDIR /app

RUN go build -o api/app

CMD ["./api/app"]