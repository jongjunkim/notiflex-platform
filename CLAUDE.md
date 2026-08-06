# Notiflex Platform

## 프로젝트 개요

Notiflex는 B2B 알림 SaaS 플랫폼이다. 이 저장소는 애플리케이션 코드와 Kubernetes 매니페스트, CI 파이프라인을 함께 관리하는 GitOps 저장소이다.

## 기술 스택

| 구분 | 내용 |
|------|------|
| 언어 | Go 1.25 (표준 라이브러리만 사용, 외부 웹 프레임워크 없음) |
| 컨테이너 | 멀티스테이지 빌드 → `scratch` 베이스 이미지 |
| 오케스트레이션 | GKE Standard (Zonal), Spot VM |
| GitOps | ArgoCD |

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
└── .github/
    └── workflows/     # CI 파이프라인
```

## 행동 규칙

1. **항상 확인 후 실행한다.** 리소스를 생성·변경·삭제하기 전에 무엇을 어떻게 바꿀지 먼저 설명한다.
2. **변경 전 현재 상태를 확인한다.** 매니페스트를 수정하기 전에 클러스터의 실제 상태(`kubectl get`)와 Git의 상태를 먼저 읽는다.
3. **kubectl에는 컨텍스트를 명시한다.** 잘못된 클러스터를 대상으로 실행하지 않도록 모든 kubectl 명령에 `--context gke-sysnet4admin_book_gitaiops`를 지정한다.
4. **클러스터에 직접 수정하지 않는다.** 배포 변경은 매니페스트를 고치고 Git에 커밋하여 ArgoCD가 반영하도록 한다.
5. **버전은 고정한다.** 이미지 태그에 `latest`를 쓰지 않고 커밋 SHA 또는 명시적 버전을 사용한다.
