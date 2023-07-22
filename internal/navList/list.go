package navList

import (
	"github.com/taofit/e-book-fyne/internal/articles"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

const PreferenceCurrentArticle = "currentArticle"

type NavSectionList struct {
	list *widget.List
}
type TappableLabel struct {
	widget.Label
	listItemID widget.ListItemID
	parent     *NavSectionList
}

func NewTappableLabel(parent *NavSectionList, text string) *TappableLabel {
	tpbLabel := &TappableLabel{
		Label:      widget.Label{Text: text},
		parent:     parent,
		listItemID: -1,
	}
	tpbLabel.ExtendBaseWidget(tpbLabel)

	return tpbLabel
}

func (t *TappableLabel) SetListItemID(id widget.ListItemID) {
	t.listItemID = id
}

func (t *TappableLabel) Tapped(pe *fyne.PointEvent) {
	t.parent.list.Unselect(t.listItemID) //need to unselect the listItemID first, as if it is selected, it won't run the code in Select func
	t.parent.list.Select(t.listItemID)
}

func (section *NavSectionList) MakeMblNav(
	setArticle func(article articles.Article),
	setSubList func(listTitle string, list fyne.CanvasObject),
	articlesForSubject []string,
	loadPrevious bool,
) fyne.CanvasObject {
	curApp := fyne.CurrentApp()

	section.list = &widget.List{
		Length: func() int {
			return len(articlesForSubject)
		},
		CreateItem: func() fyne.CanvasObject {
			return NewTappableLabel(section, "single articles")
		},
		UpdateItem: func(id widget.ListItemID, item fyne.CanvasObject) {
			subjectName := articlesForSubject[id]
			listTitle := articles.Articles[subjectName].Title
			item.(*TappableLabel).SetText(listTitle)
			item.(*TappableLabel).Wrapping = fyne.TextWrapBreak
			item.(*TappableLabel).SetListItemID(id)
			section.list.SetItemHeight(id, item.MinSize().Height)
		},
		OnSelected: func(id widget.ListItemID) {
			subjectName := articlesForSubject[id]
			listTitle := articles.Articles[subjectName].ShortenTitle()
			articlesForSubject, ok := articles.ArticleIndex[subjectName]
			if ok {
				subNavSection := NavSectionList{}
				list := subNavSection.MakeMblNav(setArticle, setSubList, articlesForSubject, false)
				setSubList(listTitle, list)
			} else {
				if a, ok := articles.Articles[subjectName]; ok {
					curApp.Preferences().SetInt(PreferenceCurrentArticle, id)
					setArticle(a)
				}
			}
		},
	}

	if loadPrevious {
		currentPref := curApp.Preferences().IntWithFallback(PreferenceCurrentArticle, 0)
		section.list.Select(currentPref)
	}

	return section.list
}
