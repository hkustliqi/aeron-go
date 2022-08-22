// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OmsCancelOrderReject struct {
	TradingHeader    TradingEventHeader
	OrderId          int64
	PortfolioId      int64
	UserId           int64
	ClientOrderId    [2]int64
	NewClientOrderId [2]int64
	RejectReason     CancelOrderRejectReasonEnum
}

func (o *OmsCancelOrderReject) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := o.TradingHeader.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.OrderId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.PortfolioId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.UserId); err != nil {
		return err
	}
	for idx := 0; idx < 2; idx++ {
		if err := _m.WriteInt64(_w, o.ClientOrderId[idx]); err != nil {
			return err
		}
	}
	for idx := 0; idx < 2; idx++ {
		if err := _m.WriteInt64(_w, o.NewClientOrderId[idx]); err != nil {
			return err
		}
	}
	if err := o.RejectReason.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (o *OmsCancelOrderReject) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if o.TradingHeaderInActingVersion(actingVersion) {
		if err := o.TradingHeader.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.OrderIdInActingVersion(actingVersion) {
		o.OrderId = o.OrderIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.OrderId); err != nil {
			return err
		}
	}
	if !o.PortfolioIdInActingVersion(actingVersion) {
		o.PortfolioId = o.PortfolioIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.PortfolioId); err != nil {
			return err
		}
	}
	if !o.UserIdInActingVersion(actingVersion) {
		o.UserId = o.UserIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.UserId); err != nil {
			return err
		}
	}
	if !o.ClientOrderIdInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			o.ClientOrderId[idx] = o.ClientOrderIdNullValue()
		}
	} else {
		for idx := 0; idx < 2; idx++ {
			if err := _m.ReadInt64(_r, &o.ClientOrderId[idx]); err != nil {
				return err
			}
		}
	}
	if !o.NewClientOrderIdInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			o.NewClientOrderId[idx] = o.NewClientOrderIdNullValue()
		}
	} else {
		for idx := 0; idx < 2; idx++ {
			if err := _m.ReadInt64(_r, &o.NewClientOrderId[idx]); err != nil {
				return err
			}
		}
	}
	if o.RejectReasonInActingVersion(actingVersion) {
		if err := o.RejectReason.Decode(_m, _r, actingVersion); err != nil {
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

func (o *OmsCancelOrderReject) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.OrderIdInActingVersion(actingVersion) {
		if o.OrderId < o.OrderIdMinValue() || o.OrderId > o.OrderIdMaxValue() {
			return fmt.Errorf("Range check failed on o.OrderId (%v < %v > %v)", o.OrderIdMinValue(), o.OrderId, o.OrderIdMaxValue())
		}
	}
	if o.PortfolioIdInActingVersion(actingVersion) {
		if o.PortfolioId < o.PortfolioIdMinValue() || o.PortfolioId > o.PortfolioIdMaxValue() {
			return fmt.Errorf("Range check failed on o.PortfolioId (%v < %v > %v)", o.PortfolioIdMinValue(), o.PortfolioId, o.PortfolioIdMaxValue())
		}
	}
	if o.UserIdInActingVersion(actingVersion) {
		if o.UserId < o.UserIdMinValue() || o.UserId > o.UserIdMaxValue() {
			return fmt.Errorf("Range check failed on o.UserId (%v < %v > %v)", o.UserIdMinValue(), o.UserId, o.UserIdMaxValue())
		}
	}
	if o.ClientOrderIdInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if o.ClientOrderId[idx] < o.ClientOrderIdMinValue() || o.ClientOrderId[idx] > o.ClientOrderIdMaxValue() {
				return fmt.Errorf("Range check failed on o.ClientOrderId[%d] (%v < %v > %v)", idx, o.ClientOrderIdMinValue(), o.ClientOrderId[idx], o.ClientOrderIdMaxValue())
			}
		}
	}
	if o.NewClientOrderIdInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if o.NewClientOrderId[idx] < o.NewClientOrderIdMinValue() || o.NewClientOrderId[idx] > o.NewClientOrderIdMaxValue() {
				return fmt.Errorf("Range check failed on o.NewClientOrderId[%d] (%v < %v > %v)", idx, o.NewClientOrderIdMinValue(), o.NewClientOrderId[idx], o.NewClientOrderIdMaxValue())
			}
		}
	}
	if err := o.RejectReason.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func OmsCancelOrderRejectInit(o *OmsCancelOrderReject) {
	return
}

