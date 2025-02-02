# Stage 1: Builder
FROM golang:alpine AS orders

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 3000

CMD [ "go", "run", "/app/services/orders/*.go" ]


FROM golang:alpine AS kitchen

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 3000

CMD [ "go", "run", "/app/services/kitchen/*.go" ]
