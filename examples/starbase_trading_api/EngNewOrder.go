// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type EngNewOrder struct {
	RequestHeader           EngUserRequestHeader
	OrderId                 int64
	InstrumentId            int64
	PortfolioId             int64
	ClientOrderId           [2]int64
	BrokerageClientId       [2]int64
	LimitPrice              int64
	StopPrice               int64
	Quantity                int64
	Funds                   int64
	OrderFlags              OrderFlags
	ExpireTime              int32
	SelfMatchPreventionMode SelfMatchPreventionModeEnum
	Side                    SideEnum
	TimeInForce             TimeInForceEnum
}

func (e *EngNewOrder) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
		if err := _m.WriteInt64(_w, e.ClientOrderId[idx]); err != nil {
			return err
		}
	}
	for idx := 0; idx < 2; idx++ {
		if err := _m.WriteInt64(_w, e.BrokerageClientId[idx]); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, e.LimitPrice); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.StopPrice); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.Quantity); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.Funds); err != nil {
		return err
	}
	if err := e.OrderFlags.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, e.ExpireTime); err != nil {
		return err
	}
	if err := e.SelfMatchPreventionMode.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.Side.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.TimeInForce.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (e *EngNewOrder) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if !e.QuantityInActingVersion(actingVersion) {
		e.Quantity = e.QuantityNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.Quantity); err != nil {
			return err
		}
	}
	if !e.FundsInActingVersion(actingVersion) {
		e.Funds = e.FundsNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.Funds); err != nil {
			return err
		}
	}
	if e.OrderFlagsInActingVersion(actingVersion) {
		if err := e.OrderFlags.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.ExpireTimeInActingVersion(actingVersion) {
		e.ExpireTime = e.ExpireTimeNullValue()
	} else {
		if err := _m.ReadInt32(_r, &e.ExpireTime); err != nil {
			return err
		}
	}
	if e.SelfMatchPreventionModeInActingVersion(actingVersion) {
		if err := e.SelfMatchPreventionMode.Decode(_m, _r, actingVersion); err != nil {
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

func (e *EngNewOrder) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if e.LimitPriceInActingVersion(actingVersion) {
		if e.LimitPrice < e.LimitPriceMinValue() || e.LimitPrice > e.LimitPriceMaxValue() {
			return fmt.Errorf("Range check failed on e.LimitPrice (%v < %v > %v)", e.LimitPriceMinValue(), e.LimitPrice, e.LimitPriceMaxValue())
		}
	}
	if e.StopPriceInActingVersion(actingVersion) {
		if e.StopPrice != e.StopPriceNullValue() && (e.StopPrice < e.StopPriceMinValue() || e.StopPrice > e.StopPriceMaxValue()) {
			return fmt.Errorf("Range check failed on e.StopPrice (%v < %v > %v)", e.StopPriceMinValue(), e.StopPrice, e.StopPriceMaxValue())
		}
	}
	if e.QuantityInActingVersion(actingVersion) {
		if e.Quantity < e.QuantityMinValue() || e.Quantity > e.QuantityMaxValue() {
			return fmt.Errorf("Range check failed on e.Quantity (%v < %v > %v)", e.QuantityMinValue(), e.Quantity, e.QuantityMaxValue())
		}
	}
	if e.FundsInActingVersion(actingVersion) {
		if e.Funds != e.FundsNullValue() && (e.Funds < e.FundsMinValue() || e.Funds > e.FundsMaxValue()) {
			return fmt.Errorf("Range check failed on e.Funds (%v < %v > %v)", e.FundsMinValue(), e.Funds, e.FundsMaxValue())
		}
	}
	if e.ExpireTimeInActingVersion(actingVersion) {
		if e.ExpireTime != e.ExpireTimeNullValue() && (e.ExpireTime < e.ExpireTimeMinValue() || e.ExpireTime > e.ExpireTimeMaxValue()) {
			return fmt.Errorf("Range check failed on e.ExpireTime (%v < %v > %v)", e.ExpireTimeMinValue(), e.ExpireTime, e.ExpireTimeMaxValue())
		}
	}
	if err := e.SelfMatchPreventionMode.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.TimeInForce.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func EngNewOrderInit(e *EngNewOrder) {
	e.StopPrice = math.MinInt64
	e.Funds = math.MinInt64
	e.ExpireTime = math.MinInt32
	return
}

func (*EngNewOrder) SbeBlockLength() (blockLength uint16) {
	return 143
}

func (*EngNewOrder) SbeTemplateId() (templateId uint16) {
	return 2000
}

func (*EngNewOrder) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*EngNewOrder) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*EngNewOrder) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*EngNewOrder) RequestHeaderId() uint16 {
	return 1
}

