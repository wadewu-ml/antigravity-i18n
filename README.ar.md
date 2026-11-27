<div align="center">
  <img src="./frontend/dist/appicon.png" width="68" height="68" alt="أيقونة Regravity">
  <h1>Regravity</h1>
  <p>أداة محمولة لنظام Windows تترجم واجهة Google Antigravity وتعرض مقاييس الجلسة وحصص النماذج وتتيح ضبط مستوى الاستدلال. تتوفر لواجهة Antigravity 18 حزمة ترجمة، وتدعم نافذة إعدادات Regravity 19 لغة.</p>
  <p>
    <a href="https://github.com/wadewu-ml/regravity/releases/latest/download/regravity.exe"><img src="https://img.shields.io/badge/Download-Latest_(Windows_x64)-0F9D58?style=flat&logo=windows&logoColor=white" alt="تنزيل الإصدار المستقر لنظام Windows x64"></a>
    <a href="https://github.com/wadewu-ml/regravity/releases"><img src="https://img.shields.io/badge/Releases-GitHub-18181B?style=flat&logo=github" alt="الإصدارات على GitHub"></a>
    <img src="https://img.shields.io/badge/Antigravity-2.36.0_tested-18181B?style=flat" alt="تم اختباره مع Antigravity 2.36.0">
    <a href="./LICENSE"><img src="https://img.shields.io/badge/License-MIT-18181B?style=flat" alt="ترخيص MIT"></a>
  </p>
  <p>
    <sub>
      <a href="./README.zh-CN.md">简体中文</a> · <a href="./README.md">English</a> · <a href="./README.zh-Hant.md">繁體中文</a> · <a href="./README.ja.md">日本語</a> · <a href="./README.ko.md">한국어</a> · <a href="./README.de.md">Deutsch</a> · <a href="./README.es.md">Español</a> · <a href="./README.fr.md">Français</a> · <a href="./README.pt-BR.md">Português (Brasil)</a> · <a href="./README.ru.md">Русский</a> · <a href="./README.tr.md">Türkçe</a> · <a href="./README.vi.md">Tiếng Việt</a> · <a href="./README.it.md">Italiano</a> · <strong>العربية</strong> · <a href="./README.id.md">Bahasa Indonesia</a> · <a href="./README.pl.md">Polski</a> · <a href="./README.cs.md">Čeština</a> · <a href="./README.hu.md">Magyar</a> · <a href="./README.bg.md">Български</a>
    </sub>
  </p>
</div>

## بدء الاستخدام

اختر **العربية** من قائمة اللغات أعلى نافذة Regravity لتغيير لغة الأداة نفسها.

1. افتح `regravity.exe`. إذا كان المسار غير صحيح، فاضغط **تغيير الموقع** واختر المجلد الذي يحتوي على `Antigravity.exe`.
2. في قسم **لغة Antigravity** اختر **العربية**. يمكنك تفعيل مقاييس الجلسة وشريط مستوى الاستدلال عند الحاجة. تعمل الميزتان أيضًا مع الواجهة الإنجليزية.
3. احفظ عملك واخرج من Antigravity ثم اضغط **تطبيق التغييرات**. بعد إعادة فتح Antigravity يمكنك إغلاق Regravity.

> [!WARNING]
> يؤدي فرض الإغلاق إلى قطع الجلسات الجارية، وقد تفقد التعديلات غير المحفوظة.

<p align="center"><a href="./docs/images/language-settings.webp"><img src="./docs/images/language-settings.webp" width="580" alt="نافذة الإعدادات"></a></p>

## حصص النماذج ومقاييس الجلسة ومستوى الاستدلال

مرّر المؤشر فوق الحلقة بجوار محدد النموذج للاطلاع على الحصة المتبقية وموعد إعادة ضبطها بالتوقيت المحلي. بعد أول استدعاء للنموذج تظهر أيضًا كمية السياق المستخدم ومتوسط سرعة توليد الرموز ونسبة إصابات ذاكرة التخزين المؤقت وزمن ظهور أول رمز. انقر لتثبيت التفاصيل. لكل محادثة مقاييسها الخاصة.

<p align="center"><a href="./docs/images/session-quota-before-first-call.webp"><img src="./docs/images/session-quota-before-first-call.webp" width="420" alt="حصص النماذج"></a></p>

في النماذج المدعومة يظهر شريط مستوى الاستدلال فوق خيار عرض الاستخدام في قائمة النموذج. اسحب المقبض ثم اتركه لتطبيق المستوى؛ يزداد المستوى باتجاه اليمين.

<p align="center"><a href="./docs/images/thinking-effort-high.webp"><img src="./docs/images/thinking-effort-high.webp" width="310" alt="مستوى الاستدلال المرتفع"></a> <a href="./docs/images/thinking-effort-max.webp"><img src="./docs/images/thinking-effort-max.webp" width="310" alt="أعلى مستوى للاستدلال"></a></p>

## التحديث والاستعادة

- بعد تحديث Regravity اضغط **إعادة التطبيق** لاستخدام الترجمة الجديدة. إذا أزال تحديث Antigravity التخصيصات، فتحقق من توافق الإصدار وأعد تطبيق الإعدادات.
- تشير نقطة زرقاء على قائمة **···** إلى توفر تحديث. يمكنك البحث عنه يدويًا من القائمة أو تعطيل البحث التلقائي من نافذة **حول**. إذا كنت تستخدم v4.2.3 أو إصدارًا أقدم، فنزّل `regravity.exe` الجديد يدويًا.
- للعودة إلى الإنجليزية الأصلية وإزالة الميزتين، اخرج من Antigravity واختر **استعادة النسخة الأصلية** من قائمة **···**.

> [!IMPORTANT]
> احتفظ بملف `resources/app.asar.clean-backup` الذي يُنشأ عند التطبيق الأول لتتمكن من استعادة النسخة الأصلية.

## PowerShell

نفّذ الأوامر التالية من المجلد الذي يحتوي على `regravity.exe`.

```powershell
.\regravity.exe locales
.\regravity.exe apply --locale ar
.\regravity.exe apply --locale en --session-metrics
.\regravity.exe apply --locale en --effort-slider
.\regravity.exe status
.\regravity.exe restore
.\regravity.exe --help
```

للبناء على Windows تحتاج إلى PowerShell 7 وGo 1.26.3 أو أحدث وأداة Wails v2 CLI. نفّذ `.\build.ps1 -Target gui` من مجلد المشروع. ستجد البرنامج في `build/bin/regravity.exe`.

إذا ظهرت رسالة تفيد بعدم وجود WebView2، فثبّت [بيئة تشغيل Microsoft Edge WebView2](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download) ثم أعد فتح Regravity.

## الترخيص

Regravity مشروع مستقل يديره المجتمع، ولا تربطه علاقة رسمية بـ Google أو Antigravity ولا يحظى باعتمادهما. تعود العلامات التجارية إلى أصحابها. الأداة مخصصة للاستخدام الشخصي والتعليمي وفق شروط البرامج المعنية.

[ترخيص MIT](./LICENSE)
