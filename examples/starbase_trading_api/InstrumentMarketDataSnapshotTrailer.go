// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type InstrumentMarketDataSnapshotTrailer struct {
	LogPosition  int64
	InstrumentId int64
	ActiveCount  int32
}

func (i *InstrumentMarketDataSnapshotTrailer) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := i.RangeCheck(i.SbeSchemaVersion(), i.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, i.LogPosition); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, i.InstrumentId); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, i.ActiveCount); err != nil {
		return err
	}
	return nil
}

func (i *InstrumentMarketDataSnapshotTrailer) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !i.LogPositionInActingVersion(actingVersion) {
		i.LogPosition = i.LogPositionNullValue()
	} else {
		if err := _m.ReadInt64(_r, &i.LogPosition); err != nil {
			return err
		}
	}
	if !i.InstrumentIdInActingVersion(actingVersion) {
		i.InstrumentId = i.InstrumentIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &i.InstrumentId); err != nil {
			return err
		}
	}
	if !i.ActiveCountInActingVersion(actingVersion) {
		i.ActiveCount = i.ActiveCountNullValue()
	} else {
		if err := _m.ReadInt32(_r, &i.ActiveCount); err != nil {
			return err
		}
	}
	if actingVersion > i.SbeSchemaVersion() && blockLength > i.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-i.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := i.RangeCheck(actingVersion, i.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (i *InstrumentMarketDataSnapshotTrailer) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if i.LogPositionInActingVersion(actingVersion) {
		if i.LogPosition < i.LogPositionMinValue() || i.LogPosition > i.LogPositionMaxValue() {
			return fmt.Errorf("Range check failed on i.LogPosition (%v < %v > %v)", i.LogPositionMinValue(), i.LogPosition, i.LogPositionMaxValue())
		}
	}
	if i.InstrumentIdInActingVersion(actingVersion) {
		if i.InstrumentId < i.InstrumentIdMinValue() || i.InstrumentId > i.InstrumentIdMaxValue() {
			return fmt.Errorf("Range check failed on i.InstrumentId (%v < %v > %v)", i.InstrumentIdMinValue(), i.InstrumentId, i.InstrumentIdMaxValue())
		}
	}
	if i.ActiveCountInActingVersion(actingVersion) {
		if i.ActiveCount < i.ActiveCountMinValue() || i.ActiveCount > i.ActiveCountMaxValue() {
			return fmt.Errorf("Range check failed on i.ActiveCount (%v < %v > %v)", i.ActiveCountMinValue(), i.ActiveCount, i.ActiveCountMaxValue())
		}
	}
	return nil
}

func InstrumentMarketDataSnapshotTrailerInit(i *InstrumentMarketDataSnapshotTrailer) {
	return
}

func (*InstrumentMarketDataSnapshotTrailer) SbeBlockLength() (blockLength uint16) {
	return 20
}

func (*InstrumentMarketDataSnapshotTrailer) SbeTemplateId() (templateId uint16) {
	return 2115
}

func (*InstrumentMarketDataSnapshotTrailer) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*InstrumentMarketDataSnapshotTrailer) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*InstrumentMarketDataSnapshotTrailer) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*InstrumentMarketDataSnapshotTrailer) LogPositionId() uint16 {
	return 1
}

func (*InstrumentMarketDataSnapshotTrailer) LogPositionSinceVersion() uint16 {
	return 0
}

func (i *InstrumentMarketDataSnapshotTrailer) LogPositionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.LogPositionSinceVersion()
}

func (*InstrumentMarketDataSnapshotTrailer) LogPositionDeprecated() uint16 {
	return 0
}

func (*InstrumentMarketDataSnapshotTrailer) LogPositionMetaAttribute(meta int) string {
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

func (*InstrumentMarketDataSnapshotTrailer) LogPositionMinValue() int64 {
	return math.MinInt64 + 1
}

func (*InstrumentMarketDataSnapshotTrailer) LogPositionMaxValue() int64 {
	return math.MaxInt64
}

func (*InstrumentMarketDataSnapshotTrailer) LogPositionNullValue() int64 {
	return math.MinInt64
}

func (*InstrumentMarketDataSnapshotTrailer) InstrumentIdId() uint16 {
	return 2
}

func (*InstrumentMarketDataSnapshotTrailer) InstrumentIdSinceVersion() uint16 {
	return 0
}

func (i *InstrumentMarketDataSnapshotTrailer) InstrumentIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.InstrumentIdSinceVersion()
}

func (*InstrumentMarketDataSnapshotTrailer) InstrumentIdDeprecated() uint16 {
	return 0
}

func (*InstrumentMarketDataSnapshotTrailer) InstrumentIdMetaAttribute(meta int) string {
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

func (*InstrumentMarketDataSnapshotTrailer) InstrumentIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*InstrumentMarketDataSnapshotTrailer) InstrumentIdMaxValue() int64 {
	return math.MaxInt64
}

func (*InstrumentMarketDataSnapshotTrailer) InstrumentIdNullValue() int64 {
	return math.MinInt64
}

func (*InstrumentMarketDataSnapshotTrailer) ActiveCountId() uint16 {
	return 3
}

func (*InstrumentMarketDataSnapshotTrailer) ActiveCountSinceVersion() uint16 {
	return 0
}

func (i *InstrumentMarketDataSnapshotTrailer) ActiveCountInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.ActiveCountSinceVersion()
}

func (*InstrumentMarketDataSnapshotTrailer) ActiveCountDeprecated() uint16 {
	return 0
}

func (*InstrumentMarketDataSnapshotTrailer) ActiveCountMetaAttribute(meta int) string {
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

func (*InstrumentMarketDataSnapshotTrailer) ActiveCountMinValue() int32 {
	return math.MinInt32 + 1
}

func (*InstrumentMarketDataSnapshotTrailer) ActiveCountMaxValue() int32 {
	return math.MaxInt32
}

func (*InstrumentMarketDataSnapshotTrailer) ActiveCountNullValue() int32 {
	return math.MinInt32
}
