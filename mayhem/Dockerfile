FROM golang:1.25 AS builder

COPY . /funnel
WORKDIR /funnel

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN apt update && apt install -y clang

ENV GOFLAGS="-mod=mod"

RUN go-fuzz-build -libfuzzer -o mayhem/fuzz.a github.com/ohsu-comp-bio/funnel/mayhem && \
    clang -fsanitize=fuzzer mayhem/fuzz.a -o mayhem/fuzz.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /funnel/mayhem/fuzz.libfuzzer /
