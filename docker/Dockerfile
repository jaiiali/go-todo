##
# builder
##
FROM golang AS builder

WORKDIR /src
COPY . /src
RUN make build BUILD_OUTPUT=./bin/app

##
# app
##
FROM debian AS app

COPY --from=builder /src/bin/app /go/bin/app
CMD ["/go/bin/app"]
