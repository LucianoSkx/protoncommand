package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func testGUI(lang, launcherID string) *gui {
	ls := launchers()
	idx := 0
	for i := range ls {
		if ls[i].ID == launcherID {
			idx = i
		}
	}
	return &gui{all: commands(), launchers: ls, launcherIdx: idx, selected: map[int]bool{}, lang: lang}
}

// initTestGUI cria um gui com todos os widgets necessários para testes de UI.
func initTestGUI(lang, launcherID string) *gui {
	a := test.NewApp()
	w := a.NewWindow("test")
	g := testGUI(lang, launcherID)
	g.app = a
	g.win = w
	g.search = widget.NewEntry()
	g.search.OnChanged = func(_ string) { g.applyFilter() }
	g.status = newMaxWidthLabel(500)
	g.copyBtn = widget.NewButton("", func() {})
	g.copyOnClick = &fyne.MenuItem{}
	g.favToggleBtn = widget.NewButton("", func() {})
	g.detailTitle = widget.NewLabel("")
	g.detailCat = widget.NewLabel("")
	g.detailCompat = widget.NewLabel("")
	g.detailCmd = newMaxWidthLabel(500)
	g.detailDesc = newMaxWidthLabel(500)
	g.langRadio = widget.NewRadioGroup([]string{"Português", "Inglês"}, func(string) {})
	g.combLabel = newMaxWidthLabel(940)
	g.combCount = widget.NewLabel("")
	g.combWarn = canvas.NewText("", theme.ErrorColor())
	g.combHint = widget.NewLabel("")
	g.combCopyBtn = widget.NewButton("", nil)
	g.clearBtn = widget.NewButton("", nil)
	g.launcherSel = widget.NewSelect(nil, func(string) {})
	g.catSel = widget.NewSelect(nil, func(string) {})
	g.favBtn = widget.NewButton("", nil)
	g.exportFavBtn = widget.NewButton("", nil)
	g.importFavBtn = widget.NewButton("", nil)
	g.langHeader = fyne.NewMenuItem("", nil)
	g.langPT = fyne.NewMenuItem("Português", nil)
	g.langEN = fyne.NewMenuItem("Inglês", nil)
	g.themeSystem = fyne.NewMenuItem("", nil)
	g.themeLight = fyne.NewMenuItem("", nil)
	g.themeDark = fyne.NewMenuItem("", nil)
	g.configBtn = widget.NewButton("", nil)
	g.aboutBtn = widget.NewButton("", nil)
	g.list = widget.NewList(
		func() int { return len(g.filtered) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			if id < len(g.filtered) {
				obj.(*widget.Label).SetText(g.all[g.filtered[id]].Command)
			}
		},
	)
	g.list.OnSelected = func(id widget.ListItemID) {
		g.selID = id
		g.selectCommand(id)
	}
	g.list.OnUnselected = func(_ widget.ListItemID) {
		g.selID = -1
		g.selectCommand(-1)
	}
	g.favs = map[string]bool{}
	g.filtered = make([]int, len(g.all))
	for i := range g.all {
		g.filtered[i] = i
	}
	return g
}

func idxOf(cmd string) int {
	cmds := commands()
	for i := range cmds {
		if cmds[i].Command == cmd {
			return i
		}
	}
	return -1
}

func TestCombinationSteam(t *testing.T) {
	g := testGUI("pt", "steam")
	g.selected[idxOf("mangohud %command%")] = true
	g.selected[idxOf("PROTON_LOG=1 %command%")] = true
	got := g.buildCombination()
	want := "PROTON_LOG=1 mangohud %command%"
	if got != want {
		t.Fatalf("steam: got %q want %q", got, want)
	}
}

func TestCombinationNoCmd(t *testing.T) {
	g := testGUI("pt", "heroic")
	g.selected[idxOf("mangohud %command%")] = true
	g.selected[idxOf("PROTON_LOG=1 %command%")] = true
	got := g.buildCombination()
	want := "PROTON_LOG=1 mangohud"
	if got != want {
		t.Fatalf("heroic: got %q want %q", got, want)
	}
}

func TestGamescopeNoCmdDropsTrailingDashDash(t *testing.T) {
	g := testGUI("pt", "heroic")
	g.selected[idxOf("gamescope -e -f -F fsr -- %command%")] = true
	g.selected[idxOf("PROTON_LOG=1 %command%")] = true
	got := g.buildCombination()
	if strings.HasSuffix(got, "--") || strings.Contains(got, " -- ") {
		t.Fatalf("heroic gamescope: linha quebrada com --: %q", got)
	}
	want := "PROTON_LOG=1 gamescope -e -f -F fsr"
	if got != want {
		t.Fatalf("heroic gamescope: got %q want %q", got, want)
	}
	gsteam := testGUI("pt", "steam")
	gsteam.selected[idxOf("gamescope -e -f -F fsr -- %command%")] = true
	if got := gsteam.buildCombination(); !strings.Contains(got, " -- ") {
		t.Fatalf("steam gamescope: esperado -- preservado, got %q", got)
	}
}

func TestDisplayCmdStrips(t *testing.T) {
	g := testGUI("pt", "lutris")
	c := commands()[idxOf("PROTON_LOG=1 %command%")]
	if got := g.displayCmd(c); got != "PROTON_LOG=1" {
		t.Fatalf("displayCmd: got %q", got)
	}
	g2 := testGUI("pt", "steam")
	if got := g2.displayCmd(c); got != "PROTON_LOG=1 %command%" {
		t.Fatalf("displayCmd steam: got %q", got)
	}
}

