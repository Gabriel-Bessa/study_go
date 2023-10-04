FROM golang:1.19

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./
# Build
RUN go build

EXPOSE 8080
# Run
CMD ["/study_go"]