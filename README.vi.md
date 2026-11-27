<div align="center">
  <img src="./frontend/dist/appicon.png" width="68" height="68" alt="Biểu tượng Regravity">
  <h1>Regravity</h1>
  <p>Công cụ chạy trực tiếp trên Windows dành cho <strong>Google Antigravity</strong>. Dịch giao diện, xem chỉ số phiên và hạn mức mô hình, điều chỉnh mức suy luận. Antigravity có 18 gói dịch. Cửa sổ cài đặt Regravity hỗ trợ 19 ngôn ngữ.</p>
  <p>
    <a href="https://github.com/wadewu-ml/regravity/releases/latest/download/regravity.exe"><img src="https://img.shields.io/badge/Download-Latest_(Windows_x64)-0F9D58?style=flat&logo=windows&logoColor=white" alt="Tải bản ổn định mới nhất cho Windows x64"></a>
    <a href="https://github.com/wadewu-ml/regravity/releases"><img src="https://img.shields.io/badge/Releases-GitHub-18181B?style=flat&logo=github" alt="Các bản phát hành trên GitHub"></a>
    <img src="https://img.shields.io/badge/Antigravity-2.36.0_tested-18181B?style=flat" alt="Đã kiểm tra với Antigravity 2.36.0">
    <a href="./LICENSE"><img src="https://img.shields.io/badge/License-MIT-18181B?style=flat" alt="Giấy phép MIT"></a>
  </p>
  <p>
    <sub>
      <a href="./README.zh-CN.md">简体中文</a> · <a href="./README.md">English</a> · <a href="./README.zh-Hant.md">繁體中文</a> · <a href="./README.ja.md">日本語</a> · <a href="./README.ko.md">한국어</a> · <a href="./README.de.md">Deutsch</a> · <a href="./README.es.md">Español</a> · <a href="./README.fr.md">Français</a> · <a href="./README.pt-BR.md">Português (Brasil)</a> · <a href="./README.ru.md">Русский</a> · <a href="./README.tr.md">Türkçe</a> · <strong>Tiếng Việt</strong> · <a href="./README.it.md">Italiano</a> · <a href="./README.ar.md">العربية</a> · <a href="./README.id.md">Bahasa Indonesia</a> · <a href="./README.pl.md">Polski</a> · <a href="./README.cs.md">Čeština</a> · <a href="./README.hu.md">Magyar</a> · <a href="./README.bg.md">Български</a>
    </sub>
  </p>
</div>

## Bắt đầu

Chọn ngôn ngữ cho cửa sổ cài đặt trong danh sách ở góc trên bên phải.

1. Mở `regravity.exe`. Nếu công cụ nhận nhầm vị trí, nhấn **Change location** và chọn thư mục chứa `Antigravity.exe`.
2. Trong **Antigravity language**, chọn **Tiếng Việt**. Bật **Session metrics** và **Thinking effort slider** nếu cần. Hai tính năng này cũng dùng được với giao diện tiếng Anh gốc.
3. Lưu công việc, thoát Antigravity rồi nhấn **Apply changes**. Khi Antigravity mở lại, bạn có thể đóng Regravity.

> [!WARNING]
> **Force quit and retry** sẽ ngắt các phiên đang chạy. Hãy thoát Antigravity theo cách thông thường để tránh mất thay đổi chưa lưu.

<p align="center">
  <a href="./docs/images/language-settings.webp"><img src="./docs/images/language-settings.webp" width="580" alt="Cửa sổ cài đặt bằng tiếng Trung, Antigravity 2.35.2"></a>
</p>

## Chỉ số phiên và mức suy luận

Di chuột lên vòng tròn cạnh bộ chọn mô hình để xem hạn mức còn lại và giờ đặt lại theo múi giờ địa phương. Sau lần gọi mô hình đầu tiên, bảng sẽ hiện lượng ngữ cảnh đã dùng, tốc độ đầu ra trung bình, tỷ lệ trúng bộ nhớ đệm và thời gian đến token đầu tiên. Nhấn vòng tròn để ghim bảng. Mỗi khung trò chuyện hiển thị số liệu của phiên tương ứng; chỉ số chưa có dữ liệu hiển thị `—`.

<p align="center">
  <a href="./docs/images/session-quota-before-first-call.webp"><img src="./docs/images/session-quota-before-first-call.webp" width="420" alt="Hạn mức mô hình trước tin nhắn đầu tiên"></a>
</p>

Với mô hình hỗ trợ điều chỉnh mức suy luận, thanh trượt xuất hiện ngay phía trên **View Usage** trong menu mô hình. Kéo đến mức mong muốn rồi thả để áp dụng.

<p align="center">
  <a href="./docs/images/thinking-effort-high.webp"><img src="./docs/images/thinking-effort-high.webp" width="310" alt="Mức suy luận cao trong menu mô hình"></a>
  <a href="./docs/images/thinking-effort-max.webp"><img src="./docs/images/thinking-effort-max.webp" width="310" alt="Mức suy luận cao nhất trong menu mô hình"></a>
</p>

## Cập nhật và khôi phục

- Sau khi cập nhật Regravity, nhấn **Reapply** để dùng bản dịch mới. Nếu giao diện trở lại tiếng Anh sau khi cập nhật Antigravity, hãy kiểm tra phiên bản được hỗ trợ rồi áp dụng lại cài đặt.
- Dấu chấm xanh trên menu **···** báo có bản cập nhật. Chọn **Check for updates** để kiểm tra thủ công, hoặc vào **About** để tắt kiểm tra tự động. Nếu đang dùng v4.2.3 trở về trước, hãy tải `regravity.exe` mới theo cách thủ công; cài đặt hiện tại được giữ lại.
- Để quay về giao diện tiếng Anh gốc và tắt hai tính năng bổ sung, hãy thoát Antigravity, chọn **··· → Restore original** rồi xác nhận.

> [!IMPORTANT]
> Hãy giữ tệp `resources/app.asar.clean-backup` được tạo khi áp dụng lần đầu.

## PowerShell

Chạy các lệnh trong thư mục chứa `regravity.exe`.

| Thao tác | Lệnh |
| --- | --- |
| Liệt kê ngôn ngữ | `.\regravity.exe locales` |
| Áp dụng tiếng Việt | `.\regravity.exe apply --locale vi` |
| Tiếng Anh và chỉ số phiên | `.\regravity.exe apply --locale en --session-metrics` |
| Tiếng Anh và thanh trượt mức suy luận | `.\regravity.exe apply --locale en --effort-slider` |
| Xem trạng thái | `.\regravity.exe status` |
| Khôi phục bản gốc | `.\regravity.exe restore` |
| Xem mọi tùy chọn | `.\regravity.exe --help` |

Để biên dịch từ mã nguồn trên Windows, cần PowerShell 7, Go 1.26.3+ và Wails v2 CLI. Chạy `.\build.ps1 -Target gui` trong thư mục dự án. Tệp kết quả nằm ở `build/bin/regravity.exe`.

Nếu công cụ báo thiếu WebView2 khi khởi động, hãy cài [Microsoft Edge WebView2 Runtime](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download) rồi mở lại công cụ.

## Giấy phép

Regravity là công cụ độc lập do cộng đồng duy trì, không trực thuộc hay được Google hoặc Antigravity xác nhận. Google Antigravity và các nhãn hiệu liên quan thuộc về chủ sở hữu tương ứng. Công cụ dành cho mục đích cá nhân và học tập; hãy tuân thủ điều khoản sử dụng của phần mềm liên quan.

[Giấy phép MIT](./LICENSE)
