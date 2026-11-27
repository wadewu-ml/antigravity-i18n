<div align="center">
  <img src="./frontend/dist/appicon.png" width="68" height="68" alt="A Regravity ikonja">
  <h1>Regravity</h1>
  <p>Hordozható Windows-eszköz a Google Antigravity felületének lefordításához. A fejlesztői változat 18 fordítási csomagot kínál, a Regravity beállítóablakában pedig 19 nyelv érhető el, köztük a magyar.</p>
  <p>
    <a href="https://github.com/wadewu-ml/regravity/releases/latest/download/regravity.exe"><img src="https://img.shields.io/badge/Download-Latest_(Windows_x64)-0F9D58?style=flat&logo=windows&logoColor=white" alt="A stabil kiadás letöltése Windows x64 rendszerre"></a>
    <a href="https://github.com/wadewu-ml/regravity/releases"><img src="https://img.shields.io/badge/Releases-GitHub-18181B?style=flat&logo=github" alt="Kiadások a GitHubon"></a>
    <img src="https://img.shields.io/badge/Antigravity-2.35.2_tested-18181B?style=flat" alt="Antigravity 2.35.2 verzióval tesztelve">
    <a href="./LICENSE"><img src="https://img.shields.io/badge/License-MIT-18181B?style=flat" alt="MIT-licenc"></a>
  </p>
  <p>
    <sub>
      <a href="./README.zh-CN.md">简体中文</a> · <a href="./README.md">English</a> · <a href="./README.zh-Hant.md">繁體中文</a> · <a href="./README.ja.md">日本語</a> · <a href="./README.ko.md">한국어</a> · <a href="./README.de.md">Deutsch</a> · <a href="./README.es.md">Español</a> · <a href="./README.fr.md">Français</a> · <a href="./README.pt-BR.md">Português (Brasil)</a> · <a href="./README.ru.md">Русский</a> · <a href="./README.tr.md">Türkçe</a> · <a href="./README.vi.md">Tiếng Việt</a> · <a href="./README.it.md">Italiano</a> · <a href="./README.ar.md">العربية</a> · <a href="./README.id.md">Bahasa Indonesia</a> · <a href="./README.pl.md">Polski</a> · <a href="./README.cs.md">Čeština</a> · <strong>Magyar</strong> · <a href="./README.bg.md">Български</a>
    </sub>
  </p>
</div>

## Első lépések

1. Indítsa el a fejlesztői változat `regravity.exe` fájlját. A felső listában válassza a **Magyar** lehetőséget a beállítóablak nyelvének módosításához.
2. Az Antigravity nyelvének beállításánál válassza a **Magyar** lehetőséget. Igény szerint kapcsolja be a **Munkamenet-mutatók** és a **Gondolkodási szint csúszkája** funkciót. Ha a program helye hibás, a **Hely módosítása** gombbal válassza ki az `Antigravity.exe` fájlt tartalmazó mappát.
3. Mentse a munkáját, zárja be az Antigravityt, majd kattintson a **Módosítások alkalmazása** gombra. Az Antigravity újraindulása után bezárhatja a Regravityt.

> [!WARNING]
> A kényszerített kilépés megszakítja a futó munkameneteket, és a nem mentett módosítások elveszhetnek.

<p align="center"><a href="./docs/images/language-settings.webp"><img src="./docs/images/language-settings.webp" width="580" alt="Beállítóablak"></a></p>

## Munkamenet-mutatók és gondolkodási szint

Vigye a mutatót a modellválasztó melletti gyűrűre. A panel megjeleníti a fennmaradó modellkeretet és a megújulás helyi idejét. Az első válasz után a felhasznált kontextust, az átlagos kimeneti sebességet, a gyorsítótár-találati arányt és az első tokenig eltelt időt is láthatja. A gyűrűre kattintva rögzítheti a panelt. A mutatók beszélgetésenként külön jelennek meg.

<p align="center"><a href="./docs/images/session-quota-before-first-call.webp"><img src="./docs/images/session-quota-before-first-call.webp" width="420" alt="Fennmaradó modellkeret"></a></p>

Ha a modell több gondolkodási szintet támogat, a modell menüjében megjelenik egy csúszka. Húzza a kívánt szintre, majd engedje el. A panel és a csúszka az Antigravityben alkalmazott nyelvet követi.

<p align="center"><a href="./docs/images/thinking-effort-high.webp"><img src="./docs/images/thinking-effort-high.webp" width="310" alt="Magas gondolkodási szint"></a> <a href="./docs/images/thinking-effort-max.webp"><img src="./docs/images/thinking-effort-max.webp" width="310" alt="Maximális szint"></a></p>

## Frissítés és visszaállítás

- A Regravity frissítése után kattintson az **Ismételt alkalmazás** gombra. Ha az Antigravity frissítése felülírja a fordítást, ellenőrizze a kompatibilitást, majd alkalmazza újra a beállításokat.
- A **···** menü melletti kék pont elérhető frissítést jelez. Itt kézzel is kereshet frissítést, a **Névjegy** ablakban pedig kikapcsolhatja az automatikus ellenőrzést. A v4.2.3-as és korábbi verziókról kézzel töltse le az új `regravity.exe` fájlt.
- Az angol felület visszaállításához és a két kiegészítő kikapcsolásához zárja be az Antigravityt, majd válassza a **··· → Eredeti verzió visszaállítása** lehetőséget.

> Őrizze meg az első alkalmazáskor létrejövő `resources/app.asar.clean-backup` fájlt. A Regravity ebből állítja vissza az eredeti alkalmazást.

## PowerShell

A parancsokat a `regravity.exe` fájlt tartalmazó mappában futtassa.

```powershell
.\regravity.exe locales
.\regravity.exe apply --locale hu
.\regravity.exe apply --locale en --session-metrics
.\regravity.exe apply --locale en --effort-slider
.\regravity.exe status
.\regravity.exe restore
.\regravity.exe --help
```

A Windows-verzió elkészítéséhez PowerShell 7, Go 1.26.3+ és Wails v2 CLI szükséges. Futtassa a projekt mappájában a `.\build.ps1 -Target gui` parancsot. Az eredmény a `build/bin/regravity.exe` fájl. Ha hiányzik a WebView2, telepítse a [Microsoft Edge WebView2 Runtime](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download) összetevőt, majd indítsa újra a Regravityt.

## Licenc

A Regravity független közösségi projekt. A Google és az Antigravity nem áll kapcsolatban vele, és hivatalosan nem támogatja. A Google Antigravity és a kapcsolódó védjegyek a tulajdonosaikat illetik. Az eszköz személyes és oktatási célra használható, az érintett szoftver feltételeinek betartásával.

[MIT-licenc](./LICENSE)
