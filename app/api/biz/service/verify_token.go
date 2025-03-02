package service

import (
	"context"

	token "github.com/beatpika/eshop/app/api/hertz_gen/basic/token"
	"github.com/cloudwego/hertz/pkg/app"
)

type VerifyTokenService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewVerifyTokenService(Context context.Context, RequestContext *app.RequestContext) *VerifyTokenService {
	return &VerifyTokenService{RequestContext: RequestContext, Context: Context}
}

func (h *VerifyTokenService) Run(req *token.VerifyTokenReq) (resp *token.VerifyTokenResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
