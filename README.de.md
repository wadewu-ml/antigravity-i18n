<div align="center">
  <img src="./frontend/dist/appicon.png" width="64" height="64" alt="Regravity-Symbol">
  <h1>Regravity</h1>
  <p>Sprachpakete, Sitzungsmetriken & Modellkontingente sowie ein Regler für den Denkaufwand für Google Antigravity unter Windows.</p>
  <p>
    <a href="https://github.com/wadewu-ml/regravity/releases/download/v4.2.3/antigravity-i18n.exe"><strong>Für Windows x64 herunterladen</strong></a>
    ·
    <a href="https://github.com/wadewu-ml/regravity/releases">Versionen</a>
    ·
    Getestet mit <strong>Antigravity App 2.35.1</strong>
  </p>
  <p>
    <a href="./README.zh-CN.md">简体中文</a> · <a href="./README.md">English</a> · <a href="./README.zh-Hant.md">繁體中文</a> · <a href="./README.ja.md">日本語</a> · <a href="./README.ko.md">한국어</a> · <strong>Deutsch</strong> · <a href="./README.es.md">Español</a> · <a href="./README.fr.md">Français</a> · <a href="./README.pt-BR.md">Português (Brasil)</a> · <a href="./README.ru.md">Русский</a>
  </p>
  <p>antigravity-i18n heißt jetzt Regravity. Oben steht v4.2.3 zum Download. Laden Sie bei der ersten Regravity-Veröffentlichung <code>regravity.exe</code> manuell herunter. Vorhandene Patches und die Einstellung für automatische Updates werden übernommen. Seit der Umbenennung kann v4.2.3 beim Versionsvergleich einen Fehler melden. Die verfügbaren Versionen finden Sie über den Link oben.</p>
</div>

<p align="center">
  <a href="./docs/images/language-settings.webp">
    <img src="./docs/images/language-settings.webp" width="560" alt="Einstellungsfenster mit Antigravity v2.35.2, vereinfachtem Chinesisch sowie aktivierten Sitzungsmetriken und Denkaufwand-Regler">
  </a>
</p>

## Funktionen

- **9 Oberflächensprachen in einer portablen EXE** — Übersetzt Einstellungen, Chat-Bedienelemente, Modellmenüs, native Menüs und Rückfragen beim Beenden ins vereinfachte und traditionelle Chinesisch, Japanische, Koreanische, Deutsche, Spanische, Französische, brasilianische Portugiesisch und Russische, ohne Ihre Eingaben oder Unterhaltungen zu verändern.
- **Optionale Sitzungsmetriken & Modellkontingente** — Ergänzt neben der Modellauswahl einen Ring für verbleibende Kontingente, lokale Rücksetzzeiten, Kontextbelegung, mittlere Ausgaberate, Cache-Trefferrate und Zeit bis zum ersten Token.
- **Optionaler Regler für den Denkaufwand** — Fügt in der Modellauswahl einen Kapselregler für Modelle mit wählbarem Denkaufwand hinzu.
- **Unabhängig & bytegenau wiederherstellbar** — Beide Schalter funktionieren unabhängig voneinander und auch mit der englischen Originaloberfläche. Nach dem Anwenden können Sie das Tool schließen und mit einer unveränderten Sicherung derselben Antigravity-Version die Originalversion wiederherstellen.

## Schnellstart

> [!NOTE]
> Die Oberfläche des Einrichtungstools ist auf vereinfachtem Chinesisch beschriftet. Die Originalbeschriftungen stehen unten in Klammern, passend zum Screenshot oben.

### Einstellungen anwenden

1. Starten Sie `antigravity-i18n.exe`. Falls der erkannte Ordner nicht stimmt, klicken Sie auf **Speicherort ändern** (`更改位置`) und wählen Sie den Ordner mit `Antigravity.exe`.
2. Wählen Sie eine **Oberflächensprache** (`界面语言`) – oder lassen Sie **English** ausgewählt, wenn Sie nur die Zusatzfunktionen nutzen möchten.
3. Aktivieren Sie unterhalb der Sprachliste nach Bedarf **Sitzungsmetriken** (`会话指标`) und/oder **Regler für den Denkaufwand** (`思考强度滑块`).
4. Speichern Sie Ihre Arbeit, beenden Sie Antigravity und klicken Sie unten rechts auf **Änderungen anwenden** (`应用更改`). Antigravity öffnet sich mit den neuen Einstellungen (oder über **Antigravity öffnen** / `打开 Antigravity`); danach können Sie das Tool schließen.

> [!WARNING]
> Falls gemeldet wird, dass Antigravity noch läuft, speichern Sie Ihre Arbeit und beenden Sie die App regulär. **Beenden erzwingen und erneut versuchen** (`强制退出并重试`) unterbricht aktive Sitzungen und kann ungespeicherte Änderungen verwerfen.

### Erneut anwenden oder wiederherstellen

- **Aktuelle Einstellungen erneut anwenden** — Klicken Sie unten rechts auf **Erneut anwenden** (`重新应用`).
- **Originalversion wiederherstellen** — Beenden Sie Antigravity, öffnen Sie oben rechts das Menü (`···`), klicken Sie auf **Originalversion wiederherstellen** (`还原官方版本`) und bestätigen Sie. Dadurch wird die englische Originaloberfläche wiederhergestellt und Anzeige sowie Regler werden entfernt.

