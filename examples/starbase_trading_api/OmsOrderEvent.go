// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OmsOrderEvent struct {
	TradingHeader           TradingEventHeader
	OrderId                 int64
	EventId                 int64
	InstrumentId            int64
	UserId                  int64
	PortfolioId             int64
	ClientOrderId           [2]int64
	BrokerageClientId       [2]int64
	Quantity                int64
	LimitPrice              int64
	StopPrice               int64
	TotalFilled             int64
	FilledVwap              int64
	ExpireTime              int64
	OrderFlags              OrderFlags
	Side                    SideEnum
	TimeInForce             TimeInForceEnum
	SelfMatchPreventionMode SelfMatchPreventionModeEnum
	EventType               OrderEventTypeEnum
	CancelReason            CancelReasonEnum
	DetailsBlockLength      uint8
}

func (o *OmsOrderEvent) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	if err := _m.WriteInt64(_w, o.EventId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.InstrumentId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.UserId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.PortfolioId); err != nil {
		return err
	}
	for idx := 0; idx < 2; idx++ {
		if err := _m.WriteInt64(_w, o.ClientOrderId[idx]); err != nil {
			return err
		}
	}
	for idx := 0; idx < 2; idx++ {
		if err := _m.WriteInt64(_w, o.BrokerageClientId[idx]); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, o.Quantity); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.LimitPrice); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.StopPrice); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.TotalFilled); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.FilledVwap); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.ExpireTime); err != nil {
		return err
	}
	if err := o.OrderFlags.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.Side.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.TimeInForce.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.SelfMatchPreventionMode.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.EventType.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.CancelReason.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, o.DetailsBlockLength); err != nil {
		return err
	}
	return nil
}

