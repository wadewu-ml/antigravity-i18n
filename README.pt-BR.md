<div align="center">
  <img src="./frontend/dist/appicon.png" width="68" height="68" alt="Ícone do Regravity">
  <h1>Regravity</h1>
  <p><em>O Antigravity decola; o Regravity traz tudo de volta para o seu idioma e controle.</em></p>
  <p>O utilitário portátil Regravity oferece 18 pacotes de tradução para o <strong>Google Antigravity</strong> no Windows, além de métricas da sessão, cotas de modelos e um controle de esforço de raciocínio. A janela de configurações do Regravity está disponível em 19 idiomas.</p>
  <p>
    <a href="https://github.com/wadewu-ml/regravity/releases/latest/download/regravity.exe"><img src="https://img.shields.io/badge/Download-Latest_(Windows_x64)-0F9D58?style=flat&logo=windows&logoColor=white" alt="Baixar a versão mais recente do Regravity para Windows x64"></a>
    <a href="https://github.com/wadewu-ml/regravity/releases"><img src="https://img.shields.io/badge/Releases-GitHub-18181B?style=flat&logo=github" alt="Versões no GitHub"></a>
    <img src="https://img.shields.io/badge/Antigravity-2.36.0_tested-18181B?style=flat" alt="Testado com Antigravity App 2.36.0">
    <a href="./LICENSE"><img src="https://img.shields.io/badge/License-MIT-18181B?style=flat" alt="Licença MIT"></a>
  </p>
  <p>
    <sub>
      <a href="./README.zh-CN.md">简体中文</a> · <a href="./README.md">English</a> · <a href="./README.zh-Hant.md">繁體中文</a> · <a href="./README.ja.md">日本語</a> · <a href="./README.ko.md">한국어</a> · <a href="./README.de.md">Deutsch</a> · <a href="./README.es.md">Español</a> · <a href="./README.fr.md">Français</a> · <strong>Português (Brasil)</strong> · <a href="./README.ru.md">Русский</a> · <a href="./README.tr.md">Türkçe</a> · <a href="./README.vi.md">Tiếng Việt</a> · <a href="./README.it.md">Italiano</a> · <a href="./README.ar.md">العربية</a> · <a href="./README.id.md">Bahasa Indonesia</a> · <a href="./README.pl.md">Polski</a> · <a href="./README.cs.md">Čeština</a> · <a href="./README.hu.md">Magyar</a> · <a href="./README.bg.md">Български</a>
    </sub>
  </p>
</div>

<br>

<p align="center">
  <a href="./docs/images/language-settings.webp">
    <img src="./docs/images/language-settings.webp" width="580" alt="Janela de configurações com Antigravity v2.35.2, chinês simplificado selecionado e métricas da sessão e controle deslizante ativados">
  </a>
  <br>
  <sub>Janela de configuração do Regravity com o Antigravity 2.35.2.</sub>
</p>

## O que está incluído

- **18 pacotes de tradução, sem tocar nos seus chats** — Traduz configurações, controles do chat, menus de modelos, menus nativos e avisos de encerramento (`zh-CN`, `zh-Hant`, `ja`, `ko`, `de`, `es`, `fr`, `pt-BR`, `ru`, `tr`, `vi`, `it`, `ar`, `id`, `pl`, `cs`, `hu`, `bg`) sem alterar suas mensagens ou entradas.
- **Métricas da sessão e cotas** *(opcional)* — Adiciona um anel ao lado do seletor de modelo com cotas restantes, horários locais de redefinição, contexto usado, velocidade média de saída, taxa de cache e tempo até o primeiro token.
- **Controle de esforço de raciocínio** *(opcional)* — Integra um controle deslizante em cápsula no seletor de modelos para modelos compatíveis.
- **Modular e reversível byte a byte** — Ambos os interruptores funcionam separadamente e também com a interface original em inglês. Restaure a versão oficial a partir de um backup íntegro da mesma versão do Antigravity quando quiser.

## Início rápido

> [!TIP]
> Os rótulos em chinês entre parênteses correspondem à captura de tela.

Escolha o idioma da ferramenta no canto superior direito.

### Aplicar configurações

1. Abra `regravity.exe`. Se a pasta detectada estiver incorreta, clique em **Alterar local** (`更改位置`) e selecione a pasta com `Antigravity.exe`.
2. Escolha um idioma em **Idioma da interface** (`界面语言`) ou mantenha **English** se quiser apenas os recursos opcionais.
3. Ative **Métricas da sessão** (`会话指标`) e/ou **Controle de esforço de raciocínio** (`思考强度滑块`) abaixo da lista de idiomas.
4. Salve seu trabalho, feche o Antigravity e clique em **Aplicar alterações** (`应用更改`) no canto inferior direito. O Antigravity reabrirá com as novas configurações (ou clique em **Abrir Antigravity** / `打开 Antigravity`), e você pode fechar a ferramenta.

