// Notiflex API — B2B 알림 SaaS 플랫폼의 ID 발급 서비스.
// Go 표준 라이브러리만 사용한다.
package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync/atomic"
	"time"
)

// version은 배포된 이미지 태그와 일치시킨다. 이미지를 새로 빌드할 때마다 함께 올린다.
const version = "v0.1.3"

// serviceName은 여러 서비스가 늘어났을 때 응답만 보고 구분하기 위한 식별자이다.
const serviceName = "notiflex-api"

// logger는 stdout으로 logfmt 형식(key=value)을 찍는다.
// 컨테이너는 stdout/stderr만 로그로 취급하므로 파일로 쓰지 않는다.
// logfmt는 사람이 읽을 수 있으면서 Loki에서 `| logfmt`로 필드 파싱이 되는 형식이다.
var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

// counter는 /id 요청마다 증가하는 인메모리 카운터이다.
// Pod마다 독립적이므로 replicas가 여러 개면 ID가 Pod별로 따로 증가한다.
var counter atomic.Uint64

// podName은 요청을 처리한 Pod을 식별한다. Kubernetes에서는 HOSTNAME이 Pod 이름이다.
var podName = hostname()

func hostname() string {
	if h := os.Getenv("HOSTNAME"); h != "" {
		return h
	}
	h, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return h
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		logger.Error("failed to encode response", "error", err)
	}
}

// handleHealth는 readiness/liveness probe가 호출하는 상태 확인 엔드포인트이다.
func handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// handleID는 순차 ID를 발급하고, 발급한 Pod 이름을 함께 반환한다.
func handleID(w http.ResponseWriter, r *http.Request) {
	id := counter.Add(1)
	writeJSON(w, http.StatusOK, map[string]string{
		"id":           strconv.FormatUint(id, 10),
		"generated_by": podName,
	})
}

// handleVersion은 지금 돌고 있는 이미지가 어느 버전인지 응답으로 확인하게 해준다.
// 롤링 업데이트 중에는 Pod마다 다른 버전이 나올 수 있어, 교체 진행 상황을 눈으로 볼 수 있다.
func handleVersion(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"service":      serviceName,
		"version":      version,
		"go_version":   runtime.Version(),
		"generated_by": podName,
	})
}

// statusRecorder는 핸들러가 쓴 상태 코드와 바이트 수를 기록한다.
// http.ResponseWriter는 이 값을 되읽을 방법을 제공하지 않으므로 감싸서 가로챈다.
type statusRecorder struct {
	http.ResponseWriter
	status int
	bytes  int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func (r *statusRecorder) Write(b []byte) (int, error) {
	// 핸들러가 WriteHeader 없이 바로 Write하면 net/http가 200을 쓴다. 그 경우를 맞춰준다.
	if r.status == 0 {
		r.status = http.StatusOK
	}
	n, err := r.ResponseWriter.Write(b)
	r.bytes += n
	return n, err
}

// withLogging은 모든 요청을 한 줄씩 기록한다.
//
// /health는 제외한다. readiness가 5초, liveness가 10초마다 호출하므로 Pod 하나당
// 분당 18줄이 쌓인다. 정보 가치는 없는데 로그 저장소를 채우고, 정작 필요한 요청 로그를
// 묻어버린다. probe 실패는 Pod 이벤트와 메트릭(4.2 대시보드)에서 확인한다.
func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/health" {
			next.ServeHTTP(w, r)
			return
		}

		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w}
		next.ServeHTTP(rec, r)

		logger.Info("request",
			"method", r.Method,
			"path", r.URL.Path,
			"status", rec.status,
			"duration_ms", float64(time.Since(start).Microseconds())/1000,
			"bytes", rec.bytes,
			"remote", r.RemoteAddr,
			"user_agent", r.UserAgent(),
			"pod", podName,
		)
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handleHealth)
	mux.HandleFunc("GET /id", handleID)
	mux.HandleFunc("GET /version", handleVersion)

	logger.Info("starting",
		"service", serviceName,
		"version", version,
		"port", port,
		"pod", podName,
	)
	if err := http.ListenAndServe(":"+port, withLogging(mux)); err != nil {
		logger.Error("server terminated", "error", err)
		os.Exit(1)
	}
}
