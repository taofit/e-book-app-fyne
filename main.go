package main

import (
	"fmt"

	"github.com/taofit/e-book-fyne/internal/articles"
	"github.com/taofit/e-book-fyne/internal/mainMenu"
	"github.com/taofit/e-book-fyne/internal/navList"
	"github.com/taofit/e-book-fyne/internal/pagination"
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

	appTitle := "eBookCollection"
	w := eBookApp.NewWindow(appTitle)
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
			child := eBookApp.NewWindow(a.Title)
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

	var setArticleWithPag func(a articles.Article, id int, articlesForSubject *[]string)
	setArticleWithPag = func(a articles.Article, id int, articlesForSubject *[]string) {
		if fyne.CurrentDevice().IsMobile() {
			child := eBookApp.NewWindow(a.Title)
			topWindow = child
			numOfAcls := len(*articlesForSubject)
			childContent := container.NewBorder(
				nil,
				pagination.MakeBottomPag(child, a, id, numOfAcls, articlesForSubject, setArticleWithPag),
				nil,
				nil,
				a.LoadFile(topWindow),
			)
			child.SetContent(childContent)
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
		// intro.SetText("brief introduction goes here")

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

	article := container.NewBorder(
		container.NewVBox(title, widget.NewSeparator()), nil, nil, nil, content)

	searchSection := search.MakeSearchEntry(setSearchResult, setArticle)
	if fyne.CurrentDevice().IsMobile() {
		topBar := makeTopBar(appTitle)
		navSection := navSection.MakeNav(setArticleWithPag, setSubList, false)
		content := container.NewBorder(topBar, searchSection, nil, nil, navSection)
		w.SetContent(content)
	} else {
		split := container.NewHSplit(navSection.MakeNav(setArticleWithPag, setSubList, true), article)
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
