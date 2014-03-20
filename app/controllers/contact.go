package controllers

import (
	"github.com/rchargel/goblog/app/domain"
	"github.com/robfig/revel"
)

type Contact struct {
	*revel.Controller
}

func (c Contact) Index() revel.Result {
	captchaKey, found := revel.Config.String("recaptcha.public")
	if !found {
		return c.NotFound("No Contact")
	}

	recaptcha_error, _ := c.Flash.Data["recaptcha_error"]
	return c.Render(captchaKey, recaptcha_error)
}

func (c Contact) Send(name, email, message, recaptcha_challenge_field, recaptcha_response_field string) revel.Result {
	success := true
	c.Validation.Required(name).Message("Please let me know your name")
	c.Validation.Required(message).Message("Don't forget to include a message")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		success = false
	}

	remote_ip := c.Request.RemoteAddr

	recaptcha := domain.NewRecaptcha()
	valid, err := recaptcha.Verify(remote_ip, recaptcha_challenge_field, recaptcha_response_field)
	if !valid && err != nil {
		c.Flash.Out["recaptcha_error"] = err.Error()
		success = false
	}
	if !success {
		c.FlashParams()
		return c.Redirect(Contact.Index)
	} else {
		return c.Render()
	}
}
