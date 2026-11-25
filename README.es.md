# <img src="./frontend/dist/appicon.png" width="32" height="32" alt=""> antigravity-i18n

Nueve paquetes de idiomas traducen los ajustes, los controles del chat, los menús y los avisos de salida de Google Antigravity para Windows. En el cuadro de entrada del chat también puedes ver cuánto contexto queda y a qué velocidad respondió el modelo. Una vez aplicados los ajustes, puedes cerrar la herramienta.

[Descargar para Windows x64](https://github.com/wadewu-ml/antigravity-i18n/releases/latest/download/antigravity-i18n.exe) · [Versiones](https://github.com/wadewu-ml/antigravity-i18n/releases)

Probado con **Antigravity App 2.35.1**.

[简体中文](./README.zh-CN.md) · [English](./README.md) · [繁體中文](./README.zh-Hant.md) · [日本語](./README.ja.md) · [한국어](./README.ko.md) · [Deutsch](./README.de.md) · **Español** · [Français](./README.fr.md) · [Português (Brasil)](./README.pt-BR.md) · [Русский](./README.ru.md)

<a href="./docs/images/antigravity-zh-cn.png"><img src="./docs/images/translated-window.webp" width="680" alt="Antigravity en chino simplificado con menús traducidos y el anillo junto al selector de modelo"></a>

## Cambiar el idioma de la interfaz

Los idiomas disponibles son chino simplificado, chino tradicional, japonés, coreano, alemán, español, francés, portugués de Brasil y ruso. La herramienta de configuración está en chino simplificado.

1. Abre `antigravity-i18n.exe`. Si la herramienta no encuentra la instalación correcta, pulsa «Cambiar ubicación» y selecciona la carpeta que contiene `Antigravity.exe`.
2. Elige un idioma en «Idioma de la interfaz». Activa «Métricas de sesión» para ver las métricas de sesión. Si quieres usarlas con la interfaz original en inglés, selecciona English.
3. Guarda tu trabajo y cierra Antigravity. Pulsa «Aplicar cambios» para aplicar los ajustes. Antigravity se abre de nuevo con los ajustes que has elegido.

Si aparece el aviso de salida, «Forzar el cierre y reintentar» fuerza el cierre de Antigravity y vuelve a intentar la operación. Esto interrumpe las sesiones activas y puede provocar la pérdida de cambios sin guardar.

Para activar el control deslizante, enciende el interruptor situado debajo de las métricas de sesión en la herramienta. Aplica los cambios y abre el selector de modelos. Si el modelo permite elegir el nivel de razonamiento, aparecerá un control deslizante encima de la opción de uso. Arrástralo y suéltalo para aplicar el nivel.

<a href="./docs/images/language-settings.webp"><img src="./docs/images/language-settings.webp" width="640" alt="Ventana de configuración con chino simplificado seleccionado y métricas de sesión activadas"></a>

En la imagen, el botón verde «重新应用» permite volver a aplicar los ajustes.

## Consultar las métricas de sesión

Cuando se registra una llamada al modelo en el chat, aparece un anillo a la derecha del selector de modelo en el cuadro de entrada. La parte rellena indica qué proporción de la ventana de contexto se estima que está en uso. Pasa el puntero sobre el anillo para ver los tokens restantes, la velocidad de salida, la tasa de aciertos de la caché y el tiempo hasta el primer token. Haz clic para fijar los detalles y vuelve a hacer clic para cerrarlos.

<a href="./docs/images/session-metrics-zh-cn.png"><img src="./docs/images/session-metrics-detail.webp" width="560" alt="El panel en chino simplificado muestra el uso del contexto, los tokens restantes y la velocidad de salida"></a>

En las vistas divididas, los chats laterales y los paneles de subagentes, cada panel muestra la llamada más reciente con datos en el registro local de su sesión. El recuento de salida incluye los tokens de razonamiento que indica el modelo. Un `—` indica que ese dato no está disponible.

## Restaurar o volver a aplicar

Para recuperar la interfaz original en inglés y quitar el indicador, cierra Antigravity, pulsa «Restaurar la versión oficial» en la herramienta y confirma con «Restaurar». La primera vez que aplicas los ajustes, la herramienta guarda una copia del archivo original. Conserva `resources/app.asar.clean-backup` en la carpeta de instalación de Antigravity para poder restaurarlo.

Si la traducción o el indicador desaparecen tras actualizar Antigravity, consulta la versión compatible al principio de esta página y vuelve a seleccionar y aplicar los ajustes. Para volver a aplicar los ajustes que ya están activos, pulsa «Volver a aplicar».

Para actualizar la herramienta, pulsa el botón de buscar actualizaciones (检查更新) en la esquina superior derecha de la ventana. Al pulsarlo se consulta GitHub. Confirma la descarga y la herramienta se volverá a abrir desde la misma ruta del EXE. Vuelve a aplicar los ajustes para usar las nuevas traducciones en Antigravity.

## Comandos de PowerShell

Abre PowerShell en la carpeta que contiene `antigravity-i18n.exe`.

| Acción | Comando |
| --- | --- |
| Listar los códigos de idioma | `.\antigravity-i18n.exe locales` |
| Aplicar español | `.\antigravity-i18n.exe apply --locale es` |
| Usar el inglés oficial con métricas | `.\antigravity-i18n.exe apply --locale en --session-metrics` |
| Ver la configuración actual | `.\antigravity-i18n.exe status` |
| Restaurar la versión oficial | `.\antigravity-i18n.exe restore` |
| Ver todas las opciones | `.\antigravity-i18n.exe --help` |

## Compilar desde el código fuente

Para compilar en Windows, instala PowerShell 7, Go 1.26.3 o posterior y la CLI de Wails v2. Ejecuta `.\build.ps1 -Target gui` desde la raíz del repositorio. El resultado se guarda en `build/bin/antigravity-i18n.exe`.

## Problemas al iniciar

Si la herramienta solicita WebView2 al iniciarse, instala el [entorno de ejecución de Microsoft Edge WebView2](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download).

## Descargo de responsabilidad

1. Este es un proyecto no oficial mantenido por la comunidad y no está afiliado, respaldado ni autorizado por Google o Antigravity.
2. Google Antigravity y las marcas comerciales relacionadas pertenecen a sus respectivos propietarios.
3. Esta herramienta se proporciona para uso personal y educativo. Asegúrese de cumplir con los términos de servicio de software relevantes.

[Licencia MIT](./LICENSE)
