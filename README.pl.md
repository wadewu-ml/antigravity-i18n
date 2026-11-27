<div align="center">
  <img src="./frontend/dist/appicon.png" width="68" height="68" alt="Ikona Regravity">
  <h1>Regravity</h1>
  <p>Przenośne narzędzie dla Windows do tłumaczenia Google Antigravity. Dodaje statystyki sesji i suwak poziomu rozumowania. Antigravity ma 18 pakietów tłumaczeń, a okno ustawień Regravity obsługuje 19 języków.</p>
  <p>
    <a href="https://github.com/wadewu-ml/regravity/releases/latest/download/regravity.exe"><img src="https://img.shields.io/badge/Download-Latest_(Windows_x64)-0F9D58?style=flat&logo=windows&logoColor=white" alt="Pobierz wydanie stabilne dla Windows x64"></a>
    <a href="https://github.com/wadewu-ml/regravity/releases"><img src="https://img.shields.io/badge/Releases-GitHub-18181B?style=flat&logo=github" alt="Wydania na GitHubie"></a>
    <img src="https://img.shields.io/badge/Antigravity-2.36.0_tested-18181B?style=flat" alt="Przetestowano z Antigravity 2.36.0">
    <a href="./LICENSE"><img src="https://img.shields.io/badge/License-MIT-18181B?style=flat" alt="Licencja MIT"></a>
  </p>
  <p>
    <sub>
      <a href="./README.zh-CN.md">简体中文</a> · <a href="./README.md">English</a> · <a href="./README.zh-Hant.md">繁體中文</a> · <a href="./README.ja.md">日本語</a> · <a href="./README.ko.md">한국어</a> · <a href="./README.de.md">Deutsch</a> · <a href="./README.es.md">Español</a> · <a href="./README.fr.md">Français</a> · <a href="./README.pt-BR.md">Português (Brasil)</a> · <a href="./README.ru.md">Русский</a> · <a href="./README.tr.md">Türkçe</a> · <a href="./README.vi.md">Tiếng Việt</a> · <a href="./README.it.md">Italiano</a> · <a href="./README.ar.md">العربية</a> · <a href="./README.id.md">Bahasa Indonesia</a> · <strong>Polski</strong> · <a href="./README.cs.md">Čeština</a> · <a href="./README.hu.md">Magyar</a> · <a href="./README.bg.md">Български</a>
    </sub>
  </p>
</div>

## Pierwsze kroki

1. Uruchom `regravity.exe`. Aby zmienić język okna ustawień, wybierz **Polski** z listy u góry.
2. Jeśli ścieżka do Antigravity jest błędna, kliknij **Zmień lokalizację** i wskaż folder z plikiem `Antigravity.exe`. W sekcji **Język Antigravity** wybierz **Polski**. W razie potrzeby włącz **Statystyki sesji** i **Suwak poziomu rozumowania**.
3. Zapisz pracę, zamknij Antigravity i kliknij **Zastosuj zmiany**. Gdy Antigravity uruchomi się ponownie, możesz zamknąć Regravity.

> [!WARNING]
> Wymuszenie zakończenia aplikacji przerywa aktywne sesje i grozi utratą niezapisanych zmian.

<p align="center"><a href="./docs/images/language-settings.webp"><img src="./docs/images/language-settings.webp" width="580" alt="Okno ustawień"></a></p>

## Statystyki sesji i suwak poziomu rozumowania

Najedź na pierścień obok listy modeli, aby zobaczyć pozostały limit modelu i godzinę jego odnowienia w czasie lokalnym. Po pierwszej odpowiedzi panel pokazuje również wykorzystanie kontekstu, średnią prędkość generowania, odsetek trafień w pamięci podręcznej i czas do pierwszego tokena. Kliknij pierścień, aby przypiąć panel. Każda rozmowa ma własne statystyki.

<p align="center"><a href="./docs/images/session-quota-before-first-call.webp"><img src="./docs/images/session-quota-before-first-call.webp" width="420" alt="Limity modeli"></a></p>

Jeśli model obsługuje kilka poziomów rozumowania, nad informacją o wykorzystaniu w menu modelu pojawi się suwak. Przeciągnij go i puść, aby wybrać poziom. Panel i suwak używają języka zastosowanego w Antigravity.

<p align="center"><a href="./docs/images/thinking-effort-high.webp"><img src="./docs/images/thinking-effort-high.webp" width="310" alt="Wysoki poziom rozumowania"></a> <a href="./docs/images/thinking-effort-max.webp"><img src="./docs/images/thinking-effort-max.webp" width="310" alt="Poziom maksymalny"></a></p>

## Aktualizacje i przywracanie

- Po aktualizacji Regravity kliknij **Zastosuj ponownie**, aby wczytać nowe tłumaczenia. Jeśli aktualizacja Antigravity usunie zmiany, sprawdź zgodność i zastosuj ustawienia ponownie.
- Niebieska kropka przy menu **···** oznacza dostępną aktualizację. W tym menu można ręcznie sprawdzić aktualizacje, a w oknie **O programie** wyłączyć sprawdzanie automatyczne. W przypadku wersji v4.2.3 lub starszej pobierz nowy plik `regravity.exe` ręcznie.
- Aby przywrócić oryginalny angielski interfejs i usunąć obie funkcje, zamknij Antigravity i wybierz **··· → Przywróć oryginał**.

> Zachowaj kopię `resources/app.asar.clean-backup` utworzoną przy pierwszym zastosowaniu zmian. Regravity używa jej do przywrócenia oryginalnej aplikacji.

## PowerShell

Uruchom polecenia z folderu zawierającego `regravity.exe`.

```powershell
.\regravity.exe locales
.\regravity.exe apply --locale pl
.\regravity.exe apply --locale en --session-metrics
.\regravity.exe apply --locale en --effort-slider
.\regravity.exe status
.\regravity.exe restore
.\regravity.exe --help
```

Do zbudowania aplikacji w Windows potrzebne są PowerShell 7, Go 1.26.3+ i Wails v2 CLI. Uruchom `.\build.ps1 -Target gui` w folderze projektu. Wynik znajdziesz w `build/bin/regravity.exe`. Jeśli brakuje WebView2, zainstaluj [Microsoft Edge WebView2 Runtime](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download) i ponownie otwórz Regravity.

## Licencja

Regravity to niezależny projekt społeczności. Nie jest powiązany z Google ani Antigravity i nie ma ich oficjalnego poparcia. Google Antigravity i powiązane znaki towarowe należą do ich właścicieli. Narzędzie służy do użytku osobistego i edukacyjnego zgodnie z warunkami korzystania z danego oprogramowania.

[Licencja MIT](./LICENSE)
