FROM golang:1.15.0-alpine AS builder
LABEL stage=builder

RUN apk add --no-cache git 

# Install UPX
ADD https://github.com/upx/upx/releases/download/v3.96/upx-3.96-amd64_linux.tar.xz /usr/local
RUN xz -d -c /usr/local/upx-3.96-amd64_linux.tar.xz | \
    tar -xOf - upx-3.96-amd64_linux/upx > /bin/upx
RUN chmod a+x /bin/upx

WORKDIR /go/src/tcphub
ENV GOPATH /go

COPY src/ ./
RUN echo $GOPATH
RUN go get 
RUN CGO_ENABLED=0 GOOS=linux go build -a tcphub.go
RUN /bin/upx tcphub



FROM alpine:3.12 AS final
WORKDIR /
COPY --from=builder /go/src/tcphub/tcphub .
COPY src/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

#RUN apk add --no-cache bash aws-cli
ENTRYPOINT ["/entrypoint.sh"]
CMD [ "/tcphub" ]
