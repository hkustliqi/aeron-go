// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type EngActiveOrder struct {
	OrderId                 int64
	ClusterSessionId        int64
	LastReqReceiveTime      int64
	LastUpdateTime          int64
	LastEventId             int64
	InstrumentId            int64
	UserId                  int64
	PortfolioId             int64
	ClientOrderId           [2]int64
	BrokerageClientId       [2]int64
	Quantity                int64
	LimitPrice              int64
	StopPrice               int64
	TriggerPrice            int64
	TotalFilled             int64
	FillQtyTimesRateSum     float64
	ExpireTime              int64
	OrderFlags              OrderFlags
	MatchingSortOrderId     int64
	Side                    SideEnum
	TimeInForce             TimeInForceEnum
	SelfMatchPreventionMode SelfMatchPreventionModeEnum
}

func (e *EngActiveOrder) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := e.RangeCheck(e.SbeSchemaVersion(), e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, e.OrderId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.ClusterSessionId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.LastReqReceiveTime); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.LastUpdateTime); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.LastEventId); err != nil {
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
	if err := _m.WriteInt64(_w, e.TriggerPrice); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.TotalFilled); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, e.FillQtyTimesRateSum); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.ExpireTime); err != nil {
		return err
	}
	if err := e.OrderFlags.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.MatchingSortOrderId); err != nil {
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
	return nil
}

