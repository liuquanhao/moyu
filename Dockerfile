FROM alpine:3.17.1 as builder
COPY . .
RUN apk add go nodejs npm
RUN cd frontend && npm install && npm run build && \
    cd ../backend && go get . && go build .

FROM alpine:3.17.1 as deploy
COPY --from=builder backend/moyu .
ENV PORT=8080
ENTRYPOINT /moyu