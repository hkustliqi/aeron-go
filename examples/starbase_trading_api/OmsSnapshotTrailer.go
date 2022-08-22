// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OmsSnapshotTrailer struct {
	Timestamp       int64
	LastMessageTime int64
}

func (o *OmsSnapshotTrailer) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, o.Timestamp); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.LastMessageTime); err != nil {
		return err
	}
	return nil
}

func (o *OmsSnapshotTrailer) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !o.TimestampInActingVersion(actingVersion) {
		o.Timestamp = o.TimestampNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.Timestamp); err != nil {
			return err
		}
	}
	if !o.LastMessageTimeInActingVersion(actingVersion) {
		o.LastMessageTime = o.LastMessageTimeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.LastMessageTime); err != nil {
			return err
		}
	}
	if actingVersion > o.SbeSchemaVersion() && blockLength > o.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-o.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := o.RangeCheck(actingVersion, o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (o *OmsSnapshotTrailer) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.TimestampInActingVersion(actingVersion) {
		if o.Timestamp < o.TimestampMinValue() || o.Timestamp > o.TimestampMaxValue() {
			return fmt.Errorf("Range check failed on o.Timestamp (%v < %v > %v)", o.TimestampMinValue(), o.Timestamp, o.TimestampMaxValue())
		}
	}
	if o.LastMessageTimeInActingVersion(actingVersion) {
		if o.LastMessageTime < o.LastMessageTimeMinValue() || o.LastMessageTime > o.LastMessageTimeMaxValue() {
			return fmt.Errorf("Range check failed on o.LastMessageTime (%v < %v > %v)", o.LastMessageTimeMinValue(), o.LastMessageTime, o.LastMessageTimeMaxValue())
		}
	}
	return nil
}

func OmsSnapshotTrailerInit(o *OmsSnapshotTrailer) {
	return
}

func (*OmsSnapshotTrailer) SbeBlockLength() (blockLength uint16) {
	return 16
}

func (*OmsSnapshotTrailer) SbeTemplateId() (templateId uint16) {
	return 1101
}

func (*OmsSnapshotTrailer) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*OmsSnapshotTrailer) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*OmsSnapshotTrailer) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*OmsSnapshotTrailer) TimestampId() uint16 {
	return 1
}

func (*OmsSnapshotTrailer) TimestampSinceVersion() uint16 {
	return 0
}

func (o *OmsSnapshotTrailer) TimestampInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TimestampSinceVersion()
}

func (*OmsSnapshotTrailer) TimestampDeprecated() uint16 {
	return 0
}

func (*OmsSnapshotTrailer) TimestampMetaAttribute(meta int) string {
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

func (*OmsSnapshotTrailer) TimestampMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsSnapshotTrailer) TimestampMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsSnapshotTrailer) TimestampNullValue() int64 {
	return math.MinInt64
}

func (*OmsSnapshotTrailer) LastMessageTimeId() uint16 {
	return 2
}

func (*OmsSnapshotTrailer) LastMessageTimeSinceVersion() uint16 {
	return 0
}

func (o *OmsSnapshotTrailer) LastMessageTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LastMessageTimeSinceVersion()
}

func (*OmsSnapshotTrailer) LastMessageTimeDeprecated() uint16 {
	return 0
}

func (*OmsSnapshotTrailer) LastMessageTimeMetaAttribute(meta int) string {
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

func (*OmsSnapshotTrailer) LastMessageTimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsSnapshotTrailer) LastMessageTimeMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsSnapshotTrailer) LastMessageTimeNullValue() int64 {
	return math.MinInt64
}