> [!WARNING]
> Se aparecer o aviso de que o Antigravity ainda está aberto, salve seu trabalho e feche-o normalmente. **Forçar encerramento e tentar novamente** (`强制退出并重试`) interrompe sessões ativas e pode descartar edições não salvas.

### Reaplicar ou restaurar

- **Reaplicar** — Após atualizar a ferramenta, clique em **Reaplicar** (`重新应用`) para usar as novas traduções.
- **Restaurar versão oficial** — Feche o Antigravity, abra o menu no canto superior direito (`···`), clique em **Restaurar versão oficial** (`还原官方版本`) e confirme.

> [!IMPORTANT]
> Na primeira aplicação, a ferramenta salva um backup limpo em `resources/app.asar.clean-backup` na pasta de instalação do Antigravity. Mantenha esse arquivo para poder restaurar depois.

## Recursos opcionais

### Métricas da sessão e cotas de modelos

- **Antes da primeira mensagem** — Passe o ponteiro sobre o anel ao lado do seletor de modelo para ver a cota restante por grupo e período (semanal, 5 horas etc.); passe sobre uma porcentagem para ver o horário local de redefinição.
- **Após uma chamada ao modelo** — O anel indica o uso estimado do contexto e a parte superior exibe **Contexto usado** (tokens, capacidade total, proporção), **Velocidade média de saída** (incluindo tokens de raciocínio), **Taxa de acertos do cache** e **Tempo até o primeiro token**. Clique no anel para fixar ou fechar.
- **Precisão por painel** — Tela dividida, chats laterais e subagentes leem a chamada mais recente com dados da própria sessão local. Um `—` indica valor indisponível.

<p align="center">
  <a href="./docs/images/session-quota-before-first-call.webp">
    <img src="./docs/images/session-quota-before-first-call.webp" width="420" alt="Painel de métricas da sessão e cota restante por modelo acima do anel ao lado do seletor de modelo">
  </a>
  <br>
  <sub>As cotas aparecem imediatamente; as métricas da sessão são preenchidas após a primeira chamada ao modelo.</sub>
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
  <br>
  <sub>Nível padrão (Alto) e nível máximo em destaque diretamente no menu de modelos.</sub>
</p>

## Atualizações e CLI

### Atualizações

- **Após atualizar o Antigravity** — Se a tradução ou os controles desaparecerem, confira a versão compatível no topo desta página e aplique novamente as configurações.
- **Atualizações automáticas** — Ao abrir, usa o resultado da verificação anterior para decidir se consulta o GitHub e aguarda ao menos 24 horas após um sucesso. Um ponto azul no menu superior direito (`···`) indica nova versão. Pode ser desativada em **Sobre** (`关于`), mantendo a verificação manual (`检查更新`). Depois de atualizar, clique em **Reaplicar** (`重新应用`).

Se você usa a v4.2.3 ou anterior, baixe e execute [`regravity.exe`](https://github.com/wadewu-ml/regravity/releases/latest/download/regravity.exe) manualmente. Suas configurações serão mantidas.

### CLI do PowerShell

Abra o PowerShell na pasta que contém `regravity.exe`:

| Ação | Comando |
| --- | --- |
| Listar códigos de idioma | `.\regravity.exe locales` |
| Aplicar português do Brasil | `.\regravity.exe apply --locale pt-BR` |
| Inglês oficial + métricas | `.\regravity.exe apply --locale en --session-metrics` |
| Inglês oficial + controle de raciocínio | `.\regravity.exe apply --locale en --effort-slider` |
| Ver configurações atuais | `.\regravity.exe status` |
| Restaurar versão oficial | `.\regravity.exe restore` |
| Ver todas as opções | `.\regravity.exe --help` |

### Compilar a partir do código-fonte

Requer Windows, PowerShell 7, Go 1.26.3+ e a CLI do Wails v2:

```powershell
.\build.ps1 -Target gui
```

Gera `build/bin/regravity.exe`.

## Solução de problemas

- **Aviso de WebView2 ausente ao iniciar** — Instale o [Runtime do Microsoft Edge WebView2](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download) e abra `regravity.exe` novamente.

---

## Isenção de responsabilidade

1. Este é um projeto não oficial mantido pela comunidade e não é afiliado, endossado ou autorizado pelo Google ou pela Antigravity.
2. Google Antigravity e marcas registradas relacionadas pertencem aos seus respectivos proprietários.
3. Esta ferramenta é fornecida para uso pessoal e educacional. Garanta a conformidade com os termos de serviço do software relevantes.

[Licença MIT](./LICENSE)
