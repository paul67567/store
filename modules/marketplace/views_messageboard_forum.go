package marketplace

import (
	"github.com/gocraft/web"
	"modules/util"
)

func (c *Context) ViewShowMessageboardImage(w web.ResponseWriter, r *web.Request) {
	size := "normal"
	if len(r.URL.Query()["size"]) > 0 {
		size = r.URL.Query()["size"][0]
	}
	util.ServeImage(r.PathParams["uuid"], size, w, r)
}
