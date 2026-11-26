<div align="center">
  <img src="./frontend/dist/appicon.png" width="68" height="68" alt="Regravity-Symbol">
  <h1>Regravity</h1>
  <p><em>Antigravity hebt ab; Regravity holt es zurück in deine Sprache und unter deine Kontrolle.</em></p>
  <p>Portables Windows-Begleittool für <strong>Google Antigravity</strong> — 9 Oberflächensprachen, Sitzungsmetriken &amp; Modellkontingente sowie ein Denkaufwand-Regler.</p>
  <p>
    <a href="https://github.com/wadewu-ml/regravity/releases/latest/download/regravity.exe"><img src="https://img.shields.io/badge/Download-Latest_(Windows_x64)-0F9D58?style=flat&logo=windows&logoColor=white" alt="Aktuelles Regravity für Windows x64 herunterladen"></a>
    <a href="https://github.com/wadewu-ml/regravity/releases"><img src="https://img.shields.io/badge/Releases-GitHub-18181B?style=flat&logo=github" alt="GitHub-Versionen"></a>
    <img src="https://img.shields.io/badge/Antigravity-2.35.1_tested-18181B?style=flat" alt="Getestet mit Antigravity App 2.35.1">
    <a href="./LICENSE"><img src="https://img.shields.io/badge/License-MIT-18181B?style=flat" alt="MIT-Lizenz"></a>
  </p>
  <p>
    <sub>
      <a href="./README.zh-CN.md">简体中文</a> · <a href="./README.md">English</a> · <a href="./README.zh-Hant.md">繁體中文</a> · <a href="./README.ja.md">日本語</a> · <a href="./README.ko.md">한국어</a> · <strong>Deutsch</strong> · <a href="./README.es.md">Español</a> · <a href="./README.fr.md">Français</a> · <a href="./README.pt-BR.md">Português (Brasil)</a> · <a href="./README.ru.md">Русский</a> · <a href="./README.tr.md">Türkçe</a> · <a href="./README.vi.md">Tiếng Việt</a>
    </sub>
  </p>
</div>

<br>

<p align="center">
  <a href="./docs/images/language-settings.webp">
    <img src="./docs/images/language-settings.webp" width="580" alt="Einstellungsfenster mit Antigravity v2.35.2, vereinfachtem Chinesisch sowie aktivierten Sitzungsmetriken und Denkaufwand-Regler">
  </a>
  <br>
  <sub>Regravity-Einstellungsfenster mit Antigravity 2.35.2.</sub>
</p>

## Funktionsumfang

- **9 Oberflächensprachen, ohne Eingriff in Chats** — Übersetzt Einstellungen, Chat-Bedienelemente, Modellmenüs, native Menüs und Rückfragen beim Beenden (`zh-CN`, `zh-Hant`, `ja`, `ko`, `de`, `es`, `fr`, `pt-BR`, `ru`), ohne Ihre Eingaben oder Unterhaltungen zu verändern.
- **Sitzungsmetriken & Modellkontingente** *(optional)* — Ergänzt neben der Modellauswahl einen Ring für verbleibende Kontingente, lokale Rücksetzzeiten, Kontextbelegung, mittlere Ausgaberate, Cache-Trefferrate und Zeit bis zum ersten Token.
- **Regler für den Denkaufwand** *(optional)* — Integriert einen Kapselregler in die Modellauswahl für Modelle mit wählbarem Denkaufwand.
- **Modular & bytegenau umkehrbar** — Beide Schalter funktionieren unabhängig voneinander und auch mit der englischen Originaloberfläche. Mit einer unveränderten Sicherung derselben Antigravity-Version lässt sich jederzeit das Original wiederherstellen.

## Schnellstart

> [!TIP]
> Die chinesischen Beschriftungen in Klammern entsprechen dem Screenshot.

In der Entwicklungsversion wählen Sie oben rechts die Sprache des Einrichtungstools. Für Antigravity stehen zusätzlich Türkisch und Vietnamesisch zur Verfügung.

### Einstellungen anwenden

1. Starten Sie `regravity.exe`. Falls der erkannte Ordner nicht stimmt, klicken Sie auf **Speicherort ändern** (`更改位置`) und wählen Sie den Ordner mit `Antigravity.exe`.
2. Wählen Sie eine **Oberflächensprache** (`界面语言`) oder behalten Sie **English**, wenn Sie nur die Zusatzfunktionen nutzen möchten.
3. Aktivieren Sie unterhalb der Sprachliste **Sitzungsmetriken** (`会话指标`) und/oder **Regler für den Denkaufwand** (`思考强度滑块`).
4. Speichern Sie Ihre Arbeit, beenden Sie Antigravity und klicken Sie unten rechts auf **Änderungen anwenden** (`应用更改`). Antigravity öffnet sich mit den neuen Einstellungen (oder über **Antigravity öffnen** / `打开 Antigravity`), danach können Sie das Tool schließen.

> [!WARNING]
> Falls gemeldet wird, dass Antigravity noch läuft, speichern Sie Ihre Arbeit und beenden Sie die App regulär. **Beenden erzwingen und erneut versuchen** (`强制退出并重试`) unterbricht aktive Sitzungen und kann ungespeicherte Änderungen verwerfen.

### Erneut anwenden oder wiederherstellen

- **Erneut anwenden** — Klicken Sie nach dem Update auf **Erneut anwenden** (`重新应用`), um die neuen Übersetzungen zu verwenden.
- **Originalversion wiederherstellen** — Beenden Sie Antigravity, öffnen Sie oben rechts das Menü (`···`), klicken Sie auf **Originalversion wiederherstellen** (`还原官方版本`) und bestätigen Sie.

