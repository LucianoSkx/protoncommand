package main

import (
	"sort"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// maxWidthLabel é um Label com quebra por palavra e largura máxima.
// Sem isso, o MinSize do Label mede o texto sem quebra e textos longos
// de uma linha só esticam o container (e a janela) além da tela.
type maxWidthLabel struct {
	widget.Label
	maxWidth float32
}

func newMaxWidthLabel(maxWidth float32) *maxWidthLabel {
	l := &maxWidthLabel{maxWidth: maxWidth}
	l.ExtendBaseWidget(l)
	l.Wrapping = fyne.TextWrapWord
	return l
}

// MinSize limita a largura e estima a altura com quebra por palavra,
// para o texto todo ficar visível sem esticar a janela.
func (l *maxWidthLabel) MinSize() fyne.Size {
	full := fyne.MeasureText(l.Text, theme.TextSize(), l.TextStyle)
	pad := theme.Padding()
	if l.maxWidth <= 0 || full.Width <= l.maxWidth {
		return fyne.NewSize(full.Width+pad, full.Height+pad)
	}
	lineH := fyne.MeasureText("Ag", theme.TextSize(), l.TextStyle).Height
	if lineH <= 0 {
		return fyne.NewSize(l.maxWidth, full.Height+pad)
	}
	spaceW := fyne.MeasureText(" ", theme.TextSize(), l.TextStyle).Width
	lines := 0
	for _, para := range strings.Split(l.Text, "\n") {
		n := 0
		w := float32(0)
		for _, word := range strings.Fields(para) {
			ww := fyne.MeasureText(word, theme.TextSize(), l.TextStyle).Width
			if n == 0 {
				w, n = ww, 1
				continue
			}
			if w+spaceW+ww > l.maxWidth {
				lines++
				w, n = ww, 1
				continue
			}
			w += spaceW + ww
		}
		lines++
	}
	lines++ // margem: a quebra do Fyne pode usar uma linha a mais
	if lines < 1 {
		lines = 1
	}
	return fyne.NewSize(l.maxWidth, float32(lines)*lineH+pad)
}

type gui struct {
	app  fyne.App
	win  fyne.Window
	lang string

	list         *widget.List
	search       *widget.Entry
	logo         *canvas.Image
	langRadio    *widget.RadioGroup
	configBtn    *widget.Button
	aboutBtn     *widget.Button
	detailTitle  *widget.Label
	detailCat    *widget.Label
	detailCompat *widget.Label
	detailDesc   *maxWidthLabel
	detailCmd    *maxWidthLabel
	copyBtn      *widget.Button
	favToggleBtn *widget.Button
	status       *maxWidthLabel

	langHeader  *fyne.MenuItem
	langPT      *fyne.MenuItem
	langEN      *fyne.MenuItem
	themeSystem *fyne.MenuItem
	themeLight  *fyne.MenuItem
	themeDark   *fyne.MenuItem
	copyOnClick *fyne.MenuItem

	combLabel   *maxWidthLabel
	combCount   *widget.Label
	combCopyBtn *widget.Button
	clearBtn    *widget.Button
	combWarn    *canvas.Text
	combHint    *widget.Label

	launcherSel *widget.Select
	launchers   []Launcher
	launcherIdx int
	catSel      *widget.Select
	catOptions  []string
	catFilterPT string
	favBtn      *widget.Button
	favOnly     bool

	all      []Command
	filtered []int
	selected map[int]bool
	favs     map[string]bool
	current  int
	selID    int
}

func main() {
	a := app.NewWithID("br.com.protoncommands")
	a.SetIcon(resourceIcon)
	w := a.NewWindow("Proton Command")
	w.Resize(fyne.NewSize(980, 640))

	g := &gui{
		app:      a,
		win:      w,
		all:      commands(),
		lang:     "pt",
		selected: map[int]bool{},
		favs:     map[string]bool{},
		current:  -1,
		selID:    -1,
	}
	g.launchers = launchers()
	g.launcherIdx = 0
	if l := a.Preferences().StringWithFallback("lang", "pt"); l == "en" || l == "pt" {
		g.lang = l
	}
	if id := a.Preferences().StringWithFallback("launcher", "steam"); id != "" {
		for i := range g.launchers {
			if g.launchers[i].ID == id {
				g.launcherIdx = i
			}
		}
	}
	for _, s := range a.Preferences().StringListWithFallback("favs", nil) {
		g.favs[s] = true
	}
	g.filtered = make([]int, len(g.all))
	for i := range g.all {
		g.filtered[i] = i
	}
	g.build()
	w.ShowAndRun()
}

func (g *gui) tr(key string) string {
	if g.lang == "en" {
		if s, ok := enTexts[key]; ok {
			return s
		}
	}
	return ptTexts[key]
}

func (g *gui) t(l Localized) string {
	if g.lang == "en" {
		return l.EN
	}
	return l.PT
}

func (g *gui) cmd(c Command) string {
	if g.lang == "en" && c.CommandEN != "" {
		return c.CommandEN
	}
	return c.Command
}

// launcher devolve o launcher atual sem guardar ponteiro para dentro
// do slice (o índice sobrevive a realocações e cópias da struct).
func (g *gui) launcher() *Launcher {
	if len(g.launchers) == 0 {
		return nil
	}
	if g.launcherIdx < 0 || g.launcherIdx >= len(g.launchers) {
		g.launcherIdx = 0
	}
	return &g.launchers[g.launcherIdx]
}

func (g *gui) build() {
	g.search = widget.NewEntry()
	g.search.OnChanged = func(_ string) { g.applyFilter() }

	g.catSel = widget.NewSelect(nil, func(v string) {
		g.catFilterPT = ""
		for _, c := range g.all {
			if g.t(c.Category) == v {
				g.catFilterPT = c.Category.PT
				break
			}
		}
		g.applyFilter()
	})

	g.favBtn = widget.NewButton("", func() {
		g.favOnly = !g.favOnly
		g.applyFavButton()
		g.applyFilter()
	})

	filterRow := container.NewBorder(nil, nil, g.catSel, g.favBtn)

	g.list = widget.NewList(
		func() int { return len(g.filtered) },
		func() fyne.CanvasObject {
			check := widget.NewCheck("", nil)
			return container.NewHBox(
				check,
				container.NewVBox(
					widget.NewLabel(""), // título
					widget.NewLabel(""), // categoria
				),
			)
		},
		func(id int, obj fyne.CanvasObject) {
			idx := g.filtered[id]
			c := g.all[idx]
			box := obj.(*fyne.Container)
			check := box.Objects[0].(*widget.Check)
			inner := box.Objects[1].(*fyne.Container)
			title := inner.Objects[0].(*widget.Label)
			cat := inner.Objects[1].(*widget.Label)
			title.SetText(g.t(c.Title))
			title.TextStyle = fyne.TextStyle{Bold: true}
			cat.SetText(g.t(c.Category) + " · " + g.t(c.Compat))
			cat.TextStyle = fyne.TextStyle{Italic: true}
			check.OnChanged = nil
			check.SetChecked(g.selected[idx])
			check.OnChanged = func(v bool) { g.toggle(idx, v) }
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

	g.detailTitle = widget.NewLabel("")
	g.detailTitle.TextStyle = fyne.TextStyle{Bold: true}
	g.detailTitle.Wrapping = fyne.TextWrapWord

	g.detailCat = widget.NewLabel("")
	g.detailCat.TextStyle = fyne.TextStyle{Italic: true}
	g.detailCat.Wrapping = fyne.TextWrapWord

	g.detailCompat = widget.NewLabel("")
	g.detailCompat.TextStyle = fyne.TextStyle{Italic: true}
	g.detailCompat.Wrapping = fyne.TextWrapWord

	g.detailCmd = newMaxWidthLabel(500)
	g.detailCmd.TextStyle = fyne.TextStyle{Monospace: true, Bold: true}

	g.detailDesc = newMaxWidthLabel(500)

	g.copyBtn = widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		if g.current < 0 {
			return
		}
		g.copyCurrent()
	})
	g.copyBtn.Importance = widget.HighImportance

	g.favToggleBtn = widget.NewButton("", func() {
		if g.current < 0 {
			return
		}
		g.toggleFav(g.current)
	})

	g.status = newMaxWidthLabel(500)
	g.status.TextStyle = fyne.TextStyle{Italic: true}

	detail := container.NewBorder(
		container.NewVBox(
			g.detailTitle,
			g.detailCat,
			g.detailCompat,
			widget.NewSeparator(),
			g.detailCmd,
			widget.NewSeparator(),
		),
		container.NewVBox(
			container.NewHBox(g.copyBtn, g.favToggleBtn),
			g.status,
		),
		nil, nil,
		container.NewVScroll(g.detailDesc),
	)
	detail.Resize(fyne.NewSize(520, 560))

	left := container.NewBorder(
		container.NewVBox(g.search, filterRow),
		nil, nil, nil,
		g.list,
	)

	split := container.NewHSplit(left, detail)
	split.Offset = 0.42

	configMenu := g.buildConfigMenu()

	g.logo = canvas.NewImageFromResource(resourceLogo)
	g.logo.FillMode = canvas.ImageFillContain
	g.logo.SetMinSize(fyne.NewSize(280, 80))
	g.logo.Resize(fyne.NewSize(280, 80))

	g.langRadio = widget.NewRadioGroup([]string{"Português", "Inglês"}, func(v string) {
		if v == "Inglês" {
			g.setLang("en", true)
		} else if v != "" {
			g.setLang("pt", true)
		}
	})
	g.langRadio.Horizontal = true

	var configBtn *widget.Button
	configBtn = widget.NewButton("", func() {
		pop := widget.NewPopUpMenu(configMenu, g.win.Canvas())
		pos := fyne.CurrentApp().Driver().AbsolutePositionForObject(configBtn)
		pop.ShowAtPosition(pos.AddXY(0, configBtn.Size().Height))
	})
	g.configBtn = configBtn

	g.aboutBtn = widget.NewButton("", func() {
		dialog.NewInformation(g.tr("about"), g.tr("aboutBody"), g.win).Show()
	})

	titleBar := container.NewBorder(
		nil, nil, nil,
		container.NewHBox(g.configBtn, g.aboutBtn),
		container.NewCenter(g.logo),
	)

	g.launcherSel = widget.NewSelect(nil, func(name string) {
		for i := range g.launchers {
			if g.t(g.launchers[i].Name) == name {
				g.setLauncher(g.launchers[i].ID)
				return
			}
		}
	})

	topRow := container.NewBorder(
		nil, nil, g.launcherSel, nil,
		container.NewCenter(g.langRadio),
	)

	g.combLabel = newMaxWidthLabel(940)

	g.combCount = widget.NewLabel("")
	g.combCount.TextStyle = fyne.TextStyle{Italic: true}

	g.combWarn = canvas.NewText("", theme.ErrorColor())
	g.combWarn.TextSize = 13
	g.combWarn.TextStyle = fyne.TextStyle{Bold: true}

	g.combHint = widget.NewLabel("")
	g.combHint.TextStyle = fyne.TextStyle{Italic: true}
	g.combHint.Wrapping = fyne.TextWrapWord

	g.combCopyBtn = widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		g.copyCombination()
	})
	g.combCopyBtn.Importance = widget.HighImportance

	g.clearBtn = widget.NewButton("", func() {
		g.clearSelection()
	})

	combBar := container.NewBorder(
		container.NewHBox(
			g.combCount,
			layout.NewSpacer(),
			g.clearBtn,
		),
		nil, nil, nil,
		container.NewVBox(
			g.combWarn,
			g.combLabel,
			g.combHint,
			g.combCopyBtn,
		),
	)

	g.win.SetContent(container.NewBorder(
		container.NewVBox(titleBar, topRow),
		combBar, nil, nil, split,
	))

	g.applyLang()
	g.copyOnClick.Checked = g.app.Preferences().BoolWithFallback("copyOnClick", false)
	g.setTheme(g.app.Preferences().StringWithFallback("theme", "system"), false)
	g.updateCombination()
	if len(g.filtered) > 0 {
		g.list.Select(0)
	}
}

