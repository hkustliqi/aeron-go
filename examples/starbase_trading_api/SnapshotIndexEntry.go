// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type SnapshotIndexEntry struct {
	Timestamp            int64
	LogPosition          int64
	EventChannelPosition int64
	SnapshotPosition     int64
	SnapshotLength       int32
}

func (s *SnapshotIndexEntry) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := s.RangeCheck(s.SbeSchemaVersion(), s.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, s.Timestamp); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, s.LogPosition); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, s.EventChannelPosition); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, s.SnapshotPosition); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, s.SnapshotLength); err != nil {
		return err
	}
	return nil
}

func (s *SnapshotIndexEntry) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !s.TimestampInActingVersion(actingVersion) {
		s.Timestamp = s.TimestampNullValue()
	} else {
		if err := _m.ReadInt64(_r, &s.Timestamp); err != nil {
			return err
		}
	}
	if !s.LogPositionInActingVersion(actingVersion) {
		s.LogPosition = s.LogPositionNullValue()
	} else {
		if err := _m.ReadInt64(_r, &s.LogPosition); err != nil {
			return err
		}
	}
	if !s.EventChannelPositionInActingVersion(actingVersion) {
		s.EventChannelPosition = s.EventChannelPositionNullValue()
	} else {
		if err := _m.ReadInt64(_r, &s.EventChannelPosition); err != nil {
			return err
		}
	}
	if !s.SnapshotPositionInActingVersion(actingVersion) {
		s.SnapshotPosition = s.SnapshotPositionNullValue()
	} else {
		if err := _m.ReadInt64(_r, &s.SnapshotPosition); err != nil {
			return err
		}
	}
	if !s.SnapshotLengthInActingVersion(actingVersion) {
		s.SnapshotLength = s.SnapshotLengthNullValue()
	} else {
		if err := _m.ReadInt32(_r, &s.SnapshotLength); err != nil {
			return err
		}
	}
	if actingVersion > s.SbeSchemaVersion() && blockLength > s.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-s.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := s.RangeCheck(actingVersion, s.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (s *SnapshotIndexEntry) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if s.TimestampInActingVersion(actingVersion) {
		if s.Timestamp < s.TimestampMinValue() || s.Timestamp > s.TimestampMaxValue() {
			return fmt.Errorf("Range check failed on s.Timestamp (%v < %v > %v)", s.TimestampMinValue(), s.Timestamp, s.TimestampMaxValue())
		}
	}
	if s.LogPositionInActingVersion(actingVersion) {
		if s.LogPosition < s.LogPositionMinValue() || s.LogPosition > s.LogPositionMaxValue() {
			return fmt.Errorf("Range check failed on s.LogPosition (%v < %v > %v)", s.LogPositionMinValue(), s.LogPosition, s.LogPositionMaxValue())
		}
	}
	if s.EventChannelPositionInActingVersion(actingVersion) {
		if s.EventChannelPosition < s.EventChannelPositionMinValue() || s.EventChannelPosition > s.EventChannelPositionMaxValue() {
			return fmt.Errorf("Range check failed on s.EventChannelPosition (%v < %v > %v)", s.EventChannelPositionMinValue(), s.EventChannelPosition, s.EventChannelPositionMaxValue())
		}
	}
	if s.SnapshotPositionInActingVersion(actingVersion) {
		if s.SnapshotPosition < s.SnapshotPositionMinValue() || s.SnapshotPosition > s.SnapshotPositionMaxValue() {
			return fmt.Errorf("Range check failed on s.SnapshotPosition (%v < %v > %v)", s.SnapshotPositionMinValue(), s.SnapshotPosition, s.SnapshotPositionMaxValue())
		}
	}
	if s.SnapshotLengthInActingVersion(actingVersion) {
		if s.SnapshotLength < s.SnapshotLengthMinValue() || s.SnapshotLength > s.SnapshotLengthMaxValue() {
			return fmt.Errorf("Range check failed on s.SnapshotLength (%v < %v > %v)", s.SnapshotLengthMinValue(), s.SnapshotLength, s.SnapshotLengthMaxValue())
		}
	}
	return nil
}

func SnapshotIndexEntryInit(s *SnapshotIndexEntry) {
	return
}

func (*SnapshotIndexEntry) SbeBlockLength() (blockLength uint16) {
	return 36
}

