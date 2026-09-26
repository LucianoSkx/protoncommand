# v0.6.7

## 🔎 Duas fontes a mais, e o que elas revelaram

- **13 fontes agora.** A discussion #2 do `vkd3d-low-latency` e a issue #2 do `low_latency_layer` entraram na lista do `--online`. Nenhuma das duas está no README, e era por isso que as afirmações abaixo pareciam sem fonte: eu as tinha rebaixado ou não tinha conferido.
- **A prioridade do `VKD3D_FRAME_RATE` voltou ao catálogo.** Na v0.6.5 eu tinha substituído "tem prioridade sobre o fps cap do Reflex, e o mantenedor é explícito" por "o upstream não afirma". O mantenedor afirma, textualmente, na discussion #2: *"VKD3D_FRAME_RATE takes priority over the Reflex fps cap if both are set"*. A parte que continua sem fonte — prioridade sobre o limite que o jogo impõe — segue registrada como tal.
- **A estimativa de 20-30% é do mantenedor e é um palpite**, não uma medição: *"my guess would be 20-30%?"*. O texto agora diz isso. A lista de jogos verificados também era imprecisa: a discussion traz 10 na lista de "explicitly verified" e o `Witchfire` aparece só em comentário — o catálogo tratava os 11 como igualmente verificados.
- **O defeito do `LOW_LATENCY_LAYER_SPOOF_NVIDIA` estava no motivo errado.** O catálogo dizia que ele é preferível ao `PROTON_FORCE_NVAPI` "porque não passa pelo `WINE_HIDE_AMD_GPU`". A fonte dá outro motivo: *"prefer `DXVK_CONFIG="dxgi.hideAmdGpu = True"`, as this option is known to break Proton's FSR4 upgrade path"* — ou seja, o mesmo defeito do FSR4, e o texto admitia isso na frase seguinte, o que o tornava autocontraditório. A âncora em `docsObrigatorias` que travava o texto errado foi trocada.
- **O "FSR4 4.1.1+ cai para FSR3" e o custo em RDNA3 foram marcados como não confirmados.** Nenhuma das 13 fontes contém 4.1.1, FSR3 como fallback nem o custo em RDNA3. O que a issue #2 documenta é o FSR4 sumindo do menu e ficando inutilizável com a camada ativa, mais a combinação que funciona (FSR4 + Reflex + `dxgi.hideAmdGpu`). O texto passou a dizer isso e a apontar, entre parênteses, que as versões e o custo vêm de fóruns e não foram confirmados.

# v0.6.6

## 🔧 Ferramenta de auditoria

- **O parser lia o literal cru da descrição.** O padrão `"(.*?)"` não truncava — como o campo exigia `", EN: "`, ele avançava até a fronteira real — mas devolvia `"` em vez de aspa e `\n` em vez de quebra. Os checks de texto liam aquilo: depois da reestruturação em blocos, nenhum dos três invariantes de texto enxergava as quebras nem as aspas reais. Agora casa o literal inteiro e devolve o valor efetivo, com `desescapar_texto` convertendo `\n` e `\t`.
- **Cinco âncoras novas travam o texto que vem depois de uma aspa escapada** (`PROTON_VKD3D_LOWLATENCY` → "Witchfire", `PROTON_FORCE_NVAPI` e `LOW_LATENCY_LAYER_SPOOF_NVIDIA` → "dxgi.hideAmdGpu", `DXVK_FRAME_RATE` → "d3d9.maxFrameRate", `PROTON_USE_OPTISCALER` → "Upscalers.Dx12Upscaler=dlss"). São as entradas que citam `DXVK_CONFIG="..."` ou a string de verificação do log. Nenhum outro check pega quando esse trecho some: a descrição continua parecendo completa.
- **11ª fonte: as release notes do fork `dxvk-low-latency`.** O README do fork é praticamente o do DXVK upstream e não menciona `DXVK_FRAME_PACE`, nem o teto de 5% abaixo do refresh, nem o `DXVK_HUD=latencydetails` — tudo o que a descrição do `DXVK_FRAME_RATE` afirma. Sem essa fonte, quem roda `--online` não conseguia checar nada desse bloco.

# v0.6.5

## 🐛 Aviso de conflito que faltava