func (g *gui) buildConfigMenu() *fyne.Menu {
	g.langHeader = fyne.NewMenuItem("", nil)
	g.langHeader.Disabled = true
	g.langPT = fyne.NewMenuItem("Português", func() {
		g.setLang("pt", true)
	})
	g.langEN = fyne.NewMenuItem("Inglês", func() {
		g.setLang("en", true)
	})
	g.themeSystem = fyne.NewMenuItem("", func() {
		g.setTheme("system", true)
	})
	g.themeLight = fyne.NewMenuItem("", func() {
		g.setTheme("light", true)
	})
	g.themeDark = fyne.NewMenuItem("", func() {
		g.setTheme("dark", true)
	})
	g.copyOnClick = fyne.NewMenuItem("", func() {
		g.copyOnClick.Checked = !g.copyOnClick.Checked
		g.app.Preferences().SetBool("copyOnClick", g.copyOnClick.Checked)
	})
	return fyne.NewMenu("",
		g.langHeader,
		g.langPT,
		g.langEN,
		fyne.NewMenuItemSeparator(),
		g.themeSystem,
		g.themeLight,
		g.themeDark,
		fyne.NewMenuItemSeparator(),
		g.copyOnClick,
	)
}

func (g *gui) setLang(lang string, persist bool) {
	if lang == "en" || lang == "pt" {
		g.lang = lang
	}
	if persist {
		g.app.Preferences().SetString("lang", g.lang)
	}
	g.applyLang()
	// Limpa antes: Select não dispara OnSelected quando o id já
	// estava selecionado, e o detalhe ficaria com o texto antigo.
	g.list.UnselectAll()
	g.list.Refresh()
	if g.selID >= 0 && g.selID < len(g.filtered) {
		g.list.Select(g.selID)
	} else if len(g.filtered) > 0 {
		g.list.Select(0)
	}
}

