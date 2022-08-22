// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type EngOrderEvent struct {
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

func (e *EngOrderEvent) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	if err := _m.WriteInt64(_w, e.EventId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.InstrumentId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.UserId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.PortfolioId); err != nil {
		return err
	}
	for idx := 0; idx < 2; idx++ {
		if err := _m.WriteInt64(_w, e.ClientOrderId[idx]); err != nil {
			return err
		}
	}
	for idx := 0; idx < 2; idx++ {
		if err := _m.WriteInt64(_w, e.BrokerageClientId[idx]); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, e.Quantity); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.LimitPrice); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.StopPrice); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.TotalFilled); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.FilledVwap); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.ExpireTime); err != nil {
		return err
	}
	if err := e.OrderFlags.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.Side.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.TimeInForce.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.SelfMatchPreventionMode.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.EventType.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.CancelReason.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, e.DetailsBlockLength); err != nil {
		return err
	}
	return nil
}

func (e *EngOrderEvent) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if !e.EventIdInActingVersion(actingVersion) {
		e.EventId = e.EventIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.EventId); err != nil {
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
	if !e.UserIdInActingVersion(actingVersion) {
		e.UserId = e.UserIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.UserId); err != nil {
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
	if !e.ClientOrderIdInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			e.ClientOrderId[idx] = e.ClientOrderIdNullValue()
		}
	} else {
		for idx := 0; idx < 2; idx++ {
			if err := _m.ReadInt64(_r, &e.ClientOrderId[idx]); err != nil {
				return err
			}
		}
	}
	if !e.BrokerageClientIdInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			e.BrokerageClientId[idx] = e.BrokerageClientIdNullValue()
		}
	} else {
		for idx := 0; idx < 2; idx++ {
			if err := _m.ReadInt64(_r, &e.BrokerageClientId[idx]); err != nil {
				return err
			}
		}
	}
	if !e.QuantityInActingVersion(actingVersion) {
		e.Quantity = e.QuantityNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.Quantity); err != nil {
			return err
		}
	}
	if !e.LimitPriceInActingVersion(actingVersion) {
		e.LimitPrice = e.LimitPriceNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.LimitPrice); err != nil {
			return err
		}
	}
	if !e.StopPriceInActingVersion(actingVersion) {
		e.StopPrice = e.StopPriceNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.StopPrice); err != nil {
			return err
		}
	}
	if !e.TotalFilledInActingVersion(actingVersion) {
		e.TotalFilled = e.TotalFilledNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.TotalFilled); err != nil {
			return err
		}
	}
	if !e.FilledVwapInActingVersion(actingVersion) {
		e.FilledVwap = e.FilledVwapNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.FilledVwap); err != nil {
			return err
		}
	}
	if !e.ExpireTimeInActingVersion(actingVersion) {
		e.ExpireTime = e.ExpireTimeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.ExpireTime); err != nil {
			return err
		}
	}
	if e.OrderFlagsInActingVersion(actingVersion) {
		if err := e.OrderFlags.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.SideInActingVersion(actingVersion) {
		if err := e.Side.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.TimeInForceInActingVersion(actingVersion) {
		if err := e.TimeInForce.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.SelfMatchPreventionModeInActingVersion(actingVersion) {
		if err := e.SelfMatchPreventionMode.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.EventTypeInActingVersion(actingVersion) {
		if err := e.EventType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.CancelReasonInActingVersion(actingVersion) {
		if err := e.CancelReason.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.DetailsBlockLengthInActingVersion(actingVersion) {
		e.DetailsBlockLength = e.DetailsBlockLengthNullValue()
	} else {
		if err := _m.ReadUint8(_r, &e.DetailsBlockLength); err != nil {
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

func (e *EngOrderEvent) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.OrderIdInActingVersion(actingVersion) {
		if e.OrderId < e.OrderIdMinValue() || e.OrderId > e.OrderIdMaxValue() {
			return fmt.Errorf("Range check failed on e.OrderId (%v < %v > %v)", e.OrderIdMinValue(), e.OrderId, e.OrderIdMaxValue())
		}
	}
	if e.EventIdInActingVersion(actingVersion) {
		if e.EventId < e.EventIdMinValue() || e.EventId > e.EventIdMaxValue() {
			return fmt.Errorf("Range check failed on e.EventId (%v < %v > %v)", e.EventIdMinValue(), e.EventId, e.EventIdMaxValue())
		}
	}
	if e.InstrumentIdInActingVersion(actingVersion) {
		if e.InstrumentId < e.InstrumentIdMinValue() || e.InstrumentId > e.InstrumentIdMaxValue() {
			return fmt.Errorf("Range check failed on e.InstrumentId (%v < %v > %v)", e.InstrumentIdMinValue(), e.InstrumentId, e.InstrumentIdMaxValue())
		}
	}
	if e.UserIdInActingVersion(actingVersion) {
		if e.UserId < e.UserIdMinValue() || e.UserId > e.UserIdMaxValue() {
			return fmt.Errorf("Range check failed on e.UserId (%v < %v > %v)", e.UserIdMinValue(), e.UserId, e.UserIdMaxValue())
		}
	}
	if e.PortfolioIdInActingVersion(actingVersion) {
		if e.PortfolioId < e.PortfolioIdMinValue() || e.PortfolioId > e.PortfolioIdMaxValue() {
			return fmt.Errorf("Range check failed on e.PortfolioId (%v < %v > %v)", e.PortfolioIdMinValue(), e.PortfolioId, e.PortfolioIdMaxValue())
		}
	}
	if e.ClientOrderIdInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if e.ClientOrderId[idx] < e.ClientOrderIdMinValue() || e.ClientOrderId[idx] > e.ClientOrderIdMaxValue() {
				return fmt.Errorf("Range check failed on e.ClientOrderId[%d] (%v < %v > %v)", idx, e.ClientOrderIdMinValue(), e.ClientOrderId[idx], e.ClientOrderIdMaxValue())
			}
		}
	}
	if e.BrokerageClientIdInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if e.BrokerageClientId[idx] < e.BrokerageClientIdMinValue() || e.BrokerageClientId[idx] > e.BrokerageClientIdMaxValue() {
				return fmt.Errorf("Range check failed on e.BrokerageClientId[%d] (%v < %v > %v)", idx, e.BrokerageClientIdMinValue(), e.BrokerageClientId[idx], e.BrokerageClientIdMaxValue())
			}
		}
	}
	if e.QuantityInActingVersion(actingVersion) {
		if e.Quantity < e.QuantityMinValue() || e.Quantity > e.QuantityMaxValue() {
			return fmt.Errorf("Range check failed on e.Quantity (%v < %v > %v)", e.QuantityMinValue(), e.Quantity, e.QuantityMaxValue())
		}
	}
	if e.LimitPriceInActingVersion(actingVersion) {
		if e.LimitPrice != e.LimitPriceNullValue() && (e.LimitPrice < e.LimitPriceMinValue() || e.LimitPrice > e.LimitPriceMaxValue()) {
			return fmt.Errorf("Range check failed on e.LimitPrice (%v < %v > %v)", e.LimitPriceMinValue(), e.LimitPrice, e.LimitPriceMaxValue())
		}
	}
	if e.StopPriceInActingVersion(actingVersion) {
		if e.StopPrice != e.StopPriceNullValue() && (e.StopPrice < e.StopPriceMinValue() || e.StopPrice > e.StopPriceMaxValue()) {
			return fmt.Errorf("Range check failed on e.StopPrice (%v < %v > %v)", e.StopPriceMinValue(), e.StopPrice, e.StopPriceMaxValue())
		}
	}
	if e.TotalFilledInActingVersion(actingVersion) {
		if e.TotalFilled < e.TotalFilledMinValue() || e.TotalFilled > e.TotalFilledMaxValue() {
			return fmt.Errorf("Range check failed on e.TotalFilled (%v < %v > %v)", e.TotalFilledMinValue(), e.TotalFilled, e.TotalFilledMaxValue())
		}
	}
	if e.FilledVwapInActingVersion(actingVersion) {
		if e.FilledVwap < e.FilledVwapMinValue() || e.FilledVwap > e.FilledVwapMaxValue() {
			return fmt.Errorf("Range check failed on e.FilledVwap (%v < %v > %v)", e.FilledVwapMinValue(), e.FilledVwap, e.FilledVwapMaxValue())
		}
	}
	if e.ExpireTimeInActingVersion(actingVersion) {
		if e.ExpireTime != e.ExpireTimeNullValue() && (e.ExpireTime < e.ExpireTimeMinValue() || e.ExpireTime > e.ExpireTimeMaxValue()) {
			return fmt.Errorf("Range check failed on e.ExpireTime (%v < %v > %v)", e.ExpireTimeMinValue(), e.ExpireTime, e.ExpireTimeMaxValue())
		}
	}
	if err := e.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.TimeInForce.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.SelfMatchPreventionMode.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.EventType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.CancelReason.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if e.DetailsBlockLengthInActingVersion(actingVersion) {
		if e.DetailsBlockLength < e.DetailsBlockLengthMinValue() || e.DetailsBlockLength > e.DetailsBlockLengthMaxValue() {
			return fmt.Errorf("Range check failed on e.DetailsBlockLength (%v < %v > %v)", e.DetailsBlockLengthMinValue(), e.DetailsBlockLength, e.DetailsBlockLengthMaxValue())
		}
	}
	return nil
}

func EngOrderEventInit(e *EngOrderEvent) {
	for idx := 0; idx < 2; idx++ {
		e.BrokerageClientId[idx] = math.MinInt64
	}
	e.LimitPrice = math.MinInt64
	e.StopPrice = math.MinInt64
	e.ExpireTime = math.MinInt64
	return
}

func (*EngOrderEvent) SbeBlockLength() (blockLength uint16) {
	return 166
}

func (*EngOrderEvent) SbeTemplateId() (templateId uint16) {
	return 2610
}

func (*EngOrderEvent) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*EngOrderEvent) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*EngOrderEvent) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*EngOrderEvent) TradingHeaderId() uint16 {
	return 1
}