func (o *OmsOrderEvent) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if !o.EventIdInActingVersion(actingVersion) {
		o.EventId = o.EventIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.EventId); err != nil {
			return err
		}
	}
	if !o.InstrumentIdInActingVersion(actingVersion) {
		o.InstrumentId = o.InstrumentIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.InstrumentId); err != nil {
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
	if !o.PortfolioIdInActingVersion(actingVersion) {
		o.PortfolioId = o.PortfolioIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.PortfolioId); err != nil {
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
	if !o.BrokerageClientIdInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			o.BrokerageClientId[idx] = o.BrokerageClientIdNullValue()
		}
	} else {
		for idx := 0; idx < 2; idx++ {
			if err := _m.ReadInt64(_r, &o.BrokerageClientId[idx]); err != nil {
				return err
			}
		}
	}
	if !o.QuantityInActingVersion(actingVersion) {
		o.Quantity = o.QuantityNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.Quantity); err != nil {
			return err
		}
	}
	if !o.LimitPriceInActingVersion(actingVersion) {
		o.LimitPrice = o.LimitPriceNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.LimitPrice); err != nil {
			return err
		}
	}
	if !o.StopPriceInActingVersion(actingVersion) {
		o.StopPrice = o.StopPriceNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.StopPrice); err != nil {
			return err
		}
	}
	if !o.TotalFilledInActingVersion(actingVersion) {
		o.TotalFilled = o.TotalFilledNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.TotalFilled); err != nil {
			return err
		}
	}
	if !o.FilledVwapInActingVersion(actingVersion) {
		o.FilledVwap = o.FilledVwapNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.FilledVwap); err != nil {
			return err
		}
	}
	if !o.ExpireTimeInActingVersion(actingVersion) {
		o.ExpireTime = o.ExpireTimeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.ExpireTime); err != nil {
			return err
		}
	}
	if o.OrderFlagsInActingVersion(actingVersion) {
		if err := o.OrderFlags.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.SideInActingVersion(actingVersion) {
		if err := o.Side.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.TimeInForceInActingVersion(actingVersion) {
		if err := o.TimeInForce.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.SelfMatchPreventionModeInActingVersion(actingVersion) {
		if err := o.SelfMatchPreventionMode.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.EventTypeInActingVersion(actingVersion) {
		if err := o.EventType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.CancelReasonInActingVersion(actingVersion) {
		if err := o.CancelReason.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.DetailsBlockLengthInActingVersion(actingVersion) {
		o.DetailsBlockLength = o.DetailsBlockLengthNullValue()
	} else {
		if err := _m.ReadUint8(_r, &o.DetailsBlockLength); err != nil {
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

func (o *OmsOrderEvent) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.OrderIdInActingVersion(actingVersion) {
		if o.OrderId < o.OrderIdMinValue() || o.OrderId > o.OrderIdMaxValue() {
			return fmt.Errorf("Range check failed on o.OrderId (%v < %v > %v)", o.OrderIdMinValue(), o.OrderId, o.OrderIdMaxValue())
		}
	}
	if o.EventIdInActingVersion(actingVersion) {
		if o.EventId < o.EventIdMinValue() || o.EventId > o.EventIdMaxValue() {
			return fmt.Errorf("Range check failed on o.EventId (%v < %v > %v)", o.EventIdMinValue(), o.EventId, o.EventIdMaxValue())
		}
	}
	if o.InstrumentIdInActingVersion(actingVersion) {
		if o.InstrumentId < o.InstrumentIdMinValue() || o.InstrumentId > o.InstrumentIdMaxValue() {
			return fmt.Errorf("Range check failed on o.InstrumentId (%v < %v > %v)", o.InstrumentIdMinValue(), o.InstrumentId, o.InstrumentIdMaxValue())
		}
	}
	if o.UserIdInActingVersion(actingVersion) {
		if o.UserId < o.UserIdMinValue() || o.UserId > o.UserIdMaxValue() {
			return fmt.Errorf("Range check failed on o.UserId (%v < %v > %v)", o.UserIdMinValue(), o.UserId, o.UserIdMaxValue())
		}
	}
	if o.PortfolioIdInActingVersion(actingVersion) {
		if o.PortfolioId < o.PortfolioIdMinValue() || o.PortfolioId > o.PortfolioIdMaxValue() {
			return fmt.Errorf("Range check failed on o.PortfolioId (%v < %v > %v)", o.PortfolioIdMinValue(), o.PortfolioId, o.PortfolioIdMaxValue())
		}
	}
	if o.ClientOrderIdInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if o.ClientOrderId[idx] < o.ClientOrderIdMinValue() || o.ClientOrderId[idx] > o.ClientOrderIdMaxValue() {
				return fmt.Errorf("Range check failed on o.ClientOrderId[%d] (%v < %v > %v)", idx, o.ClientOrderIdMinValue(), o.ClientOrderId[idx], o.ClientOrderIdMaxValue())
			}
		}
	}
	if o.BrokerageClientIdInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if o.BrokerageClientId[idx] < o.BrokerageClientIdMinValue() || o.BrokerageClientId[idx] > o.BrokerageClientIdMaxValue() {
				return fmt.Errorf("Range check failed on o.BrokerageClientId[%d] (%v < %v > %v)", idx, o.BrokerageClientIdMinValue(), o.BrokerageClientId[idx], o.BrokerageClientIdMaxValue())
			}
		}
	}
	if o.QuantityInActingVersion(actingVersion) {
		if o.Quantity < o.QuantityMinValue() || o.Quantity > o.QuantityMaxValue() {
			return fmt.Errorf("Range check failed on o.Quantity (%v < %v > %v)", o.QuantityMinValue(), o.Quantity, o.QuantityMaxValue())
		}
	}
	if o.LimitPriceInActingVersion(actingVersion) {
		if o.LimitPrice != o.LimitPriceNullValue() && (o.LimitPrice < o.LimitPriceMinValue() || o.LimitPrice > o.LimitPriceMaxValue()) {
			return fmt.Errorf("Range check failed on o.LimitPrice (%v < %v > %v)", o.LimitPriceMinValue(), o.LimitPrice, o.LimitPriceMaxValue())
		}
	}
	if o.StopPriceInActingVersion(actingVersion) {
		if o.StopPrice != o.StopPriceNullValue() && (o.StopPrice < o.StopPriceMinValue() || o.StopPrice > o.StopPriceMaxValue()) {
			return fmt.Errorf("Range check failed on o.StopPrice (%v < %v > %v)", o.StopPriceMinValue(), o.StopPrice, o.StopPriceMaxValue())
		}
	}
	if o.TotalFilledInActingVersion(actingVersion) {
		if o.TotalFilled < o.TotalFilledMinValue() || o.TotalFilled > o.TotalFilledMaxValue() {
			return fmt.Errorf("Range check failed on o.TotalFilled (%v < %v > %v)", o.TotalFilledMinValue(), o.TotalFilled, o.TotalFilledMaxValue())
		}
	}
	if o.FilledVwapInActingVersion(actingVersion) {
		if o.FilledVwap < o.FilledVwapMinValue() || o.FilledVwap > o.FilledVwapMaxValue() {
			return fmt.Errorf("Range check failed on o.FilledVwap (%v < %v > %v)", o.FilledVwapMinValue(), o.FilledVwap, o.FilledVwapMaxValue())
		}
	}
	if o.ExpireTimeInActingVersion(actingVersion) {
		if o.ExpireTime != o.ExpireTimeNullValue() && (o.ExpireTime < o.ExpireTimeMinValue() || o.ExpireTime > o.ExpireTimeMaxValue()) {
			return fmt.Errorf("Range check failed on o.ExpireTime (%v < %v > %v)", o.ExpireTimeMinValue(), o.ExpireTime, o.ExpireTimeMaxValue())
		}
	}
	if err := o.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.TimeInForce.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.SelfMatchPreventionMode.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.EventType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.CancelReason.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.DetailsBlockLengthInActingVersion(actingVersion) {
		if o.DetailsBlockLength < o.DetailsBlockLengthMinValue() || o.DetailsBlockLength > o.DetailsBlockLengthMaxValue() {
			return fmt.Errorf("Range check failed on o.DetailsBlockLength (%v < %v > %v)", o.DetailsBlockLengthMinValue(), o.DetailsBlockLength, o.DetailsBlockLengthMaxValue())
		}
	}
	return nil
}

func OmsOrderEventInit(o *OmsOrderEvent) {
	for idx := 0; idx < 2; idx++ {
		o.BrokerageClientId[idx] = math.MinInt64
	}
	o.LimitPrice = math.MinInt64
	o.StopPrice = math.MinInt64
	o.ExpireTime = math.MinInt64
	return
}

func (*OmsOrderEvent) SbeBlockLength() (blockLength uint16) {
	return 166
}

func (*OmsOrderEvent) SbeTemplateId() (templateId uint16) {
	return 1610
}

func (*OmsOrderEvent) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*OmsOrderEvent) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*OmsOrderEvent) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*OmsOrderEvent) TradingHeaderId() uint16 {
	return 1
}

