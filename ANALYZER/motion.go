package analyzer

import (
	"behnama/pkg/protomodels/analyzer/motion"

	"google.golang.org/protobuf/proto"

	"behnama/pkg/core/log"
)

type MotionsRegion struct {
	Row    uint   `json:"row"`
	Column uint   `json:"column"`
	Cells  string `json:"cells"`
}

func AnalyzerMotion(data []byte) (MotionsRegion, error) {
	motionPacket := &motion.Region{}
	var motionRegion MotionsRegion
	if err := proto.Unmarshal(data, motionPacket); err != nil {
		log.Warn("AnalyzerMotion motionPacket is wrong")
		return motionRegion, err
	}
	motionRegion.Row = uint(motionPacket.Rows)
	motionRegion.Column = uint(motionPacket.Columns)
	motionRegion.Cells = motionPacket.Cells
	return motionRegion, nil
}
