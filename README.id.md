<div align="center">
  <img src="./frontend/dist/appicon.png" width="68" height="68" alt="Ikon Regravity">
  <h1>Regravity</h1>
  <p>Alat portabel untuk Windows yang menerjemahkan Google Antigravity serta menambahkan metrik sesi dan penggeser tingkat penalaran. Tersedia 18 paket terjemahan untuk Antigravity. Jendela pengaturan Regravity mendukung 19 bahasa. Bahasa Ceko, Hungaria, dan Bulgaria tersedia dalam versi pengembangan.</p>
  <p>
    <a href="https://github.com/wadewu-ml/regravity/releases/latest/download/regravity.exe"><img src="https://img.shields.io/badge/Download-Latest_(Windows_x64)-0F9D58?style=flat&logo=windows&logoColor=white" alt="Unduh rilis stabil untuk Windows x64"></a>
    <a href="https://github.com/wadewu-ml/regravity/releases"><img src="https://img.shields.io/badge/Releases-GitHub-18181B?style=flat&logo=github" alt="Rilis di GitHub"></a>
    <img src="https://img.shields.io/badge/Antigravity-2.35.2_tested-18181B?style=flat" alt="Diuji dengan Antigravity 2.35.2">
    <a href="./LICENSE"><img src="https://img.shields.io/badge/License-MIT-18181B?style=flat" alt="Lisensi MIT"></a>
  </p>
  <p>
    <sub>
      <a href="./README.zh-CN.md">简体中文</a> · <a href="./README.md">English</a> · <a href="./README.zh-Hant.md">繁體中文</a> · <a href="./README.ja.md">日本語</a> · <a href="./README.ko.md">한국어</a> · <a href="./README.de.md">Deutsch</a> · <a href="./README.es.md">Español</a> · <a href="./README.fr.md">Français</a> · <a href="./README.pt-BR.md">Português (Brasil)</a> · <a href="./README.ru.md">Русский</a> · <a href="./README.tr.md">Türkçe</a> · <a href="./README.vi.md">Tiếng Việt</a> · <a href="./README.it.md">Italiano</a> · <a href="./README.ar.md">العربية</a> · <strong>Bahasa Indonesia</strong> · <a href="./README.pl.md">Polski</a> · <a href="./README.cs.md">Čeština</a> · <a href="./README.hu.md">Magyar</a> · <a href="./README.bg.md">Български</a>
    </sub>
  </p>
</div>

## Mulai menggunakan Regravity

1. Buka `regravity.exe`. Pilih **Bahasa Indonesia** di daftar bahasa di bagian atas untuk mengubah bahasa jendela pengaturan.
2. Jika lokasi Antigravity salah, klik **Ubah lokasi** dan pilih folder yang berisi `Antigravity.exe`. Di bagian **Bahasa Antigravity**, pilih **Bahasa Indonesia**. Aktifkan **Metrik sesi** dan **Penggeser tingkat penalaran** jika diperlukan.
3. Simpan pekerjaan Anda, keluar dari Antigravity, lalu klik **Terapkan perubahan**. Setelah Antigravity terbuka kembali, Regravity boleh ditutup.

> [!WARNING]
> Menutup Antigravity secara paksa akan menghentikan sesi aktif. Perubahan yang belum disimpan bisa hilang.

<p align="center"><a href="./docs/images/language-settings.webp"><img src="./docs/images/language-settings.webp" width="580" alt="Jendela pengaturan"></a></p>

## Metrik sesi dan penggeser tingkat penalaran

Arahkan penunjuk ke cincin di samping pemilih model untuk melihat sisa kuota model dan waktu reset dalam waktu setempat. Setelah respons pertama, panel juga menampilkan penggunaan konteks, kecepatan keluaran rata-rata, rasio hit cache, dan waktu hingga token pertama. Klik cincin agar panel tetap terbuka. Setiap percakapan memiliki metrik sendiri.

<p align="center"><a href="./docs/images/session-quota-before-first-call.webp"><img src="./docs/images/session-quota-before-first-call.webp" width="420" alt="Kuota model"></a></p>

Pada model yang mendukungnya, penggeser muncul di atas informasi penggunaan dalam menu model. Seret ke tingkat penalaran yang diinginkan lalu lepaskan. Bahasa panel dan penggeser mengikuti bahasa yang diterapkan ke Antigravity.

<p align="center"><a href="./docs/images/thinking-effort-high.webp"><img src="./docs/images/thinking-effort-high.webp" width="310" alt="Tingkat penalaran tinggi"></a> <a href="./docs/images/thinking-effort-max.webp"><img src="./docs/images/thinking-effort-max.webp" width="310" alt="Tingkat maksimum"></a></p>

## Pembaruan dan pemulihan

- Setelah memperbarui Regravity, klik **Terapkan ulang** untuk menggunakan terjemahan terbaru. Jika pembaruan Antigravity menghapus perubahan, periksa kompatibilitasnya lalu terapkan ulang.
- Titik biru pada menu **···** menandakan ada pembaruan. Periksa pembaruan dari menu tersebut atau nonaktifkan pemeriksaan otomatis melalui **Tentang**. Untuk v4.2.3 atau lebih lama, unduh `regravity.exe` terbaru secara manual.
- Untuk mengembalikan bahasa Inggris asli dan menghapus kedua fitur, keluar dari Antigravity lalu pilih **··· → Pulihkan versi asli**.

> Simpan cadangan `resources/app.asar.clean-backup` yang dibuat saat perubahan pertama kali diterapkan. Regravity membutuhkannya untuk memulihkan aplikasi asli.

## PowerShell

Jalankan perintah berikut dari folder `regravity.exe`.

```powershell
.\regravity.exe locales
.\regravity.exe apply --locale id
.\regravity.exe apply --locale en --session-metrics
.\regravity.exe apply --locale en --effort-slider
.\regravity.exe status
.\regravity.exe restore
.\regravity.exe --help
```

Untuk membangun aplikasi di Windows, gunakan PowerShell 7, Go 1.26.3+ dan Wails v2 CLI. Jalankan `.\build.ps1 -Target gui` di folder proyek. Hasilnya ada di `build/bin/regravity.exe`. Jika WebView2 belum terpasang, instal [Microsoft Edge WebView2 Runtime](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download) lalu buka kembali Regravity.

## Lisensi

Regravity adalah proyek komunitas independen tanpa afiliasi atau dukungan resmi dari Google maupun Antigravity. Google Antigravity dan merek dagangnya adalah milik pemilik masing-masing. Alat ini ditujukan untuk penggunaan pribadi dan pendidikan sesuai dengan ketentuan perangkat lunak yang digunakan.

[Lisensi MIT](./LICENSE)
