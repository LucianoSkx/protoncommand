package main

import (
	"fmt"
	"sort"
	"strings"
)

func (g *gui) isWrapper(c Command) bool {
	for _, w := range []string{"mangohud", "gamemoderun", "gamescope", "game-performance"} {
		if strings.HasPrefix(c.Command, w) {
			return true
		}
	}
	return false
}

// splitFields divide como strings.Fields, mas preserva trechos entre
// aspas simples/duplas (ex.: DXVK_CONFIG="a = b" vira um token só).
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

// displayCmd adapta o comando ao launcher escolhido, removendo o
// placeholder %command% quando o launcher não o usa.
func (g *gui) displayCmd(c Command) string {
	s := g.cmd(c)
	if !g.launcher().HasCmd {
		s = strings.TrimSpace(strings.TrimSuffix(s, "%command%"))
	}
	return s
}

// conflicts detecta variáveis duplicadas com valores diferentes e
// opções mutuamente exclusivas entre os comandos selecionados.
func (g *gui) conflicts() []string {
	var out []string
	vals := map[string]map[string]bool{}
	var keys []string
	hasAntiLag, hasReflex, hasMesaAntiLag, hasFsr4Upgrade := false, false, false, false
	hasNvidiaLibs, hasWow64 := false, false
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
		if cmdReflex {
			hasReflex = true
		} else if cmdAntiLag {
			hasAntiLag = true
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
	if hasAntiLag && hasReflex {
		out = append(out, g.tr("conflictAntiLagReflex"))
	}
	if hasMesaAntiLag && (hasAntiLag || hasReflex) {
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