- **Spoofing de GPU junto com upgrade FSR4 era montado em silêncio.** O README do low_latency_layer diz, textualmente, que `PROTON_FORCE_NVAPI` e `LOW_LATENCY_LAYER_SPOOF_NVIDIA` quebram o caminho de upgrade do FSR 4 — "known to break", não é caso raro — e o motivo está no script do Proton: `PROTON_FORCE_NVAPI` define `WINE_HIDE_AMD_GPU=1` sem condição, e é justamente esconder a AMD que impede o upgrade. O app já avisava do conflito irmão (`DISABLE_LAYER_MESA_ANTI_LAG` × `PROTON_FSR4_UPGRADE`) e não deste: as seis combinações saíam com zero aviso, e em RDNA3 o jogo cai para FSR 3.1 sem o usuário entender por quê. A regra cobre as três variáveis de upgrade, incluindo `PROTON_FFX4_UPGRADE`, que é o nome atual no CachyOS 11+ e o caminho que a primeira versão da regra deixava passar. Os dois spoofings são somados em vez de sobrescritos, senão o segundo sumia do aviso.

## 📝 Descrições reestruturadas

Dez descrições longas (de 756 a 1496 caracteres) eram um parágrafo único com várias ideias espremidas — a pior tinha seis, entre elas "não use no D3D12" e "o que fazer em vez disso", sem separação. Agora são blocos rotulados, com o conteúdo verificado preservado por bloco (nenhum identificador técnico foi perdido na reescrita, checado por script) e paridade de blocos entre PT e EN.

- O `LSFGVK_PROFILE` tinha o bloco de migração da v2.0.0 (Vulkan 1.2, FP16 2:1, incompatibilidade com a v1) **só em PT** — o usuário inglês perdia a informação de migração.
- Três erros de fato foram corrigidos no caminho: `default_cpu_limit` tem **20** jogos, não 19; a implementação nativa de `VK_AMD_anti_lag` no RADV (MR 42048) está **em aberto**, não mergeada; e o teto de 5% do modo VRR é sobrescrito **por `DXVK_FRAME_RATE`**, não ajustado dentro do próprio modo.
- A alternativa ao `PROTON_FORCE_NVAPI` estava com `DXVK_NVAPI_ALLOW_OTHER_DRIVERS=1` **dentro** das aspas do `DXVK_CONFIG` — o DXVK separa opções por `;`, então aquilo virava uma opção inválida e a variável nunca era definida. A variável fica fora das aspas.

## 📝 Afirmação sem fonte, corrigida

Duas descrições atribuíam ao mantenedor dos forks uma recomendação que não está no README de nenhum deles. Fui verificar na fonte primária antes de mexer, e o que a fonte diz é outra coisa:

- `DXVK_FRAME_RATE` dizia, nos dois idiomas, que combinar com o `DXVK_FRAME_PACE` "sobrepõe dois limitadores, que é justamente o que piora a latência — o mantenedor do fork é explícito em usar um só". O release 3.1.1 do `dxvk-low-latency` diz outra coisa: *"The fps cap is set by default to 5% below maximum refresh rate and can be overridden by setting it manually"*. Existe teto interno e ele pode ser sobrescrito — mas o upstream **não** afirma nada sobre somar um segundo limitador. O texto passou a trazer o fato verificável e a dizer explicitamente que a soma não é coberta pelo upstream.
- `VKD3D_FRAME_RATE` dizia que o limitador "tem prioridade sobre o fps cap do Reflex e sobre o in-game, e o mantenedor é explícito". O README do `vkd3d-low-latency` só diz *"FPS limiting is fully integrated into the frame pacing logic"* — integrated, sim; com prioridade declarada, não. A afirmação órfã que sobrava no `PROTON_VKD3D_LOWLATENCY` vizinho foi corrigida pelo mesmo motivo.

## ⚠️ Dívida registrada

14 entradas de **Upscaling** não têm âncora em `docsObrigatorias`, contra 7 de 7 de Latência. Várias (integer scaling, FSR strength) não citam identificador técnico na descrição, então "ancorar" seria fixar frase — o que o AGENTS.md proíbe. Enquanto isso não for resolvido reescrevendo as descrições com fonte, `upscalingSemAncora` congela a lista: entrada nova sem âncora reprova o teste.

# v0.6.4

## 🐛 Análise do estado atual

Uma revisão independente do repositório inteiro (não só de um diff) encontrou seis defeitos. Três eram de conteúdo — a regra de ouro do app — e três de teste.