func (*OmsOrderEvent) TradingHeaderSinceVersion() uint16 {
	return 0
}

func (o *OmsOrderEvent) TradingHeaderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TradingHeaderSinceVersion()
}

func (*OmsOrderEvent) TradingHeaderDeprecated() uint16 {
	return 0
}

func (*OmsOrderEvent) TradingHeaderMetaAttribute(meta int) string {
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

func (*OmsOrderEvent) OrderIdId() uint16 {
	return 2
}

func (*OmsOrderEvent) OrderIdSinceVersion() uint16 {
	return 0
}

func (o *OmsOrderEvent) OrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderIdSinceVersion()
}

func (*OmsOrderEvent) OrderIdDeprecated() uint16 {
	return 0
}

func (*OmsOrderEvent) OrderIdMetaAttribute(meta int) string {
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

func (*OmsOrderEvent) OrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsOrderEvent) OrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsOrderEvent) OrderIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsOrderEvent) EventIdId() uint16 {
	return 3
}

func (*OmsOrderEvent) EventIdSinceVersion() uint16 {
	return 0
}

func (o *OmsOrderEvent) EventIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.EventIdSinceVersion()
}

func (*OmsOrderEvent) EventIdDeprecated() uint16 {
	return 0
}

func (*OmsOrderEvent) EventIdMetaAttribute(meta int) string {
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

func (*OmsOrderEvent) EventIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsOrderEvent) EventIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsOrderEvent) EventIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsOrderEvent) InstrumentIdId() uint16 {
	return 4
}

func (*OmsOrderEvent) InstrumentIdSinceVersion() uint16 {
	return 0
}

func (o *OmsOrderEvent) InstrumentIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.InstrumentIdSinceVersion()
}

func (*OmsOrderEvent) InstrumentIdDeprecated() uint16 {
	return 0
}

func (*OmsOrderEvent) InstrumentIdMetaAttribute(meta int) string {
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

func (*OmsOrderEvent) InstrumentIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsOrderEvent) InstrumentIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsOrderEvent) InstrumentIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsOrderEvent) UserIdId() uint16 {
	return 5
}

func (*OmsOrderEvent) UserIdSinceVersion() uint16 {
	return 0
}

func (o *OmsOrderEvent) UserIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.UserIdSinceVersion()
}

func (*OmsOrderEvent) UserIdDeprecated() uint16 {
	return 0
}

func (*OmsOrderEvent) UserIdMetaAttribute(meta int) string {
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

func (*OmsOrderEvent) UserIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsOrderEvent) UserIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsOrderEvent) UserIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsOrderEvent) PortfolioIdId() uint16 {
	return 6
}

