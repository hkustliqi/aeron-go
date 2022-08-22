// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type EngPerpetualFutureSettlementSchedulerStateSnapshot struct {
	SettlementTimerId    int64
	NextSettlement       int64
	SettlementIntervalMs int64
	InstrumentId         int64
	Symbol               [80]byte
}

func (e *EngPerpetualFutureSettlementSchedulerStateSnapshot) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := e.RangeCheck(e.SbeSchemaVersion(), e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, e.SettlementTimerId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.NextSettlement); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.SettlementIntervalMs); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.InstrumentId); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.Symbol[:]); err != nil {
		return err
	}
	return nil
}

func (e *EngPerpetualFutureSettlementSchedulerStateSnapshot) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !e.SettlementTimerIdInActingVersion(actingVersion) {
		e.SettlementTimerId = e.SettlementTimerIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.SettlementTimerId); err != nil {
			return err
		}
	}
	if !e.NextSettlementInActingVersion(actingVersion) {
		e.NextSettlement = e.NextSettlementNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.NextSettlement); err != nil {
			return err
		}
	}
	if !e.SettlementIntervalMsInActingVersion(actingVersion) {
		e.SettlementIntervalMs = e.SettlementIntervalMsNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.SettlementIntervalMs); err != nil {
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
	if !e.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			e.Symbol[idx] = e.SymbolNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, e.Symbol[:]); err != nil {
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

func (e *EngPerpetualFutureSettlementSchedulerStateSnapshot) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.SettlementTimerIdInActingVersion(actingVersion) {
		if e.SettlementTimerId < e.SettlementTimerIdMinValue() || e.SettlementTimerId > e.SettlementTimerIdMaxValue() {
			return fmt.Errorf("Range check failed on e.SettlementTimerId (%v < %v > %v)", e.SettlementTimerIdMinValue(), e.SettlementTimerId, e.SettlementTimerIdMaxValue())
		}
	}
	if e.NextSettlementInActingVersion(actingVersion) {
		if e.NextSettlement < e.NextSettlementMinValue() || e.NextSettlement > e.NextSettlementMaxValue() {
			return fmt.Errorf("Range check failed on e.NextSettlement (%v < %v > %v)", e.NextSettlementMinValue(), e.NextSettlement, e.NextSettlementMaxValue())
		}
	}
	if e.SettlementIntervalMsInActingVersion(actingVersion) {
		if e.SettlementIntervalMs < e.SettlementIntervalMsMinValue() || e.SettlementIntervalMs > e.SettlementIntervalMsMaxValue() {
			return fmt.Errorf("Range check failed on e.SettlementIntervalMs (%v < %v > %v)", e.SettlementIntervalMsMinValue(), e.SettlementIntervalMs, e.SettlementIntervalMsMaxValue())
		}
	}
	if e.InstrumentIdInActingVersion(actingVersion) {
		if e.InstrumentId < e.InstrumentIdMinValue() || e.InstrumentId > e.InstrumentIdMaxValue() {
			return fmt.Errorf("Range check failed on e.InstrumentId (%v < %v > %v)", e.InstrumentIdMinValue(), e.InstrumentId, e.InstrumentIdMaxValue())
		}
	}
	if e.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			if e.Symbol[idx] < e.SymbolMinValue() || e.Symbol[idx] > e.SymbolMaxValue() {
				return fmt.Errorf("Range check failed on e.Symbol[%d] (%v < %v > %v)", idx, e.SymbolMinValue(), e.Symbol[idx], e.SymbolMaxValue())
			}
		}
	}
	for idx, ch := range e.Symbol {
		if ch > 127 {
			return fmt.Errorf("e.Symbol[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	return nil
}

func EngPerpetualFutureSettlementSchedulerStateSnapshotInit(e *EngPerpetualFutureSettlementSchedulerStateSnapshot) {
	return
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SbeBlockLength() (blockLength uint16) {
	return 112
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SbeTemplateId() (templateId uint16) {
	return 2103
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SettlementTimerIdId() uint16 {
	return 1
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SettlementTimerIdSinceVersion() uint16 {
	return 0
}

func (e *EngPerpetualFutureSettlementSchedulerStateSnapshot) SettlementTimerIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SettlementTimerIdSinceVersion()
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SettlementTimerIdDeprecated() uint16 {
	return 0
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SettlementTimerIdMetaAttribute(meta int) string {
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

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SettlementTimerIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SettlementTimerIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SettlementTimerIdNullValue() int64 {
	return math.MinInt64
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) NextSettlementId() uint16 {
	return 2
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) NextSettlementSinceVersion() uint16 {
	return 0
}

func (e *EngPerpetualFutureSettlementSchedulerStateSnapshot) NextSettlementInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NextSettlementSinceVersion()
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) NextSettlementDeprecated() uint16 {
	return 0
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) NextSettlementMetaAttribute(meta int) string {
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

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) NextSettlementMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) NextSettlementMaxValue() int64 {
	return math.MaxInt64
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) NextSettlementNullValue() int64 {
	return math.MinInt64
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SettlementIntervalMsId() uint16 {
	return 3
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SettlementIntervalMsSinceVersion() uint16 {
	return 0
}

func (e *EngPerpetualFutureSettlementSchedulerStateSnapshot) SettlementIntervalMsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SettlementIntervalMsSinceVersion()
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SettlementIntervalMsDeprecated() uint16 {
	return 0
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SettlementIntervalMsMetaAttribute(meta int) string {
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

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SettlementIntervalMsMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SettlementIntervalMsMaxValue() int64 {
	return math.MaxInt64
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SettlementIntervalMsNullValue() int64 {
	return math.MinInt64
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) InstrumentIdId() uint16 {
	return 4
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) InstrumentIdSinceVersion() uint16 {
	return 0
}

func (e *EngPerpetualFutureSettlementSchedulerStateSnapshot) InstrumentIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.InstrumentIdSinceVersion()
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) InstrumentIdDeprecated() uint16 {
	return 0
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) InstrumentIdMetaAttribute(meta int) string {
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

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) InstrumentIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) InstrumentIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) InstrumentIdNullValue() int64 {
	return math.MinInt64
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SymbolId() uint16 {
	return 5
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SymbolSinceVersion() uint16 {
	return 0
}

func (e *EngPerpetualFutureSettlementSchedulerStateSnapshot) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SymbolSinceVersion()
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SymbolDeprecated() uint16 {
	return 0
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SymbolMetaAttribute(meta int) string {
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

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SymbolMinValue() byte {
	return byte(32)
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SymbolMaxValue() byte {
	return byte(126)
}

func (*EngPerpetualFutureSettlementSchedulerStateSnapshot) SymbolNullValue() byte {
	return 0
}

func (e *EngPerpetualFutureSettlementSchedulerStateSnapshot) SymbolCharacterEncoding() string {
	return "US-ASCII"
}
