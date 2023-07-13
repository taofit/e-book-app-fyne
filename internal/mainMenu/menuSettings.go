package mainMenu

import (
	"fmt"
	"github.com/taofit/e-book-fyne/internal/themes"

	"fyne.io/fyne/v2"
)

func MakeMenu() *fyne.MainMenu {
	hzDefaultTheme := themes.HzDefaultTheme{}
	curApp := fyne.CurrentApp()

	themeMenuItems := []*fyne.MenuItem{
		fyne.NewMenuItem("白天", func() {
			hzDefaultTheme.SetThemeVariant(false)
			curApp.Settings().SetTheme(&hzDefaultTheme)
		}),
		fyne.NewMenuItem("黑夜", func() {
			hzDefaultTheme.SetThemeVariant(true)
			curApp.Settings().SetTheme(&hzDefaultTheme)
		}),
	}
	themeMenu := fyne.NewMenu("主题", themeMenuItems...)

	mainDirectoryMenuItems := []*fyne.MenuItem{
		fyne.NewMenuItem("回总目录",
			func() {
				fmt.Println("going back to main directory")
			}),
	}
	mainDirectory := fyne.NewMenu("目录", mainDirectoryMenuItems...)

	fontSizeMenuItems := []*fyne.MenuItem{
		fyne.NewMenuItem("A-", func() {
			hzDefaultTheme.SetFontSize(-1)
			curApp.Settings().SetTheme(&hzDefaultTheme)
		}),
		fyne.NewMenuItem("A", func() {
			hzDefaultTheme.SetFontSize(0)
			curApp.Settings().SetTheme(&hzDefaultTheme)
		}),
		fyne.NewMenuItem("A+", func() {
			hzDefaultTheme.SetFontSize(1)
			curApp.Settings().SetTheme(&hzDefaultTheme)
		}),
	}
	fontSize := fyne.NewMenu("字体大小", fontSizeMenuItems...)
	main := fyne.NewMainMenu(
		themeMenu,
		mainDirectory,
		fontSize,
	)

	return main
}
