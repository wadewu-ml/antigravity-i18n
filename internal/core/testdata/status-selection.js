const detected = 'C:/Detected/Antigravity';
const selected = 'D:/Portable/Antigravity';
const status = (dir, version) => ({
  AppDir: dir, AppVersion: version, CurrentLanguage: 'en',
  SessionMetrics: false, EffortSlider: false, CleanBackupStatus: 'valid',
});
let finishDetection;
let chooseDirectory;
window.go = {main: {App: {
  GetVersion: async () => 'test',
  UpdateWindowReady: () => {},
  GetStatus: () => new Promise(resolve => { finishDetection = resolve; }),
  GetLocales: async () => [{code:'en', name:'English', coverage:100}],
  SelectAppDir: () => new Promise(resolve => { chooseDirectory = resolve; }),
  SetAppDir: async path => status(path, '2.35.1'),
}}};

verify(async () => {
  const path = document.getElementById('appPathDisplay');
  const picker = document.getElementById('btnChangePath');
  await settle(200);
  check('initial detection is pending', typeof finishDetection, 'function');
  picker.click();
  check('the picker disables competing path changes', picker.disabled, true);
  chooseDirectory(selected);
  await settle();
  check('the window shows the selected path', path.textContent, selected);
  finishDetection(status(detected, '2.34.0'));
  await settle();
  check('the late detection does not replace the path', path.textContent, selected);
  check('the late detection does not replace the version', document.getElementById('appVersion').textContent, 'v2.35.1');
});
