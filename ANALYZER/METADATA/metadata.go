package metadata

import (
	"behnama/pkg/core/log"
	"github.com/golang/protobuf/proto"
)

func GetMetadata(data []byte) (*Metadata, error) {
	metadataPacket := &Metadata{}
	if err := proto.Unmarshal(data, metadataPacket); err != nil {
		log.Warn("AnalyzerMotion motionPacket is wrong")
		return metadataPacket, err
	}
	return metadataPacket, nil
}