> [!IMPORTANT]
> Beim ersten Anwenden sichert das Tool die Originaldatei unter `resources/app.asar.clean-backup` im Antigravity-Installationsordner. Bewahren Sie diese Datei für eine spätere Wiederherstellung auf.

## Optionale Funktionen

### Sitzungsmetriken & Modellkontingente

- **Kontingente schon vor der ersten Nachricht** — Bewegen Sie den Mauszeiger auf den Ring neben der Modellauswahl, um das verbleibende Kontingent je Modellgruppe und Zeitraum (z. B. wöchentlich und 5-Stunden-Fenster) zu sehen; über einem Prozentwert erscheint die lokale Rücksetzzeit.
- **Live-Metriken nach einem Modellaufruf** — Der Ring zeigt die geschätzte Kontextauslastung, und oben erscheinen **Kontextbelegung** (Token, Kapazität, Anteil), **Mittlere Ausgaberate** (inkl. Reasoning-Token), **Cache-Trefferrate** und **Zeit bis zum ersten Token**. Ein Klick fixiert oder schließt das Panel.
- **Pro Bereich genau** — Geteilte Ansichten, Nebenunterhaltungen und Subagenten lesen jeweils den letzten lokal aufgezeichneten Modellaufruf ihrer eigenen Sitzung. Ein `—` markiert nicht verfügbare Werte.

<p align="center">
  <a href="./docs/images/session-quota-before-first-call.webp">
    <img src="./docs/images/session-quota-before-first-call.webp" width="420" alt="Sitzungsmetrik- und Modellkontingent-Panel über dem Ring neben der Modellauswahl">
  </a>
  <br>
  <sub>Kontingente sind sofort sichtbar; die Sitzungsmetriken füllen sich nach dem ersten Modellaufruf.</sub>
</p>

### Regler für den Denkaufwand

Aktivieren Sie **Regler für den Denkaufwand** (`思考强度滑块`) und wenden Sie die Änderungen an. Bei Modellen mit einstellbarem Denkaufwand erscheint in der Modellauswahl über der Nutzungsanzeige ein Regler — ziehen Sie ihn zur gewünschten Stufe und lassen Sie los.

<p align="center">
  <a href="./docs/images/thinking-effort-high.webp">
    <img src="./docs/images/thinking-effort-high.webp" height="340" alt="Modellauswahl mit Denkaufwand auf Stufe Hoch">
  </a>
  &nbsp;&nbsp;
  <a href="./docs/images/thinking-effort-max.webp">
    <img src="./docs/images/thinking-effort-max.webp" height="340" alt="Modellauswahl mit Denkaufwand auf der höchsten Stufe">
  </a>
  <br>
  <sub>Reguläre Stufe (Hoch) und hervorgehobene Maximalstufe direkt in der Modellauswahl.</sub>
</p>

## Updates & CLI

### Updates

- **Nach einem Antigravity-Update** — Falls Übersetzung oder Erweiterungen verschwinden, prüfen Sie oben die getestete Version und wenden Sie Ihre Einstellungen erneut an.
- **Automatische Tool-Updates** — Beim Start entscheidet das Tool anhand des vorherigen Ergebnisses, ob GitHub kontaktiert wird, und wartet nach einer erfolgreichen Prüfung mindestens 24 Stunden. Ein blauer Punkt am Menü oben rechts (`···`) zeigt eine neue Version an. Unter **Info** (`关于`) lässt sich die automatische Prüfung umschalten; manuelle Prüfungen (`检查更新`) bleiben verfügbar. Nach einem Update **Erneut anwenden** (`重新应用`) klicken.

Wenn Sie v4.2.3 oder älter verwenden, laden Sie [`regravity.exe`](https://github.com/wadewu-ml/regravity/releases/latest/download/regravity.exe) manuell herunter und starten Sie die Datei. Ihre Einstellungen bleiben erhalten.

### PowerShell-CLI

Öffnen Sie PowerShell im Ordner mit `regravity.exe`:

| Aufgabe | Befehl |
| --- | --- |
| Sprachcodes auflisten | `.\regravity.exe locales` |
| Deutsch anwenden | `.\regravity.exe apply --locale de` |
| Englisch + Sitzungsmetriken | `.\regravity.exe apply --locale en --session-metrics` |
| Englisch + Denkaufwand-Regler | `.\regravity.exe apply --locale en --effort-slider` |
| Aktuelle Einstellungen prüfen | `.\regravity.exe status` |
| Offizielle Version wiederherstellen | `.\regravity.exe restore` |
| Alle Optionen anzeigen | `.\regravity.exe --help` |

### Aus dem Quellcode bauen

Erfordert Windows, PowerShell 7, Go 1.26.3+ und die Wails-v2-CLI:

```powershell
.\build.ps1 -Target gui
```

Erzeugt `build/bin/regravity.exe`.

## Fehlerbehebung

- **WebView2-Hinweis beim Start** — Installieren Sie die [Microsoft Edge WebView2 Runtime](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download) und starten Sie `regravity.exe` erneut.

---

## Haftungsausschluss

1. Dies ist ein inoffizielles, von der Community gepflegtes Projekt und steht in keiner Verbindung zu Google oder Antigravity, wird von diesen nicht unterstützt oder autorisiert.
2. Google Antigravity und zugehörige Marken gehören ihren jeweiligen Eigentümern.
3. Dieses Tool wird für den persönlichen und pädagogischen Gebrauch bereitgestellt. Bitte stellen Sie sicher, dass Sie die entsprechenden Software-Nutzungsbedingungen einhalten.

[MIT-Lizenz](./LICENSE)