- **A descrição de `PROTON_USE_X11_EXCLUSIVE` citava `PROTON_ENABLE_WINED3D`, variável que não existe em nenhuma das 10 fontes auditadas.** Existe só `PROTON_USE_WINED3D` (renderizador OpenGL) e `PROTON_ENABLE_WAYLAND`. O texto agora diz o que de fato acontece: sob Wine-Wayland o comando isola o `winex11.drv` no executável nomeado, e sem `PROTON_ENABLE_WAYLAND` o jogo inteiro já está no `winex11.drv`, então o comando não muda nada — a pré-condição, que faltava.
- **A descrição de `PROTON_DISABLE_NVAPI` mandava usar `PROTON_ENABLE_NVAPI=1`**, comando que o próprio catálogo removeu e que `TestNoObsoleteCommands` proíbe. A referência saiu, e o texto passou a registrar a pré-condição verificada no script do Proton 11: `DXVK_ENABLE_NVAPI=1` é definido por padrão, a menos que o compat config traga `disablenvapi`.
- **Texto PT corrompido (`se*''confundir`) e palavra inglesa no PT (`controlling`).** Nenhum teste pegava as duas — e nem pegaria, porque o texto é lido por gente. O `--offline` do auditor ganhou três invariantes: resto de edição em qualquer campo, verbo inglês na descrição PT e palavra portuguesa na descrição EN (as três com zero falso positivo hoje, calibradas contra o catálogo inteiro).
- **`TestNoConflictSameValue` era tautológico.** Marcava `selected[idxOf(...)]=true` para `WINE_ESYNC` e `WINEFSYNC`, comandos já removidos do catálogo — `idxOf` devolvia `-1` e `conflicts()` varre `g.all`, então ninguém lia a chave. O teste passava mesmo com o detector quebrado. Virou catálogo sintético, e a mutação (`len(vals[k]) >= 1`) agora o derruba.
- **`TestConflictSemConflitoDeAspas` era cópia literal de `TestConflictDuplicate`**, com o nome prometendo o contrário do corpo: nenhum teste de aspas existia no detector de conflitos. Virou `TestConflictValoresComAspas`, que quebra se o tokenizador ceder àspa.
- **O import de favoritos não migrava a chave antiga.** `carregarFavs` já migrava `Command\x00Título` desde a v0.6.3, mas `hasKnownFavKey` e `filterNewFavKey` continuavam rejeitando: um backup exportado na v0.6.2 era lido como arquivo inválido, e numa lista mista as chaves antigas sumiam sem aviso. A migração foi centralizada em `normalizeFavKey` e aplicada nos três pontos, com teste para chave nova, chave conhecida e chave já favorita.

## 📋 Dívida registrada

14 entradas de **Upscaling** não têm âncora em `docsObrigatorias`, contra 7 de 7 de Latência. Várias (integer scaling, FSR strength) não citam nenhum identificador técnico na descrição, então "ancorar" seria fixar frase — o que o AGENTS.md proíbe. Enquanto isso não é resolvido reescrevendo as descrições com verificação de fonte, `upscalingSemAncora` congela a lista: entrada nova de upscaling sem âncora reprova o teste, e entrada que ganhar âncora precisa sair da lista.

# v0.6.3

## 🐛 Correções da segunda rodada (revisão das correções anteriores)

A revisão do diff anterior encontrou defeito em quatro das sete correções, e um deles era o pior tipo: **teste verde com o código errado**.

- **O aviso de wrapper duplicado saía com `%!s(MISSING)`.** A mensagem tinha quatro verbos `%s` e a chamada passava três argumentos — o quarto era justamente o nome do segundo wrapper, que é a informação principal do aviso. O texto agora tem três verbos e foi verificado em execução.
- **Todos os favoritos de quem usou a v0.6.2 ou anterior eram apagados em silêncio no upgrade.** A chave nova é só o `Command`, mas o que está no disco de todo mundo é `Command\x00Título`. A poda jogava tudo fora. Agora a entrada antiga é migrada cortando no `\x00` antes da poda.
- **A correção da troca de idioma não existia no app real.** O `pos := g.selID` era lido **depois** de `applyLang()`, e `applyLang` chama `catSel.SetSelected`, que dispara o handler de filtro — que já zera `selID` e faz `Select(0)`. O teste passava porque o harness ligava o `catSel` com um no-op, ou seja, media um app que não existe. A posição agora é lida antes de `applyLang`, e o `initTestGUI` replica os handlers reais de `build()`.
- **Abrir o app sobrescrevia o clipboard.** `build()` lia a preferência "copiar ao clicar" e depois fazia `Select(0)`, que disparava a cópia automática. A preferência passou a ser lida depois do select inicial. *Esse caminho não é verificável no driver de teste do Fyne, porque `List.Select` não dispara callback antes do primeiro render — a proteção existe pelo mesmo mecanismo já coberto pelo filtro, mas o startup em si não tem teste.*
- **`splitFields` não tinha mudado.** A primeira tentativa de aplicar o escape de barra falhou no meio do script e nada foi gravado, mas a alteração entrou no CHANGELOG e na mensagem de commit como se tivesse sido feita. Agora está feita, com teste que prova o comportamento (`A="b\"c" B=2` são dois tokens).
- **`isWrapper` era mais frágil do que parecia:** comparava o prefixo da string inteira, então `MANGOHUD=1 mangohud %command%` nunca era reconhecida como wrapper — nem para a ordem de montagem, nem para o aviso de duplicata. Agora ela procura o token do nome do programa em qualquer posição, e há teste cobrindo as duas formas.

