package domain

import (
	"github.com/robfig/revel"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func NewRecaptcha() *Recaptcha {
	private, _ := os.Getenv("RECAPTCHA_PRIVATE")
	public, _ := revel.Config.String("recaptcha.public")
	verify, _ := revel.Config.String("recaptcha.verify")

	return &Recaptcha{public: public, private: private, verify: verify}
}

type InvalidCaptcha string

type Recaptcha struct {
	public  string
	private string
	verify  string
}

func (i InvalidCaptcha) Error() string {
	return string(i)
}

func (r *Recaptcha) Verify(remote_ip string, challenge string, response string) (bool, error) {
	resp, err := http.PostForm(r.verify, url.Values{
		"privatekey": {r.private},
		"remoteip":   {remote_ip},
		"challenge":  {challenge},
		"response":   {response},
	})
	if err != nil {
		return false, err
	} else {
		result, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return false, err
		}
		values := strings.Fields(string(result))
		if values[0] == "false" {
			return false, InvalidCaptcha(values[1])
		}
		return true, nil
	}
}
