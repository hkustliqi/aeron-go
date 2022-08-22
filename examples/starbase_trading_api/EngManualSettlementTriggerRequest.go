// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type EngManualSettlementTriggerRequest struct {
	CorrelationId int64
	InstrumentId  int64
}

func (e *EngManualSettlementTriggerRequest) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := e.RangeCheck(e.SbeSchemaVersion(), e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, e.CorrelationId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.InstrumentId); err != nil {
		return err
	}
	return nil
}

func (e *EngManualSettlementTriggerRequest) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !e.CorrelationIdInActingVersion(actingVersion) {
		e.CorrelationId = e.CorrelationIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.CorrelationId); err != nil {
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

func (e *EngManualSettlementTriggerRequest) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.CorrelationIdInActingVersion(actingVersion) {
		if e.CorrelationId < e.CorrelationIdMinValue() || e.CorrelationId > e.CorrelationIdMaxValue() {
			return fmt.Errorf("Range check failed on e.CorrelationId (%v < %v > %v)", e.CorrelationIdMinValue(), e.CorrelationId, e.CorrelationIdMaxValue())
		}
	}
	if e.InstrumentIdInActingVersion(actingVersion) {
		if e.InstrumentId < e.InstrumentIdMinValue() || e.InstrumentId > e.InstrumentIdMaxValue() {
			return fmt.Errorf("Range check failed on e.InstrumentId (%v < %v > %v)", e.InstrumentIdMinValue(), e.InstrumentId, e.InstrumentIdMaxValue())
		}
	}
	return nil
}

func EngManualSettlementTriggerRequestInit(e *EngManualSettlementTriggerRequest) {
	return
}

func (*EngManualSettlementTriggerRequest) SbeBlockLength() (blockLength uint16) {
	return 16
}

func (*EngManualSettlementTriggerRequest) SbeTemplateId() (templateId uint16) {
	return 2015
}

func (*EngManualSettlementTriggerRequest) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*EngManualSettlementTriggerRequest) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*EngManualSettlementTriggerRequest) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*EngManualSettlementTriggerRequest) CorrelationIdId() uint16 {
	return 1
}

func (*EngManualSettlementTriggerRequest) CorrelationIdSinceVersion() uint16 {
	return 0
}

func (e *EngManualSettlementTriggerRequest) CorrelationIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CorrelationIdSinceVersion()
}

func (*EngManualSettlementTriggerRequest) CorrelationIdDeprecated() uint16 {
	return 0
}

func (*EngManualSettlementTriggerRequest) CorrelationIdMetaAttribute(meta int) string {
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

func (*EngManualSettlementTriggerRequest) CorrelationIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngManualSettlementTriggerRequest) CorrelationIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngManualSettlementTriggerRequest) CorrelationIdNullValue() int64 {
	return math.MinInt64
}

func (*EngManualSettlementTriggerRequest) InstrumentIdId() uint16 {
	return 2
}

func (*EngManualSettlementTriggerRequest) InstrumentIdSinceVersion() uint16 {
	return 0
}

func (e *EngManualSettlementTriggerRequest) InstrumentIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.InstrumentIdSinceVersion()
}

func (*EngManualSettlementTriggerRequest) InstrumentIdDeprecated() uint16 {
	return 0
}

func (*EngManualSettlementTriggerRequest) InstrumentIdMetaAttribute(meta int) string {
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

func (*EngManualSettlementTriggerRequest) InstrumentIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngManualSettlementTriggerRequest) InstrumentIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngManualSettlementTriggerRequest) InstrumentIdNullValue() int64 {
	return math.MinInt64
}
