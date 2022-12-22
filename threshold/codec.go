package threshold

import (
	"coinbit/proto/pb"

	"google.golang.org/protobuf/proto"
)

type CounterCodec struct{}

func (c *CounterCodec) Encode(value interface{}) ([]byte, error) {
	return proto.Marshal(value.(*pb.Counter))
}

func (c *CounterCodec) Decode(data []byte) (interface{}, error) {
	var m pb.Counter
	return &m, proto.Unmarshal(data, &m)
}