func (e *EngActiveOrder) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !e.OrderIdInActingVersion(actingVersion) {
		e.OrderId = e.OrderIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.OrderId); err != nil {
			return err
		}
	}
	if !e.ClusterSessionIdInActingVersion(actingVersion) {
		e.ClusterSessionId = e.ClusterSessionIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.ClusterSessionId); err != nil {
			return err
		}
	}
	if !e.LastReqReceiveTimeInActingVersion(actingVersion) {
		e.LastReqReceiveTime = e.LastReqReceiveTimeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.LastReqReceiveTime); err != nil {
			return err
		}
	}
	if !e.LastUpdateTimeInActingVersion(actingVersion) {
		e.LastUpdateTime = e.LastUpdateTimeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.LastUpdateTime); err != nil {
			return err
		}
	}
	if !e.LastEventIdInActingVersion(actingVersion) {
		e.LastEventId = e.LastEventIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.LastEventId); err != nil {
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
	if !e.TriggerPriceInActingVersion(actingVersion) {
		e.TriggerPrice = e.TriggerPriceNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.TriggerPrice); err != nil {
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
	if !e.FillQtyTimesRateSumInActingVersion(actingVersion) {
		e.FillQtyTimesRateSum = e.FillQtyTimesRateSumNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &e.FillQtyTimesRateSum); err != nil {
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
	if !e.MatchingSortOrderIdInActingVersion(actingVersion) {
		e.MatchingSortOrderId = e.MatchingSortOrderIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.MatchingSortOrderId); err != nil {
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

func (e *EngActiveOrder) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.OrderIdInActingVersion(actingVersion) {
		if e.OrderId < e.OrderIdMinValue() || e.OrderId > e.OrderIdMaxValue() {
			return fmt.Errorf("Range check failed on e.OrderId (%v < %v > %v)", e.OrderIdMinValue(), e.OrderId, e.OrderIdMaxValue())
		}
	}
	if e.ClusterSessionIdInActingVersion(actingVersion) {
		if e.ClusterSessionId < e.ClusterSessionIdMinValue() || e.ClusterSessionId > e.ClusterSessionIdMaxValue() {
			return fmt.Errorf("Range check failed on e.ClusterSessionId (%v < %v > %v)", e.ClusterSessionIdMinValue(), e.ClusterSessionId, e.ClusterSessionIdMaxValue())
		}
	}
	if e.LastReqReceiveTimeInActingVersion(actingVersion) {
		if e.LastReqReceiveTime < e.LastReqReceiveTimeMinValue() || e.LastReqReceiveTime > e.LastReqReceiveTimeMaxValue() {
			return fmt.Errorf("Range check failed on e.LastReqReceiveTime (%v < %v > %v)", e.LastReqReceiveTimeMinValue(), e.LastReqReceiveTime, e.LastReqReceiveTimeMaxValue())
		}
	}
	if e.LastUpdateTimeInActingVersion(actingVersion) {
		if e.LastUpdateTime < e.LastUpdateTimeMinValue() || e.LastUpdateTime > e.LastUpdateTimeMaxValue() {
			return fmt.Errorf("Range check failed on e.LastUpdateTime (%v < %v > %v)", e.LastUpdateTimeMinValue(), e.LastUpdateTime, e.LastUpdateTimeMaxValue())
		}
	}
	if e.LastEventIdInActingVersion(actingVersion) {
		if e.LastEventId < e.LastEventIdMinValue() || e.LastEventId > e.LastEventIdMaxValue() {
			return fmt.Errorf("Range check failed on e.LastEventId (%v < %v > %v)", e.LastEventIdMinValue(), e.LastEventId, e.LastEventIdMaxValue())
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
		if e.LimitPrice < e.LimitPriceMinValue() || e.LimitPrice > e.LimitPriceMaxValue() {
			return fmt.Errorf("Range check failed on e.LimitPrice (%v < %v > %v)", e.LimitPriceMinValue(), e.LimitPrice, e.LimitPriceMaxValue())
		}
	}
	if e.StopPriceInActingVersion(actingVersion) {
		if e.StopPrice < e.StopPriceMinValue() || e.StopPrice > e.StopPriceMaxValue() {
			return fmt.Errorf("Range check failed on e.StopPrice (%v < %v > %v)", e.StopPriceMinValue(), e.StopPrice, e.StopPriceMaxValue())
		}
	}
	if e.TriggerPriceInActingVersion(actingVersion) {
		if e.TriggerPrice != e.TriggerPriceNullValue() && (e.TriggerPrice < e.TriggerPriceMinValue() || e.TriggerPrice > e.TriggerPriceMaxValue()) {
			return fmt.Errorf("Range check failed on e.TriggerPrice (%v < %v > %v)", e.TriggerPriceMinValue(), e.TriggerPrice, e.TriggerPriceMaxValue())
		}
	}
	if e.TotalFilledInActingVersion(actingVersion) {
		if e.TotalFilled < e.TotalFilledMinValue() || e.TotalFilled > e.TotalFilledMaxValue() {
			return fmt.Errorf("Range check failed on e.TotalFilled (%v < %v > %v)", e.TotalFilledMinValue(), e.TotalFilled, e.TotalFilledMaxValue())
		}
	}
	if e.FillQtyTimesRateSumInActingVersion(actingVersion) {
		if e.FillQtyTimesRateSum < e.FillQtyTimesRateSumMinValue() || e.FillQtyTimesRateSum > e.FillQtyTimesRateSumMaxValue() {
			return fmt.Errorf("Range check failed on e.FillQtyTimesRateSum (%v < %v > %v)", e.FillQtyTimesRateSumMinValue(), e.FillQtyTimesRateSum, e.FillQtyTimesRateSumMaxValue())
		}
	}
	if e.ExpireTimeInActingVersion(actingVersion) {
		if e.ExpireTime != e.ExpireTimeNullValue() && (e.ExpireTime < e.ExpireTimeMinValue() || e.ExpireTime > e.ExpireTimeMaxValue()) {
			return fmt.Errorf("Range check failed on e.ExpireTime (%v < %v > %v)", e.ExpireTimeMinValue(), e.ExpireTime, e.ExpireTimeMaxValue())
		}
	}
	if e.MatchingSortOrderIdInActingVersion(actingVersion) {
		if e.MatchingSortOrderId < e.MatchingSortOrderIdMinValue() || e.MatchingSortOrderId > e.MatchingSortOrderIdMaxValue() {
			return fmt.Errorf("Range check failed on e.MatchingSortOrderId (%v < %v > %v)", e.MatchingSortOrderIdMinValue(), e.MatchingSortOrderId, e.MatchingSortOrderIdMaxValue())
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
	return nil
}

func EngActiveOrderInit(e *EngActiveOrder) {
	e.TriggerPrice = math.MinInt64
	e.ExpireTime = math.MinInt64
	return
}

func (*EngActiveOrder) SbeBlockLength() (blockLength uint16) {
	return 171
}

func (*EngActiveOrder) SbeTemplateId() (templateId uint16) {
	return 2500
}

func (*EngActiveOrder) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*EngActiveOrder) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*EngActiveOrder) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*EngActiveOrder) OrderIdId() uint16 {
	return 1
}

func (*EngActiveOrder) OrderIdSinceVersion() uint16 {
	return 0
}

func (e *EngActiveOrder) OrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderIdSinceVersion()
}

func (*EngActiveOrder) OrderIdDeprecated() uint16 {
	return 0
}

func (*EngActiveOrder) OrderIdMetaAttribute(meta int) string {
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

func (*EngActiveOrder) OrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngActiveOrder) OrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngActiveOrder) OrderIdNullValue() int64 {
	return math.MinInt64
}

func (*EngActiveOrder) ClusterSessionIdId() uint16 {
	return 2
}

func (*EngActiveOrder) ClusterSessionIdSinceVersion() uint16 {
	return 0
}

func (e *EngActiveOrder) ClusterSessionIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ClusterSessionIdSinceVersion()
}

