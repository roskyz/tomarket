FROM golang:alpine AS builder
WORKDIR /
ARG REF
RUN apk add git &&\
    git clone https://github.com/roskyz/tomarket.git

RUN if [[ -z "${REF}" ]]; then \
        echo "No specific commit provided, use the latest one." \
    ;else \
        echo "Use commit ${REF}" &&\
        cd tomarket &&\
        git checkout ${REF} \
    ;fi

RUN cd tomarket &&\
    go build -o tomarket .

FROM alpine
WORKDIR /
RUN apk add --no-cache tzdata ca-certificates
COPY --from=builder /tomarket/tomarket /usr/local/bin/tomarket

ENTRYPOINT ["/usr/local/bin/tomarket"]