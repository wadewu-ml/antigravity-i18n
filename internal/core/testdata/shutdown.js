const calls = [];
const installed = {CurrentLanguage: 'en', SessionMetrics: false, EffortSlider: false};
const running = new Error('antigravity_running: Antigravity is running. Exit it and retry');
let nextError;

function failIfRequested() {
  const err = nextError;
  nextError = null;
  if (err) throw err;
}

window.go = {main: {App: {
  GetVersion: async () => 'test',
  GetStatus: async () => ({
    AppDir: 'C:/Antigravity', AppVersion: '2.33.0', ...installed,
    LocaleName: installed.CurrentLanguage === 'zh-CN' ? '简体中文' : 'English',
    HasCleanBackup: true, CleanBackupStatus: 'valid',
  }),
  GetLocales: async () => [
    {code: 'en', name: 'English', coverage: 100},
    {code: 'zh-CN', name: '简体中文', coverage: 100},
  ],
  ApplyCustomization: async (locale, metrics, slider, confirmed, force) => {
    calls.push(['apply', locale, metrics, slider, confirmed, force]);
    failIfRequested();
    Object.assign(installed, {CurrentLanguage: locale, SessionMetrics: metrics, EffortSlider: slider});
    return {LocaleName: 'English', Relaunched: true};
  },
  RestoreLocale: async (confirmed, force) => {
    calls.push(['restore', confirmed, force]);
    failIfRequested();
    Object.assign(installed, {CurrentLanguage: 'en', SessionMetrics: false, EffortSlider: false});
    return {Relaunched: true};
  },
}}};

verify(async () => {
  const byId = id => document.getElementById(id);
  const metrics = byId('sessionMetrics');
  const apply = byId('btnPrimaryAction');
  const message = byId('msgModal');
  const retry = byId('btnForceRetry');
  const restoreConfirm = byId('restoreConfirmModal');
  await settle(200);
  check('the installed language is the default selection',
    byId('localeChips').firstElementChild.getAttribute('aria-pressed'), 'true');
  metrics.click();
  check('English can be combined with metrics', metrics.checked && !metrics.disabled, true);
  check('unapplied preference is visible', byId('sessionMetricsDetail').textContent, '应用后开启');
  apply.click();
  check('one click starts a normal apply', calls.length === 1 && calls[0][5] === false, true);
  await settle();
  check('successful apply needs no dialog', message.open || restoreConfirm.open, false);

  nextError = running;
  apply.click();
  await settle();
  check('a running installation offers force retry', message.open && !retry.classList.contains('hidden'), true);
  check('the prompt tells users to exit Antigravity first',
    byId('msgTitle').textContent === 'Antigravity 正在运行' && byId('msgBody').textContent.includes('手动退出'), true);
  check('the log keeps the readable running-app notice',
    byId('logContent').textContent.includes('[提示] Antigravity is running'), true);
  check('retry dialog has no checkbox', message.querySelector('input') === null, true);
  check('cancel is focused by default', document.activeElement === byId('btnMsgOk'), true);
  check('failed apply preserves the preference', metrics.checked, true);
  byId('btnMsgOk').click();
  check('cancel does not retry', !message.open && calls.length === 2, true);

  nextError = running;
  apply.click();
  await settle();
  retry.click();
  await settle();
  check('force retry finishes in a single click', message.open, false);

  nextError = new Error('closing Antigravity: simulated process query failure');
  apply.click();
  await settle();
  check('other failures only show diagnostics', message.open && retry.classList.contains('hidden'), true);
  byId('btnMsgLogs').click();
  check('failure opens the exact error log', !byId('logDrawer').classList.contains('hidden') &&
    byId('logContent').textContent.includes('[错误] closing Antigravity: simulated process query failure'), true);
  byId('btnCloseLogs').click();

  byId('btnRestore').click();
  check('explicit restore has a simple confirmation', restoreConfirm.open && restoreConfirm.querySelector('input') === null, true);
  byId('btnCancelRestore').click();
  metrics.click();
  nextError = running;
  apply.click();
  check('turning off metrics starts a normal restore directly', !restoreConfirm.open && calls[5][0] === 'restore' && calls[5][2] === false, true);
  await settle();
  retry.click();
  await settle();
  check('backend receives force only after a retry click', JSON.stringify(calls), JSON.stringify([
    ['apply', 'en', true, false, true, false],
    ['apply', 'en', true, false, true, false],
    ['apply', 'en', true, false, true, false],
    ['apply', 'en', true, false, true, true],
    ['apply', 'en', true, false, true, false],
    ['restore', true, false],
    ['restore', true, true],
  ]));

  const slider = byId('effortSlider');
  check('slider starts off', slider.checked, false);
  slider.click();
  check('pending slider change', byId('effortSliderDetail').textContent, '应用后开启');
  apply.click();
  await settle();
  check('English with only the slider', installed.EffortSlider && !installed.SessionMetrics, true);
  check('successful slider apply updates the status', byId('statusDesc').textContent.includes('思考强度滑块已开启'), true);
  metrics.click();
  apply.click();
  await settle();
  slider.click();
  apply.click();
  await settle();
  check('disabling the slider keeps metrics', !installed.EffortSlider && installed.SessionMetrics, true);
  const picker = byId('localeChips');
  const pending = byId('pendingLanguage');
  const current = byId('statusLabel');
  picker.lastElementChild.click();
  check('selection previews the change', pending.textContent, '将切换为 简体中文');
  check('the installed language stays visible', current.textContent, '当前语言：官方英文');
  picker.firstElementChild.click();
  check('reselecting the installed language clears the preview', pending.textContent, '');
  picker.lastElementChild.click();
  apply.click();
  await settle();
  check('successful apply refreshes the installed language', current.textContent, '当前语言：简体中文');
  check('applied changes no longer show a preview', pending.textContent, '');
});
