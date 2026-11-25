<div align="center">
  <img src="./frontend/dist/appicon.png" width="64" height="64" alt="Ícone do antigravity-i18n">
  <h1>antigravity-i18n</h1>
  <p>Pacotes de idioma, métricas da sessão e cotas de modelos, e controle deslizante de esforço de raciocínio para Google Antigravity no Windows.</p>
  <p>
    <a href="https://github.com/wadewu-ml/antigravity-i18n/releases/latest/download/antigravity-i18n.exe"><strong>Baixar para Windows x64</strong></a>
    ·
    <a href="https://github.com/wadewu-ml/antigravity-i18n/releases">Versões</a>
    ·
    Testado com <strong>Antigravity App 2.35.1</strong>
  </p>
  <p>
    <a href="./README.zh-CN.md">简体中文</a> · <a href="./README.md">English</a> · <a href="./README.zh-Hant.md">繁體中文</a> · <a href="./README.ja.md">日本語</a> · <a href="./README.ko.md">한국어</a> · <a href="./README.de.md">Deutsch</a> · <a href="./README.es.md">Español</a> · <a href="./README.fr.md">Français</a> · <strong>Português (Brasil)</strong> · <a href="./README.ru.md">Русский</a>
  </p>
</div>

<p align="center">
  <a href="./docs/images/language-settings.webp">
    <img src="./docs/images/language-settings.webp" width="560" alt="Janela de configurações com Antigravity v2.35.1, chinês simplificado selecionado e métricas da sessão e controle deslizante ativados">
  </a>
</p>

## Recursos

- **9 idiomas em um único EXE portátil** — Traduz configurações, controles do chat, menus de modelos, menus nativos e avisos de encerramento para chinês simplificado, chinês tradicional, japonês, coreano, alemão, espanhol, francês, português do Brasil e russo sem alterar suas mensagens ou entradas.
- **Métricas da sessão e cotas (opcional)** — Adiciona um anel ao lado do seletor de modelo para consultar cotas restantes, horários locais de redefinição, contexto usado, velocidade média de saída, taxa de acertos do cache e tempo até o primeiro token.
- **Controle de esforço de raciocínio (opcional)** — Adiciona um controle deslizante em cápsula no seletor de modelos para modelos compatíveis.
- **Independente e reversível byte a byte** — Ambos os interruptores funcionam separadamente e também na interface original em inglês. Depois de aplicar, você pode fechar a ferramenta e restaurar a versão oficial quando quiser.

## Início rápido

> [!NOTE]
> A interface da ferramenta está em chinês simplificado. Os rótulos originais aparecem entre parênteses abaixo para facilitar a referência com a captura acima.

### Aplicar configurações

1. Abra `antigravity-i18n.exe`. Se a pasta detectada estiver incorreta, clique em **Alterar local** (`更改位置`) e selecione a pasta com `Antigravity.exe`.
2. Escolha um idioma em **Idioma da interface** (`界面语言`) ou mantenha **English** se quiser apenas os recursos opcionais.
3. Ative **Métricas da sessão** (`会话指标`) e/ou **Controle de esforço de raciocínio** (`思考强度滑块`) abaixo da lista de idiomas conforme desejar.
4. Salve seu trabalho, feche o Antigravity e clique em **Aplicar alterações** (`应用更改`) no canto inferior direito. O Antigravity reabrirá com as novas configurações (ou clique em **Abrir Antigravity** / `打开 Antigravity`) e você poderá fechar a ferramenta.

> [!WARNING]
> Se aparecer o aviso de que o Antigravity ainda está aberto, salve seu trabalho e feche-o normalmente. Usar **Forçar encerramento e tentar novamente** (`强制退出并重试`) interrompe sessões ativas e pode descartar edições não salvas.

### Reaplicar ou restaurar

- **Reaplicar configurações atuais** — Clique em **Reaplicar** (`重新应用`) no canto inferior direito.
- **Restaurar versão oficial** — Feche o Antigravity, abra o menu no canto superior direito (`···`), clique em **Restaurar versão oficial** (`还原官方版本`) e confirme. Isso restaura o inglês original e remove o indicador e o controle deslizante.

> [!IMPORTANT]
> Na primeira aplicação, a ferramenta salva um backup limpo em `resources/app.asar.clean-backup` na pasta de instalação do Antigravity. Mantenha esse arquivo para poder restaurar depois.

