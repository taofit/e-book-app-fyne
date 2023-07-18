package pagination

import (
	"github.com/taofit/e-book-fyne/internal/articles"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func MakeBottomPag(
	childContainer fyne.Window,
	a articles.Article,
	id int,
	numOfAcls int,
	articlesForSubject *[]string,
	setArticleWithPag func(a articles.Article, id int, articlesForSubject *[]string),
) *fyne.Container {
	nextId := id + 1
	prevId := id - 1
	var nextAcl articles.Article
	var prevAcl articles.Article

	if nextId < numOfAcls {
		nextSubjectName := (*articlesForSubject)[nextId]
		if _, ok := articles.Articles[nextSubjectName]; ok {
			nextAcl = articles.Articles[nextSubjectName]
		}
	}
	if prevId >= 0 {
		prevSubjectName := (*articlesForSubject)[prevId]
		if _, ok := articles.Articles[prevSubjectName]; ok {
			prevAcl = articles.Articles[prevSubjectName]
		}
	}

	nextButton := widget.NewButtonWithIcon("", theme.NavigateNextIcon(), func() {
		setArticleWithPag(nextAcl, nextId, articlesForSubject)
		go childContainer.Close()
	})

	if nextAcl == (articles.Article{}) {
		nextButton.Disable()
	}

	prevButton := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		setArticleWithPag(prevAcl, prevId, articlesForSubject)
		go childContainer.Close()
	})
	if prevAcl == (articles.Article{}) {
		prevButton.Disable()
	}

	return container.NewHBox(prevButton, nextButton)
}
