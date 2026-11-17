#requires -Version 7.0
param(
    [string]$AppDir = "$env:LOCALAPPDATA\Programs\antigravity",
    [string]$RendererUrl,
    [string]$BundlePath,
    [string]$Locale = 'zh-CN',
    [string]$OutputDir = (Join-Path ([IO.Path]::GetTempPath()) 'antigravity-i18n-ui-audit')
)

# Read-only audit of the static UI resources shipped by a running Antigravity.
# This is a candidate list for human review, not a percentage of translated UI.
# Running-app audits also read the public layout copy and onboarding presets.
# Bootstrap HTML is inspected in memory only: it contains session credentials.
# Requires PowerShell 7; no npm packages or app modifications are involved.
$ErrorActionPreference = 'Stop'
$repo = Split-Path -Parent $PSScriptRoot
$pack = Get-Content -LiteralPath (Join-Path $repo "assets/locales/$Locale.json") -Raw | ConvertFrom-Json -AsHashtable
$known = [Collections.Generic.HashSet[string]]::new([StringComparer]::Ordinal)
foreach ($key in $pack.text.Keys) { [void]$known.Add(($key -replace '\s+', ' ').Trim()) }
[void][IO.Directory]::CreateDirectory($OutputDir)
$version = (Get-Item -LiteralPath (Join-Path $AppDir 'Antigravity.exe')).VersionInfo.ProductVersion
$inputMode = if ($BundlePath) { 'provided-bundle' } else { 'running-app' }
$bootstrap = $null
$runtimeCopy = [Collections.Generic.List[object]]::new()
$runtimeSources = [Collections.Generic.List[string]]::new()

if (-not $BundlePath) {
    if (-not $RendererUrl) {
        $binaryPath = Join-Path $AppDir 'resources/bin/language_server.exe'
        $processIds = @(Get-Process -Name language_server -ErrorAction SilentlyContinue |
            Where-Object Path -eq $binaryPath | Select-Object -ExpandProperty Id)
        $ports = @(Get-NetTCPConnection -State Listen | Where-Object OwningProcess -in $processIds |
            Select-Object -ExpandProperty LocalPort -Unique)
        foreach ($port in $ports) {
            foreach ($scheme in @('http', 'https')) {
                $origin = "${scheme}://127.0.0.1:$port"
                try {
                    $bootstrap = Invoke-WebRequest -Uri "$origin/" -SkipCertificateCheck -TimeoutSec 3
                    if ($bootstrap.Content -match 'window\.__APP_CONFIG__' -and $bootstrap.Content -match 'src="/main\.js"') {
                        $RendererUrl = "$origin/main.js"
                        break
                    }
                } catch { continue }
            }
            if ($RendererUrl) { break }
        }
        if (-not $RendererUrl) { throw 'No running Antigravity renderer found. Start the app or pass -BundlePath / -RendererUrl.' }
    }
    $uri = [Uri]$RendererUrl
    if (-not $uri.IsLoopback) { throw 'RendererUrl must be a loopback URL serving the installed app.' }
    $BundlePath = Join-Path $OutputDir 'renderer-main.js'
    Invoke-WebRequest -Uri $uri -SkipCertificateCheck -TimeoutSec 30 -OutFile $BundlePath

    $origin = $uri.GetLeftPart([UriPartial]::Authority)
    if (-not $bootstrap) { $bootstrap = Invoke-WebRequest -Uri "$origin/" -SkipCertificateCheck -TimeoutSec 3 }
    $configMatch = [regex]::Match($bootstrap.Content, 'window\.__APP_CONFIG__\s*=\s*(\{[^\r\n]*?\});')
    if ($configMatch.Success) {
        $csrf = ($configMatch.Groups[1].Value | ConvertFrom-Json -AsHashtable).csrfToken
        $headers = @{ 'x-codeium-csrf-token' = $csrf; 'Connect-Protocol-Version' = '1' }
        $service = "$origin/exa.language_server_pb.LanguageServerService"
        try {
            $flags = Invoke-RestMethod -Uri "$service/GetMendelFlags" -Method Post -Body '{}' -ContentType 'application/json' -Headers $headers -SkipCertificateCheck -TimeoutSec 10
            foreach ($flag in $flags.experimentConfig.experiments | Where-Object { $_.keyString -eq 'project-4k-mode-copy' -and -not $_.disabled }) {
                $layout = $flag.string | ConvertFrom-Json -AsHashtable
                foreach ($mode in @('project4k', 'default')) {
                    foreach ($field in @('title', 'description')) {
                        if ($layout[$mode][$field]) { $runtimeCopy.Add([ordered]@{ text = $layout[$mode][$field]; source = "project-4k-mode-copy.$mode.$field" }) }
                    }
                }
            }
            $runtimeSources.Add('project-4k-mode-copy')
        } catch { Write-Warning 'Could not read the public layout copy; check the rendered layout cards.' }
        try {
            $onboarding = Invoke-RestMethod -Uri "$service/GetOnboardingPersonas" -Method Post -Body '{}' -ContentType 'application/json' -Headers $headers -SkipCertificateCheck -TimeoutSec 10
            # Only server-provided public presets. Never traverse customPersona,
            # customizations, account data, or conversation endpoints.
            foreach ($persona in $onboarding.personas) {
                foreach ($field in @('displayName', 'description')) {
                    if ($persona.$field) { $runtimeCopy.Add([ordered]@{ text = $persona.$field; source = "onboarding.personas.$($persona.id).$field" }) }
                }
                foreach ($suggestion in $persona.suggestions) {
                    if ($suggestion.label) { $runtimeCopy.Add([ordered]@{ text = $suggestion.label; source = "onboarding.personas.$($persona.id).suggestion" }) }
                }
            }
            $runtimeSources.Add('public-onboarding-personas')
        } catch { Write-Warning 'Could not read the public onboarding copy; check the rendered preset cards.' }
        $csrf = $null
        $headers = $null
    }
    $bootstrap = $null
}