## 🐛 Correções

Revisão de código achou sete defeitos, seis deles confirmados por execução:

- **Duas receitas do mesmo wrapper** (as duas de `gamescope`) eram combinadas sem aviso, produzindo `gamescope -w 1920 ... gamescope -e -f -F fsr -- %command%`. O primeiro gamescope lê o segundo como o nome do executável do jogo, então o jogo não abre. Agora há aviso de exclusividade, e wrappers diferentes (mangohud + gamescope) continuam sem aviso, que é o aninhamento aceito.
- **Com "copiar ao clicar" ligado, digitar na busca sobrescrevia o clipboard.** `applyFilter` termina com `Select(0)`, que dispara `OnSelected` como se fosse clique — cada tecla digitada destruía o que o usuário tinha copiado. Agora seleção programática é distinguida de clique do usuário.
- **Trocar de idioma jogava o detalhe para o primeiro item.** O guarda que preservava a posição lia `g.selID` depois do `UnselectAll`, que dispara `OnUnselected` e zera o campo: guarda morta. Trocar de idioma agora mantém o item, e o clique manual segue copiando.
- **`displayCmd` e `buildCombination` divergiam** nos launchers que não usam `%command%`: o painel mostrava o `--` que a combinação já tinha removido, então "Copiar comando" e "Copiar combinação" davam respostas diferentes para a mesma opção.
- **Falso aviso de exclusividade:** marcar a receita de Reflex junto com o Anti-Lag 2 avulso produzia "escolha apenas um", mas o resultado era só uma variável repetida. Agora exclusividade é para anti-lag puro contra reflexo puro, e a redundância ganhou aviso próprio.
- **Importar JSON vazio ou `null` dava "0 favoritos importados!"** em vez de erro: o guarda exigia lista não vazia.
- **Favorito dependia do texto do título em português.** `favKey` era `Command + título`; reescrever a redação de um título — o que aconteceu nesta sessão com o `LSFGVK_PROFILE` — apagava o favorito de quem já rodou o app, sem aviso. Agora a chave é só o `Command`, que é único por invariante. Chaves órfãs também são podadas no load, senão ficavam invisíveis na UI e ainda eram exportadas.

## 🛠️ Outros

- O título da janela passou a ser traduzido, e `setLang` o atualiza. A chave `appTitle` existia nos dois idiomas e nunca era usada.
- Removidas as chaves de i18n `selectCommand` e `launcher`, que não eram referenciadas em lugar nenhum.
- Removido um guarda morto em `maxWidthLabel.MinSize` (`if lines < 1`), inalcançável porque `lines` já recebia incremento por parágrafo mais um de margem.
- `splitFields` agora respeita `\` como escape dentro de citação, como o shell faz. Era latente: nenhum comando do catálogo tem escape hoje, mas o `tools/auditar-catalogo.py` depende do mesmo parser, então um `"` digitado por engano engoliria o resto da linha em silêncio.

## 🧪 Testes

- Cobertura em 94.2%. Sete testes novos, cada um verificado por mutação: neutralizar a correção faz o teste falhar com mensagem apontando o caso.
- O clipboard do driver de teste do Fyne é descartável (uma instância nova por chamada), então o teste de cópia observa o `status`, que `copyCurrent` preenche com "Copiado: ".

# v0.6.2

## 📚 Catálogo (104 comandos)

- `LSFGVK_PROFILE` ganhou os fatos da release v2.0.0 (setembro de 2026) depois de conferidos: é preciso trocar para o branch "lsfg-vk" do Lossless Scaling na Steam, o layout do arquivo de configuração mudou e é incompatível com a v1 (quem fizer upgrade precisa desinstalar antes, e o `lsfg-vk-cli healthcheck` acusa sobra), a v2 exige só Vulkan 1.2, e o pipeline bindless novo é 2-3x mais rápido em GPU com FP16 na proporção 2:1 (AMD laptop e handheld) usando 30% da memória. O changelog da v2.0.0 também confirma que ele não menciona a `LSFG_PROCESS` em nenhuma linha, o que reforça a remoção da afirmação de versão que estava lá.
- `LSFGVK_PROFILE` deixou de afirmar que a `LSFG_PROCESS` foi "removida no v2.0.0" — não achei essa informação em fonte nenhuma, e a wiki do projeto ainda cita a variável. A descrição agora traz o que a documentação oficial do lsfg-vk diz: o nome do perfil vai entre aspas quando tem espaço, a forma recomendada nem é a variável e sim o campo "Active In" do perfil no lsfg-vk-ui, VSync precisa estar ligado senão não há frame generation, e cuidado com qual `.exe` você marca em jogos UE4/5.

