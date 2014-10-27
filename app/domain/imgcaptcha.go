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

const CAPTCHA_COLS = 5
const CAPTCHA_ROWS = 2
const CAPTCHA_IMAGES = CAPTCHA_COLS * CAPTCHA_ROWS
const CAPTCHA_IMG_SIZE = 100
const CAPTCHA_WIDTH = CAPTCHA_COLS * CAPTCHA_IMG_SIZE
const CAPTCHA_HEIGHT = CAPTCHA_ROWS * CAPTCHA_IMG_SIZE

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

	images := make([]ImageItem, 0, 18)
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
	start := random(0, len(keys)-CAPTCHA_IMAGES)

	keys = keys[start : start+CAPTCHA_IMAGES]

	var title string
	var index int
	images := ""
	index = random(0, CAPTCHA_IMAGES)
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
	img := image.NewRGBA(image.Rect(0, 0, CAPTCHA_WIDTH, CAPTCHA_HEIGHT))
	for i, key := range strings.Split(c.Images, ",") {
		pos := i
		py := 0
		if pos >= CAPTCHA_COLS {
			py += CAPTCHA_IMG_SIZE
			pos -= CAPTCHA_COLS
		}
		px := pos * CAPTCHA_IMG_SIZE
		rec := image.Rect(px, py, px+CAPTCHA_IMG_SIZE, py+CAPTCHA_IMG_SIZE)
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
	if pos >= CAPTCHA_COLS {
		py += CAPTCHA_IMG_SIZE
		pos -= CAPTCHA_COLS
	}
	px := pos * CAPTCHA_IMG_SIZE

	if x < px || x > (px+CAPTCHA_IMG_SIZE) {
		return false
	}
	if y < py || y > (py+CAPTCHA_IMG_SIZE) {
		return false
	}
	return true
}