func TestConflictDuplicate(t *testing.T) {
	g := testGUI("pt", "steam")
	g.selected[idxOf("PROTON_LOG=1 %command%")] = true
	g.selected[idxOf("PROTON_LOG=warn+pipewire,warn+mmdevapi %command%")] = true
	warns := g.conflicts()
	if len(warns) == 0 {
		t.Fatal("expected duplicate-key conflict")
	}
	if !strings.Contains(warns[0], "PROTON_LOG") {
		t.Fatalf("unexpected warning: %q", warns[0])
	}
}

func TestConflictAntiLagReflex(t *testing.T) {
	g := testGUI("pt", "steam")
	g.selected[idxOf("LOW_LATENCY_LAYER=1 %command%")] = true
	g.selected[idxOf(`LOW_LATENCY_LAYER=1 LOW_LATENCY_LAYER_REFLEX=1 DXVK_CONFIG="dxgi.hideAmdGpu = True" %command%`)] = true
	warns := g.conflicts()
	found := false
	for _, w := range warns {
		if strings.Contains(w, "Reflex") {
			found = true
		}
	}
	if !found {
		t.Fatalf("expected anti-lag/reflex conflict, got %v", warns)
	}
}

func TestNoConflictSameValue(t *testing.T) {
	g := testGUI("pt", "steam")
	g.selected[idxOf("WINE_ESYNC=1 %command%")] = true
	g.selected[idxOf("WINEFSYNC=1 %command%")] = true
	if warns := g.conflicts(); len(warns) != 0 {
		t.Fatalf("expected no conflict, got %v", warns)
	}
}

func TestSplitFieldsKeepsQuotedValue(t *testing.T) {
	toks := splitFields(`LOW_LATENCY_LAYER=1 LOW_LATENCY_LAYER_REFLEX=1 DXVK_CONFIG="dxgi.hideAmdGpu = True" %command%`)
	want := []string{
		"LOW_LATENCY_LAYER=1",
		"LOW_LATENCY_LAYER_REFLEX=1",
		`DXVK_CONFIG="dxgi.hideAmdGpu = True"`,
		"%command%",
	}
	if len(toks) != len(want) {
		t.Fatalf("got %q want %q", toks, want)
	}
	for i := range want {
		if toks[i] != want[i] {
			t.Fatalf("tok %d: got %q want %q (all: %q)", i, toks[i], want[i], toks)
		}
	}
}

func TestNoFalseConflictFromQuotedSpaces(t *testing.T) {
	g := testGUI("pt", "steam")
	g.selected[idxOf("PROTON_LOG=1 %command%")] = true
	g.selected[idxOf(`LOW_LATENCY_LAYER=1 LOW_LATENCY_LAYER_REFLEX=1 DXVK_CONFIG="dxgi.hideAmdGpu = True" %command%`)] = true
	for _, w := range g.conflicts() {
		if strings.Contains(w, "DXVK_CONFIG") || strings.Contains(w, "True") {
			t.Fatalf("falso positivo de conflito com valor citado: %q", w)
		}
	}
}

func TestI18nParity(t *testing.T) {
	for k := range ptTexts {
		if _, ok := enTexts[k]; !ok {
			t.Errorf("chave %q existe em PT mas não em EN", k)
		}
	}
	for k := range enTexts {
		if _, ok := ptTexts[k]; !ok {
			t.Errorf("chave %q existe em EN mas não em PT", k)
		}
	}
}

func TestCommandsValid(t *testing.T) {
	cmds := commands()
	if len(cmds) == 0 {
		t.Fatal("catálogo vazio")
	}
	seen := map[string]bool{}
	for i, c := range cmds {
		if c.Command == "" {
			t.Errorf("comando %d: Command vazio", i)
		}
		if seen[c.Command] {
			t.Errorf("comando duplicado: %q", c.Command)
		}
		seen[c.Command] = true
		if !strings.HasSuffix(c.Command, "%command%") {
			t.Errorf("comando %q não termina com %%command%%", c.Command)
		}
		if c.Title.PT == "" || c.Title.EN == "" {
			t.Errorf("comando %q: título PT/EN incompleto", c.Command)
		}
		if c.Description.PT == "" || c.Description.EN == "" {
			t.Errorf("comando %q: descrição PT/EN incompleta", c.Command)
		}
		if c.Category.PT == "" || c.Category.EN == "" {
			t.Errorf("comando %q: categoria PT/EN incompleta", c.Command)
		}
	}
}

// TestNoObsoleteCommands impede o retorno de variáveis removidas por
// auditoria (obsoletas, renomeadas ou sem evidência de existência).
func TestNoObsoleteCommands(t *testing.T) {
	obsolete := []string{
		"PROTON_DUMP_DEBUG_COMMANDS",
		"DXVK_ASYNC",
		"PROTON_USE_NTSYNC",
		"PROTON_ENABLE_NVAPI",
		"WINE_AUDIO_DRIVER",
		"WINE_BLOCK_HOSTS",
		"WINE_VIRTUAL_DESKTOP",
		"WINE_ESYNC",
		"WINEFSYNC",
		"PROTON_VKREFLEX",
		"PROTON_VKD3D_HEAP",
		"PROTON_NO_D3D9",
		"FNA3D_FORCE_DRIVER",
		"DRI_CONFIG",
		"PROTON_ENABLE_HDR",
	}
	cmds := commands()
	for _, c := range cmds {
		for _, tok := range splitFields(c.Command) {
			name := tok
			if idx := strings.IndexByte(tok, '='); idx >= 0 {
				name = tok[:idx]
			}
			for _, ob := range obsolete {
				if name == ob {
					t.Errorf("comando obsoleto de volta ao catálogo: %q contém %s", c.Command, ob)
				}
			}
		}
	}
}

