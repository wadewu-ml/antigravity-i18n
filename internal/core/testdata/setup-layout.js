localStorage.setItem('autoCheckUpdates', 'off');
localStorage.setItem('setupLanguage', 'zh-CN');
let onProgress, cancelInstall;
window.runtime = {EventsOn(event, handler) { if (event === 'progress') onProgress = handler; }};
window.go = {main: {App: {
  GetVersion: async () => 'test',
  UpdateWindowReady: () => {},
  GetStatus: async () => ({
    AppDir: 'C:/Users/wade/AppData/Local/Programs/antigravity',
    AppVersion: '2.36.0', CurrentLanguage: 'en', CleanBackupStatus: 'missing',
  }),
  GetLocales: async () => Array.from(document.querySelectorAll('#setupLanguage option'),
    option => ({code: option.value, name: option.textContent, coverage: 100})),
  CheckForUpdate: async () => ({release: {available: true, version: 'v99.0.0', notes: 'Test release'}}),
  InstallUpdate: () => new Promise((_, reject) => { cancelInstall = () => reject('已取消更新。'); }),
  CancelUpdate: () => cancelInstall(),
}}};

verify(async () => {
  await settle(200);
  check('all translation packs are displayed', document.querySelectorAll('.chip').length, 19);
  check('backup status adds a visible line', document.getElementById('statusDesc').textContent.includes('missing'), true);

  // Measure the rendered app at client-area sizes, excluding the Windows title
  // bar and borders. The other browser checks use a larger outer test window.
  const frame = document.createElement('iframe');
  frame.style.cssText = 'position:fixed;left:0;top:0;border:0';
  document.body.appendChild(frame);
  const doc = frame.contentDocument;
  doc.open();
  doc.write('<!DOCTYPE html><html><head>' + document.head.querySelector('style').outerHTML + '</head><body></body></html>');
  doc.close();
  function snapshot() {
    doc.documentElement.lang = document.documentElement.lang;
    doc.documentElement.dir = document.documentElement.dir;
    doc.body.innerHTML = document.querySelector('.window-container').outerHTML;
  }
  function fits(label, width, height, allowScroll = false, actionID = 'btnPrimaryAction') {
    frame.style.width = width + 'px';
    frame.style.height = height + 'px';
    const root = doc.documentElement;
    if (allowScroll) root.scrollTop = root.scrollHeight;
    else check(label + ' vertical overflow', root.scrollHeight - root.clientHeight, 0);
    check(label + ' horizontal overflow', root.scrollWidth - root.clientWidth, 0);
    const action = doc.getElementById(actionID).getBoundingClientRect();
    check(label + ' action visible', action.height > 0 && action.bottom <= height, true);
    check(label + ' content remains scrollable if needed', ['hidden', 'clip'].includes(frame.contentWindow.getComputedStyle(root).overflowY), false);
  }
  snapshot();
  fits('default window with backup status', 744, 540);
  fits('slightly shorter client area', 736, 532);
  fits('minimum window with backup status', 624, 480, true);
  document.querySelectorAll('.chip')[17].click();
  snapshot();
  fits('pending language at default size', 744, 540);
  fits('pending language at minimum size', 624, 480, true);

  document.getElementById('btnCheckUpdate').click();
  await settle();
  document.getElementById('btnInstallUpdate').click();
  await settle();
  onProgress({stage: '正在下载更新', detail: '2.2 / 14.6 MB (15%)'});
  snapshot();
  fits('update progress with backup status', 744, 540, false, 'btnCancelDownload');
  document.getElementById('btnCancelDownload').click();
  await settle();
  snapshot();
  fits('after cancelling update', 744, 540);
});
