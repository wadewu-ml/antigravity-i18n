<div align="center">
  <img src="./frontend/dist/appicon.png" width="68" height="68" alt="Икона на Regravity">
  <h1>Regravity</h1>
  <p>Преносим инструмент за Windows, който превежда Google Antigravity и добавя показатели за сесията и плъзгач за нивото на разсъждение. Regravity предлага 18 пакета за превод и 19 езика за настройките, включително български.</p>
  <p>
    <a href="https://github.com/wadewu-ml/regravity/releases/latest/download/regravity.exe"><img src="https://img.shields.io/badge/Download-Latest_(Windows_x64)-0F9D58?style=flat&logo=windows&logoColor=white" alt="Изтегляне на стабилната версия за Windows x64"></a>
    <a href="https://github.com/wadewu-ml/regravity/releases"><img src="https://img.shields.io/badge/Releases-GitHub-18181B?style=flat&logo=github" alt="Издания в GitHub"></a>
    <img src="https://img.shields.io/badge/Antigravity-2.36.0_tested-18181B?style=flat" alt="Тествано с Antigravity 2.36.0">
    <a href="./LICENSE"><img src="https://img.shields.io/badge/License-MIT-18181B?style=flat" alt="Лиценз MIT"></a>
  </p>
  <p>
    <sub>
      <a href="./README.zh-CN.md">简体中文</a> · <a href="./README.md">English</a> · <a href="./README.zh-Hant.md">繁體中文</a> · <a href="./README.ja.md">日本語</a> · <a href="./README.ko.md">한국어</a> · <a href="./README.de.md">Deutsch</a> · <a href="./README.es.md">Español</a> · <a href="./README.fr.md">Français</a> · <a href="./README.pt-BR.md">Português (Brasil)</a> · <a href="./README.ru.md">Русский</a> · <a href="./README.tr.md">Türkçe</a> · <a href="./README.vi.md">Tiếng Việt</a> · <a href="./README.it.md">Italiano</a> · <a href="./README.ar.md">العربية</a> · <a href="./README.id.md">Bahasa Indonesia</a> · <a href="./README.pl.md">Polski</a> · <a href="./README.cs.md">Čeština</a> · <a href="./README.hu.md">Magyar</a> · <strong>Български</strong>
    </sub>
  </p>
</div>

## Първи стъпки

1. Стартирайте `regravity.exe`. Изберете **Български** от горния списък, за да смените езика на настройките на Regravity.
2. В раздела **Език на Antigravity** изберете **Български**. По желание включете **Показатели за сесията** и **Плъзгач за нивото на разсъждение**. Ако пътят е грешен, натиснете **Промяна на местоположението** и посочете папката с `Antigravity.exe`.
3. Запазете работата си, затворете Antigravity и натиснете **Прилагане на промените**. След като Antigravity се отвори отново, можете да затворите Regravity.

> [!WARNING]
> Принудителното затваряне прекъсва активните сесии и може да доведе до загуба на незапазени промени.

<p align="center"><a href="./docs/images/language-settings.webp"><img src="./docs/images/language-settings.webp" width="580" alt="Прозорец за настройки"></a></p>

## Показатели за сесията и ниво на разсъждение

Посочете пръстена до менюто за избор на модел, за да видите оставащата квота и местното време на подновяване. След първия отговор панелът показва и използвания контекст, средната скорост на генериране, дела на попаденията в кеша и времето до първия токен. Щракнете върху пръстена, за да закрепите панела. Всяка сесия има собствени показатели.

<p align="center"><a href="./docs/images/session-quota-before-first-call.webp"><img src="./docs/images/session-quota-before-first-call.webp" width="420" alt="Оставаща квота на модела"></a></p>

Ако моделът поддържа няколко нива на разсъждение, в менюто му се появява плъзгач. Преместете го до желаното ниво и го отпуснете. Панелът и плъзгачът следват езика, приложен в Antigravity.

<p align="center"><a href="./docs/images/thinking-effort-high.webp"><img src="./docs/images/thinking-effort-high.webp" width="310" alt="Високо ниво на разсъждение"></a> <a href="./docs/images/thinking-effort-max.webp"><img src="./docs/images/thinking-effort-max.webp" width="310" alt="Максимално ниво"></a></p>

## Актуализации и възстановяване

- След актуализация на Regravity натиснете **Повторно прилагане**. Ако актуализация на Antigravity премахне превода, проверете съвместимостта и приложете настройките отново.
- Синята точка до менюто **···** показва, че има актуализация. От менюто можете да проверите ръчно, а в **За програмата** можете да изключите автоматичната проверка. Ако използвате версия v4.2.3 или по-стара, изтеглете новия `regravity.exe` ръчно.
- За да възстановите английския интерфейс и да премахнете двете допълнения, затворете Antigravity и изберете **··· → Възстановяване на оригинала**.

> Запазете файла `resources/app.asar.clean-backup`, създаден при първото прилагане. Regravity го използва за възстановяване на оригиналното приложение.

## PowerShell

Изпълнете командите в папката с `regravity.exe`.

```powershell
.\regravity.exe locales
.\regravity.exe apply --locale bg
.\regravity.exe apply --locale en --session-metrics
.\regravity.exe apply --locale en --effort-slider
.\regravity.exe status
.\regravity.exe restore
.\regravity.exe --help
```

За компилиране под Windows са нужни PowerShell 7, Go 1.26.3+ и Wails v2 CLI. Изпълнете `.\build.ps1 -Target gui` в папката на проекта. Резултатът е `build/bin/regravity.exe`. Ако липсва WebView2, инсталирайте [Microsoft Edge WebView2 Runtime](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download) и отворете отново Regravity.

## Лиценз

Regravity е независим проект на общността. Няма официална връзка с Google или Antigravity и не е одобрен от тях. Google Antigravity и свързаните търговски марки принадлежат на съответните им собственици. Инструментът е предназначен за лична и образователна употреба при спазване на условията за ползване на съответния софтуер.

[Лиценз MIT](./LICENSE)
