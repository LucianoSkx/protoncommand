# AGENTS.md — protoncommand

Instruções para qualquer agente de IA (Codex, Claude Code, Cursor, opencode, Aider)
que mexer neste repositório. Vale para toda tarefa, não só para código.

## O que é o app

Catálogo de opções de inicialização (variáveis de ambiente e flags) para jogos
via Proton, Proton-GE e Proton-CachyOS. O app **não executa o jogo**: ele monta
a string de launch options e copia para a área de transferência. Por isso não
existe `os/exec`, `os.Environ` ou `syscall` no código — e não deve passar a existir.

**A regra de ouro: o valor do app é a descrição técnica estar correta.** Uma
variável descrita errado faz o usuário gastar tempo com uma configuração que não
funciona, ou que funciona e derruba o desempenho. Errar no texto é bug, mesmo que
o app compile e os testes passem.

## Idioma

- Respostas, raciocínio, commits, CHANGELOG e textos de interface em **pt-BR**.
- Strings de catálogo sempre nos **dois idiomas**: `Title`, `Category`, `Compat` e
  `Description` têm `.PT` e `.EN` preenchidos. `TestCommandsValid` falha se algum
  faltar. Texto em inglês é proibido.
- Termos técnicos (nomes de variável, extensão, arquivo, versão, flag) ficam
  **sem traduzir** nos dois idiomas.

## Mapa do repositório

| Arquivo | Papel |
|---|---|
| `commands.go` | Catálogo puro: struct `Command`/`Localized` e `commands() []Command`. Só dados literais, nenhuma lógica. |
| `combination.go` | `buildCombination` (monta a linha final com `%command%`), `splitFields` (tokenizador que respeita aspas), `isWrapper`, `conflicts`. |
| `launchers.go` | Tipo `Launcher` e a lista (`HasCmd` decide se usa `%command%`). |
| `i18n.go` | Dicionários `ptTexts`/`enTexts`. `TestI18nParity` exige chaves iguais nos dois. |
| `main.go` | `main()`, struct `gui` (todo o estado da UI), tema, idioma, favoritos, clipboard. |
| `main_logic_test.go` | Teste único do projeto. 66+ `TestXxx`, stdlib `testing` + `fyne.io/fyne/v2/test`. |
| `CHANGELOG.md` | Seções por versão. O workflow de release extrai as notas daqui. |

Ao adicionar um comando **novo**, não crie arquivo novo: edite o slice literal em
`commands.go`. Ao adicionar um conflito, edite `conflicts()` e adicione a chave
nos dois mapas de `i18n.go`.

## Invariantes do catálogo

Toda entrada precisa, obrigatoriamente:

- terminar com ` %command%`;
- ser única (sem `Command` duplicado);
- ter PT e EN em todos os quatro campos.

`CommandEN` existe para o único caso em que o valor muda entre idiomas
(ex.: `HOST_LC_ALL=pt_BR.UTF-8`). Só use com justificativa.

## Evidência antes de escrever

**Não escreva variável de ambiente, nome de extensão, caminho de arquivo, número
de versão ou afirmação de comportamento de memória.** Antes de colocar no catálogo:

1. Vá na fonte primária: README do upstream, o `*.json.in` da layer Vulkan, o
   código do driver, a discussion do mantenedor, o changelog da release.
2. Confira o **nome exato** da variável. Onde já saiu errado aqui:
   `ENABLE_LAYER_MESA_ANTI-LAG` e `ENABLE_LAYER_MESA_ANTILAG` não funcionam — o
   certo é `ENABLE_LAYER_MESA_ANTI_LAG` (e o par é `DISABLE_LAYER_MESA_ANTI_LAG`),
   que vem do `VkLayer_MESA_anti_lag.json.in` do Mesa. Do mesmo jeito,
   `PROTON_DXVK_LOWLATENCY` não tem underscore no meio, e `DXVK_LOW_LATENCY=1`
   não ativa o fork: as variáveis `DXVK_LOW_LATENCY_*` do DXVK upstream foram
   removidas no dxvk-low-latency 2.7 e viraram opções de `DXVK_CONFIG`. A
   diferença é invisível numa revisão apressada.
3. Registre a fonte dentro da própria descrição (o identificador, o caminho, o
   número da MR) para quem for corrigir no futuro ter onde procurar.
4. Se não conseguiu verificar, **não afirme**. Omitir ou marcar explicitamente
   como não confirmado. Descrição incompleta é aceitável; descrição errada não.

Preferir fonte primária a agregador. Phoronix, Reddit, forum e thread de Discord
servem para achar o caminho, nunca como única prova.

## A descrição precisa documentar a falha, não só o sucesso

A maior parte do valor das descrições está no que **não** funciona. Para cada
opção, verifique e registre o que se aplica:

- **Pré-requisito ausente na máquina do usuário** — camada opcional que nem toda
  distro empacota, build sem a flag, arquivo que precisa existir num path.
- **Quando não tem efeito** — só age em jogos que implementam a extensão, só
  funciona se o jogo chamar a API, não é global.
- **Como verificar que funcionou** — a string exata para procurar no log com
  `PROTON_LOG=1`. Menu aparecer na UI não é prova.
- **Custo ou risco** — o que quebra, o que piora, em que games dá problema.
  Indicadores de FPS que caem com a opção, conflitos com anti-cheat, depreciação
  de caminho em versões novas.
- **Ordem de tentativa**, quando há alternativas que sobem de risco: a mais
  barata primeiro, a que quebra algo por último.

Modele: leia o README do projeto upstream e pergunte "o que o autor avisa que
pode dar errado?" — é isso que falta na descrição, não o que a variável faz.

## Como proteger a documentação

