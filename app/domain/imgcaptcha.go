package domain

import (
	"encoding/json"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/nfnt/resize"
	"github.com/revel/revel"
)

const captchaCols = 5
const captchaRows = 2
const captchaImages = captchaCols * captchaRows
const captchaImgSize = 80
const captchaWidth = captchaCols * captchaImgSize
const captchaHeight = captchaRows * captchaImgSize

// ImageItem is a single image item
type ImageItem struct {
	Title string
	Image string
}

// TitledImage is the image with a title.
type TitledImage struct {
	Image image.Image
	Title string
}

// ImageCaptcha is the series of images and titles with a selected index.
type ImageCaptcha struct {
	Images string
	Title  string
	Index  int
}

var imageMap map[string]image.Image

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getImageMap() map[string]image.Image {
	if imageMap == nil {
		imageMap = newImageCaptcha()
	}
	return imageMap
}

func keys() []string {
	set := make([]string, len(getImageMap()), len(getImageMap()))
	i := 0
	for k := range getImageMap() {
		set[i] = k
		i++
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
		m := resize.Resize(captchaImgSize, captchaImgSize, image, resize.Bilinear)
		ch <- TitledImage{Image: m, Title: img.Title}
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

// NewCaptcha creates a new ImageCaptcha
func NewCaptcha() ImageCaptcha {
	rand.Seed(time.Now().Unix())
	keys := keys()
	start := random(0, len(keys)-captchaImages)

	keys = keys[start : start+captchaImages]

	var title string
	var index int
	images := ""
	index = random(0, captchaImages)
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

// Output sends the contents of the image captcha to the writer.
func (c ImageCaptcha) Output(w io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, captchaWidth, captchaHeight))
	for i, key := range strings.Split(c.Images, ",") {
		pos := i
		py := 0
		if pos >= captchaCols {
			py += captchaImgSize
			pos -= captchaCols
		}
		px := pos * captchaImgSize
		rec := image.Rect(px, py, px+captchaImgSize, py+captchaImgSize)
		draw.Draw(img, rec, getImageMap()[key], image.ZP, draw.Src)
	}
	jpeg.Encode(w, img, &jpeg.Options{Quality: 85})
}

// Apply is the controller function.
func (c ImageCaptcha) Apply(req *revel.Request, resp *revel.Response) {
	resp.Out.Header().Set("Content-Type", "image/jpeg")
	c.Output(resp.Out)
}

// IsCorrect determines if the correct image was selected.
func (c ImageCaptcha) IsCorrect(x, y int) bool {
	pos := c.Index
	py := 0
	if pos >= captchaCols {
		py += captchaImgSize
		pos -= captchaCols
	}
	px := pos * captchaImgSize

	if x < px || x > (px+captchaImgSize) {
		return false
	}
	if y < py || y > (py+captchaImgSize) {
		return false
	}
	return true
}
