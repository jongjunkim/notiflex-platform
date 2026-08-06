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
| ch3 | 3.2 GitOps 도구 | ⬜ | | |
| ch3 | 3.3 기능 추가 | ⬜ | | |
| ch3 | 3.4 CI | ⬜ | | |
| ch3 | 3.5 CI-CD 연결 | ⬜ | | |
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

## 현재 버전

| 컴포넌트 | 버전 | 변경 이력 |
|---------|------|----------|
| Go | 1.25 | 2026-08-06 최초 설정 (ch8 OTel SDK 요구사항 대비) |
| Notiflex 이미지 | v0.1.0 | 2026-08-06 최초 빌드 (digest `sha256:05b8906d…`, 2.47MB) |
| GKE | 1.35.6-gke.1250000 | 2026-08-06 클러스터 생성 |
| ArgoCD | | |
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

## 트러블슈팅 이력

독자가 겪은 문제와 해결 방법을 기록한다. 같은 문제를 다시 겪지 않도록 한다.

| 챕터 | 문제 | 해결 |
|------|------|------|
| 2.4 | GitHub 저장소가 이미 존재하고 visibility가 public이었다 | 빈 저장소여서 clone만 하고 진행. private 전환은 보류 |
| 2.5 | 클러스터 생성 직후 `GatewayClass`가 조회되지 않았다 | `gatewayApiConfig.channel=CHANNEL_STANDARD` 확인 후 대기. 약 1분 45초 뒤 4개 생성됨. `clusters update`는 불필요했다 |
| 2.5 | `get-credentials`를 재실행하니 컨텍스트가 GKE 기본 긴 이름으로 되돌아가 2개가 되었다 | 중복 컨텍스트를 삭제하고 `gke-sysnet4admin_book_gitaiops`로 `use-context` |
| 2.5 | GCP 콘솔 프로젝트 목록에 프로젝트가 보이지 않았다 | 프로젝트가 조직(`rkdlem196-org`) 밖에 있어 조직 범위 선택기에서 누락. 선택기를 "조직 없음"으로 전환하거나 `?project=` 링크로 직접 접근 |
| 2.6 | `/id`를 여러 번 호출해도 항상 같은 Pod이 응답했다 | 정상 동작. `port-forward`는 Service를 우회해 Pod 하나에 직접 연결한다. Service 분산은 클러스터 내부에서 호출해야 확인 가능 |
