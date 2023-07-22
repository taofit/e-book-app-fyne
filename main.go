package main

import (
	"fmt"

	"github.com/taofit/e-book-fyne/internal/articles"
	"github.com/taofit/e-book-fyne/internal/mainMenu"
	"github.com/taofit/e-book-fyne/internal/navList"
	"github.com/taofit/e-book-fyne/internal/navTree"
	"github.com/taofit/e-book-fyne/internal/search"
	"github.com/taofit/e-book-fyne/internal/themes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	eBookApp := app.NewWithID("io.fyne.ebook")
	eBTheme := themes.EBookDefaultTheme{}
	eBookApp.Settings().SetTheme(&eBTheme)

	appTitle := "E-book Collection"
	w := eBookApp.NewWindow(appTitle)
	topWindow := w
	articles.PopulateArticles()
	w.SetMainMenu(mainMenu.MakeMenu())
	w.SetMaster()

	navSection := navList.NavSectionList{}

	content := container.NewMax()
	title := widget.NewLabel("Article title")

	setArticle := func(a articles.Article) {
		if fyne.CurrentDevice().IsMobile() {
			child := eBookApp.NewWindow(a.ShortenTitle())
			topWindow = child
			child.SetContent(a.LoadFile(topWindow))
			child.Show()
			child.SetOnClosed(func() {
				topWindow = w
			})
			return
		}
		title.SetText(a.Title)

		content.Objects = []fyne.CanvasObject{a.LoadFile(w)}
		content.Refresh()
	}

	setSubList := func(listTitle string, list fyne.CanvasObject) {
		if fyne.CurrentDevice().IsMobile() {
			child := eBookApp.NewWindow(listTitle)
			topWindow = child
			child.SetContent(list)
			child.Show()
			child.SetOnClosed(func() {
				topWindow = w
			})
			return
		}
		title.SetText(listTitle)

		content.Objects = []fyne.CanvasObject{list}
		content.Refresh()
	}

	setSearchResult := func(resultCnt fyne.CanvasObject, input string) {
		searchTitle := fmt.Sprintf("search \" %s \" result", input)
		if fyne.CurrentDevice().IsMobile() {
			child := eBookApp.NewWindow(searchTitle)
			topWindow = child
			child.SetContent(resultCnt)
			child.Show()
			child.SetOnClosed(func() {
				topWindow = w
			})
			return
		}
		title.SetText(searchTitle)

		content.Objects = []fyne.CanvasObject{resultCnt}
		content.Refresh()
	}

	searchSection := search.MakeSearchEntry(setSearchResult, setArticle)
	rootSubjects := articles.ArticleIndex[articles.RootSubjectsKey]
	if fyne.CurrentDevice().IsMobile() {
		topBar := makeTopBar(appTitle)
		navSection := navSection.MakeMblNav(setArticle, setSubList, rootSubjects, false)
		content := container.NewBorder(topBar, searchSection, nil, nil, navSection)
		w.SetContent(content)
	} else {
		article := container.NewBorder(
			container.NewVBox(title, widget.NewSeparator()), nil, nil, nil, content)
		split := container.NewHSplit(
			navTree.MakeNav(setArticle, true),
			article,
		)
		split.Offset = 0.2
		content := container.NewBorder(nil, searchSection, nil, nil, split)
		w.SetContent(content)
	}

	w.Resize(fyne.Size{Width: 500, Height: 500})
	w.ShowAndRun()
}

func makeTopBar(titleName string) *widget.Label {
	topBar := widget.NewLabel(titleName)
	topBar.Alignment = fyne.TextAlignCenter

	return topBar
}
