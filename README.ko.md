# <img src="./frontend/dist/appicon.png" width="32" height="32" alt=""> antigravity-i18n

Windows용 Google Antigravity의 설정, 채팅 버튼, 메뉴, 종료 안내를 한국어를 비롯한 9개 언어로 번역할 수 있습니다. 채팅 입력창에는 남은 컨텍스트와 모델의 출력 속도를 표시할 수도 있습니다. 설정을 적용한 뒤에는 이 도구를 닫아도 됩니다.

[Windows x64 다운로드](https://github.com/wadewu-ml/antigravity-i18n/releases/latest/download/antigravity-i18n.exe) · [릴리스](https://github.com/wadewu-ml/antigravity-i18n/releases)

**Antigravity App 2.35.1**에서 호환성을 확인했습니다. 휴대용 EXE에 언어 팩 9개가 모두 포함되어 있습니다.

[简体中文](./README.zh-CN.md) · [English](./README.md) · [繁體中文](./README.zh-Hant.md) · [日本語](./README.ja.md) · **한국어** · [Deutsch](./README.de.md) · [Español](./README.es.md) · [Français](./README.fr.md) · [Português (Brasil)](./README.pt-BR.md) · [Русский](./README.ru.md)

<a href="./docs/images/antigravity-zh-cn.png"><img src="./docs/images/translated-window.webp" width="680" alt="중국어 간체 메뉴와 입력창의 모델 선택기 옆에 있는 세션 지표 표시기"></a>

## 표시 언어 바꾸기

언어 팩은 중국어 간체, 중국어 번체, 일본어, 한국어, 독일어, 스페인어, 프랑스어, 브라질 포르투갈어, 러시아어를 지원합니다. 설정 도구의 화면은 중국어 간체로 표시됩니다.

1. `antigravity-i18n.exe`를 엽니다. 설치 위치가 제대로 표시되지 않으면 ‘위치 변경’를 눌러 `Antigravity.exe`가 있는 폴더를 선택합니다.
2. ‘표시 언어’에서 표시 언어를 선택합니다. 세션 지표를 보려면 ‘세션 지표’를 켭니다. 영어 화면에서 지표만 사용하려면 English를 선택합니다.
3. 작업을 저장하고 Antigravity를 종료한 다음 ‘변경 사항 적용’를 누릅니다. Antigravity가 새 설정으로 다시 열립니다.

종료 안내에서 ‘강제 종료 후 다시 시도’를 선택하면 Antigravity를 강제로 종료하고 다시 시도합니다. 진행 중인 세션이 중단되며 저장하지 않은 변경 사항이 사라질 수 있습니다.

슬라이더를 사용하려면 설정 도구에서 세션 지표 아래의 스위치를 켜고 변경 사항을 적용한 뒤 모델 선택 메뉴를 여세요. 두 기능은 각각 켤 수 있으며 기본 영어 인터페이스에서도 사용할 수 있습니다. 추론 강도를 선택할 수 있는 모델에는 사용량 항목 위에 슬라이더가 표시됩니다. 슬라이더를 끌어 단계를 확인하고 놓으면 적용됩니다.

<a href="./docs/images/language-settings.webp"><img src="./docs/images/language-settings.webp" width="640" alt="중국어 간체가 선택되고 세션 지표가 켜진 설정 화면"></a>

오른쪽 아래 ‘重新应用’ 버튼을 누르면 현재 설정이 다시 적용됩니다.

<p>
  <a href="./docs/images/thinking-effort-high.webp"><img src="./docs/images/thinking-effort-high.webp" width="282" alt="추론 강도를 높음으로 설정한 모델 선택 메뉴"></a>
  <a href="./docs/images/thinking-effort-max.webp"><img src="./docs/images/thinking-effort-max.webp" width="267" alt="추론 강도를 최고로 설정한 모델 선택 메뉴"></a>
</p>

## 세션 지표 보기

모델 선택기 옆의 원형 표시기에 포인터를 올리면 모델 그룹별로 각 기간의 남은 한도를 확인할 수 있습니다. 백분율에 포인터를 올리면 다음 초기화 시각이 현지 시간으로 표시됩니다. 첫 메시지를 보내기 전에도 확인할 수 있습니다. 모델 호출이 기록되면 원이 채워진 정도로 추정 컨텍스트 사용률을 알 수 있고, 패널에는 사용한 토큰 수와 전체 용량, 출력 속도, 캐시 적중률, 첫 토큰까지 걸린 시간이 표시됩니다. 한 번 누르면 패널이 고정되고 다시 누르면 닫힙니다.

<a href="./docs/images/session-metrics-zh-cn.png"><img src="./docs/images/session-metrics-detail.webp" width="560" alt="원형 표시기 위에 열린 중국어 간체 패널에 컨텍스트 사용량, 남은 토큰 수, 출력 속도가 표시된 모습"></a>

[첫 모델 호출 전의 한도 패널](./docs/images/session-quota-before-first-call.webp)

분할 화면, 사이드 채팅, 하위 에이전트 패널은 각각 해당 세션에 저장된 가장 최근 모델 호출 기록에서 지표를 읽습니다. 출력 토큰 수에는 모델이 기록한 추론 토큰이 포함됩니다. 값이 없는 항목은 `—`로 표시됩니다.

## 복원 및 다시 적용

원래 영어 화면으로 돌아가고 지표와 슬라이더를 없애려면 Antigravity를 종료하고 창 오른쪽 위 메뉴에서 ‘공식 버전 복원’(还原官方版本)을 누른 뒤 ‘복원’으로 확인합니다. 처음 설정을 적용할 때 도구가 원본 파일을 자동으로 백업합니다. 복원에 쓰는 `resources/app.asar.clean-backup`은 Antigravity 설치 폴더에 보관하세요.

Antigravity 업데이트 후 번역이나 지표가 사라지면 이 페이지 상단에서 지원 버전을 확인한 뒤 언어와 지표를 다시 선택해 적용합니다. 이미 적용된 설정을 다시 쓰려면 ‘다시 적용’을 누릅니다.

마지막으로 업데이트를 확인한 뒤 24시간이 지났으면 도구를 열 때 다시 확인합니다. 오른쪽 위 메뉴에 파란 점이 보이면 메뉴에서 새 버전의 변경 사항을 확인하세요. 정보(关于)에서 자동 확인을 꺼도 메뉴에서 수동으로 확인할 수 있습니다(检查更新). 다운로드를 승인하면 기존 EXE와 같은 경로에서 새 버전이 열립니다. 새 번역을 Antigravity에 반영하려면 설정을 다시 적용하세요.

## PowerShell 명령

`antigravity-i18n.exe`가 있는 폴더에서 PowerShell을 엽니다.

| 작업 | 명령 |
| --- | --- |
| 언어 코드 목록 보기 | `.\antigravity-i18n.exe locales` |
| 한국어 적용 | `.\antigravity-i18n.exe apply --locale ko` |
| 공식 영어 및 지표 사용 | `.\antigravity-i18n.exe apply --locale en --session-metrics` |
| 기본 영어에서 추론 강도 슬라이더 사용 | `.\antigravity-i18n.exe apply --locale en --effort-slider` |
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
