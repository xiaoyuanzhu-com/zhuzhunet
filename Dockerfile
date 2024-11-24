FROM golang:1.23 AS api-builder
WORKDIR /zhuzhunet
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build .

FROM node:22 AS ui-builder
COPY ui /zhuzhunet/ui
WORKDIR /zhuzhunet/ui
RUN npm install
RUN npm run build

FROM debian:12
COPY --from=api-builder /zhuzhunet/zhuzhunet /zhuzhunet/zhuzhunet
COPY --from=ui-builder /zhuzhunet/ui/dist /zhuzhunet/ui/dist
WORKDIR /zhuzhunet
ENV CONFIG=/config
EXPOSE 27831
CMD ["/zhuzhunet/zhuzhunet"]