> [!IMPORTANT]
> Beim ersten Anwenden sichert das Tool die Originaldatei unter `resources/app.asar.clean-backup` im Installationsordner von Antigravity. Bewahren Sie diese Datei für eine spätere Wiederherstellung auf.

## Optionale Funktionen

### Sitzungsmetriken & Modellkontingente

- **Kontingente schon vor der ersten Nachricht** — Bewegen Sie den Mauszeiger auf den Ring neben der Modellauswahl, um das verbleibende Kontingent je Modellgruppe und Zeitraum (z. B. wöchentlich und 5-Stunden-Fenster) zu sehen. Über einem Prozentwert erscheint die lokale Rücksetzzeit.
- **Live-Metriken nach dem ersten Modellaufruf** — Sobald ein Modellaufruf aufgezeichnet wurde, zeigt der Ring die geschätzte Kontextauslastung und der obere Bereich füllt **Kontextbelegung** (genutzte Token, Gesamtkapazität und Anteil), **Mittlere Ausgaberate** (inkl. gemeldeter Reasoning-Token), **Cache-Trefferrate** und **Zeit bis zum ersten Token**. Ein Klick fixiert oder schließt das Panel.
- **Pro Bereich genau** — Geteilte Ansichten, Nebenunterhaltungen und Subagenten lesen jeweils den letzten lokal aufgezeichneten Modellaufruf mit Daten aus ihrer eigenen Sitzung. Ein `—` steht für einen nicht verfügbaren Wert.

<p align="center">
  <a href="./docs/images/session-quota-before-first-call.webp">
    <img src="./docs/images/session-quota-before-first-call.webp" width="440" alt="Sitzungsmetrik- und Modellkontingent-Panel über dem Ring neben der Modellauswahl">
  </a>
</p>

### Regler für den Denkaufwand

Aktivieren Sie **Regler für den Denkaufwand** (`思考强度滑块`) und wenden Sie die Änderungen an. Bei Modellen mit einstellbarem Denkaufwand erscheint in der Modellauswahl über der Nutzungsanzeige ein Regler — ziehen Sie ihn auf die gewünschte Stufe und lassen Sie los.

<p align="center">
  <a href="./docs/images/thinking-effort-high.webp">
    <img src="./docs/images/thinking-effort-high.webp" height="340" alt="Modellauswahl mit Denkaufwand auf Stufe Hoch">
  </a>
  &nbsp;&nbsp;
  <a href="./docs/images/thinking-effort-max.webp">
    <img src="./docs/images/thinking-effort-max.webp" height="340" alt="Modellauswahl mit Denkaufwand auf der höchsten Stufe">
  </a>
</p>

## Updates

- **Nach einem Antigravity-Update** — Falls Übersetzung oder Erweiterungen verschwinden, prüfen Sie oben die getestete Version und wenden Sie Ihre Einstellungen im Tool erneut an.
- **Automatische Tool-Updates** — Beim Start entscheidet das Ergebnis der letzten Prüfung, ob GitHub abgefragt wird. Nach erfolgreicher Prüfung vergehen mindestens 24 Stunden. Ein blauer Punkt am Menü oben rechts (`···`) zeigt eine neue Version an; dort können Sie die Hinweise lesen und das Update direkt am bisherigen EXE-Pfad laden und neu starten. Die automatische Suche lässt sich unter **Info** (`关于`) abschalten, die manuelle Suche (`检查更新`) bleibt verfügbar. Klicken Sie nach dem Update auf **Erneut anwenden** (`重新应用`), um neue Übersetzungen zu übernehmen.

## PowerShell-CLI

Öffnen Sie PowerShell im Ordner mit `antigravity-i18n.exe`:

| Aufgabe | Befehl |
| --- | --- |
| Sprachcodes auflisten | `.\antigravity-i18n.exe locales` |
| Deutsch anwenden | `.\antigravity-i18n.exe apply --locale de` |
| Englisch + Sitzungsmetriken | `.\antigravity-i18n.exe apply --locale en --session-metrics` |
| Englisch + Regler für Denkaufwand | `.\antigravity-i18n.exe apply --locale en --effort-slider` |
| Aktuelle Einstellungen prüfen | `.\antigravity-i18n.exe status` |
| Offizielle Version wiederherstellen | `.\antigravity-i18n.exe restore` |
| Alle Optionen anzeigen | `.\antigravity-i18n.exe --help` |

## Aus dem Quellcode bauen

Erfordert Windows, PowerShell 7, Go 1.26.3+ und die Wails-v2-CLI:

```powershell
.\build.ps1 -Target gui
```

Die fertige Datei liegt unter `build/bin/regravity.exe`.

## Fehlerbehebung

- **WebView2-Hinweis beim Start** — Installieren Sie die [Microsoft Edge WebView2 Runtime](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download) und starten Sie `antigravity-i18n.exe` erneut.

## Haftungsausschluss

1. Dies ist ein inoffizielles, von der Community gepflegtes Projekt und steht in keiner Verbindung zu Google oder Antigravity, wird von diesen nicht unterstützt oder autorisiert.
2. Google Antigravity und zugehörige Marken gehören ihren jeweiligen Eigentümern.
3. Dieses Tool wird für den persönlichen und pädagogischen Gebrauch bereitgestellt. Bitte stellen Sie sicher, dass Sie die entsprechenden Software-Nutzungsbedingungen einhalten.

[MIT-Lizenz](./LICENSE)
