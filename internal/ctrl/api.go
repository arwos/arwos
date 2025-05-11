package ctrl

import (
	"context"

	"go.arwos.org/arwos/internal/pkg/webcache"
	"go.osspkg.com/goppy/v2/auth/jwt"
	"go.osspkg.com/goppy/v2/web"
)

type Controller struct {
	conf     *ConfigGroup
	router   web.RouterPool
	webcache *webcache.Cache
	jwt      jwt.JWT
}

func NewAPI(r web.RouterPool, c *ConfigGroup, wc *webcache.Cache, a jwt.JWT) *Controller {
	return &Controller{
		router:   r,
		conf:     c,
		webcache: wc,
		jwt:      a,
	}
}

func (v *Controller) Up(_ context.Context) error {
	v.router.All(func(_ string, r web.Router) {
		r.Use(v.authCheckSession())

		rpcApi := r.Collection("/@/rpc")
		rpcApi.Post("/", v.rpc)

		authApi := r.Collection("/@/auth")
		authApi.Post("/pam/sign-in", v.authSignIn)
		authApi.Post("/pam/sign-up", v.authSignUp)
		authApi.Post("/logout", v.authLogout)

		r.Get("/", v.webcache.Handler)
		r.Get("/#", v.webcache.Handler)
	})

	return nil
}

func (v *Controller) Down() error {
	return nil
}
