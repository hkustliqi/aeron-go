// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type TradingEventSnapshotTrailer struct {
	Timestamp   int64
	LogPosition int64
}

func (t *TradingEventSnapshotTrailer) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := t.RangeCheck(t.SbeSchemaVersion(), t.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, t.Timestamp); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, t.LogPosition); err != nil {
		return err
	}
	return nil
}

func (t *TradingEventSnapshotTrailer) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !t.TimestampInActingVersion(actingVersion) {
		t.Timestamp = t.TimestampNullValue()
	} else {
		if err := _m.ReadInt64(_r, &t.Timestamp); err != nil {
			return err
		}
	}
	if !t.LogPositionInActingVersion(actingVersion) {
		t.LogPosition = t.LogPositionNullValue()
	} else {
		if err := _m.ReadInt64(_r, &t.LogPosition); err != nil {
			return err
		}
	}
	if actingVersion > t.SbeSchemaVersion() && blockLength > t.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-t.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := t.RangeCheck(actingVersion, t.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (t *TradingEventSnapshotTrailer) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if t.TimestampInActingVersion(actingVersion) {
		if t.Timestamp < t.TimestampMinValue() || t.Timestamp > t.TimestampMaxValue() {
			return fmt.Errorf("Range check failed on t.Timestamp (%v < %v > %v)", t.TimestampMinValue(), t.Timestamp, t.TimestampMaxValue())
		}
	}
	if t.LogPositionInActingVersion(actingVersion) {
		if t.LogPosition < t.LogPositionMinValue() || t.LogPosition > t.LogPositionMaxValue() {
			return fmt.Errorf("Range check failed on t.LogPosition (%v < %v > %v)", t.LogPositionMinValue(), t.LogPosition, t.LogPositionMaxValue())
		}
	}
	return nil
}

func TradingEventSnapshotTrailerInit(t *TradingEventSnapshotTrailer) {
	return
}

func (*TradingEventSnapshotTrailer) SbeBlockLength() (blockLength uint16) {
	return 16
}

func (*TradingEventSnapshotTrailer) SbeTemplateId() (templateId uint16) {
	return 1112
}

func (*TradingEventSnapshotTrailer) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*TradingEventSnapshotTrailer) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*TradingEventSnapshotTrailer) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*TradingEventSnapshotTrailer) TimestampId() uint16 {
	return 1
}

func (*TradingEventSnapshotTrailer) TimestampSinceVersion() uint16 {
	return 0
}

func (t *TradingEventSnapshotTrailer) TimestampInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.TimestampSinceVersion()
}

func (*TradingEventSnapshotTrailer) TimestampDeprecated() uint16 {
	return 0
}

func (*TradingEventSnapshotTrailer) TimestampMetaAttribute(meta int) string {
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

func (*TradingEventSnapshotTrailer) TimestampMinValue() int64 {
	return math.MinInt64 + 1
}

func (*TradingEventSnapshotTrailer) TimestampMaxValue() int64 {
	return math.MaxInt64
}

func (*TradingEventSnapshotTrailer) TimestampNullValue() int64 {
	return math.MinInt64
}

func (*TradingEventSnapshotTrailer) LogPositionId() uint16 {
	return 2
}

func (*TradingEventSnapshotTrailer) LogPositionSinceVersion() uint16 {
	return 0
}

func (t *TradingEventSnapshotTrailer) LogPositionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.LogPositionSinceVersion()
}

func (*TradingEventSnapshotTrailer) LogPositionDeprecated() uint16 {
	return 0
}

func (*TradingEventSnapshotTrailer) LogPositionMetaAttribute(meta int) string {
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

func (*TradingEventSnapshotTrailer) LogPositionMinValue() int64 {
	return math.MinInt64 + 1
}

func (*TradingEventSnapshotTrailer) LogPositionMaxValue() int64 {
	return math.MaxInt64
}

func (*TradingEventSnapshotTrailer) LogPositionNullValue() int64 {
	return math.MinInt64
}
