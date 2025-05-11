package ctrl

import (
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"go.osspkg.com/goppy/v2/web"
	"go.osspkg.com/logx"
	"go.osspkg.com/random"
)

func (v *Controller) authCheckSession() web.Middleware {
	return v.jwt.GuardMiddlewareWithCallback(
		func(w http.ResponseWriter, r *http.Request) bool {
			if strings.HasPrefix(r.RequestURI, "/auth") {
				return true
			}
			if strings.HasPrefix(r.RequestURI, "/@/auth") {
				return true
			}
			ext := filepath.Ext(r.RequestURI)
			return len(ext) > 0
		},
		func(w http.ResponseWriter, r *http.Request) bool {
			return false
		},
		func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.RequestURI, "/@") {
				http.Error(w, "authorization required", http.StatusUnauthorized)
				return
			}
			http.Redirect(w, r, "/auth", http.StatusMovedPermanently)
		},
	)
}

func (v *Controller) authSignIn(ctx web.Context) {
	v.jwt.SignCookie(ctx, Empty{}, 24*time.Hour)
	//ctx.JSON(200, Empty{})
}

func (v *Controller) authSignUp(ctx web.Context) {
	ctx.JSON(200, Empty{})
}

func (v *Controller) authLogout(ctx web.Context) {
	data := Empty{Data: random.String(10)}
	err := v.jwt.SignCookie(ctx, data, -3600*time.Hour)
	if err != nil {
		logx.Warn("logout", "err", err)
		ctx.JSON(500, Empty{})
	}
}
