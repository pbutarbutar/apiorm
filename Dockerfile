FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM scratch
COPY --from=builder /app/pintekid /app/

EXPOSE 9000

CMD [ "/app/pintekid" ]



#SKIP BY PARULIAN
#RUN mkdir /app
#RUN . /app
#ENV GOPATH=/app
#RUN go get github.com/qodrorid/godaemon
#RUN go get github.com/gorilla/mux
#RUN go get github.com/gorilla/handlers
#RUN go get github.com/jinzhu/gorm
#UN go get github.com/dgrijalva/jwt-go
#RUN go get github.com/joho/godotenv
#UN go get github.com/badoux/checkmail
#RUN go get golang.org/x/crypto/bcrypt
#RUN go build -o main .

