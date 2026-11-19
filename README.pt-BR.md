# antigravity-i18n


Ferramenta de localização de IU para o cliente de desktop Google Antigravity.


[简体中文](./README.zh-CN.md) | [English](./README.md) | [繁體中文](./README.zh-Hant.md) | [日本語](./README.ja.md) | [한국어](./README.ko.md) | [Deutsch](./README.de.md) | [Español](./README.es.md) | [Français](./README.fr.md) | **Português (Brasil)** | [Русский](./README.ru.md)

---

**v4.0.0** inclui nove pacotes de idiomas e foi testada com o Antigravity App **2.33.0**.

---

## Principais recursos

- **Executável portátil para Windows x64**. A interface gráfica usa o Microsoft Edge WebView2.
- **Modo duplo GUI e CLI**: abre uma janela limpa da área de trabalho por padrão, enquanto incorpora subcomandos CLI completos (`apply` / `restore` / `status` / `locales`) para scripts e automação.
- **Restauração exata de bytes**: Cria automaticamente um backup original (`app.asar.clean-backup`) antes de qualquer alteração. A restauração recupera um arquivo idêntico byte por byte (verificado pelo SHA-256) ao lançamento oficial.
- **Tradução da interface** principal, dos menus nativos e das caixas de diálogo de encerramento.
- **Métricas opcionais da sessão**. Um anel à direita do seletor de modelo mostra o uso do contexto. Os detalhes mostram o contexto restante, a velocidade de saída, a taxa de acertos do cache e o tempo até o primeiro token. O Antigravity lê os valores do banco de dados local de sessões.
- **Processamento local dos patches** com validação da sintaxe JavaScript e reversão em caso de falha na gravação.

---

## Download e uso

### Download

Baixe o `antigravity-i18n.exe` mais recente na página de lançamentos e execute-o diretamente.

### 1. Interface Gráfica (Recomendado)

Clique duas vezes em `antigravity-i18n.exe`:
1. O aplicativo detecta automaticamente a pasta de instalação do Antigravity (ou clique em ‘Alterar local’ para especificar manualmente).
2. Escolha o idioma da interface e configure as métricas da sessão separadamente. As métricas também funcionam na interface oficial em inglês.
3. Clique em “Aplicar”. Quando o Antigravity fechar, a ferramenta atualiza os recursos e abre o aplicativo novamente. “Restaurar versão oficial” recupera o arquivo original em inglês e desativa o indicador.

Quando for preciso fechar o Antigravity para aplicar ou restaurar, a ferramenta solicita o encerramento normal da instalação selecionada. Se ela continuar aberta após 30 segundos, você pode fechá-la manualmente e tentar de novo ou clicar em “Forçar o encerramento e tentar novamente”.

### 2. Interface de linha de comando

Abra o PowerShell na pasta que contém `antigravity-i18n.exe`.

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

Opções comuns:
- `--app-dir <caminho>`: Especifique manualmente o diretório de instalação do Antigravity (ou defina via `ANTIGRAVITY_APP_DIR`).
- `--locale <code>`: Código do idioma de destino (por exemplo, `zh-CN`, `zh-Hant`, `ja`, `ko`, `de`, `es`, `fr`).
- `--no-restart`: Não reinicie o Antigravity após a correção.
- `apply --session-metrics` ativa o indicador e `--session-metrics=false` o desativa. `apply --locale en --session-metrics` ativa o indicador na interface oficial em inglês.
- `--force` encerra imediatamente o Antigravity na instalação selecionada. Alterações não salvas podem ser perdidas.
- `--timeout <ms>` define o tempo de espera em milissegundos antes de cancelar a operação. O padrão é `30000`.
- `--yes`: Suprime o prompt de confirmação de desligamento.

Cada instalação aceita uma operação de aplicação ou restauração por vez. Se já houver uma operação em andamento, aguarde até que termine e tente novamente. Use `status` para consultar a instalação selecionada e seu backup. O comando retorna um erro se o arquivo não puder ser lido ou validado como um arquivo do Antigravity.

---

## Idiomas incluídos

Todos os pacotes de idiomas são incorporados diretamente no executável:

| Código | Idioma | Cobertura | Notas |
| :--- | :--- | :--- | :--- |
| `zh-CN` | Simplified Chinese (简体中文) | 100% | Cobertura completa da interface, menus nativos e diálogos |
| `zh-Hant` | Traditional Chinese (繁體中文) | 100% | Cobertura completa com terminologia técnica verificada |
| `ja` | Japanese (日本語) | 100% | Cobertura completa com terminologia técnica verificada |
| `ko` | Korean (한국어) | 100% | Cobertura completa com terminologia técnica verificada |
| `de` | German (Deutsch) | 100% | Cobertura completa com terminologia técnica verificada |
| `es` | Spanish (Español) | 100% | Cobertura completa com terminologia técnica verificada |
| `fr` | French (Français) | 100% | Cobertura completa com terminologia técnica verificada |
| `pt-BR` | Portuguese (Brasil) | 100% | Cobertura completa com terminologia técnica verificada |
| `ru` | Russian (Русский) | 100% | Cobertura completa com terminologia técnica verificada |
| `en` | English | - | Build oficial intocado (selecione para restaurar) |

---

## Perguntas frequentes e mecânica

### 1. O que acontece após uma atualização oficial do Antigravity?
As atualizações oficiais substituirão `app.asar`. Basta reabrir esta ferramenta e clicar em ‘Aplicar’ novamente. A ferramenta atualizará automaticamente seu backup limpo usando a nova versão oficial.

### 2. Troca e restauração de idioma
Você pode alternar diretamente entre qualquer idioma; a ferramenta substitui perfeitamente os patches ativos sem exigir uma restauração manual primeiro. Clicar em 'Restaurar Inglês Oficial' retorna o cliente a um estado intocado a qualquer momento.

### 3. Gerenciamento de arquivos de backup
- `app.asar.clean-backup` guarda o arquivo original usado na restauração. Ao aplicar ou restaurar uma instalação ainda sem patch, a ferramenta compara os hashes SHA-256 e atualiza o backup se forem diferentes, mesmo após reinstalar a mesma versão. **Mantenha este arquivo.**
- `app.asar.bak-*` são cópias com data e hora que podem recuperar um backup ausente ou inválido. Quando o aplicativo estiver funcionando e `status` indicar um backup válido, remova as cópias antigas se precisar liberar espaço.

---

## Construindo a partir da fonte

**Pré-requisitos**:
- Go 1.26.3+
- [Wails v2](https://wails.io/) CLI
- PowerShell 7

```powershell
# Clone the repository
git clone https://github.com/wadewu-ml/antigravity-i18n.git
cd antigravity-i18n

# Build both GUI and CLI binaries
.\build.ps1

# Descomente a próxima linha para incluir um teste com um arquivo oficial.
# $env:ANTIGRAVITY_ASAR = "$env:LOCALAPPDATA\Programs\antigravity\resources\app.asar.clean-backup"
go test ./...
```

---

## Isenção de responsabilidade

1. Este é um projeto não oficial mantido pela comunidade e não é afiliado, endossado ou autorizado pelo Google ou pela Antigravity.
2. Google Antigravity e marcas registradas relacionadas pertencem aos seus respectivos proprietários.
3. Esta ferramenta é fornecida para uso pessoal e educacional. Garanta a conformidade com os termos de serviço do software relevantes.

## Licença

[Licença MIT](./LICENSE)