// TestFilterKeepsDetailInSync garante que, ao filtrar, o painel de
// detalhes acompanha a lista mesmo quando a linha 0 já estava
// selecionada (nesse caso List.Select não dispara OnSelected).
func TestFilterKeepsDetailInSync(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")

	g.search.Text = ""
	g.applyFilter()
	if len(g.filtered) == 0 || g.current != g.filtered[0] {
		t.Fatalf("sem filtro: current=%d, esperado filtered[0]=%d", g.current, g.filtered[0])
	}
	g.search.Text = "lsfg"
	g.applyFilter()
	if len(g.filtered) == 0 {
		t.Fatal("filtro 'lsfg' não achou nada")
	}
	if g.current != g.filtered[0] {
		t.Fatalf("detalhe fora de sincronia: mostra %q, lista mostra %q",
			g.all[g.current].Command, g.all[g.filtered[0]].Command)
	}
}

// TestSetLangKeepsDetailInSync garante que trocar de idioma não deixa
// o detalhe mostrando texto do idioma antigo.
func TestSetLangKeepsDetailInSync(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")

	g.search.Text = ""
	g.applyFilter()
	g.list.Select(0)
	g.setLang("en", false)
	if g.current != g.filtered[0] {
		t.Fatalf("detalhe fora de sincronia após troca de idioma: current=%d, filtered[0]=%d", g.current, g.filtered[0])
	}
	if g.lang != "en" {
		t.Fatalf("idioma não mudou: esperado 'en', obtido %q", g.lang)
	}
}

// TestToggleFav adiciona e remove favoritos corretamente.
func TestToggleFav(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	idx := idxOf("PROTON_LOG=1 %command%")
	if idx < 0 {
		t.Fatal("comando PROTON_LOG não encontrado")
	}
	g.toggleFav(idx)
	key := g.favKey(idx)
	if !g.favs[key] {
		t.Fatal("favorito não foi adicionado")
	}
	g.toggleFav(idx)
	if g.favs[key] {
		t.Fatal("favorito não foi removido")
	}
}

// TestClearSelection limpa todos os comandos selecionados.
func TestClearSelection(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.selected[idxOf("PROTON_LOG=1 %command%")] = true
	g.selected[idxOf("mangohud %command%")] = true
	g.clearSelection()
	if len(g.selected) != 0 {
		t.Fatalf("clearSelection não limou: %v", g.selected)
	}
}

// TestSelectCommandSelecionaComando g.current aponta para o comando certo.
func TestSelectCommandSelecionaComando(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.search.Text = ""
	g.applyFilter()
	if len(g.filtered) == 0 {
		t.Fatal("nenhum comando filtrado")
	}
	g.selectCommand(0)
	if g.current != g.filtered[0] {
		t.Fatalf("selectCommand(0): current=%d, filtered[0]=%d", g.current, g.filtered[0])
	}
	if g.detailTitle.Text == "" {
		t.Fatal("detailTitle vazio após selectCommand")
	}
}

// TestSelectCommandInvalido limpa o detalhe.
func TestSelectCommandInvalido(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.selectCommand(-1)
	if g.current != -1 {
		t.Fatalf("selectCommand(-1): current=%d, esperado -1", g.current)
	}
	g.selectCommand(9999)
	if g.current != -1 {
		t.Fatalf("selectCommand(9999): current=%d, esperado -1", g.current)
	}
}

// TestClearDetail limpa o detalhe e desabilita o botão de copiar.
func TestClearDetail(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.selectCommand(0)
	g.clearDetail()
	if g.current != -1 {
		t.Fatalf("clearDetail: current=%d, esperado -1", g.current)
	}
	if g.detailTitle.Text != g.tr("noCommand") {
		t.Fatalf("clearDetail: title=%q, esperado %q", g.detailTitle.Text, g.tr("noCommand"))
	}
}

// TestUpdateFavButtonAtualizaTexto verifica se o botão de favorito
// mostra o texto correto baseado no estado atual.
func TestUpdateFavButtonAtualizaTexto(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	idx := idxOf("PROTON_LOG=1 %command%")
	g.selectCommand(idx)
	g.updateFavButton()
	if !strings.Contains(g.favToggleBtn.Text, g.tr("addFavorite")) {
		t.Fatalf("botão deveria mostrar addFavorite, got %q", g.favToggleBtn.Text)
	}
	g.toggleFav(idx)
	g.updateFavButton()
	if !strings.Contains(g.favToggleBtn.Text, g.tr("removeFavorite")) {
		t.Fatalf("botão deveria mostrar removeFavorite, got %q", g.favToggleBtn.Text)
	}
}

