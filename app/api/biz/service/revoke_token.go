package service

import (
	"context"

	token "github.com/beatpika/eshop/app/api/hertz_gen/basic/token"
	"github.com/cloudwego/hertz/pkg/app"
)

type RevokeTokenService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRevokeTokenService(Context context.Context, RequestContext *app.RequestContext) *RevokeTokenService {
	return &RevokeTokenService{RequestContext: RequestContext, Context: Context}
}

func (h *RevokeTokenService) Run(req *token.RevokeTokenReq) (resp *token.RevokeTokenResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
