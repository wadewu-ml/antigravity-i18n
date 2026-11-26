<div align="center">
  <img src="./frontend/dist/appicon.png" width="68" height="68" alt="Icône Regravity">
  <h1>Regravity</h1>
  <p><em>Antigravity décolle ; Regravity le ramène dans votre langue et sous votre contrôle.</em></p>
  <p>Compagnon portable Windows pour <strong>Google Antigravity</strong> — 9 langues d’interface, métriques de session &amp; quotas de modèles, et curseur d’effort de raisonnement.</p>
  <p>
    <a href="https://github.com/wadewu-ml/regravity/releases/download/v4.2.3/antigravity-i18n.exe"><img src="https://img.shields.io/badge/Download-v4.2.3_(Windows_x64)-0F9D58?style=flat&logo=windows&logoColor=white" alt="Télécharger v4.2.3 pour Windows x64"></a>
    <a href="https://github.com/wadewu-ml/regravity/releases"><img src="https://img.shields.io/badge/Releases-GitHub-18181B?style=flat&logo=github" alt="Versions sur GitHub"></a>
    <img src="https://img.shields.io/badge/Antigravity-2.35.1_tested-18181B?style=flat" alt="Testé avec Antigravity App 2.35.1">
    <a href="./LICENSE"><img src="https://img.shields.io/badge/License-MIT-18181B?style=flat" alt="Licence MIT"></a>
  </p>
  <p><sub>Fichier Windows actuel · <code>antigravity-i18n.exe</code></sub></p>
  <p>
    <sub>
      <a href="./README.zh-CN.md">简体中文</a> · <a href="./README.md">English</a> · <a href="./README.zh-Hant.md">繁體中文</a> · <a href="./README.ja.md">日本語</a> · <a href="./README.ko.md">한국어</a> · <a href="./README.de.md">Deutsch</a> · <a href="./README.es.md">Español</a> · <strong>Français</strong> · <a href="./README.pt-BR.md">Português (Brasil)</a> · <a href="./README.ru.md">Русский</a>
    </sub>
  </p>
</div>

<br>

<p align="center">
  <a href="./docs/images/language-settings.webp">
    <img src="./docs/images/language-settings.webp" width="580" alt="Fenêtre de configuration affichant Antigravity v2.35.2, le chinois simplifié sélectionné ainsi que les métriques de session et le curseur activés">
  </a>
  <br>
  <sub>Fenêtre de configuration de la version de développement de Regravity avec Antigravity 2.35.2.</sub>
</p>

## Ce que propose l’outil

- **9 langues d’interface, sans toucher à vos chats** — Traduit réglages, commandes du chat, menus de modèles, menus natifs et confirmations de fermeture (`zh-CN`, `zh-Hant`, `ja`, `ko`, `de`, `es`, `fr`, `pt-BR`, `ru`) sans modifier vos messages ni vos saisies.
- **Métriques de session et quotas** *(facultatif)* — Ajoute un anneau près du sélecteur de modèle pour consulter quotas restants, heures locales de réinitialisation, contexte utilisé, débit moyen, taux de cache et délai avant le premier jeton.
- **Curseur d’effort de raisonnement** *(facultatif)* — Intègre un curseur capsule dans le sélecteur de modèles pour les modèles compatibles.
- **Modulaire et réversible à l’octet près** — Les deux interrupteurs fonctionnent indépendamment et avec l’interface anglaise d’origine. Restaurez la version officielle depuis une sauvegarde intacte de la même version d’Antigravity à tout moment.

## Démarrage rapide

> [!TIP]
> L’interface de l’outil est en chinois simplifié. Les libellés d’origine figurent entre parenthèses pour correspondre à la capture ci-dessus.

### Appliquer les réglages

1. Lancez `antigravity-i18n.exe`. Si le dossier détecté n’est pas le bon, cliquez sur **Changer d’emplacement** (`更改位置`) et choisissez le dossier contenant `Antigravity.exe`.
2. Sélectionnez une langue dans **Langue de l’interface** (`界面语言`), ou gardez **English** pour n’utiliser que les options facultatives.
3. Activez **Métriques de session** (`会话指标`) et/ou **Curseur d’effort de raisonnement** (`思考强度滑块`) sous la liste des langues.
4. Enregistrez votre travail, quittez Antigravity puis cliquez sur **Appliquer les modifications** (`应用更改`) en bas à droite. Antigravity redémarre avec vos réglages (ou cliquez sur **Ouvrir Antigravity** / `打开 Antigravity`), et vous pouvez fermer l’outil.

> [!WARNING]
> Si l’outil indique qu’Antigravity est encore ouvert, enregistrez votre travail et fermez-le normalement. **Forcer la fermeture et réessayer** (`强制退出并重试`) interrompt les sessions actives et peut faire perdre des modifications non enregistrées.

### Réappliquer ou restaurer

- **Réappliquer** — Cliquez sur **Réappliquer** (`重新应用`) en bas à droite pour réappliquer les réglages actifs.
- **Restaurer la version officielle** — Quittez Antigravity, ouvrez le menu en haut à droite (`···`), cliquez sur **Restaurer la version officielle** (`还原官方版本`) puis confirmez.

