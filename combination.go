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
			case "PROTON_NVIDIA_LIBS":
				hasNvidiaLibs = true
			case "PROTON_USE_WOW64":
				hasWow64 = true
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
	if hasNvidiaLibs && hasWow64 {
		out = append(out, g.tr("conflictNvidiaLibsWow64"))
	}
	return out
}
