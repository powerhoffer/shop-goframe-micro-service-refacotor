package user_info

import (
	"context"
	v1 "shop-goframe-micro-service-refacotor/app/user/api/user_info/v1"
	"shop-goframe-micro-service-refacotor/app/user/internal/consts"
	v2 "shop-goframe-micro-service-refacotor/app/user/internal/logic/user_info"
	"shop-goframe-micro-service-refacotor/app/user/internal/model/entity"
	"time"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Register(s *grpcx.GrpcServer) {
	v1.RegisterUserInfoServer(s.Server, &Controller{})
}

type Controller struct {
	v1.UnimplementedUserInfoServer
}

func (*Controller) Login(ctx context.Context, req *v1.UserInfoLoginReq) (res *v1.UserInfoLoginRes, err error) {
	// 调用logic层
	token, expireIn, userInfo, err := v2.Login(ctx, req.Name, req.Password)
	// 错误类型
	infoError := consts.InfoError(consts.UserInfo, consts.LoginFail)
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 计算过期时间
	expireTime := time.Now().Add(time.Duration(expireIn) * time.Second)
	expireProto := timestamppb.New(expireTime)
	if err := expireProto.CheckValid(); err != nil {
		return nil, err
	}

	// 返回响应
	return &v1.UserInfoLoginRes{
		Type:     "Bearer",
		Token:    token,
		ExpireIn: uint32(expireIn),
		UserInfo: &v1.UserInfoBase{
			Id:     uint32(userInfo.Id),
			Name:   userInfo.Name,
			Avatar: userInfo.Avatar,
			Sex:    uint32(userInfo.Sex),
			Sign:   userInfo.Sign,
			Status: uint32(userInfo.Status),
		},
	}, nil
}

func (c *Controller) Register(ctx context.Context, req *v1.UserInfoRegisterReq) (*v1.UserInfoRegisterRes, error) {
	var registerData *entity.UserInfo
	// 将请求参数req转换为实体对象consigneeInfo
	if err := gconv.Struct(req, &registerData); err != nil {
		return nil, err
	}
	// 错误类型
	infoError := consts.InfoError(consts.UserInfo, consts.RegisterFail)
	// 调用logic层注册
	userInfo, err := v2.Register(ctx, registerData)
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回响应
	return &v1.UserInfoRegisterRes{
		Id: uint32(userInfo.Id),
	}, nil
}

func (c *Controller) UpdatePassword(ctx context.Context, req *v1.UserInfoUpdatePasswordReq) (*v1.UserInfoUpdatePasswordRes, error) {
	// 调用logic层修改密码
	err := v2.UpdatePassword(ctx, int(req.Id), req.Password, req.SecretAnswer)
	// 错误类型
	infoError := consts.InfoError(consts.UserInfo, consts.UpdatePasswordFail)
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回响应
	return &v1.UserInfoUpdatePasswordRes{
		Id: req.Id,
	}, nil
}

func (*Controller) GetUserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	// 调用logic层获取用户信息
	userInfo, err := v2.GetUserInfo(ctx, int(req.Id))
	// 错误类型
	infoError := consts.InfoError(consts.UserInfo, consts.GetUserInfoFail)
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回响应
	return &v1.UserInfoRes{
		UserInfo: &v1.UserInfoBase{
			Id:     uint32(userInfo.Id),
			Name:   userInfo.Name,
			Avatar: userInfo.Avatar,
			Sex:    uint32(userInfo.Sex),
			Sign:   userInfo.Sign,
			Status: uint32(userInfo.Status),
		},
	}, nil
}
