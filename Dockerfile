FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go get gopkg.in/tucnak/telebot.v2
RUN go get github.com/vjeantet/jodaTime
RUN go build -o main . 


EXPOSE 80

CMD ["/app/main"]