FROM golang:1.20
WORKDIR /usr/src/app
COPY go.mod ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o whoami
CMD [ "./whoami" ]

EXPOSE 8088