> [!IMPORTANT]
> Lors de la première application, l’outil sauvegarde le fichier original sous `resources/app.asar.clean-backup` dans le dossier d’installation d’Antigravity. Conservez ce fichier pour pouvoir restaurer.

## Options facultatives

### Métriques de session et quotas de modèles

- **Avant le premier message** — Survolez l’anneau à côté du sélecteur de modèle pour consulter le quota restant par groupe et par période (hebdomadaire, 5 heures, etc.) ; survolez un pourcentage pour l’heure locale de réinitialisation.
- **Après un appel au modèle** — L’anneau reflète l’utilisation estimée du contexte et le haut du panneau affiche **Contexte utilisé** (jetons, capacité totale, part), **Débit moyen** (jetons de raisonnement inclus), **Taux de réussite du cache** et **Délai avant le premier jeton**. Cliquez pour épingler ou fermer.
- **Suivi par panneau** — En vue partagée, conversations latérales et sous-agents, chaque panneau lit le dernier appel avec données de sa propre session locale. Un `—` indique une valeur indisponible.

<p align="center">
  <a href="./docs/images/session-quota-before-first-call.webp">
    <img src="./docs/images/session-quota-before-first-call.webp" width="420" alt="Panneau des métriques de session et des quotas restants par modèle au-dessus de l’anneau">
  </a>
  <br>
  <sub>Les quotas s’affichent immédiatement ; les métriques se remplissent après le premier appel au modèle.</sub>
</p>

### Curseur d’effort de raisonnement

Activez **Curseur d’effort de raisonnement** (`思考强度滑块`) et appliquez les changements. Dans le sélecteur de modèles d’Antigravity, les modèles compatibles affichent un curseur au-dessus de l’option d’utilisation : faites-le glisser pour prévisualiser et relâchez pour l’appliquer.

<p align="center">
  <a href="./docs/images/thinking-effort-high.webp">
    <img src="./docs/images/thinking-effort-high.webp" height="340" alt="Sélecteur de modèles avec effort de raisonnement élevé">
  </a>
  &nbsp;&nbsp;
  <a href="./docs/images/thinking-effort-max.webp">
    <img src="./docs/images/thinking-effort-max.webp" height="340" alt="Sélecteur de modèles avec effort de raisonnement maximal">
  </a>
  <br>
  <sub>Palier standard (Élevé) et niveau maximal mis en évidence directement dans le menu des modèles.</sub>
</p>

## Mises à jour et CLI

### Mises à jour

- **Après une mise à jour d’Antigravity** — Si la traduction ou les options disparaissent, vérifiez la version prise en charge en haut de cette page puis réappliquez vos réglages.
- **Mises à jour automatiques** — Au démarrage, l’outil s’appuie sur la vérification précédente pour décider de contacter GitHub et attend au moins 24 heures après un succès. Un point bleu sur le menu en haut à droite (`···`) signale une nouvelle version. Désactivable dans **À propos** (`关于`), avec vérification manuelle (`检查更新`) toujours disponible. Après mise à jour, cliquez sur **Réappliquer** (`重新应用`).

Dès la première version de Regravity, téléchargez `regravity.exe` manuellement depuis [Versions](https://github.com/wadewu-ml/regravity/releases) — vos correctifs et préférences de mise à jour seront conservés. Depuis le renommage du dépôt, la vérification intégrée de la `v4.2.3` peut signaler une erreur de comparaison des versions.

### CLI PowerShell

Ouvrez PowerShell dans le dossier contenant `antigravity-i18n.exe` :

| Action | Commande |
| --- | --- |
| Lister les codes de langue | `.\antigravity-i18n.exe locales` |
| Appliquer le français | `.\antigravity-i18n.exe apply --locale fr` |
| Anglais officiel + métriques | `.\antigravity-i18n.exe apply --locale en --session-metrics` |
| Anglais officiel + curseur de raisonnement | `.\antigravity-i18n.exe apply --locale en --effort-slider` |
| Voir les réglages actuels | `.\antigravity-i18n.exe status` |
| Restaurer la version officielle | `.\antigravity-i18n.exe restore` |
| Voir toutes les options | `.\antigravity-i18n.exe --help` |

### Compiler depuis les sources

Nécessite Windows, PowerShell 7, Go 1.26.3+ et la CLI Wails v2 :

```powershell
.\build.ps1 -Target gui
```

Génère `build/bin/regravity.exe`.

## Dépannage

- **WebView2 manquant au démarrage** — Installez le [runtime Microsoft Edge WebView2](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download) puis relancez `antigravity-i18n.exe`.

---

## Clause de non-responsabilité

1. Il s'agit d'un projet non officiel maintenu par la communauté et n'est ni affilié, ni approuvé ni autorisé par Google ou Antigravity.
2. Google Antigravity et les marques associées appartiennent à leurs propriétaires respectifs.
3. Cet outil est fourni pour un usage personnel et éducatif. Veuillez vous assurer du respect des conditions de service du logiciel concerné.

[Licence MIT](./LICENSE)
