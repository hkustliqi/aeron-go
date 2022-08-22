// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OmsSnapshotHeader struct {
	Timestamp                                       int64
	OmsEventsPosition                               int64
	LastEventSequenceNumber                         int64
	LastPortfolioBalanceUpdateRequestSequenceNumber int64
	SnowflakeIdFactoryState                         StarbaseIdFactoryStateSnapshot
	ActiveOrderRepoCapacity                         int32
}

func (o *OmsSnapshotHeader) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, o.Timestamp); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.OmsEventsPosition); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.LastEventSequenceNumber); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.LastPortfolioBalanceUpdateRequestSequenceNumber); err != nil {
		return err
	}
	if err := o.SnowflakeIdFactoryState.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, o.ActiveOrderRepoCapacity); err != nil {
		return err
	}
	return nil
}

func (o *OmsSnapshotHeader) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !o.TimestampInActingVersion(actingVersion) {
		o.Timestamp = o.TimestampNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.Timestamp); err != nil {
			return err
		}
	}
	if !o.OmsEventsPositionInActingVersion(actingVersion) {
		o.OmsEventsPosition = o.OmsEventsPositionNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.OmsEventsPosition); err != nil {
			return err
		}
	}
	if !o.LastEventSequenceNumberInActingVersion(actingVersion) {
		o.LastEventSequenceNumber = o.LastEventSequenceNumberNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.LastEventSequenceNumber); err != nil {
			return err
		}
	}
	if !o.LastPortfolioBalanceUpdateRequestSequenceNumberInActingVersion(actingVersion) {
		o.LastPortfolioBalanceUpdateRequestSequenceNumber = o.LastPortfolioBalanceUpdateRequestSequenceNumberNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.LastPortfolioBalanceUpdateRequestSequenceNumber); err != nil {
			return err
		}
	}
	if o.SnowflakeIdFactoryStateInActingVersion(actingVersion) {
		if err := o.SnowflakeIdFactoryState.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.ActiveOrderRepoCapacityInActingVersion(actingVersion) {
		o.ActiveOrderRepoCapacity = o.ActiveOrderRepoCapacityNullValue()
	} else {
		if err := _m.ReadInt32(_r, &o.ActiveOrderRepoCapacity); err != nil {
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

func (o *OmsSnapshotHeader) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.TimestampInActingVersion(actingVersion) {
		if o.Timestamp < o.TimestampMinValue() || o.Timestamp > o.TimestampMaxValue() {
			return fmt.Errorf("Range check failed on o.Timestamp (%v < %v > %v)", o.TimestampMinValue(), o.Timestamp, o.TimestampMaxValue())
		}
	}
	if o.OmsEventsPositionInActingVersion(actingVersion) {
		if o.OmsEventsPosition < o.OmsEventsPositionMinValue() || o.OmsEventsPosition > o.OmsEventsPositionMaxValue() {
			return fmt.Errorf("Range check failed on o.OmsEventsPosition (%v < %v > %v)", o.OmsEventsPositionMinValue(), o.OmsEventsPosition, o.OmsEventsPositionMaxValue())
		}
	}
	if o.LastEventSequenceNumberInActingVersion(actingVersion) {
		if o.LastEventSequenceNumber < o.LastEventSequenceNumberMinValue() || o.LastEventSequenceNumber > o.LastEventSequenceNumberMaxValue() {
			return fmt.Errorf("Range check failed on o.LastEventSequenceNumber (%v < %v > %v)", o.LastEventSequenceNumberMinValue(), o.LastEventSequenceNumber, o.LastEventSequenceNumberMaxValue())
		}
	}
	if o.LastPortfolioBalanceUpdateRequestSequenceNumberInActingVersion(actingVersion) {
		if o.LastPortfolioBalanceUpdateRequestSequenceNumber < o.LastPortfolioBalanceUpdateRequestSequenceNumberMinValue() || o.LastPortfolioBalanceUpdateRequestSequenceNumber > o.LastPortfolioBalanceUpdateRequestSequenceNumberMaxValue() {
			return fmt.Errorf("Range check failed on o.LastPortfolioBalanceUpdateRequestSequenceNumber (%v < %v > %v)", o.LastPortfolioBalanceUpdateRequestSequenceNumberMinValue(), o.LastPortfolioBalanceUpdateRequestSequenceNumber, o.LastPortfolioBalanceUpdateRequestSequenceNumberMaxValue())
		}
	}
	if o.ActiveOrderRepoCapacityInActingVersion(actingVersion) {
		if o.ActiveOrderRepoCapacity < o.ActiveOrderRepoCapacityMinValue() || o.ActiveOrderRepoCapacity > o.ActiveOrderRepoCapacityMaxValue() {
			return fmt.Errorf("Range check failed on o.ActiveOrderRepoCapacity (%v < %v > %v)", o.ActiveOrderRepoCapacityMinValue(), o.ActiveOrderRepoCapacity, o.ActiveOrderRepoCapacityMaxValue())
		}
	}
	return nil
}

func OmsSnapshotHeaderInit(o *OmsSnapshotHeader) {
	return
}

func (*OmsSnapshotHeader) SbeBlockLength() (blockLength uint16) {
	return 48
}

func (*OmsSnapshotHeader) SbeTemplateId() (templateId uint16) {
	return 1100
}

func (*OmsSnapshotHeader) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*OmsSnapshotHeader) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*OmsSnapshotHeader) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*OmsSnapshotHeader) TimestampId() uint16 {
	return 1
}

