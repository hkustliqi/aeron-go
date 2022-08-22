// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type EngCancelOrder struct {
	RequestHeader    EngUserRequestHeader
	OrderId          int64
	InstrumentId     int64
	PortfolioId      int64
	NewClientOrderId [2]int64
}

func (e *EngCancelOrder) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := e.RangeCheck(e.SbeSchemaVersion(), e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := e.RequestHeader.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.OrderId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.InstrumentId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.PortfolioId); err != nil {
		return err
	}
	for idx := 0; idx < 2; idx++ {
		if err := _m.WriteInt64(_w, e.NewClientOrderId[idx]); err != nil {
			return err
		}
	}
	return nil
}

func (e *EngCancelOrder) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if e.RequestHeaderInActingVersion(actingVersion) {
		if err := e.RequestHeader.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.OrderIdInActingVersion(actingVersion) {
		e.OrderId = e.OrderIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.OrderId); err != nil {
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
	if !e.PortfolioIdInActingVersion(actingVersion) {
		e.PortfolioId = e.PortfolioIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.PortfolioId); err != nil {
			return err
		}
	}
	if !e.NewClientOrderIdInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			e.NewClientOrderId[idx] = e.NewClientOrderIdNullValue()
		}
	} else {
		for idx := 0; idx < 2; idx++ {
			if err := _m.ReadInt64(_r, &e.NewClientOrderId[idx]); err != nil {
				return err
			}
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

func (e *EngCancelOrder) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.OrderIdInActingVersion(actingVersion) {
		if e.OrderId < e.OrderIdMinValue() || e.OrderId > e.OrderIdMaxValue() {
			return fmt.Errorf("Range check failed on e.OrderId (%v < %v > %v)", e.OrderIdMinValue(), e.OrderId, e.OrderIdMaxValue())
		}
	}
	if e.InstrumentIdInActingVersion(actingVersion) {
		if e.InstrumentId < e.InstrumentIdMinValue() || e.InstrumentId > e.InstrumentIdMaxValue() {
			return fmt.Errorf("Range check failed on e.InstrumentId (%v < %v > %v)", e.InstrumentIdMinValue(), e.InstrumentId, e.InstrumentIdMaxValue())
		}
	}
	if e.PortfolioIdInActingVersion(actingVersion) {
		if e.PortfolioId < e.PortfolioIdMinValue() || e.PortfolioId > e.PortfolioIdMaxValue() {
			return fmt.Errorf("Range check failed on e.PortfolioId (%v < %v > %v)", e.PortfolioIdMinValue(), e.PortfolioId, e.PortfolioIdMaxValue())
		}
	}
	if e.NewClientOrderIdInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if e.NewClientOrderId[idx] < e.NewClientOrderIdMinValue() || e.NewClientOrderId[idx] > e.NewClientOrderIdMaxValue() {
				return fmt.Errorf("Range check failed on e.NewClientOrderId[%d] (%v < %v > %v)", idx, e.NewClientOrderIdMinValue(), e.NewClientOrderId[idx], e.NewClientOrderIdMaxValue())
			}
		}
	}
	return nil
}

func EngCancelOrderInit(e *EngCancelOrder) {
	return
}

func (*EngCancelOrder) SbeBlockLength() (blockLength uint16) {
	return 80
}

func (*EngCancelOrder) SbeTemplateId() (templateId uint16) {
	return 2001
}

func (*EngCancelOrder) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*EngCancelOrder) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*EngCancelOrder) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*EngCancelOrder) RequestHeaderId() uint16 {
	return 1
}

func (*EngCancelOrder) RequestHeaderSinceVersion() uint16 {
	return 0
}

func (e *EngCancelOrder) RequestHeaderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.RequestHeaderSinceVersion()
}

func (*EngCancelOrder) RequestHeaderDeprecated() uint16 {
	return 0
}

func (*EngCancelOrder) RequestHeaderMetaAttribute(meta int) string {
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

func (*EngCancelOrder) OrderIdId() uint16 {
	return 2
}

func (*EngCancelOrder) OrderIdSinceVersion() uint16 {
	return 0
}

func (e *EngCancelOrder) OrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderIdSinceVersion()
}

func (*EngCancelOrder) OrderIdDeprecated() uint16 {
	return 0
}

func (*EngCancelOrder) OrderIdMetaAttribute(meta int) string {
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

func (*EngCancelOrder) OrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCancelOrder) OrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCancelOrder) OrderIdNullValue() int64 {
	return math.MinInt64
}

func (*EngCancelOrder) InstrumentIdId() uint16 {
	return 3
}

func (*EngCancelOrder) InstrumentIdSinceVersion() uint16 {
	return 0
}

func (e *EngCancelOrder) InstrumentIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.InstrumentIdSinceVersion()
}

func (*EngCancelOrder) InstrumentIdDeprecated() uint16 {
	return 0
}

func (*EngCancelOrder) InstrumentIdMetaAttribute(meta int) string {
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

func (*EngCancelOrder) InstrumentIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCancelOrder) InstrumentIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCancelOrder) InstrumentIdNullValue() int64 {
	return math.MinInt64
}

func (*EngCancelOrder) PortfolioIdId() uint16 {
	return 4
}

func (*EngCancelOrder) PortfolioIdSinceVersion() uint16 {
	return 0
}

func (e *EngCancelOrder) PortfolioIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PortfolioIdSinceVersion()
}

func (*EngCancelOrder) PortfolioIdDeprecated() uint16 {
	return 0
}

func (*EngCancelOrder) PortfolioIdMetaAttribute(meta int) string {
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

func (*EngCancelOrder) PortfolioIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCancelOrder) PortfolioIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCancelOrder) PortfolioIdNullValue() int64 {
	return math.MinInt64
}

func (*EngCancelOrder) NewClientOrderIdId() uint16 {
	return 5
}

func (*EngCancelOrder) NewClientOrderIdSinceVersion() uint16 {
	return 0
}

func (e *EngCancelOrder) NewClientOrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NewClientOrderIdSinceVersion()
}

func (*EngCancelOrder) NewClientOrderIdDeprecated() uint16 {
	return 0
}

func (*EngCancelOrder) NewClientOrderIdMetaAttribute(meta int) string {
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

func (*EngCancelOrder) NewClientOrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCancelOrder) NewClientOrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCancelOrder) NewClientOrderIdNullValue() int64 {
	return math.MinInt64
}
