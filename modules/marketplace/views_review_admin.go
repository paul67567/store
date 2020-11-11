package marketplace

import (
	"github.com/gocraft/web"

	"github.com/paul67567/store/modules/util"
)

func (c *Context) AdminReviews(w web.ResponseWriter, r *web.Request) {
	// c.Reviews = GetAllReviews()
	util.RenderTemplate(w, "reviews/admin/reviews", c)
}
