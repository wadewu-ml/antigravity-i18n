# antigravity-i18n


Outil de localisation d'interface utilisateur pour le client de bureau Google Antigravity.


[简体中文](./README.zh-CN.md) | [English](./README.md) | [繁體中文](./README.zh-Hant.md) | [日本語](./README.ja.md) | [한국어](./README.ko.md) | [Deutsch](./README.de.md) | [Español](./README.es.md) | **Français** | [Português (Brasil)](./README.pt-BR.md) | [Русский](./README.ru.md)

---

Un exécutable autonome regroupant une interface graphique moderne, un moteur Go asar pur et des packs de langue complets. Fonctionne immédiatement sans aucune dépendance sur Node.js ou les bibliothèques d'exécution, et prend en charge la restauration exacte en octets vers la version officielle à tout moment.

---

## Principales fonctionnalités

- **Binaire autonome unique** : exécutable autonome d'environ 12 Mo – aucun environnement d'exécution Node.js, npm ou Python requis.
- **GUI & CLI Dual Mode** : Ouvre une fenêtre de bureau propre par défaut, tout en intégrant des sous-commandes CLI complètes (`apply` / `restore` / `status` / `locales`) pour les scripts et l'automatisation.
- **Restauration exacte en octets** : crée automatiquement une sauvegarde vierge (`app.asar.clean-backup`) avant toute modification. La restauration récupère une archive identique octet par octet (vérifiée SHA-256) à la version officielle.
- **Conception non invasive** : localise uniquement l'interface utilisateur du shell, les menus natifs et les boîtes de dialogue de fermeture. Les éditeurs de code (Monaco), les terminaux intégrés (xterm) et les flux de conversations IA restent strictement intacts.
- **Hors ligne et fiable** : zéro activité réseau, zéro collecte de données. Équipé d'une validation initiale de la syntaxe JavaScript et d'une restauration transactionnelle en cas d'erreurs inattendues.

---

## Téléchargement et utilisation

### Télécharger

Téléchargez le dernier « antigravity-i18n.exe » depuis la page des versions et exécutez-le directement.

### 1. Interface graphique (recommandée)

