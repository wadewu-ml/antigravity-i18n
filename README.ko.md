# <img src="./frontend/dist/appicon.png" width="32" height="32" alt=""> antigravity-i18n

Windows용 Google Antigravity의 설정, 채팅 버튼, 메뉴, 종료 안내를 한국어를 비롯한 9개 언어로 번역할 수 있습니다. 채팅 입력창에는 남은 컨텍스트와 모델의 출력 속도를 표시할 수도 있습니다. 설정을 적용한 뒤에는 이 도구를 닫아도 됩니다.

[Windows x64 다운로드](https://github.com/wadewu-ml/antigravity-i18n/releases/latest/download/antigravity-i18n.exe) · [릴리스](https://github.com/wadewu-ml/antigravity-i18n/releases)

**Antigravity App 2.35.1**에서 호환성을 확인했습니다.

[简体中文](./README.zh-CN.md) · [English](./README.md) · [繁體中文](./README.zh-Hant.md) · [日本語](./README.ja.md) · **한국어** · [Deutsch](./README.de.md) · [Español](./README.es.md) · [Français](./README.fr.md) · [Português (Brasil)](./README.pt-BR.md) · [Русский](./README.ru.md)

<a href="./docs/images/antigravity-zh-cn.png"><img src="./docs/images/translated-window.webp" width="680" alt="중국어 간체 메뉴와 입력창의 모델 선택기 옆에 있는 세션 지표 표시기"></a>

## 표시 언어 바꾸기

언어 팩은 중국어 간체, 중국어 번체, 일본어, 한국어, 독일어, 스페인어, 프랑스어, 브라질 포르투갈어, 러시아어를 지원합니다. 설정 도구의 화면은 중국어 간체로 표시됩니다.

1. `antigravity-i18n.exe`를 엽니다. 설치 위치가 제대로 표시되지 않으면 ‘위치 변경’를 눌러 `Antigravity.exe`가 있는 폴더를 선택합니다.
2. ‘표시 언어’에서 표시 언어를 선택합니다. 세션 지표를 보려면 ‘세션 지표’를 켭니다. 영어 화면에서 지표만 사용하려면 English를 선택합니다.
3. 작업을 저장하고 Antigravity를 종료한 다음 ‘변경 사항 적용’를 누릅니다. Antigravity가 새 설정으로 다시 열립니다.

종료 안내에서 ‘강제 종료 후 다시 시도’를 선택하면 Antigravity를 강제로 종료하고 다시 시도합니다. 진행 중인 세션이 중단되며 저장하지 않은 변경 사항이 사라질 수 있습니다.

<a href="./docs/images/language-settings.webp"><img src="./docs/images/language-settings.webp" width="640" alt="중국어 간체가 선택되고 세션 지표가 켜진 설정 화면"></a>

화면 오른쪽 아래의 초록색 ‘重新应用’ 버튼을 누르면 설정이 다시 적용됩니다.

## 세션 지표 보기

채팅에 모델 호출 기록이 생기면 입력창의 모델 선택기 오른쪽에 원형 표시기가 나타납니다. 원이 채워진 정도는 컨텍스트를 얼마나 사용했는지 보여 주는 추정치입니다. 포인터를 올리면 남은 토큰 수, 출력 속도, 캐시 적중률, 첫 토큰까지 걸린 시간을 볼 수 있습니다. 한 번 누르면 세부 정보가 고정되고 다시 누르면 닫힙니다.

<a href="./docs/images/session-metrics-zh-cn.png"><img src="./docs/images/session-metrics-detail.webp" width="560" alt="원형 표시기 위에 열린 중국어 간체 패널에 컨텍스트 사용량, 남은 토큰 수, 출력 속도가 표시된 모습"></a>

분할 화면, 사이드 채팅, 하위 에이전트 패널은 각각 해당 세션에 저장된 가장 최근 모델 호출 기록에서 지표를 읽습니다. 출력 토큰 수에는 모델이 기록한 추론 토큰이 포함됩니다. 값이 없는 항목은 `—`로 표시됩니다.

## 복원 및 다시 적용

원래 영어 화면으로 돌아가고 지표를 없애려면 Antigravity를 종료하고 도구에서 ‘공식 버전 복원’을 누른 뒤 ‘복원’으로 확인합니다. 처음 설정을 적용할 때 도구가 원본 파일을 자동으로 백업합니다. 복원에 쓰는 `resources/app.asar.clean-backup`은 Antigravity 설치 폴더에 보관하세요.

Antigravity 업데이트 후 번역이나 지표가 사라지면 이 페이지 상단에서 지원 버전을 확인한 뒤 언어와 지표를 다시 선택해 적용합니다. 이미 적용된 설정을 다시 쓰려면 ‘다시 적용’을 누릅니다.

도구를 업데이트하려면 창 상단의 버전 번호 옆에 있는 업데이트 확인 버튼을 누르세요. 버튼을 누르면 GitHub에서 새 버전을 확인합니다. 다운로드를 승인하면 기존 EXE와 같은 경로에서 새 버전이 열립니다. 새 번역을 Antigravity에 반영하려면 설정을 다시 적용하세요.

## PowerShell 명령

`antigravity-i18n.exe`가 있는 폴더에서 PowerShell을 엽니다.

| 작업 | 명령 |
| --- | --- |
| 언어 코드 목록 보기 | `.\antigravity-i18n.exe locales` |
| 한국어 적용 | `.\antigravity-i18n.exe apply --locale ko` |
| 공식 영어 및 지표 사용 | `.\antigravity-i18n.exe apply --locale en --session-metrics` |
| 현재 설정 확인 | `.\antigravity-i18n.exe status` |
| 공식 버전 복원 | `.\antigravity-i18n.exe restore` |
| 모든 옵션 보기 | `.\antigravity-i18n.exe --help` |

## 소스에서 빌드

Windows에서 빌드하려면 PowerShell 7, Go 1.26.3 이상, Wails v2 CLI가 필요합니다. 저장소 루트에서 `.\build.ps1 -Target gui`를 실행하면 `build/bin/antigravity-i18n.exe`가 생성됩니다.

## 시작 문제 해결

도구를 실행할 때 WebView2가 필요하다는 안내가 뜨면 [Microsoft Edge WebView2 런타임](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download)을 설치하세요.

## 면책 조항

1. 본 프로젝트는 커뮤니티에서 유지 관리하는 비공식 도구이며 Google 또는 Antigravity 공식과 아무런 관련이 없습니다.
2. Google Antigravity 및 관련 상표의 모든 권리는 원저작권자에게 있습니다.
3. 본 도구는 학습 및 개인적인 용도로 제공됩니다. 관련 소프트웨어 이용 약관을 준수하여 사용하시기 바랍니다.

[MIT 라이선스](./LICENSE)
