// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OmsNewOrder struct {
	RequestHeader           OmsUserRequestHeader
	InstrumentId            int64
	ClientOrderId           [2]int64
	PortfolioId             int64
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

func (o *OmsNewOrder) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := o.RequestHeader.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.InstrumentId); err != nil {
		return err
	}
	for idx := 0; idx < 2; idx++ {
		if err := _m.WriteInt64(_w, o.ClientOrderId[idx]); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, o.PortfolioId); err != nil {
		return err
	}
	for idx := 0; idx < 2; idx++ {
		if err := _m.WriteInt64(_w, o.BrokerageClientId[idx]); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, o.LimitPrice); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.StopPrice); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.Quantity); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.Funds); err != nil {
		return err
	}
	if err := o.OrderFlags.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, o.ExpireTime); err != nil {
		return err
	}
	if err := o.SelfMatchPreventionMode.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.Side.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.TimeInForce.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (o *OmsNewOrder) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if o.RequestHeaderInActingVersion(actingVersion) {
		if err := o.RequestHeader.Decode(_m, _r, actingVersion); err != nil {
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
	if !o.PortfolioIdInActingVersion(actingVersion) {
		o.PortfolioId = o.PortfolioIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.PortfolioId); err != nil {
			return err
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
	if !o.QuantityInActingVersion(actingVersion) {
		o.Quantity = o.QuantityNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.Quantity); err != nil {
			return err
		}
	}
	if !o.FundsInActingVersion(actingVersion) {
		o.Funds = o.FundsNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.Funds); err != nil {
			return err
		}
	}
	if o.OrderFlagsInActingVersion(actingVersion) {
		if err := o.OrderFlags.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.ExpireTimeInActingVersion(actingVersion) {
		o.ExpireTime = o.ExpireTimeNullValue()
	} else {
		if err := _m.ReadInt32(_r, &o.ExpireTime); err != nil {
			return err
		}
	}
	if o.SelfMatchPreventionModeInActingVersion(actingVersion) {
		if err := o.SelfMatchPreventionMode.Decode(_m, _r, actingVersion); err != nil {
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

func (o *OmsNewOrder) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.InstrumentIdInActingVersion(actingVersion) {
		if o.InstrumentId < o.InstrumentIdMinValue() || o.InstrumentId > o.InstrumentIdMaxValue() {
			return fmt.Errorf("Range check failed on o.InstrumentId (%v < %v > %v)", o.InstrumentIdMinValue(), o.InstrumentId, o.InstrumentIdMaxValue())
		}
	}
	if o.ClientOrderIdInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if o.ClientOrderId[idx] < o.ClientOrderIdMinValue() || o.ClientOrderId[idx] > o.ClientOrderIdMaxValue() {
				return fmt.Errorf("Range check failed on o.ClientOrderId[%d] (%v < %v > %v)", idx, o.ClientOrderIdMinValue(), o.ClientOrderId[idx], o.ClientOrderIdMaxValue())
			}
		}
	}
	if o.PortfolioIdInActingVersion(actingVersion) {
		if o.PortfolioId < o.PortfolioIdMinValue() || o.PortfolioId > o.PortfolioIdMaxValue() {
			return fmt.Errorf("Range check failed on o.PortfolioId (%v < %v > %v)", o.PortfolioIdMinValue(), o.PortfolioId, o.PortfolioIdMaxValue())
		}
	}
	if o.BrokerageClientIdInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if o.BrokerageClientId[idx] < o.BrokerageClientIdMinValue() || o.BrokerageClientId[idx] > o.BrokerageClientIdMaxValue() {
				return fmt.Errorf("Range check failed on o.BrokerageClientId[%d] (%v < %v > %v)", idx, o.BrokerageClientIdMinValue(), o.BrokerageClientId[idx], o.BrokerageClientIdMaxValue())
			}
		}
	}
	if o.LimitPriceInActingVersion(actingVersion) {
		if o.LimitPrice < o.LimitPriceMinValue() || o.LimitPrice > o.LimitPriceMaxValue() {
			return fmt.Errorf("Range check failed on o.LimitPrice (%v < %v > %v)", o.LimitPriceMinValue(), o.LimitPrice, o.LimitPriceMaxValue())
		}
	}
	if o.StopPriceInActingVersion(actingVersion) {
		if o.StopPrice != o.StopPriceNullValue() && (o.StopPrice < o.StopPriceMinValue() || o.StopPrice > o.StopPriceMaxValue()) {
			return fmt.Errorf("Range check failed on o.StopPrice (%v < %v > %v)", o.StopPriceMinValue(), o.StopPrice, o.StopPriceMaxValue())
		}
	}
	if o.QuantityInActingVersion(actingVersion) {
		if o.Quantity < o.QuantityMinValue() || o.Quantity > o.QuantityMaxValue() {
			return fmt.Errorf("Range check failed on o.Quantity (%v < %v > %v)", o.QuantityMinValue(), o.Quantity, o.QuantityMaxValue())
		}
	}
	if o.FundsInActingVersion(actingVersion) {
		if o.Funds != o.FundsNullValue() && (o.Funds < o.FundsMinValue() || o.Funds > o.FundsMaxValue()) {
			return fmt.Errorf("Range check failed on o.Funds (%v < %v > %v)", o.FundsMinValue(), o.Funds, o.FundsMaxValue())
		}
	}
	if o.ExpireTimeInActingVersion(actingVersion) {
		if o.ExpireTime != o.ExpireTimeNullValue() && (o.ExpireTime < o.ExpireTimeMinValue() || o.ExpireTime > o.ExpireTimeMaxValue()) {
			return fmt.Errorf("Range check failed on o.ExpireTime (%v < %v > %v)", o.ExpireTimeMinValue(), o.ExpireTime, o.ExpireTimeMaxValue())
		}
	}
	if err := o.SelfMatchPreventionMode.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.TimeInForce.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func OmsNewOrderInit(o *OmsNewOrder) {
	o.StopPrice = math.MinInt64
	o.Funds = math.MinInt64
	o.ExpireTime = math.MinInt32
	return
}