func (*SnapshotIndexEntry) SbeTemplateId() (templateId uint16) {
	return 1110
}

func (*SnapshotIndexEntry) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*SnapshotIndexEntry) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*SnapshotIndexEntry) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*SnapshotIndexEntry) TimestampId() uint16 {
	return 1
}

func (*SnapshotIndexEntry) TimestampSinceVersion() uint16 {
	return 0
}

func (s *SnapshotIndexEntry) TimestampInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TimestampSinceVersion()
}

func (*SnapshotIndexEntry) TimestampDeprecated() uint16 {
	return 0
}

func (*SnapshotIndexEntry) TimestampMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*SnapshotIndexEntry) TimestampMinValue() int64 {
	return math.MinInt64 + 1
}

func (*SnapshotIndexEntry) TimestampMaxValue() int64 {
	return math.MaxInt64
}

func (*SnapshotIndexEntry) TimestampNullValue() int64 {
	return math.MinInt64
}

func (*SnapshotIndexEntry) LogPositionId() uint16 {
	return 2
}

func (*SnapshotIndexEntry) LogPositionSinceVersion() uint16 {
	return 0
}

func (s *SnapshotIndexEntry) LogPositionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.LogPositionSinceVersion()
}

func (*SnapshotIndexEntry) LogPositionDeprecated() uint16 {
	return 0
}

func (*SnapshotIndexEntry) LogPositionMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*SnapshotIndexEntry) LogPositionMinValue() int64 {
	return math.MinInt64 + 1
}

func (*SnapshotIndexEntry) LogPositionMaxValue() int64 {
	return math.MaxInt64
}

func (*SnapshotIndexEntry) LogPositionNullValue() int64 {
	return math.MinInt64
}

func (*SnapshotIndexEntry) EventChannelPositionId() uint16 {
	return 3
}

func (*SnapshotIndexEntry) EventChannelPositionSinceVersion() uint16 {
	return 0
}

func (s *SnapshotIndexEntry) EventChannelPositionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.EventChannelPositionSinceVersion()
}

func (*SnapshotIndexEntry) EventChannelPositionDeprecated() uint16 {
	return 0
}

func (*SnapshotIndexEntry) EventChannelPositionMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*SnapshotIndexEntry) EventChannelPositionMinValue() int64 {
	return math.MinInt64 + 1
}

func (*SnapshotIndexEntry) EventChannelPositionMaxValue() int64 {
	return math.MaxInt64
}

func (*SnapshotIndexEntry) EventChannelPositionNullValue() int64 {
	return math.MinInt64
}

func (*SnapshotIndexEntry) SnapshotPositionId() uint16 {
	return 4
}

func (*SnapshotIndexEntry) SnapshotPositionSinceVersion() uint16 {
	return 0
}

func (s *SnapshotIndexEntry) SnapshotPositionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SnapshotPositionSinceVersion()
}

func (*SnapshotIndexEntry) SnapshotPositionDeprecated() uint16 {
	return 0
}

func (*SnapshotIndexEntry) SnapshotPositionMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*SnapshotIndexEntry) SnapshotPositionMinValue() int64 {
	return math.MinInt64 + 1
}

func (*SnapshotIndexEntry) SnapshotPositionMaxValue() int64 {
	return math.MaxInt64
}

func (*SnapshotIndexEntry) SnapshotPositionNullValue() int64 {
	return math.MinInt64
}

func (*SnapshotIndexEntry) SnapshotLengthId() uint16 {
	return 5
}

func (*SnapshotIndexEntry) SnapshotLengthSinceVersion() uint16 {
	return 0
}

func (s *SnapshotIndexEntry) SnapshotLengthInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SnapshotLengthSinceVersion()
}

func (*SnapshotIndexEntry) SnapshotLengthDeprecated() uint16 {
	return 0
}

func (*SnapshotIndexEntry) SnapshotLengthMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*SnapshotIndexEntry) SnapshotLengthMinValue() int32 {
	return math.MinInt32 + 1
}

func (*SnapshotIndexEntry) SnapshotLengthMaxValue() int32 {
	return math.MaxInt32
}

func (*SnapshotIndexEntry) SnapshotLengthNullValue() int32 {
	return math.MinInt32
}