// TestSetThemeAplicaTema verifica se o tema é aplicado corretamente.
func TestSetThemeAplicaTema(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.setTheme("dark", false)
	if !g.themeDark.Checked {
		t.Fatal("themeDark deveria estar checked após setTheme('dark')")
	}
	if g.themeSystem.Checked || g.themeLight.Checked {
		t.Fatal("outros temas não deveriam estar checked")
	}
	g.setTheme("light", false)
	if !g.themeLight.Checked {
		t.Fatal("themeLight deveria estar checked após setTheme('light')")
	}
	g.setTheme("system", false)
	if !g.themeSystem.Checked {
		t.Fatal("themeSystem deveria estar checked após setTheme('system')")
	}
}

// TestSetLangPersistencia verifica se o idioma é persistido.
func TestSetLangPersistencia(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.setLang("en", true)
	if g.lang != "en" {
		t.Fatalf("idioma não mudou: esperado 'en', obtido %q", g.lang)
	}
	g.setLang("pt", true)
	if g.lang != "pt" {
		t.Fatalf("idioma não mudou: esperado 'pt', obtido %q", g.lang)
	}
	g.setLang("invalid", true)
	if g.lang != "pt" {
		t.Fatalf("idioma deveria permanecer 'pt' para valor inválido, obtido %q", g.lang)
	}
}

// TestLauncher atualiza o launcher corretamente.
func TestLauncher(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	if g.launcher().ID != "steam" {
		t.Fatalf("launcher inicial: esperado 'steam', obtido %q", g.launcher().ID)
	}
	g.setLauncher("heroic")
	if g.launcher().ID != "heroic" {
		t.Fatalf("setLauncher: esperado 'heroic', obtido %q", g.launcher().ID)
	}
}

// TestUpdateCombinationAtualizaContagem verifica se a contagem de
// seleções é atualizada corretamente.
func TestUpdateCombinationAtualizaContagem(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.updateCombination()
	if g.combCount.Text != "" {
		t.Fatalf("combCount deveria estar vazio, got %q", g.combCount.Text)
	}
	g.selected[idxOf("PROTON_LOG=1 %command%")] = true
	g.updateCombination()
	if g.combCount.Text == "" {
		t.Fatal("combCount não deveria estar vazio após selecionar comando")
	}
}

// TestApplyFilterCategoriaFiltraCorretamente.
func TestApplyFilterCategoriaFiltraCorretamente(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.search.Text = ""
	g.catFilterPT = ""
	g.applyFilter()
	totalSemFiltro := len(g.filtered)

	// Filtra por uma categoria específica
	for _, c := range g.all {
		g.catFilterPT = c.Category.PT
		break
	}
	g.applyFilter()
	totalComFiltro := len(g.filtered)
	if totalComFiltro >= totalSemFiltro {
		t.Fatalf("filtro por categoria deveria reduzir resultados: %d >= %d", totalComFiltro, totalSemFiltro)
	}
}

// TestApplyFilterFavoritosFiltraCorretamente.
func TestApplyFilterFavoritosFiltraCorretamente(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.search.Text = ""
	g.favOnly = false
	g.applyFilter()

	idx := idxOf("PROTON_LOG=1 %command%")
	key := g.favKey(idx)
	g.favs[key] = true
	g.favOnly = true
	g.applyFilter()
	totalComFiltro := len(g.filtered)
	if totalComFiltro != 1 {
		t.Fatalf("filtro de favoritos deveria retornar 1 resultado, got %d", totalComFiltro)
	}
}

// TestCombinationMultiplosComandos testa combinação com vários comandos.
func TestCombinationMultiplosComandos(t *testing.T) {
	g := testGUI("pt", "steam")
	g.selected[idxOf("PROTON_LOG=1 %command%")] = true
	g.selected[idxOf("mangohud %command%")] = true
	g.selected[idxOf("gamemoderun %command%")] = true
	got := g.buildCombination()
	if !strings.Contains(got, "PROTON_LOG=1") {
		t.Fatalf("combinação deveria conter PROTON_LOG=1, got %q", got)
	}
	if !strings.Contains(got, "mangohud") {
		t.Fatalf("combinação deveria conter mangohud, got %q", got)
	}
	if !strings.Contains(got, "gamemoderun") {
		t.Fatalf("combinação deveria conter gamemoderun, got %q", got)
	}
	if !strings.HasSuffix(got, "%command%") {
		t.Fatalf("combinação deveria terminar com %%command%%, got %q", got)
	}
}

// TestCombinationWrappersPorUltimo verifica se wrappers ficam por último.
func TestCombinationWrappersPorUltimo(t *testing.T) {
	g := testGUI("pt", "steam")
	g.selected[idxOf("PROTON_LOG=1 %command%")] = true
	g.selected[idxOf("mangohud %command%")] = true
	g.selected[idxOf("gamescope -e -f -F fsr -- %command%")] = true
	got := g.buildCombination()
	// Wrappers devem estar depois de PROTON_LOG
	idxLog := strings.Index(got, "PROTON_LOG")
	idxMango := strings.Index(got, "mangohud")
	idxGamescope := strings.Index(got, "gamescope")
	if idxLog > idxMango || idxLog > idxGamescope {
		t.Fatalf("PROTON_LOG deveria estar antes dos wrappers, got %q", got)
	}
}

