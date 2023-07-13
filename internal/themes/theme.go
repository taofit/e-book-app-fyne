//go:generate fyne bundle -o bundled.go GochiHand.ttf
//go:generate fyne bundle -append -o bundled.go Icon.png

package themes

import (
	"image/color"

	"github.com/taofit/e-book-fyne/internal/bundled"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type HzDefaultTheme struct {
	useDark  bool
	fontSize float32
}

var _ fyne.Theme = (*HzDefaultTheme)(nil)

func (m *HzDefaultTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNameBackground, theme.ColorNameInputBackground,
		theme.ColorNameOverlayBackground, theme.ColorNameMenuBackground:
		if m.useDark || v == theme.VariantDark {
			return &color.NRGBA{R: 0x0C, G: 0x02, B: 0x09, A: 0xFF}
		}
		return &color.NRGBA{R: 0xFF, G: 0xF8, B: 0xDC, A: 0xFF}
	case theme.ColorNameForeground:
		if m.useDark || v == theme.VariantDark {
			return &color.NRGBA{R: 0xFF, G: 0xF8, B: 0xDC, A: 0xFF}
		}
		return &color.NRGBA{R: 0x0C, G: 0x02, B: 0x09, A: 0xFF}
	case theme.ColorNamePrimary:
		return &color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xAA}
	case theme.ColorNameButton, theme.ColorNameFocus, theme.ColorNameSelection:
		return &color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x66}
	}

	return theme.DefaultTheme().Color(n, v)
}

func (m *HzDefaultTheme) Font(s fyne.TextStyle) fyne.Resource {
	return bundled.ResourceFZKaiZ03RegularTtf
}

func (m *HzDefaultTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (m *HzDefaultTheme) Size(n fyne.ThemeSizeName) float32 {
	switch n {
	case theme.SizeNameLineSpacing:
		return 2
	case theme.SizeNameText:
		return theme.DefaultTheme().Size(n) + 5 + m.fontSize
	}

	return theme.DefaultTheme().Size(n)
}

func (m *HzDefaultTheme) SetThemeVariant(useDark bool) {
	m.useDark = useDark
}

func (m *HzDefaultTheme) SetFontSize(size float32) {
	m.fontSize = size
}