func (*EngActiveOrder) ClusterSessionIdDeprecated() uint16 {
	return 0
}

func (*EngActiveOrder) ClusterSessionIdMetaAttribute(meta int) string {
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

func (*EngActiveOrder) ClusterSessionIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngActiveOrder) ClusterSessionIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngActiveOrder) ClusterSessionIdNullValue() int64 {
	return math.MinInt64
}

func (*EngActiveOrder) LastReqReceiveTimeId() uint16 {
	return 3
}

func (*EngActiveOrder) LastReqReceiveTimeSinceVersion() uint16 {
	return 0
}

func (e *EngActiveOrder) LastReqReceiveTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LastReqReceiveTimeSinceVersion()
}

func (*EngActiveOrder) LastReqReceiveTimeDeprecated() uint16 {
	return 0
}

func (*EngActiveOrder) LastReqReceiveTimeMetaAttribute(meta int) string {
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

func (*EngActiveOrder) LastReqReceiveTimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngActiveOrder) LastReqReceiveTimeMaxValue() int64 {
	return math.MaxInt64
}

func (*EngActiveOrder) LastReqReceiveTimeNullValue() int64 {
	return math.MinInt64
}

func (*EngActiveOrder) LastUpdateTimeId() uint16 {
	return 4
}

func (*EngActiveOrder) LastUpdateTimeSinceVersion() uint16 {
	return 0
}

func (e *EngActiveOrder) LastUpdateTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LastUpdateTimeSinceVersion()
}

func (*EngActiveOrder) LastUpdateTimeDeprecated() uint16 {
	return 0
}

func (*EngActiveOrder) LastUpdateTimeMetaAttribute(meta int) string {
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

func (*EngActiveOrder) LastUpdateTimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngActiveOrder) LastUpdateTimeMaxValue() int64 {
	return math.MaxInt64
}

func (*EngActiveOrder) LastUpdateTimeNullValue() int64 {
	return math.MinInt64
}

func (*EngActiveOrder) LastEventIdId() uint16 {
	return 5
}

func (*EngActiveOrder) LastEventIdSinceVersion() uint16 {
	return 0
}

func (e *EngActiveOrder) LastEventIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LastEventIdSinceVersion()
}

func (*EngActiveOrder) LastEventIdDeprecated() uint16 {
	return 0
}

func (*EngActiveOrder) LastEventIdMetaAttribute(meta int) string {
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

func (*EngActiveOrder) LastEventIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngActiveOrder) LastEventIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngActiveOrder) LastEventIdNullValue() int64 {
	return math.MinInt64
}

func (*EngActiveOrder) InstrumentIdId() uint16 {
	return 6
}

func (*EngActiveOrder) InstrumentIdSinceVersion() uint16 {
	return 0
}

func (e *EngActiveOrder) InstrumentIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.InstrumentIdSinceVersion()
}

func (*EngActiveOrder) InstrumentIdDeprecated() uint16 {
	return 0
}

func (*EngActiveOrder) InstrumentIdMetaAttribute(meta int) string {
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

func (*EngActiveOrder) InstrumentIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngActiveOrder) InstrumentIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngActiveOrder) InstrumentIdNullValue() int64 {
	return math.MinInt64
}

func (*EngActiveOrder) UserIdId() uint16 {
	return 7
}

func (*EngActiveOrder) UserIdSinceVersion() uint16 {
	return 0
}

func (e *EngActiveOrder) UserIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.UserIdSinceVersion()
}

func (*EngActiveOrder) UserIdDeprecated() uint16 {
	return 0
}

func (*EngActiveOrder) UserIdMetaAttribute(meta int) string {
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

func (*EngActiveOrder) UserIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngActiveOrder) UserIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngActiveOrder) UserIdNullValue() int64 {
	return math.MinInt64
}

func (*EngActiveOrder) PortfolioIdId() uint16 {
	return 8
}

func (*EngActiveOrder) PortfolioIdSinceVersion() uint16 {
	return 0
}

func (e *EngActiveOrder) PortfolioIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PortfolioIdSinceVersion()
}

func (*EngActiveOrder) PortfolioIdDeprecated() uint16 {
	return 0
}

