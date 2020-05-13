FROM zenika/alpine-chrome:81

USER root
RUN apk add --no-cache gcc git make musl-dev go

RUN go get gopkg.in/alessio/shellescape.v1
RUN go get github.com/CovenantSQL/CookieScanner

ENV GOPATH="/root/go"
ENV PATH="$PATH:$GOPATH/bin"

ENTRYPOINT ["go", "run", "server.go"]