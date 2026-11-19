# <img src="./frontend/dist/appicon.png" width="32" height="32" alt=""> antigravity-i18n

Neuf packs de langue traduisent les réglages, les commandes du chat, les menus et les demandes de confirmation à la fermeture de Google Antigravity pour Windows. Un indicateur facultatif dans la zone de saisie permet de voir l’espace restant dans la fenêtre de contexte et la vitesse de réponse du modèle. Une fois les réglages appliqués, vous pouvez fermer l’outil.

[Télécharger pour Windows x64](https://github.com/wadewu-ml/antigravity-i18n/releases/latest/download/antigravity-i18n.exe) · [Versions](https://github.com/wadewu-ml/antigravity-i18n/releases)

Testé avec **Antigravity App 2.33.0**. Si l’outil demande WebView2 au démarrage, installez le [runtime Microsoft Edge WebView2](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download).

[简体中文](./README.zh-CN.md) · [English](./README.md) · [繁體中文](./README.zh-Hant.md) · [日本語](./README.ja.md) · [한국어](./README.ko.md) · [Deutsch](./README.de.md) · [Español](./README.es.md) · **Français** · [Português (Brasil)](./README.pt-BR.md) · [Русский](./README.ru.md)

## Changer la langue de l’interface

Les langues proposées sont le chinois simplifié, le chinois traditionnel, le japonais, le coréen, l’allemand, l’espagnol, le français, le portugais du Brésil et le russe. L’outil de configuration est en chinois simplifié.

1. Ouvrez `antigravity-i18n.exe`. Si l’outil ne trouve pas la bonne installation, cliquez sur **更改位置** et sélectionnez le dossier qui contient `Antigravity.exe`.
2. Choisissez une langue dans **界面语言**. Activez **会话指标** pour afficher les métriques de session. Pour les utiliser avec l’interface anglaise d’origine, sélectionnez **English**.
3. Enregistrez votre travail et quittez Antigravity, puis cliquez sur **应用更改** pour appliquer les réglages. Antigravity s’ouvre à nouveau avec ces réglages.

Si l’outil vous demande de quitter Antigravity, le bouton **强制退出并重试** force sa fermeture et relance l’opération. Les sessions en cours sont interrompues et les modifications non enregistrées risquent d’être perdues.

## Consulter les métriques de session

Dès qu’un appel au modèle est enregistré pour la conversation, un anneau apparaît à droite du sélecteur de modèle dans la zone de saisie. Son remplissage indique la proportion estimée de la fenêtre de contexte déjà utilisée. Survolez-le pour voir le nombre de tokens restants, la vitesse de sortie, le taux de succès du cache et le délai avant le premier token. Cliquez pour épingler les détails, puis cliquez à nouveau pour les refermer.

Chaque panneau utilise le dernier appel au modèle pour lequel sa session dispose de données locales, y compris en vue partagée, dans les conversations latérales et pour les sous-agents. Le nombre de tokens de sortie inclut les tokens de raisonnement signalés par le modèle. Un `—` indique qu’une valeur est indisponible.

## Restaurer ou réappliquer

Pour retrouver l’interface anglaise d’origine et retirer l’indicateur, quittez Antigravity, cliquez sur **还原官方版本** dans l’outil et confirmez avec **还原**. L’outil sauvegarde le fichier d’origine à la première application des réglages. Conservez `resources/app.asar.clean-backup` dans le dossier d’installation d’Antigravity pour la restauration.

Si la traduction ou l’indicateur disparaît après une mise à jour d’Antigravity, vérifiez la version prise en charge en haut de cette page, puis sélectionnez et appliquez à nouveau vos réglages. Pour réappliquer les réglages déjà actifs, cliquez sur **重新应用**.

<details>
<summary>Commandes PowerShell</summary>

Ouvrez PowerShell dans le dossier contenant `antigravity-i18n.exe`.

| Action | Commande |
| --- | --- |
| Afficher les codes de langue | `.\antigravity-i18n.exe locales` |
| Appliquer le français | `.\antigravity-i18n.exe apply --locale fr` |
| Utiliser l’anglais officiel avec les métriques | `.\antigravity-i18n.exe apply --locale en --session-metrics` |
| Voir les réglages actuels | `.\antigravity-i18n.exe status` |
| Restaurer la version officielle | `.\antigravity-i18n.exe restore` |
| Voir toutes les options | `.\antigravity-i18n.exe --help` |

</details>

<details>
<summary>Compiler à partir des sources</summary>

Pour compiler sous Windows, installez PowerShell 7, Go 1.26.3 ou ultérieur et la CLI de Wails v2. Exécutez `.\build.ps1 -Target gui` à la racine du dépôt. Le fichier est généré dans `build/bin/antigravity-i18n.exe`.

</details>

Outil non officiel maintenu par la communauté. [Licence MIT](./LICENSE)