# v0.6.1

## 🐛 Correções

Auditoria do catálogo inteiro contra as fontes primárias (script do Proton 11, README e CHANGELOG do Proton-CachyOS, README do GE, README e release notes do DXVK, `dxvk.conf`, `radv_instance.c` e o `umu-protonfixes`) encontrou duas entradas que prometiam coisa que não existe mais:

- **`DXVK_FRAME_RATE`** foi **removida no DXVK 3.0** — está no release note, com o próprio doitsujin mandando usar limitador externo ou opção de configuração. A entrada dizia "Todos" e prometia funcionar. Agora avisa que só volta a existir no fork dxvk-low-latency (ou emulada pelo Proton-EM) e aponta a alternativa.
- **`PROTON_FRAME_RATE`** não é limitador do Proton, e o compat "Proton 8+" estava errado: ela é do **Proton-EM**, que só a converte nas opções `dxgi.maxFrameRate` e `d3d9.maxFrameRate` do `DXVK_CONFIG` — o mesmo código que emula `DXVK_FRAME_RATE` e `VKD3D_FRAME_RATE`. Não existe no Proton upstream, nem no GE, nem no CachyOS.
- `VKD3D_FRAME_RATE` ficou mais preciso: existe só no fork vkd3d-low-latency, e a emulação do Proton-EM não alcança D3D12 mesmo.
- `RADV_DEBUG=nofastclears` passou a citar a origem (`radv_debug_options`, primeira entrada) e as irmãs que resolvem artefato parecido, com o aviso de que flag de debug custa desempenho.

## 📚 Catálogo (104 comandos)

- **Novo**: `DXVK_CONFIG="dxgi.maxFrameRate=60;d3d9.maxFrameRate=60"` — a forma suportada de limitar FPS no DXVK atual, já que a variável saiu do upstream.

- `LSFGVK_PROFILE` deixou de afirmar que a `LSFG_PROCESS` foi "removida no v2.0.0" — não achei essa informação em lugar nenhum, e a wiki do projeto ainda cita a variável. A descrição agora traz o que a documentação oficial do lsfg-vk diz: o nome do perfil vai entre aspas quando tem espaço, a forma recomendada nem é a variável e sim o campo "Active In" do perfil no lsfg-vk-ui, VSync precisa estar ligado senão não há frame generation, e cuidado com qual `.exe` você marca em jogos UE4/5.

## 🛠️ Ferramenta

- **`tools/auditar-catalogo.py`** — a auditoria que encontrou os erros acima, virada ferramenta. `tools/auditar-catalogo.py --offline` confere os invariantes do catálogo sem rede e **roda no CI** (comando duplicado, `%command%` faltando, PT/EN incompleto, env var repetida com valores diferentes, wrapper fora da primeira posição). `tools/auditar-catalogo.py --online` baixa 10 fontes do upstream e lista as env vars do catálogo que nenhuma menciona, apontando onde cada uma deveria ser conferida; não falha o build, porque variável sem menção é normal quando ela pertence a outro projeto (DXVK, MangoHud, Mesa). Tem também um mapa `REMOVIDAS` para o caso de o upstream tirar ou renomear uma variável.

## ✅ Verificados sem mudança

- `RADV_DEBUG=nofastclears` e `DXVK_FILTER_DEVICE_NAME` confirmados no código e no README do upstream.
- `LSFGVK_PROFILE` confirmado no `docs/Configuration.md` do lsfg-vk.
- `PROTON_MEDIA_FORCE_GST` está correto: foi renomeado de `PROTON_MEDIA_USE_GST`, e o nome novo é o atual.
- `PROTON_CPU_TOPOLOGY` confirmado no script do Proton, junto com o `default_cpu_limit` de 19 jogos.

# v0.6.0

## 📚 Catálogo (103 comandos)

- **Novo**: `PROTON_CPU_TOPOLOGY` — limita o jogo a um subconjunto de núcleos. O script do Proton repassa a variável para `WINE_CPU_TOPOLOGY`, e o formato `nCPUs:lista` aceita a lista até três vezes para separar grupos de kernel e usuário.
- **Novos 11**: a família de controle do Proton-GE (`PROTON_STEAMINPUT_FALLBACK` e as três `PROTON_STEAMINPUT_LAYOUT_*`, mais `PROTON_SONY_AUTO_XINPUT`, `PROTON_SONY_DUALSENSE_AS_DUALSHOCK4`, `PROTON_SONY_DUALSENSE_EDGE_AS_DUALSENSE`, `PROTON_SONY_DUALSHOCK4_V2_AS_V1`, `PROTON_SONY_HIDRAW_XINPUT`), e duas de plataforma (`PROTON_USE_X11_EXCLUSIVE`, `PROTON_NO_WM_DECORATION`).

