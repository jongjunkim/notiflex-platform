// Notiflex API — B2B 알림 SaaS 플랫폼의 ID 발급 서비스.
// Go 표준 라이브러리만 사용한다.
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync/atomic"
)

// version은 배포된 이미지 태그와 일치시킨다. 이미지를 새로 빌드할 때마다 함께 올린다.
const version = "v0.1.1"

// serviceName은 여러 서비스가 늘어났을 때 응답만 보고 구분하기 위한 식별자이다.
const serviceName = "notiflex-api"

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
		log.Printf("failed to encode response: %v", err)
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

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handleHealth)
	mux.HandleFunc("GET /id", handleID)
	mux.HandleFunc("GET /version", handleVersion)

	log.Printf("%s %s listening on :%s (pod=%s)", serviceName, version, port, podName)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("server terminated: %v", err)
	}
}
