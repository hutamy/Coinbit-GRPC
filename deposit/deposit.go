package deposit

import (
	"coinbit/proto/pb"

	"github.com/lovoo/goka"
	"google.golang.org/protobuf/proto"
)

var (
	DepositStream goka.Stream = "deposit"
)

type DepositCodec struct{}

func (c *DepositCodec) Encode(value interface{}) ([]byte, error) {
	return proto.Marshal(value.(*pb.Deposit))
}

func (c *DepositCodec) Decode(data []byte) (interface{}, error) {
	var m pb.Deposit
	return &m, proto.Unmarshal(data, &m)
}
