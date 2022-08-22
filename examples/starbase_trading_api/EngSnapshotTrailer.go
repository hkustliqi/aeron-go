// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type EngSnapshotTrailer struct {
	Timestamp       int64
	LastMessageTime int64
}

func (e *EngSnapshotTrailer) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := e.RangeCheck(e.SbeSchemaVersion(), e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, e.Timestamp); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.LastMessageTime); err != nil {
		return err
	}
	return nil
}

func (e *EngSnapshotTrailer) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !e.TimestampInActingVersion(actingVersion) {
		e.Timestamp = e.TimestampNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.Timestamp); err != nil {
			return err
		}
	}
	if !e.LastMessageTimeInActingVersion(actingVersion) {
		e.LastMessageTime = e.LastMessageTimeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.LastMessageTime); err != nil {
			return err
		}
	}
	if actingVersion > e.SbeSchemaVersion() && blockLength > e.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-e.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := e.RangeCheck(actingVersion, e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (e *EngSnapshotTrailer) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.TimestampInActingVersion(actingVersion) {
		if e.Timestamp < e.TimestampMinValue() || e.Timestamp > e.TimestampMaxValue() {
			return fmt.Errorf("Range check failed on e.Timestamp (%v < %v > %v)", e.TimestampMinValue(), e.Timestamp, e.TimestampMaxValue())
		}
	}
	if e.LastMessageTimeInActingVersion(actingVersion) {
		if e.LastMessageTime < e.LastMessageTimeMinValue() || e.LastMessageTime > e.LastMessageTimeMaxValue() {
			return fmt.Errorf("Range check failed on e.LastMessageTime (%v < %v > %v)", e.LastMessageTimeMinValue(), e.LastMessageTime, e.LastMessageTimeMaxValue())
		}
	}
	return nil
}

func EngSnapshotTrailerInit(e *EngSnapshotTrailer) {
	return
}

func (*EngSnapshotTrailer) SbeBlockLength() (blockLength uint16) {
	return 16
}

func (*EngSnapshotTrailer) SbeTemplateId() (templateId uint16) {
	return 2101
}

func (*EngSnapshotTrailer) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*EngSnapshotTrailer) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*EngSnapshotTrailer) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*EngSnapshotTrailer) TimestampId() uint16 {
	return 1
}

func (*EngSnapshotTrailer) TimestampSinceVersion() uint16 {
	return 0
}

func (e *EngSnapshotTrailer) TimestampInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TimestampSinceVersion()
}

func (*EngSnapshotTrailer) TimestampDeprecated() uint16 {
	return 0
}

func (*EngSnapshotTrailer) TimestampMetaAttribute(meta int) string {
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

func (*EngSnapshotTrailer) TimestampMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngSnapshotTrailer) TimestampMaxValue() int64 {
	return math.MaxInt64
}

func (*EngSnapshotTrailer) TimestampNullValue() int64 {
	return math.MinInt64
}

func (*EngSnapshotTrailer) LastMessageTimeId() uint16 {
	return 2
}

func (*EngSnapshotTrailer) LastMessageTimeSinceVersion() uint16 {
	return 0
}

func (e *EngSnapshotTrailer) LastMessageTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LastMessageTimeSinceVersion()
}

func (*EngSnapshotTrailer) LastMessageTimeDeprecated() uint16 {
	return 0
}

func (*EngSnapshotTrailer) LastMessageTimeMetaAttribute(meta int) string {
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

func (*EngSnapshotTrailer) LastMessageTimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngSnapshotTrailer) LastMessageTimeMaxValue() int64 {
	return math.MaxInt64
}

func (*EngSnapshotTrailer) LastMessageTimeNullValue() int64 {
	return math.MinInt64
}