func (*OmsOrderEvent) PortfolioIdSinceVersion() uint16 {
	return 0
}

func (o *OmsOrderEvent) PortfolioIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PortfolioIdSinceVersion()
}

func (*OmsOrderEvent) PortfolioIdDeprecated() uint16 {
	return 0
}

func (*OmsOrderEvent) PortfolioIdMetaAttribute(meta int) string {
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

func (*OmsOrderEvent) PortfolioIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsOrderEvent) PortfolioIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsOrderEvent) PortfolioIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsOrderEvent) ClientOrderIdId() uint16 {
	return 7
}

func (*OmsOrderEvent) ClientOrderIdSinceVersion() uint16 {
	return 0
}

func (o *OmsOrderEvent) ClientOrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ClientOrderIdSinceVersion()
}

func (*OmsOrderEvent) ClientOrderIdDeprecated() uint16 {
	return 0
}

func (*OmsOrderEvent) ClientOrderIdMetaAttribute(meta int) string {
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

func (*OmsOrderEvent) ClientOrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsOrderEvent) ClientOrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsOrderEvent) ClientOrderIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsOrderEvent) BrokerageClientIdId() uint16 {
	return 8
}

func (*OmsOrderEvent) BrokerageClientIdSinceVersion() uint16 {
	return 0
}

func (o *OmsOrderEvent) BrokerageClientIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.BrokerageClientIdSinceVersion()
}

func (*OmsOrderEvent) BrokerageClientIdDeprecated() uint16 {
	return 0
}

func (*OmsOrderEvent) BrokerageClientIdMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*OmsOrderEvent) BrokerageClientIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsOrderEvent) BrokerageClientIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsOrderEvent) BrokerageClientIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsOrderEvent) QuantityId() uint16 {
	return 9
}

func (*OmsOrderEvent) QuantitySinceVersion() uint16 {
	return 0
}

func (o *OmsOrderEvent) QuantityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.QuantitySinceVersion()
}

func (*OmsOrderEvent) QuantityDeprecated() uint16 {
	return 0
}

func (*OmsOrderEvent) QuantityMetaAttribute(meta int) string {
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

func (*OmsOrderEvent) QuantityMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsOrderEvent) QuantityMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsOrderEvent) QuantityNullValue() int64 {
	return math.MinInt64
}

func (*OmsOrderEvent) LimitPriceId() uint16 {
	return 10
}

func (*OmsOrderEvent) LimitPriceSinceVersion() uint16 {
	return 0
}

func (o *OmsOrderEvent) LimitPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LimitPriceSinceVersion()
}

func (*OmsOrderEvent) LimitPriceDeprecated() uint16 {
	return 0
}

func (*OmsOrderEvent) LimitPriceMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*OmsOrderEvent) LimitPriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsOrderEvent) LimitPriceMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsOrderEvent) LimitPriceNullValue() int64 {
	return math.MinInt64
}

func (*OmsOrderEvent) StopPriceId() uint16 {
	return 11
}

func (*OmsOrderEvent) StopPriceSinceVersion() uint16 {
	return 0
}

func (o *OmsOrderEvent) StopPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StopPriceSinceVersion()
}

func (*OmsOrderEvent) StopPriceDeprecated() uint16 {
	return 0
}

func (*OmsOrderEvent) StopPriceMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*OmsOrderEvent) StopPriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsOrderEvent) StopPriceMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsOrderEvent) StopPriceNullValue() int64 {
	return math.MinInt64
}

func (*OmsOrderEvent) TotalFilledId() uint16 {
	return 12
}

func (*OmsOrderEvent) TotalFilledSinceVersion() uint16 {
	return 0
}

func (o *OmsOrderEvent) TotalFilledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TotalFilledSinceVersion()
}

func (*OmsOrderEvent) TotalFilledDeprecated() uint16 {
	return 0
}

func (*OmsOrderEvent) TotalFilledMetaAttribute(meta int) string {
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

func (*OmsOrderEvent) TotalFilledMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsOrderEvent) TotalFilledMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsOrderEvent) TotalFilledNullValue() int64 {
	return math.MinInt64
}

func (*OmsOrderEvent) FilledVwapId() uint16 {
	return 13
}

func (*OmsOrderEvent) FilledVwapSinceVersion() uint16 {
	return 0
}

func (o *OmsOrderEvent) FilledVwapInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.FilledVwapSinceVersion()
}

