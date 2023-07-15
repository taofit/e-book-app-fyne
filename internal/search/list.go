package search

import (
	"github.com/taofit/e-book-fyne/internal/articles"
	"github.com/taofit/e-book-fyne/internal/navList"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeResultList(
	input string,
	searchKeyword func(word string) []resultItem,
	setArticle func(article articles.Article)) fyne.CanvasObject {
	resultSlice := searchKeyword(input)
	curApp := fyne.CurrentApp()

	list := &widget.List{
		Length: func() int {
			return len(resultSlice)
		},
		CreateItem: func() fyne.CanvasObject {
			return container.NewVBox(widget.NewLabel("title"), widget.NewRichTextWithText("result text"))
		},
		UpdateItem: func(id widget.ListItemID, item fyne.CanvasObject) {
			articleKey := resultSlice[id].key
			text := resultSlice[id].text
			articleTitle := "<<" + articles.Articles[articleKey].Title + ">>"

			item.(*fyne.Container).Objects[0].(*widget.Label).SetText(articleTitle)
			item.(*fyne.Container).Objects[1].(*widget.RichText).Segments[0] = &widget.TextSegment{
				Text: text,
			}
			item.(*fyne.Container).Objects[1].(*widget.RichText).Wrapping = fyne.TextWrapWord
			item.(*fyne.Container).Objects[1].Refresh()
		},
		OnSelected: func(id widget.ListItemID) {
			articleKey := resultSlice[id].key
			if a, ok := articles.Articles[articleKey]; ok {
				curApp.Preferences().SetInt(navList.PreferenceCurrentArticle, id)
				setArticle(a)
			}
		},
	}

	if list.Length() == 0 {
		resultLabel := widget.NewLabel("no result")
		resultLabel.Alignment = fyne.TextAlignCenter

		return resultLabel
	}

	return list
}
