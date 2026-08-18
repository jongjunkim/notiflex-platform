# Notiflex 여정 기록

이 파일은 독자가 실제로 진행한 내용을 기록한다. AI가 각 챕터 완료 시 자동으로 업데이트한다.

## 진행 현황

| 챕터 | 서브챕터 | 상태 | 완료일 | 비고 |
|------|---------|------|--------|------|
| ch2 | 2.2 설치 확인 | ✅ | 2026-08-06 | Claude Code 동작 확인 |
| ch2 | 2.3 gcloud 설정 | ✅ | 2026-08-06 | project=gitaiops-notiflex-385f98, region=asia-northeast3, zone=asia-northeast3-a |
| ch2 | 2.4 GitHub 저장소 | ✅ | 2026-08-06 | notiflex-platform (public), CLAUDE.md + 디렉터리 구조 생성 |
| ch2 | 2.5 GKE 클러스터 | ✅ | 2026-08-06 | notiflex-cluster, e2-medium×2 Spot, Gateway API standard |
| ch2 | 2.6 빌드/배포 | ✅ | 2026-08-06 | api:v0.1.0 푸시, Pod 2개 Running, /health·/id 검증 |
| ch2 | 2.7 첫 커밋 | ✅ | 2026-08-06 | 초기 커밋 및 push |
| ch2 | 커스텀 스킬 `/update-docs` | ✅ | 2026-08-12 | `.claude/commands/update-docs.md` — 프로젝트 스코프 슬래시 커맨드 |
| ch3 | 3.2 GitOps 도구 | ✅ | 2026-08-12 | ArgoCD v3.5.1 설치, notiflex-smb Application → Synced/Healthy |
| ch3 | 3.3 기능 추가 | ✅ | 2026-08-17 | `/version` 엔드포인트 추가, api:v0.1.1 롤링 업데이트 (ArgoCD auto-sync). `git revert`로 v0.1.0 롤백 → 재복구 검증 완료 |
| ch3 | 3.4 CI | ✅ | 2026-08-17 | GitHub Actions `.github/workflows/ci.yaml` — app/ 변경 시 빌드→AR 푸시, 태그 `sha-<7자리>`. 첫 실행 1분 16초 성공 |
| ch3 | 3.5 CI-CD 연결 | ✅ | 2026-08-17 | CI가 빌드 후 매니페스트 태그를 갱신·커밋 → ArgoCD 자동 배포. 엔드투엔드 검증 완료 (코드 push → 약 4분 뒤 Pod 교체) |
| ch4 | 4.2 메트릭 모니터링 | ✅ | 2026-08-18 | kube-prometheus-stack (Helm) — Prometheus·Grafana·Alertmanager·kube-state-metrics·node-exporter. 수집 대상 16개 전부 up, Notiflex 대시보드 4패널 |
| ch4 | 4.3 로그 수집 | ✅ | 2026-08-18 | Loki 3.6.12 (SingleBinary, 5Gi PVC, 보존 72h) + Fluent Bit 5.1.1 DaemonSet ×2 → Grafana Loki 데이터소스. argocd·kube-system·monitoring 로그 수집 확인. **notiflex는 앱이 요청 로그를 찍지 않아 수집할 로그가 없다** |
| ch4 | 4.4 알림 | ⬜ | | |
| ch5 | 5.2 트래픽 관리 | ⬜ | | |
| ch5 | 5.3 무중단 배포 | ⬜ | | |
| ch6 | 6.1 캐시 | ⬜ | | |
| ch6 | 6.2 시크릿 관리 | ⬜ | | |
| ch6 | 6.3 Canary 전환 | ⬜ | | |
| ch7 | 7.2 멀티 노드풀 | ⬜ | | |
| ch7 | 7.3 App of Apps | ⬜ | | |
| ch7 | 7.4 멀티테넌시 | ⬜ | | |
| ch8 | 8.1 메시징 | ⬜ | | |
| ch8 | 8.2 트레이싱 | ⬜ | | |
| ch8 | 8.3 CronJob | ⬜ | | |
| ch9 | 9.1 저장소 분석 | ⬜ | | |
| ch9 | 9.2 회고 | ⬜ | | |
| ch9 | 9.3 온보딩 문서 | ⬜ | | |
| ch9 | 9.4 GitAIOps 분석 | ⬜ | | |
| ch9 | 9.5 마무리 | ⬜ | | |

## 도구 선택 기록

독자가 3-프롬프트 패턴(탐색→비교→실행)에서 실제로 선택한 도구와 이유를 기록한다.

