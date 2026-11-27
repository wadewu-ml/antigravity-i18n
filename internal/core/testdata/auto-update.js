const hour = 3600000;
const start = Date.now();
const cached = { current: '4.2.1', available: 'v4.2.2', nextCheck: start + 24 * hour };
let calls = 0, installs = 0, finish, cancelInstall, onProgress;
window.runtime = { EventsOn(event, handler) { if (event === 'progress') onProgress = handler; } };
if (scenario === 'cached' || scenario === 'upgraded') localStorage.setItem('updateCheck', JSON.stringify(cached));
if (scenario === 'disabled') localStorage.setItem('autoCheckUpdates', 'off');
window.go = { main: { App: {
  GetVersion: async () => scenario === 'upgraded' ? '4.2.2' : '4.2.1',
  UpdateWindowReady() {},
  GetStatus: async () => scenario === 'installing'
    ? {AppDir:'C:/Antigravity', AppVersion:'2.35.2', CurrentLanguage:'zh-CN', LocaleName:'简体中文', SessionMetrics:true, EffortSlider:true}
    : {AppDir:'C:/Antigravity', AppVersion:'2.35.1', CurrentLanguage:'en'},
  GetLocales: async () => scenario === 'installing'
    ? Array.from(document.querySelectorAll('#setupLanguage option'), option => ({code:option.value, name:option.textContent, coverage:100}))
    : [{code:'en', name:'English', coverage:100}],
  CheckForUpdate: () => { calls++; return new Promise(resolve => { finish = resolve; }); },
  InstallUpdate: async () => {
    installs++;
    if (scenario === 'installing') await new Promise((_, reject) => { cancelInstall = () => reject('已取消更新。'); });
  },
  CancelUpdate: () => cancelInstall(),
} } };

verify(async () => {
  const byId = id => document.getElementById(id);
  const dot = () => byId('moreActions').classList.contains('has-update');
  const snapshot = () => JSON.parse(localStorage.getItem('updateCheck'));
  const manual = () => byId('btnCheckUpdate').click();
  const toggle = () => byId('autoCheckUpdates').click();
  await settle(170);

  if (scenario === 'installing') {
    const root = document.documentElement;
    const overflow = () => root.scrollHeight - root.clientHeight;
    const visible = element => element.getClientRects().length > 0 && element.getBoundingClientRect().bottom <= root.clientHeight;
    const actionsTop = byId('actionRow').getBoundingClientRect().top;
    check(`setup fits before updating (${root.clientWidth}×${root.clientHeight})`, overflow(), 0);
    finish({release:{available:true, version:'v4.2.2', notes:'Release notes'}});
    await settle();
    manual();
    await settle();
    byId('btnInstallUpdate').click();
    await settle();
    onProgress({stage:'正在下载更新', detail:'2.2 / 14.6 MB (15%)'});
    check('one installation started', installs, 1);
    check('download does not increase page height', overflow(), 0);
    check('download uses the existing footer space', byId('actionRow').getBoundingClientRect().top, actionsTop);
    check('download detail stays visible', visible(byId('progressDetail')), true);
    check('cancel stays visible', visible(byId('btnCancelDownload')), true);
    const longDetail = 'C:/Antigravity/' + 'long-file-name-'.repeat(20) + '.exe';
    onProgress({stage:'正在下载更新', detail:longDetail});
    check('long details do not enlarge the window content', overflow(), 0);
    check('full progress details remain available', byId('progressDetail').title, longDetail);
    byId('btnCancelDownload').click();
    await settle();
    check('setup fits after cancel', overflow(), 0);
    check('normal primary action is visible again', visible(byId('btnPrimaryAction')), true);
    return;
  }
  if (scenario === 'cached' || scenario === 'upgraded') {
    check('fresh cache skips the network', calls, 0);
    check('badge belongs to this installed version', dot(), scenario === 'cached');
    if (scenario === 'cached') {
      manual();
      await settle();
      check('manual check bypasses the daily interval', calls, 1);
      finish({release:{available:false, version:'v4.2.1'}});
      await settle();
      check('latest check clears stale badge', dot(), false);
    }
    return;
  }
  if (scenario === 'disabled') {
    check('disabled setting survives startup', byId('autoCheckUpdates').checked, false);
    check('disabled skips network', calls, 0);
    manual();
    await settle();
    check('manual check still works', calls, 1);
    return;
  }
  check('opens with one background check', calls, 1);
  check('directory picker stays usable during background check', byId('btnChangePath').disabled, false);
  if (scenario === 'available') {
    finish({release:{available:true, version:'v4.2.2', notes:'Release notes'}});
    await settle();
    check('badge appears', dot(), true);
    check('menu names the version', byId('btnCheckUpdate').textContent, '发现新版本 v4.2.2');
    check('startup stays quiet', document.querySelector('dialog[open]'), null);
    check('next check is 24 hours later', Math.abs(snapshot().nextCheck - Date.now() - 24 * hour) < 1000, true);
    manual();
    await settle();
    check('manual opens release notes', byId('updateModal').open, true);
    check('notes reuse the verified response', calls, 1);
    check('no automatic install', installs, 0);
    byId('btnLater').click();
    toggle();
    check('turning off hides the dot', dot(), false);
    check('off is saved', localStorage.getItem('autoCheckUpdates'), 'off');
    toggle();
    check('turning on restores the cached dot', dot(), true);
    check('preferences keep daily interval', calls, 1);
  } else {
    const limited = scenario === 'limited';
    finish({error:'暂时无法检查更新', retryAt:limited ? start + 2 * hour : 0});
    await settle();
    check('no error dialog at startup', document.querySelector('dialog[open]'), null);
    check('failed check schedules retry', Math.abs(snapshot().nextCheck - start - (limited ? 2 : 1) * hour) < 1000, true);
    toggle();
    toggle();
    check('retry deadline prevents automatic check', calls, 1);
    manual();
    await settle();
    check('only server throttling limits manual retries', calls, limited ? 1 : 2);
  }
});
