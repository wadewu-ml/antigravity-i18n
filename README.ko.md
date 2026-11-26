<div align="center">
  <img src="./frontend/dist/appicon.png" width="64" height="64" alt="Regravity 아이콘">
  <h1>Regravity</h1>
  <p>Windows용 Google Antigravity를 위한 다국어 팩, 세션 지표 및 모델 한도 패널, 추론 강도 슬라이더.</p>
  <p>
    <a href="https://github.com/wadewu-ml/regravity/releases/download/v4.2.3/antigravity-i18n.exe"><strong>Windows x64 다운로드</strong></a>
    ·
    <a href="https://github.com/wadewu-ml/regravity/releases">릴리스</a>
    ·
    호환 확인: <strong>Antigravity App 2.35.1</strong>
  </p>
  <p>
    <a href="./README.zh-CN.md">简体中文</a> · <a href="./README.md">English</a> · <a href="./README.zh-Hant.md">繁體中文</a> · <a href="./README.ja.md">日本語</a> · <strong>한국어</strong> · <a href="./README.de.md">Deutsch</a> · <a href="./README.es.md">Español</a> · <a href="./README.fr.md">Français</a> · <a href="./README.pt-BR.md">Português (Brasil)</a> · <a href="./README.ru.md">Русский</a>
  </p>
  <p>antigravity-i18n의 새 이름은 Regravity입니다. 위 다운로드는 v4.2.3입니다. Regravity 첫 버전이 출시되면 <code>regravity.exe</code>를 직접 다운로드하세요. 기존 패치와 자동 업데이트 설정은 유지됩니다. 저장소 이름 변경으로 v4.2.3에서 버전 비교 오류가 발생할 수 있습니다. 위 릴리스 링크에서 확인하세요.</p>
</div>

<p align="center">
  <a href="./docs/images/language-settings.webp">
    <img src="./docs/images/language-settings.webp" width="560" alt="Antigravity v2.35.2, 중국어 간체 선택, 세션 지표와 추론 강도 슬라이더가 모두 켜진 설정 화면">
  </a>
</p>

## 주요 기능

- **단일 휴대용 EXE에 9개 언어 포함** — 설정, 채팅 버튼, 모델 선택 메뉴, 기본 메뉴, 종료 안내를 중국어 간체, 중국어 번체, 일본어, 한국어, 독일어, 스페인어, 프랑스어, 브라질 포르투갈어, 러시아어로 번역하며 사용자 입력이나 대화 내용은 변경하지 않습니다.
- **세션 지표 및 모델 한도 패널(선택)** — 모델 선택기 옆에 원형 표시기를 추가해 남은 한도, 현지 초기화 시각, 컨텍스트 사용량, 평균 출력 속도, 캐시 적중률, 첫 토큰 지연 시간을 확인할 수 있습니다.
- **추론 강도 슬라이더(선택)** — 추론 강도를 선택할 수 있는 모델의 선택 메뉴 안에 캡슐형 슬라이더를 추가합니다.
- **독립 조합 및 바이트 단위 복원** — 두 스위치는 각각 켤 수 있으며 기본 영어 UI에서도 사용할 수 있습니다. 적용 후에는 도구를 닫아도 되며 설치된 Antigravity와 같은 버전의 원본 백업으로 복원할 수 있습니다.

## 빠른 시작

> [!NOTE]
> 설정 도구 화면은 중국어 간체로 표시됩니다. 위 스크린샷과 대조할 수 있도록 아래 단계에 원래 중국어 간체 라벨을 함께 표기했습니다.

### 설정 적용

1. `antigravity-i18n.exe`를 실행합니다. 설치 위치가 다르게 표시되면 **위치 변경**(`更改位置`)을 눌러 `Antigravity.exe`가 있는 폴더를 선택합니다.
2. **표시 언어**(`界面语言`)에서 언어를 선택합니다(영어 화면에서 부가 기능만 쓰려면 **English** 선택).
3. 언어 목록 아래의 **세션 지표**(`会话指标`)와 **추론 강도 슬라이더**(`思考强度滑块`)를 필요에 따라 켭니다.
4. 작업을 저장하고 Antigravity를 종료한 뒤 오른쪽 아래 **변경 사항 적용**(`应用更改`)을 누릅니다. Antigravity가 새 설정으로 다시 열리며(**Antigravity 열기** / `打开 Antigravity` 버튼도 사용 가능) 이후에는 이 도구를 닫아도 됩니다.

> [!WARNING]
> Antigravity가 실행 중이라는 안내가 뜨면 작업을 저장하고 정상적으로 종료하세요. **강제 종료 후 다시 시도**(`强制退出并重试`)를 사용하면 진행 중인 세션이 중단되고 저장하지 않은 변경 사항이 사라질 수 있습니다.

### 다시 적용 및 복원

- **현재 설정 다시 적용** — 오른쪽 아래 **다시 적용**(`重新应用`)을 누릅니다.
- **공식 버전 복원** — Antigravity를 종료한 뒤 오른쪽 위 메뉴(`···`)에서 **공식 버전 복원**(`还原官方版本`)을 누르고 확인합니다. 원래 영어 화면으로 돌아가며 지표 원형 표시기와 슬라이더도 제거됩니다.

