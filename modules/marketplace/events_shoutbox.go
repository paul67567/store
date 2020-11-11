package marketplace

import (
	"fmt"
	"strings"

	"github.com/paul67567/store/tree/master/modules/apis"
)

func EventNewShoutboxPost(user User, message Message) {
	var (
		marketUrl = MARKETPLACE_SETTINGS.SiteURL
		text      = fmt.Sprintf("[@%s](%s/user/%s) has posted on shoutbox:\n> %s",
			user.Username, marketUrl, user.Username,
			strings.Replace(message.Text, "\n", "\n > ", -1), //------------|
		)
	)

	go apis.PostMattermostEvent(MARKETPLACE_SETTINGS.MattermostIncomingHookShoutbox, text)
}
