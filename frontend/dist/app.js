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
  var elMoreActions = document.getElementById('moreActions');
  var elCheckUpdate = document.getElementById('btnCheckUpdate');
  var elUpdateStatus = document.getElementById('updateStatus');
  var elUpdateModal = document.getElementById('updateModal');
  var elCancelDownload = document.getElementById('btnCancelDownload');
  var elReleaseFallback = document.getElementById('btnReleaseFallback');
  var offeredVersion = '';
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
  var elPendingLanguage = document.getElementById('pendingLanguage');
  var elBtnSecondaryAction = document.getElementById('btnSecondaryAction');
  var elBtnRestore = document.getElementById('btnRestore');
  var elLocaleChips = document.getElementById('localeChips');
  var elLocaleNote = document.getElementById('localeNote');
  var elSessionMetrics = document.getElementById('sessionMetrics');
  var elEffortSlider = document.getElementById('effortSlider');
  var elEffortSliderDetail = document.getElementById('effortSliderDetail');
  var elSessionMetricsDetail = document.getElementById('sessionMetricsDetail');
  var elLogModal = document.getElementById('logModal');
  var elLogContent = document.getElementById('logContent');
  var elBtnCopyLogs = document.getElementById('btnCopyLogs');
  var elBtnCloseLogs = document.getElementById('btnCloseLogs');
  var elAboutModal = document.getElementById('aboutModal');
  var elBtnCloseAbout = document.getElementById('btnCloseAbout');
  var elRestoreConfirmModal = document.getElementById('restoreConfirmModal');
  var elMsgModal = document.getElementById('msgModal');
  var elMsgTitle = document.getElementById('msgTitle');
  var elMsgBody = document.getElementById('msgBody');
  var elBtnMsgLogs = document.getElementById('btnMsgLogs');
  var elBtnMsgOk = document.getElementById('btnMsgOk');
  var elBtnForceRetry = document.getElementById('btnForceRetry');

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

  function showMessage(title, body, showLogs, retry) {
    elReleaseFallback.classList.add('hidden');
    elMsgTitle.textContent = title;
    elMsgBody.textContent = body;
    elBtnMsgLogs.classList.toggle('hidden', !showLogs);
    elBtnMsgOk.textContent = retry ? '取消' : '知道了';
    elBtnMsgOk.className = retry ? 'btn btn-ghost' : 'btn btn-primary';
    elBtnForceRetry.classList.toggle('hidden', !retry);
    elBtnForceRetry.onclick = retry ? function () {
      elMsgModal.close();
      retry();
    } : null;
    elMsgModal.showModal();
    elBtnMsgOk.focus();
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
    [elBtnPrimaryAction, elBtnSecondaryAction, elBtnRestore, elBtnChangePath, elCheckUpdate].forEach(function (btn) {
      btn.disabled = operating;
    });
    Array.prototype.forEach.call(elLocaleChips.querySelectorAll('.chip'), function (chip) {
      chip.disabled = operating;
    });
    renderPrimaryLabel();
  }

  function isCustomized(st) {
    return !!(st && (st.SessionMetrics || st.EffortSlider || st.CurrentLanguage && st.CurrentLanguage !== 'en'));
  }

  function setStatus(st, selectInstalled) {
    currentStatus = st;
    if (selectInstalled || !selectedLocale) {
      selectedLocale = st.CurrentLanguage;
      elSessionMetrics.checked = Boolean(st.SessionMetrics);
      elEffortSlider.checked = Boolean(st.EffortSlider);
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

    var currentName = st.CurrentLanguage === 'en' ? '官方英文' : (st.LocaleName || st.CurrentLanguage);
    elStatusLabel.textContent = '当前语言：' + currentName;
    elStatusDesc.textContent = '';
    if (isCustomized(st)) {
      elStatusPill.classList.add('success');
      elBtnSecondaryAction.classList.remove('hidden');
      elBtnSecondaryAction.textContent = '打开 Antigravity';
      elBtnRestore.classList.remove('hidden');
    } else if (st.HasCleanBackup) {
      elStatusPill.classList.add('default');
      elBtnSecondaryAction.classList.remove('hidden');
      elBtnSecondaryAction.textContent = '打开 Antigravity';
    }

    if (st.CleanBackupStatus && st.CleanBackupStatus !== 'valid') {
      elStatusDesc.textContent += ' 干净备份状态：' + st.CleanBackupStatus + '。';
    }
    if (st.StateIsStale) {
      elStatusDesc.textContent += ' 状态文件与实际内容不一致，以 app.asar 的实际内容为准。';
    }
  }
  function renderPrimaryLabel() {
    var languageChanged = currentStatus && selectedLocale && currentStatus.CurrentLanguage !== selectedLocale;
    var pending = languageChanged && locales.find(function (loc) { return loc.code === selectedLocale; });
    elPendingLanguage.textContent = pending ? '将切换为 ' + pending.name : '';
    elEffortSlider.disabled = elSessionMetrics.disabled = isOperating || !currentStatus;
    elBtnPrimaryAction.title = currentStatus ? '应用后会重启 Antigravity' : '';
    if (!currentStatus) {
      elBtnPrimaryAction.textContent = '重新检测';
      elBtnPrimaryAction.disabled = isOperating;
      return;
    }
    var metricsChanged = Boolean(currentStatus.SessionMetrics) !== elSessionMetrics.checked;
    var sliderChanged = Boolean(currentStatus.EffortSlider) !== elEffortSlider.checked;
    var hasChanges = languageChanged || metricsChanged || sliderChanged;
    var canReapply = !hasChanges && isCustomized(currentStatus);
    elBtnPrimaryAction.textContent = canReapply ? '重新应用' : '应用更改';
    elBtnPrimaryAction.disabled = isOperating || (!hasChanges && !canReapply);
    elEffortSliderDetail.textContent = sliderChanged
      ? (elEffortSlider.checked ? '应用后开启' : '应用后关闭')
      : '';
    elSessionMetricsDetail.textContent = metricsChanged
      ? (elSessionMetrics.checked ? '应用后开启' : '应用后关闭')
      : '';
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
      chip.appendChild(name);
      if (loc.coverage < 100) {
        var cov = document.createElement('span');
        cov.className = 'cov';
        cov.textContent = loc.coverage + '%';
        chip.appendChild(cov);
      }

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
    var slider = elEffortSlider.checked;
    if (code === 'en' && !metrics && !slider) {
      doRestore(false);
      return;
    }
    runApply(app, code, metrics, slider, false);
  }

  function showOperationError(title, err, retry) {
    var msg = errorText(err);
    var runningPrefix = 'antigravity_running: ';
    if (retry && msg.startsWith(runningPrefix)) {
      appendLog('[提示] ' + msg.slice(runningPrefix.length));
      showMessage('Antigravity 正在运行', '请先手动退出 Antigravity 再试。强制退出会中断正在进行的会话，未保存的编辑可能丢失。', true, retry);
      return;
    }
    appendLog('[错误] ' + msg);
    showMessage(title, msg, true);
  }

  function runApply(app, code, metrics, slider, force) {
    setOperatingState(true, '正在应用更改', '正在准备补丁...');
    appendLog('[操作] 应用 ' + code + ' · 会话指标' + (metrics ? '开启' : '关闭') +
      ' · 思考强度滑块' + (slider ? '开启' : '关闭') + (force ? ' · 强制退出' : ''));

    app.ApplyCustomization(code, metrics, slider, true, force).then(function (res) {
      var name = res && res.LocaleName ? res.LocaleName : code;
      appendLog('[完成] 已应用 ' + name);
      if (!res || !res.Relaunched) showMessage('已应用', '请手动打开 Antigravity。');
    }).catch(function (err) {
      showOperationError('应用失败', err, force ? null : function () { runApply(app, code, metrics, slider, true); });
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
      elEffortSlider.checked = false;
      appendLog('[完成] 已还原为官方英文');
      if (res && !res.AlreadyEnglish && !res.Relaunched) showMessage('已还原', '请手动打开 Antigravity。');
    }).catch(function (err) {
      showOperationError('还原失败', err, force ? null : function () { doRestore(true); });
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
  function closeMoreActions() {
    elMoreActions.querySelector('summary').focus();
    elMoreActions.open = false;
  }
  elMoreActions.addEventListener('keydown', function (event) {
    if (elMoreActions.open && event.key === 'Escape') {
      event.preventDefault();
      closeMoreActions();
    }
  });
  elMoreActions.querySelector('.action-list').addEventListener('click', function (event) {
    if (event.target.closest('button')) closeMoreActions();
  }, true);
  document.addEventListener('click', function (event) {
    if (!elMoreActions.contains(event.target)) elMoreActions.open = false;
  });
  elMoreActions.addEventListener('focusout', function (event) {
    if (!elMoreActions.contains(event.relatedTarget)) elMoreActions.open = false;
  });
  elBtnSecondaryAction.addEventListener('click', launchApp);
  elBtnChangePath.addEventListener('click', handleChangePath);
  elBtnRestore.addEventListener('click', function () {
    elRestoreConfirmModal.showModal();
  });
  elSessionMetrics.addEventListener('change', renderPrimaryLabel);
  elEffortSlider.addEventListener('change', renderPrimaryLabel);
  document.getElementById('btnCancelRestore').addEventListener('click', function () {
    elRestoreConfirmModal.close();
  });
  document.getElementById('btnConfirmRestore').addEventListener('click', function () {
    elRestoreConfirmModal.close();
    doRestore(false);
  });

  elBtnAbout.addEventListener('click', function () {
    elAboutModal.showModal();
  });
  elBtnCloseAbout.addEventListener('click', function () {
    elAboutModal.close();
  });

  elBtnMsgOk.addEventListener('click', function () {
    elMsgModal.close();
  });
  function openLogs() {
    elLogModal.showModal();
    elLogContent.scrollTop = elLogContent.scrollHeight;
  }

  elBtnMsgLogs.addEventListener('click', function () {
    elMsgModal.close();
    openLogs();
  });

  document.getElementById('btnOpenLogs').addEventListener('click', openLogs);
  elBtnCloseLogs.addEventListener('click', function () {
    elLogModal.close();
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

  function showUpdateError(title, err) {
    showMessage(title, errorText(err));
    elReleaseFallback.classList.remove('hidden');
  }

  elCheckUpdate.addEventListener('click', async function () {
    var app = backend();
    if (!app || isOperating) return;
    elUpdateStatus.textContent = '';
    setOperatingState(true, '正在检查更新', '正在查询 GitHub 发布记录…');
    try {
      var release = await app.CheckForUpdate();
      if (!release.available) {
        var currentVersion = elAppVersionBadge.textContent.trim();
        showMessage('已是最新版本', currentVersion ? '当前版本为 ' + currentVersion + '。' : '当前工具已是最新版本。');
        return;
      }
      offeredVersion = release.version;
      document.getElementById('updateTitle').textContent = '发现新版本 ' + release.version;
      document.getElementById('updateSummary').textContent = '当前版本 ' + elAppVersionBadge.textContent + '。更新后将重新打开本工具。';
      document.getElementById('updateNotes').textContent = release.notes;
      elUpdateModal.showModal();
    } catch (err) {
      showUpdateError('检查更新失败', err);
    } finally {
      setOperatingState(false);
    }
  });

  function openReleases() { backend().OpenReleases(); }
  elReleaseFallback.addEventListener('click', openReleases);
  document.getElementById('btnReleaseNotes').addEventListener('click', openReleases);
  document.getElementById('btnLater').addEventListener('click', function () { elUpdateModal.close(); });
  document.getElementById('btnInstallUpdate').addEventListener('click', async function () {
    elUpdateModal.close();
    setOperatingState(true, '正在下载更新', '正在准备下载…');
    elCancelDownload.disabled = false;
    elCancelDownload.classList.remove('hidden');
    try {
      await backend().InstallUpdate(offeredVersion);
    } catch (err) {
      if (errorText(err) === '已取消更新。') elUpdateStatus.textContent = '已取消更新';
      else showUpdateError('更新失败', err);
    } finally {
      elCancelDownload.classList.add('hidden');
      setOperatingState(false);
    }
  });
  elCancelDownload.addEventListener('click', function () {
    elCancelDownload.disabled = true;
    backend().CancelUpdate();
  });

  function setupWailsEvents() {
    if (window.runtime && window.runtime.EventsOn) {
      window.runtime.EventsOn('progress', function (data) {
        if (!data || !isOperating) return;
        if (data.stage === '正在重启工具') elCancelDownload.classList.add('hidden');
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
          if (v) {
            elAppVersionBadge.textContent = 'v' + v;
            app.UpdateWindowReady();
          }
        }).catch(function () {});
      }
      refreshStatus().then(loadLocales);
    }, 120);
  });
})();
