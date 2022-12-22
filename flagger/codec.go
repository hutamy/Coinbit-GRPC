package flagger

import (
	"coinbit/proto/pb"

	"google.golang.org/protobuf/proto"
)

type FlagEventCodec struct{}

func (c *FlagEventCodec) Encode(value interface{}) ([]byte, error) {
	return proto.Marshal(value.(*pb.FlagEvent))
}

func (c *FlagEventCodec) Decode(data []byte) (interface{}, error) {
	var m pb.FlagEvent
	return &m, proto.Unmarshal(data, &m)
}

type FlagValueCodec struct{}

func (c *FlagValueCodec) Encode(value interface{}) ([]byte, error) {
	return proto.Marshal(value.(*pb.FlagValue))
}

func (c *FlagValueCodec) Decode(data []byte) (interface{}, error) {
	var m pb.FlagValue
	return &m, proto.Unmarshal(data, &m)
}
