package main

import (
	"strings"
	"testing"
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