Toda opção de latência/upscaling tem uma entrada em `docsObrigatorias`
(`main_logic_test.go`) com os identificadores técnicos que a descrição precisa
citar. Três regras ao mexer nisso:

- **Ancore em identificador, não em frase.** `VkLayer_MESA_anti_lag.json`,
  `VK_NV_low_latency2`, `NvAPI_D3D_SetSleepMode` são estáveis. "muito menos jogos"
  ou "não é à prova de falhas" quebram na primeira reescrita de redação e
  transformam o teste em Obstáculo.
- A mesma âncora vale para **PT e EN** — é por isso que existe uma lista só.
- Ao adicionar um aviso que importa, **junte a âncora correspondente** ao mesmo
  tempo, senão a documentação e o teste divergem.

Antes de fechar, prove que o teste pega regressão: quebre uma âncora de
propósito, veja falhar com mensagem útil, reverta. Teste que nunca falhou não
foi verificado.

## Revisar antes de dizer que terminou

Dois passos obrigatórios, nesta ordem:

1. **Segunda opinião.** Chame um subagente de revisão com o diff e a lista do
   que você mudou. Ele acha o que você não pensou — nesta repo a revisão
   pegou um `UnselectAll()` que zerava o campo que o guarda seguinte ia ler,
   um `Select(0)` programático disparando a cópia automática, e um
   `fmt.Sprintf` com 3 argumentos para 4 verbos `%s`. Nenhum disso aparece em
   build ou em leitura.
2. **Mutação em cada teste novo.** Neutralize a correção e confirme que o
   teste falha com mensagem útil. Reverter. Um teste que nunca falhou não foi
   verificado, e a mutação que não aplica (string procurada errada) não prova
   nada — confira que ela compilou antes de confiar no "ok".

Onde não dá para testar, registre. O clipboard do driver de teste do Fyne é
descartável (nova instância a cada chamada), então a cópia se observa pelo
`status`, que `copyCurrent` preenche com `"Copiado: "`.

## Verificação

```bash
gofmt -l .          # tem que sair vazio
go vet ./...
go test -cover ./...
go build -o /tmp/protoncommand .
python3 tools/auditar-catalogo.py --offline
```

Tudo limpo antes de considerar pronto. A cobertura está em **94%** e não deve
cair. `go test` sozinho não basta: formatação é o que o CI reprova, e a
auditoria do catálogo é o que impede o tipo de erro que o app mais sofre.

## Auditar o catálogo contra o upstream

`tools/auditar-catalogo.py` existe porque o app documenta variável de outros
projetos e o upstream as remove sem avisar — foi assim que
`DXVK_FRAME_RATE` (removida no DXVK 3.0) e `PROTON_FRAME_RATE` (que nunca
existiu no Proton, só no Proton-EM) ficaram garantindo coisa que não acontece.

```bash
python3 tools/auditar-catalogo.py --offline   # invariantes locais, roda no CI
python3 tools/auditar-catalogo.py --online    # baixa o upstream e compara
```

O `--offline` falha o build e confere: comando terminando em `%command%`,
duplicata, PT/EN completo, env var repetida com valores diferentes, wrapper na
primeira posição. É rápido e não usa rede.

O `--online` baixa 10 fontes (script do Proton 11, README e CHANGELOG do
Proton-CachyOS, README do GE, README e `dxvk.conf` do DXVK, README do
low_latency_layer e dos dois forks, docs do lsfg-vk) e lista as env vars do
catálogo que nenhuma delas menciona, dizendo onde cada uma deveria ser
conferida. Ele **não** falha o build: variável sem menção é quase sempre
normal, porque DXVK_*, MANGOHUD* e LOW_LATENCY_LAYER* vivem em outros
repositórios. Ele também tem uma lista de `REMOVIDAS` que é atualizada à mão
— ao descobrir que o upstream tirou ou renomeou uma variável, acrescente lá
**e** corrija a entrada.

Quando o `--online` acusar algo, não acredite no script: abra a fonte
apontada. Se for mesmo um erro, corrija a descrição e chame a variável de
removida no mapa `REMOVIDAS`, senão o aviso volta.

## Versionamento e release

- A versão **não** está em `FyneApp.toml` (que só tem `[Migrations]`). Ela vem da
  tag: `build-appimage.sh` usa `git describe --tags`, e o workflow recebe
  `VERSION` do `ref_name`.
- O `CHANGELOG.md` precisa ter a seção `# vX.Y.Z` da versão **antes** do push da
  tag — o workflow extrai as release notes desse cabeçalho. Sem a seção, o release
  sai sem notas.
- Atualize a contagem de comandos no `README.md` junto com o catálogo.
- `git push origin <tag>` dispara o workflow **Release AppImage** e publica
  publicamente. Trate como ação externa: só faça com autorização explícita na
  mesma conversa, mesmo que o usuário já tenha autorizado o push anterior.

## Git

- **Nunca commite sem pedido explícito.** Preparar, verificar e mostrar o diff é
  sempre permitido; `git commit` não.
- Não mexa em `git config`, não use `--force`, não pule hooks, não faça `--amend`
  em commit já existente.
- Mensagem de commit em **pt-BR**, no formato do histórico: `tipo: o que mudou e
  por quê`. O corpo explica a decisão, não o diff.
- Examine `git status`, `git diff` e `git log --oneline -10` antes de commitar, e
  versione só o que era para ir.

## Fora do escopo do repositório

Não instale nem altere pacote de sistema, driver, Steam, Proton ou configuração
do usuário para "validar" uma mudança. A validação do app é `gofmt`/`vet`/`test`/
`build`. Se uma afirmação do catálogo só pode ser confirmada numa máquina com GPU
AMD, registre o caminho de verificação na descrição em vez de tentar replicar o
ambiente.

Nunca imprima, copie ou commite token, chave ou credencial.