| 영역 | 선택 | 검토한 대안 | 선택 이유 |
|------|------|-----------|----------|
| — | — | — | ch2는 선택지가 없는 환경 구성 단계 |
| GitOps 배포 도구 (ch3.2) | ArgoCD | Flux, Jenkins X, Spinnaker | Web UI로 Sync 상태를 눈으로 확인할 수 있어 학습에 유리. Flux는 ~100MB로 가볍지만 UI가 없다. Jenkins X·Spinnaker는 e2-medium 2대에 과중 |
| CI 도구 (ch3.4) | GitHub Actions | Cloud Build, GitLab CI, Jenkins | 코드가 이미 GitHub에 있어 별도 서버·웹훅 없이 YAML 한 파일로 동작. public 저장소라 실행 시간 무료. Cloud Build는 GitHub 트리거를 따로 붙여야 하고 로그를 GitHub 밖에서 봐야 한다 |
| 메트릭 모니터링 (ch4.2) | Prometheus + Grafana (kube-prometheus-stack) | Datadog, Google Cloud Monitoring, CloudWatch | K8s 모니터링 표준이고 비용이 없다. Helm 차트 하나로 6개 컴포넌트를 검증된 조합으로 설치한다. 4.3 Loki·8.2 Tempo가 같은 Grafana로 합쳐져 도구가 파편화되지 않는다. 관리형 도구는 PromQL·임계값 설정 과정을 감춰 4.4 알림 실습에 부족하다 |
| 로그 수집 (ch4.3) | Loki + Fluent Bit | ELK Stack, Google Cloud Logging, CloudWatch | Loki는 192Mi로 동작해 4GB 노드에 들어간다 (Elasticsearch는 최소 2Gi로 불가). 4.2의 Grafana에 데이터소스만 추가하면 메트릭과 같은 화면에서 시각을 맞춰 볼 수 있다. 라벨만 인덱싱해 저장 비용이 낮다. Cloud Logging은 이미 수집 중이지만 Grafana에 통합되지 않고 LogQL 실습이 불가능하다 |

## 현재 버전

| 컴포넌트 | 버전 | 변경 이력 |
|---------|------|----------|
| Go | 1.25 | 2026-08-06 최초 설정 (ch8 OTel SDK 요구사항 대비) |
| Notiflex 이미지 | `sha-372fa2c` (앱 버전 v0.1.2) | 2026-08-06 v0.1.0 최초 빌드 (digest `sha256:05b8906d…`, 2.47MB) → 2026-08-17 v0.1.1 (`/version` 추가, `sha256:bba1f801…`) → 2026-08-17 `sha-372fa2c` (CI 최초 자동 배포, `sha256:de666f35…`) |
| GKE | 1.35.6-gke.1250000 | 2026-08-06 클러스터 생성 |
| ArgoCD | v3.5.1 | 2026-08-12 설치 (`quay.io/argoproj/argocd:v3.5.1`) |
| kube-prometheus-stack | chart 88.3.0 (operator v0.93.0) | 2026-08-18 설치, Helm revision 2. Prometheus v3.13.2, Grafana 13.1.3, Alertmanager v0.33.1 — requests 축소 적용 |
| Loki | chart 7.3.0 / Loki 3.6.12 | 2026-08-18 설치. SingleBinary, PVC 5Gi(standard-rwo), 보존 72h |
| Fluent Bit | chart 0.58.1 / v5.1.1 | 2026-08-18 설치. DaemonSet 2개, loki 출력 플러그인 (helm revision 3) |
| Kafka | | |
| OTel SDK | | |

## 현재 리소스

| 노드풀 | 머신 타입 | 노드 수 | 주요 워크로드 |
|--------|----------|---------|-------------|
| default-pool | e2-medium (Spot, 30GB) | 2 | notiflex-api ×2, ArgoCD 7개, 관측 가능성 스택 7개 |

**CPU requests 실측 (2026-08-18, ch4.2 완료 후)**

| 항목 | 값 |
|------|-----|
| 노드당 allocatable | **940m** (e2-medium 2 vCPU 중 GKE 시스템 예약분 제외) |
| 총 가용 | **1880m** |
| kube-system | 1198m |
| ch4.3 설치 후 총 요청 | 1569m (노드1 873m 92% / 노드2 696m 74%) |
| 잔여 | **약 311m** |

> ⚠️ `shared/resource-budget.md`는 총 가용을 3200m으로 가정하지만 **실측은 1880m**이다. GKE
> 시스템 예약이 예산표보다 크다. ArgoCD는 requests를 지정하지 않아 예산표의 500m이 실제로는
> 0m으로 잡히는 반면, kube-system이 1198m을 점유해 순효과는 예산표보다 **훨씬 빠듯하다**.
> ch6에서 CSI DaemonSet 240m(축소 불가)이 들어오면 잔여 341m으로는 어렵다. 그 전에
> 관측 가능성 스택 requests를 5m대로 더 줄여야 한다.

