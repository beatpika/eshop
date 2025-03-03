package service

import (
	"context"

	user "github.com/beatpika/eshop/app/api/hertz_gen/basic/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdatePasswordService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdatePasswordService(Context context.Context, RequestContext *app.RequestContext) *UpdatePasswordService {
	return &UpdatePasswordService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdatePasswordService) Run(req *user.UpdatePasswordReq) (resp *user.UpdatePasswordResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
