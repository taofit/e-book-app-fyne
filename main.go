package main

import (
	"github.com/taofit/e-book-fyne/internal/articles"
	"github.com/taofit/e-book-fyne/internal/mainMenu"
	"github.com/taofit/e-book-fyne/internal/navList"
	"github.com/taofit/e-book-fyne/internal/themes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	hzApp := app.NewWithID("io.fyne.hz")
	hzTheme := themes.HzDefaultTheme{}
	hzApp.Settings().SetTheme(&hzTheme)

	appTitle := "eBookCollection"
	w := hzApp.NewWindow(appTitle)
	topWindow := w
	articles.PopulateArticles()
	w.SetMainMenu(mainMenu.MakeMenu())
	w.SetMaster()

	navSection := navList.NavSectionList{}

	content := container.NewMax()
	title := widget.NewLabel("Article title")
	// intro := widget.NewLabel("Introduction goes here")
	// intro.Wrapping = fyne.TextWrapWord

	setArticle := func(a articles.Article) {
		if fyne.CurrentDevice().IsMobile() {
			child := hzApp.NewWindow(a.Title)
			topWindow = child
			child.SetContent(a.LoadFile(topWindow))
			child.Show()
			child.SetOnClosed(func() {
				topWindow = w
			})
			return
		}
		title.SetText(a.Title)
		// intro.SetText(a.Intro)

		content.Objects = []fyne.CanvasObject{a.LoadFile(w)}
		content.Refresh()
	}

	setSubList := func(listTitle string, list fyne.CanvasObject) {
		if fyne.CurrentDevice().IsMobile() {
			child := hzApp.NewWindow(listTitle)
			topWindow = child
			child.SetContent(list)
			child.Show()
			child.SetOnClosed(func() {
				topWindow = w
			})
			return
		}
		title.SetText(listTitle)
		// intro.SetText("brief introduction goes here")

		content.Objects = []fyne.CanvasObject{list}
		content.Refresh()
	}

	article := container.NewBorder(
		container.NewVBox(title, widget.NewSeparator()), nil, nil, nil, content)

	if fyne.CurrentDevice().IsMobile() {
		topBar := makeTopBar(appTitle)
		navSection := navSection.MakeNav(setArticle, setSubList, false)
		content := container.NewBorder(topBar, nil, nil, nil, navSection)
		w.SetContent(content)
	} else {
		split := container.NewHSplit(navSection.MakeNav(setArticle, setSubList, true), article)
		split.Offset = 0.2
		w.SetContent(split)
	}

	w.Resize(fyne.Size{Width: 500, Height: 500})
	w.ShowAndRun()
}

func makeTopBar(titleName string) *widget.Label {
	topBar := widget.NewLabel(titleName)
	topBar.Alignment = fyne.TextAlignCenter

	return topBar
}