func (*EngOrderEvent) TradingHeaderSinceVersion() uint16 {
	return 0
}

func (e *EngOrderEvent) TradingHeaderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TradingHeaderSinceVersion()
}

func (*EngOrderEvent) TradingHeaderDeprecated() uint16 {
	return 0
}

func (*EngOrderEvent) TradingHeaderMetaAttribute(meta int) string {
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

func (*EngOrderEvent) OrderIdId() uint16 {
	return 2
}

func (*EngOrderEvent) OrderIdSinceVersion() uint16 {
	return 0
}

func (e *EngOrderEvent) OrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderIdSinceVersion()
}

func (*EngOrderEvent) OrderIdDeprecated() uint16 {
	return 0
}

func (*EngOrderEvent) OrderIdMetaAttribute(meta int) string {
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

func (*EngOrderEvent) OrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngOrderEvent) OrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngOrderEvent) OrderIdNullValue() int64 {
	return math.MinInt64
}

func (*EngOrderEvent) EventIdId() uint16 {
	return 3
}

func (*EngOrderEvent) EventIdSinceVersion() uint16 {
	return 0
}

func (e *EngOrderEvent) EventIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.EventIdSinceVersion()
}

func (*EngOrderEvent) EventIdDeprecated() uint16 {
	return 0
}

