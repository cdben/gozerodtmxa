package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/dtmgrpc"
	"xa/trans/internal/cores"

	"xa/trans/internal/svc"
	"xa/trans/trans"

	"github.com/zeromicro/go-zero/core/logx"
)

type TransInXaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTransInXaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransInXaLogic {
	return &TransInXaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TransInXaLogic) TransInXa(in *trans.AdjustInfo) (*trans.Response, error) {
	err := dtmgrpc.XaLocalTransaction(l.ctx, cores.ToConfig(l.svcCtx.Config.Mysql), func(db *sql.DB, xa *dtmgrpc.XaGrpc) error {
		r, err := l.svcCtx.UserAccountModel.AdjustBalance(db, in.UserID, in.Amount)

		if err == nil && r == 0 {
			return fmt.Errorf("update error, balance not enough")
		}
		return nil
	})
	return &trans.Response{}, err

	return &trans.Response{}, nil
}
