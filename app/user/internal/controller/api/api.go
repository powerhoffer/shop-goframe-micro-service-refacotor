package api

import (
	"context"

	v1 "shop-goframe-micro-service-refacotor/app/user/api/consignee_info/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedConsigneeInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterConsigneeInfoServer(s.Server, &Controller{})
}

func (*Controller) GetList(ctx context.Context, req *v1.ConsigneeInfoGetListReq) (res *v1.ConsigneeInfoGetListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) Create(ctx context.Context, req *v1.ConsigneeInfoCreateReq) (res *v1.ConsigneeInfoCreateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) Update(ctx context.Context, req *v1.ConsigneeInfoUpdateReq) (res *v1.ConsigneeInfoUpdateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) Delete(ctx context.Context, req *v1.ConsigneeInfoDeleteReq) (res *v1.ConsigneeInfoDeleteRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