func (*EngNewOrder) RequestHeaderSinceVersion() uint16 {
	return 0
}

func (e *EngNewOrder) RequestHeaderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.RequestHeaderSinceVersion()
}

func (*EngNewOrder) RequestHeaderDeprecated() uint16 {
	return 0
}

func (*EngNewOrder) RequestHeaderMetaAttribute(meta int) string {
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

func (*EngNewOrder) OrderIdId() uint16 {
	return 2
}

func (*EngNewOrder) OrderIdSinceVersion() uint16 {
	return 0
}

func (e *EngNewOrder) OrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderIdSinceVersion()
}

func (*EngNewOrder) OrderIdDeprecated() uint16 {
	return 0
}

func (*EngNewOrder) OrderIdMetaAttribute(meta int) string {
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

func (*EngNewOrder) OrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngNewOrder) OrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngNewOrder) OrderIdNullValue() int64 {
	return math.MinInt64
}

func (*EngNewOrder) InstrumentIdId() uint16 {
	return 3
}

func (*EngNewOrder) InstrumentIdSinceVersion() uint16 {
	return 0
}

func (e *EngNewOrder) InstrumentIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.InstrumentIdSinceVersion()
}

func (*EngNewOrder) InstrumentIdDeprecated() uint16 {
	return 0
}

func (*EngNewOrder) InstrumentIdMetaAttribute(meta int) string {
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

func (*EngNewOrder) InstrumentIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngNewOrder) InstrumentIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngNewOrder) InstrumentIdNullValue() int64 {
	return math.MinInt64
}

func (*EngNewOrder) PortfolioIdId() uint16 {
	return 4
}

func (*EngNewOrder) PortfolioIdSinceVersion() uint16 {
	return 0
}

func (e *EngNewOrder) PortfolioIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PortfolioIdSinceVersion()
}

func (*EngNewOrder) PortfolioIdDeprecated() uint16 {
	return 0
}

func (*EngNewOrder) PortfolioIdMetaAttribute(meta int) string {
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

func (*EngNewOrder) PortfolioIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngNewOrder) PortfolioIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngNewOrder) PortfolioIdNullValue() int64 {
	return math.MinInt64
}

func (*EngNewOrder) ClientOrderIdId() uint16 {
	return 5
}

func (*EngNewOrder) ClientOrderIdSinceVersion() uint16 {
	return 0
}

func (e *EngNewOrder) ClientOrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ClientOrderIdSinceVersion()
}

func (*EngNewOrder) ClientOrderIdDeprecated() uint16 {
	return 0
}

func (*EngNewOrder) ClientOrderIdMetaAttribute(meta int) string {
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

func (*EngNewOrder) ClientOrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngNewOrder) ClientOrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngNewOrder) ClientOrderIdNullValue() int64 {
	return math.MinInt64
}

func (*EngNewOrder) BrokerageClientIdId() uint16 {
	return 6
}

func (*EngNewOrder) BrokerageClientIdSinceVersion() uint16 {
	return 0
}

func (e *EngNewOrder) BrokerageClientIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.BrokerageClientIdSinceVersion()
}

func (*EngNewOrder) BrokerageClientIdDeprecated() uint16 {
	return 0
}

func (*EngNewOrder) BrokerageClientIdMetaAttribute(meta int) string {
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

func (*EngNewOrder) BrokerageClientIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngNewOrder) BrokerageClientIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngNewOrder) BrokerageClientIdNullValue() int64 {
	return math.MinInt64
}

func (*EngNewOrder) LimitPriceId() uint16 {
	return 7
}

func (*EngNewOrder) LimitPriceSinceVersion() uint16 {
	return 0
}

func (e *EngNewOrder) LimitPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LimitPriceSinceVersion()
}

func (*EngNewOrder) LimitPriceDeprecated() uint16 {
	return 0
}

func (*EngNewOrder) LimitPriceMetaAttribute(meta int) string {
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

func (*EngNewOrder) LimitPriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngNewOrder) LimitPriceMaxValue() int64 {
	return math.MaxInt64
}

func (*EngNewOrder) LimitPriceNullValue() int64 {
	return math.MinInt64
}

