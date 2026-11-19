# <img src="./frontend/dist/appicon.png" width="32" height="32" alt=""> antigravity-i18n

Neun Sprachpakete übersetzen Einstellungen, Bedienelemente im Chat, Menüs und Rückfragen beim Beenden von Google Antigravity für Windows. Auf Wunsch sehen Sie direkt am Eingabefeld, wie viel Kontext noch verfügbar ist und wie schnell das Modell geantwortet hat. Nach dem Anwenden können Sie das Tool schließen.

[Für Windows x64 herunterladen](https://github.com/wadewu-ml/antigravity-i18n/releases/latest/download/antigravity-i18n.exe) · [Versionen](https://github.com/wadewu-ml/antigravity-i18n/releases)

Getestet mit **Antigravity App 2.33.0**.

[简体中文](./README.zh-CN.md) · [English](./README.md) · [繁體中文](./README.zh-Hant.md) · [日本語](./README.ja.md) · [한국어](./README.ko.md) · **Deutsch** · [Español](./README.es.md) · [Français](./README.fr.md) · [Português (Brasil)](./README.pt-BR.md) · [Русский](./README.ru.md)

<a href="./docs/images/antigravity-zh-cn.png"><img src="./docs/images/translated-window.webp" width="680" alt="Antigravity mit Menüs auf vereinfachtem Chinesisch und dem Ring neben der Modellauswahl"></a>

## Oberflächensprache ändern

Zur Auswahl stehen vereinfachtes und traditionelles Chinesisch, Japanisch, Koreanisch, Deutsch, Spanisch, Französisch, brasilianisches Portugiesisch und Russisch. Das Tool selbst ist auf vereinfachtem Chinesisch beschriftet.

1. Öffnen Sie `antigravity-i18n.exe`. Falls das Tool die Installation nicht findet oder den falschen Ordner anzeigt, klicken Sie auf **更改位置** und wählen Sie den Ordner mit `Antigravity.exe`.
2. Wählen Sie unter **界面语言** eine Sprache. Mit **会话指标** schalten Sie die Sitzungsmetriken ein. Für Metriken mit der englischen Originaloberfläche wählen Sie **English**.
3. Speichern Sie Ihre Arbeit und beenden Sie Antigravity. Klicken Sie dann auf **应用更改**, um die Einstellungen zu übernehmen. Antigravity öffnet sich mit den neuen Einstellungen.

Mit **强制退出并重试** können Sie bei einer entsprechenden Aufforderung das Beenden erzwingen und den Vorgang wiederholen. Laufende Sitzungen werden dabei unterbrochen und ungespeicherte Änderungen können verloren gehen.

<a href="./docs/images/language-settings.webp"><img src="./docs/images/language-settings.webp" width="640" alt="Einstellungsfenster mit vereinfachtem Chinesisch als Sprache und aktivierten Sitzungsmetriken"></a>


## Sitzungsmetriken ablesen

Sobald ein Modellaufruf für den Chat aufgezeichnet wurde, erscheint rechts neben der Modellauswahl im Eingabefeld ein Ring. Er zeigt an, welcher Anteil des Kontextfensters schätzungsweise belegt ist. Bewegen Sie den Mauszeiger auf den Ring, um die verbleibenden Token, die Ausgabegeschwindigkeit, die Cache-Trefferquote und die Wartezeit bis zum ersten Token zu sehen. Ein Klick hält die Details offen, ein weiterer schließt sie.

<a href="./docs/images/session-metrics-zh-cn.png"><img src="./docs/images/session-metrics-detail.webp" width="560" alt="Das chinesische Metrikfenster über dem Ring zeigt Kontextnutzung, verbleibende Token und Ausgabegeschwindigkeit"></a>

In geteilten Ansichten, Nebenunterhaltungen und Bereichen für Subagenten stammen die Werte jeweils aus dem letzten lokal aufgezeichneten Modellaufruf der betreffenden Sitzung, für den Daten vorliegen. Die Ausgabemenge umfasst auch die vom Modell gemeldeten Reasoning-Token. Ein `—` kennzeichnet einen fehlenden Wert.

## Wiederherstellen oder erneut anwenden

Um zur englischen Originaloberfläche zurückzukehren und die Anzeige zu entfernen, beenden Sie Antigravity, klicken Sie im Tool auf **还原官方版本** und bestätigen Sie mit **还原**. Beim ersten Anwenden sichert das Tool die Originaldatei. Bewahren Sie `resources/app.asar.clean-backup` im Installationsordner von Antigravity auf.

Falls Übersetzung oder Anzeige nach einem Antigravity-Update verschwinden, vergleichen Sie Ihre Version mit der oben angegebenen unterstützten Version. Wählen Sie dann Sprache und Metriken erneut aus und wenden Sie die Einstellungen an. Bereits aktive Einstellungen können Sie mit **重新应用** erneut anwenden.

## PowerShell-Befehle

Öffnen Sie PowerShell im Ordner mit `antigravity-i18n.exe`.

| Aufgabe | Befehl |
| --- | --- |
| Sprachcodes auflisten | `.\antigravity-i18n.exe locales` |
| Deutsch anwenden | `.\antigravity-i18n.exe apply --locale de` |
| Englisch mit Sitzungsmetriken verwenden | `.\antigravity-i18n.exe apply --locale en --session-metrics` |
| Aktuelle Einstellungen anzeigen | `.\antigravity-i18n.exe status` |
| Offizielle Version wiederherstellen | `.\antigravity-i18n.exe restore` |
| Alle Optionen anzeigen | `.\antigravity-i18n.exe --help` |

## Aus dem Quellcode bauen

Für den Windows-Build benötigen Sie PowerShell 7, Go 1.26.3 oder neuer und die Wails-v2-CLI. Führen Sie `.\build.ps1 -Target gui` im Repository-Stamm aus. Die Datei entsteht unter `build/bin/antigravity-i18n.exe`.

## Probleme beim Start

Falls das Tool beim Start nach WebView2 fragt, installieren Sie die [Microsoft Edge WebView2 Runtime](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download).

## Haftungsausschluss

1. Dies ist ein inoffizielles, von der Community gepflegtes Projekt und steht in keiner Verbindung zu Google oder Antigravity, wird von diesen nicht unterstützt oder autorisiert.
2. Google Antigravity und zugehörige Marken gehören ihren jeweiligen Eigentümern.
3. Dieses Tool wird für den persönlichen und pädagogischen Gebrauch bereitgestellt. Bitte stellen Sie sicher, dass Sie die entsprechenden Software-Nutzungsbedingungen einhalten.

[MIT-Lizenz](./LICENSE)
