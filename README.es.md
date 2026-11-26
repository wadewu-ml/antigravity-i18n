<div align="center">
  <img src="./frontend/dist/appicon.png" width="68" height="68" alt="Icono de Regravity">
  <h1>Regravity</h1>
  <p><em>Antigravity despega; Regravity lo trae de vuelta a tu idioma y bajo tu control.</em></p>
  <p>Complemento portátil para <strong>Google Antigravity</strong> en Windows — 9 idiomas de interfaz, métricas de sesión y cuotas de modelos, y control de nivel de razonamiento.</p>
  <p>
    <a href="https://github.com/wadewu-ml/regravity/releases/latest/download/regravity.exe"><img src="https://img.shields.io/badge/Download-Latest_(Windows_x64)-0F9D58?style=flat&logo=windows&logoColor=white" alt="Descargar la última versión de Regravity para Windows x64"></a>
    <a href="https://github.com/wadewu-ml/regravity/releases"><img src="https://img.shields.io/badge/Releases-GitHub-18181B?style=flat&logo=github" alt="Versiones en GitHub"></a>
    <img src="https://img.shields.io/badge/Antigravity-2.35.1_tested-18181B?style=flat" alt="Probado con Antigravity App 2.35.1">
    <a href="./LICENSE"><img src="https://img.shields.io/badge/License-MIT-18181B?style=flat" alt="Licencia MIT"></a>
  </p>
  <p>
    <sub>
      <a href="./README.zh-CN.md">简体中文</a> · <a href="./README.md">English</a> · <a href="./README.zh-Hant.md">繁體中文</a> · <a href="./README.ja.md">日本語</a> · <a href="./README.ko.md">한국어</a> · <a href="./README.de.md">Deutsch</a> · <strong>Español</strong> · <a href="./README.fr.md">Français</a> · <a href="./README.pt-BR.md">Português (Brasil)</a> · <a href="./README.ru.md">Русский</a> · <a href="./README.tr.md">Türkçe</a> · <a href="./README.vi.md">Tiếng Việt</a>
    </sub>
  </p>
</div>

<br>

<p align="center">
  <a href="./docs/images/language-settings.webp">
    <img src="./docs/images/language-settings.webp" width="580" alt="Ventana de configuración con Antigravity v2.35.2, chino simplificado seleccionado y las opciones de métricas de sesión y control deslizante activadas">
  </a>
  <br>
  <sub>Ventana de configuración de Regravity con Antigravity 2.35.2.</sub>
</p>

## Qué incluye

- **9 idiomas de interfaz, sin tocar tus chats** — Traduce ajustes, controles del chat, menús de modelos, menús nativos y avisos de salida (`zh-CN`, `zh-Hant`, `ja`, `ko`, `de`, `es`, `fr`, `pt-BR`, `ru`) sin alterar tus entradas ni conversaciones.
- **Métricas de sesión y cuotas** *(opcional)* — Añade un anillo junto al selector de modelo con cuotas restantes, horarios locales de restablecimiento, contexto utilizado, velocidad media de salida, aciertos de caché y tiempo hasta el primer token.
- **Control de nivel de razonamiento** *(opcional)* — Integra un control deslizante tipo cápsula en el selector de modelos para modelos compatibles.
- **Modular y reversible byte a byte** — Ambos interruptores funcionan por separado y también con la interfaz oficial en inglés. Restaura la versión oficial desde una copia limpia de la misma versión de Antigravity en cualquier momento.

## Inicio rápido

> [!TIP]
> Las etiquetas chinas entre paréntesis corresponden a la captura.

En la versión de desarrollo puedes elegir el idioma de la herramienta arriba a la derecha. Antigravity también admite turco y vietnamita.

### Aplicar ajustes

1. Abre `regravity.exe`. Si la ruta detectada no es correcta, pulsa **Cambiar ubicación** (`更改位置`) y selecciona la carpeta que contiene `Antigravity.exe`.
2. Elige un idioma en **Idioma de la interfaz** (`界面语言`), o deja **English** si solo quieres los controles opcionales.
3. Activa **Métricas de sesión** (`会话指标`) y/o **Control de nivel de razonamiento** (`思考强度滑块`) debajo de la lista de idiomas.
4. Guarda tu trabajo, cierra Antigravity y pulsa **Aplicar cambios** (`应用更改`) abajo a la derecha. Antigravity se reabrirá con tus ajustes (o ábrelo con **Abrir Antigravity** / `打开 Antigravity`), y podrás cerrar la herramienta.

> [!WARNING]
> Si aparece un aviso de que Antigravity sigue abierto, guarda tu trabajo y ciérralo normalmente. **Forzar el cierre y reintentar** (`强制退出并重试`) interrumpe sesiones activas y puede perder cambios sin guardar.

### Volver a aplicar o restaurar

- **Volver a aplicar** — Tras actualizar la herramienta, pulsa **Volver a aplicar** (`重新应用`) para usar las nuevas traducciones.
- **Restaurar la versión oficial** — Cierra Antigravity, abre el menú superior derecho (`···`), pulsa **Restaurar la versión oficial** (`还原官方版本`) y confirma.

