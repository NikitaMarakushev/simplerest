FROM 1.12-alpine3.11

LABEL maintainer="NikitaMarakushev <nikitamarakushev@yandex.ru>"

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

RUN mkdir /var/projects/simplerest
WORKDIR  /var/projects/simplerest

COPY . .
COPY .env .

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build -o /build

EXPOSE 8080

CMD [ "/build" ]