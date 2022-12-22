package deposit

import (
	"coinbit/proto/pb"

	"google.golang.org/protobuf/proto"
)

type DepositListCodec struct{}

func (c *DepositListCodec) Encode(value interface{}) ([]byte, error) {
	return proto.Marshal(value.(*pb.DepositHistory))
}

func (c *DepositListCodec) Decode(data []byte) (interface{}, error) {
	var m pb.DepositHistory
	return &m, proto.Unmarshal(data, &m)
}
