// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type EngCancelOrderReject struct {
	TradingHeader    TradingEventHeader
	OrderId          int64
	PortfolioId      int64
	UserId           int64
	NewClientOrderId [2]int64
	RejectReason     CancelOrderRejectReasonEnum
}

func (e *EngCancelOrderReject) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := e.RangeCheck(e.SbeSchemaVersion(), e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := e.TradingHeader.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.OrderId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.PortfolioId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.UserId); err != nil {
		return err
	}
	for idx := 0; idx < 2; idx++ {
		if err := _m.WriteInt64(_w, e.NewClientOrderId[idx]); err != nil {
			return err
		}
	}
	if err := e.RejectReason.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (e *EngCancelOrderReject) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if e.TradingHeaderInActingVersion(actingVersion) {
		if err := e.TradingHeader.Decode(_m, _r, actingVersion); err != nil {
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
	if !e.PortfolioIdInActingVersion(actingVersion) {
		e.PortfolioId = e.PortfolioIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.PortfolioId); err != nil {
			return err
		}
	}
	if !e.UserIdInActingVersion(actingVersion) {
		e.UserId = e.UserIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.UserId); err != nil {
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
	if e.RejectReasonInActingVersion(actingVersion) {
		if err := e.RejectReason.Decode(_m, _r, actingVersion); err != nil {
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

func (e *EngCancelOrderReject) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.OrderIdInActingVersion(actingVersion) {
		if e.OrderId < e.OrderIdMinValue() || e.OrderId > e.OrderIdMaxValue() {
			return fmt.Errorf("Range check failed on e.OrderId (%v < %v > %v)", e.OrderIdMinValue(), e.OrderId, e.OrderIdMaxValue())
		}
	}
	if e.PortfolioIdInActingVersion(actingVersion) {
		if e.PortfolioId < e.PortfolioIdMinValue() || e.PortfolioId > e.PortfolioIdMaxValue() {
			return fmt.Errorf("Range check failed on e.PortfolioId (%v < %v > %v)", e.PortfolioIdMinValue(), e.PortfolioId, e.PortfolioIdMaxValue())
		}
	}
	if e.UserIdInActingVersion(actingVersion) {
		if e.UserId < e.UserIdMinValue() || e.UserId > e.UserIdMaxValue() {
			return fmt.Errorf("Range check failed on e.UserId (%v < %v > %v)", e.UserIdMinValue(), e.UserId, e.UserIdMaxValue())
		}
	}
	if e.NewClientOrderIdInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if e.NewClientOrderId[idx] < e.NewClientOrderIdMinValue() || e.NewClientOrderId[idx] > e.NewClientOrderIdMaxValue() {
				return fmt.Errorf("Range check failed on e.NewClientOrderId[%d] (%v < %v > %v)", idx, e.NewClientOrderIdMinValue(), e.NewClientOrderId[idx], e.NewClientOrderIdMaxValue())
			}
		}
	}
	if err := e.RejectReason.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func EngCancelOrderRejectInit(e *EngCancelOrderReject) {
	return
}

func (*EngCancelOrderReject) SbeBlockLength() (blockLength uint16) {
	return 73
}

func (*EngCancelOrderReject) SbeTemplateId() (templateId uint16) {
	return 2630
}

func (*EngCancelOrderReject) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*EngCancelOrderReject) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*EngCancelOrderReject) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*EngCancelOrderReject) TradingHeaderId() uint16 {
	return 1
}

func (*EngCancelOrderReject) TradingHeaderSinceVersion() uint16 {
	return 0
}

func (e *EngCancelOrderReject) TradingHeaderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TradingHeaderSinceVersion()
}

func (*EngCancelOrderReject) TradingHeaderDeprecated() uint16 {
	return 0
}

func (*EngCancelOrderReject) TradingHeaderMetaAttribute(meta int) string {
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

func (*EngCancelOrderReject) OrderIdId() uint16 {
	return 2
}

func (*EngCancelOrderReject) OrderIdSinceVersion() uint16 {
	return 0
}

func (e *EngCancelOrderReject) OrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderIdSinceVersion()
}

func (*EngCancelOrderReject) OrderIdDeprecated() uint16 {
	return 0
}

func (*EngCancelOrderReject) OrderIdMetaAttribute(meta int) string {
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

func (*EngCancelOrderReject) OrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCancelOrderReject) OrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCancelOrderReject) OrderIdNullValue() int64 {
	return math.MinInt64
}

func (*EngCancelOrderReject) PortfolioIdId() uint16 {
	return 3
}

func (*EngCancelOrderReject) PortfolioIdSinceVersion() uint16 {
	return 0
}

func (e *EngCancelOrderReject) PortfolioIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PortfolioIdSinceVersion()
}

func (*EngCancelOrderReject) PortfolioIdDeprecated() uint16 {
	return 0
}

func (*EngCancelOrderReject) PortfolioIdMetaAttribute(meta int) string {
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

func (*EngCancelOrderReject) PortfolioIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCancelOrderReject) PortfolioIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCancelOrderReject) PortfolioIdNullValue() int64 {
	return math.MinInt64
}

func (*EngCancelOrderReject) UserIdId() uint16 {
	return 4
}

func (*EngCancelOrderReject) UserIdSinceVersion() uint16 {
	return 0
}

func (e *EngCancelOrderReject) UserIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.UserIdSinceVersion()
}

func (*EngCancelOrderReject) UserIdDeprecated() uint16 {
	return 0
}

func (*EngCancelOrderReject) UserIdMetaAttribute(meta int) string {
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

func (*EngCancelOrderReject) UserIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCancelOrderReject) UserIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCancelOrderReject) UserIdNullValue() int64 {
	return math.MinInt64
}

func (*EngCancelOrderReject) NewClientOrderIdId() uint16 {
	return 5
}

func (*EngCancelOrderReject) NewClientOrderIdSinceVersion() uint16 {
	return 0
}

func (e *EngCancelOrderReject) NewClientOrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NewClientOrderIdSinceVersion()
}

func (*EngCancelOrderReject) NewClientOrderIdDeprecated() uint16 {
	return 0
}

func (*EngCancelOrderReject) NewClientOrderIdMetaAttribute(meta int) string {
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

func (*EngCancelOrderReject) NewClientOrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCancelOrderReject) NewClientOrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCancelOrderReject) NewClientOrderIdNullValue() int64 {
	return math.MinInt64
}

func (*EngCancelOrderReject) RejectReasonId() uint16 {
	return 6
}

func (*EngCancelOrderReject) RejectReasonSinceVersion() uint16 {
	return 0
}

func (e *EngCancelOrderReject) RejectReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.RejectReasonSinceVersion()
}

func (*EngCancelOrderReject) RejectReasonDeprecated() uint16 {
	return 0
}

func (*EngCancelOrderReject) RejectReasonMetaAttribute(meta int) string {
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
