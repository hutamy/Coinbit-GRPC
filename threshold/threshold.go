package threshold

import (
	"coinbit/deposit"
	"coinbit/flagger"
	"coinbit/proto/pb"
	"context"
	"time"

	"github.com/lovoo/goka"
)

const (
	rollingPeriod = 120 // 2 minutes
	maxAmount     = 10000
)

var (
	group goka.Group = "threshold"
)

func Run(ctx context.Context, brokers []string) func() error {
	return func() error {
		g := goka.DefineGroup(
			group,
			goka.Input(
				deposit.DepositStream,
				new(deposit.DepositCodec),
				validateThreshold,
			),
			goka.Output(flagger.Stream, new(flagger.FlagEventCodec)),
			goka.Persist(new(CounterCodec)),
		)

		p, err := goka.NewProcessor(brokers, g)
		if err != nil {
			return err
		}

		return p.Run(ctx)
	}
}

func validateThreshold(ctx goka.Context, msg interface{}) {
	ctxVal := ctx.Value()
	c := &pb.Counter{}
	if ctxVal != nil {
		c = ctxVal.(*pb.Counter)
	}

	m := msg.(*pb.Deposit)
	c.Recieved += m.Amount
	if c.StartPeriod == 0 {
		c.StartPeriod = time.Now().Unix()
	} else {
		if time.Now().Unix()-c.StartPeriod > rollingPeriod {
			c.StartPeriod = 0
			c.Recieved = 0
		}
	}
	ctx.SetValue(c)

	if (c.Recieved > maxAmount) && (c.StartPeriod != 0) {
		ctx.Emit(flagger.Stream, ctx.Key(), &pb.FlagEvent{
			FlagRemoved: false,
			StartPeriod: c.StartPeriod,
		})
	} else {
		ctx.Emit(flagger.Stream, ctx.Key(), &pb.FlagEvent{
			FlagRemoved: true,
		})
	}
}
