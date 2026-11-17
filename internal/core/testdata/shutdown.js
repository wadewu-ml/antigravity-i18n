const calls = [];
window.go = {main: {App: {
  GetVersion: async () => 'test',
  GetStatus: async () => ({
    AppDir: 'C:/Antigravity', AppVersion: '2.32.1', CurrentLanguage: 'zh-CN',
    HasCleanBackup: true, CleanBackupStatus: 'valid',
  }),
  GetLocales: async () => [
    {code: 'en', name: 'English', coverage: 100},
    {code: 'zh-CN', name: '简体中文', coverage: 100},
  ],
  ApplyLocale: async (locale, confirmed, force) => {
    calls.push(['apply', locale, confirmed, force]);
    return {LocaleName: '简体中文', Relaunched: true};
  },
  RestoreLocale: async (confirmed, force) => {
    calls.push(['restore', confirmed, force]);
    return {Relaunched: true};
  },
}}};

verify(async () => {
  const byId = id => document.getElementById(id);
  const force = byId('forceShutdown');
  const confirm = byId('btnConfirmApply');
  await settle(200);
  byId('btnPrimaryAction').click();
  check('default force selection', force.checked, false);
  check('timeout notice remains visible', byId('applyConfirmModal').querySelector('.modal-sub').textContent.includes('取消操作'), true);
  confirm.click();
  await settle();
  byId('btnMsgOk').click();
  byId('btnRestore').click();
  force.click();
  confirm.click();
  await settle();
  byId('btnMsgOk').click();
  byId('btnRestore').click();
  check('force is reset for each operation', force.checked, false);
  check('backend receives explicit consent', JSON.stringify(calls), JSON.stringify([
    ['apply', 'zh-CN', true, false],
    ['restore', true, true],
  ]));
});
