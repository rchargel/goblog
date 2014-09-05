package domain

import (
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	"image"
	"image/draw"
	"image/jpeg"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

type ImageItem struct {
	Title string
	Image string
}

type TitledImage struct {
	Image image.Image
	Title string
}

type ImageCaptcha struct {
	Images string
	Title  string
	Index  int
}

var imageMap map[string]image.Image = newImageCaptcha()

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func keys() []string {
	set := make([]string, len(imageMap), len(imageMap))
	i := 0
	for k, _ := range imageMap {
		set[i] = k
		i += 1
	}

	p := rand.Perm(len(set))

	dest := make([]string, len(set))

	for i, v := range p {
		dest[v] = set[i]
	}
	return dest
}

func readImages(images []ImageItem, ch chan TitledImage) {
	for _, img := range images {
		file, _ := os.Open("./public/img/captcha/" + img.Image)
		defer file.Close()

		image, _ := jpeg.Decode(file)
		ch <- TitledImage{Image: image, Title: img.Title}
	}
	close(ch)
}

func putImages(i map[string]image.Image, ch chan TitledImage) {
	for img := range ch {
		i[img.Title] = img.Image
	}
}

func initImages(images []ImageItem, i map[string]image.Image) {
	ch := make(chan TitledImage)

	go readImages(images, ch)
	putImages(i, ch)
}

func newImageCaptcha() map[string]image.Image {
	file, _ := os.Open("./public/resources/images.json")
	defer file.Close()

	decoder := json.NewDecoder(file)

	images := make([]ImageItem, 0, 12)
	for {
		if err := decoder.Decode(&images); err == nil {
			// resume will get set
		} else if err == io.EOF {
			break
		} else {
			panic(err)
		}
	}

	i := make(map[string]image.Image)
	initImages(images, i)
	return i
}

func NewCaptcha() ImageCaptcha {
	rand.Seed(time.Now().Unix())
	keys := keys()
	start := random(0, len(keys)-6)

	keys = keys[start : start+6]

	var title string
	var index int
	images := ""
	index = random(0, 6)
	for i, key := range keys {
		if i == index {
			title = key
		}
		if i == 0 {
			images += key
		} else {
			images += fmt.Sprintf(",%v", key)
		}
	}

	return ImageCaptcha{Images: images, Title: title, Index: index}
}

func (c ImageCaptcha) Output(w io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, 300, 200))
	for i, key := range strings.Split(c.Images, ",") {
		pos := i
		py := 0
		if pos >= 3 {
			py = 100
			pos -= 3
		}
		px := pos * 100
		rec := image.Rect(px, py, px+100, py+100)
		draw.Draw(img, rec, imageMap[key], image.ZP, draw.Src)
	}
	jpeg.Encode(w, img, &jpeg.Options{85})
}

func (c ImageCaptcha) Apply(req *revel.Request, resp *revel.Response) {
	resp.Out.Header().Set("Content-Type", "image/jpeg")
	c.Output(resp.Out)
}

func (c ImageCaptcha) IsCorrect(x, y int) bool {
	pos := c.Index
	py := 0
	if pos >= 3 {
		py = 100
		pos -= 3
	}
	px := pos * 100

	if x < px || x > (px+100) {
		return false
	}
	if y < py || y > (py+100) {
		return false
	}
	return true
}