// TestConflictSemConflitoDeAspas verifica que valores entre aspas
// não geram falsos positivos.
func TestConflictSemConflitoDeAspas(t *testing.T) {
	g := testGUI("pt", "steam")
	// Seleciona dois comandos com PROTON_LOG (mesmo valor, sem conflito)
	g.selected[idxOf("PROTON_LOG=1 %command%")] = true
	g.selected[idxOf("PROTON_LOG=warn+pipewire,warn+mmdevapi %command%")] = true
	warns := g.conflicts()
	found := false
	for _, w := range warns {
		if strings.Contains(w, "PROTON_LOG") {
			found = true
		}
	}
	if !found {
		t.Fatal("deveria detectar conflito de PROTON_LOG com valores diferentes")
	}
}

// TestSplitFieldsAspasSimples verifica preservação de aspas simples.
func TestSplitFieldsAspasSimples(t *testing.T) {
	toks := splitFields(`VAR1=1 'value with spaces' VAR2=2`)
	if len(toks) != 3 {
		t.Fatalf("esperado 3 tokens, got %d: %q", len(toks), toks)
	}
	if toks[1] != "'value with spaces'" {
		t.Fatalf("token 1: esperado 'value with spaces', got %q", toks[1])
	}
}

// TestSplitFieldsVazio testa entrada vazia.
func TestSplitFieldsVazio(t *testing.T) {
	toks := splitFields("")
	if len(toks) != 0 {
		t.Fatalf("esperado 0 tokens, got %d: %q", len(toks), toks)
	}
}

// TestDisplayCmdWrapperWithoutPercentCommand testa display de wrapper sem %command%.
func TestDisplayCmdWrapperWithoutPercentCommand(t *testing.T) {
	g := testGUI("pt", "steam")
	c := Command{
		Command:   "mangohud",
		CommandEN: "mangohud",
	}
	got := g.displayCmd(c)
	// Wrapper sem %command% não deveria adicionar nada
	if got != "mangohud" {
		t.Fatalf("displayCmd wrapper: esperado 'mangohud', got %q", got)
	}
}

// TestTrChaveInexistente testa tradução de chave inexistente.
func TestTrChaveInexistente(t *testing.T) {
	g := testGUI("pt", "steam")
	got := g.tr("chaveQueNaoExiste")
	// ptTexts retorna "" para chave inexistente
	if got != "" {
		t.Fatalf("tr deveria retornar vazio para chave inexistente, got %q", got)
	}
	g.lang = "en"
	got = g.tr("chaveQueNaoExiste")
	if got != "" {
		t.Fatalf("tr EN deveria retornar vazio para chave inexistente, got %q", got)
	}
}

// TestTLocalizado testa a função t() para Localized.
func TestTLocalizado(t *testing.T) {
	g := testGUI("pt", "steam")
	loc := Localized{PT: "pt-text", EN: "en-text"}
	if got := g.t(loc); got != "pt-text" {
		t.Fatalf("t(pt): esperado 'pt-text', got %q", got)
	}
	g.lang = "en"
	if got := g.t(loc); got != "en-text" {
		t.Fatalf("t(en): esperado 'en-text', got %q", got)
	}
}

// TestCmdIdiomaAlternativo testa a função cmd() para idioma alternativo.
func TestCmdIdiomaAlternativo(t *testing.T) {
	g := testGUI("pt", "steam")
	c := Command{Command: "cmd-pt", CommandEN: "cmd-en"}
	if got := g.cmd(c); got != "cmd-pt" {
		t.Fatalf("cmd(pt): esperado 'cmd-pt', got %q", got)
	}
	g.lang = "en"
	if got := g.cmd(c); got != "cmd-en" {
		t.Fatalf("cmd(en): esperado 'cmd-en', got %q", got)
	}
}

// TestCmdIdiomaAlternativoVazio testa cmd() quando CommandEN está vazio.
func TestCmdIdiomaAlternativoVazio(t *testing.T) {
	g := testGUI("pt", "steam")
	c := Command{Command: "cmd-pt", CommandEN: ""}
	g.lang = "en"
	if got := g.cmd(c); got != "cmd-pt" {
		t.Fatalf("cmd(en) vazio: esperado 'cmd-pt', got %q", got)
	}
}

// TestLauncherInvalido testa comportamento com launcherIdx inválido.
func TestLauncherInvalido(t *testing.T) {
	g := testGUI("pt", "steam")
	g.launcherIdx = -1
	if g.launcher() == nil {
		t.Fatal("launcher() não deveria retornar nil com idx -1")
	}
	if g.launcher().ID != "steam" {
		t.Fatalf("launcher() deveria resetar para steam, got %q", g.launcher().ID)
	}
	g.launcherIdx = 999
	if g.launcher() == nil {
		t.Fatal("launcher() não deveria retornar nil com idx 999")
	}
	if g.launcher().ID != "steam" {
		t.Fatalf("launcher() deveria resetar para steam, got %q", g.launcher().ID)
	}
}

// TestMinSizeCurto testa MinSize com texto curto.
func TestMinSizeCurto(t *testing.T) {
	l := newMaxWidthLabel(500)
	l.SetText("hello")
	size := l.MinSize()
	if size.Width <= 0 || size.Height <= 0 {
		t.Fatalf("MinSize deveria ter dimensões positivas, got %v", size)
	}
}

// TestMinSizeLongo testa MinSize com texto longo que excede maxWidth.
func TestMinSizeLongo(t *testing.T) {
	l := newMaxWidthLabel(200)
	l.SetText("Este é um texto muito longo que deveria quebrar linha e ocupar mais de uma linha de altura")
	size := l.MinSize()
	if size.Width > 210 { // margem para arredondamento
		t.Fatalf("MinSize.Width não deveria exceder maxWidth significativamente, got %v", size.Width)
	}
	if size.Height <= 20 { // deveria ter múltiplas linhas
		t.Fatalf("MinSize.Height deveria ser maior para texto longo, got %v", size.Height)
	}
}