## 🛠️ Por baixo do pano

- O CI agora também roda em push de tag (`v*`), então uma release nunca sai sem passar por `gofmt`/`vet`/`test`/`build`.

## 📝 Decisões

- `PROTON_USE_D9VK` e `PROTON_USE_SECCOMP` **não** entraram no catálogo: estão obsoletas no upstream desde o Proton 5.0 e 5.13 respectivamente. O próprio README do Proton-CachyOS marca as duas como obsoletas.
- `PROTON_ENABLE_MEDIACONV` não entrou: a documentação upstream diz literalmente "for debugging purposes, do not use".
- As quatro variantes individuais `PROTON_NVIDIA_NVCUDA`, `_NVENC`, `_NVML` e `_NVOPTIX` foram deixadas dentro da descrição de `PROTON_NVIDIA_LIBS` em vez de virarem entradas — "use só a biblioteca X" é granularidade excessiva para quem consulta o catálogo, e a de NVML já vem ligada por padrão.

## 📚 Catálogo (91 comandos)

- **Novos 4**: `PROTON_FFX4_UPGRADE` (nome atual do upgrade de FSR 4), `PROTON_MLFG_UPGRADE` (frame generation MLFG/Redstone), `PROTON_NVIDIA_LIBS` e `PROTON_NVIDIA_LIBS_NO_32BIT` (bibliotecas NVIDIA alternativas do nvidia-libs).

## 🛠️ Por baixo do pano

- Novo aviso de conflito: `PROTON_NVIDIA_LIBS` + `PROTON_USE_WOW64=1`, porque o Proton desliga as libs automaticamente quando o wow64 está ativo.

## 📝 Correção de auditoria

- A auditoria anterior acusou uma lacuna em `PROTON_VKREFLEX` que **não existe**: esse é o *compat config string* no README do Proton-CachyOS, não uma variável de ambiente. A variável é `DXVK_NVAPI_VKREFLEX`, que o catálogo já tinha. Nada foi criado a partir desse engano.

## 📚 Catálogo (87 comandos)

- **Novos 3**: `LOW_LATENCY_LAYER_SPOOF_NVIDIA`, `LOW_LATENCY_LAYER_FORCE_DECOUPLED` e `VKD3D_FRAME_RATE` — variáveis que até aqui só apareciam dentro da descrição de outra entrada e não davam para copiar.

## 🐛 Correções

- `PROTON_FORCE_NVAPI`: a descrição dizia só "quebra o upgrade FSR 4". O script do Proton define **três** variáveis de uma vez e sem condição — `DXVK_NVAPI_ALLOW_OTHER_DRIVERS=1`, `DXVK_NVAPI_DRIVER_VERSION=99999` e `WINE_HIDE_AMD_GPU=1` — então o FSR 4 quebra sempre, e a alternativa que expõe o mesmo Reflex sem o `WINE_HIDE_AMD_GPU` (spoofing direto por `DXVK_CONFIG`) passa a ser a recomendada.
- `PROTON_FSR4_UPGRADE`: "Desativa Anti-Lag 2" era impreciso. O Proton-CachyOS define `DISABLE_LAYER_MESA_ANTI_LAG=1` junto — só a camada do Mesa, o `low_latency_layer` continua funcionando. O flag já foi removido no 10.0-20250919 e voltou no 10.0-20251007.
- `DXVK_FRAME_RATE`: não vale para D3D12 (use `VKD3D_FRAME_RATE`) e conflita com o limitador do `dxvk-low-latency`.
- `ENABLE_LAYER_MESA_ANTI_LAG`: avisa que o `PROTON_FSR4_UPGRADE` desliga essa camada por conta própria.

## 🛠️ Por baixo do pano

- Novo aviso de conflito: `ENABLE_LAYER_MESA_ANTI_LAG` + `PROTON_FSR4_UPGRADE`, porque o segundo define `DISABLE_LAYER_MESA_ANTI_LAG` por dentro.
- Cobertura de testes em 94%, com as novas entradas e as novas âncoras de documentação.

# v0.5.6

## 📚 Catálogo (82 comandos)

- **Novos 4**: `DXVK_FILTER_DEVICE_NAME` (forçar GPU NVIDIA/AMD pelo nome), `PROTON_FRAME_RATE` (limite de FPS no nível do Proton) e `MANGOHUD_CONFIG` (limite de FPS via MangoHud).

## ✨ Novidades

