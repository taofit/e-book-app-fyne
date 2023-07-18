package mainMenu

import (
	"github.com/taofit/e-book-fyne/internal/themes"

	"fyne.io/fyne/v2"
)

func MakeMenu() *fyne.MainMenu {
	eBDefaultTheme := themes.EBookDefaultTheme{}
	curApp := fyne.CurrentApp()

	themeMenuItems := []*fyne.MenuItem{
		fyne.NewMenuItem("light", func() {
			eBDefaultTheme.SetThemeVariant(false)
			curApp.Settings().SetTheme(&eBDefaultTheme)
		}),
		fyne.NewMenuItem("dark", func() {
			eBDefaultTheme.SetThemeVariant(true)
			curApp.Settings().SetTheme(&eBDefaultTheme)
		}),
	}
	themeMenu := fyne.NewMenu("theme", themeMenuItems...)

	fontSizeMenuItems := []*fyne.MenuItem{
		fyne.NewMenuItem("A-", func() {
			eBDefaultTheme.SetFontSize(-1)
			curApp.Settings().SetTheme(&eBDefaultTheme)
		}),
		fyne.NewMenuItem("A", func() {
			eBDefaultTheme.SetFontSize(0)
			curApp.Settings().SetTheme(&eBDefaultTheme)
		}),
		fyne.NewMenuItem("A+", func() {
			eBDefaultTheme.SetFontSize(1)
			curApp.Settings().SetTheme(&eBDefaultTheme)
		}),
	}
	fontSize := fyne.NewMenu("font size", fontSizeMenuItems...)
	main := fyne.NewMainMenu(
		themeMenu,
		fontSize,
	)

	return main
}
