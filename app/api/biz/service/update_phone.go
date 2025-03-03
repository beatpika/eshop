package service

import (
	"context"

	user "github.com/beatpika/eshop/app/api/hertz_gen/basic/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdatePhoneService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdatePhoneService(Context context.Context, RequestContext *app.RequestContext) *UpdatePhoneService {
	return &UpdatePhoneService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdatePhoneService) Run(req *user.UpdatePhoneReq) (resp *user.UpdatePhoneResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