func (*EngOrderEvent) EventIdMetaAttribute(meta int) string {
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

func (*EngOrderEvent) EventIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngOrderEvent) EventIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngOrderEvent) EventIdNullValue() int64 {
	return math.MinInt64
}

func (*EngOrderEvent) InstrumentIdId() uint16 {
	return 4
}

func (*EngOrderEvent) InstrumentIdSinceVersion() uint16 {
	return 0
}

func (e *EngOrderEvent) InstrumentIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.InstrumentIdSinceVersion()
}

func (*EngOrderEvent) InstrumentIdDeprecated() uint16 {
	return 0
}

func (*EngOrderEvent) InstrumentIdMetaAttribute(meta int) string {
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

func (*EngOrderEvent) InstrumentIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngOrderEvent) InstrumentIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngOrderEvent) InstrumentIdNullValue() int64 {
	return math.MinInt64
}

func (*EngOrderEvent) UserIdId() uint16 {
	return 5
}

func (*EngOrderEvent) UserIdSinceVersion() uint16 {
	return 0
}

func (e *EngOrderEvent) UserIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.UserIdSinceVersion()
}

func (*EngOrderEvent) UserIdDeprecated() uint16 {
	return 0
}

func (*EngOrderEvent) UserIdMetaAttribute(meta int) string {
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

func (*EngOrderEvent) UserIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngOrderEvent) UserIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngOrderEvent) UserIdNullValue() int64 {
	return math.MinInt64
}

func (*EngOrderEvent) PortfolioIdId() uint16 {
	return 6
}

func (*EngOrderEvent) PortfolioIdSinceVersion() uint16 {
	return 0
}

func (e *EngOrderEvent) PortfolioIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PortfolioIdSinceVersion()
}

func (*EngOrderEvent) PortfolioIdDeprecated() uint16 {
	return 0
}

func (*EngOrderEvent) PortfolioIdMetaAttribute(meta int) string {
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

func (*EngOrderEvent) PortfolioIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngOrderEvent) PortfolioIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngOrderEvent) PortfolioIdNullValue() int64 {
	return math.MinInt64
}

