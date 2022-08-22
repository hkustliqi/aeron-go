// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"math"
)

type StarbaseIdFactoryStateSnapshot struct {
	LastId int64
	NodeId int32
}

func (s *StarbaseIdFactoryStateSnapshot) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, s.LastId); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, s.NodeId); err != nil {
		return err
	}
	return nil
}

func (s *StarbaseIdFactoryStateSnapshot) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !s.LastIdInActingVersion(actingVersion) {
		s.LastId = s.LastIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &s.LastId); err != nil {
			return err
		}
	}
	if !s.NodeIdInActingVersion(actingVersion) {
		s.NodeId = s.NodeIdNullValue()
	} else {
		if err := _m.ReadInt32(_r, &s.NodeId); err != nil {
			return err
		}
	}
	return nil
}

func (s *StarbaseIdFactoryStateSnapshot) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if s.LastIdInActingVersion(actingVersion) {
		if s.LastId < s.LastIdMinValue() || s.LastId > s.LastIdMaxValue() {
			return fmt.Errorf("Range check failed on s.LastId (%v < %v > %v)", s.LastIdMinValue(), s.LastId, s.LastIdMaxValue())
		}
	}
	if s.NodeIdInActingVersion(actingVersion) {
		if s.NodeId < s.NodeIdMinValue() || s.NodeId > s.NodeIdMaxValue() {
			return fmt.Errorf("Range check failed on s.NodeId (%v < %v > %v)", s.NodeIdMinValue(), s.NodeId, s.NodeIdMaxValue())
		}
	}
	return nil
}

func StarbaseIdFactoryStateSnapshotInit(s *StarbaseIdFactoryStateSnapshot) {
	return
}

func (*StarbaseIdFactoryStateSnapshot) EncodedLength() int64 {
	return 12
}

func (*StarbaseIdFactoryStateSnapshot) LastIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*StarbaseIdFactoryStateSnapshot) LastIdMaxValue() int64 {
	return math.MaxInt64
}

func (*StarbaseIdFactoryStateSnapshot) LastIdNullValue() int64 {
	return math.MinInt64
}

func (*StarbaseIdFactoryStateSnapshot) LastIdSinceVersion() uint16 {
	return 0
}

func (s *StarbaseIdFactoryStateSnapshot) LastIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.LastIdSinceVersion()
}

func (*StarbaseIdFactoryStateSnapshot) LastIdDeprecated() uint16 {
	return 0
}

func (*StarbaseIdFactoryStateSnapshot) NodeIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*StarbaseIdFactoryStateSnapshot) NodeIdMaxValue() int32 {
	return math.MaxInt32
}

func (*StarbaseIdFactoryStateSnapshot) NodeIdNullValue() int32 {
	return math.MinInt32
}

func (*StarbaseIdFactoryStateSnapshot) NodeIdSinceVersion() uint16 {
	return 0
}

func (s *StarbaseIdFactoryStateSnapshot) NodeIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.NodeIdSinceVersion()
}

func (*StarbaseIdFactoryStateSnapshot) NodeIdDeprecated() uint16 {
	return 0
}
