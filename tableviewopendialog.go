package main

import (
	"fmt"
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type myMainWindow struct {
	*walk.MainWindow
	searchResultTableView      *walk.TableView
	searchResultTableViewModel *SearchResultTableViewModel
}

func main() {
	mw := &myMainWindow{searchResultTableViewModel: NewSearchResultTableViewModel()}

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "QQ:583081713 mabiao",
		Size:     Size{Width: 600, Height: 300},
		Layout:   VBox{},
		Children: []Widget{
			TableView{
				AssignTo:              &mw.searchResultTableView,
				AlternatingRowBGColor: walk.RGB(240, 240, 240),
				Columns: []TableViewColumn{
					{Title: "Index", Name: "Index", Width: 50},
					{Title: "Word", Name: "Word", Format: "%s", Width: 260},
					{Title: "More", Name: "ResultMore", Format: "%s", Width: 100},
				},
				Model:           mw.searchResultTableViewModel,
				OnItemActivated: mw.OnItemActivated_SearchResultTableView, // Double click to open the dialog
			},
		},
	}.Create()); err != nil {
		log.Fatal("create windows err：", err)
	}

	// load ico
	if icon, err := walk.NewIconFromResourceId(3); err == nil {
		mw.SetIcon(icon)
	}

	mw.Run()
}

type SearchResultTableViewModel struct {
	walk.SortedReflectTableModelBase
	items []*SearchResult
}

type SearchResult struct {
	Index      int     // index
	Word       string  // word
	ResultMore string  //
	KP         []*Line //
}

func NewSearchResultTableViewModel() *SearchResultTableViewModel {
	results := []*SearchResult{
		&SearchResult{0, "number1", "1, 2, 5", []*Line{
			&Line{Index: 1, Title: "num1", IsHave: true},
			&Line{Index: 2, Title: "num1", IsHave: true},
			&Line{Index: 3, Title: "num1", IsHave: false},
			&Line{Index: 4, Title: "num1", IsHave: false},
			&Line{Index: 5, Title: "num1", IsHave: true},
		}},
		&SearchResult{1, "number2", "1, 3, 4, 5", []*Line{
			&Line{Index: 1, Title: "num2", IsHave: true},
			&Line{Index: 2, Title: "num2", IsHave: false},
			&Line{Index: 3, Title: "num2", IsHave: true},
			&Line{Index: 4, Title: "num2", IsHave: false},
			&Line{Index: 5, Title: "num2", IsHave: true},
		}},
		&SearchResult{2, "number3", "1, 8", []*Line{
			&Line{Index: 1, Title: "num3", IsHave: true},
			&Line{Index: 2, Title: "num3", IsHave: false},
			&Line{Index: 3, Title: "num31", IsHave: false},
			&Line{Index: 4, Title: "num3", IsHave: false},
			&Line{Index: 5, Title: "num3", IsHave: false},
			&Line{Index: 6, Title: "num3", IsHave: false},
			&Line{Index: 7, Title: "num3", IsHave: false},
			&Line{Index: 8, Title: "num3", IsHave: true},
		}},
	}
	sr := new(SearchResultTableViewModel)
	sr.items = results
	return sr
}

func (sr *SearchResultTableViewModel) Items() interface{} {
	return sr.items
}

type LineTableViewModel struct {
	walk.SortedReflectTableModelBase
	items []*Line
}

type Line struct {
	Index  int    // idx
	Title  string // title
	IsHave bool   // is have
}

func NewLineTableViewModel() *LineTableViewModel {
	return new(LineTableViewModel)
}

func (kp *LineTableViewModel) Items() interface{} {
	return kp.items
}

// double click
func (mw *myMainWindow) OnItemActivated_SearchResultTableView() {
	if index := mw.searchResultTableView.CurrentIndex(); index > -1 {
		info := mw.searchResultTableViewModel.items[index]
		if cmd, err := RunTableViewDialog(mw, info); err != nil {
			log.Print(err)
		} else if cmd == walk.DlgCmdOK {
			fmt.Println("DlgOk")
		}
	}
}

func RunTableViewDialog(owner walk.Form, sr *SearchResult) (int, error) {
	var dlg *walk.Dialog
	kp := NewLineTableViewModel()
	kp.items = sr.KP
	return Dialog{
		AssignTo: &dlg,
		Title:    sr.Word,
		MinSize:  Size{450, 250},
		Layout:   VBox{},
		Children: []Widget{
			TableView{
				AlternatingRowBGColor: walk.RGB(240, 240, 240), // 行的颜色
				Columns: []TableViewColumn{
					{Title: "Number", Name: "Index", Width: 80},
					{Title: "Title", Name: "Title", Format: "%s", Width: 220},
					{Title: "IsHave", Name: "IsHave", Format: "%T", Width: 80},
				},
				Model: kp,
			},
		},
	}.Run(owner)
}
