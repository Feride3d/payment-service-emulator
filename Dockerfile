FROM golang:1.16-alpine as builder

ENV BIN_FILE /opt/payment-service-emulator/payment-service-emulator
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} ${CODE_DIR}/cmd/payment-service-emulator/*

FROM alpine:3

LABEL SERVICE="payment-service-emulator"
LABEL MAINTAINERS="feride3d@icloud.com"

ENV BIN_FILE "/opt/payment-service-emulator/payment-service-emulator"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

ENV CONFIG_FILE /etc/payment-service-emulator/config.yaml
COPY ./configs/config.yaml ${CONFIG_FILE}

CMD ${BIN_FILE} -config ${CONFIG_FILE}