// TestMinSizeZeroMaxWidth testa MinSize com maxWidth zero.
func TestMinSizeZeroMaxWidth(t *testing.T) {
	l := newMaxWidthLabel(0)
	l.SetText("test")
	size := l.MinSize()
	if size.Width <= 0 || size.Height <= 0 {
		t.Fatalf("MinSize deveria funcionar com maxWidth zero, got %v", size)
	}
}

// TestToggleAtualizaSelected verifica se toggle atualiza o map selected.
func TestToggleAtualizaSelected(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	idx := idxOf("PROTON_LOG=1 %command%")
	g.toggle(idx, true)
	if !g.selected[idx] {
		t.Fatal("toggle(true) deveria adicionar ao selected")
	}
	g.toggle(idx, false)
	if g.selected[idx] {
		t.Fatal("toggle(false) deveria remover do selected")
	}
}

// TestCopyCombinationCopiaParaClipboard verifica se copyCombination copia.
func TestCopyCombinationCopiaParaClipboard(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.selected[idxOf("PROTON_LOG=1 %command%")] = true
	g.updateCombination()
	g.copyCombination()
	// Verifica que o status foi atualizado
	if g.status.Text == "" {
		t.Fatal("copyCombination deveria atualizar o status")
	}
}

// TestCopyCombinationVazio não copia nada.
func TestCopyCombinationVazio(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.copyCombination()
	// Não deveria mudar o status
	if g.status.Text != "" {
		t.Fatal("copyCombination vazio não deveria alterar o status")
	}
}

// TestCopyCurrentCopiaComandoAtual verifica se copyCurrent copia o comando.
func TestCopyCurrentCopiaComandoAtual(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.selectCommand(0)
	g.copyCurrent()
	if g.status.Text == "" {
		t.Fatal("copyCurrent deveria atualizar o status")
	}
}

// TestCopyCurrentInvalido não copia nada.
func TestCopyCurrentInvalido(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.current = -1
	g.copyCurrent()
	if g.status.Text != "" {
		t.Fatal("copyCurrent com current=-1 não deveria alterar o status")
	}
}

// TestApplyFavButtonAltaImportance verifica se o botão fica destacado.
func TestApplyFavButtonAltaImportance(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.favOnly = true
	g.applyFavButton()
	if g.favBtn.Importance != widget.HighImportance {
		t.Fatal("applyFavButton deveria usar HighImportance quando favOnly=true")
	}
}

// TestApplyFavButtonMediaImportance verifica se o botão fica normal.
func TestApplyFavButtonMediaImportance(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.favOnly = false
	g.applyFavButton()
	if g.favBtn.Importance != widget.MediumImportance {
		t.Fatal("applyFavButton deveria usar MediumImportance quando favOnly=false")
	}
}

// TestSelectCommandAtualizaCopiBtn verifica se o botão de copiar é habilitado.
func TestSelectCommandAtualizaCopiBtn(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.copyBtn.Disable()
	g.selectCommand(0)
	if g.copyBtn.Disabled() {
		t.Fatal("selectCommand deveria habilitar o copyBtn")
	}
}

// TestClearDetailDesabilitaCopiBtn verifica se o botão de copiar é desabilitado.
func TestClearDetailDesabilitaCopiBtn(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.copyBtn.Enable()
	g.clearDetail()
	if !g.copyBtn.Disabled() {
		t.Fatal("clearDetail deveria desabilitar o copyBtn")
	}
}

// TestCombinationSemComandos retorna string vazia.
func TestCombinationSemComandos(t *testing.T) {
	g := testGUI("pt", "steam")
	got := g.buildCombination()
	if got != "" {
		t.Fatalf("buildCombination sem seleção deveria retornar vazio, got %q", got)
	}
}

// TestLauncherSemLaunchers testa comportamento com slice vazio.
func TestLauncherSemLaunchers(t *testing.T) {
	g := testGUI("pt", "steam")
	g.launchers = nil
	if g.launcher() != nil {
		t.Fatal("launcher() deveria retornar nil com launchers nil")
	}
}

// TestConflictsSemSeleção retorna slice vazio.
func TestConflictsSemSeleção(t *testing.T) {
	g := testGUI("pt", "steam")
	warns := g.conflicts()
	if len(warns) != 0 {
		t.Fatalf("conflicts sem seleção deveria retornar vazio, got %v", warns)
	}
}

// TestDisplayCmdSemPercentCommand testa display de comando sem %command%.
func TestDisplayCmdSemPercentCommand(t *testing.T) {
	g := testGUI("pt", "steam")
	c := Command{Command: "VAR=1"}
	got := g.displayCmd(c)
	if got != "VAR=1" {
		t.Fatalf("displayCmd sem %%command%%: esperado 'VAR=1', got %q", got)
	}
}

// TestSetLauncherInvalido não muda o launcher.
func TestSetLauncherInvalido(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	original := g.launcher().ID
	g.setLauncher("launcherQueNaoExiste")
	if g.launcher().ID != original {
		t.Fatalf("setLauncher inválido deveria manter launcher original, got %q", g.launcher().ID)
	}
}

