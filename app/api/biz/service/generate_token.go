package service

import (
	"context"

	token "github.com/beatpika/eshop/app/api/hertz_gen/basic/token"
	"github.com/cloudwego/hertz/pkg/app"
)

type GenerateTokenService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGenerateTokenService(Context context.Context, RequestContext *app.RequestContext) *GenerateTokenService {
	return &GenerateTokenService{RequestContext: RequestContext, Context: Context}
}

func (h *GenerateTokenService) Run(req *token.GenerateTokenReq) (resp *token.GenerateTokenResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