Double-cliquez sur « antigravity-i18n.exe » :
1. L'application détecte automatiquement le dossier d'installation d'Antigravity (ou cliquez sur « Modifier l'emplacement » pour le spécifier manuellement).
2. Cliquez sur votre langue préférée (la sélection de l'anglais restaure la version officielle).
3. Cliquez sur « Appliquer ». Une fois Antigravity fermé, l’outil met à jour les ressources et relance l’application.

Si l’application ou la restauration nécessite de fermer Antigravity, l’outil demande à l’installation sélectionnée de se fermer normalement. Si elle est encore ouverte après 30 secondes, vous pouvez la fermer manuellement et réessayer ou cliquer sur « Forcer la fermeture et réessayer ».

### 2. Interface de ligne de commande

Ouvrez PowerShell dans le dossier qui contient `antigravity-i18n.exe`.

```powershell
# Check current localization status and backups
.\antigravity-i18n.exe status

# Install a language pack (Simplified Chinese by default)
.\antigravity-i18n.exe apply --locale zh-CN

# Restore the official untranslated English version
.\antigravity-i18n.exe restore

# List all bundled language packs with coverage
.\antigravity-i18n.exe locales
```

Options courantes :
- `--app-dir <path>` : spécifiez manuellement le répertoire d'installation d'Antigravity (ou définissez via `ANTIGRAVITY_APP_DIR`).
- `--locale <code>` : code de la langue cible (par exemple, `zh-CN`, `zh-Hant`, `ja`, `ko`, `de`, `es`, `fr`).
- `--no-restart` : Ne relancez pas Antigravity après l'application du correctif.
- `--force` ferme immédiatement l’installation sélectionnée. Les modifications non enregistrées peuvent être perdues.
- `--timeout <ms>` fixe le délai d’attente en millisecondes avant l’annulation de l’opération. La valeur par défaut est `30000`.
- `--yes` : supprime l'invite de confirmation d'arrêt.

Une seule opération d’application ou de restauration est possible à la fois pour chaque installation. Si une autre opération est en cours, attendez qu’elle se termine avant de réessayer. Utilisez `status` pour vérifier l’installation sélectionnée et sa sauvegarde. Une archive illisible ou dont l’identité Antigravity ne peut pas être vérifiée produit une erreur.

---

## Langues groupées

Tous les modules linguistiques sont intégrés directement dans l'exécutable :

| Code | Langue | Couverture | Remarques |
| :--- | :--- | :--- | :--- |
| `zh-CN` | Simplified Chinese (简体中文) | 100% | Couverture complète de l'interface, des menus et des dialogues |
| `zh-Hant` | Traditional Chinese (繁體中文) | 100% | Couverture complète avec terminologie technique vérifiée |
| `ja` | Japanese (日本語) | 100% | Couverture complète avec terminologie technique vérifiée |
| `ko` | Korean (한국어) | 100% | Couverture complète avec terminologie technique vérifiée |
| `de` | German (Deutsch) | 100% | Couverture complète avec terminologie technique vérifiée |
| `es` | Spanish (Español) | 100% | Couverture complète avec terminologie technique vérifiée |
| `fr` | French (Français) | 100% | Couverture complète avec terminologie technique vérifiée |
| `pt-BR` | Portuguese (Brasil) | 100% | Couverture complète avec terminologie technique vérifiée |
| `ru` | Russian (Русский) | 100% | Couverture complète avec terminologie technique vérifiée |
| `en` | English | - | Version officielle intacte (sélectionnez pour restaurer) |

---

## FAQ et mécanique

### 1. Que se passe-t-il après une mise à jour officielle d'Antigravity ?
Les mises à jour officielles écraseront « app.asar ». Rouvrez simplement cet outil et cliquez à nouveau sur « Appliquer ». L'outil actualisera automatiquement sa sauvegarde propre à l'aide de la nouvelle version officielle.

### 2. Changement de langue et restauration
Vous pouvez basculer directement entre n’importe quelle langue ; l'outil remplace de manière transparente les correctifs actifs sans nécessiter de restauration manuelle au préalable. En cliquant sur « Restaurer l'anglais officiel », le client revient à tout moment à un état intact.

### 3. Gestion des fichiers de sauvegarde
- `app.asar.clean-backup` contient l’archive d’origine nécessaire à la restauration. Lors d’une application ou d’une restauration sur une installation non modifiée, l’outil compare les empreintes SHA-256 et actualise la sauvegarde si elles diffèrent, y compris après la réinstallation de la même version. **Conservez ce fichier.**
- `app.asar.bak-*` désigne les copies horodatées que l’outil peut utiliser pour récupérer une sauvegarde manquante ou inutilisable. Une fois que le client fonctionne et que `status` indique une sauvegarde valide, vous pouvez supprimer les anciennes copies pour libérer de l’espace.

---

## Construire à partir de la source

**Prérequis** :
- Go 1.26.3+
- [Wails v2](https://wails.io/) CLI
- PowerShell 7

```powershell
# Clone the repository
git clone https://github.com/wadewu-ml/antigravity-i18n.git
cd antigravity-i18n

# Build both GUI and CLI binaries
.\build.ps1

# Décommentez la ligne suivante pour inclure un test avec une archive officielle.
# $env:ANTIGRAVITY_ASAR = "$env:LOCALAPPDATA\Programs\antigravity\resources\app.asar.clean-backup"
go test ./...
```

---

## Clause de non-responsabilité

1. Il s'agit d'un projet non officiel maintenu par la communauté et n'est ni affilié, ni approuvé ni autorisé par Google ou Antigravity.
2. Google Antigravity et les marques associées appartiennent à leurs propriétaires respectifs.
3. Cet outil est fourni pour un usage personnel et éducatif. Veuillez vous assurer du respect des conditions de service du logiciel concerné.

## Licence

[Licence MIT](./LICENSE)