func (*OmsCancelOrderReject) SbeBlockLength() (blockLength uint16) {
	return 89
}

func (*OmsCancelOrderReject) SbeTemplateId() (templateId uint16) {
	return 1630
}

func (*OmsCancelOrderReject) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*OmsCancelOrderReject) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*OmsCancelOrderReject) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*OmsCancelOrderReject) TradingHeaderId() uint16 {
	return 1
}

func (*OmsCancelOrderReject) TradingHeaderSinceVersion() uint16 {
	return 0
}

func (o *OmsCancelOrderReject) TradingHeaderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TradingHeaderSinceVersion()
}

func (*OmsCancelOrderReject) TradingHeaderDeprecated() uint16 {
	return 0
}

func (*OmsCancelOrderReject) TradingHeaderMetaAttribute(meta int) string {
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

func (*OmsCancelOrderReject) OrderIdId() uint16 {
	return 2
}

func (*OmsCancelOrderReject) OrderIdSinceVersion() uint16 {
	return 0
}

func (o *OmsCancelOrderReject) OrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderIdSinceVersion()
}

func (*OmsCancelOrderReject) OrderIdDeprecated() uint16 {
	return 0
}

func (*OmsCancelOrderReject) OrderIdMetaAttribute(meta int) string {
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

func (*OmsCancelOrderReject) OrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsCancelOrderReject) OrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsCancelOrderReject) OrderIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsCancelOrderReject) PortfolioIdId() uint16 {
	return 3
}

func (*OmsCancelOrderReject) PortfolioIdSinceVersion() uint16 {
	return 0
}

func (o *OmsCancelOrderReject) PortfolioIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PortfolioIdSinceVersion()
}

func (*OmsCancelOrderReject) PortfolioIdDeprecated() uint16 {
	return 0
}

func (*OmsCancelOrderReject) PortfolioIdMetaAttribute(meta int) string {
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

func (*OmsCancelOrderReject) PortfolioIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsCancelOrderReject) PortfolioIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsCancelOrderReject) PortfolioIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsCancelOrderReject) UserIdId() uint16 {
	return 4
}

func (*OmsCancelOrderReject) UserIdSinceVersion() uint16 {
	return 0
}

func (o *OmsCancelOrderReject) UserIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.UserIdSinceVersion()
}

func (*OmsCancelOrderReject) UserIdDeprecated() uint16 {
	return 0
}

func (*OmsCancelOrderReject) UserIdMetaAttribute(meta int) string {
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

func (*OmsCancelOrderReject) UserIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsCancelOrderReject) UserIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsCancelOrderReject) UserIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsCancelOrderReject) ClientOrderIdId() uint16 {
	return 5
}

func (*OmsCancelOrderReject) ClientOrderIdSinceVersion() uint16 {
	return 0
}

func (o *OmsCancelOrderReject) ClientOrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ClientOrderIdSinceVersion()
}

func (*OmsCancelOrderReject) ClientOrderIdDeprecated() uint16 {
	return 0
}

func (*OmsCancelOrderReject) ClientOrderIdMetaAttribute(meta int) string {
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

func (*OmsCancelOrderReject) ClientOrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsCancelOrderReject) ClientOrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsCancelOrderReject) ClientOrderIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsCancelOrderReject) NewClientOrderIdId() uint16 {
	return 6
}

func (*OmsCancelOrderReject) NewClientOrderIdSinceVersion() uint16 {
	return 0
}

func (o *OmsCancelOrderReject) NewClientOrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.NewClientOrderIdSinceVersion()
}

func (*OmsCancelOrderReject) NewClientOrderIdDeprecated() uint16 {
	return 0
}

func (*OmsCancelOrderReject) NewClientOrderIdMetaAttribute(meta int) string {
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

func (*OmsCancelOrderReject) NewClientOrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsCancelOrderReject) NewClientOrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsCancelOrderReject) NewClientOrderIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsCancelOrderReject) RejectReasonId() uint16 {
	return 7
}

func (*OmsCancelOrderReject) RejectReasonSinceVersion() uint16 {
	return 0
}

func (o *OmsCancelOrderReject) RejectReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.RejectReasonSinceVersion()
}

func (*OmsCancelOrderReject) RejectReasonDeprecated() uint16 {
	return 0
}

func (*OmsCancelOrderReject) RejectReasonMetaAttribute(meta int) string {
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
