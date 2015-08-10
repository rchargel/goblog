package domain

import "testing"

func Test_SlideShow_Init(t *testing.T) {
	slideShowMap := &SlideShowMap{}
	slideShowMap.init("../../public/resources/slideshow.yml")
	if len(slideShowMap.Items) == 0 {
		t.Error("There were no items in the slideshow")
	}
}