- Exportar/importar favoritos em JSON (o import ignora chaves inválidas e duplicadas).
- Busca avançada com prefixos: `cat:`, `compat:` e `cmd:` (ex.: `cat:GPU`).
- Botões de exportar/importar com texto e ícone; campo de busca com ícone.
- Barras de filtro e topo com rolagem horizontal (não cortam em janela estreita).

## 🛠️ Por baixo do pano

- Favoritos: persistência extraída para `saveFavs()`; import valida as chaves contra o catálogo.
- Cobertura de testes em 93%+, com testes para import/export, busca por prefixo e montagem da interface.

# v0.5.5

Favoritos mais robustos e cobertura de testes quase triplicada.

## 🐛 Correções

- Favoritos usam chave composta (comando + título) para evitar colisão quando dois comandos diferentes têm o mesmo shell.

## 🛠️ Por baixo do pano

- Cobertura de testes subiu de 19.8% para 38.7%.
- Novos testes: toggleFav, clearSelection, selectCommand, clearDetail, updateFavButton, setTheme, setLang, launcher, updateCombination, applyFilter, combinação, conflitos, splitFields, displayCmd, tr, t, cmd.

# v0.5.4

Detalhe acompanha a lista ao pesquisar e ao trocar de idioma.

## 🐛 Correções / Fixes

- Novo `maxWidthLabel`: combinação, status e detalhes quebram linha dentro da janela em vez de esticá-la além da tela.
- Removida a rolagem (`VScroll`) da barra de combinação.
- Pesquisa: painel de detalhes agora acompanha a lista mesmo quando a linha 0 já estava selecionada.
- Idioma: detalhe atualiza o texto ao trocar PT/EN (mesmo bug do search).

# v0.5.3

Aparece só na aba Jogos do menu do sistema.

## 🖥️ Menu

- `Categories=Game;` no `.desktop` (removido `Utility`) — o app não aparece mais em Utilitários.

# v0.5.2

lsfg-vk atualizado para v2.0.0+ e barra de combinação com rolagem.

## 📚 Catálogo (78 comandos)

- **Removido**: `LSFG_PROCESS=steam` (variável removida no lsfg-vk v2.0.0).

## 🐛 Correções

- Barra de combinação com rolagem (`NewVScroll`) — não expande mais a janela além da tela com muitos comandos.
- README com exemplo de combinação e versão atualizados.

# v0.5.1

Ordem correta dos wrappers na combinação e lsfg-vk v2.0.0+.

## 📚 Catálogo

- **Novo**: `LSFGVK_PROFILE=steam` (variável oficial do lsfg-vk v2.0.0+ para selecionar o perfil).

## 🐛 Correções

- Wrappers (`game-performance`, `mangohud`, `gamemoderun`, `gamescope`) agora vão por último, antes do `%command%`: `PROTON_LOG=1 mangohud gamemoderun %command%`.
- Label da combinação contido (`maxWidthLabel`) para não esticar a janela.

# v0.5.0

Catálogo auditado, AppImage como formato único (com atualização delta) e releases automáticos pelo GitHub.

## 📚 Catálogo (78 comandos)

- **Removidos 18**: variáveis obsoletas, renomeadas ou sem efeito nas versões atuais — `DXVK_ASYNC`, `PROTON_USE_NTSYNC`, `PROTON_ENABLE_NVAPI`, `PROTON_VKREFLEX`, `PROTON_VKD3D_HEAP`, `PROTON_NO_D3D9`, `PROTON_ENABLE_HDR`, `PROTON_DUMP_DEBUG_COMMANDS`, `WINE_AUDIO_DRIVER`, `WINE_BLOCK_HOSTS`, `WINE_VIRTUAL_DESKTOP`, `WINE_ESYNC`, `WINEFSYNC`, `FNA3D_FORCE_DRIVER`, `DRI_CONFIG` — e as 3 entradas de cracks (Online-Fix/SteamFix).
- **Novos 4**: `PROTON_FFX3_UPGRADE` (nome atual do FSR 3.1 no CachyOS 11+), `PROTON_VKD3D_LOWLATENCY` (D3D12), `DXVK_FRAME_RATE=60` (limite de FPS sem overlay), `PROTON_D7VK_DDRAW` (jogos DX7 ou anteriores).
- **Correções**: `DXVK_HUD` (`gpu` → `gpuload`); escopo explícito DX8–11 vs DX12 no low-latency; compatibilidades revisadas (RDNA3, indicadores FSR4/DLSS, LSFG).
- Removida a duplicata `PROTON_XESS_UPGRADE`.

## 📦 Distribuição

