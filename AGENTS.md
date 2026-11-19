# antigravity-i18n

This portable Windows tool localizes the Antigravity desktop interface and adds optional session metrics. The GUI and CLI share the service in `internal/core`.

## Development

- Scope shutdown and archive writes to the selected installation. Preserve the clean backup, installation lock, archive checks, and atomic replacement used by apply and restore. Keep forced shutdown an explicit choice.
- Describe and change translations according to the UI paths the patch actually reaches. Preserve user input and conversation content. Keep session metrics optional and read the local session database in read-only mode.
- Build the Windows executable with `.\build.ps1 -Target gui`. The script reads `VERSION` and embeds the window icon. Run the relevant Go tests for code changes. For an archive round trip using the real app, set `ANTIGRAVITY_ASAR` to a clean official archive and run `go test -count=1 ./...`. The test uses a copy.

## Project documentation

- `README.md` is the default English guide. Start with what the tool does, where to download it, and how to apply and restore settings. Keep affected language guides in sync. Use actual control labels and describe compatibility within the verified scope.
- The GitHub About description should identify the app, platform, and main uses in one short sentence.
- For a release, check the commits since the previous tag. Write concise Chinese and English notes about user-visible changes and any upgrade steps. Publish the checksum as a separate attachment.
