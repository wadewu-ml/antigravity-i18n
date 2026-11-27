<div align="center">
  <img src="./frontend/dist/appicon.png" width="68" height="68" alt="Regravity 아이콘">
  <h1>Regravity</h1>
  <p><em>Antigravity가 날아오를 때, Regravity가 당신의 언어와 제어권을 되찾아 줍니다.</em></p>
  <p>휴대용 도구인 Regravity는 Windows용 <strong>Google Antigravity</strong>에 번역 팩 18개와 세션 지표, 모델 한도, 추론 강도 슬라이더를 제공합니다. Regravity 설정 창은 19개 언어를 지원합니다.</p>
  <p>
    <a href="https://github.com/wadewu-ml/regravity/releases/latest/download/regravity.exe"><img src="https://img.shields.io/badge/Download-Latest_(Windows_x64)-0F9D58?style=flat&logo=windows&logoColor=white" alt="최신 Windows x64용 Regravity 다운로드"></a>
    <a href="https://github.com/wadewu-ml/regravity/releases"><img src="https://img.shields.io/badge/Releases-GitHub-18181B?style=flat&logo=github" alt="GitHub 릴리스"></a>
    <img src="https://img.shields.io/badge/Antigravity-2.35.2_tested-18181B?style=flat" alt="Antigravity App 2.35.2 호환 확인">
    <a href="./LICENSE"><img src="https://img.shields.io/badge/License-MIT-18181B?style=flat" alt="MIT 라이선스"></a>
  </p>
  <p>
    <sub>
      <a href="./README.zh-CN.md">简体中文</a> · <a href="./README.md">English</a> · <a href="./README.zh-Hant.md">繁體中文</a> · <a href="./README.ja.md">日本語</a> · <strong>한국어</strong> · <a href="./README.de.md">Deutsch</a> · <a href="./README.es.md">Español</a> · <a href="./README.fr.md">Français</a> · <a href="./README.pt-BR.md">Português (Brasil)</a> · <a href="./README.ru.md">Русский</a> · <a href="./README.tr.md">Türkçe</a> · <a href="./README.vi.md">Tiếng Việt</a> · <a href="./README.it.md">Italiano</a> · <a href="./README.ar.md">العربية</a> · <a href="./README.id.md">Bahasa Indonesia</a> · <a href="./README.pl.md">Polski</a> · <a href="./README.cs.md">Čeština</a> · <a href="./README.hu.md">Magyar</a> · <a href="./README.bg.md">Български</a>
    </sub>
  </p>
</div>

<br>

<p align="center">
  <a href="./docs/images/language-settings.webp">
    <img src="./docs/images/language-settings.webp" width="580" alt="Antigravity v2.35.2, 중국어 간체 선택, 세션 지표와 추론 강도 슬라이더가 모두 켜진 설정 화면">
  </a>
  <br>
  <sub>Antigravity 2.35.2를 사용한 Regravity 설정 화면입니다.</sub>
</p>

## 주요 기능

- **번역 팩 18개, 대화 내용 비간섭** — 설정, 채팅 제어, 모델 메뉴, 기본 메뉴, 종료 안내를 `zh-CN`, `zh-Hant`, `ja`, `ko`, `de`, `es`, `fr`, `pt-BR`, `ru`, `tr`, `vi`, `it`, `ar`, `id`, `pl`, `cs`, `hu`, `bg`로 번역하며 입력이나 대화 내용은 변경하지 않습니다.
- **세션 지표 및 모델 한도** *(선택)* — 모델 선택기 옆 원형 표시기로 남은 한도, 현지 초기화 시각, 컨텍스트 사용량, 평균 출력 속도, 캐시 적중률, 첫 토큰 지연 시간을 확인합니다.
- **추론 강도 슬라이더** *(선택)* — 추론 강도를 선택할 수 있는 모델의 선택 메뉴 안에 캡슐형 슬라이더를 추가합니다.
- **독립 조합 및 바이트 단위 복원** — 두 스위치는 각각 켤 수 있고 기본 영어 UI와도 호환됩니다. 설치된 Antigravity와 같은 버전의 원본 백업으로 언제든 복원할 수 있습니다.

## 빠른 시작

> [!TIP]
> 괄호 안의 중국어 표기는 위 스크린샷의 버튼 이름입니다.

오른쪽 위에서 설정 창의 언어를 선택할 수 있습니다.

### 설정 적용

1. `regravity.exe`를 실행합니다. 설치 위치가 다르게 표시되면 **위치 변경**(`更改位置`)을 눌러 `Antigravity.exe`가 있는 폴더를 선택합니다.
2. **표시 언어**(`界面语言`)에서 언어를 선택합니다(영어 화면에서 부가 기능만 쓰려면 **English** 선택).
3. 언어 목록 아래의 **세션 지표**(`会话指标`)와 **추론 강도 슬라이더**(`思考强度滑块`)를 필요에 따라 켭니다.
4. 작업을 저장하고 Antigravity를 종료한 뒤 오른쪽 아래 **변경 사항 적용**(`应用更改`)을 누릅니다. Antigravity가 새 설정으로 열리면(**Antigravity 열기** / `打开 Antigravity` 지원) 도구는 닫아도 됩니다.

> [!WARNING]
> Antigravity가 실행 중이라는 안내가 뜨면 작업을 저장하고 정상적으로 종료하세요. **강제 종료 후 다시 시도**(`强制退出并重试`)를 사용하면 진행 중인 세션이 중단되고 저장하지 않은 변경 사항이 사라질 수 있습니다.

### 다시 적용 및 복원