> [!IMPORTANT]
> 처음 적용할 때 도구가 Antigravity 설치 폴더에 원본 파일을 `resources/app.asar.clean-backup`으로 자동 백업합니다. 나중에 복원하려면 이 파일을 보관하세요.

## 선택 기능 안내

### 세션 지표 및 모델별 남은 한도

- **첫 메시지 전에도 한도 확인** — 모델 선택기 옆의 원형 표시기에 포인터를 올리면 모델 그룹별로 각 기간(주간, 5시간 등)의 남은 한도를 볼 수 있고, 백분율 위에 올리면 현지 시간 기준 다음 초기화 시각이 표시됩니다.
- **모델 호출 후 실시간 지표** — 모델 호출이 기록되면 원형 표시기가 추정 컨텍스트 사용률을 보여주고, 패널 상단에 **컨텍스트 사용량**(사용 토큰 / 전체 용량 및 비율), **평균 출력 속도**(추론 토큰 포함), **캐시 적중률**, **첫 토큰 지연 시간**이 채워집니다. 원형 표시기를 누르면 패널이 고정되거나 닫힙니다.
- **패널별 독립 집계** — 분할 화면, 사이드 채팅, 하위 에이전트 패널은 각각 해당 세션의 로컬 기록에서 가장 최근 모델 호출 데이터를 표시합니다. 값이 없으면 `—`로 표시됩니다.

<p align="center">
  <a href="./docs/images/session-quota-before-first-call.webp">
    <img src="./docs/images/session-quota-before-first-call.webp" width="440" alt="원형 표시기 위에 열린 세션 지표 및 모델별 남은 한도 패널">
  </a>
</p>

### 추론 강도 슬라이더

설정 도구에서 **추론 강도 슬라이더**(`思考强度滑块`)를 켜고 변경 사항을 적용하세요. Antigravity 모델 선택 메뉴에서 추론 강도를 선택할 수 있는 모델은 사용량 항목 위에 슬라이더가 표시됩니다. 슬라이더를 끌어 단계를 확인하고 놓으면 적용됩니다.

<p align="center">
  <a href="./docs/images/thinking-effort-high.webp">
    <img src="./docs/images/thinking-effort-high.webp" height="340" alt="추론 강도가 높음으로 설정된 모델 선택 메뉴">
  </a>
  &nbsp;&nbsp;
  <a href="./docs/images/thinking-effort-max.webp">
    <img src="./docs/images/thinking-effort-max.webp" height="340" alt="추론 강도가 최고로 설정된 모델 선택 메뉴">
  </a>
</p>

## 업데이트

- **Antigravity 업데이트 후** — 번역이나 기능이 사라지면 이 페이지 상단의 호환 버전을 확인한 뒤 도구를 다시 열어 설정을 적용하세요.
- **도구 자동 업데이트** — 실행 시 이전 확인 결과에 따라 GitHub에 새 버전을 요청합니다. 확인에 성공하면 다음 자동 확인까지 최소 24시간 기다립니다. 오른쪽 위 메뉴(`···`)에 파란 점이 보이면 메뉴에서 변경 사항을 확인하고 승인해 같은 경로에서 새 버전으로 다시 시작할 수 있습니다. **정보**(`关于`)에서 자동 확인을 끄거나 메뉴에서 수동 확인(`检查更新`)할 수 있습니다. 도구 업데이트 후 **다시 적용**(`重新应用`)을 눌러 새 번역을 반영하세요.

## PowerShell CLI

`antigravity-i18n.exe`가 있는 폴더에서 PowerShell을 열어 실행합니다:

| 작업 | 명령 |
| --- | --- |
| 언어 코드 목록 보기 | `.\antigravity-i18n.exe locales` |
| 한국어 적용 | `.\antigravity-i18n.exe apply --locale ko` |
| 기본 영어 + 세션 지표 | `.\antigravity-i18n.exe apply --locale en --session-metrics` |
| 기본 영어 + 추론 강도 슬라이더 | `.\antigravity-i18n.exe apply --locale en --effort-slider` |
| 현재 설정 확인 | `.\antigravity-i18n.exe status` |
| 공식 버전 복원 | `.\antigravity-i18n.exe restore` |
| 모든 옵션 보기 | `.\antigravity-i18n.exe --help` |

## 소스에서 빌드

Windows, PowerShell 7, Go 1.26.3 이상, Wails v2 CLI가 필요합니다:

```powershell
.\build.ps1 -Target gui
```

빌드된 실행 파일은 `build/bin/regravity.exe`에 생성됩니다.

## 문제 해결

- **실행 시 WebView2가 필요하다는 안내가 표시될 때** — [Microsoft Edge WebView2 런타임](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download)을 설치한 뒤 `antigravity-i18n.exe`를 다시 실행하세요.

## 면책 조항

1. 본 프로젝트는 커뮤니티에서 유지 관리하는 비공식 도구이며 Google 또는 Antigravity 공식과 아무런 관련이 없습니다.
2. Google Antigravity 및 관련 상표의 모든 권리는 원저작권자에게 있습니다.
3. 본 도구는 학습 및 개인적인 용도로 제공됩니다. 관련 소프트웨어 이용 약관을 준수하여 사용하시기 바랍니다.

[MIT 라이선스](./LICENSE)
