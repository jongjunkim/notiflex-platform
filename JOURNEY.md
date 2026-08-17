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
| ch4 | 4.2 메트릭 모니터링 | ⬜ | | |
| ch4 | 4.3 로그 수집 | ⬜ | | |
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
| CI 도구 (ch3.4) | GitHub Actions | Cloud Build, GitLab CI, Jenkins | 코드가 이미 GitHub에 있어 별도 서버·웹훅 없이 YAML 한 파일로 동작. public 저장소라 실행 시간 무료. Cloud Build는 GitHub 트리거를 따로 붙여야 하고 로그를 GitHub 밖에서 봐야 한다 |
| GitOps 배포 도구 (ch3.2) | ArgoCD | Flux, Jenkins X, Spinnaker | Web UI로 Sync 상태를 눈으로 확인할 수 있어 학습에 유리. Flux는 ~100MB로 가볍지만 UI가 없다. Jenkins X·Spinnaker는 e2-medium 2대에 과중 |

## 현재 버전

| 컴포넌트 | 버전 | 변경 이력 |
|---------|------|----------|
| Go | 1.25 | 2026-08-06 최초 설정 (ch8 OTel SDK 요구사항 대비) |
| Notiflex 이미지 | `sha-372fa2c` (앱 버전 v0.1.2) | 2026-08-06 v0.1.0 최초 빌드 (digest `sha256:05b8906d…`, 2.47MB) → 2026-08-17 v0.1.1 (`/version` 추가, `sha256:bba1f801…`) → 2026-08-17 `sha-372fa2c` (CI 최초 자동 배포, `sha256:de666f35…`) |
| GKE | 1.35.6-gke.1250000 | 2026-08-06 클러스터 생성 |
| ArgoCD | v3.5.1 | 2026-08-12 설치 (`quay.io/argoproj/argocd:v3.5.1`) |
| Kafka | | |
| OTel SDK | | |

## 현재 리소스

| 노드풀 | 머신 타입 | 노드 수 | 주요 워크로드 |
|--------|----------|---------|-------------|
| default-pool | e2-medium (Spot, 30GB) | 2 | notiflex-api ×2 |

**CPU requests 누적**: 100m / 가용 약 3200m (잔여 약 3100m)

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
| 3.5 | 워크플로를 고쳐 push했는데 실행이 아예 생성되지 않았다 | 커밋 메시지 **본문**에 `[skip` `ci]` 문자열을 설명용으로 적은 것을 GitHub이 실제 지시로 읽었다. 제목뿐 아니라 본문까지 스캔한다. 이 키워드는 커밋 메시지에 언급하지 않는다 |
| 3.5 | CI가 액션 다운로드에서 `429 Too Many Requests`로 실패 경고 | 일시적 GitHub 혼잡. Actions가 자동 백오프 후 재시도해 성공했다. 실패로 처리하지 않아도 된다 |
| 2.x | 선점 직후 `kube-dns`가 `FailedScheduling — Insufficient cpu`로 뜨지 못했다 | 노드 1대만 Ready인 동안 CPU가 부족해 발생. 두 번째 노드가 Ready가 되면 해소된다. 상시 발생하면 `shared/resource-budget.md` 기준으로 노드 수·머신 타입을 재검토 |