func (*EngActiveOrder) PortfolioIdMetaAttribute(meta int) string {
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

func (*EngActiveOrder) PortfolioIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngActiveOrder) PortfolioIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngActiveOrder) PortfolioIdNullValue() int64 {
	return math.MinInt64
}

func (*EngActiveOrder) ClientOrderIdId() uint16 {
	return 9
}

func (*EngActiveOrder) ClientOrderIdSinceVersion() uint16 {
	return 0
}

func (e *EngActiveOrder) ClientOrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ClientOrderIdSinceVersion()
}

func (*EngActiveOrder) ClientOrderIdDeprecated() uint16 {
	return 0
}

func (*EngActiveOrder) ClientOrderIdMetaAttribute(meta int) string {
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

func (*EngActiveOrder) ClientOrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngActiveOrder) ClientOrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngActiveOrder) ClientOrderIdNullValue() int64 {
	return math.MinInt64
}

func (*EngActiveOrder) BrokerageClientIdId() uint16 {
	return 10
}

func (*EngActiveOrder) BrokerageClientIdSinceVersion() uint16 {
	return 0
}

func (e *EngActiveOrder) BrokerageClientIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.BrokerageClientIdSinceVersion()
}

func (*EngActiveOrder) BrokerageClientIdDeprecated() uint16 {
	return 0
}

func (*EngActiveOrder) BrokerageClientIdMetaAttribute(meta int) string {
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

func (*EngActiveOrder) BrokerageClientIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngActiveOrder) BrokerageClientIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngActiveOrder) BrokerageClientIdNullValue() int64 {
	return math.MinInt64
}

func (*EngActiveOrder) QuantityId() uint16 {
	return 11
}

func (*EngActiveOrder) QuantitySinceVersion() uint16 {
	return 0
}

func (e *EngActiveOrder) QuantityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.QuantitySinceVersion()
}

func (*EngActiveOrder) QuantityDeprecated() uint16 {
	return 0
}

func (*EngActiveOrder) QuantityMetaAttribute(meta int) string {
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

func (*EngActiveOrder) QuantityMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngActiveOrder) QuantityMaxValue() int64 {
	return math.MaxInt64
}

func (*EngActiveOrder) QuantityNullValue() int64 {
	return math.MinInt64
}

func (*EngActiveOrder) LimitPriceId() uint16 {
	return 12
}

func (*EngActiveOrder) LimitPriceSinceVersion() uint16 {
	return 0
}

func (e *EngActiveOrder) LimitPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LimitPriceSinceVersion()
}

func (*EngActiveOrder) LimitPriceDeprecated() uint16 {
	return 0
}

func (*EngActiveOrder) LimitPriceMetaAttribute(meta int) string {
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

func (*EngActiveOrder) LimitPriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngActiveOrder) LimitPriceMaxValue() int64 {
	return math.MaxInt64
}

func (*EngActiveOrder) LimitPriceNullValue() int64 {
	return math.MinInt64
}

func (*EngActiveOrder) StopPriceId() uint16 {
	return 13
}

func (*EngActiveOrder) StopPriceSinceVersion() uint16 {
	return 0
}

func (e *EngActiveOrder) StopPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.StopPriceSinceVersion()
}

func (*EngActiveOrder) StopPriceDeprecated() uint16 {
	return 0
}

func (*EngActiveOrder) StopPriceMetaAttribute(meta int) string {
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

func (*EngActiveOrder) StopPriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngActiveOrder) StopPriceMaxValue() int64 {
	return math.MaxInt64
}

func (*EngActiveOrder) StopPriceNullValue() int64 {
	return math.MinInt64
}

func (*EngActiveOrder) TriggerPriceId() uint16 {
	return 14
}

func (*EngActiveOrder) TriggerPriceSinceVersion() uint16 {
	return 0
}

func (e *EngActiveOrder) TriggerPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TriggerPriceSinceVersion()
}

func (*EngActiveOrder) TriggerPriceDeprecated() uint16 {
	return 0
}

func (*EngActiveOrder) TriggerPriceMetaAttribute(meta int) string {
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

func (*EngActiveOrder) TriggerPriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngActiveOrder) TriggerPriceMaxValue() int64 {
	return math.MaxInt64
}

func (*EngActiveOrder) TriggerPriceNullValue() int64 {
	return math.MinInt64
}

func (*EngActiveOrder) TotalFilledId() uint16 {
	return 15
}

func (*EngActiveOrder) TotalFilledSinceVersion() uint16 {
	return 0
}

func (e *EngActiveOrder) TotalFilledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TotalFilledSinceVersion()
}

