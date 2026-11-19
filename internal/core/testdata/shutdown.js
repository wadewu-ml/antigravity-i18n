const calls = [];
const installed = {CurrentLanguage: 'en', SessionMetrics: false};
let rejectForcedApply = true;
window.go = {main: {App: {
  GetVersion: async () => 'test',
  GetStatus: async () => ({
    AppDir: 'C:/Antigravity', AppVersion: '2.33.0', ...installed,
    HasCleanBackup: true, CleanBackupStatus: 'valid',
  }),
  GetLocales: async () => [
    {code: 'en', name: 'English', coverage: 100},
    {code: 'zh-CN', name: '简体中文', coverage: 100},
  ],
  ApplyCustomization: async (locale, metrics, confirmed, force) => {
    calls.push(['apply', locale, metrics, confirmed, force]);
    if (force && rejectForcedApply) {
      rejectForcedApply = false;
      throw new Error('closing Antigravity: simulated shutdown failure');
    }
    Object.assign(installed, {CurrentLanguage: locale, SessionMetrics: metrics});
    return {LocaleName: 'English', Relaunched: true};
  },
  RestoreLocale: async (confirmed, force) => {
    calls.push(['restore', confirmed, force]);
    Object.assign(installed, {CurrentLanguage: 'en', SessionMetrics: false});
    return {Relaunched: true};
  },
}}};

verify(async () => {
  const byId = id => document.getElementById(id);
  const force = byId('forceShutdown');
  const confirm = byId('btnConfirmApply');
  const metrics = byId('sessionMetrics');
  await settle(200);
  check('the installed language is the default selection',
    byId('localeChips').firstElementChild.getAttribute('aria-pressed'), 'true');
  metrics.click();
  check('English can be combined with metrics', metrics.checked && !metrics.disabled, true);
  check('unapplied preference is visible', byId('sessionMetricsDetail').textContent, '应用后开启');
  byId('btnPrimaryAction').click();
  check('confirmation describes only the changed feature', byId('applyConfirmBody').textContent, '开启会话指标。');
  check('default force selection', force.checked, false);
  check('timeout notice remains visible', byId('applyConfirmModal').querySelector('.modal-sub').textContent.includes('30 秒'), true);
  force.click();
  confirm.click();
  await settle();
  check('failure names the stage', byId('msgBody').textContent.includes('closing Antigravity'), true);
  check('failed apply preserves the preference', metrics.checked, true);
  byId('btnMsgLogs').click();
  check('failure opens the exact error log', !byId('logDrawer').classList.contains('hidden') &&
    byId('logContent').textContent.includes('[错误] closing Antigravity'), true);
  byId('btnCloseLogs').click();
  byId('btnPrimaryAction').click();
  check('force is reset for each operation', force.checked, false);
  confirm.click();
  await settle();
  check('successful apply needs no additional dismissal', byId('msgModal').classList.contains('hidden'), true);

  metrics.click();
  byId('btnPrimaryAction').click();
  check('turning off the indicator is an independent action', byId('applyConfirmBody').textContent, '关闭会话指标。');
  force.click();
  confirm.click();
  await settle();
  check('backend receives explicit consent', JSON.stringify(calls), JSON.stringify([
    ['apply', 'en', true, true, true],
    ['apply', 'en', true, true, false],
    ['restore', true, true],
  ]));
});
