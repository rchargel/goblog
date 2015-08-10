package domain

import "testing"

func Test_ReadBookmarks(t *testing.T) {
	filepath := "../../public/resources/bookmarks.yml"
	bookmarks, err := readBookmarks(filepath)
	if err == nil {
		if len(bookmarks.Bookmarks) == 0 {
			t.Error("There are no values in the bookmarks")
		}
	} else {
		t.Error(err)
	}
}
