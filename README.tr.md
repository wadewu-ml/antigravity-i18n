<div align="center">
  <img src="./frontend/dist/appicon.png" width="68" height="68" alt="Regravity simgesi">
  <h1>Regravity</h1>
  <p><strong>Google Antigravity</strong> için taşınabilir Windows aracı. Arayüzü çevirir, oturum ölçümlerini ve model kotalarını gösterir, düşünme düzeyini ayarlamanızı sağlar. Antigravity için 18 çeviri paketi sunar. Regravity ayar penceresi 19 dili destekler.</p>
  <p>
    <a href="https://github.com/wadewu-ml/regravity/releases/latest/download/regravity.exe"><img src="https://img.shields.io/badge/Download-Latest_(Windows_x64)-0F9D58?style=flat&logo=windows&logoColor=white" alt="Windows x64 için son kararlı sürümü indir"></a>
    <a href="https://github.com/wadewu-ml/regravity/releases"><img src="https://img.shields.io/badge/Releases-GitHub-18181B?style=flat&logo=github" alt="GitHub sürümleri"></a>
    <img src="https://img.shields.io/badge/Antigravity-2.36.0_tested-18181B?style=flat" alt="Antigravity 2.36.0 ile test edildi">
    <a href="./LICENSE"><img src="https://img.shields.io/badge/License-MIT-18181B?style=flat" alt="MIT Lisansı"></a>
  </p>
  <p>
    <sub>
      <a href="./README.zh-CN.md">简体中文</a> · <a href="./README.md">English</a> · <a href="./README.zh-Hant.md">繁體中文</a> · <a href="./README.ja.md">日本語</a> · <a href="./README.ko.md">한국어</a> · <a href="./README.de.md">Deutsch</a> · <a href="./README.es.md">Español</a> · <a href="./README.fr.md">Français</a> · <a href="./README.pt-BR.md">Português (Brasil)</a> · <a href="./README.ru.md">Русский</a> · <strong>Türkçe</strong> · <a href="./README.vi.md">Tiếng Việt</a> · <a href="./README.it.md">Italiano</a> · <a href="./README.ar.md">العربية</a> · <a href="./README.id.md">Bahasa Indonesia</a> · <a href="./README.pl.md">Polski</a> · <a href="./README.cs.md">Čeština</a> · <a href="./README.hu.md">Magyar</a> · <a href="./README.bg.md">Български</a>
    </sub>
  </p>
</div>

## Başlarken

Ayar penceresinin dilini sağ üstteki listeden seçin.

1. `regravity.exe` dosyasını açın. Yanlış klasör seçildiyse **Change location** ile `Antigravity.exe` dosyasının bulunduğu klasörü seçin.
2. **Antigravity language** altında **Türkçe** seçin. İsterseniz **Session metrics** ve **Thinking effort slider** seçeneklerini açın. Bu özellikleri orijinal İngilizce arayüzle de kullanabilirsiniz.
3. Çalışmanızı kaydedip Antigravity’yi kapatın ve **Apply changes** düğmesine tıklayın. Antigravity yeniden açılınca Regravity’yi kapatabilirsiniz.

> [!WARNING]
> **Force quit and retry** açık oturumları keser. Kaydedilmemiş değişiklikleri korumak için önce Antigravity’den normal şekilde çıkın.

<p align="center">
  <a href="./docs/images/language-settings.webp"><img src="./docs/images/language-settings.webp" width="580" alt="Çince ayar penceresi, Antigravity 2.35.2"></a>
</p>

## Oturum ölçümleri ve düşünme düzeyi

Model seçicinin yanındaki halkanın üzerine gelerek kalan kotayı ve yerel sıfırlanma zamanını görebilirsiniz. İlk model çağrısından sonra bağlam kullanımı, ortalama çıktı hızı, önbellek isabet oranı ve ilk tokene kadar geçen süre de gösterilir. Ayrıntıları sabitlemek için halkaya tıklayın. Bölünmüş görünümler kendi oturumlarının ölçümlerini gösterir; veri henüz yoksa `—` görünür.

<p align="center">
  <a href="./docs/images/session-quota-before-first-call.webp"><img src="./docs/images/session-quota-before-first-call.webp" width="420" alt="İlk iletiden önce model kotası"></a>
</p>

Düşünme düzeyini destekleyen modellerde kaydırıcı, model menüsünde **View Usage** seçeneğinin üzerinde görünür. Yeni düzeyi seçmek için kaydırıp bırakın.

<p align="center">
  <a href="./docs/images/thinking-effort-high.webp"><img src="./docs/images/thinking-effort-high.webp" width="310" alt="Model menüsünde yüksek düşünme düzeyi"></a>
  <a href="./docs/images/thinking-effort-max.webp"><img src="./docs/images/thinking-effort-max.webp" width="310" alt="Model menüsünde en yüksek düşünme düzeyi"></a>
</p>

## Güncelleme ve geri yükleme

- Regravity’yi güncelledikten sonra yeni çeviriler için **Reapply** düğmesine tıklayın. Antigravity güncellemesinden sonra çeviriler kaybolursa desteklenen sürümü kontrol edip ayarlarınızı yeniden uygulayın.
- Üstteki **···** menüsünde mavi nokta görürseniz yeni sürüm vardır. **Check for updates** ile elle kontrol edebilir, otomatik kontrolü **About** bölümünden kapatabilirsiniz. v4.2.3 veya daha eski bir sürüm kullanıyorsanız yeni `regravity.exe` dosyasını elle indirin; ayarlarınız korunur.
- Orijinal İngilizce arayüze dönmek ve iki ek özelliği kaldırmak için Antigravity’yi kapatıp **··· → Restore original** seçeneğini onaylayın.

> [!IMPORTANT]
> İlk uygulamada oluşturulan `resources/app.asar.clean-backup` dosyasını saklayın.

## PowerShell

Komutları `regravity.exe` dosyasının bulunduğu klasörde çalıştırın.

| İşlem | Komut |
| --- | --- |
| Dilleri listele | `.\regravity.exe locales` |
| Türkçe uygula | `.\regravity.exe apply --locale tr` |
| İngilizce ve oturum ölçümleri | `.\regravity.exe apply --locale en --session-metrics` |
| İngilizce ve düşünme düzeyi kaydırıcısı | `.\regravity.exe apply --locale en --effort-slider` |
| Durumu göster | `.\regravity.exe status` |
| Orijinal sürüme dön | `.\regravity.exe restore` |
| Tüm seçenekler | `.\regravity.exe --help` |

Windows’ta kaynaktan derlemek için PowerShell 7, Go 1.26.3+ ve Wails v2 CLI gerekir. Proje klasöründe `.\build.ps1 -Target gui` çalıştırın. Çıktı: `build/bin/regravity.exe`.

Başlatırken WebView2 eksik uyarısı alırsanız [Microsoft Edge WebView2 Runtime](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download) kurup aracı tekrar açın.

## Lisans

Regravity, topluluk tarafından geliştirilen bağımsız bir araçtır. Google veya Antigravity ile bağlantısı ya da bu kuruluşların onayı yoktur. Google Antigravity ve ilgili ticari markalar kendi sahiplerine aittir. Araç kişisel ve eğitim amaçlıdır; ilgili yazılımın kullanım koşullarına uyun.

[MIT Lisansı](./LICENSE)
