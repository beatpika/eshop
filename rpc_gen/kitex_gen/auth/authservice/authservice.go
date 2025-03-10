// Code generated by Kitex v0.9.1. DO NOT EDIT.

package authservice

import (
	"context"
	"errors"
	auth "github.com/beatpika/eshop/rpc_gen/kitex_gen/auth"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"GenerateToken": kitex.NewMethodInfo(
		generateTokenHandler,
		newGenerateTokenArgs,
		newGenerateTokenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"VerifyToken": kitex.NewMethodInfo(
		verifyTokenHandler,
		newVerifyTokenArgs,
		newVerifyTokenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"RefreshToken": kitex.NewMethodInfo(
		refreshTokenHandler,
		newRefreshTokenArgs,
		newRefreshTokenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"RevokeToken": kitex.NewMethodInfo(
		revokeTokenHandler,
		newRevokeTokenArgs,
		newRevokeTokenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	authServiceServiceInfo                = NewServiceInfo()
	authServiceServiceInfoForClient       = NewServiceInfoForClient()
	authServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return authServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return authServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return authServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "AuthService"
	handlerType := (*auth.AuthService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "auth",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func generateTokenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(auth.GenerateTokenRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(auth.AuthService).GenerateToken(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GenerateTokenArgs:
		success, err := handler.(auth.AuthService).GenerateToken(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GenerateTokenResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGenerateTokenArgs() interface{} {
	return &GenerateTokenArgs{}
}

func newGenerateTokenResult() interface{} {
	return &GenerateTokenResult{}
}

type GenerateTokenArgs struct {
	Req *auth.GenerateTokenRequest
}

func (p *GenerateTokenArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(auth.GenerateTokenRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GenerateTokenArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GenerateTokenArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GenerateTokenArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GenerateTokenArgs) Unmarshal(in []byte) error {
	msg := new(auth.GenerateTokenRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GenerateTokenArgs_Req_DEFAULT *auth.GenerateTokenRequest

func (p *GenerateTokenArgs) GetReq() *auth.GenerateTokenRequest {
	if !p.IsSetReq() {
		return GenerateTokenArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GenerateTokenArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GenerateTokenArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GenerateTokenResult struct {
	Success *auth.GenerateTokenResponse
}

var GenerateTokenResult_Success_DEFAULT *auth.GenerateTokenResponse

func (p *GenerateTokenResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(auth.GenerateTokenResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GenerateTokenResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GenerateTokenResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GenerateTokenResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GenerateTokenResult) Unmarshal(in []byte) error {
	msg := new(auth.GenerateTokenResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GenerateTokenResult) GetSuccess() *auth.GenerateTokenResponse {
	if !p.IsSetSuccess() {
		return GenerateTokenResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GenerateTokenResult) SetSuccess(x interface{}) {
	p.Success = x.(*auth.GenerateTokenResponse)
}

func (p *GenerateTokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GenerateTokenResult) GetResult() interface{} {
	return p.Success
}

func verifyTokenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(auth.VerifyTokenRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(auth.AuthService).VerifyToken(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *VerifyTokenArgs:
		success, err := handler.(auth.AuthService).VerifyToken(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*VerifyTokenResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newVerifyTokenArgs() interface{} {
	return &VerifyTokenArgs{}
}

func newVerifyTokenResult() interface{} {
	return &VerifyTokenResult{}
}

type VerifyTokenArgs struct {
	Req *auth.VerifyTokenRequest
}

func (p *VerifyTokenArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(auth.VerifyTokenRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *VerifyTokenArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *VerifyTokenArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *VerifyTokenArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *VerifyTokenArgs) Unmarshal(in []byte) error {
	msg := new(auth.VerifyTokenRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var VerifyTokenArgs_Req_DEFAULT *auth.VerifyTokenRequest

func (p *VerifyTokenArgs) GetReq() *auth.VerifyTokenRequest {
	if !p.IsSetReq() {
		return VerifyTokenArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *VerifyTokenArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *VerifyTokenArgs) GetFirstArgument() interface{} {
	return p.Req
}

type VerifyTokenResult struct {
	Success *auth.VerifyTokenResponse
}

var VerifyTokenResult_Success_DEFAULT *auth.VerifyTokenResponse

func (p *VerifyTokenResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(auth.VerifyTokenResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *VerifyTokenResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *VerifyTokenResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *VerifyTokenResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *VerifyTokenResult) Unmarshal(in []byte) error {
	msg := new(auth.VerifyTokenResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *VerifyTokenResult) GetSuccess() *auth.VerifyTokenResponse {
	if !p.IsSetSuccess() {
		return VerifyTokenResult_Success_DEFAULT
	}
	return p.Success
}

func (p *VerifyTokenResult) SetSuccess(x interface{}) {
	p.Success = x.(*auth.VerifyTokenResponse)
}

func (p *VerifyTokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *VerifyTokenResult) GetResult() interface{} {
	return p.Success
}

func refreshTokenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(auth.RefreshTokenRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(auth.AuthService).RefreshToken(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *RefreshTokenArgs:
		success, err := handler.(auth.AuthService).RefreshToken(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RefreshTokenResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newRefreshTokenArgs() interface{} {
	return &RefreshTokenArgs{}
}

func newRefreshTokenResult() interface{} {
	return &RefreshTokenResult{}
}

type RefreshTokenArgs struct {
	Req *auth.RefreshTokenRequest
}

func (p *RefreshTokenArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(auth.RefreshTokenRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RefreshTokenArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RefreshTokenArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RefreshTokenArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *RefreshTokenArgs) Unmarshal(in []byte) error {
	msg := new(auth.RefreshTokenRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RefreshTokenArgs_Req_DEFAULT *auth.RefreshTokenRequest

func (p *RefreshTokenArgs) GetReq() *auth.RefreshTokenRequest {
	if !p.IsSetReq() {
		return RefreshTokenArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RefreshTokenArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RefreshTokenArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RefreshTokenResult struct {
	Success *auth.RefreshTokenResponse
}

var RefreshTokenResult_Success_DEFAULT *auth.RefreshTokenResponse

func (p *RefreshTokenResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(auth.RefreshTokenResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RefreshTokenResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RefreshTokenResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RefreshTokenResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *RefreshTokenResult) Unmarshal(in []byte) error {
	msg := new(auth.RefreshTokenResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RefreshTokenResult) GetSuccess() *auth.RefreshTokenResponse {
	if !p.IsSetSuccess() {
		return RefreshTokenResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RefreshTokenResult) SetSuccess(x interface{}) {
	p.Success = x.(*auth.RefreshTokenResponse)
}

func (p *RefreshTokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RefreshTokenResult) GetResult() interface{} {
	return p.Success
}

func revokeTokenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(auth.RevokeTokenRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(auth.AuthService).RevokeToken(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *RevokeTokenArgs:
		success, err := handler.(auth.AuthService).RevokeToken(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RevokeTokenResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newRevokeTokenArgs() interface{} {
	return &RevokeTokenArgs{}
}

func newRevokeTokenResult() interface{} {
	return &RevokeTokenResult{}
}

type RevokeTokenArgs struct {
	Req *auth.RevokeTokenRequest
}

func (p *RevokeTokenArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(auth.RevokeTokenRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RevokeTokenArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RevokeTokenArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RevokeTokenArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *RevokeTokenArgs) Unmarshal(in []byte) error {
	msg := new(auth.RevokeTokenRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RevokeTokenArgs_Req_DEFAULT *auth.RevokeTokenRequest

func (p *RevokeTokenArgs) GetReq() *auth.RevokeTokenRequest {
	if !p.IsSetReq() {
		return RevokeTokenArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RevokeTokenArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RevokeTokenArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RevokeTokenResult struct {
	Success *auth.RevokeTokenResponse
}

var RevokeTokenResult_Success_DEFAULT *auth.RevokeTokenResponse

func (p *RevokeTokenResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(auth.RevokeTokenResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RevokeTokenResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RevokeTokenResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RevokeTokenResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *RevokeTokenResult) Unmarshal(in []byte) error {
	msg := new(auth.RevokeTokenResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RevokeTokenResult) GetSuccess() *auth.RevokeTokenResponse {
	if !p.IsSetSuccess() {
		return RevokeTokenResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RevokeTokenResult) SetSuccess(x interface{}) {
	p.Success = x.(*auth.RevokeTokenResponse)
}

func (p *RevokeTokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RevokeTokenResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GenerateToken(ctx context.Context, Req *auth.GenerateTokenRequest) (r *auth.GenerateTokenResponse, err error) {
	var _args GenerateTokenArgs
	_args.Req = Req
	var _result GenerateTokenResult
	if err = p.c.Call(ctx, "GenerateToken", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VerifyToken(ctx context.Context, Req *auth.VerifyTokenRequest) (r *auth.VerifyTokenResponse, err error) {
	var _args VerifyTokenArgs
	_args.Req = Req
	var _result VerifyTokenResult
	if err = p.c.Call(ctx, "VerifyToken", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RefreshToken(ctx context.Context, Req *auth.RefreshTokenRequest) (r *auth.RefreshTokenResponse, err error) {
	var _args RefreshTokenArgs
	_args.Req = Req
	var _result RefreshTokenResult
	if err = p.c.Call(ctx, "RefreshToken", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RevokeToken(ctx context.Context, Req *auth.RevokeTokenRequest) (r *auth.RevokeTokenResponse, err error) {
	var _args RevokeTokenArgs
	_args.Req = Req
	var _result RevokeTokenResult
	if err = p.c.Call(ctx, "RevokeToken", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
