# antigravity-i18n


UI-Lokalisierungstool für den Google Antigravity-Desktop-Client.


[简体中文](./README.zh-CN.md) | [English](./README.md) | [繁體中文](./README.zh-Hant.md) | [日本語](./README.ja.md) | [한국어](./README.ko.md) | **Deutsch** | [Español](./README.es.md) | [Français](./README.fr.md) | [Português (Brasil)](./README.pt-BR.md) | [Русский](./README.ru.md)

---

Eine eigenständige ausführbare Datei, die eine moderne grafische Benutzeroberfläche, eine reine Go-Asar-Engine und vollständige Sprachpakete bündelt. Läuft sofort ohne Abhängigkeiten von Node.js oder Laufzeitbibliotheken und unterstützt jederzeit eine bytegenaue Wiederherstellung des offiziellen Builds.

---

## Hauptmerkmale

- **Einzelne eigenständige Binärdatei**: ~12 MB eigenständige ausführbare Datei – keine Node.js-, npm- oder Python-Laufzeit erforderlich.
- **GUI- und CLI-Dual-Modus**: Öffnet standardmäßig ein sauberes Desktop-Fenster und bettet gleichzeitig vollständige CLI-Unterbefehle („apply“ / „restore“ / „status“ / „locales“) für Skripterstellung und Automatisierung ein.
- **Bytegenaue Wiederherstellung**: Erstellt automatisch ein makelloses Backup („app.asar.clean-backup“), bevor Änderungen vorgenommen werden. Beim Wiederherstellen wird ein Archiv wiederhergestellt, das Byte für Byte mit der offiziellen Version identisch ist (SHA-256-verifiziert).
- **Nicht-invasives Design**: Lokalisiert nur die Shell-Benutzeroberfläche, native Menüs und Beenden-Dialoge. Code-Editoren (Monaco), integrierte Terminals (xterm) und KI-Konversationsströme bleiben strikt unberührt.
- **Offline und zuverlässig**: Keine Netzwerkaktivität, keine Datenerfassung. Ausgestattet mit einer JavaScript-Syntaxvalidierung im Voraus und einem Transaktions-Rollback bei unerwarteten Fehlern.

---

## Download und Nutzung

### Herunterladen

Laden Sie die neueste „antigravity-i18n.exe“ von der Seite „Releases“ herunter und führen Sie sie direkt aus.

### 1. Grafische Oberfläche (empfohlen)

Doppelklicken Sie auf „antigravity-i18n.exe“:
1. Die App erkennt den Antigravity-Installationsordner automatisch (oder klicken Sie auf „Speicherort ändern“, um ihn manuell anzugeben).
2. Klicken Sie auf Ihre bevorzugte Sprache (durch Auswahl von Englisch wird die offizielle Version wiederhergestellt).
3. Klicken Sie auf „Übernehmen“ und bestätigen Sie den Vorgang. Sobald Antigravity beendet ist, aktualisiert das Tool die Ressourcen und startet Antigravity neu.

Wenn Antigravity für das Anwenden oder Wiederherstellen geschlossen werden muss, fordert das Tool die ausgewählte Installation zum regulären Beenden auf. Läuft sie nach 30 Sekunden noch, wird der Vorgang abgebrochen. Im Bestätigungsdialog können Sie das sofortige Beenden ausdrücklich auswählen.

### 2. Befehlszeilenschnittstelle

Öffnen Sie PowerShell im Ordner mit `antigravity-i18n.exe`.

```powershell
# Check current localization status and backups
.\antigravity-i18n.exe status

# Install a language pack (Simplified Chinese by default)
.\antigravity-i18n.exe apply --locale zh-CN

# Restore the official untranslated English version
.\antigravity-i18n.exe restore

# List all bundled language packs with coverage
.\antigravity-i18n.exe locales
```

Häufige Optionen:
- `--app-dir <Pfad>`: Geben Sie das Antigravity-Installationsverzeichnis manuell an (oder legen Sie es über `ANTIGRAVITY_APP_DIR` fest).
- `--locale <code>`: Zielsprachencode (z. B. `zh-CN`, `zh-Hant`, `ja`, `ko`, `de`, `es`, `fr`).
- „--no-restart“: Antigravity nach dem Patchen nicht neu starten.
- `--force` beendet die ausgewählte Installation sofort. Ungespeicherte Änderungen können verloren gehen.
- `--timeout <ms>` legt die Wartezeit bis zum Abbruch in Millisekunden fest. Der Standardwert ist `30000`.
- „--yes“: Unterdrückt die Bestätigungsaufforderung zum Herunterfahren.