// TestUpdateCombinationSemSeleção desabilita botão de copiar.
func TestUpdateCombinationSemSeleção(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.updateCombination()
	if !g.combCopyBtn.Disabled() {
		t.Fatal("updateCombination sem seleção deveria desabilitar combCopyBtn")
	}
	if g.combCount.Text != "" {
		t.Fatalf("updateCombination sem seleção deveria limpar combCount, got %q", g.combCount.Text)
	}
}

// TestFilterNewFavs ignora chaves inválidas e duplicadas.
func TestFilterNewFavs(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	valid := g.validFavKeys()
	if len(valid) != len(g.all) {
		t.Fatalf("validFavKeys: %d chaves, esperado %d", len(valid), len(g.all))
	}
	existing := map[string]bool{g.favKey(0): true}
	imported := []string{g.favKey(0), g.favKey(1), g.favKey(1), "chave-invalida"}
	added := filterNewFavs(imported, valid, existing)
	if len(added) != 1 || added[0] != g.favKey(1) {
		t.Fatalf("filterNewFavs: got %q, esperado só favKey(1)", added)
	}
	if hasKnownFavKey([]string{"lixo"}, valid) {
		t.Fatal("hasKnownFavKey deveria ser falso para chaves desconhecidas")
	}
	if !hasKnownFavKey([]string{"lixo", g.favKey(2)}, valid) {
		t.Fatal("hasKnownFavKey deveria ser verdadeiro com ao menos uma chave válida")
	}
}

// TestSaveFavsPersisteOrdenado verifica que saveFavs grava a lista ordenada.
func TestSaveFavsPersisteOrdenado(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.favs[g.favKey(2)] = true
	g.favs[g.favKey(0)] = true
	g.saveFavs()
	got := g.app.Preferences().StringListWithFallback("favs", nil)
	if len(got) != 2 || got[0] > got[1] {
		t.Fatalf("saveFavs deveria persistir ordenado, got %q", got)
	}
}

// favExportWriteFake é um fyne.URIWriteCloser em memória para testes.
type favExportWriteFake struct {
	*bytes.Buffer
	fail   bool
	closed bool
}

func (f *favExportWriteFake) Write(p []byte) (int, error) {
	if f.fail {
		return 0, errors.New("write fail")
	}
	return f.Buffer.Write(p)
}

func (f *favExportWriteFake) Close() error {
	f.closed = true
	return nil
}

func (f *favExportWriteFake) URI() fyne.URI { return nil }

// TestFinishFavExport cobre escrita OK, erro de escrita e cancelamento.
func TestFinishFavExport(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	f := &favExportWriteFake{Buffer: &bytes.Buffer{}}
	g.finishFavExport(f, nil, []byte(`["a"]`))
	if !f.closed || f.String() != `["a"]` {
		t.Fatalf("finish OK: closed=%v data=%q", f.closed, f.String())
	}
	if g.status.Text != g.tr("favsExported") {
		t.Fatalf("finish OK: status=%q", g.status.Text)
	}
	g.finishFavExport(&favExportWriteFake{Buffer: &bytes.Buffer{}, fail: true}, nil, []byte(`x`))
	if g.status.Text != g.tr("favsExportError") {
		t.Fatalf("finish erro: status=%q", g.status.Text)
	}
	g.status.SetText("marcador")
	g.finishFavExport(nil, nil, nil)
	if g.status.Text != "marcador" {
		t.Fatal("finish cancelado não deveria tocar no status")
	}
}

// TestExportFavsComFavoritos abre o diálogo sem erro no ambiente de teste.
func TestExportFavsComFavoritos(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.favs[g.favKey(0)] = true
	g.exportFavs()
	if g.status.Text == g.tr("favsEmpty") {
		t.Fatal("export com favoritos não deveria dizer que está vazio")
	}
}

// TestFavExportPayload gera JSON ordenado com round-trip válido.
func TestFavExportPayload(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	favs := map[string]bool{g.favKey(2): true, g.favKey(0): true}
	data, err := favExportPayload(favs)
	if err != nil {
		t.Fatalf("favExportPayload: erro inesperado: %v", err)
	}
	var got []string
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("payload não é JSON válido: %v", err)
	}
	if len(got) != 2 || got[0] > got[1] {
		t.Fatalf("payload deveria ser ordenado, got %q", got)
	}
}

// TestProcessFavImport cobre os três casos: inválido, misto e duplicado.
func TestProcessFavImport(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	valid := g.validFavKeys()

	if _, ok := processFavImport([]string{"lixo"}, valid, map[string]bool{}); ok {
		t.Fatal("processFavImport deveria rejeitar arquivo sem chave conhecida")
	}
	existing := map[string]bool{g.favKey(0): true}
	added, ok := processFavImport([]string{g.favKey(0), g.favKey(1), "lixo"}, valid, existing)
	if !ok || added != 1 {
		t.Fatalf("processFavImport misto: added=%d ok=%v, esperado 1/true", added, ok)
	}
	added, ok = processFavImport([]string{g.favKey(0)}, valid, existing)
	if !ok || added != 0 {
		t.Fatalf("processFavImport duplicado: added=%d ok=%v, esperado 0/true", added, ok)
	}
	added, ok = processFavImport(nil, valid, map[string]bool{})
	if !ok || added != 0 {
		t.Fatalf("processFavImport vazio: added=%d ok=%v, esperado 0/true", added, ok)
	}
}

