# <img src="./frontend/dist/appicon.png" width="32" height="32" alt=""> antigravity-i18n

Nove pacotes de idiomas traduzem as configurações, os controles do chat, os menus e os avisos de encerramento do Google Antigravity para Windows. Um indicador opcional na caixa de mensagem mostra quanto contexto ainda está disponível e com que velocidade o modelo respondeu. Depois de aplicar as configurações, você pode fechar a ferramenta.

[Baixar para Windows x64](https://github.com/wadewu-ml/antigravity-i18n/releases/latest/download/antigravity-i18n.exe) · [Versões](https://github.com/wadewu-ml/antigravity-i18n/releases)

Compatibilidade verificada com o **Antigravity App 2.35.1**.

[简体中文](./README.zh-CN.md) · [English](./README.md) · [繁體中文](./README.zh-Hant.md) · [日本語](./README.ja.md) · [한국어](./README.ko.md) · [Deutsch](./README.de.md) · [Español](./README.es.md) · [Français](./README.fr.md) · **Português (Brasil)** · [Русский](./README.ru.md)

<a href="./docs/images/antigravity-zh-cn.png"><img src="./docs/images/translated-window.webp" width="680" alt="Antigravity em chinês simplificado com os menus traduzidos e o anel ao lado do seletor de modelo"></a>

## Alterar o idioma da interface

Os idiomas disponíveis são chinês simplificado, chinês tradicional, japonês, coreano, alemão, espanhol, francês, português do Brasil e russo. A ferramenta de configuração está em chinês simplificado.

1. Abra `antigravity-i18n.exe`. Se a ferramenta não localizar a instalação correta, clique em “Alterar local” e selecione a pasta que contém `Antigravity.exe`.
2. Escolha o idioma em “Idioma da interface”. Ative “Métricas da sessão” para ver as métricas da sessão. Para usá-las com a interface original em inglês, selecione English.
3. Salve seu trabalho, encerre o Antigravity e clique em “Aplicar alterações” para aplicar as configurações. O Antigravity é reaberto com as novas configurações.

Se a ferramenta pedir que você encerre o Antigravity, a opção “Forçar encerramento e tentar novamente” força o encerramento e tenta de novo. Isso interrompe as sessões em andamento e pode descartar alterações que ainda não foram salvas.

<a href="./docs/images/language-settings.webp"><img src="./docs/images/language-settings.webp" width="640" alt="Janela de configurações com chinês simplificado selecionado e métricas da sessão ativadas"></a>

Na imagem, o botão verde “重新应用” reaplica as configurações.

## Consultar as métricas da sessão

Assim que uma chamada ao modelo é registrada no chat, aparece um anel à direita do seletor de modelo na caixa de mensagem. O preenchimento do anel mostra a proporção estimada da janela de contexto em uso. Passe o ponteiro sobre ele para ver os tokens restantes, a velocidade de saída, a taxa de acertos do cache e o tempo até o primeiro token. Clique para fixar os detalhes e clique de novo para fechá-los.

<a href="./docs/images/session-metrics-zh-cn.png"><img src="./docs/images/session-metrics-detail.webp" width="560" alt="O painel em chinês simplificado mostra o uso do contexto, os tokens restantes e a velocidade de saída"></a>

Na tela dividida, nos chats laterais e nos painéis de subagentes, cada painel usa a chamada ao modelo mais recente com dados no registro local da sua sessão. A contagem de saída inclui os tokens de raciocínio informados pelo modelo. Um `—` indica que aquele dado não está disponível.

## Restaurar ou reaplicar

Para voltar à interface original em inglês e remover o indicador, encerre o Antigravity, clique em “Restaurar versão oficial” na ferramenta e confirme com “Restaurar”. A ferramenta cria uma cópia do arquivo original na primeira vez que você aplica as configurações. Mantenha `resources/app.asar.clean-backup` na pasta de instalação do Antigravity para poder restaurá-lo.

Se a tradução ou o indicador desaparecer após uma atualização do Antigravity, confira a versão compatível no início desta página e volte a selecionar e aplicar as configurações. Para reaplicar as configurações que já estão em uso, clique em “Reaplicar”.

## Comandos do PowerShell

Abra o PowerShell na pasta que contém `antigravity-i18n.exe`.

| Ação | Comando |
| --- | --- |
| Listar os códigos de idioma | `.\antigravity-i18n.exe locales` |
| Aplicar português do Brasil | `.\antigravity-i18n.exe apply --locale pt-BR` |
| Usar o inglês oficial com métricas | `.\antigravity-i18n.exe apply --locale en --session-metrics` |
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