## 인프라 현황

| 항목 | 값 |
|------|-----|
| GCP 프로젝트 | `gitaiops-notiflex-385f98` (조직 미소속) |
| 리전 / 존 | `asia-northeast3` / `asia-northeast3-a` |
| GKE 클러스터 | `notiflex-cluster` (Standard, Zonal) |
| kubectl 컨텍스트 | `gke-sysnet4admin_book_gitaiops` |
| Artifact Registry | `asia-northeast3-docker.pkg.dev/gitaiops-notiflex-385f98/notiflex` |
| Gateway API | CHANNEL_STANDARD (GatewayClass 4개 Accepted) |
| 배포 전략 | Rolling Update (기본) |
| GitOps | ArgoCD (`argocd` ns) → `k8s/smb` 자동 동기화 (prune·selfHeal 활성) |
| 관측 가능성 | Prometheus + Grafana + Alertmanager + Loki + Fluent Bit (`monitoring` ns, Helm 관리 — ArgoCD 밖) |
| CI/CD | GitHub Actions `.github/workflows/ci.yaml` (`app/**` 변경 시 트리거) → 빌드·푸시 → 매니페스트 태그 갱신 커밋 → ArgoCD가 배포 |
| Actions 권한 | 저장소 `default_workflow_permissions=write` + 워크플로 `contents: write` (둘 다 필요) |
| CI 서비스 계정 | `notiflex-ci@gitaiops-notiflex-385f98.iam.gserviceaccount.com` — 권한 `roles/artifactregistry.writer` 하나 |
| GitHub Secrets | `GCP_SA_KEY`(SA JSON 키), `GCP_PROJECT_ID` |
| 이미지 태그 규칙 | 수동 배포는 `v0.1.x`, CI 빌드는 `sha-<커밋 7자리>` |

## 트러블슈팅 이력

독자가 겪은 문제와 해결 방법을 기록한다. 같은 문제를 다시 겪지 않도록 한다.