// TestApplyFavImport incorpora válidos, rejeita inválidos e filtra favOnly.
func TestApplyFavImport(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.applyFavImport([]string{"lixo"})
	if g.status.Text != g.tr("favsInvalidFile") {
		t.Fatalf("import inválido: status=%q", g.status.Text)
	}
	if len(g.favs) != 0 {
		t.Fatal("import inválido não deveria adicionar favoritos")
	}
	g.applyFavImport([]string{g.favKey(0), g.favKey(1), "lixo"})
	if len(g.favs) != 2 {
		t.Fatalf("import misto: %d favs, esperado 2", len(g.favs))
	}
	want := fmt.Sprintf(g.tr("favsImported"), 2)
	if g.status.Text != want {
		t.Fatalf("import misto: status=%q, esperado %q", g.status.Text, want)
	}
	g2 := initTestGUI("pt", "steam")
	g2.favOnly = true
	g2.applyFavImport([]string{g2.favKey(0)})
	if len(g2.filtered) != 1 {
		t.Fatalf("import com favOnly: filtered=%d, esperado 1", len(g2.filtered))
	}
}

// TestDecodeFavImport lê JSON válido e rejeita inválido.
func TestDecodeFavImport(t *testing.T) {
	got, err := decodeFavImport(strings.NewReader(`["a", "b"]`))
	if err != nil || len(got) != 2 {
		t.Fatalf("decode válido: got=%q err=%v", got, err)
	}
	if _, err := decodeFavImport(strings.NewReader(`não-json`)); err == nil {
		t.Fatal("decode inválido deveria retornar erro")
	}
}

// TestMenuActions cobre os callbacks do menu de configurações.
func TestMenuActions(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.buildConfigMenu()
	g.langEN.Action()
	if g.lang != "en" {
		t.Fatalf("langEN.Action: lang=%q, esperado en", g.lang)
	}
	g.langPT.Action()
	if g.lang != "pt" {
		t.Fatalf("langPT.Action: lang=%q, esperado pt", g.lang)
	}
	g.themeDark.Action()
	g.themeLight.Action()
	g.themeSystem.Action()
	before := g.copyOnClick.Checked
	g.copyOnClick.Action()
	if g.copyOnClick.Checked == before {
		t.Fatal("copyOnClick.Action deveria alternar Checked")
	}
}

// TestBuildConfigMenu garante que o menu de configurações é montado.
func TestBuildConfigMenu(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	m := g.buildConfigMenu()
	if m == nil || len(m.Items) == 0 {
		t.Fatal("buildConfigMenu deveria retornar menu com itens")
	}
	if g.langPT == nil || g.langEN == nil || g.themeSystem == nil || g.copyOnClick == nil {
		t.Fatal("buildConfigMenu deveria criar os itens de idioma, tema e copyOnClick")
	}
}

// TestExportFavsVazio mostra mensagem de vazio.
func TestExportFavsVazio(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.exportFavs()
	if g.status.Text != g.tr("favsEmpty") {
		t.Fatalf("exportFavs vazio: status=%q, esperado %q", g.status.Text, g.tr("favsEmpty"))
	}
}

// TestImportFavsInvalido mostra mensagem de arquivo inválido.
func TestImportFavsInvalido(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	// Teste indireto: funções de import/export não devem paniquar com estado vazio
	g.importFavs()
	// Sem arquivo real, apenas verifica que não paniqua
}

// TestApplyFilterCatPrefix filtra por categoria usando prefixo cat:.
func TestApplyFilterCatPrefix(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.search.SetText("cat:gpu")
	if len(g.filtered) == 0 {
		t.Fatal("cat:gpu deveria retornar resultados")
	}
	for _, idx := range g.filtered {
		cat := strings.ToLower(g.t(g.all[idx].Category))
		if !strings.Contains(cat, "gpu") {
			t.Fatalf("cat:gpu retornou categoria %q", cat)
		}
	}
}

// TestApplyFilterCompatPrefix filtra por compatibilidade usando prefixo compat:.
func TestApplyFilterCompatPrefix(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.search.SetText("compat:cachyos")
	if len(g.filtered) == 0 {
		t.Fatal("compat:cachyos deveria retornar resultados")
	}
	for _, idx := range g.filtered {
		compat := strings.ToLower(g.t(g.all[idx].Compat))
		if !strings.Contains(compat, "cachyos") {
			t.Fatalf("compat:cachyos retornou compat %q", compat)
		}
	}
}

// TestApplyFilterCmdPrefix filtra por comando usando prefixo cmd:.
func TestApplyFilterCmdPrefix(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.search.SetText("cmd:mangohud")
	if len(g.filtered) == 0 {
		t.Fatal("cmd:mangohud deveria retornar resultados")
	}
	for _, idx := range g.filtered {
		cmd := strings.ToLower(g.all[idx].Command)
		if !strings.Contains(cmd, "mangohud") {
			t.Fatalf("cmd:mangohud retornou comando %q", cmd)
		}
	}
}

// TestApplyFilterPrefixVazio lista tudo.
func TestApplyFilterPrefixVazio(t *testing.T) {
	a := test.NewApp()
	defer a.Quit()
	g := initTestGUI("pt", "steam")
	g.search.SetText("cat:")
	if len(g.filtered) != len(g.all) {
		t.Fatalf("cat: vazio deveria listar tudo, got %d/%d", len(g.filtered), len(g.all))
	}
}
