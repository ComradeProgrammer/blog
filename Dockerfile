FROM node:18-alpine as frontendbuilder
WORKDIR /blog
COPY . .
RUN  cd web && npm install && npm run build

FROM golang:1.18 as backendbuilder
WORKDIR /blog
COPY . .
RUN go build -o blog cmd/myblog/main.go


FROM debian:latest 
WORKDIR /blog
COPY --from=backendbuilder /blog/blog .
COPY --from=frontendbuilder /blog/web/build ./web/build
CMD ["./blog"]