func (*EngOrderEvent) ClientOrderIdId() uint16 {
	return 7
}

func (*EngOrderEvent) ClientOrderIdSinceVersion() uint16 {
	return 0
}

func (e *EngOrderEvent) ClientOrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ClientOrderIdSinceVersion()
}

func (*EngOrderEvent) ClientOrderIdDeprecated() uint16 {
	return 0
}

func (*EngOrderEvent) ClientOrderIdMetaAttribute(meta int) string {
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

func (*EngOrderEvent) ClientOrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngOrderEvent) ClientOrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngOrderEvent) ClientOrderIdNullValue() int64 {
	return math.MinInt64
}

func (*EngOrderEvent) BrokerageClientIdId() uint16 {
	return 8
}

func (*EngOrderEvent) BrokerageClientIdSinceVersion() uint16 {
	return 0
}

func (e *EngOrderEvent) BrokerageClientIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.BrokerageClientIdSinceVersion()
}

func (*EngOrderEvent) BrokerageClientIdDeprecated() uint16 {
	return 0
}

func (*EngOrderEvent) BrokerageClientIdMetaAttribute(meta int) string {
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

func (*EngOrderEvent) BrokerageClientIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngOrderEvent) BrokerageClientIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngOrderEvent) BrokerageClientIdNullValue() int64 {
	return math.MinInt64
}

func (*EngOrderEvent) QuantityId() uint16 {
	return 9
}

func (*EngOrderEvent) QuantitySinceVersion() uint16 {
	return 0
}

func (e *EngOrderEvent) QuantityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.QuantitySinceVersion()
}

func (*EngOrderEvent) QuantityDeprecated() uint16 {
	return 0
}

func (*EngOrderEvent) QuantityMetaAttribute(meta int) string {
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

func (*EngOrderEvent) QuantityMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngOrderEvent) QuantityMaxValue() int64 {
	return math.MaxInt64
}

func (*EngOrderEvent) QuantityNullValue() int64 {
	return math.MinInt64
}

func (*EngOrderEvent) LimitPriceId() uint16 {
	return 10
}

func (*EngOrderEvent) LimitPriceSinceVersion() uint16 {
	return 0
}

func (e *EngOrderEvent) LimitPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LimitPriceSinceVersion()
}

func (*EngOrderEvent) LimitPriceDeprecated() uint16 {
	return 0
}

func (*EngOrderEvent) LimitPriceMetaAttribute(meta int) string {
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

func (*EngOrderEvent) LimitPriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngOrderEvent) LimitPriceMaxValue() int64 {
	return math.MaxInt64
}

func (*EngOrderEvent) LimitPriceNullValue() int64 {
	return math.MinInt64
}

func (*EngOrderEvent) StopPriceId() uint16 {
	return 11
}

func (*EngOrderEvent) StopPriceSinceVersion() uint16 {
	return 0
}

func (e *EngOrderEvent) StopPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.StopPriceSinceVersion()
}

func (*EngOrderEvent) StopPriceDeprecated() uint16 {
	return 0
}

func (*EngOrderEvent) StopPriceMetaAttribute(meta int) string {
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

func (*EngOrderEvent) StopPriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngOrderEvent) StopPriceMaxValue() int64 {
	return math.MaxInt64
}

func (*EngOrderEvent) StopPriceNullValue() int64 {
	return math.MinInt64
}

func (*EngOrderEvent) TotalFilledId() uint16 {
	return 12
}

func (*EngOrderEvent) TotalFilledSinceVersion() uint16 {
	return 0
}

func (e *EngOrderEvent) TotalFilledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TotalFilledSinceVersion()
}

func (*EngOrderEvent) TotalFilledDeprecated() uint16 {
	return 0
}

func (*EngOrderEvent) TotalFilledMetaAttribute(meta int) string {
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

func (*EngOrderEvent) TotalFilledMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngOrderEvent) TotalFilledMaxValue() int64 {
	return math.MaxInt64
}

func (*EngOrderEvent) TotalFilledNullValue() int64 {
	return math.MinInt64
}

func (*EngOrderEvent) FilledVwapId() uint16 {
	return 13
}

func (*EngOrderEvent) FilledVwapSinceVersion() uint16 {
	return 0
}

func (e *EngOrderEvent) FilledVwapInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FilledVwapSinceVersion()
}

func (*EngOrderEvent) FilledVwapDeprecated() uint16 {
	return 0
}

