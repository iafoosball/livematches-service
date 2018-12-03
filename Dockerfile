# STEP 1 build executable binary

FROM golang:1.10 as builder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

#Copy the service
RUN mkdir -p /go/src/github.com/iafoosball/livematches-service
WORKDIR /go/src/github.com/iafoosball/livematches-service
COPY . .

#Run dep ensure
RUN dep ensure -vendor-only

#Install the service
WORKDIR /go/src/github.com/iafoosball/livematches-service/main
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o livematches .

# STEP 2 build a small image
# start from scratch
# FROM scratch


FROM alpine:latest
RUN apk --no-cache add ca-certificates
ARG mHost
ARG mPort

# Copy our static executable
COPY --from=builder /go/src/github.com/iafoosball/livematches-service/main/livematches .
CMD ./livematches --matchesHost=$mHost --matchesPort=$mPort