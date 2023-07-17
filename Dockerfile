ARG IMAGE_URL
FROM ${IMAGE_URL}/baseimages/codefin-alpine-base-image:1.18.0 as builder
RUN update-ca-certificates
# RUN GO111MODULE=on \
#         go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1 \
#         google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0

WORKDIR /app
ADD cmd /app/cmd
ADD cert /app/cert
ADD api /app/api
ADD pkg /app/pkg
ADD third_party /app/third_party
ADD internal /app/internal
ADD vendor /app/vendor
ADD Makefile /app/Makefile
ADD go.mod /app/go.mod
ADD go.sum /app/go.sum

# RUN make bin
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o bin/server-linux-amd64 -mod=vendor cmd/server/*.go
RUN	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o bin/campaignctl -mod=vendor cmd/tools/campaignctl.go

FROM ${IMAGE_URL}/baseimages/alpine:latest
COPY --from=builder /app/bin /app/bin
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
RUN cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime
ADD ./cert /app/cert
ADD ./tmp_image /app/tmp_image
WORKDIR /app
CMD [ "/app/bin/server-linux-amd64" ]
