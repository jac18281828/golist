ARG VERSION=102221

FROM jac18281828/godev:${VERSION} 

ARG PROJECT=golist
WORKDIR /workspaces/${PROJECT}

COPY linkedlist linkedlist/
RUN go mod init github.com/jac18281828/golist
RUN go build ./...
RUN go test ./...
CMD go test ./...
