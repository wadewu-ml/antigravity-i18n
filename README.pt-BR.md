# <img src="./frontend/dist/appicon.png" width="32" height="32" alt=""> antigravity-i18n

Nove pacotes de idiomas traduzem as configurações, os controles do chat, os menus e os avisos de encerramento do Google Antigravity para Windows. Um indicador opcional na caixa de mensagem mostra quanto contexto ainda está disponível e com que velocidade o modelo respondeu. Depois de aplicar as configurações, você pode fechar a ferramenta.

[Baixar para Windows x64](https://github.com/wadewu-ml/antigravity-i18n/releases/latest/download/antigravity-i18n.exe) · [Versões](https://github.com/wadewu-ml/antigravity-i18n/releases)

Compatibilidade verificada com o **Antigravity App 2.35.1**. O EXE portátil inclui os nove pacotes de idiomas.

[简体中文](./README.zh-CN.md) · [English](./README.md) · [繁體中文](./README.zh-Hant.md) · [日本語](./README.ja.md) · [한국어](./README.ko.md) · [Deutsch](./README.de.md) · [Español](./README.es.md) · [Français](./README.fr.md) · **Português (Brasil)** · [Русский](./README.ru.md)

<a href="./docs/images/antigravity-zh-cn.png"><img src="./docs/images/translated-window.webp" width="680" alt="Antigravity em chinês simplificado com os menus traduzidos e o anel ao lado do seletor de modelo"></a>

## Alterar o idioma da interface

Os idiomas disponíveis são chinês simplificado, chinês tradicional, japonês, coreano, alemão, espanhol, francês, português do Brasil e russo. A ferramenta de configuração está em chinês simplificado.

1. Abra `antigravity-i18n.exe`. Se a ferramenta não localizar a instalação correta, clique em “Alterar local” e selecione a pasta que contém `Antigravity.exe`.
2. Escolha o idioma em “Idioma da interface”. Ative “Métricas da sessão” para ver as métricas da sessão. Para usá-las com a interface original em inglês, selecione English.
3. Salve seu trabalho, encerre o Antigravity e clique em “Aplicar alterações” para aplicar as configurações. O Antigravity é reaberto com as novas configurações.

Se a ferramenta pedir que você encerre o Antigravity, a opção “Forçar encerramento e tentar novamente” força o encerramento e tenta de novo. Isso interrompe as sessões em andamento e pode descartar alterações que ainda não foram salvas.

Para ativar o controle deslizante, ligue o interruptor abaixo das métricas da sessão na ferramenta. Os dois recursos podem ser usados separadamente, inclusive com a interface em inglês. Aplique as alterações e abra o seletor de modelos. Se o modelo permitir ajustar o esforço de raciocínio, um controle deslizante aparecerá acima da opção de uso. Arraste e solte para aplicar o nível escolhido.

<a href="./docs/images/language-settings.webp"><img src="./docs/images/language-settings.webp" width="640" alt="Janela de configurações com chinês simplificado selecionado e métricas da sessão ativadas"></a>

O botão “重新应用”, no canto inferior direito, reaplica as configurações atuais.

<p>
  <a href="./docs/images/thinking-effort-high.webp"><img src="./docs/images/thinking-effort-high.webp" width="282" alt="Seletor de modelos com esforço de raciocínio alto"></a>
  <a href="./docs/images/thinking-effort-max.webp"><img src="./docs/images/thinking-effort-max.webp" width="267" alt="Seletor de modelos com esforço de raciocínio máximo"></a>
</p>

## Consultar as métricas da sessão

Passe o ponteiro sobre o anel ao lado do seletor de modelo para ver a cota restante de cada grupo por período. Passe o ponteiro sobre uma porcentagem para ver o horário local da próxima redefinição. Você pode consultar a cota antes de enviar a primeira mensagem. Após a primeira chamada registrada, o anel indica o uso estimado do contexto e o painel mostra os tokens usados e a capacidade total, a velocidade de saída, a taxa de acertos do cache e o tempo até o primeiro token. Clique para fixar o painel e clique novamente para fechá-lo.

<a href="./docs/images/session-metrics-zh-cn.png"><img src="./docs/images/session-metrics-detail.webp" width="560" alt="O painel em chinês simplificado mostra o uso do contexto, os tokens restantes e a velocidade de saída"></a>

[Painel de cotas antes da primeira chamada ao modelo](./docs/images/session-quota-before-first-call.webp)

Na tela dividida, nos chats laterais e nos painéis de subagentes, cada painel usa a chamada ao modelo mais recente com dados no registro local da sua sessão. A contagem de saída inclui os tokens de raciocínio informados pelo modelo. Um `—` indica que aquele dado não está disponível.

## Restaurar ou reaplicar

Para voltar à interface original em inglês e remover o indicador e o controle deslizante, encerre o Antigravity, abra o menu no canto superior direito da janela, clique em “Restaurar versão oficial” (还原官方版本) e confirme com “Restaurar”. A ferramenta cria uma cópia do arquivo original na primeira vez que você aplica as configurações. Mantenha `resources/app.asar.clean-backup` na pasta de instalação do Antigravity para poder restaurá-lo.

Se a tradução ou o indicador desaparecer após uma atualização do Antigravity, confira a versão compatível no início desta página e volte a selecionar e aplicar as configurações. Para reaplicar as configurações que já estão em uso, clique em “Reaplicar”.

Para atualizar a ferramenta, abra o menu no canto superior direito da janela e clique em verificar atualizações (检查更新). A ferramenta consulta o GitHub quando você clica. Confirme o download para abrir a nova versão pelo mesmo caminho do EXE. Para usar as novas traduções no Antigravity, reaplique as configurações.

## Comandos do PowerShell

Abra o PowerShell na pasta que contém `antigravity-i18n.exe`.

| Ação | Comando |
| --- | --- |
| Listar os códigos de idioma | `.\antigravity-i18n.exe locales` |
| Aplicar português do Brasil | `.\antigravity-i18n.exe apply --locale pt-BR` |
| Usar o inglês oficial com métricas | `.\antigravity-i18n.exe apply --locale en --session-metrics` |
| Usar o inglês oficial com o controle de esforço de raciocínio | `.\antigravity-i18n.exe apply --locale en --effort-slider` |
| Ver as configurações atuais | `.\antigravity-i18n.exe status` |
| Restaurar a versão oficial | `.\antigravity-i18n.exe restore` |
| Ver todas as opções | `.\antigravity-i18n.exe --help` |

## Compilar a partir do código-fonte

Para compilar no Windows, instale o PowerShell 7, Go 1.26.3 ou posterior e a CLI do Wails v2. Rode `.\build.ps1 -Target gui` na raiz do repositório. O executável fica em `build/bin/antigravity-i18n.exe`.

## Problemas ao iniciar

Se a ferramenta solicitar o WebView2 ao iniciar, instale o [Runtime do Microsoft Edge WebView2](https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download).

## Isenção de responsabilidade

1. Este é um projeto não oficial mantido pela comunidade e não é afiliado, endossado ou autorizado pelo Google ou pela Antigravity.
2. Google Antigravity e marcas registradas relacionadas pertencem aos seus respectivos proprietários.
3. Esta ferramenta é fornecida para uso pessoal e educacional. Garanta a conformidade com os termos de serviço do software relevantes.

[Licença MIT](./LICENSE)
