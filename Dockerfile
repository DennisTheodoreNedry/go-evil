FROM golang:1.19-bullseye

WORKDIR /app

COPY . ./

RUN apt update && apt install upx-ucl  -y
RUN make dependencies

RUN make

ENTRYPOINT ["/app/gevil"]