func (g *gui) applyLang() {
	if g.lang == "en" {
		g.langRadio.SetSelected("Inglês")
	} else {
		g.langRadio.SetSelected("Português")
	}
	g.search.SetPlaceHolder(g.tr("searchPlaceholder"))
	g.configBtn.Text = g.tr("settings")
	g.configBtn.Refresh()
	g.aboutBtn.Text = g.tr("about")
	g.aboutBtn.Refresh()
	g.copyBtn.Text = g.tr("copyCommand")
	g.copyBtn.Refresh()
	g.langHeader.Label = g.tr("language")
	g.langPT.Checked = g.lang == "pt"
	g.langEN.Checked = g.lang == "en"
	g.themeSystem.Label = g.tr("themeSystem")
	g.themeLight.Label = g.tr("themeLight")
	g.themeDark.Label = g.tr("themeDark")
	g.copyOnClick.Label = g.tr("copyOnClick")
	g.combCopyBtn.Text = g.tr("copyCombination")
	g.combCopyBtn.Refresh()
	g.clearBtn.Text = g.tr("clearSelection")
	g.clearBtn.Refresh()
	g.favBtn.Text = "★ " + g.tr("favoritesOnly")
	g.favBtn.Refresh()
	g.applyFavButton()

	names := make([]string, len(g.launchers))
	for i := range g.launchers {
		names[i] = g.t(g.launchers[i].Name)
	}
	g.launcherSel.Options = names
	if l := g.launcher(); l != nil {
		g.launcherSel.SetSelected(g.t(l.Name))
	}

	cats := map[string]bool{}
	for _, c := range g.all {
		cats[g.t(c.Category)] = true
	}
	g.catOptions = []string{g.tr("allCategories")}
	for c := range cats {
		g.catOptions = append(g.catOptions, c)
	}
	sort.Strings(g.catOptions[1:])
	g.catSel.Options = g.catOptions
	if g.catFilterPT != "" {
		idx := -1
		for i, o := range g.catOptions {
			for _, c := range g.all {
				if g.t(c.Category) == o && c.Category.PT == g.catFilterPT {
					idx = i
				}
			}
		}
		if idx > 0 {
			g.catSel.SetSelected(g.catOptions[idx])
		} else {
			g.catFilterPT = ""
			g.catSel.SetSelected(g.catOptions[0])
		}
	} else {
		g.catSel.SetSelected(g.catOptions[0])
	}

	g.updateCombination()
	g.refreshDetail()
	g.updateFavButton()
}