func (*OmsSnapshotHeader) TimestampSinceVersion() uint16 {
	return 0
}

func (o *OmsSnapshotHeader) TimestampInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TimestampSinceVersion()
}

func (*OmsSnapshotHeader) TimestampDeprecated() uint16 {
	return 0
}

func (*OmsSnapshotHeader) TimestampMetaAttribute(meta int) string {
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

func (*OmsSnapshotHeader) TimestampMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsSnapshotHeader) TimestampMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsSnapshotHeader) TimestampNullValue() int64 {
	return math.MinInt64
}

func (*OmsSnapshotHeader) OmsEventsPositionId() uint16 {
	return 2
}

func (*OmsSnapshotHeader) OmsEventsPositionSinceVersion() uint16 {
	return 0
}

func (o *OmsSnapshotHeader) OmsEventsPositionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OmsEventsPositionSinceVersion()
}

func (*OmsSnapshotHeader) OmsEventsPositionDeprecated() uint16 {
	return 0
}

func (*OmsSnapshotHeader) OmsEventsPositionMetaAttribute(meta int) string {
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

func (*OmsSnapshotHeader) OmsEventsPositionMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsSnapshotHeader) OmsEventsPositionMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsSnapshotHeader) OmsEventsPositionNullValue() int64 {
	return math.MinInt64
}

func (*OmsSnapshotHeader) LastEventSequenceNumberId() uint16 {
	return 3
}

func (*OmsSnapshotHeader) LastEventSequenceNumberSinceVersion() uint16 {
	return 0
}

func (o *OmsSnapshotHeader) LastEventSequenceNumberInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LastEventSequenceNumberSinceVersion()
}

func (*OmsSnapshotHeader) LastEventSequenceNumberDeprecated() uint16 {
	return 0
}

func (*OmsSnapshotHeader) LastEventSequenceNumberMetaAttribute(meta int) string {
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

func (*OmsSnapshotHeader) LastEventSequenceNumberMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsSnapshotHeader) LastEventSequenceNumberMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsSnapshotHeader) LastEventSequenceNumberNullValue() int64 {
	return math.MinInt64
}

func (*OmsSnapshotHeader) LastPortfolioBalanceUpdateRequestSequenceNumberId() uint16 {
	return 4
}

func (*OmsSnapshotHeader) LastPortfolioBalanceUpdateRequestSequenceNumberSinceVersion() uint16 {
	return 0
}

func (o *OmsSnapshotHeader) LastPortfolioBalanceUpdateRequestSequenceNumberInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LastPortfolioBalanceUpdateRequestSequenceNumberSinceVersion()
}

func (*OmsSnapshotHeader) LastPortfolioBalanceUpdateRequestSequenceNumberDeprecated() uint16 {
	return 0
}

func (*OmsSnapshotHeader) LastPortfolioBalanceUpdateRequestSequenceNumberMetaAttribute(meta int) string {
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

func (*OmsSnapshotHeader) LastPortfolioBalanceUpdateRequestSequenceNumberMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsSnapshotHeader) LastPortfolioBalanceUpdateRequestSequenceNumberMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsSnapshotHeader) LastPortfolioBalanceUpdateRequestSequenceNumberNullValue() int64 {
	return math.MinInt64
}

func (*OmsSnapshotHeader) SnowflakeIdFactoryStateId() uint16 {
	return 5
}

func (*OmsSnapshotHeader) SnowflakeIdFactoryStateSinceVersion() uint16 {
	return 0
}

func (o *OmsSnapshotHeader) SnowflakeIdFactoryStateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SnowflakeIdFactoryStateSinceVersion()
}

func (*OmsSnapshotHeader) SnowflakeIdFactoryStateDeprecated() uint16 {
	return 0
}

func (*OmsSnapshotHeader) SnowflakeIdFactoryStateMetaAttribute(meta int) string {
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

func (*OmsSnapshotHeader) ActiveOrderRepoCapacityId() uint16 {
	return 6
}

func (*OmsSnapshotHeader) ActiveOrderRepoCapacitySinceVersion() uint16 {
	return 0
}

func (o *OmsSnapshotHeader) ActiveOrderRepoCapacityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ActiveOrderRepoCapacitySinceVersion()
}

func (*OmsSnapshotHeader) ActiveOrderRepoCapacityDeprecated() uint16 {
	return 0
}

func (*OmsSnapshotHeader) ActiveOrderRepoCapacityMetaAttribute(meta int) string {
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

func (*OmsSnapshotHeader) ActiveOrderRepoCapacityMinValue() int32 {
	return math.MinInt32 + 1
}

func (*OmsSnapshotHeader) ActiveOrderRepoCapacityMaxValue() int32 {
	return math.MaxInt32
}

func (*OmsSnapshotHeader) ActiveOrderRepoCapacityNullValue() int32 {
	return math.MinInt32
}
