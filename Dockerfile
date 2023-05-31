FROM golang:1.19

WORKDIR /signup

COPY pb/ pb/
COPY deploy/ deploy/
COPY api/ api/
COPY database/ database/

COPY go.mod go.sum main.go .env ./
RUN go mod download -x

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o app main.go" --command="./app"