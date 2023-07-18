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

func (section *NavSectionList) MakeNav(
	setArticleWithPag func(article articles.Article, id int, articlesForSubject *[]string),
	setSubList func(listTitle string, list fyne.CanvasObject),
	loadPrevious bool,
) fyne.CanvasObject {
	curApp := fyne.CurrentApp()

	section.list = &widget.List{
		Length: func() int {
			return len(articles.ArticleIndex["rootSubjects"])
		},
		CreateItem: func() fyne.CanvasObject {
			return NewTappableLabel(section, "single articles")
		},
		UpdateItem: func(id widget.ListItemID, item fyne.CanvasObject) {
			subjectName := articles.ArticleIndex["rootSubjects"][id]
			listTitle := articles.Articles[subjectName].Title
			item.(*TappableLabel).SetText(listTitle)
			item.(*TappableLabel).SetListItemID(id)
		},
		OnSelected: func(id widget.ListItemID) {
			subjectName := articles.ArticleIndex["rootSubjects"][id]
			listTitle := articles.Articles[subjectName].Title
			articlesForSubject, ok := articles.ArticleIndex[subjectName]
			if ok {
				subNavSection := NavSectionList{}
				list := subNavSection.generateSubList(setArticleWithPag, setSubList, articlesForSubject)				
				setSubList(listTitle, list)
			} else {
				if a, ok := articles.Articles[subjectName]; ok {
					curApp.Preferences().SetInt(PreferenceCurrentArticle, id)
					setArticleWithPag(a, id, &[]string{})
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

func (section *NavSectionList) generateSubList(
	setArticleWithPag func(article articles.Article, id int, articlesForSubject *[]string),
	setSubList func(listTitle string, list fyne.CanvasObject),
	articlesForSubject []string,
) fyne.CanvasObject {
	curApp := fyne.CurrentApp()

	section.list = &widget.List{
		Length: func() int {
			return len(articlesForSubject)
		},
		CreateItem: func() fyne.CanvasObject {
			return NewTappableLabel(section, "single article")
		},
		UpdateItem: func(id widget.ListItemID, item fyne.CanvasObject) {
			subjectName := articlesForSubject[id]
			title := articles.Articles[subjectName].Title
			item.(*TappableLabel).SetText(title)
			item.(*TappableLabel).SetListItemID(id)
		},
		OnSelected: func(id widget.ListItemID) {
			subjectName := articlesForSubject[id]
			listTitle := articles.Articles[subjectName].Title
			aclsForSubject, ok := articles.ArticleIndex[subjectName]
			if ok {
				subNavSection := NavSectionList{}
				list := subNavSection.generateSubList(setArticleWithPag, setSubList, aclsForSubject)
				setSubList(listTitle, list)
			} else {
				if a, ok := articles.Articles[subjectName]; ok {
					curApp.Preferences().SetString(PreferenceCurrentArticle, subjectName)
					setArticleWithPag(a, id, &articlesForSubject)
				}
			}
		},
	}

	return section.list
}
