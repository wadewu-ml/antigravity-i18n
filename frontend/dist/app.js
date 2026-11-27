// Frontend controller for Regravity.
//
// Field naming follows what Wails actually puts on the wire: types declared
// with json tags (LocaleInfo) arrive lower camel case, while plain structs
// (Status, ApplyResult, RestoreResult) arrive with their Go field names.
(function () {
  'use strict';

  var englishMessages = {
    "未知错误": "Unknown error",
    "取消": "Cancel",
    "知道了": "OK",
    "正在处理...": "Working…",
    "请稍候...": "Please wait…",
    "无法连接": "Disconnected",
    "窗口运行库没有就绪，请重新启动本工具。": "The window runtime is unavailable. Restart Regravity.",
    "读取失败": "Could not read status",
    "[错误] 读取状态失败: ": "[Error] Could not read status: ",
    "未检测到 Antigravity 安装位置": "Antigravity installation not found",
    "版本未知": "Unknown version",
    "官方英文": "Original English",
    "当前语言：": "Current language: ",
    "打开 Antigravity": "Open Antigravity",
    " 干净备份状态：": " Backup status: ",
    "。": ".",
    " 状态文件与实际内容不一致，以 app.asar 的实际内容为准。": " The saved state differs from the installed archive. Showing the archive’s settings.",
    "将切换为 ": "Switch to ",
    "应用后会重启 Antigravity": "Applying changes restarts Antigravity",
    "重新检测": "Detect again",
    "重新应用": "Reapply",
    "应用更改": "Apply changes",
    "应用后开启": "Will turn on",
    "应用后关闭": "Will turn off",
    "没有可用的语言包": "No languages available",
    "读取语言列表失败": "Could not load languages",
    "[错误] 读取语言列表失败: ": "[Error] Could not load languages: ",
    "[提示] ": "[Notice] ",
    "Antigravity 正在运行": "Antigravity is running",
    "请先手动退出 Antigravity 再试。强制退出会中断正在进行的会话，未保存的编辑可能丢失。": "Quit Antigravity and try again. Forcing it to quit interrupts active sessions and may discard unsaved edits.",
    "[错误] ": "[Error] ",
    "正在应用更改": "Applying changes",
    "正在准备补丁...": "Preparing the patch…",
    "[操作] 应用 ": "[Action] Apply ",
    " · 会话指标": " · Session metrics: ",
    "开启": "On",
    "关闭": "Off",
    " · 思考强度滑块": " · Thinking effort slider: ",
    " · 强制退出": " · Force quit",
    "[完成] 已应用 ": "[Done] Applied ",
    "已应用": "Changes applied",
    "请手动打开 Antigravity。": "Open Antigravity to continue.",
    "应用失败": "Could not apply changes",
    "正在还原官方版本": "Restoring original app",
    "正在恢复干净备份并重算校验...": "Restoring and verifying the backup…",
    "[操作] 还原为官方英文": "[Action] Restore original English",
    "[完成] 已还原为官方英文": "[Done] Restored original English",
    "已还原": "Original app restored",
    "还原失败": "Could not restore the original app",
    "[操作] 已启动 Antigravity": "[Action] Opened Antigravity",
    "启动失败": "Could not open Antigravity",
    "正在切换安装位置": "Changing installation",
    "请选择 Antigravity 安装目录": "Choose the Antigravity installation folder",
    "切换位置失败": "Could not change location",
    "已复制": "Copied",
    "运行日志已复制到剪贴板。": "Activity log copied to the clipboard.",
    "复制失败": "Copy failed",
    "无法访问剪贴板，请手动选择日志内容。": "Clipboard unavailable. Select the log text and copy it manually.",
    "发现新版本 ": "Update available: ",
    "检查更新": "Check for updates",
    "更多操作，": "More actions, ",
    "更多操作": "More actions",
    "[更新] 无法保存检查时间：": "[Update] Could not save the check time: ",
    "GitHub 暂时限制了更新请求，请在 ": "GitHub has limited update requests. Try again after ",
    " 后重试。": ".",
    "[更新] ": "[Update] ",
    "[更新] 无法读取更新设置：": "[Update] Could not read update settings: ",
    "保存失败": "Could not save settings",
    "正在检查更新": "Checking for updates",
    "正在查询 GitHub 发布记录…": "Checking GitHub releases…",
    "已是最新版本": "Up to date",
    "当前版本为 ": "Current version: ",
    "当前工具已是最新版本。": "You have the latest version.",
    "当前版本 ": "Current version: ",
    "。更新后将重新打开本工具。": ". Regravity will restart after the update.",
    "检查更新失败": "Could not check for updates",
    "正在下载更新": "Downloading update",
    "正在准备下载…": "Preparing download…",
    "已取消更新": "Update canceled",
    "更新失败": "Update failed",
    "[进度] ": "[Progress] ",
    "已取消更新。": "Update canceled.",
    "正在重启工具": "Restarting Regravity",
    "正在关闭 Antigravity": "Closing Antigravity",
    "正在校验补丁": "Verifying patch",
    "正在准备干净备份": "Preparing original backup",
    "正在校验补丁内容": "Verifying patch contents",
    "设置已生效，无需写回": "Settings already applied",
    "已创建安全备份": "Safety backup created",
    "正在写回 app.asar": "Writing app.asar",
    "正在定位干净备份": "Finding the original backup",
    "已经是官方英文版本": "Already using original English",
    "正在校验版本": "Verifying version",
    "正在执行另一项操作。": "Another operation is in progress.",
    "请先检查并确认更新版本。": "Check for updates and confirm the version first.",
    "另一个窗口正在更新这个工具，请稍后重试。": "Another Regravity window is updating. Try again later.",
    "等待新版窗口就绪超时。": "The updated window timed out.",
    "新版程序在窗口就绪前退出。": "The updated app closed before its window was ready.",
    "无效的更新下载重定向": "Invalid update download redirect.",
    "GitHub 暂时限制了更新请求，请稍后重试，或打开发布页手动下载。": "GitHub has limited update requests. Try again later or download the update from Releases.",
    "无法比较当前工具与正式发布的版本号。": "Could not compare versions. Open Releases to download the latest version.",
    "无法确认更新程序的文件大小。": "Could not verify the update size.",
    "获取的发布信息与正式版本不符。": "The release information does not match the expected version.",
    "该版本缺少 Windows 程序或 SHA-256 校验信息。": "The release is missing its Windows executable or checksum.",
    "发布附件的摘要格式无效。": "Invalid release checksum format.",
    "校验文件中缺少 Windows 程序的 SHA-256。": "The checksum file has no entry for the Windows executable.",
    "无效的 SHA-256 校验值。": "Invalid SHA-256 checksum.",
    "下载的文件与发布版本的大小或 SHA-256 不符，请重新下载。": "The downloaded file does not match the published size or checksum. Download it again.",
    "请先确认是否继续。": "Confirm the operation first.",
    "完整更新说明可在 GitHub 发布页查看。": "See the full notes on GitHub Releases."
  };
  var diagnosticPatterns = [
    [/^更新请求失败（HTTP (\d+)）$/, "Update request failed (HTTP $1)."],
    [/^程序版本不符，预期 (.+)。$/, "Unexpected version. Expected $1."],
    [/^无法在工具所在目录创建更新记录：(.+)$/, "Could not create an update record: $1"],
    [/^更新记录无效，请检查 (.+)$/, "Invalid update record: $1"],
    [/^上次更新保留的旧版位于 (.+)$/, "Previous version saved at $1"],
    [/^上次更新（(.+) → (.+)）的记录位于 (.+)$/, "Update record ($1 \u2192 $2): $3"],
    [/^请重新打开当前工具再更新：(.+)$/, "Reopen Regravity before updating: $1"],
    [/^新版启动失败（(.+)），可直接打开旧版 (.+)：(.+)$/, "Updated app could not start ($1). Open the previous version at $2: $3"],
    [/^更新失败（(.+)），可直接打开旧版 (.+)：(.+)$/, "Update failed ($1). Open the previous version at $2: $3"],
    [/^已恢复旧版：(.+)$/, "Previous version restored: $1"]
  ];

  var currentStatus = null;
  var statusRequest = 0;
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
  var elAutoCheck = document.getElementById('autoCheckUpdates');
  var checkInFlight = null;
  var latestRelease = null;
  var hour = 60 * 60 * 1000;
  var updateCache = {};
  var toolVersion = '';
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

  var languageSelect = document.getElementById('setupLanguage');
  var setupLanguage = 'zh-CN';
  try {
    var savedLanguage = localStorage.getItem('setupLanguage');
    if (Array.from(languageSelect.options).some(option => option.value === savedLanguage)) setupLanguage = savedLanguage;
  } catch (err) { /* Keep the default when storage is unavailable. */ }

  var staticLabels = Array.from(document.querySelectorAll('[data-en], [data-en-title]'), function (element) {
    var attribute = element.hasAttribute('data-en-title') ? 'title' : 'textContent';
    return {element: element, attribute: attribute, chinese: element[attribute],
      english: attribute === 'title' ? element.dataset.enTitle : element.dataset.en};
  });

  function fromEnglish(text) {
    return setupLanguage === 'en' ? text : setupMessages[setupLanguage]?.[text] ?? text;
  }

  function t(text) {
    return setupLanguage !== 'zh-CN' && Object.hasOwn(englishMessages, text) ? fromEnglish(englishMessages[text]) : text;
  }

  function translateDiagnostic(message) {
    if (setupLanguage === 'zh-CN') return message;
    return message.split('\n').map(function (line) {
      if (Object.hasOwn(englishMessages, line)) return t(line);
      for (var [pattern, replacement] of diagnosticPatterns) {
        const match = pattern.exec(line);
        if (match) {
          return fromEnglish(replacement).replace(/\$(\d+)/g, (_, index) => translateDiagnostic(match[index]));
        }
      }
      return line;
    }).join('\n');
  }

  function renderSetupLanguage() {
    document.documentElement.lang = setupLanguage;
    document.documentElement.dir = setupLanguage === 'ar' ? 'rtl' : 'ltr';
    if (elUpdateStatus.textContent) elUpdateStatus.textContent = t('已取消更新');
    languageSelect.value = setupLanguage;
    languageSelect.lang = setupLanguage;
    languageSelect.setAttribute('aria-label', setupLanguage === 'zh-CN' ? '设置窗口语言' : fromEnglish('Setup language'));
    staticLabels.forEach(function (label) {
      label.element[label.attribute] = setupLanguage === 'zh-CN' ? label.chinese : fromEnglish(label.english);
    });
    renderUpdateNotice();
  }

  languageSelect.addEventListener('change', function () {
    setupLanguage = languageSelect.value;
    try { localStorage.setItem('setupLanguage', setupLanguage); } catch (err) { /* Session only. */ }
    renderSetupLanguage();
    if (currentStatus) renderStatus(currentStatus);
    else refreshStatus();
    renderChips();
    if (operationLogs.length) elLogContent.textContent = operationLogs.join('\n');
  });

  function backend() {
    return window.go && window.go.main && window.go.main.App ? window.go.main.App : null;
  }

  function errorText(err) {
    if (!err) return t('未知错误');
    return translateDiagnostic(typeof err === 'string' ? err : err.message || String(err));
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
    elBtnMsgOk.textContent = retry ? t('取消') : t('知道了');
    elBtnMsgOk.className = retry ? 'btn btn-ghost' : 'btn btn-primary';
    elBtnForceRetry.classList.toggle('hidden', !retry);
    elBtnForceRetry.onclick = retry ? function () {
      elMsgModal.close();
      retry();
    } : null;
    elMsgModal.showModal();
    elBtnMsgOk.focus();
  }

  function setProgressDetail(detail) {
    elProgressDetail.textContent = detail;
    elProgressDetail.title = detail;
  }

  function setOperatingState(operating, stage, detail) {
    isOperating = operating;
    if (operating) {
      elProgressBox.classList.remove('hidden');
      elProgressStage.textContent = stage || t('正在处理...');
      setProgressDetail(detail || t('请稍候...'));
    } else {
      elProgressBox.classList.add('hidden');
    }
    [elBtnPrimaryAction, elBtnSecondaryAction, elBtnRestore, elBtnChangePath, elCheckUpdate, languageSelect].forEach(function (btn) {
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
      elStatusLabel.textContent = t('无法连接');
      elStatusDesc.textContent = t('窗口运行库没有就绪，请重新启动本工具。');
      return Promise.resolve();
    }
    var request = ++statusRequest;
    return app.GetStatus().then(function (st) {
      if (request === statusRequest) setStatus(st, false);
    }).catch(function (err) {
      if (request !== statusRequest) return;
      currentStatus = null;
      renderPrimaryLabel();
      elStatusPill.className = 'status-pill danger';
      elStatusLabel.textContent = t('读取失败');
      elStatusDesc.textContent = errorText(err);
      appendLog(t('[错误] 读取状态失败: ') + errorText(err));
    });
  }

  function renderStatus(st) {
    if (!st) return;

    elAppPathDisplay.textContent = st.AppDir || t('未检测到 Antigravity 安装位置');
    elAppPathDisplay.dir = st.AppDir ? 'ltr' : 'auto';
    elAppPathDisplay.title = st.AppDir || '';
    elAppVersion.textContent = st.AppVersion ? 'v' + st.AppVersion : t('版本未知');

    elStatusPill.className = 'status-pill';
    elBtnSecondaryAction.classList.add('hidden');
    elBtnRestore.classList.add('hidden');

    var currentName = st.CurrentLanguage === 'en' ? t('官方英文') : (st.LocaleName || st.CurrentLanguage);
    elStatusLabel.textContent = t('当前语言：') + currentName;
    elStatusDesc.textContent = '';
    if (isCustomized(st)) {
      elStatusPill.classList.add('success');
      elBtnSecondaryAction.classList.remove('hidden');
      elBtnSecondaryAction.textContent = t('打开 Antigravity');
      elBtnRestore.classList.remove('hidden');
    } else if (st.HasCleanBackup) {
      elStatusPill.classList.add('default');
      elBtnSecondaryAction.classList.remove('hidden');
      elBtnSecondaryAction.textContent = t('打开 Antigravity');
    }

    if (st.CleanBackupStatus && st.CleanBackupStatus !== 'valid') {
      elStatusDesc.textContent += t(' 干净备份状态：') + st.CleanBackupStatus + t('。');
    }
    if (st.StateIsStale) {
      elStatusDesc.textContent += t(' 状态文件与实际内容不一致，以 app.asar 的实际内容为准。');
    }
  }
  function renderPrimaryLabel() {
    var languageChanged = currentStatus && selectedLocale && currentStatus.CurrentLanguage !== selectedLocale;
    var pending = languageChanged && locales.find(function (loc) { return loc.code === selectedLocale; });
    elPendingLanguage.textContent = pending ? t('将切换为 ') + pending.name : '';
    elEffortSlider.disabled = elSessionMetrics.disabled = isOperating || !currentStatus;
    elBtnPrimaryAction.title = currentStatus ? t('应用后会重启 Antigravity') : '';
    if (!currentStatus) {
      elBtnPrimaryAction.textContent = t('重新检测');
      elBtnPrimaryAction.disabled = isOperating;
      return;
    }
    var metricsChanged = Boolean(currentStatus.SessionMetrics) !== elSessionMetrics.checked;
    var sliderChanged = Boolean(currentStatus.EffortSlider) !== elEffortSlider.checked;
    var hasChanges = languageChanged || metricsChanged || sliderChanged;
    var canReapply = !hasChanges && isCustomized(currentStatus);
    elBtnPrimaryAction.textContent = canReapply ? t('重新应用') : t('应用更改');
    elBtnPrimaryAction.disabled = isOperating || (!hasChanges && !canReapply);
    elEffortSliderDetail.textContent = sliderChanged
      ? (elEffortSlider.checked ? t('应用后开启') : t('应用后关闭'))
      : '';
    elSessionMetricsDetail.textContent = metricsChanged
      ? (elSessionMetrics.checked ? t('应用后开启') : t('应用后关闭'))
      : '';
  }

  function renderChips() {
    elLocaleChips.innerHTML = '';
    if (!locales.length) {
      elLocaleNote.textContent = t('没有可用的语言包');
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
      name.dir = 'auto';
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
      elLocaleNote.textContent = t('读取语言列表失败');
      appendLog(t('[错误] 读取语言列表失败: ') + errorText(err));
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
      appendLog(t('[提示] ') + msg.slice(runningPrefix.length));
      showMessage(t('Antigravity 正在运行'), t('请先手动退出 Antigravity 再试。强制退出会中断正在进行的会话，未保存的编辑可能丢失。'), true, retry);
      return;
    }
    appendLog(t('[错误] ') + msg);
    showMessage(title, msg, true);
  }

  function runApply(app, code, metrics, slider, force) {
    setOperatingState(true, t('正在应用更改'), t('正在准备补丁...'));
    appendLog(t('[操作] 应用 ') + code + t(' · 会话指标') + (metrics ? t('开启') : t('关闭')) +
      t(' · 思考强度滑块') + (slider ? t('开启') : t('关闭')) + (force ? t(' · 强制退出') : ''));

    app.ApplyCustomization(code, metrics, slider, true, force).then(function (res) {
      var name = res && res.LocaleName ? res.LocaleName : code;
      appendLog(t('[完成] 已应用 ') + name);
      if (!res || !res.Relaunched) showMessage(t('已应用'), t('请手动打开 Antigravity。'));
    }).catch(function (err) {
      showOperationError(t('应用失败'), err, force ? null : function () { runApply(app, code, metrics, slider, true); });
    }).then(function () {
      setOperatingState(false);
      return refreshStatus();
    });
  }

  function doRestore(force) {
    if (isOperating) return;
    var app = backend();
    if (!app) return;

    setOperatingState(true, t('正在还原官方版本'), t('正在恢复干净备份并重算校验...'));
    appendLog(t('[操作] 还原为官方英文'));

    app.RestoreLocale(true, force).then(function (res) {
      selectedLocale = 'en';
      elSessionMetrics.checked = false;
      elEffortSlider.checked = false;
      appendLog(t('[完成] 已还原为官方英文'));
      if (res && !res.AlreadyEnglish && !res.Relaunched) showMessage(t('已还原'), t('请手动打开 Antigravity。'));
    }).catch(function (err) {
      showOperationError(t('还原失败'), err, force ? null : function () { doRestore(true); });
    }).then(function () {
      setOperatingState(false);
      return refreshStatus();
    });
  }

  function launchApp() {
    var app = backend();
    if (!app) return;
    app.LaunchAntigravity().then(function () {
      appendLog(t('[操作] 已启动 Antigravity'));
    }).catch(function (err) {
      showMessage(t('启动失败'), errorText(err));
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
    if (!app || isOperating) return;
    setOperatingState(true, t('正在切换安装位置'), t('请选择 Antigravity 安装目录'));
    app.SelectAppDir(t('请选择 Antigravity 安装目录')).then(function (selected) {
      if (!selected) return null;
      setProgressDetail(selected);
      return app.SetAppDir(selected).then(function (st) {
        // Invalidate any status request for the previously selected directory.
        statusRequest++;
        setStatus(st, true);
      });
    }).catch(function (err) {
      var msg = errorText(err);
      appendLog(t('[错误] ') + msg);
      showMessage(t('切换位置失败'), msg);
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
        showMessage(t('已复制'), t('运行日志已复制到剪贴板。'));
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
            showMessage(ok === false ? t('复制失败') : t('已复制'),
              ok === false ? t('无法访问剪贴板，请手动选择日志内容。') : t('运行日志已复制到剪贴板。'));
          }).catch(function () {
            showMessage(t('复制失败'), t('无法访问剪贴板，请手动选择日志内容。'));
          });
          return;
        }
        showMessage(t('已复制'), t('运行日志已复制到剪贴板。'));
        return;
      } catch (err) {
        // Fall through to the manual-selection hint below.
      }
    }
    showMessage(t('复制失败'), t('无法访问剪贴板，请手动选择日志内容。'));
  }

  function showUpdateError(title, err) {
    showMessage(title, errorText(err));
    elReleaseFallback.classList.remove('hidden');
  }

  function renderUpdateNotice() {
    var version = elAutoCheck.checked && toolVersion && updateCache.current === toolVersion && updateCache.available;
    var label = version ? t('发现新版本 ') + version : t('检查更新');
    elCheckUpdate.textContent = label;
    elMoreActions.classList.toggle('has-update', Boolean(version));
    var summary = elMoreActions.querySelector('summary');
    summary.setAttribute('aria-label', version ? t('更多操作，') + label : t('更多操作'));
    summary.title = summary.getAttribute('aria-label');
  }

  function saveUpdateCache() {
    try {
      localStorage.setItem('updateCheck', JSON.stringify(updateCache));
    } catch (err) {
      appendLog(t('[更新] 无法保存检查时间：') + errorText(err));
    }
  }

  function queryUpdate() {
    if (checkInFlight) return checkInFlight;
    if (Date.now() < updateCache.retryAt) {
      return Promise.reject(new Error(t('GitHub 暂时限制了更新请求，请在 ') +
        new Date(updateCache.retryAt).toLocaleTimeString() + t(' 后重试。')));
    }
    // Retain the last known version if the window closes or the network drops.
    updateCache.nextCheck = Date.now() + hour;
    saveUpdateCache();
    checkInFlight = Promise.resolve().then(function () {
      return backend().CheckForUpdate();
    }).then(function (result) {
      if (result.error) {
        updateCache.retryAt = result.retryAt || 0;
        throw new Error(result.error);
      }
      latestRelease = result.release;
      updateCache = {
        current: toolVersion,
        available: latestRelease.available ? latestRelease.version : '',
        nextCheck: Date.now() + 24 * hour,
        retryAt: latestRelease.retryAt || 0
      };
      return latestRelease;
    }).catch(function (err) {
      updateCache.nextCheck = Math.max(updateCache.retryAt || 0, Date.now() + hour);
      throw err;
    }).finally(function () {
      checkInFlight = null;
      saveUpdateCache();
      renderUpdateNotice();
    });
    return checkInFlight;
  }

  function checkOnStartup() {
    renderUpdateNotice();
    if (!toolVersion || !elAutoCheck.checked || Date.now() < updateCache.retryAt ||
        Date.now() < updateCache.nextCheck) return;
    queryUpdate().catch(function (err) {
      appendLog(t('[更新] ') + errorText(err));
    });
  }

  try {
    elAutoCheck.checked = localStorage.getItem('autoCheckUpdates') !== 'off';
    var saved = JSON.parse(localStorage.getItem('updateCheck'));
    if (saved && typeof saved === 'object') updateCache = saved;
  } catch (err) {
    appendLog(t('[更新] 无法读取更新设置：') + errorText(err));
  }
  elAutoCheck.addEventListener('change', function () {
    try {
      localStorage.setItem('autoCheckUpdates', elAutoCheck.checked ? 'on' : 'off');
    } catch (err) {
      elAutoCheck.checked = !elAutoCheck.checked;
      showMessage(t('保存失败'), errorText(err));
      return;
    }
    checkOnStartup();
  });

  elCheckUpdate.addEventListener('click', async function () {
    var app = backend();
    if (!app || isOperating) return;
    elUpdateStatus.textContent = '';
    setOperatingState(true, t('正在检查更新'), t('正在查询 GitHub 发布记录…'));
    try {
      var release = elMoreActions.classList.contains('has-update') && latestRelease && latestRelease.available
        ? latestRelease : await queryUpdate();
      if (!release.available) {
        var currentVersion = elAppVersionBadge.textContent.trim();
        showMessage(t('已是最新版本'), currentVersion ? t('当前版本为 ') + currentVersion + t('。') : t('当前工具已是最新版本。'));
        return;
      }
      offeredVersion = release.version;
      document.getElementById('updateTitle').textContent = t('发现新版本 ') + release.version;
      document.getElementById('updateSummary').textContent = t('当前版本 ') + elAppVersionBadge.textContent + t('。更新后将重新打开本工具。');
      document.getElementById('updateNotes').textContent = t(release.notes);
      elUpdateModal.showModal();
    } catch (err) {
      showUpdateError(t('检查更新失败'), err);
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
    setOperatingState(true, t('正在下载更新'), t('正在准备下载…'));
    elCancelDownload.disabled = false;
    elCancelDownload.classList.remove('hidden');
    try {
      await backend().InstallUpdate(offeredVersion);
    } catch (err) {
      if (errorText(err) === t('已取消更新。')) elUpdateStatus.textContent = t('已取消更新');
      else showUpdateError(t('更新失败'), err);
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
        if (data.stage) elProgressStage.textContent = t(data.stage);
        if (data.detail) setProgressDetail(data.detail);
        appendLog(t('[进度] ') + t(data.stage || '') + (data.detail ? ': ' + data.detail : ''));
      });
    }
  }

  window.addEventListener('DOMContentLoaded', function () {
    renderSetupLanguage();
    setupWailsEvents();
    // Wait briefly for the Wails IPC layer to attach window.go.
    setTimeout(function () {
      var app = backend();
      if (app) {
        app.GetVersion().then(function (v) {
          if (v) {
            elAppVersionBadge.textContent = 'v' + v;
            document.getElementById('menuVersion').textContent = elAppVersionBadge.textContent;
            toolVersion = v;
            app.UpdateWindowReady();
            checkOnStartup();
          }
        }).catch(function () {});
      }
      refreshStatus().then(loadLocales);
    }, 120);
  });
})();
