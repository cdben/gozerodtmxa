package logic

import (
	"context"
	"fmt"
	"github.com/dtm-labs/dtmdriver"
	driver "github.com/dtm-labs/dtmdriver-gozero"
	"github.com/dtm-labs/dtmgrpc"
	"xa/api/internal/svc"
	"xa/api/internal/types"
	"xa/trans/transclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TransLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTransLogic(ctx context.Context, svcCtx *svc.ServiceContext) TransLogic {
	return TransLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TransLogic) Trans(req types.TransRequest) (resp *types.TransResponse, err error) {
	transRpcBusiServer, err := l.svcCtx.Config.TransRpc.BuildTarget()
	if err != nil {
		return nil, err
	}
	// dtm 服务的 etcd 注册地址
	var dtmServer = "etcd://localhost:2379/dtmservice"
	// 创建一个gid
	gid := dtmgrpc.MustGenGid(dtmServer)
	if err := dtmdriver.Use(driver.DriverName); err != nil {
		return nil, err
	}
	err = dtmgrpc.XaGlobalTransaction(dtmServer, gid, func(xa *dtmgrpc.XaGrpc) error {
		var r transclient.Response
		var r1 transclient.Response
		err1 := xa.CallBranch(&transclient.AdjustInfo{
			UserID: req.UserId,
			Amount: req.Amount,
		},
			transRpcBusiServer+"/transclient.Trans/TransOutXa", &r)

		err2 := xa.CallBranch(&transclient.AdjustInfo{
			UserID: req.ToUserId,
			Amount: req.Amount,
		},
			transRpcBusiServer+"/transclient.Trans/TransInXa", &r1)

		if err2 != nil || err1 != nil {
			return fmt.Errorf("xa2 error")
		}
		return nil
	})
	return
}
