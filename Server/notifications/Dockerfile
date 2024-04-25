FROM golang:1.22-alpine

WORKDIR /notifications

RUN go get -u -v github.com/appleboy/gorush

EXPOSE ${port}

CMD ["gorush -android -m your message -k API Key -t Device token"]

