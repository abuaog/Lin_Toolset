package media

import (
	"behnama/pkg/core/log"
	"github.com/golang/protobuf/proto"
)

func ManualFaceUnmarshal(data []byte) (*ManualFaceMessage, error) {
	result := &ManualFaceMessage{}
	err := proto.Unmarshal(data, result)
	if err != nil {
		log.Warn("AnalyzerMotion motionPacket is wrong")
		return result, err
	}
	return result, nil
}
