package service

import (
	"context"

	user "github.com/beatpika/eshop/app/api/hertz_gen/basic/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeactivateAccountService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeactivateAccountService(Context context.Context, RequestContext *app.RequestContext) *DeactivateAccountService {
	return &DeactivateAccountService{RequestContext: RequestContext, Context: Context}
}

func (h *DeactivateAccountService) Run(req *user.DeactivateAccountReq) (resp *user.DeactivateAccountResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
