package main

type Localized struct {
	PT string
	EN string
}

type Command struct {
	Command     string
	CommandEN   string
	Title       Localized
	Category    Localized
	Compat      Localized
	Description Localized
}

func commands() []Command {
	return []Command{
		{
			Command: "PROTON_LOG=1 %command%",
			Title: Localized{
				PT: "Log detalhado do Proton",
				EN: "Detailed Proton log",
			},
			Category: Localized{
				PT: "Diagnóstico",
				EN: "Diagnostics",
			},
			Compat: Localized{
				PT: "Proton padrão, GE e CachyOS",
				EN: "Standard Proton, GE and CachyOS",
			},
			Description: Localized{
				PT: "Gera um arquivo de log detalhado (steam-<appid>.log) na sua pasta pessoal. Use quando o jogo trava, não abre ou tem erros de desempenho misteriosos — o log mostra o que falhou. O arquivo só aparece depois de fechar o jogo.",
				EN: "Generates a detailed log file (steam-<appid>.log) in your home folder. Use when a game crashes, won't launch, or has mysterious performance issues — the log shows what failed. The file only appears after closing the game.",
			},
		},
		{
			Command: "PROTON_LOG=warn+pipewire,warn+mmdevapi %command%",
			Title: Localized{
				PT: "Log de problemas de áudio",
				EN: "Audio issue log",
			},
			Category: Localized{
				PT: "Diagnóstico",
				EN: "Diagnostics",
			},
			Compat: Localized{
				PT: "CachyOS",
				EN: "CachyOS",
			},
			Description: Localized{
				PT: "Variação do PROTON_LOG focada em áudio, para diagnosticar problemas de som (clipping, estalo ou áudio ausente) no driver PipeWire. Use quando o áudio do jogo está estranho. Não suba para +pipewire completo: o tracing total atrapalha o timing do áudio e causa mais estalos.",
				EN: "PROTON_LOG variant focused on audio, to diagnose sound issues (clipping, crackling or missing audio) on the PipeWire driver. Use when game audio is weird. Don't go above warn+: full tracing perturbs audio timing and causes more crackling.",
			},
		},
		{
			Command: "PROTON_LOG_DIR=~/proton-logs %command%",
			Title: Localized{
				PT: "Pasta dos logs",
				EN: "Log folder",
			},
			Category: Localized{
				PT: "Diagnóstico",
				EN: "Diagnostics",
			},
			Compat: Localized{
				PT: "Proton padrão, GE e CachyOS",
				EN: "Standard Proton, GE and CachyOS",
			},
			Description: Localized{
				PT: "Redireciona os logs do Proton para a pasta que você indicar (ex.: ~/proton-logs) em vez da pasta pessoal. Use para manter um log por jogo e evitar poluir o home quando você testa vários títulos.",
				EN: "Redirects Proton logs to the folder you specify (e.g.: ~/proton-logs) instead of your home directory. Use to keep one log per game and avoid cluttering your home when testing several titles.",
			},
		},
		{
			Command: "PROTON_CRASH_REPORT_DIR=~/crash-reports %command%",
			Title: Localized{
				PT: "Relatórios de crash",
				EN: "Crash reports",
			},
			Category: Localized{
				PT: "Diagnóstico",
				EN: "Diagnostics",
			},
			Compat: Localized{
				PT: "Proton padrão, GE e CachyOS",
				EN: "Standard Proton, GE and CachyOS",
			},
			Description: Localized{
				PT: "Grava logs de crash na pasta que você indicar (ex.: ~/crash-reports). Útil para inspecionar o motivo de travamentos repetidos. Atenção: não limpa logs antigos, então pode encher o disco se não for monitorado de tempos em tempos.",
				EN: "Writes crash logs to the folder you specify (e.g.: ~/crash-reports). Useful to inspect why a game keeps crashing. Warning: it does not clean old logs, so it can fill your disk if not cleared occasionally.",
			},
		},
		{
			Command: "DXVK_HUD=fps %command%",
			Title: Localized{
				PT: "HUD de FPS (DXVK)",
				EN: "FPS HUD (DXVK)",
			},
			Category: Localized{
				PT: "Overlay e desempenho",
				EN: "Overlay & Performance",
			},
			Compat: Localized{
				PT: "Todos",
				EN: "All",
			},
			Description: Localized{
				PT: "Mostra apenas os quadros por segundo (FPS) do jogo na tela. É o overlay mais leve do DXVK — ideal para um teste rápido de desempenho sem poluir a tela com informações demais.",
				EN: "Shows only the frames per second (FPS) on screen. The lightest DXVK overlay — ideal for a quick performance check without cluttering the screen with too much info.",
			},
		},
		{
			Command: "DXVK_HUD=fps,gpuload,api,devinfo %command%",
			Title: Localized{
				PT: "HUD completo (DXVK)",
				EN: "Full HUD (DXVK)",
			},
			Category: Localized{
				PT: "Overlay e desempenho",
				EN: "Overlay & Performance",
			},
			Compat: Localized{
				PT: "Todos",
				EN: "All",
			},
			Description: Localized{
				PT: "Painel detalhado com FPS, carga da GPU, API gráfica em uso (Vulkan/D3D) e informações do driver. Combina as métricas: fps, gpuload, api, devinfo, version, drawcalls, memory. Separe os itens com vírgula.",
				EN: "Detailed panel with FPS, GPU load, graphics API in use (Vulkan/D3D) and driver info. Combine metrics: fps, gpuload, api, devinfo, version, drawcalls, memory. Separate items with commas.",
			},
		},
		{
			Command: "DXVK_HUD=version %command%",
			Title: Localized{
				PT: "Versão do DXVK",
				EN: "DXVK version",
			},
			Category: Localized{
				PT: "Overlay e desempenho",
				EN: "Overlay & Performance",
			},
			Compat: Localized{
				PT: "Todos",
				EN: "All",
			},
			Description: Localized{
				PT: "Exibe a versão do DXVK em uso no canto da tela. Útil para confirmar qual build de DXVK o Proton está carregando antes de reportar problemas.",
				EN: "Shows the DXVK version in use in the corner of the screen. Useful to confirm which DXVK build Proton is loading before reporting issues.",
			},
		},
		{
			Command: "mangohud %command%",
			Title: Localized{
				PT: "MangoHud",
				EN: "MangoHud",
			},
			Category: Localized{
				PT: "Overlay e desempenho",
				EN: "Overlay & Performance",
			},
			Compat: Localized{
				PT: "Todos (exige mangohud instalado)",
				EN: "All (requires mangohud)",
			},
			Description: Localized{
				PT: "Painel avançado na tela com métricas completas: uso de CPU, GPU, memória RAM, temperaturas, frequências e FPS. Instale com: sudo pacman -S mangohud. Configurações em ~/.config/MangoHud/MangoHud.conf.",
				EN: "Advanced on-screen panel with full metrics: CPU/GPU usage, RAM, temperatures, frequencies and FPS. Install with: sudo pacman -S mangohud. Settings in ~/.config/MangoHud/MangoHud.conf.",
			},
		},
		{
			Command: "gamescope -w 1920 -h 1080 -r 60 %command%",
			Title: Localized{
				PT: "Gamescope",
				EN: "Gamescope",
			},
			Category: Localized{
				PT: "Overlay e desempenho",
				EN: "Overlay & Performance",
			},
			Compat: Localized{
				PT: "Todos (exige gamescope instalado)",
				EN: "All (requires gamescope)",
			},
			Description: Localized{
				PT: "Compositor micro do Steam/Valve que isola o jogo e permite upscaling, limite de FPS, FSR e janela redimensionável. Ajuste -w/-h para a resolução interna do jogo e -r para o limite de FPS.",
				EN: "Valve's micro compositor that isolates the game and allows upscaling, FPS limit, FSR and a resizable window. Adjust -w/-h for the game's internal resolution and -r for the FPS limit.",
			},
		},
		{
			Command: "gamemoderun %command%",
			Title: Localized{
				PT: "GameMode",
				EN: "GameMode",
			},
			Category: Localized{
				PT: "Overlay e desempenho",
				EN: "Overlay & Performance",
			},
			Compat: Localized{
				PT: "Todos (exige gamemode instalado)",
				EN: "All (requires gamemode)",
			},
			Description: Localized{
				PT: "Ativa o GameMode (da Feral Interactive), que aplica otimizações temporárias de CPU/GPU enquanto o jogo roda. Instale com: sudo pacman -S gamemode lib32-gamemode.",
				EN: "Enables GameMode (by Feral Interactive), applying temporary CPU/GPU optimizations while the game runs. Install with: sudo pacman -S gamemode lib32-gamemode.",
			},
		},
		{
			Command: "WINE_FULLSCREEN_FSR=1 %command%",
			Title: Localized{
				PT: "FidelityFX Super Resolution (FSR)",
				EN: "FidelityFX Super Resolution (FSR)",
			},
			Category: Localized{
				PT: "Upscaling",
				EN: "Upscaling",
			},
			Compat: Localized{
				PT: "Todos (Vulkan via DXVK)",
				EN: "All (Vulkan via DXVK)",
			},
			Description: Localized{
				PT: "Ativa o upscaling FSR 1 da AMD no modo fullscreen: o jogo renderiza em resolução menor e o FSR sobe a imagem, ganhando FPS em placas mais fracas. Combine com WINE_FULLSCREEN_FSR_STRENGTH=2 para ajustar a nitidez (0 = máxima, 5 = mínima). Funciona nos jogos renderizados via Vulkan — inclui títulos D3D11/D3D12 traduzidos pelo DXVK/VKD3D; não funciona em OpenGL (wined3d). Recurso legado: prefira o Gamescope (-F fsr) ou o upgrade FSR 4.",
				EN: "Enables AMD FSR 1 upscaling in fullscreen: the game renders at a lower resolution and FSR upscales the image, gaining FPS on weaker GPUs. Combine with WINE_FULLSCREEN_FSR_STRENGTH=2 for sharpness (0 = max, 5 = min). Works in games rendered through Vulkan — including D3D11/D3D12 titles translated by DXVK/VKD3D; does not work in OpenGL (wined3d). Legacy feature: prefer Gamescope (-F fsr) or the FSR 4 upgrade.",
			},
		},
		{
			Command: "WINE_FULLSCREEN_INTEGER_SCALING=1 %command%",
			Title: Localized{
				PT: "Escala inteira (pixels nítidos)",
				EN: "Integer scaling (sharp pixels)",
			},
			Category: Localized{
				PT: "Upscaling",
				EN: "Upscaling",
			},
			Compat: Localized{
				PT: "Proton padrão",
				EN: "Standard Proton",
			},
			Description: Localized{
				PT: "Ativa a escala inteira em fullscreen: pixels nítidos e quadradinhos ao subir a resolução, sem blur. Útil em jogos antigos ou pixel art.",
				EN: "Enables integer scaling in fullscreen: sharp, blocky pixels when upscaling, no blur. Useful in old or pixel-art games.",
			},
		},
		{
			Command: "PROTON_FSR4_UPGRADE=1 %command%",
			Title: Localized{
				PT: "Upgrade FSR 3.1 → FSR 4",
				EN: "FSR 3.1 → FSR 4 upgrade",
			},
			Category: Localized{
				PT: "Upscaling",
				EN: "Upscaling",
			},
			Compat: Localized{
				PT: "GE e CachyOS (GPU AMD RDNA3+/RDNA4)",
				EN: "GE and CachyOS (AMD RDNA3+/RDNA4 GPU)",
			},
			Description: Localized{
				PT: "Baixa automaticamente a amdxcffx64.dll e atualiza jogos com FSR 3.1 para FSR 4. Versão customizável: PROTON_FSR4_UPGRADE=\"4.0.2\" (default 4.0.2 no GE, 4.1.1 no CachyOS). Para RDNA3 use PROTON_FSR4_RDNA3_UPGRADE. No Proton-CachyOS esta variável vem acompanhada de DISABLE_LAYER_MESA_ANTI_LAG=1, ou seja, ela desliga a camada Anti-Lag 2 do Mesa — porque as duas brigavam por causa da issue 47. O flag já foi removido uma vez (10.0-20250919, quando acharam que o Anti-Lag 2 tinha melhorado) e voltou no 10.0-20251007, então espere por ele. Isso NÃO afeta o low_latency_layer: o Reflex/Anti-Lag 2 do LOW_LATENCY_LAYER=1 continua funcionando, porque é outra camada.",
				EN: "Automatically downloads amdxcffx64.dll and upgrades games with FSR 3.1 to FSR 4. Custom version: PROTON_FSR4_UPGRADE=\"4.0.2\" (default 4.0.2 on GE, 4.1.1 on CachyOS). For RDNA3 use PROTON_FSR4_RDNA3_UPGRADE. On Proton-CachyOS this variable ships alongside DISABLE_LAYER_MESA_ANTI_LAG=1, i.e. it turns off the Mesa Anti-Lag 2 layer — the two used to clash over issue 47. The flag was removed once (10.0-20250919, when Anti-Lag 2 looked fixed) and came back in 10.0-20251007, so expect it. This does NOT affect low_latency_layer: the Reflex/Anti-Lag 2 from LOW_LATENCY_LAYER=1 keeps working, because that's a different layer.",
			},
		},
		{
			Command: "PROTON_DLSS_UPGRADE=1 %command%",
			Title: Localized{
				PT: "Upgrade DLSS (versão customizável)",
				EN: "DLSS upgrade (custom version)",
			},
			Category: Localized{
				PT: "Upscaling",
				EN: "Upscaling",
			},
			Compat: Localized{
				PT: "GE e CachyOS (GPU NVIDIA)",
				EN: "GE and CachyOS (NVIDIA GPU)",
			},
			Description: Localized{
				PT: "Baixa automaticamente versões mais novas das DLLs nvngx_dlss e substitui as do jogo. Fixe a versão: PROTON_DLSS_UPGRADE=\"310.2\". Também define DXVK_NVAPI_DRS_SETTINGS para o preset mais recente.",
				EN: "Automatically downloads newer nvngx_dlss DLLs and replaces the game's. Pin version: PROTON_DLSS_UPGRADE=\"310.2\". Also sets DXVK_NVAPI_DRS_SETTINGS to latest preset.",
			},
		},
		{
			Command: "PROTON_USE_WINED3D=1 %command%",
			Title: Localized{
				PT: "Forçar OpenGL (wined3d)",
				EN: "Force OpenGL (wined3d)",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "Todos",
				EN: "All",
			},
			Description: Localized{
				PT: "Força o Proton a usar a tradução OpenGL do Wine (wined3d) em vez do Vulkan (DXVK). Geralmente perde desempenho; use só quando o DXVK trava o jogo ou a placa é tão antiga que nem Vulkan 1.0 tem. Útil também para isolar se um bug é do DXVK ou do jogo.",
				EN: "Forces Proton to use Wine's OpenGL translation (wined3d) instead of Vulkan (DXVK). Usually loses performance; use only when DXVK crashes the game or the GPU is so old it lacks even Vulkan 1.0. Also useful to isolate whether a bug is in DXVK or the game.",
			},
		},
		{
			Command: "PROTON_DXVK_SAREK=1 %command%",
			Title: Localized{
				PT: "DXVK-Sarek (placas antigas)",
				EN: "DXVK-Sarek (old GPUs)",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "CachyOS",
				EN: "CachyOS",
			},
			Description: Localized{
				PT: "Usa o fork dxvk-sarek do DXVK, feito para GPUs antigas que só suportam Vulkan 1.1/1.2 (em vez de 1.3). Usa o branch async e NÃO deve ser usado em jogos multiplayer ou com anti-cheat.",
				EN: "Uses the dxvk-sarek DXVK fork, made for old GPUs that only support Vulkan 1.1/1.2 (instead of 1.3). Uses the async branch and MUST NOT be used in multiplayer or anti-cheat games.",
			},
		},
		{
			Command: "PROTON_DXVK_LOWLATENCY=1 %command%",
			Title: Localized{
				PT: "DXVK de baixa latência (DX8/9/10/11)",
				EN: "Low latency DXVK (DX8/9/10/11)",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "CachyOS (ou qualquer Proton com as DLLs do dxvk-low-latency instaladas)",
				EN: "CachyOS (or any Proton with the dxvk-low-latency DLLs installed)",
			},
			Description: Localized{
				PT: "Troca o DXVK pelo fork dxvk-low-latency (a partir da 3.1.1) nos jogos Direct3D 8/9/10/11: substitui o max-frame-latency upstream por um pacing que atrasa a CPU só o necessário, reduzindo bastante o input lag e a variação de latência. Já vem no Proton-CachyOS; em outros Protons dá para instalar manualmente trocando as DLLs em files/lib/wine/dxvk dentro de um Proton em ~/.local/share/Steam/compatibilitytools.d. Para VRR, combine com DXVK_FRAME_PACE=low-latency-vrr. Diagnóstico: DXVK_HUD=latencydetails. Para jogos Direct3D 12 use PROTON_VKD3D_LOWLATENCY.",
				EN: "Swaps DXVK for the dxvk-low-latency fork (3.1.1 and newer) in Direct3D 8/9/10/11 games: it replaces upstream's max-frame-latency with pacing that delays the CPU only as much as needed, cutting input lag and latency variance considerably. Bundled with Proton-CachyOS; in other Protons you can install it manually by replacing the DLLs in files/lib/wine/dxvk of a Proton under ~/.local/share/Steam/compatibilitytools.d. For VRR, combine with DXVK_FRAME_PACE=low-latency-vrr. Diagnostics: DXVK_HUD=latencydetails. For Direct3D 12 games use PROTON_VKD3D_LOWLATENCY.",
			},
		},
		{
			Command: "PROTON_VKD3D_LOWLATENCY=1 %command%",
			Title: Localized{
				PT: "VKD3D de baixa latência (D3D12)",
				EN: "Low latency VKD3D (D3D12)",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "CachyOS (ou Proton com vkd3d-proton substituído pelo fork)",
				EN: "CachyOS (or a Proton whose vkd3d-proton was replaced by the fork)",
			},
			Description: Localized{
				PT: "Usa o fork vkd3d-low-latency nos jogos Direct3D 12: o equivalente do dxvk-low-latency para D3D12. O pacing de baixa latência só é ativado se o jogo usar OU a API Reflex OU swapchains DXGI aguardáveis (waitable) — apenas cerca de 20-30% dos títulos DX12 usam swapchains aguardáveis, então confirme com PROTON_LOG=1 procurando \"waitable dxgi swapchain present percentage\" no log do proton (perto de 100% = ativo, 0% ou ausente = o jogo não usa). Se o jogo tiver ambos, o Reflex tem prioridade. Já foram verificados Resident Evil 2, Resident Evil 7, Street Fighter 6, Warframe, Devil May Cry 5, Monster Hunter Rise, Forza Horizon 4, The Division 2, Overwatch, Dead Space Remake e Witchfire; a lista completa está na discussion do projeto. Dois avisos: em jogos Unreal Engine 4 com r.OneFrameThreadLag=1 (o padrão) as swapchains aguardáveis causam queda forte de desempenho — solte um limite de fps (VKD3D_FRAME_RATE) para evitar, e lembre que esse limitador entra em conflito com o do próprio fork; e Anti-Lag 2 não é suportado aqui. Em GPU AMD o caminho do Reflex funciona por spoofing, mas isso quebra o upgrade FSR4 — ver a entrada de Reflex do low_latency_layer. Para jogos Direct3D 8/9/10/11 use PROTON_DXVK_LOWLATENCY.",
				EN: "Uses the vkd3d-low-latency fork in Direct3D 12 games: the D3D12 equivalent of dxvk-low-latency. The low-latency pacing only kicks in if the game uses EITHER the Reflex API OR waitable DXGI swapchains — only around 20-30% of DX12 titles use waitable swapchains, so verify with PROTON_LOG=1 and look for \"waitable dxgi swapchain present percentage\" in the proton log (near 100% = active, 0% or missing = the game doesn't use it). If a game has both, Reflex wins. Verified so far in Resident Evil 2, Resident Evil 7, Street Fighter 6, Warframe, Devil May Cry 5, Monster Hunter Rise, Forza Horizon 4, The Division 2, Overwatch, Dead Space Remake and Witchfire; the full list lives in the project's discussion thread. Two warnings: in Unreal Engine 4 games with r.OneFrameThreadLag=1 (the default) waitable swapchains cause a heavy performance drop — set an fps limit (VKD3D_FRAME_RATE) to avoid it, and note that limiter conflicts with the fork's own one; and Anti-Lag 2 is not supported here. On AMD GPUs the Reflex path works through spoofing, but that breaks the FSR4 upgrade — see the low_latency_layer Reflex entry. For Direct3D 8/9/10/11 games use PROTON_DXVK_LOWLATENCY.",
			},
		},
		{
			Command: "DXVK_FRAME_PACE=low-latency-vrr %command%",
			Title: Localized{
				PT: "Pacing de baixa latência com VRR (DX8/9/10/11)",
				EN: "Low-latency pacing with VRR (DX8/9/10/11)",
			},
			Category: Localized{
				PT: "Latência",
				EN: "Latency",
			},
			Compat: Localized{
				PT: "CachyOS (requer PROTON_DXVK_LOWLATENCY=1 e monitor VRR)",
				EN: "CachyOS (requires PROTON_DXVK_LOWLATENCY=1 and a VRR monitor)",
			},
			Description: Localized{
				PT: "Ativa o modo de VRR do dxvk-low-latency 3.1.1+, que usa VK_EXT_present_timing para usar timings de VBlank precisos e detectar o refresh do monitor automaticamente: elimina o buffer de V-Sync e mantém o pacing suave. Exige PROTON_DXVK_LOWLATENCY=1 e o monitor configurado em frequência variável. O limite de fps é fixado em 5% abaixo do refresh máximo e pode ser ajustado com DXVK_FRAME_RATE (ex.: 225 em 240 Hz). No Wayland precisa de wp_presentation v2, que a maioria dos compositors já oferece. Em versões anteriores à 3.1.1 o modo equivalente era low-latency-vrr-<refresh>, ex.: low-latency-vrr-240. Diagnóstico: DXVK_HUD=latencydetails.",
				EN: "Enables the VRR mode of dxvk-low-latency 3.1.1+, which uses VK_EXT_present_timing for precise VBlank timings and automatic refresh rate detection: it removes V-Sync buffering while keeping pacing smooth. Requires PROTON_DXVK_LOWLATENCY=1 and a monitor configured for variable refresh rate. The fps cap is fixed at 5% below the max refresh rate and can be overridden with DXVK_FRAME_RATE (e.g. 225 on 240 Hz). On Wayland it needs wp_presentation v2, which most compositors already provide. Before 3.1.1 the equivalent mode was low-latency-vrr-<refresh>, e.g. low-latency-vrr-240. Diagnostics: DXVK_HUD=latencydetails.",
			},
		},
		{
			Command: "DXVK_HDR=1 %command%",
			Title: Localized{
				PT: "HDR via DXVK",
				EN: "HDR via DXVK",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "CachyOS (exige monitor e compositor com HDR)",
				EN: "CachyOS (requires HDR monitor and compositor)",
			},
			Description: Localized{
				PT: "Ativa HDR nos jogos via DXVK. Use quando o jogo tem opção de HDR mas não acende sozinho. Em NVIDIA com drivers mais antigos, combine com ENABLE_HDR_WSI=1.",
				EN: "Enables HDR in games via DXVK. Use when the game has an HDR option but won't turn on by itself. On NVIDIA with older drivers, combine with ENABLE_HDR_WSI=1.",
			},
		},
		{
			Command: "PROTON_NO_D3D11=1 %command%",
			Title: Localized{
				PT: "Desabilitar D3D11",
				EN: "Disable D3D11",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "Todos",
				EN: "All",
			},
			Description: Localized{
				PT: "Desativa a d3d11.dll. Só serve para jogos que rodam melhor caindo para o fallback D3D9. Se o jogo não tiver fallback, ele simplesmente não abre.",
				EN: "Disables d3d11.dll. Only useful for games that run better falling back to D3D9. If the game has no fallback, it simply won't start.",
			},
		},
		{
			Command: "PROTON_NO_ESYNC=1 %command%",
			Title: Localized{
				PT: "Desabilitar ESync",
				EN: "Disable ESync",
			},
			Category: Localized{
				PT: "Sincronização",
				EN: "Synchronization",
			},
			Compat: Localized{
				PT: "Todos",
				EN: "All",
			},
			Description: Localized{
				PT: "Desativa a sincronização por eventfd (ESync). Use como teste quando o jogo trava, congela ou tem travamentos esporádicos — se desativar resolver, o problema é na sincronização. Não muda nada se o jogo já roda bem.",
				EN: "Disables eventfd-based synchronization (ESync). Use as a test when a game crashes, freezes, or has sporadic hitches — if disabling fixes it, sync was the culprit. Does nothing if the game already runs fine.",
			},
		},
		{
			Command: "PROTON_NO_FSYNC=1 %command%",
			Title: Localized{
				PT: "Desabilitar FSync",
				EN: "Disable FSync",
			},
			Category: Localized{
				PT: "Sincronização",
				EN: "Synchronization",
			},
			Compat: Localized{
				PT: "Todos",
				EN: "All",
			},
			Description: Localized{
				PT: "Desativa a sincronização por futex (FSync). Use como teste se o jogo trava ou tem travamentos estranhos — alterna com PROTON_NO_ESYNC=1 para descobrir qual sincronização é a culpada. (Em kernels sem FUTEX_WAIT_MULTIPLE, já vem desativado.)",
				EN: "Disables futex-based synchronization (FSync). Use as a test if the game crashes or hitches weirdly — toggle with PROTON_NO_ESYNC=1 to find which sync is the culprit. (On kernels without FUTEX_WAIT_MULTIPLE, it's already off.)",
			},
		},
		{
			Command: "DRI_PRIME=1 %command%",
			Title: Localized{
				PT: "Forçar GPU dedicada",
				EN: "Force discrete GPU",
			},
			Category: Localized{
				PT: "GPU",
				EN: "GPU",
			},
			Compat: Localized{
				PT: "Todos (laptops híbridos)",
				EN: "All (hybrid laptops)",
			},
			Description: Localized{
				PT: "Em laptops com GPU integrada + dedicada, força o jogo a rodar na placa de vídeo dedicada (mais forte). Use quando o jogo está rodando na integrada e ficando lento. O índice exato da GPU aparece em `lspci | grep VGA`; ex.: DRI_PRIME=pci-0000_01_00_0.",
				EN: "On laptops with integrated + discrete GPUs, forces the game to run on the dedicated (stronger) GPU. Use when the game is running on the integrated one and feels slow. The exact GPU index shows in `lspci | grep VGA`; e.g.: DRI_PRIME=pci-0000_01_00_0.",
			},
		},
		{
			Command: "PROTON_ENABLE_WAYLAND=1 %command%",
			Title: Localized{
				PT: "Driver nativo Wayland",
				EN: "Native Wayland driver",
			},
			Category: Localized{
				PT: "Wayland",
				EN: "Wayland",
			},
			Compat: Localized{
				PT: "GE e CachyOS (experimental)",
				EN: "GE and CachyOS (experimental)",
			},
			Description: Localized{
				PT: "Ativa o driver winewayland (janela nativa em Wayland, sem XWayland). É experimental: launchers Electron (Battle.net, EA App, Ubisoft Connect) podem abrir janela branca; nesse caso tente adicionar --in-process-gpu aos argumentos. Se houver problema com controle, use PROTON_USE_SDL=1.",
				EN: "Enables the winewayland driver (native Wayland window, no XWayland). It's experimental: Electron launchers (Battle.net, EA App, Ubisoft Connect) may show a white window; try adding --in-process-gpu to the arguments then. If you have controller issues, use PROTON_USE_SDL=1.",
			},
		},
		{
			Command: "PROTON_USE_PIPEWIRE=0 %command%",
			Title: Localized{
				PT: "Trocar driver de áudio",
				EN: "Switch audio driver",
			},
			Category: Localized{
				PT: "Áudio",
				EN: "Audio",
			},
			Compat: Localized{
				PT: "CachyOS",
				EN: "CachyOS",
			},
			Description: Localized{
				PT: "Desativa o driver de áudio winepipewire (ligado por padrão no proton-cachyos) e volta para o winepulse. Útil quando o jogo tem áudio estalando ou cortando.",
				EN: "Disables the winepipewire audio driver (default in proton-cachyos) and falls back to winepulse. Useful when a game's audio crackles or cuts out.",
			},
		},
		{
			Command: "PROTON_USE_SDL=1 %command%",
			Title: Localized{
				PT: "Input via SDL",
				EN: "SDL input",
			},
			Category: Localized{
				PT: "Input",
				EN: "Input",
			},
			Compat: Localized{
				PT: "GE e CachyOS",
				EN: "GE and CachyOS",
			},
			Description: Localized{
				PT: "Usa o input do SDL em vez de HIDRAW/Steam Input. Resolve problemas de controle que não é detectado ou se comporta mal (comum com o driver Wayland).",
				EN: "Uses SDL input instead of HIDRAW/Steam Input. Fixes controllers that are not detected or behave badly (common with the Wayland driver).",
			},
		},
		{
			Command:   "HOST_LC_ALL=pt_BR.UTF-8 %command%",
			CommandEN: "HOST_LC_ALL=en_US.UTF-8 %command%",
			Title: Localized{
				PT: "Idioma do jogo",
				EN: "Game language",
			},
			Category: Localized{
				PT: "Outros",
				EN: "Other",
			},
			Compat: Localized{
				PT: "Todos",
				EN: "All",
			},
			Description: Localized{
				PT: "Força um locale específico para o jogo, sobrescrevendo todos os outros ajustes de idioma. Troque pt_BR.UTF-8 pelo locale desejado (ex.: en_US.UTF-8) para jogos que pegam o idioma errado.",
				EN: "Forces a specific locale for the game, overriding all other language settings. Replace pt_BR.UTF-8 with the desired locale (e.g.: en_US.UTF-8) for games that pick the wrong language.",
			},
		},
		{
			Command: "PROTON_USE_OPTISCALER=1 %command%",
			Title: Localized{
				PT: "OptiScaler (injeção automática)",
				EN: "OptiScaler (auto injection)",
			},
			Category: Localized{
				PT: "Upscaling",
				EN: "Upscaling",
			},
			Compat: Localized{
				PT: "CachyOS",
				EN: "CachyOS",
			},
			Description: Localized{
				PT: "Ativa a injeção automática do OptiScaler, permitindo FSR/DLSS/XeSS em jogos sem suporte nativo. Configure a DLL com PROTON_OPTISCALER_NAME (dxgi.dll, d3d11.dll, d3d12.dll) e o upscaler via PROTON_OPTISCALER_CONFIG=\"Upscalers.Dx11Upscaler=fsr31;Upscalers.Dx12Upscaler=dlss\". Em desenvolvimento: nem todos os jogos funcionam.",
				EN: "Enables automatic OptiScaler injection, letting you use FSR/DLSS/XeSS in games without native support. Configure DLL with PROTON_OPTISCALER_NAME (dxgi.dll, d3d11.dll, d3d12.dll) and upscaler via PROTON_OPTISCALER_CONFIG=\"Upscalers.Dx11Upscaler=fsr31;Upscalers.Dx12Upscaler=dlss\". Work in progress: not all games work.",
			},
		},
		{
			Command: "ENABLE_LAYER_MESA_ANTI_LAG=1 %command%",
			Title: Localized{
				PT: "Anti-Lag 2 (camada do Mesa)",
				EN: "Anti-Lag 2 (Mesa layer)",
			},
			Category: Localized{
				PT: "Latência",
				EN: "Latency",
			},
			Compat: Localized{
				PT: "Mesa 25.3+ com a camada anti-lag instalada (nem toda distro empacota)",
				EN: "Mesa 25.3+ with the anti-lag layer installed (not every distro ships it)",
			},
			Description: Localized{
				PT: "Liga a camada implícita VK_LAYER_MESA_anti_lag, a implementação open-source da extensão VK_AMD_anti_lag que o Mesa traz desde a 25.3. Duas ressalvas: (1) a camada é OPCIONAL — precisa estar em /usr/share/vulkan/implicit_layer.d/VkLayer_MESA_anti_lag.json, e vários builds não a incluem (runtime Flatpak do Mesa, por exemplo), então a variável não faz nada; confira com vulkaninfo; (2) ela não é global — só age em jogos que realmente chamam vkAntiLagUpdateAMD, ou seja, que implementam a extensão, muito menos jogos que o número de menus de Anti-Lag sugere. Nos testes do low_latency_layer ela chegou a parecer no-op. Desligue com DISABLE_LAYER_MESA_ANTI_LAG=1 — e note que o Proton-CachyOS já faz isso sozinho quando PROTON_FSR4_UPGRADE está ativo, porque as duas brigam. O RADV agora recebe implementação nativa de VK_AMD_anti_lag (MR 42048), o que torna essa camada desnecessária nas versões que já a trouxeram — prefira LOW_LATENCY_LAYER=1 com LOW_LATENCY_LAYER_REFLEX=1, que tem o mesmo efeito e bem mais jogos.",
				EN: "Enables the implicit VK_LAYER_MESA_anti_lag layer, the open-source implementation of the VK_AMD_anti_lag extension that has shipped with Mesa since 25.3. Two caveats: (1) the layer is OPTIONAL — it must be present as /usr/share/vulkan/implicit_layer.d/VkLayer_MESA_anti_lag.json, and many builds don't include it (the Mesa Flatpak runtime, for example), so the variable does nothing; check with vulkaninfo; (2) it is not global — it only acts on games that actually call vkAntiLagUpdateAMD, i.e. those implementing the extension, far fewer than the number of Anti-Lag menus suggests. In low_latency_layer's benchmarks it looked like a no-op. Turn it off with DISABLE_LAYER_MESA_ANTI_LAG=1 — and note that Proton-CachyOS already does that on its own whenever PROTON_FSR4_UPGRADE is active, because the two clash. RADV now gets a native VK_AMD_anti_lag implementation (MR 42048), which makes this layer unnecessary on versions that include it — prefer LOW_LATENCY_LAYER=1 with LOW_LATENCY_LAYER_REFLEX=1, same effect and far more games.",
			},
		},
		{
			Command: "LOW_LATENCY_LAYER=1 %command%",
			Title: Localized{
				PT: "AMD Anti-Lag 2 (qualquer GPU)",
				EN: "AMD Anti-Lag 2 (any GPU)",
			},
			Category: Localized{
				PT: "Latência",
				EN: "Latency",
			},
			Compat: Localized{
				PT: "Todos (GPU AMD/Intel; requer low_latency_layer)",
				EN: "All (AMD/Intel GPU; requires low_latency_layer)",
			},
			Description: Localized{
				PT: "Ativa o low_latency_layer, que expõe a extensão VK_AMD_anti_lag em GPUs AMD e Intel — o Anti-Lag 2 passa a funcionar em jogos Vulkan (CS2 nativo, e via dxvk-nvapi em jogos Proton com proton-cachyos/GE, que já embutem o layer). Use para reduzir o input lag em jogos competitivos. Desative com DISABLE_LOW_LATENCY_LAYER=1 se causar travamentos. Notas: o Reflex tem a mesma performance e funciona em muito mais jogos, então prefira a entrada de Reflex; em Cyberpunk 2077 o Anti-Lag 2 não funciona por bug do próprio jogo (a camada nunca recebe a chamada) — use o caminho do Reflex; em Marvel Rivals é preciso LOW_LATENCY_LAYER_FORCE_DECOUPLED=1 (o layer já aplica sozinho nesse título); LOW_LATENCY_LAYER_SPOOF_NVIDIA=1 é uma alternativa mais suave ao PROTON_FORCE_NVAPI=1 para expor o menu.",
				EN: "Enables low_latency_layer, which exposes the VK_AMD_anti_lag extension on AMD and Intel GPUs — Anti-Lag 2 now works in Vulkan games (native CS2, and via dxvk-nvapi in Proton games with proton-cachyos/GE, which already bundle the layer). Use to reduce input lag in competitive games. Disable with DISABLE_LOW_LATENCY_LAYER=1 if it causes crashes. Notes: Reflex has the same performance and works in far more games, so prefer the Reflex entry; in Cyberpunk 2077 Anti-Lag 2 doesn't work due to a game bug (the layer never gets the call) — use the Reflex path; Marvel Rivals needs LOW_LATENCY_LAYER_FORCE_DECOUPLED=1 (the layer already does it for that title on its own); LOW_LATENCY_LAYER_SPOOF_NVIDIA=1 is a gentler alternative to PROTON_FORCE_NVAPI=1 for exposing the menu.",
			},
		},
		{
			Command: "LOW_LATENCY_LAYER_SPOOF_NVIDIA=1 %command%",
			Title: Localized{
				PT: "Spoofing NVIDIA no low_latency_layer",
				EN: "NVIDIA spoof in low_latency_layer",
			},
			Category: Localized{
				PT: "Latência",
				EN: "Latency",
			},
			Compat: Localized{
				PT: "Todos (requer LOW_LATENCY_LAYER=1)",
				EN: "All (requires LOW_LATENCY_LAYER=1)",
			},
			Description: Localized{
				PT: "Faz o low_latency_layer reportar a GPU como NVIDIA para a aplicação, independente do hardware real. É a alternativa mais suave ao PROTON_FORCE_NVAPI=1 para destravar o menu Reflex em GPU AMD, porque não passa pelo WINE_HIDE_AMD_GPU. Use junto de LOW_LATENCY_LAYER=1 e LOW_LATENCY_LAYER_REFLEX=1. O próprio autor marca como não recomendado, então trate como último recurso — e os mesmos avisos valem: quebra o upgrade FSR4 (PROTON_FSR4_UPGRADE) e não é à prova de anti-cheat. Compare sempre antes com DXVK_CONFIG=\"dxgi.hideAmdGpu = True\", que faz o mesmo pelo DXVK sem mexer na camada.",
				EN: "Makes low_latency_layer report the GPU as NVIDIA to the application, whatever the real hardware. It's the gentler alternative to PROTON_FORCE_NVAPI=1 for unlocking the Reflex menu on AMD, because it doesn't go through WINE_HIDE_AMD_GPU. Use it together with LOW_LATENCY_LAYER=1 and LOW_LATENCY_LAYER_REFLEX=1. The author marks it as not recommended, so treat it as a last resort — and the same warnings apply: it breaks the FSR4 upgrade (PROTON_FSR4_UPGRADE) and isn't anti-cheat safe. Always compare it first with DXVK_CONFIG=\"dxgi.hideAmdGpu = True\", which does the same thing through DXVK without touching the layer.",
			},
		},
		{
			Command: "LOW_LATENCY_LAYER_FORCE_DECOUPLED=1 %command%",
			Title: Localized{
				PT: "Fila de simulação separada (Marvel Rivals)",
				EN: "Decoupled sim queue (Marvel Rivals)",
			},
			Category: Localized{
				PT: "Latência",
				EN: "Latency",
			},
			Compat: Localized{
				PT: "Todos (requer LOW_LATENCY_LAYER=1; Marvel Rivals)",
				EN: "All (requires LOW_LATENCY_LAYER=1; Marvel Rivals)",
			},
			Description: Localized{
				PT: "Força a mitigação de fila de simulação e render desacopladas no low_latency_layer. Existe por causa do Marvel Rivals, que usa essa arquitetura e precisa de estatísticas e de atraso extra para bater a implementação do Windows. O layer já aplica isso automaticamente só nesse título, então esta variável serve para os outros jogos UE5 com o mesmo problema — ou para forçar quando a detecção falha. Só faz sentido junto de LOW_LATENCY_LAYER=1.",
				EN: "Forces low_latency_layer to mitigate a decoupled simulation and render queue. It exists because of Marvel Rivals, which uses that architecture and needs extra statistics and delay to match the Windows implementation. The layer already does this automatically for that title only, so this variable is for other UE5 games with the same problem — or to force it when detection fails. Only meaningful together with LOW_LATENCY_LAYER=1.",
			},
		},
		{
			Command: "VKD3D_FRAME_RATE=60 %command%",
			Title: Localized{
				PT: "Limite de FPS (D3D12, vkd3d-low-latency)",
				EN: "FPS cap (D3D12, vkd3d-low-latency)",
			},
			Category: Localized{
				PT: "Desempenho",
				EN: "Performance",
			},
			Compat: Localized{
				PT: "CachyOS (requer PROTON_VKD3D_LOWLATENCY=1)",
				EN: "CachyOS (requires PROTON_VKD3D_LOWLATENCY=1)",
			},
			Description: Localized{
				PT: "Limita o FPS dos jogos D3D12, equivalente ao DXVK_FRAME_RATE do outro lado. Só existe dentro do fork vkd3d-low-latency, então não combine sem PROTON_VKD3D_LOWLATENCY=1. Atenção: este limitador tem prioridade sobre o fps cap do Reflex e sobre o in-game, e o mantenedor é explícito — sobrepor dois limitadores é justamente o que piora a latência, então use um só. Em jogos UE4 com swapchain aguardável ele também evita a queda de desempenho do r.OneFrameThreadLag=1 (veja a entrada do PROTON_VKD3D_LOWLATENCY).",
				EN: "Caps FPS in D3D12 games, the D12 counterpart of DXVK_FRAME_RATE. It only exists inside the vkd3d-low-latency fork, so don't use it without PROTON_VKD3D_LOWLATENCY=1. Note that this limiter takes priority over the Reflex fps cap and over the in-game one, and the maintainer is blunt about it — layering two limiters is exactly what makes latency worse, so pick one. In UE4 games with waitable swapchains it also avoids the r.OneFrameThreadLag=1 performance drop (see the PROTON_VKD3D_LOWLATENCY entry).",
			},
		},
		{
			Command: "LOW_LATENCY_LAYER=1 LOW_LATENCY_LAYER_REFLEX=1 DXVK_CONFIG=\"dxgi.hideAmdGpu = True\" %command%",
			Title: Localized{
				PT: "NVIDIA Reflex em qualquer GPU",
				EN: "NVIDIA Reflex on any GPU",
			},
			Category: Localized{
				PT: "Latência",
				EN: "Latency",
			},
			Compat: Localized{
				PT: "Todos (requer low_latency_layer; jogos com Reflex)",
				EN: "All (requires low_latency_layer; Reflex games)",
			},
			Description: Localized{
				PT: "Faz o low_latency_layer expor VK_NV_low_latency2 em vez de anti-lag, ativando o Reflex: mesma latência do Anti-Lag 2, mas com suporte em muito mais jogos (é a via preferida). O spoofing de GPU é um recurso de último recurso e costuma falhar, então tente nesta ordem: LOW_LATENCY_LAYER_REFLEX=1 sozinho; depois acrescente DXVK_CONFIG=\"dxgi.hideAmdGpu = True\" (às vezes também DXVK_NVAPI_ALLOW_OTHER_DRIVERS=1); e por último PROTON_FORCE_NVAPI=1 e/ou LOW_LATENCY_LAYER_SPOOF_NVIDIA=1. Se mesmo assim o jogo recusar a API, o mantenedor do vkd3d-low-latency sugere forjar o device inteiro por DXVK_CONFIG: dxgi.customVendorId=10de, dxgi.customDeviceId=2206, dxgi.customDeviceDesc=\"NVIDIA GeForce RTX 3080\" (também dá para pôr num dxvk.conf na pasta do jogo). AVISOS: PROTON_FORCE_NVAPI=1 e LOW_LATENCY_LAYER_SPOOF_NVIDIA=1 quebram o upgrade automático do FSR4 (PROTON_FSR4_UPGRADE) — no FSR4 4.1.1+ o jogo cai para FSR3, e em RDNA3 isso costuma custar bastante desempenho, então vale testar FSR 4.1.0 ou inferior; o spoofing também não é à prova de falhas, jogos com verificações extras de GPU (THE FINALS tem uma exceção no próprio Proton, Test Drive: Solar Crown checa o device) podem recusar a API mesmo com o menu visível; e em jogos com anti-cheat o spoofing pode ser barrado ou causar kick — não use com anti-cheat. Confira se está realmente funcionando procurando \"nvapi64:<-NvAPI_D3D_SetSleepMode (Enabled/0us): OK\" no log com PROTON_LOG=1; o menu aparecer não garante que esteja ativo.",
				EN: "Makes low_latency_layer expose VK_NV_low_latency2 instead of anti-lag, enabling Reflex: the same latency as Anti-Lag 2 but supported in far more games (it's the preferred path). GPU spoofing is a last resort and isn't foolproof, so try in this order: LOW_LATENCY_LAYER_REFLEX=1 alone; then add DXVK_CONFIG=\"dxgi.hideAmdGpu = True\" (sometimes DXVK_NVAPI_ALLOW_OTHER_DRIVERS=1 too); and only then PROTON_FORCE_NVAPI=1 and/or LOW_LATENCY_LAYER_SPOOF_NVIDIA=1. If the game still refuses the API, the vkd3d-low-latency maintainer suggests forging the whole device via DXVK_CONFIG: dxgi.customVendorId=10de, dxgi.customDeviceId=2206, dxgi.customDeviceDesc=\"NVIDIA GeForce RTX 3080\" (you can also drop these in a dxvk.conf next to the game). WARNINGS: PROTON_FORCE_NVAPI=1 and LOW_LATENCY_LAYER_SPOOF_NVIDIA=1 break the FSR4 automatic upgrade (PROTON_FSR4_UPGRADE) — on FSR4 4.1.1+ the game falls back to FSR3, and on RDNA3 that usually costs a lot of performance, so consider FSR 4.1.0 or older; spoofing is also not foolproof — games with extra GPU checks (THE FINALS has an exception in Proton itself, Test Drive: Solar Crown checks the device) may refuse the API even with the menu visible; and in games with anti-cheat spoofing may be blocked or get you kicked — don't use it with anti-cheat. Verify it's actually working by looking for \"nvapi64:<-NvAPI_D3D_SetSleepMode (Enabled/0us): OK\" in the PROTON_LOG=1 log; the menu showing up doesn't guarantee it's active.",
			},
		},
		{
			Command: "WINE_FULLSCREEN_FSR_STRENGTH=2 %command%",
			Title: Localized{
				PT: "Nitidez do FSR",
				EN: "FSR sharpness",
			},
			Category: Localized{
				PT: "Upscaling",
				EN: "Upscaling",
			},
			Compat: Localized{
				PT: "Todos (use junto com WINE_FULLSCREEN_FSR=1)",
				EN: "All (use with WINE_FULLSCREEN_FSR=1)",
			},
			Description: Localized{
				PT: "Ajusta a nitidez do FSR em fullscreen: 0 = nitidez máxima, 5 = mínima (2 é o recomendado pela AMD). Sem essa variável, o FSR usa o valor padrão.",
				EN: "Adjusts FSR sharpness in fullscreen: 0 = maximum sharpness, 5 = minimum (2 is AMD's recommendation). Without this variable, FSR uses the default value.",
			},
		},
		{
			Command: "VKD3D_CONFIG=dxr %command%",
			Title: Localized{
				PT: "Ray tracing (DXR) via vkd3d",
				EN: "Ray tracing (DXR) via vkd3d",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "Todos (GPU com suporte a ray tracing)",
				EN: "All (ray tracing capable GPU)",
			},
			Description: Localized{
				PT: "Ativa o suporte a ray tracing (DXR) nos jogos Direct3D 12 que usam vkd3d-proton. Requer GPU com ray tracing habilitado no driver. Outras opções podem ser separadas por vírgula (ex.: dxr,force_bindless_texel_buffer).",
				EN: "Enables ray tracing (DXR) support in Direct3D 12 games using vkd3d-proton. Requires a ray tracing capable GPU with RT enabled in the driver. Other options can be comma separated (e.g.: dxr,force_bindless_texel_buffer).",
			},
		},
		{
			Command: "WINEDLLOVERRIDES=\"dinput8=n,b\" %command%",
			Title: Localized{
				PT: "Overrides de DLL (WINEDLLOVERRIDES)",
				EN: "DLL overrides (WINEDLLOVERRIDES)",
			},
			Category: Localized{
				PT: "Outros",
				EN: "Other",
			},
			Compat: Localized{
				PT: "Todos",
				EN: "All",
			},
			Description: Localized{
				PT: "Controla de onde cada DLL do jogo vem: n = nativa (a do próprio jogo, usada por mods/wrappers como dinput8 em jogos antigos), b = builtin (a do Proton). Ex.: WINEDLLOVERRIDES=\"dinput8=n,b\" força a DLL nativa do jogo antes da builtin. Separe várias DLLs com ponto-e-vírgula. Útil para fixar controles ou contornar bibliotecas quebradas.",
				EN: "Controls where each game DLL comes from: n = native (the game's own, used by mods/wrappers like dinput8 in old games), b = builtin (Proton's). E.g.: WINEDLLOVERRIDES=\"dinput8=n,b\" forces the game's native DLL before the builtin. Separate multiple DLLs with semicolons. Useful to fix controllers or bypass broken libraries.",
			},
		},
		{
			Command: "PROTON_NO_D3D10=1 %command%",
			Title: Localized{
				PT: "Desabilitar D3D10",
				EN: "Disable D3D10",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "GE e CachyOS",
				EN: "GE and CachyOS",
			},
			Description: Localized{
				PT: "Desativa a d3d10.dll e a dxgi.dll, para jogos D3D10 que conseguem cair para D3D9 com mais desempenho. Se o jogo não tiver fallback, não abre.",
				EN: "Disables d3d10.dll and dxgi.dll, for D3D10 games that can fall back to D3D9 with better performance. If the game has no fallback, it won't start.",
			},
		},
		{
			Command: "RADV_DEBUG=nofastclears %command%",
			Title: Localized{
				PT: "RADV: corrigir artefatos (nofastclears)",
				EN: "RADV: fix artifacts (nofastclears)",
			},
			Category: Localized{
				PT: "GPU",
				EN: "GPU",
			},
			Compat: Localized{
				PT: "Todos (GPU AMD / Mesa RADV)",
				EN: "All (AMD GPU / Mesa RADV)",
			},
			Description: Localized{
				PT: "Desativa os fast clears no driver RADV, corrigindo artefatos visuais (tela piscando, linhas estranhas) em alguns jogos AMD. Se o jogo sumir no HUD, é sintoma de fast clear.",
				EN: "Disables fast clears in the RADV driver, fixing visual artifacts (flickering, weird lines) in some AMD games. If a game disappears from the HUD, it's a fast clear symptom.",
			},
		},
		{
			Command: "MESA_VK_WSI_PRESENT_MODE=mailbox %command%",
			Title: Localized{
				PT: "Mesa: modo de apresentação Vulkan",
				EN: "Mesa: Vulkan present mode",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "Todos (Mesa Vulkan)",
				EN: "All (Mesa Vulkan)",
			},
			Description: Localized{
				PT: "Força o modo de apresentação do driver Mesa para todos os jogos Vulkan: mailbox (sem vsync, baixa latência, sem tearing quando o compositor suporta) ou immediate (vsync desligado de vez). Opções: fifo (vsync), mailbox, immediate.",
				EN: "Forces the Mesa driver's present mode for all Vulkan games: mailbox (no vsync, low latency, no tearing when the compositor supports it) or immediate (vsync fully off). Options: fifo (vsync), mailbox, immediate.",
			},
		},
		{
			Command: "__GL_SHADER_DISK_CACHE_SKIP_CLEANUP=1 %command%",
			Title: Localized{
				PT: "NVIDIA: cache de shaders sem limpeza",
				EN: "NVIDIA: shader cache without cleanup",
			},
			Category: Localized{
				PT: "GPU",
				EN: "GPU",
			},
			Compat: Localized{
				PT: "NVIDIA",
				EN: "NVIDIA",
			},
			Description: Localized{
				PT: "Impede o driver NVIDIA de limpar o cache de shaders ao sair do jogo, evitando stutter recorrente na primeira execução. Recomendado pela wiki do CachyOS. Lembre de limpar o cache manualmente de tempos em tempos.",
				EN: "Prevents the NVIDIA driver from cleaning the shader cache on game exit, avoiding recurring stutter on first runs. Recommended by the CachyOS wiki. Remember to clean the cache manually from time to time.",
			},
		},
		{
			Command: "PROTON_USE_WOW64=1 %command%",
			Title: Localized{
				PT: "Prefixos WOW64",
				EN: "WOW64 prefixes",
			},
			Category: Localized{
				PT: "Outros",
				EN: "Other",
			},
			Compat: Localized{
				PT: "GE e CachyOS",
				EN: "GE and CachyOS",
			},
			Description: Localized{
				PT: "Usa o modo wow64 (novo modelo de prefixo do Wine, 32 e 64 bits sem camada de tradução separada). Mais moderno e em alguns casos com melhor desempenho, mas exige criar um prefixo novo (trocar a versão do Proton no jogo).",
				EN: "Uses wow64 mode (Wine's new prefix model, 32 and 64 bit without a separate translation layer). More modern and sometimes faster, but requires a fresh prefix (switch the game's Proton version).",
			},
		},
		{
			Command: "PROTON_DISCORD_BRIDGE=1 %command%",
			Title: Localized{
				PT: "Rich Presence no Discord",
				EN: "Discord Rich Presence",
			},
			Category: Localized{
				PT: "Outros",
				EN: "Other",
			},
			Compat: Localized{
				PT: "CachyOS",
				EN: "CachyOS",
			},
			Description: Localized{
				PT: "Ativa o rpc-bridge, que permite jogos rodando no Proton exibirem Rich Presence (\"jogando X\") no Discord.",
				EN: "Enables rpc-bridge, letting games running in Proton show Rich Presence (\"playing X\") on Discord.",
			},
		},
		{
			Command:   "PROTON_WAIT_ATTACH=1 %command%",
			CommandEN: "PROTON_WAIT_ATTACH=1 %command%",
			Title: Localized{
				PT: "Esperar depurador anexar",
				EN: "Wait for debugger attach",
			},
			Category: Localized{
				PT: "Diagnóstico",
				EN: "Diagnostics",
			},
			Compat: Localized{
				PT: "Proton padrão",
				EN: "Standard Proton",
			},
			Description: Localized{
				PT: "Espera um depurador anexar ao steam.exe antes de iniciar o processo do jogo. Para anexar ao jogo no início, configure o depurador para seguir processos filhos.",
				EN: "Waits for a debugger to attach to steam.exe before launching the game process. To attach to the game at startup, set debuggers to follow child processes.",
			},
		},
		{
			Command:   "PROTON_DXVK_D3D8=1 %command%",
			CommandEN: "PROTON_DXVK_D3D8=1 %command%",
			Title: Localized{
				PT: "D3D8 via DXVK",
				EN: "D3D8 via DXVK",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "Proton padrão, GE e CachyOS",
				EN: "Standard Proton, GE and CachyOS",
			},
			Description: Localized{
				PT: "Usa o d3d8.dll do DXVK em vez do d3d8 do Wine para jogos Direct3D 8. Use quando um jogo D3D8 antigo trava ou não renderiza — o DXVK costuma ser mais compatível que a tradução nativa do Wine.",
				EN: "Uses DXVK's d3d8.dll instead of Wine's d3d8 for Direct3D 8 games. Use when an old D3D8 game crashes or fails to render — DXVK is usually more compatible than Wine's native translation.",
			},
		},
		{
			Command:   "PROTON_D7VK_DDRAW=1 %command%",
			CommandEN: "PROTON_D7VK_DDRAW=1 %command%",
			Title: Localized{
				PT: "D7VK para jogos DX7 ou anteriores",
				EN: "D7VK for DX7 and older games",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "CachyOS",
				EN: "CachyOS",
			},
			Description: Localized{
				PT: "Usa o ddraw.dll do D7VK nos jogos DirectDraw/Direct3D 7 ou anteriores, em vez da implementação do Wine. Use quando um clássico trava, pisca ou não renderiza.",
				EN: "Uses D7VK's ddraw.dll in DirectDraw/Direct3D 7 and older games instead of Wine's implementation. Use when a classic crashes, flickers or fails to render.",
			},
		},
		{
			Command:   "PROTON_NO_NTSYNC=1 %command%",
			CommandEN: "PROTON_NO_NTSYNC=1 %command%",
			Title: Localized{
				PT: "Desativar ntsync",
				EN: "Disable ntsync",
			},
			Category: Localized{
				PT: "Sincronização",
				EN: "Synchronization",
			},
			Compat: Localized{
				PT: "Proton padrão, GE e CachyOS",
				EN: "Standard Proton, GE and CachyOS",
			},
			Description: Localized{
				PT: "Desativa o ntsync (sincronização no estilo Windows NT). Use se algum jogo apresentar problema com o ntsync ativo. O ntsync agora é o padrão no Proton 11+.",
				EN: "Disables ntsync (Windows NT style synchronization). Use it if a game misbehaves with ntsync enabled. ntsync is now the default in Proton 11+.",
			},
		},
		{
			Command:   "PROTON_FORCE_LARGE_ADDRESS_AWARE=1 %command%",
			CommandEN: "PROTON_FORCE_LARGE_ADDRESS_AWARE=1 %command%",
			Title: Localized{
				PT: "Forçar LARGE_ADDRESS_AWARE",
				EN: "Force LARGE_ADDRESS_AWARE",
			},
			Category: Localized{
				PT: "Outros",
				EN: "Other",
			},
			Compat: Localized{
				PT: "Proton padrão, GE e CachyOS",
				EN: "Standard Proton, GE and CachyOS",
			},
			Description: Localized{
				PT: "Força a flag LARGE_ADDRESS_AWARE em todos os executáveis, permitindo usar mais de 2 GB de RAM (já habilitado por padrão no Proton).",
				EN: "Forces the LARGE_ADDRESS_AWARE flag on all executables, allowing more than 2 GB of RAM (enabled by default in Proton).",
			},
		},
		{
			Command:   "PROTON_HEAP_DELAY_FREE=1 %command%",
			CommandEN: "PROTON_HEAP_DELAY_FREE=1 %command%",
			Title: Localized{
				PT: "Atrasar liberação de memória",
				EN: "Delay memory freeing",
			},
			Category: Localized{
				PT: "Outros",
				EN: "Other",
			},
			Compat: Localized{
				PT: "Proton padrão, GE e CachyOS",
				EN: "Standard Proton, GE and CachyOS",
			},
			Description: Localized{
				PT: "Atrasa a liberação de parte da memória para contornar bugs de use-after-free em alguns jogos.",
				EN: "Delays freeing some memory to work around use-after-free bugs in some games.",
			},
		},
		{
			Command:   "PROTON_SET_GAME_DRIVE=1 %command%",
			CommandEN: "PROTON_SET_GAME_DRIVE=1 %command%",
			Title: Localized{
				PT: "Unidade S: do jogo",
				EN: "Game S: drive",
			},
			Category: Localized{
				PT: "Outros",
				EN: "Other",
			},
			Compat: Localized{
				PT: "Proton padrão, GE e CachyOS",
				EN: "Standard Proton, GE and CachyOS",
			},
			Description: Localized{
				PT: "Cria uma unidade S: apontando para a biblioteca Steam que contém o jogo. Útil para jogos que procuram arquivos em caminhos fixos.",
				EN: "Creates an S: drive pointing to the Steam Library which contains the game. Useful for games that look for files in fixed paths.",
			},
		},
		{
			Command:   "PROTON_OLD_GL_STRING=1 %command%",
			CommandEN: "PROTON_OLD_GL_STRING=1 %command%",
			Title: Localized{
				PT: "String GL limitada",
				EN: "Limit GL extension string",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "Proton padrão, GE e CachyOS",
				EN: "Standard Proton, GE and CachyOS",
			},
			Description: Localized{
				PT: "Aplica overrides no driver para limitar o tamanho da string de extensões GL, para jogos antigos que travam com strings muito longas.",
				EN: "Sets driver overrides to limit the length of the GL extension string, for old games that crash on very long extension strings.",
			},
		},
		{
			Command:   "WINE_DO_NOT_CREATE_DXGI_DEVICE_MANAGER=1 %command%",
			CommandEN: "WINE_DO_NOT_CREATE_DXGI_DEVICE_MANAGER=1 %command%",
			Title: Localized{
				PT: "Sem DXGI device manager",
				EN: "Skip DXGI device manager",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "Proton padrão, GE e CachyOS",
				EN: "Standard Proton, GE and CachyOS",
			},
			Description: Localized{
				PT: "Hack para contornar problemas de vídeo em alguns jogos causados por suporte incompleto a IMFDXGIDeviceManager.",
				EN: "Hack to work around video issues in some games due to incomplete IMFDXGIDeviceManager support.",
			},
		},
		{
			Command:   "WINE_DISABLE_VULKAN_OPWR=1 %command%",
			CommandEN: "WINE_DISABLE_VULKAN_OPWR=1 %command%",
			Title: Localized{
				PT: "Sem Vulkan OPWR",
				EN: "Disable Vulkan OPWR",
			},
			Category: Localized{
				PT: "Wayland",
				EN: "Wayland",
			},
			Compat: Localized{
				PT: "Wayland",
				EN: "Wayland",
			},
			Description: Localized{
				PT: "Desativa o render de janelas de outros processos via Vulkan (other process window rendering), que às vezes causa atraso de um frame no Wayland.",
				EN: "Disables Vulkan other process window rendering, which sometimes causes issues on Wayland due to blit being one frame behind.",
			},
		},
		{
			Command:   "PROTON_HIDE_NVIDIA_GPU=1 %command%",
			CommandEN: "PROTON_HIDE_NVIDIA_GPU=1 %command%",
			Title: Localized{
				PT: "Ocultar GPU NVIDIA",
				EN: "Hide NVIDIA GPU",
			},
			Category: Localized{
				PT: "GPU",
				EN: "GPU",
			},
			Compat: Localized{
				PT: "NVIDIA",
				EN: "NVIDIA",
			},
			Description: Localized{
				PT: "Faz a GPU NVIDIA ser sempre reportada como AMD. Alguns jogos exigem isso quando dependem de funcionalidade do driver NVIDIA que só existe no Windows.",
				EN: "Forces NVIDIA GPUs to always be reported as AMD GPUs. Some games require this if they depend on Windows-only NVIDIA driver functionality.",
			},
		},
		{
			Command:   "WINE_USE_KWIN_HACKS=1 %command%",
			CommandEN: "WINE_USE_KWIN_HACKS=1 %command%",
			Title: Localized{
				PT: "Hacks para KDE",
				EN: "KDE windowing hacks",
			},
			Category: Localized{
				PT: "Wayland",
				EN: "Wayland",
			},
			Compat: Localized{
				PT: "Proton padrão (KDE)",
				EN: "Standard Proton (KDE)",
			},
			Description: Localized{
				PT: "Ativa hacks específicos do KDE que melhoram a experiência com KDE mais antigo que 6.4 no Wayland e 6.6 no X11.",
				EN: "Enables KDE-specific windowing hacks that may improve experience with KDE older than 6.4 on Wayland and 6.6 on X11.",
			},
		},
		{
			Command:   "PROTON_USE_XALIA=1 %command%",
			CommandEN: "PROTON_USE_XALIA=1 %command%",
			Title: Localized{
				PT: "Xalia (UI de gamepad)",
				EN: "Xalia (gamepad UI)",
			},
			Category: Localized{
				PT: "Input",
				EN: "Input",
			},
			Compat: Localized{
				PT: "Proton padrão, GE e CachyOS",
				EN: "Standard Proton, GE and CachyOS",
			},
			Description: Localized{
				PT: "Ativa o Xalia, que adiciona UI de gamepad para algumas interfaces de teclado/mouse. Por padrão o Proton decide dinamicamente; use 0 para desativar.",
				EN: "Enables Xalia, which adds a gamepad UI for some keyboard/mouse interfaces. The default is dynamic; set to 0 to disable.",
			},
		},
		{
			Command:   "MANGOHUD=1 %command%",
			CommandEN: "MANGOHUD=1 %command%",
			Title: Localized{
				PT: "MangoHud via variável de ambiente",
				EN: "MangoHud via environment variable",
			},
			Category: Localized{
				PT: "Overlay e desempenho",
				EN: "Overlay & Performance",
			},
			Compat: Localized{
				PT: "Todos (exige mangohud instalado)",
				EN: "All (requires mangohud)",
			},
			Description: Localized{
				PT: "Alternativa ao wrapper mangohud: injeta o MangoHud via variável de ambiente. Útil em launchers que não aceitam wrappers (Heroic, Bottles) ou quando você precisa combinar com outros wrappers.",
				EN: "Alternative to the mangohud wrapper: injects MangoHud via an environment variable. Useful in launchers that don't accept wrappers (Heroic, Bottles) or when you need to combine with other wrappers.",
			},
		},
		{
			Command:   "MANGOHUD_CONFIG=cpu_temp,gpu_temp,fps,frametime %command%",
			CommandEN: "MANGOHUD_CONFIG=cpu_temp,gpu_temp,fps,frametime %command%",
			Title: Localized{
				PT: "MangoHud: configuração via env",
				EN: "MangoHud: config via env",
			},
			Category: Localized{
				PT: "Overlay e desempenho",
				EN: "Overlay & Performance",
			},
			Compat: Localized{
				PT: "Todos (exige mangohud instalado)",
				EN: "All (requires mangohud)",
			},
			Description: Localized{
				PT: "Define as métricas do MangoHud por jogo, sem mexer no MangoHud.conf global. Exemplo: cpu_temp,gpu_temp,fps,frametime. O MANGOHUD_CONFIG global do jogo tem prioridade sobre o do arquivo. Separe os itens com vírgula.",
				EN: "Sets per-game MangoHud metrics without touching the global MangoHud.conf. Example: cpu_temp,gpu_temp,fps,frametime. Per-game MANGOHUD_CONFIG takes priority over the file's. Separate items with commas.",
			},
		},
		{
			Command:   "PROTON_FORCE_NVAPI=1 %command%",
			CommandEN: "PROTON_FORCE_NVAPI=1 %command%",
			Title: Localized{
				PT: "Forçar NVAPI (com patente do Reflex)",
				EN: "Force NVAPI (with Reflex patents)",
			},
			Category: Localized{
				PT: "GPU",
				EN: "GPU",
			},
			Compat: Localized{
				PT: "GE e CachyOS (GPU NVIDIA ou AMD com spoofing)",
				EN: "GE and CachyOS (NVIDIA or spoofed AMD GPU)",
			},
			Description: Localized{
				PT: "Habilita o suporte a patentes do NVIDIA Reflex no dxvk-nvapi. Útil quando o menu de Reflex não aparece em jogos com suporte. No script do Proton isso define três variáveis de uma vez, sem condição: DXVK_NVAPI_ALLOW_OTHER_DRIVERS=1, DXVK_NVAPI_DRIVER_VERSION=99999 e WINE_HIDE_AMD_GPU=1. Ou seja, funciona em GPU AMD (é o que abre o Reflex lá), mas o WINE_HIDE_AMD_GPU é justamente o que impede o jogo de enxergar a AMD e por isso quebra o upgrade FSR 4 (PROTON_FSR4_UPGRADE) — sempre, não às vezes. Como o DXVK_NVAPI_ALLOW_OTHER_DRIVERS=1 já vem junto, prefira spoofing direto com DXVK_CONFIG=\"dxgi.hideAmdGpu=True\" DXVK_NVAPI_ALLOW_OTHER_DRIVERS=1, que expõe o mesmo Reflex sem o WINE_HIDE_AMD_GPU. Atenção também ao anti-cheat: forjar o device pode ser barrado. Evite em THE FINALS se possível — o Proton já desliga o NVAPI em GPU não-NVIDIA nesse título, e o PROTON_FORCE_NVAPI é a única forma de furar isso, ao custo do FSR 4.",
				EN: "Enables NVIDIA Reflex patent support in dxvk-nvapi. Useful when the Reflex menu doesn't show in supported games. The Proton script sets three variables at once, unconditionally: DXVK_NVAPI_ALLOW_OTHER_DRIVERS=1, DXVK_NVAPI_DRIVER_VERSION=99999 and WINE_HIDE_AMD_GPU=1. In other words it works on AMD GPUs (that's what gets Reflex showing there), but the WINE_HIDE_AMD_GPU is exactly what stops the game from seeing the AMD, and that's why it breaks the FSR 4 upgrade (PROTON_FSR4_UPGRADE) — always, not sometimes. Since DXVK_NVAPI_ALLOW_OTHER_DRIVERS=1 comes bundled, prefer spoofing directly with DXVK_CONFIG=\"dxgi.hideAmdGpu=True\" DXVK_NVAPI_ALLOW_OTHER_DRIVERS=1, which exposes the same Reflex without the WINE_HIDE_AMD_GPU. Watch out for anti-cheat too: faking the device can get blocked. Avoid it in THE FINALS if you can — Proton already disables NVAPI on non-NVIDIA GPUs for that title, and PROTON_FORCE_NVAPI is the only way around, at the cost of FSR 4.",
			},
		},
		{
			Command:   "PROTON_DISABLE_NVAPI=1 %command%",
			CommandEN: "PROTON_DISABLE_NVAPI=1 %command%",
			Title: Localized{
				PT: "Desabilitar NVAPI",
				EN: "Disable NVAPI",
			},
			Category: Localized{
				PT: "GPU",
				EN: "GPU",
			},
			Compat: Localized{
				PT: "GE (GPU NVIDIA)",
				EN: "GE (NVIDIA GPU)",
			},
			Description: Localized{
				PT: "Desativa a biblioteca NVAPI dentro do Proton. Use se o NVAPI (PROTON_ENABLE_NVAPI=1) estiver causando crashes ou bugs em algum jogo.",
				EN: "Disables the NVAPI library inside Proton. Use if NVAPI (PROTON_ENABLE_NVAPI=1) is causing crashes or bugs in some game.",
			},
		},
		{
			Command:   "ENABLE_HDR_WSI=1 %command%",
			CommandEN: "ENABLE_HDR_WSI=1 %command%",
			Title: Localized{
				PT: "HDR WSI (NVIDIA)",
				EN: "HDR WSI (NVIDIA)",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "NVIDIA (use com DXVK_HDR=1)",
				EN: "NVIDIA (use with DXVK_HDR=1)",
			},
			Description: Localized{
				PT: "Habilita a cadeia de swap HDR (WSI) no driver NVIDIA. Necessário para HDR em jogos Vulkan via DXVK quando o driver não ativa o HDR automaticamente.",
				EN: "Enables the HDR swapchain (WSI) on the NVIDIA driver. Required for HDR in Vulkan games via DXVK when the driver doesn't enable HDR automatically.",
			},
		},
		{
			Command:   "DXVK_LOG_LEVEL=none %command%",
			CommandEN: "DXVK_LOG_LEVEL=none %command%",
			Title: Localized{
				PT: "Silenciar log do DXVK",
				EN: "Silence DXVK log",
			},
			Category: Localized{
				PT: "Diagnóstico",
				EN: "Diagnostics",
			},
			Compat: Localized{
				PT: "Todos",
				EN: "All",
			},
			Description: Localized{
				PT: "Impede o DXVK de escrever o arquivo de log (d3d11.log / dxgi.log). Reduz escrita no disco e ganha um pouco de desempenho; útil quando o log não é necessário.",
				EN: "Prevents DXVK from writing its log file (d3d11.log / dxgi.log). Reduces disk writes and slightly improves performance; useful when the log isn't needed.",
			},
		},
		{
			Command:   "DXVK_FRAME_RATE=60 %command%",
			CommandEN: "DXVK_FRAME_RATE=60 %command%",
			Title: Localized{
				PT: "Limite de FPS (DXVK)",
				EN: "FPS cap (DXVK)",
			},
			Category: Localized{
				PT: "Overlay e desempenho",
				EN: "Overlay & Performance",
			},
			Compat: Localized{
				PT: "Todos",
				EN: "All",
			},
			Description: Localized{
				PT: "Limita o FPS dos jogos D3D11/D3D10/D3D9 (via DXVK) direto no driver, sem overlay. Troque 60 pelo limite desejado (ex.: 120, 144); -1 desativa. Útil em telas 60 Hz ou para reduzir consumo/ruído quando o jogo passa folgado do refresh. Não vale para D3D12 — nesses use VKD3D_FRAME_RATE, que só existe no vkd3d-low-latency. E cuidado: com PROTON_DXVK_LOWLATENCY=1 o limitador do fork entra em conflito com este, sobrepondo o fps cap do Reflex/V-Sync e normalmente piorando a latência. O próprio limitador do fork é o DXVK_FRAME_PACE=low-latency-vrr, que já vem com cap embutido.",
				EN: "Caps FPS in D3D11/D3D10/D3D9 games (via DXVK) right in the driver, with no overlay. Replace 60 with the desired cap (e.g.: 120, 144); -1 disables. Useful on 60 Hz displays or to cut power/noise when the game runs well past refresh. It does not apply to D3D12 — use VKD3D_FRAME_RATE there, which only exists in vkd3d-low-latency. And be careful: with PROTON_DXVK_LOWLATENCY=1 the fork's limiter clashes with this one, overriding the Reflex/V-Sync fps cap and usually making latency worse. The fork's own limiter is DXVK_FRAME_PACE=low-latency-vrr, which already ships with a cap baked in.",
			},
		},
		{
			Command:   "gamescope -e -f -F fsr -- %command%",
			CommandEN: "gamescope -e -f -F fsr -- %command%",
			Title: Localized{
				PT: "Gamescope exclusivo com FSR",
				EN: "Exclusive gamescope with FSR",
			},
			Category: Localized{
				PT: "Upscaling",
				EN: "Upscaling",
			},
			Compat: Localized{
				PT: "Todos (exige gamescope instalado)",
				EN: "All (requires gamescope)",
			},
			Description: Localized{
				PT: "Roda o jogo em modo exclusivo (sem compositor) dentro do gamescope com FSR integrado: o jogo renderiza na resolução interna e o FSR sobe a imagem para a resolução da tela. Ajuste a resolução com -w/-h (interna) e -W/-H (tela).",
				EN: "Runs the game in exclusive mode (no compositor) inside gamescope with built-in FSR: the game renders at its internal resolution and FSR upscales to the screen resolution. Tune with -w/-h (internal) and -W/-H (screen).",
			},
		},
		{
			Command:   "game-performance %command%",
			CommandEN: "game-performance %command%",
			Title: Localized{
				PT: "Modo de jogo do CachyOS",
				EN: "CachyOS game mode",
			},
			Category: Localized{
				PT: "Overlay e desempenho",
				EN: "Overlay & Performance",
			},
			Compat: Localized{
				PT: "CachyOS (requer cachyos-settings)",
				EN: "CachyOS (requires cachyos-settings)",
			},
			Description: Localized{
				PT: "Ativa o modo de desempenho do CachyOS enquanto o jogo roda: define o perfil de energia e o governor da CPU como \"performance\" durante o jogo e restaura ao fechar (via powerprofilesctl launch). Instale com: sudo pacman -S cachyos-settings. Para manter o screensaver ativo durante o jogo, use GAME_PERFORMANCE_SCREENSAVER_ON=1.",
				EN: "Enables CachyOS performance mode while the game runs: sets the power profile and CPU governor to \"performance\" during the game and restores on exit (via powerprofilesctl launch). Install with: sudo pacman -S cachyos-settings. To keep the screensaver active during the game, use GAME_PERFORMANCE_SCREENSAVER_ON=1.",
			},
		},
		{
			Command:   "LSFGVK_PROFILE=steam %command%",
			CommandEN: "LSFGVK_PROFILE=steam %command%",
			Title: Localized{
				PT: "lsfg-vk: perfil Steam (v2.0.0+)",
				EN: "lsfg-vk: Steam profile (v2.0.0+)",
			},
			Category: Localized{
				PT: "Overlay e desempenho",
				EN: "Overlay & Performance",
			},
			Compat: Localized{
				PT: "Todos (exige lsfg-vk 2.0+ instalado)",
				EN: "All (requires lsfg-vk 2.0+)",
			},
			Description: Localized{
				PT: "Seleciona o perfil \"steam\" do lsfg-vk (v2.0.0+), ativando o frame generation para jogos lançados pelo Steam. Variável oficial do lsfg-vk: LSFGVK_PROFILE substitui a detecção automática de processo. Prefira LSFGVK_PROFILE sobre a variante antiga LSFG_PROCESS, removida no v2.0.0. Requer Lossless Scaling ativado no Steam e lsfg-vk 2.0+ instalado.",
				EN: "Selects the lsfg-vk \"steam\" profile (v2.0.0+), enabling frame generation for Steam-launched games. Official lsfg-vk variable: LSFGVK_PROFILE overrides automatic process detection. Use LSFGVK_PROFILE instead of the removed LSFG_PROCESS (deprecated in v2.0.0). Requires Lossless Scaling enabled in Steam and lsfg-vk 2.0+ installed.",
			},
		},
		{
			Command:   "PROTON_FSR4_INDICATOR=1 %command%",
			CommandEN: "PROTON_FSR4_INDICATOR=1 %command%",
			Title: Localized{
				PT: "Indicador FSR 4 (watermark)",
				EN: "FSR 4 indicator (watermark)",
			},
			Category: Localized{
				PT: "Upscaling",
				EN: "Upscaling",
			},
			Compat: Localized{
				PT: "GE e CachyOS (GPU AMD)",
				EN: "GE and CachyOS (AMD GPU)",
			},
			Description: Localized{
				PT: "Mostra um watermark no canto da tela confirmando que o FSR 4 está ativo. Útil para verificar se o PROTON_FSR4_UPGRADE funcionou.",
				EN: "Shows a watermark in the corner confirming FSR 4 is active. Useful to verify PROTON_FSR4_UPGRADE worked.",
			},
		},
		{
			Command:   "PROTON_FSR4_RDNA3_UPGRADE=1 %command%",
			CommandEN: "PROTON_FSR4_RDNA3_UPGRADE=1 %command%",
			Title: Localized{
				PT: "Upgrade FSR 4 para RDNA 3",
				EN: "FSR 4 upgrade for RDNA 3",
			},
			Category: Localized{
				PT: "Upscaling",
				EN: "Upscaling",
			},
			Compat: Localized{
				PT: "GE (GPU AMD RDNA 3)",
				EN: "GE (AMD RDNA 3 GPU)",
			},
			Description: Localized{
				PT: "Upgrade FSR 3.1 para FSR 4 em GPUs RDNA 3 (RX 7000). Aplica workarounds específicos RDNA 3. Versão customizável: PROTON_FSR4_RDNA3_UPGRADE=\"4.0.2\". Removida no Proton-CachyOS 11+ (workaround não mais necessário).",
				EN: "FSR 3.1 to FSR 4 upgrade on RDNA 3 GPUs (RX 7000). Applies RDNA 3-specific workarounds. Custom version: PROTON_FSR4_RDNA3_UPGRADE=\"4.0.2\". Removed in Proton-CachyOS 11+ (workaround no longer needed).",
			},
		},
		{
			Command:   "PROTON_FSR3_UPGRADE=1 %command%",
			CommandEN: "PROTON_FSR3_UPGRADE=1 %command%",
			Title: Localized{
				PT: "Upgrade para FSR 3.1",
				EN: "FSR 3.1 upgrade",
			},
			Category: Localized{
				PT: "Upscaling",
				EN: "Upscaling",
			},
			Compat: Localized{
				PT: "GE e CachyOS (GPU AMD)",
				EN: "GE and CachyOS (AMD GPU)",
			},
			Description: Localized{
				PT: "Baixa automaticamente a DLL do FSR 3.1 e atualiza jogos para usá-la. Versão customizável via PROTON_FSR3_UPGRADE=\"versão\". Renomeada para PROTON_FFX3_UPGRADE no Proton-CachyOS 11+.",
				EN: "Automatically downloads the FSR 3.1 DLL and upgrades games to use it. Custom version via PROTON_FSR3_UPGRADE=\"version\". Renamed to PROTON_FFX3_UPGRADE in Proton-CachyOS 11+.",
			},
		},
		{
			Command:   "PROTON_FFX3_UPGRADE=1 %command%",
			CommandEN: "PROTON_FFX3_UPGRADE=1 %command%",
			Title: Localized{
				PT: "Upgrade para FSR 3.1 (nome atual)",
				EN: "FSR 3.1 upgrade (current name)",
			},
			Category: Localized{
				PT: "Upscaling",
				EN: "Upscaling",
			},
			Compat: Localized{
				PT: "CachyOS 11+ (GPU AMD)",
				EN: "CachyOS 11+ (AMD GPU)",
			},
			Description: Localized{
				PT: "Nome atual do upgrade de FSR 3.1 no Proton-CachyOS 11+ (antes PROTON_FSR3_UPGRADE). Baixa automaticamente a DLL do FSR 3.1 e atualiza jogos para usá-la. Versão customizável via PROTON_FFX3_UPGRADE=\"versão\".",
				EN: "Current name of the FSR 3.1 upgrade in Proton-CachyOS 11+ (formerly PROTON_FSR3_UPGRADE). Automatically downloads the FSR 3.1 DLL and upgrades games to use it. Custom version via PROTON_FFX3_UPGRADE=\"version\".",
			},
		},
		{
			Command:   "PROTON_XESS_UPGRADE=1 %command%",
			CommandEN: "PROTON_XESS_UPGRADE=1 %command%",
			Title: Localized{
				PT: "Upgrade do XeSS (Intel)",
				EN: "XeSS upgrade (Intel)",
			},
			Category: Localized{
				PT: "Upscaling",
				EN: "Upscaling",
			},
			Compat: Localized{
				PT: "GE e CachyOS (GPU Intel)",
				EN: "GE and CachyOS (Intel GPU)",
			},
			Description: Localized{
				PT: "Baixa automaticamente a DLL do XeSS (Intel) e atualiza jogos para a versão mais recente do upscaler. Versão customizável via PROTON_XESS_UPGRADE=\"versão\".",
				EN: "Automatically downloads the XeSS (Intel) DLL and upgrades games to the latest upscaler version. Custom version via PROTON_XESS_UPGRADE=\"version\".",
			},
		},
		{
			Command:   "PROTON_DLSS_INDICATOR=1 %command%",
			CommandEN: "PROTON_DLSS_INDICATOR=1 %command%",
			Title: Localized{
				PT: "Indicador DLSS (overlay)",
				EN: "DLSS indicator (overlay)",
			},
			Category: Localized{
				PT: "Upscaling",
				EN: "Upscaling",
			},
			Compat: Localized{
				PT: "GE e CachyOS (GPU NVIDIA)",
				EN: "GE and CachyOS (NVIDIA GPU)",
			},
			Description: Localized{
				PT: "Mostra um overlay DLSS no canto inferior esquerdo da tela. Mesmo efeito de FSR4_WATERMARK=1. Útil para confirmar se o DLSS está ativo.",
				EN: "Shows a DLSS overlay at the bottom left of the screen. Same effect as FSR4_WATERMARK=1. Useful to confirm DLSS is active.",
			},
		},
		{
			Command:   "PROTON_ADD_CONFIG=config1,config2 %command%",
			CommandEN: "PROTON_ADD_CONFIG=config1,config2 %command%",
			Title: Localized{
				PT: "Adicionar compat configs",
				EN: "Add compat configs",
			},
			Category: Localized{
				PT: "Outros",
				EN: "Other",
			},
			Compat: Localized{
				PT: "CachyOS",
				EN: "CachyOS",
			},
			Description: Localized{
				PT: "Passa uma lista de configs de compatibilidade (separadas por vírgula) via variável de ambiente. As configs são a primeira coluna das tabelas de variáveis na documentação do CachyOS.",
				EN: "Pass a comma-separated list of compat configs via environment variable. The configs are the first column in the environment variable tables in CachyOS docs.",
			},
		},
		{
			Command:   "PROTON_LOCAL_SHADER_CACHE=1 %command%",
			CommandEN: "PROTON_LOCAL_SHADER_CACHE=1 %command%",
			Title: Localized{
				PT: "Shader cache local por jogo",
				EN: "Per-game local shader cache",
			},
			Category: Localized{
				PT: "Outros",
				EN: "Other",
			},
			Compat: Localized{
				PT: "CachyOS",
				EN: "CachyOS",
			},
			Description: Localized{
				PT: "Habilita cache de shader por jogo mesmo se o pré-cache de shaders do Steam estiver desligado. Cria cache em <steamlibrary>/shadercache/<appid>. Valores de env do usuário têm prioridade.",
				EN: "Enables per-game shader cache even if Steam's Shader Pre-Caching is off. Creates cache at <steamlibrary>/shadercache/<appid>. User-set env values take priority.",
			},
		},
		{
			Command:   "PROTON_MEDIA_FORCE_GST=1 %command%",
			CommandEN: "PROTON_MEDIA_FORCE_GST=1 %command%",
			Title: Localized{
				PT: "Forçar GStreamer (mídia)",
				EN: "Force GStreamer (media)",
			},
			Category: Localized{
				PT: "Áudio",
				EN: "Audio",
			},
			Compat: Localized{
				PT: "CachyOS",
				EN: "CachyOS",
			},
			Description: Localized{
				PT: "Força o uso do GStreamer para reprodução de vídeo/áudio em cutscenes. Corrige problemas de mídia em alguns jogos.",
				EN: "Forces GStreamer for video/audio playback in cutscenes. Fixes media issues in some games.",
			},
		},
		{
			Command:   "PROTON_GST_VIDEO_ORIENTATION=90 %command%",
			CommandEN: "PROTON_GST_VIDEO_ORIENTATION=90 %command%",
			Title: Localized{
				PT: "Rotar vídeo (GStreamer)",
				EN: "Rotate video (GStreamer)",
			},
			Category: Localized{
				PT: "Áudio",
				EN: "Audio",
			},
			Compat: Localized{
				PT: "CachyOS",
				EN: "CachyOS",
			},
			Description: Localized{
				PT: "Muda a orientação/rotação de vídeos renderizados via GStreamer. Valores aceitos: os mesmos do plugin videoflip do GStreamer (ex.: 90, 180, 270, identity).",
				EN: "Changes the orientation/rotation of videos rendered via GStreamer. Accepted values: same as GStreamer's videoflip plugin (e.g.: 90, 180, 270, identity).",
			},
		},
		{
			Command:   "PROTON_OPTISCALER_NAME=dxgi.dll %command%",
			CommandEN: "PROTON_OPTISCALER_NAME=dxgi.dll %command%",
			Title: Localized{
				PT: "Escolher DLL do OptiScaler",
				EN: "Choose OptiScaler DLL",
			},
			Category: Localized{
				PT: "Upscaling",
				EN: "Upscaling",
			},
			Compat: Localized{
				PT: "CachyOS (com PROTON_USE_OPTISCALER=1)",
				EN: "CachyOS (with PROTON_USE_OPTISCALER=1)",
			},
			Description: Localized{
				PT: "Define qual DLL o OptiScaler deve injetar. Opções: dxgi.dll, d3d11.dll, d3d12.dll. Combine com PROTON_OPTISCALER_CONFIG para configuração completa.",
				EN: "Sets which DLL OptiScaler should inject. Options: dxgi.dll, d3d11.dll, d3d12.dll. Combine with PROTON_OPTISCALER_CONFIG for full configuration.",
			},
		},
		{
			Command:   "PROTON_OPTISCALER_CONFIG=\"Upscalers.Dx11Upscaler=fsr31;Upscalers.Dx12Upscaler=dlss\" %command%",
			CommandEN: "PROTON_OPTISCALER_CONFIG=\"Upscalers.Dx11Upscaler=fsr31;Upscalers.Dx12Upscaler=dlss\" %command%",
			Title: Localized{
				PT: "Configuração do OptiScaler",
				EN: "OptiScaler configuration",
			},
			Category: Localized{
				PT: "Upscaling",
				EN: "Upscaling",
			},
			Compat: Localized{
				PT: "CachyOS (com PROTON_USE_OPTISCALER=1)",
				EN: "CachyOS (with PROTON_USE_OPTISCALER=1)",
			},
			Description: Localized{
				PT: "Escreve configuração arbitrária no OptiScaler.ini via variável de ambiente (separada por ponto-e-vírgula). Exemplo: Upscalers.Dx11Upscaler=fsr31;Upscalers.Dx12Upscaler=dlss. Evita erros, use apenas chaves válidas do OptiScaler.ini.",
				EN: "Writes arbitrary OptiScaler.ini config via environment variable (semicolon-separated). Example: Upscalers.Dx11Upscaler=fsr31;Upscalers.Dx12Upscaler=dlss. Avoid mistakes, only use valid OptiScaler.ini keys.",
			},
		},
		{
			Command:   "DXVK_NVAPI_VKREFLEX=1 %command%",
			CommandEN: "DXVK_NVAPI_VKREFLEX=1 %command%",
			Title: Localized{
				PT: "NVIDIA Reflex em Vulkan (VKREFLEX)",
				EN: "NVIDIA Reflex in Vulkan (VKREFLEX)",
			},
			Category: Localized{
				PT: "Latência",
				EN: "Latency",
			},
			Compat: Localized{
				PT: "GE e CachyOS (GPU NVIDIA; jogos Vulkan com Reflex)",
				EN: "GE and CachyOS (NVIDIA GPU; Vulkan games with Reflex)",
			},
			Description: Localized{
				PT: "Habilita a layer Vulkan Reflex do dxvk-nvapi para jogos como Portal RTX, Path of Exile 1/2, Doom TDA. Requer GPU NVIDIA e jogo com suporte nativo a Reflex em Vulkan.",
				EN: "Enables dxvk-nvapi's Vulkan Reflex layer for games like Portal RTX, Path of Exile 1/2, Doom TDA. Requires NVIDIA GPU and game with native Vulkan Reflex support.",
			},
		},
		{
			Command:   "PROTON_USE_WAYLAND=1 %command%",
			CommandEN: "PROTON_USE_WAYLAND=1 %command%",
			Title: Localized{
				PT: "Driver Wayland (alias)",
				EN: "Wayland driver (alias)",
			},
			Category: Localized{
				PT: "Wayland",
				EN: "Wayland",
			},
			Compat: Localized{
				PT: "CachyOS (experimental)",
				EN: "CachyOS (experimental)",
			},
			Description: Localized{
				PT: "Alias de PROTON_ENABLE_WAYLAND. Ativa o driver winewayland (janela nativa Wayland, sem XWayland). Mesma função, nome diferente — use o que preferir. Veja PROTON_ENABLE_WAYLAND para os detalhes e limitações.",
				EN: "Alias for PROTON_ENABLE_WAYLAND. Enables the winewayland driver (native Wayland window, no XWayland). Same function, different name — pick whichever you like. See PROTON_ENABLE_WAYLAND for details and limitations.",
			},
		},
		{
			Command:   "PROTON_PREFER_SDL=1 %command%",
			CommandEN: "PROTON_PREFER_SDL=1 %command%",
			Title: Localized{
				PT: "Input via SDL (alias)",
				EN: "SDL input (alias)",
			},
			Category: Localized{
				PT: "Input",
				EN: "Input",
			},
			Compat: Localized{
				PT: "CachyOS",
				EN: "CachyOS",
			},
			Description: Localized{
				PT: "Alias de PROTON_USE_SDL. Usa input SDL em vez de HIDRAW/Steam Input. Mesma função, nome diferente. Útil quando o controle não é detectado ou se comporta mal — comum com o driver Wayland ativo.",
				EN: "Alias for PROTON_USE_SDL. Uses SDL input instead of HIDRAW/Steam Input. Same function, different name. Useful when the controller isn't detected or misbehaves — common with the Wayland driver active.",
			},
		},
		{
			Command:   "PROTON_NO_STEAMINPUT=1 %command%",
			CommandEN: "PROTON_NO_STEAMINPUT=1 %command%",
			Title: Localized{
				PT: "Desabilitar Steam Input",
				EN: "Disable Steam Input",
			},
			Category: Localized{
				PT: "Input",
				EN: "Input",
			},
			Compat: Localized{
				PT: "CachyOS",
				EN: "CachyOS",
			},
			Description: Localized{
				PT: "Desabilita completamente o Steam Input no Proton. Útil quando o Steam Input conflita com o controle nativo do jogo.",
				EN: "Completely disables Steam Input in Proton. Useful when Steam Input conflicts with the game's native controller support.",
			},
		},
		{
			Command:   "DXVK_FILTER_DEVICE_NAME=\"NVIDIA\" %command%",
			CommandEN: "DXVK_FILTER_DEVICE_NAME=\"NVIDIA\" %command%",
			Title: Localized{
				PT: "Filtrar GPU por nome",
				EN: "Filter GPU by name",
			},
			Category: Localized{
				PT: "GPU",
				EN: "GPU",
			},
			Compat: Localized{
				PT: "DXVK (todos)",
				EN: "DXVK (all)",
			},
			Description: Localized{
				PT: "Força o DXVK a usar uma GPU específica pelo nome. Útil em notebooks com GPU integrada + dedicada — substitua \"NVIDIA\" pelo nome exato da sua GPU (ex: \"NVIDIA GeForce RTX 3070\"). Use DXVK_HUD=devinfo para ver o nome da GPU.",
				EN: "Forces DXVK to use a specific GPU by name. Useful on laptops with integrated + dedicated GPU — replace \"NVIDIA\" with your exact GPU name (e.g. \"NVIDIA GeForce RTX 3070\"). Use DXVK_HUD=devinfo to see GPU name.",
			},
		},
		{
			Command:   "PROTON_FRAME_RATE=60 %command%",
			CommandEN: "PROTON_FRAME_RATE=60 %command%",
			Title: Localized{
				PT: "Limitar FPS via Proton",
				EN: "Cap FPS via Proton",
			},
			Category: Localized{
				PT: "Desempenho",
				EN: "Performance",
			},
			Compat: Localized{
				PT: "Proton 8+",
				EN: "Proton 8+",
			},
			Description: Localized{
				PT: "Limita o framerate no nível do Proton, antes do jogo renderizar. Mais leve que limitadores de overlay. Substitua 60 pelo valor desejado. Funciona melhor que V-Sync em monitores com refresh rate alto.",
				EN: "Caps framerate at the Proton level, before the game renders. Lighter than overlay cappers. Replace 60 with your desired value. Works better than V-Sync on high refresh rate monitors.",
			},
		},
		{
			Command:   "DXVK_FILTER_DEVICE_NAME=\"AMD\" %command%",
			CommandEN: "DXVK_FILTER_DEVICE_NAME=\"AMD\" %command%",
			Title: Localized{
				PT: "Forçar GPU AMD",
				EN: "Force AMD GPU",
			},
			Category: Localized{
				PT: "GPU",
				EN: "GPU",
			},
			Compat: Localized{
				PT: "DXVK (todos)",
				EN: "DXVK (all)",
			},
			Description: Localized{
				PT: "Força o DXVK a usar a GPU AMD. Em sistemas com AMD + Intel, garante que a GPU dedicada seja usada. Substitua \"AMD\" pelo nome exato da sua GPU.",
				EN: "Forces DXVK to use the AMD GPU. On systems with AMD + Intel, ensures the dedicated GPU is used. Replace \"AMD\" with your exact GPU name.",
			},
		},
		{
			Command:   "MANGOHUD_CONFIG=\"fps_limit=60\" %command%",
			CommandEN: "MANGOHUD_CONFIG=\"fps_limit=60\" %command%",
			Title: Localized{
				PT: "MangoHUD: limitar FPS",
				EN: "MangoHUD: cap FPS",
			},
			Category: Localized{
				PT: "Desempenho",
				EN: "Performance",
			},
			Compat: Localized{
				PT: "MangoHUD instalado",
				EN: "MangoHUD installed",
			},
			Description: Localized{
				PT: "Configura o MangoHUD via variável de ambiente. Substitua 60 pelo FPS desejado. Outras opções: \"fps_limit=60,120\" (lista), \"toggle_fps_limit=Shift_L+F1\" (atalho). O MangoHUD precisa estar instalado no sistema.",
				EN: "Configures MangoHUD via environment variable. Replace 60 with desired FPS. Other options: \"fps_limit=60,120\" (list), \"toggle_fps_limit=Shift_L+F1\" (hotkey). MangoHUD must be installed on the system.",
			},
		},
	}
}
