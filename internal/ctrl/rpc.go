package ctrl

import "go.osspkg.com/goppy/v2/web"

func (v *Controller) rpc(ctx web.Context) {
	ctx.JSON(200, Empty{})
}
