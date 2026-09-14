FROM golang:1.27.1-bookworm

WORKDIR /HR_system

COPY . .

RUN go mod tidy
RUN go build -o /HR_system/exe main.go

CMD ["/HR_system/exe"]