# 빌드 스테이지 — Go 툴체인은 최종 이미지에 포함하지 않는다.
FROM golang:1.25-alpine AS builder

WORKDIR /src

# 의존성 파일을 먼저 복사해 레이어 캐시를 활용한다.
COPY go.mod ./
RUN go mod download

COPY . .

# CGO를 끄면 libc에 의존하지 않는 정적 바이너리가 나온다. scratch에서 실행하려면 필수다.
# -ldflags="-s -w"로 디버그 심볼을 제거해 바이너리 크기를 줄인다.
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /notiflex-api .

# 실행 스테이지 — 셸도 패키지 매니저도 없는 빈 이미지.
FROM scratch

COPY --from=builder /notiflex-api /notiflex-api

EXPOSE 8080

ENTRYPOINT ["/notiflex-api"]
