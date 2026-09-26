package main

import (
	"fmt"
	"sort"
	"strings"
)

var wrapperNames = []string{"mangohud", "gamemoderun", "gamescope", "game-performance"}

// wrapperBin devolve o nome do programa wrapper que a receita invoca, ou
// "" se não for wrapper. Procura em qualquer posição e compara o token
// inteiro: MANGOHUD=1 mangohud %command% é wrapper mesmo com a variável
// na frente, e testar só o prefixo da string não pega esse caso.
func (g *gui) wrapperBin(c Command) string {
	for _, tok := range splitFields(c.Command) {
		t := strings.ToLower(tok)
		for _, w := range wrapperNames {
			if t == w {
				return t
			}
		}
	}
	return ""
}

// uniq preserva a ordem de-appearance e remove repetidos, para não
// duplicar aviso quando a mesma variável vem em receitas diferentes.
func uniq(in []string) []string {
	visto := map[string]bool{}
	var out []string
	for _, v := range in {
		if visto[v] {
			continue
		}
		visto[v] = true
		out = append(out, v)
	}
	return out
}

func (g *gui) isWrapper(c Command) bool {
	return g.wrapperBin(c) != ""
}

// splitFields divide como strings.Fields, mas preserva trechos entre
// aspas simples/duplas (ex.: DXVK_CONFIG="a = b" vira um token só) e
// respeita barra invertida como escape dentro da citação, como o shell.
func splitFields(s string) []string {
	var out []string
	var cur strings.Builder
	var quote byte
	flush := func() {
		if cur.Len() > 0 {
			out = append(out, cur.String())
			cur.Reset()
		}
	}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if quote != 0 {
			if ch == '\\' && i+1 < len(s) {
				i++
				cur.WriteByte(s[i])
				continue
			}
			cur.WriteByte(ch)
			if ch == quote {
				quote = 0
			}
			continue
		}
		switch {
		case ch == '"' || ch == '\'':
			quote = ch
			cur.WriteByte(ch)
		case ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r':
			flush()
		default:
			cur.WriteByte(ch)
		}
	}
	flush()
	return out
}

// adapt normaliza um comando para o launcher escolhido. No Steam e no
// Faugus o %command% é preservado, porque é ele que o launcher expande;
// nos demais o placeholder e o separador "--" saem, já que não têm
// sentido fora deles.
func (g *gui) adapt(s string) string {
	if g.launcher().HasCmd {
		return s
	}
	s = strings.TrimSpace(strings.TrimSuffix(s, "%command%"))
	return strings.TrimSpace(strings.TrimSuffix(s, "--"))
}

func (g *gui) buildCombination() string {
	var wrappers, envs []string
	for i := range g.all {
		if !g.selected[i] {
			continue
		}
		c := g.all[i]
		base := strings.TrimSpace(strings.TrimSuffix(g.cmd(c), "%command%"))
		if !g.launcher().HasCmd {
			base = strings.TrimSpace(strings.TrimSuffix(base, "--"))
		}
		if g.isWrapper(c) {
			wrappers = append(wrappers, base)
		} else {
			envs = append(envs, base)
		}
	}
	parts := append(envs, wrappers...)
	if len(parts) == 0 {
		return ""
	}
	line := strings.Join(parts, " ")
	if g.launcher().HasCmd {
		line += " %command%"
	}
	return line
}

// displayCmd mostra o mesmo que buildCombination monta, para o painel de
// detalhe e o botão de copiar não divergirem da combinação.
func (g *gui) displayCmd(c Command) string {
	return g.adapt(g.cmd(c))
}

