package marketplace

import (
	"github.com/dchest/captcha"
	"github.com/gocraft/web"
	"github.com/mojocn/base64Captcha"

	"github.com/paul67567/store/modules/util"
)

func (c *Context) ViewListPrivateMessages(w web.ResponseWriter, r *web.Request) {
	if len(r.URL.Query()["section"]) > 0 {
		section := r.URL.Query()["section"][0]
		c.SelectedSection = section
	} else {
		c.SelectedSection = ""
	}
	if c.SelectedSection == "inbox" {
		vps := []ViewPrivateThread{}
		for _, vp := range c.ViewPrivateThreads {
			if !vp.IsRead {
				vps = append(vps, vp)
			}
		}
		c.ViewPrivateThreads = vps
	}

	util.RenderTemplateOrAPIResponse(w, r, "board/messages", c, c.IsAPIRequest)

}

func (c *Context) ViewShowPrivateMessageGET(w web.ResponseWriter, r *web.Request) {
	c.CaptchaId = captcha.New()
	c.SelectedSection = c.ViewThread.Uuid

	for _, m := range c.ViewThread.Messages {
		if m.RecieverUserUuid == c.ViewUser.Uuid && !m.IsReadByReciever {
			m.IsReadByReciever = !m.IsReadByReciever
			m.Save()
		}
	}
	UpdateThreadPerusalStatus(c.ViewThread.Uuid, c.ViewUser.Uuid)

	// hack to make displayed thread viewd as read
	for i := range c.ViewThreads {
		th := c.ViewThreads[i]
		if th.Uuid == c.ViewThread.Uuid {
			c.ViewThreads[i].IsRead = true
			break
		}
	}

	util.RenderTemplateOrAPIResponse(w, r, "board/message", c, c.IsAPIRequest)
}

func (c *Context) ViewShowPrivateMessagePOST(w web.ResponseWriter, r *web.Request) {
	isCaptchaValid := base64Captcha.VerifyCaptcha(r.FormValue("captcha_id"), r.FormValue("captcha")) || c.ViewUser.IsAdmin || c.ViewUser.IsStaff || c.IsAPIRequest

	if !isCaptchaValid {
		c.Error = "Invalid captcha"
		c.ViewShowPrivateMessageGET(w, r)
		return
	}

	message, err := CreateMessage(r.FormValue("text"), c.Thread, *c.ViewUser.User)
	if err != nil {
		c.Error = err.Error()
		c.ViewShowPrivateMessageGET(w, r)
		return
	}

	err = message.AddImage(r)
	if err != nil {
		c.Error = err.Error()
		c.ViewShowPrivateMessageGET(w, r)
		return
	}
	CacheClearPrivateThreads(c.ViewUser.Uuid)
	c.MessagesMiddleware(w, r, c.ViewShowPrivateMessageGET)
}
