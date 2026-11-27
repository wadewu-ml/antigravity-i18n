<div align="center">
  <img src="./frontend/dist/appicon.png" width="68" height="68" alt="Ikona Regravity">
  <h1>Regravity</h1>
  <p>Přenosný nástroj pro Windows překládá Google Antigravity a přidává metriky relace a posuvník úrovně uvažování. Regravity nabízí 18 překladových balíčků a 19 jazyků nastavení včetně češtiny.</p>
  <p>
    <a href="https://github.com/wadewu-ml/regravity/releases/latest/download/regravity.exe"><img src="https://img.shields.io/badge/Download-Latest_(Windows_x64)-0F9D58?style=flat&logo=windows&logoColor=white" alt="Stáhnout stabilní verzi pro Windows x64"></a>
    <a href="https://github.com/wadewu-ml/regravity/releases"><img src="https://img.shields.io/badge/Releases-GitHub-18181B?style=flat&logo=github" alt="Vydání na GitHubu"></a>
    <img src="https://img.shields.io/badge/Antigravity-2.36.0_tested-18181B?style=flat" alt="Ověřeno s Antigravity 2.36.0">
    <a href="./LICENSE"><img src="https://img.shields.io/badge/License-MIT-18181B?style=flat" alt="Licence MIT"></a>
  </p>
  <p>
    <sub>
      <a href="./README.zh-CN.md">简体中文</a> · <a href="./README.md">English</a> · <a href="./README.zh-Hant.md">繁體中文</a> · <a href="./README.ja.md">日本語</a> · <a href="./README.ko.md">한국어</a> · <a href="./README.de.md">Deutsch</a> · <a href="./README.es.md">Español</a> · <a href="./README.fr.md">Français</a> · <a href="./README.pt-BR.md">Português (Brasil)</a> · <a href="./README.ru.md">Русский</a> · <a href="./README.tr.md">Türkçe</a> · <a href="./README.vi.md">Tiếng Việt</a> · <a href="./README.it.md">Italiano</a> · <a href="./README.ar.md">العربية</a> · <a href="./README.id.md">Bahasa Indonesia</a> · <a href="./README.pl.md">Polski</a> · <strong>Čeština</strong> · <a href="./README.hu.md">Magyar</a> · <a href="./README.bg.md">Български</a>
    </sub>
  </p>
</div>

## Začínáme

1. Spusťte `regravity.exe` a v horní nabídce vyberte **Čeština**. Tím změníte jazyk nastavení Regravity.
2. V části **Jazyk Antigravity** vyberte **Čeština**. Podle potřeby zapněte **Metriky relace** a **Posuvník úrovně uvažování**. Pokud cesta nesouhlasí, klikněte na **Změnit umístění** a vyberte složku s `Antigravity.exe`.
3. Uložte práci, zavřete Antigravity a klikněte na **Použít změny**. Jakmile se Antigravity znovu spustí, můžete Regravity zavřít.

> [!WARNING]
> Vynucené ukončení přeruší běžící relace a může vést ke ztrátě neuložených změn.

<p align="center"><a href="./docs/images/language-settings.webp"><img src="./docs/images/language-settings.webp" width="580" alt="Okno nastavení"></a></p>

## Metriky relace a úroveň uvažování

Najeďte na kroužek vedle výběru modelu. Zobrazí se zbývající kvóta a místní čas jejího obnovení. Po první odpovědi uvidíte také využití kontextu, průměrnou rychlost výstupu, podíl zásahů do mezipaměti a čas do prvního tokenu. Kliknutím na kroužek panel připnete. Každá konverzace má vlastní metriky.

<p align="center"><a href="./docs/images/session-quota-before-first-call.webp"><img src="./docs/images/session-quota-before-first-call.webp" width="420" alt="Zbývající kvóta modelu"></a></p>

Pokud model nabízí více úrovní uvažování, objeví se v nabídce modelu posuvník. Přetažením vyberte úroveň a uvolněním ji potvrďte. Panel i posuvník používají jazyk nastavený pro Antigravity.

<p align="center"><a href="./docs/images/thinking-effort-high.webp"><img src="./docs/images/thinking-effort-high.webp" width="310" alt="Vysoká úroveň uvažování"></a> <a href="./docs/images/thinking-effort-max.webp"><img src="./docs/images/thinking-effort-max.webp" width="310" alt="Maximální úroveň"></a></p>

## Aktualizace a obnovení

- Po aktualizaci Regravity klikněte na **Použít znovu**. Pokud aktualizace Antigravity přepíše překlad, ověřte kompatibilitu a změny znovu použijte.
- Modrá tečka u nabídky **···** upozorňuje na aktualizaci. V nabídce ji můžete vyhledat ručně. Automatické kontroly lze vypnout v okně **O aplikaci**. U verze v4.2.3 nebo starší stáhněte nový `regravity.exe` ručně.
- Chcete-li vrátit anglické rozhraní a odebrat obě funkce, zavřete Antigravity a vyberte **··· → Obnovit původní verzi**.

> Ponechte si soubor `resources/app.asar.clean-backup` vytvořený při prvním použití. Regravity z něj obnovuje původní aplikaci.

## PowerShell

Příkazy spusťte ve složce s `regravity.exe`.

```powershell
.\regravity.exe locales
.\regravity.exe apply --locale cs
.\regravity.exe apply --locale en --session-metrics
.\regravity.exe apply --locale en --effort-slider
.\regravity.exe status
.\regravity.exe restore
.\regravity.exe --help
```

K sestavení ve Windows potřebujete PowerShell 7, Go 1.26.3+ a Wails v2 CLI. Ve složce projektu spusťte `.\build.ps1 -Target gui`. Výsledek najdete v `build/bin/regravity.exe`. Pokud chybí WebView2, nainstalujte [Microsoft Edge WebView2 Runtime](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download) a znovu otevřete Regravity.

## Licence

Regravity je nezávislý komunitní projekt bez oficiálního spojení se společností Google či projektem Antigravity nebo jejich podpory. Google Antigravity a související ochranné známky patří příslušným vlastníkům. Nástroj je určen k osobnímu a vzdělávacímu použití v souladu s podmínkami daného softwaru.

[Licence MIT](./LICENSE)