## Recursos opcionais

### Métricas da sessão e cotas de modelos

- **Cotas antes da primeira mensagem** — Passe o ponteiro sobre o anel ao lado do seletor de modelo para ver a cota restante por grupo e período (como semanal e 5 horas). Passe sobre uma porcentagem para ver o horário local de redefinição.
- **Métricas ao vivo após uma chamada** — Após a primeira chamada registrada, o anel mostra o uso estimado do contexto e a parte superior exibe **Contexto usado** (tokens usados, capacidade total e porcentagem), **Velocidade média de saída** (incluindo tokens de raciocínio informados), **Taxa de acertos do cache** e **Tempo até o primeiro token**. Clique no anel para fixar ou fechar o painel.
- **Precisão por painel** — Tela dividida, chats laterais e painéis de subagentes leem a chamada mais recente com dados do registro local de sua própria sessão. Um `—` indica valor indisponível.

<p align="center">
  <a href="./docs/images/session-quota-before-first-call.webp">
    <img src="./docs/images/session-quota-before-first-call.webp" width="440" alt="Painel de métricas da sessão e cota restante por modelo acima do anel ao lado do seletor de modelo">
  </a>
</p>

### Controle de esforço de raciocínio

Ative **Controle de esforço de raciocínio** (`思考强度滑块`) e aplique as alterações. No seletor de modelos do Antigravity, modelos compatíveis exibem um controle deslizante acima da opção de uso: arraste para pré-visualizar e solte para aplicar.

<p align="center">
  <a href="./docs/images/thinking-effort-high.webp">
    <img src="./docs/images/thinking-effort-high.webp" height="340" alt="Seletor de modelos com esforço de raciocínio alto">
  </a>
  &nbsp;&nbsp;
  <a href="./docs/images/thinking-effort-max.webp">
    <img src="./docs/images/thinking-effort-max.webp" height="340" alt="Seletor de modelos com esforço de raciocínio máximo">
  </a>
</p>

## Atualizações

- **Após atualizar o Antigravity** — Se a tradução ou os controles desaparecerem, confira a versão compatível no topo desta página e aplique novamente as configurações.
- **Atualizações automáticas da ferramenta** — Ao abrir, verifica o GitHub a cada 24 horas. Um ponto azul no menu superior direito (`···`) indica nova versão; abra o menu para ler as notas e confirmar o download e reinício pelo mesmo EXE. É possível desativar a verificação automática em **Sobre** (`关于`) ou verificar manualmente (`检查更新`). Depois de atualizar, clique em **Reaplicar** (`重新应用`) para enviar as novas traduções ao Antigravity.

## CLI do PowerShell

Abra o PowerShell na pasta que contém `antigravity-i18n.exe`:

| Ação | Comando |
| --- | --- |
| Listar códigos de idioma | `.\antigravity-i18n.exe locales` |
| Aplicar português do Brasil | `.\antigravity-i18n.exe apply --locale pt-BR` |
| Inglês oficial + métricas | `.\antigravity-i18n.exe apply --locale en --session-metrics` |
| Inglês oficial + controle de raciocínio | `.\antigravity-i18n.exe apply --locale en --effort-slider` |
| Ver configurações atuais | `.\antigravity-i18n.exe status` |
| Restaurar versão oficial | `.\antigravity-i18n.exe restore` |
| Ver todas as opções | `.\antigravity-i18n.exe --help` |

## Compilar a partir do código-fonte

Requer Windows, PowerShell 7, Go 1.26.3+ e a CLI do Wails v2:

```powershell
.\build.ps1 -Target gui
```

O executável é salvo em `build/bin/antigravity-i18n.exe`.

## Solução de problemas

- **Aviso de WebView2 ausente ao iniciar** — Instale o [Runtime do Microsoft Edge WebView2](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download) e abra `antigravity-i18n.exe` novamente.

## Isenção de responsabilidade

1. Este é um projeto não oficial mantido pela comunidade e não é afiliado, endossado ou autorizado pelo Google ou pela Antigravity.
2. Google Antigravity e marcas registradas relacionadas pertencem aos seus respectivos proprietários.
3. Esta ferramenta é fornecida para uso pessoal e educacional. Garanta a conformidade com os termos de serviço do software relevantes.

[Licença MIT](./LICENSE)
