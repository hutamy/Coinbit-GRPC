package flagger

import (
	"coinbit/proto/pb"
	"context"

	"github.com/lovoo/goka"
)

var (
	group  goka.Group  = "flagger"
	Table  goka.Table  = goka.GroupTable(group)
	Stream goka.Stream = "flag-wallet"
)

func Run(ctx context.Context, brokers []string) func() error {
	return func() error {
		g := goka.DefineGroup(
			group,
			goka.Input(
				Stream,
				new(FlagEventCodec),
				flag,
			),
			goka.Persist(new(FlagValueCodec)),
		)

		p, err := goka.NewProcessor(brokers, g)
		if err != nil {
			return err
		}

		return p.Run(ctx)
	}
}

func flag(ctx goka.Context, msg interface{}) {
	var s *pb.FlagValue
	v := ctx.Value()
	if v == nil {
		s = new(pb.FlagValue)
	} else {
		s = v.(*pb.FlagValue)
	}

	flagEvent := msg.(*pb.FlagEvent)
	if flagEvent.FlagRemoved {
		s.Flagged = false
		s.StartPeriod = 0
	} else {
		s.Flagged = true
		s.StartPeriod = flagEvent.StartPeriod
	}
	ctx.SetValue(s)
}