func (*OmsNewOrder) SbeBlockLength() (blockLength uint16) {
	return 127
}

func (*OmsNewOrder) SbeTemplateId() (templateId uint16) {
	return 1000
}

func (*OmsNewOrder) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*OmsNewOrder) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*OmsNewOrder) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*OmsNewOrder) RequestHeaderId() uint16 {
	return 1
}

func (*OmsNewOrder) RequestHeaderSinceVersion() uint16 {
	return 0
}

func (o *OmsNewOrder) RequestHeaderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.RequestHeaderSinceVersion()
}

func (*OmsNewOrder) RequestHeaderDeprecated() uint16 {
	return 0
}

func (*OmsNewOrder) RequestHeaderMetaAttribute(meta int) string {
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

func (*OmsNewOrder) InstrumentIdId() uint16 {
	return 2
}

func (*OmsNewOrder) InstrumentIdSinceVersion() uint16 {
	return 0
}

func (o *OmsNewOrder) InstrumentIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.InstrumentIdSinceVersion()
}

func (*OmsNewOrder) InstrumentIdDeprecated() uint16 {
	return 0
}

func (*OmsNewOrder) InstrumentIdMetaAttribute(meta int) string {
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

func (*OmsNewOrder) InstrumentIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsNewOrder) InstrumentIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsNewOrder) InstrumentIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsNewOrder) ClientOrderIdId() uint16 {
	return 3
}

func (*OmsNewOrder) ClientOrderIdSinceVersion() uint16 {
	return 0
}

func (o *OmsNewOrder) ClientOrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ClientOrderIdSinceVersion()
}

func (*OmsNewOrder) ClientOrderIdDeprecated() uint16 {
	return 0
}

func (*OmsNewOrder) ClientOrderIdMetaAttribute(meta int) string {
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

func (*OmsNewOrder) ClientOrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsNewOrder) ClientOrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsNewOrder) ClientOrderIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsNewOrder) PortfolioIdId() uint16 {
	return 4
}

func (*OmsNewOrder) PortfolioIdSinceVersion() uint16 {
	return 0
}

func (o *OmsNewOrder) PortfolioIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PortfolioIdSinceVersion()
}

func (*OmsNewOrder) PortfolioIdDeprecated() uint16 {
	return 0
}

func (*OmsNewOrder) PortfolioIdMetaAttribute(meta int) string {
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

func (*OmsNewOrder) PortfolioIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsNewOrder) PortfolioIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsNewOrder) PortfolioIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsNewOrder) BrokerageClientIdId() uint16 {
	return 5
}

func (*OmsNewOrder) BrokerageClientIdSinceVersion() uint16 {
	return 0
}

func (o *OmsNewOrder) BrokerageClientIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.BrokerageClientIdSinceVersion()
}

func (*OmsNewOrder) BrokerageClientIdDeprecated() uint16 {
	return 0
}

func (*OmsNewOrder) BrokerageClientIdMetaAttribute(meta int) string {
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

func (*OmsNewOrder) BrokerageClientIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsNewOrder) BrokerageClientIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsNewOrder) BrokerageClientIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsNewOrder) LimitPriceId() uint16 {
	return 6
}

func (*OmsNewOrder) LimitPriceSinceVersion() uint16 {
	return 0
}

