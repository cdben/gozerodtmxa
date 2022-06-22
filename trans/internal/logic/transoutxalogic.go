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

type TransOutXaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTransOutXaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransOutXaLogic {
	return &TransOutXaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TransOutXaLogic) TransOutXa(in *trans.AdjustInfo) (*trans.Response, error) {
	err := dtmgrpc.XaLocalTransaction(
		l.ctx,
		cores.ToConfig(l.svcCtx.Config.Mysql),
		func(db *sql.DB, xa *dtmgrpc.XaGrpc) error {
			r, err := l.svcCtx.UserAccountModel.AdjustBalance(db, in.UserID, -in.Amount)
			if err != nil {
				return err
			}
			if r == 0 {
				return fmt.Errorf("update error, balance not enough")
			}
			return nil
		})
	return &trans.Response{}, err
}
