# Notiflex Platform

## 프로젝트 개요

Notiflex는 B2B 알림 SaaS 플랫폼이다. 이 저장소는 애플리케이션 코드와 Kubernetes 매니페스트, CI 파이프라인을 함께 관리하는 GitOps 저장소이다.

## 기술 스택

| 구분 | 내용 |
|------|------|
| 언어 | Go 1.25 (표준 라이브러리만 사용, 외부 웹 프레임워크 없음) |
| 컨테이너 | 멀티스테이지 빌드 → `scratch` 베이스 이미지 |
| 오케스트레이션 | GKE Standard (Zonal), Spot VM |
| GitOps | ArgoCD (자동 동기화, prune·selfHeal 활성) |
| CI/CD | GitHub Actions → Artifact Registry → 매니페스트 태그 갱신 커밋 |

## GCP 설정

| 항목 | 값 |
|------|-----|
| 프로젝트 ID | `gitaiops-notiflex-385f98` |
| 리전 | `asia-northeast3` (서울) |
| 존 | `asia-northeast3-a` |
| Artifact Registry | `asia-northeast3-docker.pkg.dev/gitaiops-notiflex-385f98/notiflex` |

## 디렉터리 구조

```
notiflex-platform/
├── CLAUDE.md
├── app/               # Go 애플리케이션
├── k8s/
│   └── smb/           # K8s 매니페스트
├── argocd/            # ArgoCD Application 정의
└── .github/
    └── workflows/
        └── ci.yaml    # 빌드 → 푸시 → 매니페스트 태그 갱신
```

## 행동 규칙

1. **항상 확인 후 실행한다.** 리소스를 생성·변경·삭제하기 전에 무엇을 어떻게 바꿀지 먼저 설명한다.
2. **변경 전 현재 상태를 확인한다.** 매니페스트를 수정하기 전에 클러스터의 실제 상태(`kubectl get`)와 Git의 상태를 먼저 읽는다.
3. **kubectl에는 컨텍스트를 명시한다.** 잘못된 클러스터를 대상으로 실행하지 않도록 모든 kubectl 명령에 `--context gke-sysnet4admin_book_gitaiops`를 지정한다.
4. **클러스터에 직접 수정하지 않는다.** 배포 변경은 매니페스트를 고치고 Git에 커밋하여 ArgoCD가 반영하도록 한다.
5. **버전은 고정한다.** 이미지 태그에 `latest`를 쓰지 않는다. CI가 빌드하는 이미지는 `sha-<커밋 7자리>`, 손으로 배포하는 이미지는 `v0.1.x`를 쓴다.
6. **매니페스트 이미지 태그는 CI가 관리한다.** `app/`을 고쳐 push하면 CI가 `k8s/smb/deployment.yaml`의 태그를 갱신해 커밋한다. 사람이 같은 줄을 동시에 고치면 충돌하므로, 손으로 바꾸기 전에 `git pull`을 먼저 한다.

## CI/CD 파이프라인

`app/**`이 바뀐 push만 CI를 트리거한다 (문서·매니페스트만 바뀌면 빌드하지 않는다).

```
코드 push → go vet/test → 이미지 빌드 → Artifact Registry 푸시
         → k8s/smb/deployment.yaml 태그 갱신 커밋 → ArgoCD 폴링(최대 3분) → 롤링 배포
```

- **CI는 클러스터에 접근하지 않는다.** CI 서비스 계정 권한은 `roles/artifactregistry.writer` 하나뿐이고, 배포는 클러스터 안의 ArgoCD가 수행한다.
- **Actions 권한은 두 곳 모두 write여야 한다.** 워크플로의 `contents: write`와 저장소의 `default_workflow_permissions`. 하나라도 `read`면 매니페스트 push가 403으로 실패한다.
- **CI 커밋 메시지에 `[skip` `ci]` 키워드를 언급하지 않는다.** GitHub은 커밋 메시지 본문까지 스캔해 실행을 건너뛴다.
- **롤백은 `git revert`로 한다.** `selfHeal: true`이므로 `kubectl rollout undo`나 ArgoCD UI 롤백은 다음 동기화에서 되돌려진다.
- **`Synced`인데 반영이 안 되면** `status.sync.revision`을 `git ls-remote origin main`과 대조한다. 다르면 repo-server 캐시 문제이므로 `argocd.argoproj.io/refresh=hard` 어노테이션으로 무효화한다.
