// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type EngSnapshotHeader struct {
	Timestamp               int64
	LastEventSequenceNumber int64
	MarketDataEventPosition int64
	SnowflakeIdFactoryState StarbaseIdFactoryStateSnapshot
	ActiveOrderRepoCapacity int32
}

func (e *EngSnapshotHeader) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := e.RangeCheck(e.SbeSchemaVersion(), e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, e.Timestamp); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.LastEventSequenceNumber); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.MarketDataEventPosition); err != nil {
		return err
	}
	if err := e.SnowflakeIdFactoryState.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, e.ActiveOrderRepoCapacity); err != nil {
		return err
	}
	return nil
}

func (e *EngSnapshotHeader) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !e.TimestampInActingVersion(actingVersion) {
		e.Timestamp = e.TimestampNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.Timestamp); err != nil {
			return err
		}
	}
	if !e.LastEventSequenceNumberInActingVersion(actingVersion) {
		e.LastEventSequenceNumber = e.LastEventSequenceNumberNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.LastEventSequenceNumber); err != nil {
			return err
		}
	}
	if !e.MarketDataEventPositionInActingVersion(actingVersion) {
		e.MarketDataEventPosition = e.MarketDataEventPositionNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.MarketDataEventPosition); err != nil {
			return err
		}
	}
	if e.SnowflakeIdFactoryStateInActingVersion(actingVersion) {
		if err := e.SnowflakeIdFactoryState.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.ActiveOrderRepoCapacityInActingVersion(actingVersion) {
		e.ActiveOrderRepoCapacity = e.ActiveOrderRepoCapacityNullValue()
	} else {
		if err := _m.ReadInt32(_r, &e.ActiveOrderRepoCapacity); err != nil {
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

func (e *EngSnapshotHeader) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.TimestampInActingVersion(actingVersion) {
		if e.Timestamp < e.TimestampMinValue() || e.Timestamp > e.TimestampMaxValue() {
			return fmt.Errorf("Range check failed on e.Timestamp (%v < %v > %v)", e.TimestampMinValue(), e.Timestamp, e.TimestampMaxValue())
		}
	}
	if e.LastEventSequenceNumberInActingVersion(actingVersion) {
		if e.LastEventSequenceNumber < e.LastEventSequenceNumberMinValue() || e.LastEventSequenceNumber > e.LastEventSequenceNumberMaxValue() {
			return fmt.Errorf("Range check failed on e.LastEventSequenceNumber (%v < %v > %v)", e.LastEventSequenceNumberMinValue(), e.LastEventSequenceNumber, e.LastEventSequenceNumberMaxValue())
		}
	}
	if e.MarketDataEventPositionInActingVersion(actingVersion) {
		if e.MarketDataEventPosition < e.MarketDataEventPositionMinValue() || e.MarketDataEventPosition > e.MarketDataEventPositionMaxValue() {
			return fmt.Errorf("Range check failed on e.MarketDataEventPosition (%v < %v > %v)", e.MarketDataEventPositionMinValue(), e.MarketDataEventPosition, e.MarketDataEventPositionMaxValue())
		}
	}
	if e.ActiveOrderRepoCapacityInActingVersion(actingVersion) {
		if e.ActiveOrderRepoCapacity < e.ActiveOrderRepoCapacityMinValue() || e.ActiveOrderRepoCapacity > e.ActiveOrderRepoCapacityMaxValue() {
			return fmt.Errorf("Range check failed on e.ActiveOrderRepoCapacity (%v < %v > %v)", e.ActiveOrderRepoCapacityMinValue(), e.ActiveOrderRepoCapacity, e.ActiveOrderRepoCapacityMaxValue())
		}
	}
	return nil
}

func EngSnapshotHeaderInit(e *EngSnapshotHeader) {
	return
}

func (*EngSnapshotHeader) SbeBlockLength() (blockLength uint16) {
	return 40
}

func (*EngSnapshotHeader) SbeTemplateId() (templateId uint16) {
	return 2100
}

func (*EngSnapshotHeader) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*EngSnapshotHeader) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*EngSnapshotHeader) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*EngSnapshotHeader) TimestampId() uint16 {
	return 1
}

