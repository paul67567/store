package marketplace

import (
	"fmt"

	"github.com/paul67567/store/modules/apis"
)

func EventUserLoggedIn(user User) {
	return
	var (
		marketUrl = MARKETPLACE_SETTINGS.SiteURL
		text      = fmt.Sprintf("[@%s](%s/user/%s) has logged in", user.Username, marketUrl, user.Username)
	)
	go apis.PostMattermostEvent(MARKETPLACE_SETTINGS.MattermostIncomingHookAuthentication, text)
}

func EventUserRegistred(user User) {
	var (
		marketUrl = MARKETPLACE_SETTINGS.SiteURL
		text      = fmt.Sprintf("[@%s](%s/user/%s) has registered", user.Username, marketUrl, user.Username)
	)
	go apis.PostMattermostEvent(MARKETPLACE_SETTINGS.MattermostIncomingHookAuthentication, text)
}

func EventUserSubscription(user User, plan, currency, address, rawTx string) {
	var (
		marketUrl = MARKETPLACE_SETTINGS.SiteURL
		text      = fmt.Sprintf("[@%s](%s/user/%s) has purchased %s account via %s: %s \t raw_tx: %s", user.Username, marketUrl, user.Username, plan, currency, address, rawTx)
	)
	go apis.PostMattermostEvent(MARKETPLACE_SETTINGS.MattermostIncomingHookSubscriptions, text)
}