func (g *gui) refreshDetail() {
	if g.current >= 0 {
		c := g.all[g.current]
		g.detailTitle.SetText(g.t(c.Title))
		g.detailCat.SetText(g.tr("category") + g.t(c.Category))
		g.detailCompat.SetText(g.tr("compatible") + g.t(c.Compat))
		g.detailCmd.SetText(g.displayCmd(c))
		g.detailDesc.SetText(g.t(c.Description))
	} else {
		g.detailTitle.SetText(g.tr("noCommand"))
		g.detailCat.SetText("")
		g.detailCompat.SetText("")
		g.detailCmd.SetText("")
		g.detailDesc.SetText("")
	}
	g.updateFavButton()
}

func (g *gui) setTheme(name string, persist bool) {
	switch name {
	case "light":
		g.app.Settings().SetTheme(theme.LightTheme())
	case "dark":
		g.app.Settings().SetTheme(theme.DarkTheme())
	default:
		g.app.Settings().SetTheme(theme.DefaultTheme())
	}
	g.themeSystem.Checked = name == "system"
	g.themeLight.Checked = name == "light"
	g.themeDark.Checked = name == "dark"
	if persist {
		g.app.Preferences().SetString("theme", name)
	}
}

func (g *gui) toggle(idx int, v bool) {
	if v {
		g.selected[idx] = true
	} else {
		delete(g.selected, idx)
	}
	g.updateCombination()
}