func (*OmsOrderEvent) FilledVwapDeprecated() uint16 {
	return 0
}

func (*OmsOrderEvent) FilledVwapMetaAttribute(meta int) string {
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

func (*OmsOrderEvent) FilledVwapMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsOrderEvent) FilledVwapMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsOrderEvent) FilledVwapNullValue() int64 {
	return math.MinInt64
}

func (*OmsOrderEvent) ExpireTimeId() uint16 {
	return 14
}

func (*OmsOrderEvent) ExpireTimeSinceVersion() uint16 {
	return 0
}

func (o *OmsOrderEvent) ExpireTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ExpireTimeSinceVersion()
}

func (*OmsOrderEvent) ExpireTimeDeprecated() uint16 {
	return 0
}

func (*OmsOrderEvent) ExpireTimeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*OmsOrderEvent) ExpireTimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsOrderEvent) ExpireTimeMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsOrderEvent) ExpireTimeNullValue() int64 {
	return math.MinInt64
}

func (*OmsOrderEvent) OrderFlagsId() uint16 {
	return 15
}

func (*OmsOrderEvent) OrderFlagsSinceVersion() uint16 {
	return 0
}

func (o *OmsOrderEvent) OrderFlagsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderFlagsSinceVersion()
}

func (*OmsOrderEvent) OrderFlagsDeprecated() uint16 {
	return 0
}

func (*OmsOrderEvent) OrderFlagsMetaAttribute(meta int) string {
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

func (*OmsOrderEvent) SideId() uint16 {
	return 16
}

func (*OmsOrderEvent) SideSinceVersion() uint16 {
	return 0
}

func (o *OmsOrderEvent) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SideSinceVersion()
}

func (*OmsOrderEvent) SideDeprecated() uint16 {
	return 0
}

func (*OmsOrderEvent) SideMetaAttribute(meta int) string {
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

func (*OmsOrderEvent) TimeInForceId() uint16 {
	return 17
}

func (*OmsOrderEvent) TimeInForceSinceVersion() uint16 {
	return 0
}

func (o *OmsOrderEvent) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TimeInForceSinceVersion()
}

func (*OmsOrderEvent) TimeInForceDeprecated() uint16 {
	return 0
}

func (*OmsOrderEvent) TimeInForceMetaAttribute(meta int) string {
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

func (*OmsOrderEvent) SelfMatchPreventionModeId() uint16 {
	return 18
}

func (*OmsOrderEvent) SelfMatchPreventionModeSinceVersion() uint16 {
	return 0
}

func (o *OmsOrderEvent) SelfMatchPreventionModeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SelfMatchPreventionModeSinceVersion()
}

func (*OmsOrderEvent) SelfMatchPreventionModeDeprecated() uint16 {
	return 0
}

func (*OmsOrderEvent) SelfMatchPreventionModeMetaAttribute(meta int) string {
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

func (*OmsOrderEvent) EventTypeId() uint16 {
	return 19
}

func (*OmsOrderEvent) EventTypeSinceVersion() uint16 {
	return 0
}

func (o *OmsOrderEvent) EventTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.EventTypeSinceVersion()
}

func (*OmsOrderEvent) EventTypeDeprecated() uint16 {
	return 0
}

func (*OmsOrderEvent) EventTypeMetaAttribute(meta int) string {
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

func (*OmsOrderEvent) CancelReasonId() uint16 {
	return 20
}

func (*OmsOrderEvent) CancelReasonSinceVersion() uint16 {
	return 0
}

func (o *OmsOrderEvent) CancelReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CancelReasonSinceVersion()
}

func (*OmsOrderEvent) CancelReasonDeprecated() uint16 {
	return 0
}

func (*OmsOrderEvent) CancelReasonMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*OmsOrderEvent) DetailsBlockLengthId() uint16 {
	return 21
}

func (*OmsOrderEvent) DetailsBlockLengthSinceVersion() uint16 {
	return 0
}

func (o *OmsOrderEvent) DetailsBlockLengthInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.DetailsBlockLengthSinceVersion()
}

func (*OmsOrderEvent) DetailsBlockLengthDeprecated() uint16 {
	return 0
}

func (*OmsOrderEvent) DetailsBlockLengthMetaAttribute(meta int) string {
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

func (*OmsOrderEvent) DetailsBlockLengthMinValue() uint8 {
	return 0
}

func (*OmsOrderEvent) DetailsBlockLengthMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*OmsOrderEvent) DetailsBlockLengthNullValue() uint8 {
	return math.MaxUint8
}
