FROM golang:1.24.5-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o go-email

FROM build AS prod

WORKDIR /prod

COPY --from=build /app/go-email ./

EXPOSE 8000

CMD [ "/prod/go-email","-port=:8000" ]