func (*EngNewOrder) StopPriceId() uint16 {
	return 8
}

func (*EngNewOrder) StopPriceSinceVersion() uint16 {
	return 0
}

func (e *EngNewOrder) StopPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.StopPriceSinceVersion()
}

func (*EngNewOrder) StopPriceDeprecated() uint16 {
	return 0
}

func (*EngNewOrder) StopPriceMetaAttribute(meta int) string {
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

func (*EngNewOrder) StopPriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngNewOrder) StopPriceMaxValue() int64 {
	return math.MaxInt64
}

func (*EngNewOrder) StopPriceNullValue() int64 {
	return math.MinInt64
}

func (*EngNewOrder) QuantityId() uint16 {
	return 9
}

func (*EngNewOrder) QuantitySinceVersion() uint16 {
	return 0
}

func (e *EngNewOrder) QuantityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.QuantitySinceVersion()
}

func (*EngNewOrder) QuantityDeprecated() uint16 {
	return 0
}

func (*EngNewOrder) QuantityMetaAttribute(meta int) string {
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

func (*EngNewOrder) QuantityMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngNewOrder) QuantityMaxValue() int64 {
	return math.MaxInt64
}

func (*EngNewOrder) QuantityNullValue() int64 {
	return math.MinInt64
}

func (*EngNewOrder) FundsId() uint16 {
	return 11
}

func (*EngNewOrder) FundsSinceVersion() uint16 {
	return 0
}

func (e *EngNewOrder) FundsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FundsSinceVersion()
}

func (*EngNewOrder) FundsDeprecated() uint16 {
	return 0
}

func (*EngNewOrder) FundsMetaAttribute(meta int) string {
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

func (*EngNewOrder) FundsMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngNewOrder) FundsMaxValue() int64 {
	return math.MaxInt64
}

func (*EngNewOrder) FundsNullValue() int64 {
	return math.MinInt64
}

func (*EngNewOrder) OrderFlagsId() uint16 {
	return 12
}

func (*EngNewOrder) OrderFlagsSinceVersion() uint16 {
	return 0
}

func (e *EngNewOrder) OrderFlagsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderFlagsSinceVersion()
}

func (*EngNewOrder) OrderFlagsDeprecated() uint16 {
	return 0
}

func (*EngNewOrder) OrderFlagsMetaAttribute(meta int) string {
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

func (*EngNewOrder) ExpireTimeId() uint16 {
	return 13
}

func (*EngNewOrder) ExpireTimeSinceVersion() uint16 {
	return 0
}

func (e *EngNewOrder) ExpireTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExpireTimeSinceVersion()
}

func (*EngNewOrder) ExpireTimeDeprecated() uint16 {
	return 0
}

func (*EngNewOrder) ExpireTimeMetaAttribute(meta int) string {
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

func (*EngNewOrder) ExpireTimeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*EngNewOrder) ExpireTimeMaxValue() int32 {
	return math.MaxInt32
}

func (*EngNewOrder) ExpireTimeNullValue() int32 {
	return math.MinInt32
}

func (*EngNewOrder) SelfMatchPreventionModeId() uint16 {
	return 14
}

func (*EngNewOrder) SelfMatchPreventionModeSinceVersion() uint16 {
	return 0
}

func (e *EngNewOrder) SelfMatchPreventionModeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SelfMatchPreventionModeSinceVersion()
}

func (*EngNewOrder) SelfMatchPreventionModeDeprecated() uint16 {
	return 0
}

func (*EngNewOrder) SelfMatchPreventionModeMetaAttribute(meta int) string {
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

func (*EngNewOrder) SideId() uint16 {
	return 15
}

func (*EngNewOrder) SideSinceVersion() uint16 {
	return 0
}

func (e *EngNewOrder) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SideSinceVersion()
}

func (*EngNewOrder) SideDeprecated() uint16 {
	return 0
}

func (*EngNewOrder) SideMetaAttribute(meta int) string {
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

func (*EngNewOrder) TimeInForceId() uint16 {
	return 16
}

func (*EngNewOrder) TimeInForceSinceVersion() uint16 {
	return 0
}

func (e *EngNewOrder) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TimeInForceSinceVersion()
}

func (*EngNewOrder) TimeInForceDeprecated() uint16 {
	return 0
}

func (*EngNewOrder) TimeInForceMetaAttribute(meta int) string {
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
