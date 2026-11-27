localStorage.setItem('autoCheckUpdates', 'off');
window.go = {main: {App: {
  GetVersion: async () => 'test',
  UpdateWindowReady: () => {},
  GetStatus: async () => ({AppDir: 'D:/Antigravity', AppVersion: 'test', CurrentLanguage: 'en'}),
  GetLocales: async () => [
    {code: 'en', name: 'English', coverage: 100},
    {code: 'vi', name: 'Tiếng Việt', coverage: 100},
    {code: 'cs', name: 'Čeština', coverage: 100},
    {code: 'hu', name: 'Magyar', coverage: 100},
    {code: 'bg', name: 'Български', coverage: 100},
  ],
}}};

verify(async () => {
  function selectLanguage(code) {
    const control = document.getElementById('setupLanguage');
    control.value = code;
    control.dispatchEvent(new Event('change'));
  }
  await settle(220);
  const byId = id => document.getElementById(id);
  byId('localeChips').querySelectorAll('button')[1].click();
  byId('sessionMetrics').click();
  byId('effortSlider').click();
  selectLanguage('en');
  check('English document', document.documentElement.lang, 'en');
  check('English action', byId('btnPrimaryAction').textContent, 'Apply changes');
  check('independent pending language', byId('pendingLanguage').textContent, 'Switch to Tiếng Việt');
  check('feature selection survives', byId('sessionMetrics').checked, true);
  check('English saved', localStorage.getItem('setupLanguage'), 'en');
  check('updates unaffected', localStorage.getItem('autoCheckUpdates'), 'off');
  byId('btnAbout').click();
  check('About dialog translated', byId('btnOpenLogs').textContent, 'Activity log');
  byId('btnCloseAbout').click();
  for (const option of byId('setupLanguage').options) {
    selectLanguage(option.value);
    check('window direction', document.documentElement.dir, option.value === 'ar' ? 'rtl' : 'ltr');
    if (option.value === 'ar') check('installation path direction', byId('appPathDisplay').dir, 'ltr');
    check('menu keeps the tool version', byId('btnAbout').querySelector('#menuVersion').textContent, 'vtest');
    check('document language', document.documentElement.lang, option.value);
    check('saved language', localStorage.getItem('setupLanguage'), option.value);
    check('pending selection survives', byId('pendingLanguage').textContent.includes('Tiếng Việt'), true);
    if (option.value !== 'zh-CN' && option.value !== 'en') {
      for (const [id, key] of [['sessionMetrics', 'Session metrics'], ['effortSlider', 'Thinking effort slider']]) {
        const control = byId(id);
        check('localized switch ' + key, control.labels[0].querySelector('[data-en]').textContent, setupMessages[option.value][key]);
        check('switch stays on ' + key, control.checked, true);
      }
      check('translated action', byId('btnPrimaryAction').textContent, setupMessages[option.value]['Apply changes']);
    }
  }
  for (const [code, name] of [['cs', 'Čeština'], ['hu', 'Magyar'], ['bg', 'Български']]) {
    selectLanguage(code);
    const chip = [...byId('localeChips').querySelectorAll('button')].find(el => el.textContent === name);
    chip.click();
    check('pending Antigravity language ' + code, byId('pendingLanguage').textContent.includes(name), true);
  }
  [...byId('localeChips').querySelectorAll('button')].find(el => el.textContent === 'Tiếng Việt').click();
  check('all setup languages', byId('setupLanguage').options.length, 19);
  selectLanguage('zh-CN');
  check('Chinese document', document.documentElement.lang, 'zh-CN');
  check('Chinese action', byId('btnPrimaryAction').textContent, '应用更改');
  check('pending language retained', byId('pendingLanguage').textContent, '将切换为 Tiếng Việt');
});