| 챕터 | 문제 | 해결 |
|------|------|------|
| 2.4 | GitHub 저장소가 이미 존재하고 visibility가 public이었다 | 빈 저장소여서 clone만 하고 진행. private 전환은 보류 |
| 2.5 | 클러스터 생성 직후 `GatewayClass`가 조회되지 않았다 | `gatewayApiConfig.channel=CHANNEL_STANDARD` 확인 후 대기. 약 1분 45초 뒤 4개 생성됨. `clusters update`는 불필요했다 |
| 2.5 | `get-credentials`를 재실행하니 컨텍스트가 GKE 기본 긴 이름으로 되돌아가 2개가 되었다 | 중복 컨텍스트를 삭제하고 `gke-sysnet4admin_book_gitaiops`로 `use-context` |
| 2.5 | GCP 콘솔 프로젝트 목록에 프로젝트가 보이지 않았다 | 프로젝트가 조직(`rkdlem196-org`) 밖에 있어 조직 범위 선택기에서 누락. 선택기를 "조직 없음"으로 전환하거나 `?project=` 링크로 직접 접근 |
| 2.6 | `/id`를 여러 번 호출해도 항상 같은 Pod이 응답했다 | 정상 동작. `port-forward`는 Service를 우회해 Pod 하나에 직접 연결한다. Service 분산은 클러스터 내부에서 호출해야 확인 가능 |
| 2.x | Spot VM 선점으로 노드 2대가 재생성되고 `Error` 상태 Pod이 남았다 (2026-08-12) | Spot의 정상 동작. 새 노드에 Pod이 재스케줄되면 복구된다. 잔여 `Error` Pod은 `kubectl delete pod --field-selector status.phase=Failed -n notiflex`로 정리. 재발이 잦으면 노드풀을 온디맨드로 전환 |
| 3.3 | git push 후에도 5분 가까이 v0.1.0 Pod이 그대로였다 | 정상. ArgoCD auto-sync는 기본 3분 주기로 Git을 폴링한다. 즉시 반영하려면 UI에서 Sync 또는 `argocd app sync notiflex-smb`. 3.5에서 CI가 push하면 같은 흐름으로 자동 반영된다 |
| 3.3 | 롤백 후 재복구 push를 했는데 5분이 지나도 ArgoCD가 v0.1.0에 머물렀다 (`status.sync.revision`이 이전 커밋 `93c9381`에 고정, reconcile은 계속 돌고 있었음) | repo-server가 캐시된 Git 리비전을 붙들고 있었다. `kubectl annotate application notiflex-smb -n argocd argocd.argoproj.io/refresh=hard --overwrite`로 캐시 무효화 → 5초 만에 최신 커밋 인식. **증상 구분법**: Application이 `Synced`인데 `status.sync.revision`이 `git ls-remote origin main`과 다르면 캐시 문제다 |
| 3.4 | 서비스 계정 생성 직후 `add-iam-policy-binding`이 `Service account ... does not exist`로 실패했다 | IAM 전파 지연. 생성과 바인딩을 연달아 실행하면 발생한다. 재시도 루프(10초 간격)로 해결 |
| 3.4 | `actions/setup-go`에 `cache-dependency-path: app/go.sum`을 주면 실패한다 | 이 앱은 표준 라이브러리만 써서 `go.sum`이 없다. `cache: false`로 설정. 나중에 외부 의존성이 생기면 `go.sum`이 만들어지므로 캐시를 켠다 |
| 4.3 | Loki 차트 기본값으로 설치하면 Pod이 뜨지 않는다 | 기본 캐시가 memcached에 **chunksCache 8192Mi + resultsCache 1024Mi**를 요청한다(노드 전체가 4GB). values에서 `chunksCache.enabled: false`, `resultsCache.enabled: false`. 함께 `gateway`·`lokiCanary`·`test`도 끄고 `read/write/backend` replicas를 0으로 둔다 (SingleBinary 모드) |
| 4.3 | 가드레일의 `grafana.datasource.isDefault` 키가 loki 차트 7.x에 없다 | 이 버전은 Grafana 데이터소스를 만들어주지 않는다. `k8s/monitoring/loki-datasource.yaml`(label `grafana_datasource=1`)로 직접 정의하고 **`isDefault: false`**를 명시한다. Prometheus가 이미 default라 둘 다 default면 Grafana가 기동에 실패한다 |
| 4.3 | Grafana에 로그가 `{"time":...,"_p":"F","log":"..."}` JSON으로 보인다 | CRI 멀티라인 파서가 붙이는 `_p`와 수집 시각 `time`이 남아 키가 여러 개라 `drop_single_key`가 작동하지 않는다. `remove_keys kubernetes, stream, time, _p` + `drop_single_key raw` (`on`은 따옴표로 감싼 JSON 문자열이 된다) |
| 4.2 | Grafana에 올바른 비밀번호를 넣어도 `Invalid username or password` | **비밀번호 문제가 아니라 계정 잠금.** Grafana는 연속 실패 5회 이상이면 계정을 5분간 차단하고, 차단 중에도 브라우저에는 똑같이 "Invalid username or password"로 표시한다. 진단은 `kubectl logs deploy/kube-prometheus-grafana -c grafana \| grep locked` — `too many consecutive incorrect login attempts`가 보이면 잠금이다. 해결: 5분 대기하거나 `kubectl rollout restart deploy/kube-prometheus-grafana -n monitoring` (이 설치는 emptyDir이라 재시작 시 잠금 기록이 사라지고 비밀번호는 Secret에서 재주입된다. 대시보드는 ConfigMap, 데이터소스는 프로비저닝이라 유실 없음). 재시작 후에는 **port-forward도 다시 띄워야 한다** (옛 Pod을 가리킴) |
| 4.2 | `helm`이 설치돼 있지 않았다 | `brew install helm` (v4.2.4). 2장에서 다루지 않는 도구다 |
| 4.2 | Prometheus 수집 대상 중 `coredns` 2개가 계속 Down | **GKE는 CoreDNS가 아니라 kube-dns를 쓴다.** 차트는 CoreDNS 규격대로 9153 포트를 긁지만 kube-dns는 거기에 메트릭이 없다. values에 `coreDns.enabled: false`. 같은 이유로 `kubeEtcd`·`kubeControllerManager`·`kubeScheduler`·`kubeProxy`도 끈다 (GKE 관리형 컨트롤 플레인은 노출하지 않음) |
| 3.5 | 워크플로를 고쳐 push했는데 실행이 아예 생성되지 않았다 | 커밋 메시지 **본문**에 `[skip` `ci]` 문자열을 설명용으로 적은 것을 GitHub이 실제 지시로 읽었다. 제목뿐 아니라 본문까지 스캔한다. 이 키워드는 커밋 메시지에 언급하지 않는다 |
| 3.5 | CI가 액션 다운로드에서 `429 Too Many Requests`로 실패 경고 | 일시적 GitHub 혼잡. Actions가 자동 백오프 후 재시도해 성공했다. 실패로 처리하지 않아도 된다 |
| 2.x | 선점 직후 `kube-dns`가 `FailedScheduling — Insufficient cpu`로 뜨지 못했다 | 노드 1대만 Ready인 동안 CPU가 부족해 발생. 두 번째 노드가 Ready가 되면 해소된다. 상시 발생하면 `shared/resource-budget.md` 기준으로 노드 수·머신 타입을 재검토 |