> [!IMPORTANT]
> La primera vez que aplicas los ajustes, la herramienta guarda una copia intacta en `resources/app.asar.clean-backup` dentro de la carpeta de instalación de Antigravity. Conserva ese archivo para poder restaurarlo.

## Controles opcionales

### Métricas de sesión y cuotas de modelos

- **Antes del primer mensaje** — Pasa el puntero sobre el anillo junto al selector de modelo para ver la cuota restante por grupo y periodo (semanal, 5 horas, etc.); pasa sobre un porcentaje para ver la hora local de restablecimiento.
- **Tras la primera llamada** — El anillo refleja el uso estimado del contexto y la parte superior muestra **Contexto utilizado** (tokens, capacidad total y porcentaje), **Velocidad media de salida** (incluidos tokens de razonamiento), **Tasa de aciertos de caché** y **Tiempo hasta el primer token**. Haz clic en el anillo para fijar o cerrar.
- **Precisión por panel** — Vistas divididas, chats laterales y subagentes leen la llamada más reciente con datos de su propia sesión local. Un `—` indica dato no disponible.

<p align="center">
  <a href="./docs/images/session-quota-before-first-call.webp">
    <img src="./docs/images/session-quota-before-first-call.webp" width="420" alt="Panel de métricas de sesión y cuota restante por modelo sobre el anillo junto al selector de modelo">
  </a>
  <br>
  <sub>Las cuotas están disponibles de inmediato; las métricas se rellenan tras la primera llamada al modelo.</sub>
</p>

### Control de nivel de razonamiento

Activa **Control de nivel de razonamiento** (`思考强度滑块`) y aplica los cambios. En el selector de modelos de Antigravity, los modelos compatibles muestran un control deslizante encima de la opción de uso: arrástralo para previsualizar y suéltalo para aplicarlo.

<p align="center">
  <a href="./docs/images/thinking-effort-high.webp">
    <img src="./docs/images/thinking-effort-high.webp" height="340" alt="Selector de modelos con el nivel de razonamiento en Alto">
  </a>
  &nbsp;&nbsp;
  <a href="./docs/images/thinking-effort-max.webp">
    <img src="./docs/images/thinking-effort-max.webp" height="340" alt="Selector de modelos con el nivel de razonamiento en Máximo">
  </a>
  <br>
  <sub>Nivel estándar (Alto) y nivel máximo destacado directamente en el menú de modelos.</sub>
</p>

## Actualizaciones y CLI

### Actualizaciones

- **Tras actualizar Antigravity** — Si desaparecen la traducción o los controles, comprueba la versión compatible arriba y vuelve a aplicar tus ajustes.
- **Actualizaciones automáticas** — Al abrirse, decide si consultar GitHub según el resultado anterior y espera al menos 24 horas tras una comprobación correcta. Un punto azul en el menú superior derecho (`···`) señala una nueva versión. Puedes desactivar la búsqueda automática en **Acerca de** (`关于`) o buscar manualmente (`检查更新`). Tras actualizar, pulsa **Volver a aplicar** (`重新应用`).

Si usas la versión v4.2.3 o anterior, descarga y ejecuta [`regravity.exe`](https://github.com/wadewu-ml/regravity/releases/latest/download/regravity.exe) manualmente. Se conservarán tus ajustes.

### CLI de PowerShell

Ejecuta desde PowerShell en la carpeta que contiene `regravity.exe`:

| Acción | Comando |
| --- | --- |
| Listar códigos de idioma | `.\regravity.exe locales` |
| Aplicar español | `.\regravity.exe apply --locale es` |
| Inglés oficial + métricas | `.\regravity.exe apply --locale en --session-metrics` |
| Inglés oficial + control de razonamiento | `.\regravity.exe apply --locale en --effort-slider` |
| Ver configuración actual | `.\regravity.exe status` |
| Restaurar versión oficial | `.\regravity.exe restore` |
| Ver todas las opciones | `.\regravity.exe --help` |

### Compilar desde el código fuente

Requiere Windows, PowerShell 7, Go 1.26.3+ y la CLI de Wails v2:

```powershell
.\build.ps1 -Target gui
```

Genera `build/bin/regravity.exe`.

## Solución de problemas

- **Aviso de WebView2 faltante al iniciar** — Instala el [entorno de ejecución de Microsoft Edge WebView2](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download) y vuelve a abrir `regravity.exe`.

---

## Descargo de responsabilidad

1. Este es un proyecto no oficial mantenido por la comunidad y no está afiliado, respaldado ni autorizado por Google o Antigravity.
2. Google Antigravity y las marcas comerciales relacionadas pertenecen a sus respectivos propietarios.
3. Esta herramienta se proporciona para uso personal y educativo. Asegúrese de cumplir con los términos de servicio de software relevantes.

[Licencia MIT](./LICENSE)