$source = [IO.File]::ReadAllText((Resolve-Path -LiteralPath $BundlePath))
$candidates = [Collections.Generic.Dictionary[string,object]]::new([StringComparer]::Ordinal)
# The release bundle uses JSON-compatible double-quoted string literals. Capture
# source context too so reviewers can distinguish UI from library descriptions,
# tool prompts, examples, user content and fragments split by inline elements.
$literals = [regex]::Matches($source, '"(?:[^"\\\r\n]|\\.)*"')
foreach ($literal in $literals) {
    try { $value = ConvertFrom-Json -InputObject $literal.Value -NoEnumerate } catch { continue }
    if ($value -isnot [string] -or $value.Length -lt 2 -or $value.Length -gt 1800) { continue }
    $value = ($value -replace '\s+', ' ').Trim()
    if ($value -notmatch '[A-Za-z]' -or $known.Contains($value)) { continue }
    $beforeStart = [Math]::Max(0, $literal.Index - 160)
    $before = $source.Substring($beforeStart, $literal.Index - $beforeStart)
    $afterStart = $literal.Index + $literal.Length
    $after = $source.Substring($afterStart, [Math]::Min(160, $source.Length - $afterStart))
    $property = [regex]::Match($before, '(?:^|[,\{])(?:"(?<quoted>aria-label|data-tooltip-content)"|(?<plain>children|label|title|description|placeholder|tooltip|content|message|detail|text|name))\s*:\s*$')
    $kind = if ($property.Success) {
        if ($property.Groups['quoted'].Success) { $property.Groups['quoted'].Value } else { $property.Groups['plain'].Value }
    } elseif ($value -match '[A-Za-z]{2,}[ ,][A-Za-z]{2,}' -and $value -notmatch '[{}<>]|https?://|^\s*function\b') { 'other-string' } else { continue }
    if ($candidates.ContainsKey($value)) {
        $candidates[$value].occurrences++
        if ($kind -eq 'other-string') { continue }
        if ($candidates[$value].kind -ne 'other-string') { continue }
    }
    $candidates[$value] = [ordered]@{ text = $value; kind = $kind; offset = $literal.Index; before = $before; after = $after; occurrences = 1 }
}

foreach ($copy in $runtimeCopy) {
    $value = ($copy.text -replace '\s+', ' ').Trim()
    if ($value -and -not $known.Contains($value) -and -not $candidates.ContainsKey($value)) {
        $candidates[$value] = [ordered]@{ text = $value; kind = 'runtime-copy'; offset = $null; before = $copy.source; after = ''; occurrences = 1 }
    }
}
[IO.File]::WriteAllText((Join-Path $OutputDir 'runtime-ui.json'), (ConvertTo-Json -InputObject @($runtimeCopy) -Depth 4) + "`n")

$rows = @($candidates.Values | Sort-Object @{ Expression = { $_.kind -eq 'other-string' } }, @{ Expression = { $_.text } })
$allPath = Join-Path $OutputDir 'candidates.json'
$uiPath = Join-Path $OutputDir 'ui-candidates.json'
[IO.File]::WriteAllText($allPath, (ConvertTo-Json -InputObject $rows -Depth 6) + "`n")
$uiRows = @($rows | Where-Object { $_.kind -ne 'other-string' })
[IO.File]::WriteAllText($uiPath, (ConvertTo-Json -InputObject $uiRows -Depth 6) + "`n")
$metadata = [ordered]@{
    installedAppVersion = $version
    inputMode = $inputMode
    locale = $Locale
    bundleSHA256 = (Get-FileHash -LiteralPath $BundlePath -Algorithm SHA256).Hash
    uiCandidates = $uiRows.Count
    allCandidates = $rows.Count
    runtimeSources = @($runtimeSources)
    runtimeStrings = $runtimeCopy.Count
    note = 'Static literals plus the listed public runtime sources; other runtime copy still requires checking the rendered UI. Candidates require contextual review; library strings and partial sentences may be intentionally excluded. Existing regex translations are not counted as covered.'
}
[IO.File]::WriteAllText((Join-Path $OutputDir 'metadata.json'), (ConvertTo-Json -InputObject $metadata) + "`n")
[PSCustomObject]@{ Version = $version; UICandidates = $uiRows.Count; AllCandidates = $rows.Count; OutputDir = $OutputDir }