func (*EngActiveOrder) TotalFilledDeprecated() uint16 {
	return 0
}

func (*EngActiveOrder) TotalFilledMetaAttribute(meta int) string {
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

func (*EngActiveOrder) TotalFilledMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngActiveOrder) TotalFilledMaxValue() int64 {
	return math.MaxInt64
}

func (*EngActiveOrder) TotalFilledNullValue() int64 {
	return math.MinInt64
}

func (*EngActiveOrder) FillQtyTimesRateSumId() uint16 {
	return 16
}

func (*EngActiveOrder) FillQtyTimesRateSumSinceVersion() uint16 {
	return 0
}

func (e *EngActiveOrder) FillQtyTimesRateSumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FillQtyTimesRateSumSinceVersion()
}

func (*EngActiveOrder) FillQtyTimesRateSumDeprecated() uint16 {
	return 0
}

func (*EngActiveOrder) FillQtyTimesRateSumMetaAttribute(meta int) string {
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

func (*EngActiveOrder) FillQtyTimesRateSumMinValue() float64 {
	return -math.MaxFloat64
}

func (*EngActiveOrder) FillQtyTimesRateSumMaxValue() float64 {
	return math.MaxFloat64
}

func (*EngActiveOrder) FillQtyTimesRateSumNullValue() float64 {
	return math.NaN()
}

func (*EngActiveOrder) ExpireTimeId() uint16 {
	return 17
}

func (*EngActiveOrder) ExpireTimeSinceVersion() uint16 {
	return 0
}

func (e *EngActiveOrder) ExpireTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExpireTimeSinceVersion()
}

func (*EngActiveOrder) ExpireTimeDeprecated() uint16 {
	return 0
}

func (*EngActiveOrder) ExpireTimeMetaAttribute(meta int) string {
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

func (*EngActiveOrder) ExpireTimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngActiveOrder) ExpireTimeMaxValue() int64 {
	return math.MaxInt64
}

func (*EngActiveOrder) ExpireTimeNullValue() int64 {
	return math.MinInt64
}

func (*EngActiveOrder) OrderFlagsId() uint16 {
	return 18
}

func (*EngActiveOrder) OrderFlagsSinceVersion() uint16 {
	return 0
}

func (e *EngActiveOrder) OrderFlagsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderFlagsSinceVersion()
}

func (*EngActiveOrder) OrderFlagsDeprecated() uint16 {
	return 0
}

func (*EngActiveOrder) OrderFlagsMetaAttribute(meta int) string {
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

func (*EngActiveOrder) MatchingSortOrderIdId() uint16 {
	return 19
}

func (*EngActiveOrder) MatchingSortOrderIdSinceVersion() uint16 {
	return 0
}

func (e *EngActiveOrder) MatchingSortOrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MatchingSortOrderIdSinceVersion()
}

func (*EngActiveOrder) MatchingSortOrderIdDeprecated() uint16 {
	return 0
}

func (*EngActiveOrder) MatchingSortOrderIdMetaAttribute(meta int) string {
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

func (*EngActiveOrder) MatchingSortOrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngActiveOrder) MatchingSortOrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngActiveOrder) MatchingSortOrderIdNullValue() int64 {
	return math.MinInt64
}

func (*EngActiveOrder) SideId() uint16 {
	return 20
}

func (*EngActiveOrder) SideSinceVersion() uint16 {
	return 0
}

func (e *EngActiveOrder) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SideSinceVersion()
}

func (*EngActiveOrder) SideDeprecated() uint16 {
	return 0
}

func (*EngActiveOrder) SideMetaAttribute(meta int) string {
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

func (*EngActiveOrder) TimeInForceId() uint16 {
	return 21
}

func (*EngActiveOrder) TimeInForceSinceVersion() uint16 {
	return 0
}

func (e *EngActiveOrder) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TimeInForceSinceVersion()
}

func (*EngActiveOrder) TimeInForceDeprecated() uint16 {
	return 0
}

func (*EngActiveOrder) TimeInForceMetaAttribute(meta int) string {
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

func (*EngActiveOrder) SelfMatchPreventionModeId() uint16 {
	return 22
}

func (*EngActiveOrder) SelfMatchPreventionModeSinceVersion() uint16 {
	return 0
}

func (e *EngActiveOrder) SelfMatchPreventionModeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SelfMatchPreventionModeSinceVersion()
}

func (*EngActiveOrder) SelfMatchPreventionModeDeprecated() uint16 {
	return 0
}

func (*EngActiveOrder) SelfMatchPreventionModeMetaAttribute(meta int) string {
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
