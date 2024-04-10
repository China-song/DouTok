// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	"fmt"
	user "github.com/TremblingV5/DouTok/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Register":         kitex.NewMethodInfo(registerHandler, newRegisterArgs, newRegisterResult, false),
		"Login":            kitex.NewMethodInfo(loginHandler, newLoginArgs, newLoginResult, false),
		"GetUserById":      kitex.NewMethodInfo(getUserByIdHandler, newGetUserByIdArgs, newGetUserByIdResult, false),
		"GetUserListByIds": kitex.NewMethodInfo(getUserListByIdsHandler, newGetUserListByIdsArgs, newGetUserListByIdsResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.DouyinUserRegisterRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).Register(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *RegisterArgs:
		success, err := handler.(user.UserService).Register(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RegisterResult)
		realResult.Success = success
	}
	return nil
}
func newRegisterArgs() interface{} {
	return &RegisterArgs{}
}

func newRegisterResult() interface{} {
	return &RegisterResult{}
}

type RegisterArgs struct {
	Req *user.DouyinUserRegisterRequest
}

func (p *RegisterArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.DouyinUserRegisterRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RegisterArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RegisterArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RegisterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in RegisterArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *RegisterArgs) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserRegisterRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RegisterArgs_Req_DEFAULT *user.DouyinUserRegisterRequest

func (p *RegisterArgs) GetReq() *user.DouyinUserRegisterRequest {
	if !p.IsSetReq() {
		return RegisterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RegisterArgs) IsSetReq() bool {
	return p.Req != nil
}

type RegisterResult struct {
	Success *user.DouyinUserRegisterResponse
}

var RegisterResult_Success_DEFAULT *user.DouyinUserRegisterResponse

func (p *RegisterResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.DouyinUserRegisterResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RegisterResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RegisterResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RegisterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in RegisterResult")
	}
	return proto.Marshal(p.Success)
}

func (p *RegisterResult) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserRegisterResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RegisterResult) GetSuccess() *user.DouyinUserRegisterResponse {
	if !p.IsSetSuccess() {
		return RegisterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RegisterResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.DouyinUserRegisterResponse)
}

func (p *RegisterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.DouyinUserLoginRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).Login(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *LoginArgs:
		success, err := handler.(user.UserService).Login(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*LoginResult)
		realResult.Success = success
	}
	return nil
}
func newLoginArgs() interface{} {
	return &LoginArgs{}
}

func newLoginResult() interface{} {
	return &LoginResult{}
}

type LoginArgs struct {
	Req *user.DouyinUserLoginRequest
}

func (p *LoginArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.DouyinUserLoginRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *LoginArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *LoginArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *LoginArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in LoginArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *LoginArgs) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserLoginRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var LoginArgs_Req_DEFAULT *user.DouyinUserLoginRequest

func (p *LoginArgs) GetReq() *user.DouyinUserLoginRequest {
	if !p.IsSetReq() {
		return LoginArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *LoginArgs) IsSetReq() bool {
	return p.Req != nil
}

type LoginResult struct {
	Success *user.DouyinUserLoginResponse
}

var LoginResult_Success_DEFAULT *user.DouyinUserLoginResponse

func (p *LoginResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.DouyinUserLoginResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *LoginResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *LoginResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *LoginResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in LoginResult")
	}
	return proto.Marshal(p.Success)
}

func (p *LoginResult) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserLoginResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *LoginResult) GetSuccess() *user.DouyinUserLoginResponse {
	if !p.IsSetSuccess() {
		return LoginResult_Success_DEFAULT
	}
	return p.Success
}

func (p *LoginResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.DouyinUserLoginResponse)
}

func (p *LoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getUserByIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.DouyinUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).GetUserById(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetUserByIdArgs:
		success, err := handler.(user.UserService).GetUserById(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserByIdResult)
		realResult.Success = success
	}
	return nil
}
func newGetUserByIdArgs() interface{} {
	return &GetUserByIdArgs{}
}

func newGetUserByIdResult() interface{} {
	return &GetUserByIdResult{}
}

type GetUserByIdArgs struct {
	Req *user.DouyinUserRequest
}

func (p *GetUserByIdArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.DouyinUserRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetUserByIdArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetUserByIdArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetUserByIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUserByIdArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserByIdArgs) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserByIdArgs_Req_DEFAULT *user.DouyinUserRequest

func (p *GetUserByIdArgs) GetReq() *user.DouyinUserRequest {
	if !p.IsSetReq() {
		return GetUserByIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserByIdArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUserByIdResult struct {
	Success *user.DouyinUserResponse
}

var GetUserByIdResult_Success_DEFAULT *user.DouyinUserResponse

func (p *GetUserByIdResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.DouyinUserResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetUserByIdResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetUserByIdResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetUserByIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUserByIdResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserByIdResult) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserByIdResult) GetSuccess() *user.DouyinUserResponse {
	if !p.IsSetSuccess() {
		return GetUserByIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserByIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.DouyinUserResponse)
}

func (p *GetUserByIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getUserListByIdsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.DouyinUserListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).GetUserListByIds(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetUserListByIdsArgs:
		success, err := handler.(user.UserService).GetUserListByIds(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserListByIdsResult)
		realResult.Success = success
	}
	return nil
}
func newGetUserListByIdsArgs() interface{} {
	return &GetUserListByIdsArgs{}
}

func newGetUserListByIdsResult() interface{} {
	return &GetUserListByIdsResult{}
}

type GetUserListByIdsArgs struct {
	Req *user.DouyinUserListRequest
}

func (p *GetUserListByIdsArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.DouyinUserListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetUserListByIdsArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetUserListByIdsArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetUserListByIdsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUserListByIdsArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserListByIdsArgs) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserListByIdsArgs_Req_DEFAULT *user.DouyinUserListRequest

func (p *GetUserListByIdsArgs) GetReq() *user.DouyinUserListRequest {
	if !p.IsSetReq() {
		return GetUserListByIdsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserListByIdsArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUserListByIdsResult struct {
	Success *user.DouyinUserListResponse
}

var GetUserListByIdsResult_Success_DEFAULT *user.DouyinUserListResponse

func (p *GetUserListByIdsResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.DouyinUserListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetUserListByIdsResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetUserListByIdsResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetUserListByIdsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUserListByIdsResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserListByIdsResult) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserListByIdsResult) GetSuccess() *user.DouyinUserListResponse {
	if !p.IsSetSuccess() {
		return GetUserListByIdsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserListByIdsResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.DouyinUserListResponse)
}

func (p *GetUserListByIdsResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, Req *user.DouyinUserRegisterRequest) (r *user.DouyinUserRegisterResponse, err error) {
	var _args RegisterArgs
	_args.Req = Req
	var _result RegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, Req *user.DouyinUserLoginRequest) (r *user.DouyinUserLoginResponse, err error) {
	var _args LoginArgs
	_args.Req = Req
	var _result LoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserById(ctx context.Context, Req *user.DouyinUserRequest) (r *user.DouyinUserResponse, err error) {
	var _args GetUserByIdArgs
	_args.Req = Req
	var _result GetUserByIdResult
	if err = p.c.Call(ctx, "GetUserById", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserListByIds(ctx context.Context, Req *user.DouyinUserListRequest) (r *user.DouyinUserListResponse, err error) {
	var _args GetUserListByIdsArgs
	_args.Req = Req
	var _result GetUserListByIdsResult
	if err = p.c.Call(ctx, "GetUserListByIds", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