- **다시 적용** — 도구를 업데이트한 뒤 **다시 적용**(`重新应用`)을 누르면 새 번역을 사용할 수 있습니다.
- **공식 버전 복원** — Antigravity를 종료한 뒤 오른쪽 위 메뉴(`···`)에서 **공식 버전 복원**(`还原官方版本`)을 누르고 확인하면 기본 영어 화면으로 돌아가며 지표와 슬라이더도 제거됩니다.

> [!IMPORTANT]
> 처음 적용할 때 도구가 Antigravity 설치 폴더에 원본 파일을 `resources/app.asar.clean-backup`으로 자동 백업합니다. 나중에 복원하려면 이 파일을 보관하세요.

## 선택 기능

### 세션 지표 및 모델별 남은 한도

- **첫 메시지 전에도 한도 확인** — 모델 선택기 옆 원형 표시기에 포인터를 올리면 모델 그룹별로 각 기간(주간, 5시간 등)의 남은 한도를 볼 수 있고, 백분율 위에 올리면 현지 초기화 시각이 표시됩니다.
- **모델 호출 후 실시간 지표** — 모델 호출이 기록되면 원형 표시기가 추정 컨텍스트 사용률을 보여주고, 상단에 **컨텍스트 사용량**(사용 토큰 / 전체 용량 및 비율), **평균 출력 속도**(추론 토큰 포함), **캐시 적중률**, **첫 토큰 지연 시간**이 채워집니다. 원형 표시기를 누르면 고정하거나 닫을 수 있습니다.
- **패널별 독립 집계** — 분할 화면, 사이드 채팅, 하위 에이전트 패널은 각각 해당 세션 로컬 기록의 가장 최근 모델 호출 데이터를 표시하며 값이 없으면 `—`로 나타납니다.

<p align="center">
  <a href="./docs/images/session-quota-before-first-call.webp">
    <img src="./docs/images/session-quota-before-first-call.webp" width="420" alt="원형 표시기 위에 열린 세션 지표 및 모델별 남은 한도 패널">
  </a>
  <br>
  <sub>첫 호출 전에도 한도를 바로 확인할 수 있으며, 모델 호출 후 상단 지표가 자동으로 채워집니다.</sub>
</p>

### 추론 강도 슬라이더

**추론 강도 슬라이더**(`思考强度滑块`)를 켜고 적용하면, 추론 강도를 선택할 수 있는 모델의 메뉴 안(사용량 항목 위)에 슬라이더가 표시됩니다. 끌어 단계를 확인하고 놓으면 적용됩니다.

<p align="center">
  <a href="./docs/images/thinking-effort-high.webp">
    <img src="./docs/images/thinking-effort-high.webp" height="340" alt="추론 강도가 높음으로 설정된 모델 선택 메뉴">
  </a>
  &nbsp;&nbsp;
  <a href="./docs/images/thinking-effort-max.webp">
    <img src="./docs/images/thinking-effort-max.webp" height="340" alt="추론 강도가 최고로 설정된 모델 선택 메뉴">
  </a>
  <br>
  <sub>일반 단계(높음)와 강조된 최고 단계(최고)를 모델 메뉴에서 바로 전환할 수 있습니다.</sub>
</p>

## 업데이트 및 CLI

### 업데이트

- **Antigravity 업데이트 후** — 번역이나 기능이 사라지면 페이지 상단의 호환 버전을 확인한 뒤 도구를 다시 열어 설정을 적용하세요.
- **도구 자동 업데이트** — 실행 시 이전 확인 결과를 바탕으로 GitHub 조회 여부를 결정하며 성공 후 최소 24시간 간격을 둡니다. 오른쪽 위 메뉴(`···`)에 파란 점이 보이면 변경 사항을 확인하고 같은 경로에서 바로 업데이트할 수 있습니다. **정보**(`关于`)에서 자동 확인을 끄거나 수동 확인(`检查更新`)할 수 있으며, 업데이트 후 **다시 적용**(`重新应用`)으로 새 번역을 반영하세요.

v4.2.3 이하 버전을 사용 중이라면 [`regravity.exe`](https://github.com/wadewu-ml/regravity/releases/latest/download/regravity.exe)를 직접 다운로드해 실행하세요. 기존 설정은 그대로 유지됩니다.

### PowerShell CLI

`regravity.exe`가 있는 폴더에서 PowerShell을 열어 실행합니다:

| 작업 | 명령 |
| --- | --- |
| 언어 코드 목록 보기 | `.\regravity.exe locales` |
| 한국어 적용 | `.\regravity.exe apply --locale ko` |
| 기본 영어 + 세션 지표 | `.\regravity.exe apply --locale en --session-metrics` |
| 기본 영어 + 추론 강도 슬라이더 | `.\regravity.exe apply --locale en --effort-slider` |
| 현재 설정 확인 | `.\regravity.exe status` |
| 공식 버전 복원 | `.\regravity.exe restore` |
| 모든 옵션 보기 | `.\regravity.exe --help` |

### 소스에서 빌드

Windows, PowerShell 7, Go 1.26.3 이상, Wails v2 CLI가 필요합니다:

```powershell
.\build.ps1 -Target gui
```

`build/bin/regravity.exe`에 생성됩니다.

## 문제 해결

- **실행 시 WebView2가 필요하다는 안내가 표시될 때** — [Microsoft Edge WebView2 런타임](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download)을 설치한 뒤 `regravity.exe`를 다시 실행하세요.

---

## 면책 조항

1. 본 프로젝트는 커뮤니티에서 유지 관리하는 비공식 도구이며 Google 또는 Antigravity 공식과 아무런 관련이 없습니다.
2. Google Antigravity 및 관련 상표의 모든 권리는 원저작권자에게 있습니다.
3. 본 도구는 학습 및 개인적인 용도로 제공됩니다. 관련 소프트웨어 이용 약관을 준수하여 사용하시기 바랍니다.

[MIT 라이선스](./LICENSE)
