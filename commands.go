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
				PT: "Baixa automaticamente a amdxcffx64.dll e atualiza jogos com FSR 3.1 para FSR 4. Versão customizável: PROTON_FSR4_UPGRADE=\"4.0.2\" (default 4.0.2 no GE, 4.1.1 no CachyOS). Para RDNA3 use PROTON_FSR4_RDNA3_UPGRADE. Desativa Anti-Lag 2.",
				EN: "Automatically downloads amdxcffx64.dll and upgrades games with FSR 3.1 to FSR 4. Custom version: PROTON_FSR4_UPGRADE=\"4.0.2\" (default 4.0.2 on GE, 4.1.1 on CachyOS). For RDNA3 use PROTON_FSR4_RDNA3_UPGRADE. Disables Anti-Lag 2.",
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
				PT: "CachyOS",
				EN: "CachyOS",
			},
			Description: Localized{
				PT: "Usa o fork dxvk-low-latency nos jogos Direct3D 8/9/10/11: adiciona frame pacing de baixa latência, melhorando a responsividade (input lag) e a estabilidade da latência. Para jogos Direct3D 12 use PROTON_VKD3D_LOWLATENCY.",
				EN: "Uses the dxvk-low-latency fork in Direct3D 8/9/10/11 games: adds low-latency frame pacing, improving responsiveness (input lag) and latency stability. For Direct3D 12 games use PROTON_VKD3D_LOWLATENCY.",
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
				PT: "CachyOS",
				EN: "CachyOS",
			},
			Description: Localized{
				PT: "Usa o fork vkd3d-low-latency nos jogos Direct3D 12: o equivalente do dxvk-low-latency para D3D12, com frame pacing de menor latência. Para jogos Direct3D 8/9/10/11 use PROTON_DXVK_LOWLATENCY. Combine com a layer LOW_LATENCY_LAYER para Anti-Lag/Reflex.",
				EN: "Uses the vkd3d-low-latency fork in Direct3D 12 games: the D3D12 equivalent of dxvk-low-latency, with lower-latency frame pacing. For Direct3D 8/9/10/11 games use PROTON_DXVK_LOWLATENCY. Combine with the LOW_LATENCY_LAYER layer for Anti-Lag/Reflex.",
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
				PT: "Ativa o low_latency_layer, que expõe a extensão VK_AMD_anti_lag em GPUs AMD e Intel — o Anti-Lag 2 passa a funcionar em jogos Vulkan (CS2 nativo, e via dxvk-nvapi em jogos Proton com proton-cachyos/GE, que já embutem o layer). Use para reduzir o input lag em jogos competitivos. Desative com DISABLE_LOW_LATENCY_LAYER=1 se causar travamentos.",
				EN: "Enables low_latency_layer, which exposes the VK_AMD_anti_lag extension on AMD and Intel GPUs — Anti-Lag 2 now works in Vulkan games (native CS2, and via dxvk-nvapi in Proton games with proton-cachyos/GE, which already bundle the layer). Use to reduce input lag in competitive games. Disable with DISABLE_LOW_LATENCY_LAYER=1 if it causes crashes.",
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
				PT: "Faz o low_latency_layer expor VK_NV_low_latency2 em vez de anti-lag, ativando o Reflex (mesmo desempenho do Anti-Lag 2, mas com suporte em muito mais jogos). Se o menu Reflex não aparecer, tente também PROTON_FORCE_NVAPI=1 (atenção: quebra o upgrade FSR4).",
				EN: "Makes low_latency_layer expose VK_NV_low_latency2 instead of anti-lag, enabling Reflex (same performance as Anti-Lag 2, but supported in far more games). If the Reflex menu doesn't appear, also try PROTON_FORCE_NVAPI=1 (warning: breaks FSR4 upgrade).",
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
			Command:   "LSFG_PROCESS=steam %command%",
			CommandEN: "LSFG_PROCESS=steam %command%",
			Title: Localized{
				PT: "lsfg-vk: frame generation",
				EN: "lsfg-vk: frame generation",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "Todos (exige lsfg-vk instalado)",
				EN: "All (requires lsfg-vk)",
			},
			Description: Localized{
				PT: "Ativa a layer de frame generation do lsfg-vk usando o profile \"steam\". Crie o profile e ajuste multiplicador, modo de desempenho, HDR etc. na interface gráfica do lsfg-vk (comando lsfg-vk-ui). Requer lsfg-vk instalado e Lossless Scaling na Steam.",
				EN: "Enables the lsfg-vk frame generation layer using the \"steam\" profile. Create the profile and tune multiplier, performance mode, HDR etc. in the lsfg-vk GUI (lsfg-vk-ui command). Requires lsfg-vk installed and Lossless Scaling on Steam.",
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
				PT: "GE (GPU NVIDIA)",
				EN: "GE (NVIDIA GPU)",
			},
			Description: Localized{
				PT: "Habilita o suporte a patentes do NVIDIA Reflex no dxvk-nvapi. Útil quando o menu de Reflex não aparece em jogos com suporte. Atenção: quebra o upgrade FSR 4 (PROTON_FSR4_UPGRADE).",
				EN: "Enables NVIDIA Reflex patent support in dxvk-nvapi. Useful when the Reflex menu doesn't show in supported games. Warning: breaks the FSR 4 upgrade (PROTON_FSR4_UPGRADE).",
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
				PT: "Limita o FPS dos jogos D3D11/D3D10/D3D9 (via DXVK) direto no driver, sem overlay. Troque 60 pelo limite desejado (ex.: 120, 144); -1 desativa. Útil em telas 60 Hz ou para reduzir consumo/ruído quando o jogo passa folgado do refresh.",
				EN: "Caps FPS in D3D11/D3D10/D3D9 games (via DXVK) right in the driver, with no overlay. Replace 60 with the desired cap (e.g.: 120, 144); -1 disables. Useful on 60 Hz displays or to cut power/noise when the game runs well past refresh.",
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
	}
}
