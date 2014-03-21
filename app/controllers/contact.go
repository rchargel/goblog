package controllers

import (
	"fmt"
	"github.com/rchargel/goblog/app/domain"
	"github.com/robfig/revel"
	"strconv"
)

type Contact struct {
	*revel.Controller
}

func (c Contact) getCaptcha() domain.ImageCaptcha {
	title := c.Session["captcha_title"]
	images := c.Session["captcha_images"]
	index, _ := strconv.ParseInt(c.Session["captcha_index"], 10, 32)
	return domain.ImageCaptcha{Title: title, Images: images, Index: int(index)}
}

func (c Contact) Index() revel.Result {
	imageCaptcha := domain.NewCaptcha()
	c.Session["captcha_title"] = imageCaptcha.Title
	c.Session["captcha_images"] = imageCaptcha.Images
	c.Session["captcha_index"] = fmt.Sprintf("%v", imageCaptcha.Index)

	captcha_title := imageCaptcha.Title

	return c.Render(captcha_title)
}

func (c Contact) RenderImg() revel.Result {
	imageCaptcha := c.getCaptcha()
	return imageCaptcha
}

func (c Contact) Send(name, email, message string, x, y int) revel.Result {
	success := true
	c.Validation.Required(name).Message("Please let me know your name")
	c.Validation.Required(message).Message("Don't forget to include a message")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		success = false
	}

	captcha := c.getCaptcha()
	if !captcha.IsCorrect(x, y) {
		c.Flash.Out["captcha_error"] = "Please click on the correct image"
		success = false
	}

	if !success {
		c.FlashParams()
		return c.Redirect(Contact.Index)
	} else {
		return c.Render()
	}
}
