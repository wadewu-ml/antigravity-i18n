# antigravity-i18n


[Google Antigravity](https://antigravity.google/) 데스크톱 클라이언트를 위한 다국어 로컬라이제이션 도구입니다.


[简体中文](./README.zh-CN.md) | [English](./README.md) | [繁體中文](./README.zh-Hant.md) | [日本語](./README.ja.md) | **한국어** | [Deutsch](./README.de.md) | [Español](./README.es.md) | [Français](./README.fr.md) | [Português (Brasil)](./README.pt-BR.md) | [Русский](./README.ru.md)

---

현대적인 GUI, 순수 Go로 구현된 asar 엔진, 완전한 언어 팩을 단일 실행 파일에 통합했습니다. Node.js나 런타임 설치 없이 즉시 실행 가능하며, 언제든지 공식 영문 빌드로 바이트 단위 완벽 복원을 지원합니다.

---

## 핵심 기능

- **단일 파일 무설치**: 약 12 MB의 독립 실행형 프로그램으로 Node.js, npm, Python 등의 환경 설정이 전혀 필요 없습니다.
- **GUI 및 CLI 듀얼 모드**: 기본적으로 직관적인 데스크톱 창을 실행하며, 스크립트 및 자동화를 위한 완전한 CLI 하위 명령(`apply` / `restore` / `status` / `locales`)을 내장하고 있습니다.
- **바이트 단위 완벽 복원**: 최초 적용 전 공식 원본 백업(`app.asar.clean-backup`)을 자동 저장합니다. 복원 시 공식 원본 파일과 SHA-256 해시가 완벽히 일치합니다.
- **비침습적 설계**: 쉘 인터페이스, 기본 메뉴 및 종료 확인 창만 번역합니다. 코드 편집기(Monaco), 터미널(xterm), AI 대화 스트림은 원본 그대로 안전하게 유지됩니다.
- **오프라인 및 안정성**: 네트워크 통신 및 원격 측정 데이터 수집이 일절 없습니다. 스크립트 문법 사전 검증 및 트랜잭션 롤백 시스템을 탑재하여 오류 시 안전하게 복구됩니다.

---

## 다운로드 및 사용법

### 다운로드

저장소의 [Releases](../../releases) 페이지에서 최신 `antigravity-i18n.exe`를 다운로드하여 바로 실행하세요.

### 1. 그래픽 인터페이스 (권장)

`antigravity-i18n.exe`를 더블 클릭하여 실행합니다:
1. 프로그램이 Antigravity 설치 경로를 자동으로 감지합니다(미감지 시 '위치 변경'을 클릭하여 직접 지정).
2. 언어 목록에서 원하는 언어를 선택합니다(**English** 선택 시 공식 원본으로 복원).
3. ‘적용’을 누르고 확인합니다. Antigravity가 종료되면 리소스를 갱신하고 앱을 다시 실행합니다.

적용 또는 복원을 위해 Antigravity를 닫아야 하면 선택한 설치의 앱에 정상 종료를 요청합니다. 30초 후에도 실행 중이면 작업을 취소합니다. 즉시 종료하려면 확인 창에서 강제 종료 옵션을 선택하세요.

### 2. 명령줄(CLI) 모드

`antigravity-i18n.exe`가 있는 폴더에서 PowerShell을 엽니다.

```powershell
# Check current localization status and backups
.\antigravity-i18n.exe status

# Install a language pack (Simplified Chinese by default)
.\antigravity-i18n.exe apply --locale zh-CN

# Restore the official untranslated English version
.\antigravity-i18n.exe restore

# List all bundled language packs with coverage
.\antigravity-i18n.exe locales
```

자주 사용하는 옵션:
- `--app-dir <path>`: Antigravity 설치 디렉터리 수동 지정(환경 변수 `ANTIGRAVITY_APP_DIR` 사용 가능).
- `--locale <code>`: 대상 언어 코드(`zh-CN`, `zh-Hant`, `ja`, `ko`, `de`, `es`, `fr` 등).
- `--no-restart`: 적용 완료 후 Antigravity를 자동으로 재실행하지 않음.
- `--force`를 지정하면 선택한 설치의 Antigravity를 즉시 강제 종료합니다. 저장하지 않은 작업이 손실될 수 있습니다.
- `--timeout <ms>`로 작업을 취소하기까지 기다리는 시간을 밀리초 단위로 설정합니다. 기본값은 `30000`입니다.
- `--yes`: 프로세스 종료 확인 메시지 자동 승인.

설치 하나에서는 적용 또는 복원 작업을 한 번에 하나만 실행할 수 있습니다. 다른 작업이 진행 중이면 완료될 때까지 기다린 후 다시 시도하세요. `status`로 선택한 설치와 백업을 확인할 수 있습니다. 아카이브를 읽을 수 없거나 Antigravity 아카이브임을 확인할 수 없으면 오류가 표시됩니다.

---

## 내장 언어 지원

모든 언어 팩은 바이나리에 직접 내장되어 있습니다:

| 코드 | 언어 | 적용 범위 | 설명 |
| :--- | :--- | :--- | :--- |
| `zh-CN` | Simplified Chinese (简体中文) | 100% | 쉘 UI, 기본 메뉴 및 대화 상자 완벽 지원 |
| `zh-Hant` | Traditional Chinese (繁體中文) | 100% | 검증된 기술 용어를 포함한 100% 완전 지원 |
| `ja` | Japanese (日本語) | 100% | 검증된 기술 용어를 포함한 100% 완전 지원 |
| `ko` | Korean (한국어) | 100% | 검증된 기술 용어를 포함한 100% 완전 지원 |
| `de` | German (Deutsch) | 100% | 검증된 기술 용어를 포함한 100% 완전 지원 |
| `es` | Spanish (Español) | 100% | 검증된 기술 용어를 포함한 100% 완전 지원 |
| `fr` | French (Français) | 100% | 검증된 기술 용어를 포함한 100% 완전 지원 |
| `pt-BR` | Portuguese (Brasil) | 100% | 검증된 기술 용어를 포함한 100% 완전 지원 |
| `ru` | Russian (Русский) | 100% | 검증된 기술 용어를 포함한 100% 완전 지원 |
| `en` | English | - | 공식 원본 영문 빌드 (복원하려면 선택) |

---

## 동작 원리 및 자주 묻는 질문

### 1. Antigravity 공식 업데이트 후 작동하지 않나요?
공식 업데이트 시 `app.asar`가 덮어씌워집니다. 업데이트가 완료된 후 본 도구를 다시 열고 '적용'을 누르면 즉시 다시 적용됩니다(도구가 새 버전을 기준으로 깨끗한 백업을 자동 갱신합니다).

### 2. 언어 전환 및 복원
언어 간 전환 시 사전 복원 없이 바로 다른 언어를 선택하여 적용할 수 있습니다. 언제든 '공식 영문으로 복원'을 선택하면 순정 상태로 돌아갑니다.

### 3. 백업 파일 관리 안내
- `app.asar.clean-backup`은 복원에 사용하는 원본 아카이브입니다. 패치되지 않은 설치에 적용하거나 복원할 때 설치된 아카이브와 SHA-256 해시를 비교합니다. 같은 버전을 다시 설치했더라도 해시가 다르면 백업을 갱신합니다. **이 파일을 보관하세요.**
- `app.asar.bak-*`는 타임스탬프가 붙은 복사본입니다. 원본 백업이 없거나 유효하지 않을 때 복구에 사용할 수 있습니다. 앱이 정상 작동하고 `status`에서 백업이 유효하다고 표시되면 오래된 복사본을 삭제해 공간을 확보할 수 있습니다.

---

## 소스 빌드

**사전 요구 사항**:
- Go 1.26.3+
- [Wails v2](https://wails.io/) CLI
- PowerShell 7

```powershell
# Clone the repository
git clone https://github.com/wadewu-ml/antigravity-i18n.git
cd antigravity-i18n

# Build both GUI and CLI binaries
.\build.ps1

# 공식 아카이브 복원 테스트를 포함하려면 다음 줄의 주석을 해제합니다.
# $env:ANTIGRAVITY_ASAR = "$env:LOCALAPPDATA\Programs\antigravity\resources\app.asar.clean-backup"
go test ./...
```

---

## 면책 조항

1. 본 프로젝트는 커뮤니티에서 유지 관리하는 비공식 도구이며 Google 또는 Antigravity 공식과 아무런 관련이 없습니다.
2. Google Antigravity 및 관련 상표의 모든 권리는 원저작권자에게 있습니다.
3. 본 도구는 학습 및 개인적인 용도로 제공됩니다. 관련 소프트웨어 이용 약관을 준수하여 사용하시기 바랍니다.

## 라이선스

[MIT License](./LICENSE)
