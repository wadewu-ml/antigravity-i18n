# antigravity-i18n


Herramienta de localización de UI para el cliente de escritorio Google Antigravity.


[简体中文](./README.zh-CN.md) | [English](./README.md) | [繁體中文](./README.zh-Hant.md) | [日本語](./README.ja.md) | [한국어](./README.ko.md) | [Deutsch](./README.de.md) | **Español** | [Français](./README.fr.md) | [Português (Brasil)](./README.pt-BR.md) | [Русский](./README.ru.md)

---

Un ejecutable independiente que incluye una interfaz gráfica moderna, un motor Go asar puro y paquetes de idiomas completos. Se ejecuta de inmediato sin dependencias de Node.js o bibliotecas de tiempo de ejecución, y admite la restauración exacta de bytes a la compilación oficial en cualquier momento.

---

## Características clave

- **Binario autónomo único**: ejecutable independiente de ~12 MB; no se requiere tiempo de ejecución de Node.js, npm o Python.
- **Modo dual GUI y CLI**: abre una ventana de escritorio limpia de forma predeterminada, al tiempo que incorpora subcomandos CLI completos (`apply`/`restore`/`status`/`locales`) para secuencias de comandos y automatización.
- **Restauración exacta en bytes**: crea automáticamente una copia de seguridad impecable (`app.asar.clean-backup`) antes de cualquier cambio. La restauración recupera un archivo idéntico byte por byte (verificado SHA-256) a la versión oficial.
- **Diseño no invasivo**: localiza solo la interfaz de usuario del shell, los menús nativos y los cuadros de diálogo para salir. Los editores de código (Mónaco), los terminales integrados (xterm) y los flujos de conversaciones de IA permanecen estrictamente intactos.
- **Sin conexión y confiable**: Cero actividad de red, cero recopilación de datos. Equipado con validación inicial de sintaxis de JavaScript y reversión transaccional en caso de errores inesperados.

---

## Descarga y uso

### Descargar

Descargue la última versión `antigravity-i18n.exe` de la página de Versiones y ejecútela directamente.

### 1. Interfaz gráfica (recomendada)

Haga doble clic en `antigravity-i18n.exe`:
1. La aplicación detecta automáticamente la carpeta de instalación de Antigravity (o haga clic en "Cambiar ubicación" para especificarla manualmente).
2. Haga clic en su idioma preferido (al seleccionar inglés se restaura la versión oficial).
3. Haga clic en «Aplicar» y confirme. Una vez cerrado Antigravity, la herramienta actualiza los recursos y lo vuelve a iniciar.

Si es necesario cerrar Antigravity para aplicar o restaurar, la herramienta solicita el cierre normal de la instalación seleccionada. La operación se cancela si sigue abierta tras 30 segundos. Para cerrarla inmediatamente, seleccione la opción de cierre forzado en el cuadro de confirmación.

### 2. Interfaz de línea de comando

Abra PowerShell en la carpeta que contiene `antigravity-i18n.exe`.

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

Opciones comunes:
- `--app-dir <ruta>`: especifique manualmente el directorio de instalación de Antigravity (o configúrelo mediante `ANTIGRAVITY_APP_DIR`).
- `--locale <código>`: código del idioma de destino (por ejemplo, `zh-CN`, `zh-Hant`, `ja`, `ko`, `de`, `es`, `fr`).
- `--no-restart`: No reinicie Antigravity después de aplicar el parche.
- `--force` cierra de inmediato la instalación seleccionada. Puede perder los cambios sin guardar.
- `--timeout <ms>` fija el tiempo de espera en milisegundos antes de cancelar la operación. El valor predeterminado es `30000`.
- `--yes`: suprime el mensaje de confirmación de apagado.

Cada instalación admite una sola operación de aplicación o restauración a la vez. Si ya hay una operación en curso, espere a que termine y vuelva a intentarlo. Consulte la instalación seleccionada y su copia de seguridad con `status`. Si el archivo no se puede leer o no se puede verificar como archivo de Antigravity, el comando devuelve un error.

---

## Idiomas incluidos

Todos los paquetes de idiomas están integrados directamente en el ejecutable:

| Código | Idioma | Cobertura | Notas |
| :--- | :--- | :--- | :--- |
| `zh-CN` | Simplified Chinese (简体中文) | 100% | Cobertura completa de la interfaz, menús nativos y cuadros de diálogo |
| `zh-Hant` | Traditional Chinese (繁體中文) | 100% | Cobertura completa con terminología técnica verificada |
| `ja` | Japanese (日本語) | 100% | Cobertura completa con terminología técnica verificada |
| `ko` | Korean (한국어) | 100% | Cobertura completa con terminología técnica verificada |
| `de` | German (Deutsch) | 100% | Cobertura completa con terminología técnica verificada |
| `es` | Spanish (Español) | 100% | Cobertura completa con terminología técnica verificada |
| `fr` | French (Français) | 100% | Cobertura completa con terminología técnica verificada |
| `pt-BR` | Portuguese (Brasil) | 100% | Cobertura completa con terminología técnica verificada |
| `ru` | Russian (Русский) | 100% | Cobertura completa con terminología técnica verificada |
| `en` | English | - | Compilación oficial original (seleccione para restaurar) |

---

## Preguntas frecuentes y mecánica

### 1. ¿Qué sucede después de una actualización oficial de Antigravity?
Las actualizaciones oficiales sobrescribirán `app.asar`. Simplemente vuelva a abrir esta herramienta y haga clic en "Aplicar" nuevamente. La herramienta actualizará automáticamente su copia de seguridad limpia utilizando la nueva versión oficial.

### 2. Cambio y restauración de idioma
Puedes cambiar directamente entre cualquier idioma; la herramienta reemplaza perfectamente los parches activos sin necesidad de una restauración manual primero. Al hacer clic en 'Restaurar inglés oficial', el cliente vuelve a un estado intacto en cualquier momento.

### 3. Gestión de archivos de respaldo
- `app.asar.clean-backup` guarda el archivo original para la restauración. Al aplicar o restaurar una instalación sin parches, la herramienta compara los valores SHA-256 y actualiza la copia si difieren, incluso tras reinstalar la misma versión. **Conserve este archivo.**
- `app.asar.bak-*` son copias con marca de tiempo que permiten recuperar una copia de seguridad ausente o no válida. Si el cliente funciona y `status` indica que la copia es válida, puede eliminar las copias antiguas para liberar espacio.

---

## Construyendo desde la fuente

**Requisitos previos**:
- Go 1.26.3+
- [Wails v2](https://wails.io/) CLI
- PowerShell 7

```powershell
# Clone the repository
git clone https://github.com/wadewu-ml/antigravity-i18n.git
cd antigravity-i18n

# Build both GUI and CLI binaries
.\build.ps1

# Quite el comentario de la línea siguiente para incluir una prueba con un archivo oficial.
# $env:ANTIGRAVITY_ASAR = "$env:LOCALAPPDATA\Programs\antigravity\resources\app.asar.clean-backup"
go test ./...
```

---

## Descargo de responsabilidad

1. Este es un proyecto no oficial mantenido por la comunidad y no está afiliado, respaldado ni autorizado por Google o Antigravity.
2. Google Antigravity y las marcas comerciales relacionadas pertenecen a sus respectivos propietarios.
3. Esta herramienta se proporciona para uso personal y educativo. Asegúrese de cumplir con los términos de servicio de software relevantes.

## Licencia

[Licencia MIT](./LICENSE)