func (o *OmsNewOrder) LimitPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LimitPriceSinceVersion()
}

func (*OmsNewOrder) LimitPriceDeprecated() uint16 {
	return 0
}

func (*OmsNewOrder) LimitPriceMetaAttribute(meta int) string {
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

func (*OmsNewOrder) LimitPriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsNewOrder) LimitPriceMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsNewOrder) LimitPriceNullValue() int64 {
	return math.MinInt64
}

func (*OmsNewOrder) StopPriceId() uint16 {
	return 7
}

func (*OmsNewOrder) StopPriceSinceVersion() uint16 {
	return 0
}

func (o *OmsNewOrder) StopPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StopPriceSinceVersion()
}

func (*OmsNewOrder) StopPriceDeprecated() uint16 {
	return 0
}

func (*OmsNewOrder) StopPriceMetaAttribute(meta int) string {
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

func (*OmsNewOrder) StopPriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsNewOrder) StopPriceMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsNewOrder) StopPriceNullValue() int64 {
	return math.MinInt64
}

func (*OmsNewOrder) QuantityId() uint16 {
	return 8
}

func (*OmsNewOrder) QuantitySinceVersion() uint16 {
	return 0
}

func (o *OmsNewOrder) QuantityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.QuantitySinceVersion()
}

func (*OmsNewOrder) QuantityDeprecated() uint16 {
	return 0
}

func (*OmsNewOrder) QuantityMetaAttribute(meta int) string {
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

func (*OmsNewOrder) QuantityMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsNewOrder) QuantityMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsNewOrder) QuantityNullValue() int64 {
	return math.MinInt64
}

func (*OmsNewOrder) FundsId() uint16 {
	return 9
}

func (*OmsNewOrder) FundsSinceVersion() uint16 {
	return 0
}

func (o *OmsNewOrder) FundsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.FundsSinceVersion()
}

func (*OmsNewOrder) FundsDeprecated() uint16 {
	return 0
}

func (*OmsNewOrder) FundsMetaAttribute(meta int) string {
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

func (*OmsNewOrder) FundsMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsNewOrder) FundsMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsNewOrder) FundsNullValue() int64 {
	return math.MinInt64
}

func (*OmsNewOrder) OrderFlagsId() uint16 {
	return 10
}

func (*OmsNewOrder) OrderFlagsSinceVersion() uint16 {
	return 0
}

func (o *OmsNewOrder) OrderFlagsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderFlagsSinceVersion()
}

func (*OmsNewOrder) OrderFlagsDeprecated() uint16 {
	return 0
}

func (*OmsNewOrder) OrderFlagsMetaAttribute(meta int) string {
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

func (*OmsNewOrder) ExpireTimeId() uint16 {
	return 11
}

func (*OmsNewOrder) ExpireTimeSinceVersion() uint16 {
	return 0
}

func (o *OmsNewOrder) ExpireTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ExpireTimeSinceVersion()
}

func (*OmsNewOrder) ExpireTimeDeprecated() uint16 {
	return 0
}

func (*OmsNewOrder) ExpireTimeMetaAttribute(meta int) string {
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

func (*OmsNewOrder) ExpireTimeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*OmsNewOrder) ExpireTimeMaxValue() int32 {
	return math.MaxInt32
}

func (*OmsNewOrder) ExpireTimeNullValue() int32 {
	return math.MinInt32
}

func (*OmsNewOrder) SelfMatchPreventionModeId() uint16 {
	return 12
}

func (*OmsNewOrder) SelfMatchPreventionModeSinceVersion() uint16 {
	return 0
}

func (o *OmsNewOrder) SelfMatchPreventionModeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SelfMatchPreventionModeSinceVersion()
}

func (*OmsNewOrder) SelfMatchPreventionModeDeprecated() uint16 {
	return 0
}

func (*OmsNewOrder) SelfMatchPreventionModeMetaAttribute(meta int) string {
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

func (*OmsNewOrder) SideId() uint16 {
	return 13
}

func (*OmsNewOrder) SideSinceVersion() uint16 {
	return 0
}

func (o *OmsNewOrder) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SideSinceVersion()
}

func (*OmsNewOrder) SideDeprecated() uint16 {
	return 0
}

func (*OmsNewOrder) SideMetaAttribute(meta int) string {
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

func (*OmsNewOrder) TimeInForceId() uint16 {
	return 14
}

func (*OmsNewOrder) TimeInForceSinceVersion() uint16 {
	return 0
}

func (o *OmsNewOrder) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TimeInForceSinceVersion()
}

func (*OmsNewOrder) TimeInForceDeprecated() uint16 {
	return 0
}

func (*OmsNewOrder) TimeInForceMetaAttribute(meta int) string {
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
