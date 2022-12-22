package history

import (
	"coinbit/deposit"
	"coinbit/proto/pb"
	"context"

	"github.com/lovoo/goka"
)

var (
	group goka.Group = "balance"
	Table goka.Table = goka.GroupTable(group)
)

func Run(ctx context.Context, brokers []string) func() error {
	return func() error {
		g := goka.DefineGroup(
			group,
			goka.Input(
				deposit.DepositStream,
				new(deposit.DepositCodec),
				balanceHistory,
			),
			goka.Persist(new(deposit.DepositListCodec)),
		)

		p, err := goka.NewProcessor(brokers, g)
		if err != nil {
			return err
		}

		return p.Run(ctx)
	}
}

func balanceHistory(ctx goka.Context, msg interface{}) {
	dh := &pb.DepositHistory{}
	ctxVal := ctx.Value()
	if ctxVal != nil {
		dh = ctxVal.(*pb.DepositHistory)
	}

	d := msg.(*pb.Deposit)
	dh.WalletId = d.WalletId
	dh.Deposits = append(dh.Deposits, d)
	ctx.SetValue(dh)
}
