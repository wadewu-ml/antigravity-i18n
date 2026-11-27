# Regravity

This portable Windows tool localizes the Antigravity desktop interface and adds optional session metrics and a thinking-effort slider. The GUI and CLI share the service in `internal/core`.

## Development

- Scope shutdown and archive writes to the selected installation. Preserve the clean backup, installation lock, archive checks, and atomic replacement used by apply and restore. Keep forced shutdown an explicit choice.
- Let automatic detection fill an unset installation path. If a status request finishes after the user selects another installation, keep the user's selection in both the Go backend and the GUI.
- Add further languages when users request them.
- Describe and change translations according to the UI paths the patch actually reaches. Preserve user input and conversation content. Keep session metrics optional. Read the local session database in read-only mode inside Antigravity so users can close the setup tool after applying their settings.
- Build the Windows executable with `.\build.ps1 -Target gui`. The script reads `VERSION` and embeds the window icon. Run the relevant Go tests for code changes. For an archive round trip using the real app, set `ANTIGRAVITY_ASAR` to a clean official archive and run `go test -count=1 ./...`. The test uses a copy.

## Project documentation

- `README.md` is the default English guide. Start with what the tool does, where to download it, and how to apply and restore settings. Keep affected language guides in sync. State the verified client version near the download and put missing-runtime help under troubleshooting.
- Write steps in the guide’s language. When the setup tool uses a different language, describe the control by its action or location and give the original label beside the screenshot if needed. Display compact screenshots and short command or build sections inline.
- When shortening a guide, preserve its three disclaimer clauses on affiliation, trademark ownership, and scope of use.
- The GitHub About description should identify the app, platform, and main uses in one short sentence.
- For a release, check the commits since the previous tag. Write concise Chinese and English notes about user-visible changes and any upgrade steps. Publish the checksum as a separate attachment.