func (*EngOrderEvent) FilledVwapMetaAttribute(meta int) string {
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

func (*EngOrderEvent) FilledVwapMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngOrderEvent) FilledVwapMaxValue() int64 {
	return math.MaxInt64
}

func (*EngOrderEvent) FilledVwapNullValue() int64 {
	return math.MinInt64
}

func (*EngOrderEvent) ExpireTimeId() uint16 {
	return 14
}

func (*EngOrderEvent) ExpireTimeSinceVersion() uint16 {
	return 0
}

func (e *EngOrderEvent) ExpireTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExpireTimeSinceVersion()
}

func (*EngOrderEvent) ExpireTimeDeprecated() uint16 {
	return 0
}

func (*EngOrderEvent) ExpireTimeMetaAttribute(meta int) string {
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

func (*EngOrderEvent) ExpireTimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngOrderEvent) ExpireTimeMaxValue() int64 {
	return math.MaxInt64
}

func (*EngOrderEvent) ExpireTimeNullValue() int64 {
	return math.MinInt64
}

func (*EngOrderEvent) OrderFlagsId() uint16 {
	return 15
}

func (*EngOrderEvent) OrderFlagsSinceVersion() uint16 {
	return 0
}

func (e *EngOrderEvent) OrderFlagsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderFlagsSinceVersion()
}

func (*EngOrderEvent) OrderFlagsDeprecated() uint16 {
	return 0
}

func (*EngOrderEvent) OrderFlagsMetaAttribute(meta int) string {
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

func (*EngOrderEvent) SideId() uint16 {
	return 16
}

func (*EngOrderEvent) SideSinceVersion() uint16 {
	return 0
}

func (e *EngOrderEvent) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SideSinceVersion()
}

func (*EngOrderEvent) SideDeprecated() uint16 {
	return 0
}

func (*EngOrderEvent) SideMetaAttribute(meta int) string {
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

func (*EngOrderEvent) TimeInForceId() uint16 {
	return 17
}

func (*EngOrderEvent) TimeInForceSinceVersion() uint16 {
	return 0
}

func (e *EngOrderEvent) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TimeInForceSinceVersion()
}

func (*EngOrderEvent) TimeInForceDeprecated() uint16 {
	return 0
}

func (*EngOrderEvent) TimeInForceMetaAttribute(meta int) string {
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

func (*EngOrderEvent) SelfMatchPreventionModeId() uint16 {
	return 18
}

func (*EngOrderEvent) SelfMatchPreventionModeSinceVersion() uint16 {
	return 0
}

func (e *EngOrderEvent) SelfMatchPreventionModeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SelfMatchPreventionModeSinceVersion()
}

func (*EngOrderEvent) SelfMatchPreventionModeDeprecated() uint16 {
	return 0
}

func (*EngOrderEvent) SelfMatchPreventionModeMetaAttribute(meta int) string {
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

func (*EngOrderEvent) EventTypeId() uint16 {
	return 19
}

func (*EngOrderEvent) EventTypeSinceVersion() uint16 {
	return 0
}

func (e *EngOrderEvent) EventTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.EventTypeSinceVersion()
}

func (*EngOrderEvent) EventTypeDeprecated() uint16 {
	return 0
}

func (*EngOrderEvent) EventTypeMetaAttribute(meta int) string {
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

func (*EngOrderEvent) CancelReasonId() uint16 {
	return 20
}

func (*EngOrderEvent) CancelReasonSinceVersion() uint16 {
	return 0
}

func (e *EngOrderEvent) CancelReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CancelReasonSinceVersion()
}

func (*EngOrderEvent) CancelReasonDeprecated() uint16 {
	return 0
}

func (*EngOrderEvent) CancelReasonMetaAttribute(meta int) string {
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

func (*EngOrderEvent) DetailsBlockLengthId() uint16 {
	return 21
}

func (*EngOrderEvent) DetailsBlockLengthSinceVersion() uint16 {
	return 0
}

func (e *EngOrderEvent) DetailsBlockLengthInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.DetailsBlockLengthSinceVersion()
}

func (*EngOrderEvent) DetailsBlockLengthDeprecated() uint16 {
	return 0
}

func (*EngOrderEvent) DetailsBlockLengthMetaAttribute(meta int) string {
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

func (*EngOrderEvent) DetailsBlockLengthMinValue() uint8 {
	return 0
}

func (*EngOrderEvent) DetailsBlockLengthMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*EngOrderEvent) DetailsBlockLengthNullValue() uint8 {
	return math.MaxUint8
}