func (*EngSnapshotHeader) TimestampSinceVersion() uint16 {
	return 0
}

func (e *EngSnapshotHeader) TimestampInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TimestampSinceVersion()
}

func (*EngSnapshotHeader) TimestampDeprecated() uint16 {
	return 0
}

func (*EngSnapshotHeader) TimestampMetaAttribute(meta int) string {
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

func (*EngSnapshotHeader) TimestampMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngSnapshotHeader) TimestampMaxValue() int64 {
	return math.MaxInt64
}

func (*EngSnapshotHeader) TimestampNullValue() int64 {
	return math.MinInt64
}

func (*EngSnapshotHeader) LastEventSequenceNumberId() uint16 {
	return 2
}

func (*EngSnapshotHeader) LastEventSequenceNumberSinceVersion() uint16 {
	return 0
}

func (e *EngSnapshotHeader) LastEventSequenceNumberInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LastEventSequenceNumberSinceVersion()
}

func (*EngSnapshotHeader) LastEventSequenceNumberDeprecated() uint16 {
	return 0
}

func (*EngSnapshotHeader) LastEventSequenceNumberMetaAttribute(meta int) string {
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

func (*EngSnapshotHeader) LastEventSequenceNumberMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngSnapshotHeader) LastEventSequenceNumberMaxValue() int64 {
	return math.MaxInt64
}

func (*EngSnapshotHeader) LastEventSequenceNumberNullValue() int64 {
	return math.MinInt64
}

func (*EngSnapshotHeader) MarketDataEventPositionId() uint16 {
	return 3
}

func (*EngSnapshotHeader) MarketDataEventPositionSinceVersion() uint16 {
	return 0
}

func (e *EngSnapshotHeader) MarketDataEventPositionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MarketDataEventPositionSinceVersion()
}

func (*EngSnapshotHeader) MarketDataEventPositionDeprecated() uint16 {
	return 0
}

func (*EngSnapshotHeader) MarketDataEventPositionMetaAttribute(meta int) string {
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

func (*EngSnapshotHeader) MarketDataEventPositionMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngSnapshotHeader) MarketDataEventPositionMaxValue() int64 {
	return math.MaxInt64
}

func (*EngSnapshotHeader) MarketDataEventPositionNullValue() int64 {
	return math.MinInt64
}

func (*EngSnapshotHeader) SnowflakeIdFactoryStateId() uint16 {
	return 4
}

func (*EngSnapshotHeader) SnowflakeIdFactoryStateSinceVersion() uint16 {
	return 0
}

func (e *EngSnapshotHeader) SnowflakeIdFactoryStateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SnowflakeIdFactoryStateSinceVersion()
}

func (*EngSnapshotHeader) SnowflakeIdFactoryStateDeprecated() uint16 {
	return 0
}

func (*EngSnapshotHeader) SnowflakeIdFactoryStateMetaAttribute(meta int) string {
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

func (*EngSnapshotHeader) ActiveOrderRepoCapacityId() uint16 {
	return 5
}

func (*EngSnapshotHeader) ActiveOrderRepoCapacitySinceVersion() uint16 {
	return 0
}

func (e *EngSnapshotHeader) ActiveOrderRepoCapacityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ActiveOrderRepoCapacitySinceVersion()
}

func (*EngSnapshotHeader) ActiveOrderRepoCapacityDeprecated() uint16 {
	return 0
}

func (*EngSnapshotHeader) ActiveOrderRepoCapacityMetaAttribute(meta int) string {
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

func (*EngSnapshotHeader) ActiveOrderRepoCapacityMinValue() int32 {
	return math.MinInt32 + 1
}

func (*EngSnapshotHeader) ActiveOrderRepoCapacityMaxValue() int32 {
	return math.MaxInt32
}

func (*EngSnapshotHeader) ActiveOrderRepoCapacityNullValue() int32 {
	return math.MinInt32
}