// conflicts detecta variáveis duplicadas com valores diferentes e
// opções mutuamente exclusivas entre os comandos selecionados.
func (g *gui) conflicts() []string {
	var out []string
	vals := map[string]map[string]bool{}
	var keys []string
	var antiLagSo, reflexSo, antiLagComReflex bool
	hasMesaAntiLag, hasFsr4Upgrade := false, false
	hasNvidiaLibs, hasWow64 := false, false
	var spoofNvidia []string
	var fsr4Upgrades []string
	wrapperVisto := map[string]string{}
	for i := range g.all {
		if !g.selected[i] {
			continue
		}
		s := g.cmd(g.all[i])
		toks := splitFields(s)
		cmdReflex, cmdAntiLag := false, false
		for _, tok := range toks {
			name := tok
			if idx := strings.IndexByte(tok, '='); idx >= 0 {
				name = tok[:idx]
			}
			switch name {
			case "LOW_LATENCY_LAYER_REFLEX":
				cmdReflex = true
			case "LOW_LATENCY_LAYER":
				cmdAntiLag = true
			case "ENABLE_LAYER_MESA_ANTI_LAG":
				hasMesaAntiLag = true
			case "PROTON_FSR4_UPGRADE":
				hasFsr4Upgrade = true
				fsr4Upgrades = append(fsr4Upgrades, name)
			case "PROTON_FSR4_RDNA3_UPGRADE", "PROTON_FFX4_UPGRADE":
				// PROTON_FFX4_UPGRADE é o nome atual do upgrade de FSR 4 no
				// CachyOS 11+ e controla as mesmas versões (CHANGELOG do
				// CachyOS). Sem ele aqui, a regra pegava o nome antigo e
				// deixava passar o caminho que está em uso hoje.
				fsr4Upgrades = append(fsr4Upgrades, name)
			case "PROTON_NVIDIA_LIBS":
				hasNvidiaLibs = true
			case "PROTON_USE_WOW64":
				hasWow64 = true
			case "PROTON_FORCE_NVAPI":
				spoofNvidia = append(spoofNvidia, "PROTON_FORCE_NVAPI=1")
			case "LOW_LATENCY_LAYER_SPOOF_NVIDIA":
				spoofNvidia = append(spoofNvidia, "LOW_LATENCY_LAYER_SPOOF_NVIDIA=1")
			}
			if tok == "%command%" || tok == "--" {
				continue
			}
			k, v, ok := strings.Cut(tok, "=")
			if !ok || k == "" {
				continue
			}
			if vals[k] == nil {
				vals[k] = map[string]bool{}
				keys = append(keys, k)
			}
			vals[k][v] = true
		}
		switch {
		case cmdReflex && cmdAntiLag:
			antiLagComReflex = true
		case cmdReflex:
			reflexSo = true
		case cmdAntiLag:
			antiLagSo = true
		}
		if bin := g.wrapperBin(g.all[i]); bin != "" {
			{
				if outro, repete := wrapperVisto[bin]; repete {
					out = append(out, fmt.Sprintf(g.tr("conflictSameWrapper"), bin, outro, g.all[i].Command))
				} else {
					wrapperVisto[bin] = g.all[i].Command
				}
			}
		}
	}
	for _, k := range keys {
		if len(vals[k]) > 1 {
			vs := make([]string, 0, len(vals[k]))
			for v := range vals[k] {
				vs = append(vs, v)
			}
			sort.Strings(vs)
			out = append(out, fmt.Sprintf(g.tr("conflictDuplicate"), k, strings.Join(vs, ", ")))
		}
	}
	if antiLagSo && reflexSo {
		out = append(out, g.tr("conflictAntiLagReflex"))
	}
	if antiLagSo && antiLagComReflex {
		out = append(out, g.tr("conflictAntiLagRedundant"))
	}
	if hasMesaAntiLag && (antiLagSo || antiLagComReflex || reflexSo) {
		out = append(out, g.tr("conflictMesaAntiLagLayer"))
	}
	if hasMesaAntiLag && hasFsr4Upgrade {
		out = append(out, g.tr("conflictMesaAntiLagFsr4"))
	}
	// Spoofing de GPU + upgrade FSR4: o README do low_latency_layer diz que
	// PROTON_FORCE_NVAPI e LOW_LATENCY_LAYER_SPOOF_NVIDIA quebram o caminho
	// de upgrade FSR4, e não é caso raro. Sem esta regra o app montava a
	// combinação sem aviso e o jogo caía para FSR 3.1 sozinho.
	if len(spoofNvidia) > 0 {
		spoofs := uniq(spoofNvidia)
		for _, up := range uniq(fsr4Upgrades) {
			out = append(out, fmt.Sprintf(g.tr("conflictFsr4NvapiSpoof"), strings.Join(spoofs, " / "), up))
		}
	}
	if hasNvidiaLibs && hasWow64 {
		out = append(out, g.tr("conflictNvidiaLibsWow64"))
	}
	return out
}