Pro Installation ist jeweils nur ein Vorgang zum Anwenden oder Wiederherstellen möglich. Falls bereits ein Vorgang läuft, warten Sie auf dessen Abschluss und versuchen Sie es erneut. Mit `status` prüfen Sie die ausgewählte Installation und ihre Sicherung. Ein unlesbares Archiv oder eine nicht überprüfbare Anwendungsidentität führt zu einer Fehlermeldung.

---

## Gebündelte Sprachen

Alle Sprachpakete sind direkt in die ausführbare Datei eingebettet:

| Code | Sprache | Abdeckung | Hinweise |
| :--- | :--- | :--- | :--- |
| `zh-CN` | Simplified Chinese (简体中文) | 100% | Vollständige Abdeckung der Benutzeroberfläche, Menüs und Dialoge |
| `zh-Hant` | Traditional Chinese (繁體中文) | 100% | Vollständige Abdeckung mit verifizierten Fachbegriffen |
| `ja` | Japanese (日本語) | 100% | Vollständige Abdeckung mit verifizierten Fachbegriffen |
| `ko` | Korean (한국어) | 100% | Vollständige Abdeckung mit verifizierten Fachbegriffen |
| `de` | German (Deutsch) | 100% | Vollständige Abdeckung mit verifizierten Fachbegriffen |
| `es` | Spanish (Español) | 100% | Vollständige Abdeckung mit verifizierten Fachbegriffen |
| `fr` | French (Français) | 100% | Vollständige Abdeckung mit verifizierten Fachbegriffen |
| `pt-BR` | Portuguese (Brasil) | 100% | Vollständige Abdeckung mit verifizierten Fachbegriffen |
| `ru` | Russian (Русский) | 100% | Vollständige Abdeckung mit verifizierten Fachbegriffen |
| `en` | English | - | Offizieller Original-Build (zum Wiederherstellen auswählen) |

---

## FAQ & Mechanik

### 1. Was passiert nach einem offiziellen Antigravity-Update?
Offizielle Updates überschreiben „app.asar“. Öffnen Sie einfach dieses Tool erneut und klicken Sie erneut auf „Übernehmen“. Das Tool aktualisiert sein sauberes Backup automatisch mit der neuen offiziellen Version.

### 2. Sprachumschaltung und Wiederherstellung
Sie können direkt zwischen jeder Sprache wechseln; Das Tool ersetzt aktive Patches nahtlos, ohne dass zunächst eine manuelle Wiederherstellung erforderlich ist. Durch Klicken auf „Offizielles Englisch wiederherstellen“ kann der Client jederzeit in einen unberührten Zustand zurückversetzt werden.

### 3. Verwaltung der Sicherungsdateien
- `app.asar.clean-backup` enthält das Originalarchiv für die Wiederherstellung. Beim Anwenden oder Wiederherstellen einer ungepatchten Installation vergleicht das Tool die SHA-256-Prüfsummen und aktualisiert die Sicherung bei Abweichungen, auch nach einer Neuinstallation derselben Version. **Bewahren Sie diese Datei auf.**
- `app.asar.bak-*` sind Kopien mit Zeitstempel, aus denen das Tool eine fehlende oder ungültige Sicherung wiederherstellen kann. Sobald der Client funktioniert und `status` eine gültige Sicherung meldet, können Sie ältere Kopien löschen, um Speicherplatz freizugeben.

---

## Bauen aus der Quelle

**Voraussetzungen**:
- Go 1.26.3+
- [Wails v2](https://wails.io/) CLI
- PowerShell 7

```powershell
# Clone the repository
git clone https://github.com/wadewu-ml/antigravity-i18n.git
cd antigravity-i18n

# Build both GUI and CLI binaries
.\build.ps1

# Entfernen Sie das Kommentarzeichen der nächsten Zeile, um ein offizielles Archiv zu testen.
# $env:ANTIGRAVITY_ASAR = "$env:LOCALAPPDATA\Programs\antigravity\resources\app.asar.clean-backup"
go test ./...
```

---

## Haftungsausschluss

1. Dies ist ein inoffizielles, von der Community gepflegtes Projekt und steht in keiner Verbindung zu Google oder Antigravity, wird von diesen nicht unterstützt oder autorisiert.
2. Google Antigravity und zugehörige Marken gehören ihren jeweiligen Eigentümern.
3. Dieses Tool wird für den persönlichen und pädagogischen Gebrauch bereitgestellt. Bitte stellen Sie sicher, dass Sie die entsprechenden Software-Nutzungsbedingungen einhalten.

## Lizenz

[MIT-Lizenz](./LICENSE)
