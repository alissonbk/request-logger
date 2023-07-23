FROM golang:1.16-alpine

# Set destination for COPY
WORKDIR /go/src
ENV PATH="go/bin:${PATH}"
ENV CGO_ENABLED=1

#OS Dependencies
RUN apk add build-base mpc1-dev gcc glib-dev linux-headers musl-dev libpcap-dev

#Donwload
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
ADD src ./src


#Build
RUN go build -o /requestlogger

EXPOSE 8080

CMD [ "/requestlogger" ]