- Só AppImage: removidos `.deb`, `.rpm` e scripts por formato.
- Nome com versão (`protoncommand-0.5.0-x86_64.AppImage`) e update info embutida (zsync) — atualize por delta com AppImageUpdate ou Gear Lever.
- Release publicado automaticamente pelo GitHub a cada tag `v*`.

## 🛠️ Por baixo do pano

- Combinação: wrappers vão por último, antes de `%command%` (ordem correta: env vars → wrappers → %command%).
- Detector de conflitos respeita aspas e compara nomes exatos.
- CI em todo push/PR (`gofmt`, `vet`, `test`, `build`) com retry de download e fallback de proxy; testes anti-regressão do catálogo e i18n.

## 🛠️ Instalação rápida

```bash
chmod +x protoncommand-0.5.0-x86_64.AppImage && ./protoncommand-0.5.0-x86_64.AppImage
```

# v0.4.0

Correções de nomenclatura e internacionalização do rótulo de idioma.

## 🐛 Correções

- **Licença**: copyright atualizado de `ProtonBox contributors` para `protoncommand contributors` no `LICENSE`.
- **README**: seção em Português agora diz "descrição em Português e Inglês" (antes "English").
- **App**: o seletor de idioma e o menu de configurações passam a exibir **Inglês** em vez de "English" (PT e EN).

## 📦 Pacotes

| Formato | Arquivo | Uso |
|---|---|---|
| AppImage | `protoncommand-0.4.0.x86_64.AppImage` | qualquer distro |
| Debian/Ubuntu | `protoncommand-0.4.0.x86_64.deb` | `sudo dpkg -i` |
| Fedora/openSUSE | `protoncommand-0.4.0.x86_64.rpm` | `sudo rpm -i` |

## 🛠️ Instalação rápida

```bash
# AppImage
chmod +x protoncommand-0.4.0.x86_64.AppImage && ./protoncommand-0.4.0.x86_64.AppImage

# Debian/Ubuntu
sudo dpkg -i protoncommand-0.4.0.x86_64.deb

# Fedora/openSUSE
sudo rpm -i protoncommand-0.4.0.x86_64.rpm
```

# v0.3.0

Catálogo expandido e descrições reescritas com foco prático (sintoma → solução, exemplos de valores e consistência PT/EN).

## ✨ Novos comandos

### Upscalers e upgrades
- `PROTON_FSR4_INDICATOR=1`, `PROTON_FSR4_RDNA3_UPGRADE=1`, `PROTON_FSR3_UPGRADE=1` — indicador e upgrades FSR4/FSR3
- `PROTON_XESS_UPGRADE=1` — XeSS (atualizado)
- `PROTON_DLSS_INDICATOR=1` — indicador DLSS
- `PROTON_OPTISCALER_NAME=dxgi.dll`, `PROTON_OPTISCALER_CONFIG=...` — OptiScaler

### Latência e HDR
- `DXVK_NVAPI_VKREFLEX=1`, `PROTON_VKREFLEX=1` — NVIDIA Reflex em jogos Vulkan
- `PROTON_ENABLE_HDR=1` — HDR via Proton
- `PROTON_VKD3D_HEAP=1` — heap dedicada do vkd3d

### Display e input
- `PROTON_USE_WAYLAND=1`, `PROTON_PREFER_SDL=1`, `PROTON_NO_STEAMINPUT=1`

### Diversos
- `PROTON_ADD_CONFIG=config1,config2` — aplica configs do Proton sem editar arquivos
- `PROTON_LOCAL_SHADER_CACHE=1` — cache local de shaders
- `PROTON_MEDIA_FORCE_GST=1`, `PROTON_GST_VIDEO_ORIENTATION=90`, `WINE_BLOCK_HOSTS=...`

### Obsoleto
- `PROTON_USE_NTSYNC=1` — marcado como obsoleto (ntsync é o padrão no Proton 11+; use `PROTON_NO_NTSYNC=1` para desativar)

## 📦 Pacotes

| Formato | Arquivo | Uso |
|---|---|---|
| AppImage | `protoncommand-0.3.0.x86_64.AppImage` | qualquer distro |
| Debian/Ubuntu | `protoncommand-0.3.0.x86_64.deb` | `sudo dpkg -i` |
| Fedora/openSUSE | `protoncommand-0.3.0.x86_64.rpm` | `sudo rpm -i` |

## 🛠️ Instalação rápida

```bash
# AppImage
chmod +x protoncommand-0.3.0.x86_64.AppImage && ./protoncommand-0.3.0.x86_64.AppImage

# Debian/Ubuntu
sudo dpkg -i protoncommand-0.3.0.x86_64.deb

# Fedora/openSUSE
sudo rpm -i protoncommand-0.3.0.x86_64.rpm
```
