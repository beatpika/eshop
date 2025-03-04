package service

import (
	"context"
	"testing"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

func TestUpdateSKUStock_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateSKUStockService(ctx)
	// init req and assert value

	req := &product.UpdateSKUStockReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
