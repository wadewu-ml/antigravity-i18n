// Frontend controller for antigravity-i18n.
//
// Field naming follows what Wails actually puts on the wire: types declared
// with json tags (LocaleInfo) arrive lower camel case, while plain structs
// (Status, ApplyResult, RestoreResult) arrive with their Go field names.
(function () {
  'use strict';

  var currentStatus = null;
  var locales = [];
  var selectedLocale = '';
  var isOperating = false;
  var operationLogs = [];

  var elAppVersionBadge = document.getElementById('appVersionBadge');
  var elAppVersion = document.getElementById('appVersion');
  var elAppPathDisplay = document.getElementById('appPathDisplay');
  var elBtnChangePath = document.getElementById('btnChangePath');
  var elBtnAbout = document.getElementById('btnAbout');
  var elStatusPill = document.getElementById('statusPill');
  var elStatusLabel = document.getElementById('statusLabel');
  var elStatusDesc = document.getElementById('statusDesc');
  var elProgressBox = document.getElementById('progressBox');
  var elProgressStage = document.getElementById('progressStage');
  var elProgressDetail = document.getElementById('progressDetail');
  var elBtnPrimaryAction = document.getElementById('btnPrimaryAction');
  var elBtnSecondaryAction = document.getElementById('btnSecondaryAction');
  var elBtnRestore = document.getElementById('btnRestore');
  var elLocaleChips = document.getElementById('localeChips');
  var elLocaleNote = document.getElementById('localeNote');
  var elSessionMetrics = document.getElementById('sessionMetrics');
  var elSessionMetricsDetail = document.getElementById('sessionMetricsDetail');
  var elBtnToggleLogs = document.getElementById('btnToggleLogs');
  var elLogDrawer = document.getElementById('logDrawer');
  var elLogContent = document.getElementById('logContent');
  var elBtnCopyLogs = document.getElementById('btnCopyLogs');
  var elBtnCloseLogs = document.getElementById('btnCloseLogs');
  var elAboutModal = document.getElementById('aboutModal');
  var elBtnCloseAbout = document.getElementById('btnCloseAbout');
  var elApplyConfirmModal = document.getElementById('applyConfirmModal');
  var elApplyConfirmTitle = document.getElementById('applyConfirmTitle');
  var elApplyConfirmBody = document.getElementById('applyConfirmBody');
  var elForceShutdown = document.getElementById('forceShutdown');
  var elBtnCancelApply = document.getElementById('btnCancelApply');
  var elBtnConfirmApply = document.getElementById('btnConfirmApply');
  var elMsgModal = document.getElementById('msgModal');
  var elMsgTitle = document.getElementById('msgTitle');
  var elMsgBody = document.getElementById('msgBody');
  var elBtnMsgLogs = document.getElementById('btnMsgLogs');
  var elBtnMsgOk = document.getElementById('btnMsgOk');

  function backend() {
    return window.go && window.go.main && window.go.main.App ? window.go.main.App : null;
  }

  function errorText(err) {
    if (!err) return '未知错误';
    if (typeof err === 'string') return err;
    return err.message || String(err);
  }

  function appendLog(line) {
    operationLogs.push(line);
    elLogContent.textContent = operationLogs.join('\n');
    elLogContent.scrollTop = elLogContent.scrollHeight;
  }

  function showMessage(title, body, showLogs) {
    elMsgTitle.textContent = title;
    elMsgBody.textContent = body;
    elBtnMsgLogs.classList.toggle('hidden', !showLogs);
    elMsgModal.classList.remove('hidden');
  }

  function setOperatingState(operating, stage, detail) {
    isOperating = operating;
    if (operating) {
      elProgressBox.classList.remove('hidden');
      elProgressStage.textContent = stage || '正在处理...';
      elProgressDetail.textContent = detail || '请稍候...';
    } else {
      elProgressBox.classList.add('hidden');
    }
    [elBtnPrimaryAction, elBtnSecondaryAction, elBtnRestore, elBtnChangePath].forEach(function (btn) {
      btn.disabled = operating;
    });
    Array.prototype.forEach.call(elLocaleChips.querySelectorAll('.chip'), function (chip) {
      chip.disabled = operating;
    });
    renderPrimaryLabel();
  }

  function isCustomized(st) {
    return !!(st && (st.SessionMetrics || st.CurrentLanguage && st.CurrentLanguage !== 'en'));
  }

  function setStatus(st, selectInstalled) {
    currentStatus = st;
    if (selectInstalled || !selectedLocale) {
      selectedLocale = st.CurrentLanguage;
      elSessionMetrics.checked = Boolean(st.SessionMetrics);
    }
    renderStatus(st);
    renderChips();
  }

  function refreshStatus() {
    var app = backend();
    if (!app) {
      elStatusPill.className = 'status-pill danger';
      elStatusLabel.textContent = '无法连接';
      elStatusDesc.textContent = '窗口运行库没有就绪，请重新启动本工具。';
      return Promise.resolve();
    }
    return app.GetStatus().then(function (st) {
      setStatus(st, false);
    }).catch(function (err) {
      currentStatus = null;
      renderPrimaryLabel();
      elStatusPill.className = 'status-pill danger';
      elStatusLabel.textContent = '读取失败';
      elStatusDesc.textContent = errorText(err);
      appendLog('[错误] 读取状态失败: ' + errorText(err));
    });
  }

  function renderStatus(st) {
    if (!st) return;

    elAppPathDisplay.textContent = st.AppDir || '未检测到 Antigravity 安装位置';
    elAppPathDisplay.title = st.AppDir || '';
    elAppVersion.textContent = st.AppVersion ? 'v' + st.AppVersion : '版本未知';

    elStatusPill.className = 'status-pill';
    elBtnSecondaryAction.classList.add('hidden');
    elBtnRestore.classList.add('hidden');

    if (isCustomized(st)) {
      elStatusPill.classList.add('success');
      elStatusLabel.textContent = st.CurrentLanguage === 'en' ? '官方英文' : '已本地化 · ' + st.CurrentLanguage;
      elStatusDesc.textContent = '当前界面语言为 ' + (st.LocaleName || st.CurrentLanguage) + '。' +
        (st.SessionMetrics ? '会话指标已开启。' : '');
      elBtnSecondaryAction.classList.remove('hidden');
      elBtnSecondaryAction.textContent = '打开 Antigravity';
      elBtnRestore.classList.remove('hidden');
    } else if (st.HasCleanBackup) {
      elStatusPill.classList.add('default');
      elStatusLabel.textContent = '官方英文';
      elStatusDesc.textContent = '干净备份可用。';
      elBtnSecondaryAction.classList.remove('hidden');
      elBtnSecondaryAction.textContent = '打开 Antigravity';
    } else {
      elStatusPill.classList.add('warning');
      elStatusLabel.textContent = '官方英文';
      elStatusDesc.textContent = '应用前会自动创建干净备份。';
    }

    if (st.CleanBackupStatus && st.CleanBackupStatus !== 'valid') {
      elStatusDesc.textContent += ' 干净备份状态：' + st.CleanBackupStatus + '。';
    }
    if (st.StateIsStale) {
      elStatusDesc.textContent += ' 状态文件与实际内容不一致，以 app.asar 的实际内容为准。';
    }
  }
  function selectedLabel() {
    for (var i = 0; i < locales.length; i++) {
      if (locales[i].code === selectedLocale) return locales[i].name;
    }
    return selectedLocale;
  }

  function renderPrimaryLabel() {
    elSessionMetrics.disabled = isOperating || !currentStatus;
    if (!currentStatus) {
      elBtnPrimaryAction.textContent = '重新检测';
      elBtnPrimaryAction.disabled = isOperating;
      return;
    }
    var metricsChanged = Boolean(currentStatus.SessionMetrics) !== elSessionMetrics.checked;
    var hasChanges = currentStatus.CurrentLanguage !== selectedLocale || metricsChanged;
    var canReapply = !hasChanges && isCustomized(currentStatus);
    elBtnPrimaryAction.textContent = canReapply ? '重新应用' : '应用更改';
    elBtnPrimaryAction.disabled = isOperating || (!hasChanges && !canReapply);
    elSessionMetricsDetail.textContent = metricsChanged
      ? (elSessionMetrics.checked ? '应用后开启' : '应用后关闭')
      : '在 Antigravity 的模型选择器旁显示上下文用量。';
  }

  function renderChips() {
    elLocaleChips.innerHTML = '';
    if (!locales.length) {
      elLocaleNote.textContent = '没有可用的语言包';
      return;
    }
    locales.forEach(function (loc) {
      var chip = document.createElement('button');
      chip.type = 'button';
      chip.className = 'chip' + (loc.coverage < 100 ? ' partial' : '');
      chip.setAttribute('aria-pressed', String(loc.code === selectedLocale));
      chip.disabled = isOperating;
      chip.title = loc.englishName;

      var name = document.createElement('span');
      name.textContent = loc.name;
      var cov = document.createElement('span');
      cov.className = 'cov';
      cov.textContent = loc.coverage + '%';
      chip.appendChild(name);
      chip.appendChild(cov);

      chip.addEventListener('click', function () {
        selectedLocale = loc.code;
        renderChips();
      });
      elLocaleChips.appendChild(chip);
    });
    elLocaleNote.textContent = '';
    renderPrimaryLabel();
  }

  function loadLocales() {
    var app = backend();
    if (!app) return Promise.resolve();
    return app.GetLocales().then(function (list) {
      locales = list || [];
      renderChips();
    }).catch(function (err) {
      elLocaleNote.textContent = '读取语言列表失败';
      appendLog('[错误] 读取语言列表失败: ' + errorText(err));
    });
  }

  function doApply() {
    if (isOperating) return;
    var app = backend();
    if (!app) return;
    var code = selectedLocale || 'zh-CN';

    var metrics = elSessionMetrics.checked;
    var changes = [];
    if (currentStatus.CurrentLanguage !== code) changes.push('界面语言设为 ' + selectedLabel());
    if (Boolean(currentStatus.SessionMetrics) !== metrics) changes.push((metrics ? '开启' : '关闭') + '会话指标');
    var title = changes.length ? '应用更改？' : '重新应用？';
    var description = changes.length ? changes.join('，') + '。' : '重新写入当前的语言和会话指标设置。';
    var apply = (code === 'en' && !metrics) ? doRestore : function (force) { runApply(app, code, metrics, force); };
    confirmAction(title, description, apply);
  }

  function confirmRestore() {
    confirmAction('还原官方版本？', 'Antigravity 将恢复官方英文界面，会话指标将关闭。', doRestore);
  }

  function confirmAction(title, description, onConfirm) {
    elApplyConfirmTitle.textContent = title;
    elApplyConfirmBody.textContent = description;
    elForceShutdown.checked = false;
    elBtnConfirmApply.onclick = function () {
      elApplyConfirmModal.classList.add('hidden');
      onConfirm(elForceShutdown.checked);
    };
    elApplyConfirmModal.classList.remove('hidden');
  }

  function runApply(app, code, metrics, force) {
    setOperatingState(true, '正在应用更改', '正在准备补丁...');
    appendLog('[操作] 应用 ' + code + ' · 会话指标' + (metrics ? '开启' : '关闭') + (force ? ' · 强制退出' : ''));

    app.ApplyCustomization(code, metrics, true, force).then(function (res) {
      var name = res && res.LocaleName ? res.LocaleName : code;
      appendLog('[完成] 已应用 ' + name);
      if (!res || !res.Relaunched) showMessage('已应用', '请手动打开 Antigravity。');
    }).catch(function (err) {
      var msg = errorText(err);
      appendLog('[错误] ' + msg);
      showMessage('应用失败', msg, true);
    }).then(function () {
      setOperatingState(false);
      return refreshStatus();
    });
  }

  function doRestore(force) {
    if (isOperating) return;
    var app = backend();
    if (!app) return;

    setOperatingState(true, '正在还原官方版本', '正在恢复干净备份并重算校验...');
    appendLog('[操作] 还原为官方英文');

    app.RestoreLocale(true, force).then(function (res) {
      selectedLocale = 'en';
      elSessionMetrics.checked = false;
      appendLog('[完成] 已还原为官方英文');
      if (res && !res.AlreadyEnglish && !res.Relaunched) showMessage('已还原', '请手动打开 Antigravity。');
    }).catch(function (err) {
      var msg = errorText(err);
      appendLog('[错误] ' + msg);
      showMessage('还原失败', msg, true);
    }).then(function () {
      setOperatingState(false);
      return refreshStatus();
    });
  }

  function launchApp() {
    var app = backend();
    if (!app) return;
    app.LaunchAntigravity().then(function () {
      appendLog('[操作] 已启动 Antigravity');
    }).catch(function (err) {
      showMessage('启动失败', errorText(err));
    });
  }

  function handlePrimaryAction() {
    if (isOperating) return;
    if (!currentStatus) {
      refreshStatus();
      return;
    }
    if (!currentStatus.AppDir) {
      handleChangePath();
      return;
    }
    doApply();
  }

  function handleChangePath() {
    var app = backend();
    if (!app) return;
    app.SelectAppDir().then(function (selected) {
      if (!selected) return null;
      setOperatingState(true, '正在切换安装位置', selected);
      return app.SetAppDir(selected).then(function (st) {
        setStatus(st, true);
      });
    }).catch(function (err) {
      var msg = errorText(err);
      appendLog('[错误] ' + msg);
      showMessage('切换位置失败', msg);
    }).then(function () {
      setOperatingState(false);
    });
  }

  elBtnPrimaryAction.addEventListener('click', handlePrimaryAction);
  elBtnSecondaryAction.addEventListener('click', launchApp);
  elBtnChangePath.addEventListener('click', handleChangePath);
  elBtnRestore.addEventListener('click', confirmRestore);
  elSessionMetrics.addEventListener('change', renderPrimaryLabel);
  elBtnCancelApply.addEventListener('click', function () {
    elApplyConfirmModal.classList.add('hidden');
    elBtnConfirmApply.onclick = null;
  });

  elBtnAbout.addEventListener('click', function () {
    elAboutModal.classList.remove('hidden');
  });
  elBtnCloseAbout.addEventListener('click', function () {
    elAboutModal.classList.add('hidden');
  });

  elBtnMsgOk.addEventListener('click', function () {
    elMsgModal.classList.add('hidden');
  });
  elBtnMsgLogs.addEventListener('click', function () {
    elMsgModal.classList.add('hidden');
    elLogDrawer.classList.remove('hidden');
  });

  elBtnToggleLogs.addEventListener('click', function () {
    elLogDrawer.classList.toggle('hidden');
  });
  elBtnCloseLogs.addEventListener('click', function () {
    elLogDrawer.classList.add('hidden');
  });
  elBtnCopyLogs.addEventListener('click', function () {
    var text = operationLogs.join('\n');
    // navigator.clipboard is only exposed in a secure context and can be
    // absent in older WebView2 builds. Wails ships its own clipboard call, so
    // fall back to it instead of leaving the button doing nothing.
    if (navigator.clipboard && navigator.clipboard.writeText) {
      navigator.clipboard.writeText(text).then(function () {
        showMessage('已复制', '运行日志已复制到剪贴板。');
      }).catch(function () {
        copyLogsWithRuntime(text);
      });
      return;
    }
    copyLogsWithRuntime(text);
  });

  function copyLogsWithRuntime(text) {
    var rt = window.runtime;
    if (rt && typeof rt.ClipboardSetText === 'function') {
      try {
        var result = rt.ClipboardSetText(text);
        if (result && typeof result.then === 'function') {
          result.then(function (ok) {
            showMessage(ok === false ? '复制失败' : '已复制',
              ok === false ? '无法访问剪贴板，请手动选择日志内容。' : '运行日志已复制到剪贴板。');
          }).catch(function () {
            showMessage('复制失败', '无法访问剪贴板，请手动选择日志内容。');
          });
          return;
        }
        showMessage('已复制', '运行日志已复制到剪贴板。');
        return;
      } catch (err) {
        // Fall through to the manual-selection hint below.
      }
    }
    showMessage('复制失败', '无法访问剪贴板，请手动选择日志内容。');
  }

  function setupWailsEvents() {
    if (window.runtime && window.runtime.EventsOn) {
      window.runtime.EventsOn('progress', function (data) {
        if (!data || !isOperating) return;
        if (data.stage) elProgressStage.textContent = data.stage;
        if (data.detail) elProgressDetail.textContent = data.detail;
        appendLog('[进度] ' + (data.stage || '') + (data.detail ? ': ' + data.detail : ''));
      });
    }
  }

  window.addEventListener('DOMContentLoaded', function () {
    setupWailsEvents();
    // Wait briefly for the Wails IPC layer to attach window.go.
    setTimeout(function () {
      var app = backend();
      if (app) {
        app.GetVersion().then(function (v) {
          if (v) elAppVersionBadge.textContent = 'v' + v;
        }).catch(function () {});
      }
      refreshStatus().then(loadLocales);
    }, 120);
  });
})();