func (g *gui) updateCombination() {
	n := len(g.selected)
	comb := g.buildCombination()
	g.combHint.SetText(g.t(g.launcher().Hint))
	if n == 0 {
		g.combLabel.SetText(g.tr("noCommandSelected"))
		g.combCount.SetText("")
		g.combWarn.Text = ""
		g.combWarn.Refresh()
		g.combCopyBtn.Disable()
		return
	}
	g.combLabel.SetText(comb)
	g.combCount.SetText(g.tr("selectedCount") + strconv.Itoa(n))
	g.combCopyBtn.Enable()
	if warns := g.conflicts(); len(warns) > 0 {
		g.combWarn.Text = g.tr("warning") + strings.Join(warns, "\n")
	} else {
		g.combWarn.Text = ""
	}
	g.combWarn.Refresh()
}

func (g *gui) copyCombination() {
	comb := g.buildCombination()
	if comb == "" {
		return
	}
	g.win.Clipboard().SetContent(comb)
	g.status.SetText(g.tr("copied") + comb)
	g.status.Refresh()
}

func (g *gui) setLauncher(id string) {
	for i := range g.launchers {
		if g.launchers[i].ID == id {
			g.launcherIdx = i
			break
		}
	}
	g.app.Preferences().SetString("launcher", g.launcher().ID)
	g.updateCombination()
	g.refreshDetail()
}

func (g *gui) toggleFav(idx int) {
	key := g.all[idx].Command
	if g.favs[key] {
		delete(g.favs, key)
	} else {
		g.favs[key] = true
	}
	ids := make([]string, 0, len(g.favs))
	for k := range g.favs {
		ids = append(ids, k)
	}
	sort.Strings(ids)
	g.app.Preferences().SetStringList("favs", ids)
	g.updateFavButton()
	if g.favOnly {
		g.applyFilter()
	}
}

func (g *gui) updateFavButton() {
	if g.current >= 0 && g.favs[g.all[g.current].Command] {
		g.favToggleBtn.Text = "★ " + g.tr("removeFavorite")
	} else {
		g.favToggleBtn.Text = "☆ " + g.tr("addFavorite")
	}
	g.favToggleBtn.Refresh()
}

func (g *gui) applyFavButton() {
	if g.favOnly {
		g.favBtn.Importance = widget.HighImportance
	} else {
		g.favBtn.Importance = widget.MediumImportance
	}
	g.favBtn.Refresh()
}

func (g *gui) clearSelection() {
	g.selected = map[int]bool{}
	g.list.Refresh()
	g.updateCombination()
	g.status.SetText("")
}

func (g *gui) copyCurrent() {
	if g.current < 0 {
		return
	}
	cmd := g.displayCmd(g.all[g.current])
	g.win.Clipboard().SetContent(cmd)
	g.status.SetText(g.tr("copied") + cmd)
	g.status.Refresh()
}

func (g *gui) applyFilter() {
	query := strings.ToLower(strings.TrimSpace(g.search.Text))
	g.filtered = g.filtered[:0]
	for i, c := range g.all {
		if g.catFilterPT != "" && c.Category.PT != g.catFilterPT {
			continue
		}
		if g.favOnly && !g.favs[c.Command] {
			continue
		}
		hay := strings.ToLower(c.Command + " " + c.CommandEN + " " + g.t(c.Title) + " " + g.t(c.Category) + " " + g.t(c.Description) + " " + g.t(c.Compat))
		if query == "" || strings.Contains(hay, query) {
			g.filtered = append(g.filtered, i)
		}
	}
	g.selID = -1
	// Limpa a seleção antes: Select não dispara OnSelected quando o id
	// já estava selecionado, e o detalhe mostraria outro comando.
	g.list.UnselectAll()
	g.list.Refresh()
	if len(g.filtered) > 0 {
		g.list.Select(0)
	} else {
		g.clearDetail()
	}
}

func (g *gui) selectCommand(id widget.ListItemID) {
	if id < 0 || id >= len(g.filtered) {
		g.clearDetail()
		return
	}
	g.current = g.filtered[id]
	g.refreshDetail()
	g.status.SetText("")
	g.copyBtn.Enable()
	if g.copyOnClick.Checked {
		g.copyCurrent()
	}
}

func (g *gui) clearDetail() {
	g.current = -1
	g.refreshDetail()
	g.status.SetText("")
	g.copyBtn.Disable()
}
