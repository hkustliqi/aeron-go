// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type EngSettlementTrigger struct {
	Timestamp    int64
	InstrumentId int64
	FundingRate  int64
}

func (e *EngSettlementTrigger) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := e.RangeCheck(e.SbeSchemaVersion(), e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, e.Timestamp); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.InstrumentId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.FundingRate); err != nil {
		return err
	}
	return nil
}

func (e *EngSettlementTrigger) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !e.TimestampInActingVersion(actingVersion) {
		e.Timestamp = e.TimestampNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.Timestamp); err != nil {
			return err
		}
	}
	if !e.InstrumentIdInActingVersion(actingVersion) {
		e.InstrumentId = e.InstrumentIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.InstrumentId); err != nil {
			return err
		}
	}
	if !e.FundingRateInActingVersion(actingVersion) {
		e.FundingRate = e.FundingRateNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.FundingRate); err != nil {
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

func (e *EngSettlementTrigger) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.TimestampInActingVersion(actingVersion) {
		if e.Timestamp < e.TimestampMinValue() || e.Timestamp > e.TimestampMaxValue() {
			return fmt.Errorf("Range check failed on e.Timestamp (%v < %v > %v)", e.TimestampMinValue(), e.Timestamp, e.TimestampMaxValue())
		}
	}
	if e.InstrumentIdInActingVersion(actingVersion) {
		if e.InstrumentId < e.InstrumentIdMinValue() || e.InstrumentId > e.InstrumentIdMaxValue() {
			return fmt.Errorf("Range check failed on e.InstrumentId (%v < %v > %v)", e.InstrumentIdMinValue(), e.InstrumentId, e.InstrumentIdMaxValue())
		}
	}
	if e.FundingRateInActingVersion(actingVersion) {
		if e.FundingRate < e.FundingRateMinValue() || e.FundingRate > e.FundingRateMaxValue() {
			return fmt.Errorf("Range check failed on e.FundingRate (%v < %v > %v)", e.FundingRateMinValue(), e.FundingRate, e.FundingRateMaxValue())
		}
	}
	return nil
}

func EngSettlementTriggerInit(e *EngSettlementTrigger) {
	return
}

func (*EngSettlementTrigger) SbeBlockLength() (blockLength uint16) {
	return 24
}

func (*EngSettlementTrigger) SbeTemplateId() (templateId uint16) {
	return 2014
}

func (*EngSettlementTrigger) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*EngSettlementTrigger) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*EngSettlementTrigger) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*EngSettlementTrigger) TimestampId() uint16 {
	return 1
}

func (*EngSettlementTrigger) TimestampSinceVersion() uint16 {
	return 0
}

func (e *EngSettlementTrigger) TimestampInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TimestampSinceVersion()
}

func (*EngSettlementTrigger) TimestampDeprecated() uint16 {
	return 0
}

func (*EngSettlementTrigger) TimestampMetaAttribute(meta int) string {
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

func (*EngSettlementTrigger) TimestampMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngSettlementTrigger) TimestampMaxValue() int64 {
	return math.MaxInt64
}

func (*EngSettlementTrigger) TimestampNullValue() int64 {
	return math.MinInt64
}

func (*EngSettlementTrigger) InstrumentIdId() uint16 {
	return 2
}

func (*EngSettlementTrigger) InstrumentIdSinceVersion() uint16 {
	return 0
}

func (e *EngSettlementTrigger) InstrumentIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.InstrumentIdSinceVersion()
}

func (*EngSettlementTrigger) InstrumentIdDeprecated() uint16 {
	return 0
}

func (*EngSettlementTrigger) InstrumentIdMetaAttribute(meta int) string {
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

func (*EngSettlementTrigger) InstrumentIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngSettlementTrigger) InstrumentIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngSettlementTrigger) InstrumentIdNullValue() int64 {
	return math.MinInt64
}

func (*EngSettlementTrigger) FundingRateId() uint16 {
	return 3
}

func (*EngSettlementTrigger) FundingRateSinceVersion() uint16 {
	return 0
}

func (e *EngSettlementTrigger) FundingRateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FundingRateSinceVersion()
}

func (*EngSettlementTrigger) FundingRateDeprecated() uint16 {
	return 0
}

func (*EngSettlementTrigger) FundingRateMetaAttribute(meta int) string {
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

func (*EngSettlementTrigger) FundingRateMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngSettlementTrigger) FundingRateMaxValue() int64 {
	return math.MaxInt64
}

func (*EngSettlementTrigger) FundingRateNullValue() int64 {
	return math.MinInt64
}
