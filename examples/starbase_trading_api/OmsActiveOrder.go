// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OmsActiveOrder struct {
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
	FilledVwap              int64
	ExpireTime              int64
	OrderFlags              OrderFlags
	Side                    SideEnum
	TimeInForce             TimeInForceEnum
	SelfMatchPreventionMode SelfMatchPreventionModeEnum
}

func (o *OmsActiveOrder) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, o.OrderId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.ClusterSessionId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.LastReqReceiveTime); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.LastUpdateTime); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.LastEventId); err != nil {
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
	if err := _m.WriteInt64(_w, o.TriggerPrice); err != nil {
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
	return nil
}

func (o *OmsActiveOrder) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !o.OrderIdInActingVersion(actingVersion) {
		o.OrderId = o.OrderIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.OrderId); err != nil {
			return err
		}
	}
	if !o.ClusterSessionIdInActingVersion(actingVersion) {
		o.ClusterSessionId = o.ClusterSessionIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.ClusterSessionId); err != nil {
			return err
		}
	}
	if !o.LastReqReceiveTimeInActingVersion(actingVersion) {
		o.LastReqReceiveTime = o.LastReqReceiveTimeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.LastReqReceiveTime); err != nil {
			return err
		}
	}
	if !o.LastUpdateTimeInActingVersion(actingVersion) {
		o.LastUpdateTime = o.LastUpdateTimeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.LastUpdateTime); err != nil {
			return err
		}
	}
	if !o.LastEventIdInActingVersion(actingVersion) {
		o.LastEventId = o.LastEventIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.LastEventId); err != nil {
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
	if !o.TriggerPriceInActingVersion(actingVersion) {
		o.TriggerPrice = o.TriggerPriceNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.TriggerPrice); err != nil {
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

func (o *OmsActiveOrder) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.OrderIdInActingVersion(actingVersion) {
		if o.OrderId < o.OrderIdMinValue() || o.OrderId > o.OrderIdMaxValue() {
			return fmt.Errorf("Range check failed on o.OrderId (%v < %v > %v)", o.OrderIdMinValue(), o.OrderId, o.OrderIdMaxValue())
		}
	}
	if o.ClusterSessionIdInActingVersion(actingVersion) {
		if o.ClusterSessionId < o.ClusterSessionIdMinValue() || o.ClusterSessionId > o.ClusterSessionIdMaxValue() {
			return fmt.Errorf("Range check failed on o.ClusterSessionId (%v < %v > %v)", o.ClusterSessionIdMinValue(), o.ClusterSessionId, o.ClusterSessionIdMaxValue())
		}
	}
	if o.LastReqReceiveTimeInActingVersion(actingVersion) {
		if o.LastReqReceiveTime < o.LastReqReceiveTimeMinValue() || o.LastReqReceiveTime > o.LastReqReceiveTimeMaxValue() {
			return fmt.Errorf("Range check failed on o.LastReqReceiveTime (%v < %v > %v)", o.LastReqReceiveTimeMinValue(), o.LastReqReceiveTime, o.LastReqReceiveTimeMaxValue())
		}
	}
	if o.LastUpdateTimeInActingVersion(actingVersion) {
		if o.LastUpdateTime < o.LastUpdateTimeMinValue() || o.LastUpdateTime > o.LastUpdateTimeMaxValue() {
			return fmt.Errorf("Range check failed on o.LastUpdateTime (%v < %v > %v)", o.LastUpdateTimeMinValue(), o.LastUpdateTime, o.LastUpdateTimeMaxValue())
		}
	}
	if o.LastEventIdInActingVersion(actingVersion) {
		if o.LastEventId < o.LastEventIdMinValue() || o.LastEventId > o.LastEventIdMaxValue() {
			return fmt.Errorf("Range check failed on o.LastEventId (%v < %v > %v)", o.LastEventIdMinValue(), o.LastEventId, o.LastEventIdMaxValue())
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
		if o.LimitPrice < o.LimitPriceMinValue() || o.LimitPrice > o.LimitPriceMaxValue() {
			return fmt.Errorf("Range check failed on o.LimitPrice (%v < %v > %v)", o.LimitPriceMinValue(), o.LimitPrice, o.LimitPriceMaxValue())
		}
	}
	if o.StopPriceInActingVersion(actingVersion) {
		if o.StopPrice < o.StopPriceMinValue() || o.StopPrice > o.StopPriceMaxValue() {
			return fmt.Errorf("Range check failed on o.StopPrice (%v < %v > %v)", o.StopPriceMinValue(), o.StopPrice, o.StopPriceMaxValue())
		}
	}
	if o.TriggerPriceInActingVersion(actingVersion) {
		if o.TriggerPrice != o.TriggerPriceNullValue() && (o.TriggerPrice < o.TriggerPriceMinValue() || o.TriggerPrice > o.TriggerPriceMaxValue()) {
			return fmt.Errorf("Range check failed on o.TriggerPrice (%v < %v > %v)", o.TriggerPriceMinValue(), o.TriggerPrice, o.TriggerPriceMaxValue())
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
	return nil
}

func OmsActiveOrderInit(o *OmsActiveOrder) {
	o.TriggerPrice = math.MinInt64
	o.ExpireTime = math.MinInt64
	return
}

func (*OmsActiveOrder) SbeBlockLength() (blockLength uint16) {
	return 163
}

func (*OmsActiveOrder) SbeTemplateId() (templateId uint16) {
	return 1500
}

func (*OmsActiveOrder) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*OmsActiveOrder) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*OmsActiveOrder) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*OmsActiveOrder) OrderIdId() uint16 {
	return 1
}

func (*OmsActiveOrder) OrderIdSinceVersion() uint16 {
	return 0
}

func (o *OmsActiveOrder) OrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderIdSinceVersion()
}

func (*OmsActiveOrder) OrderIdDeprecated() uint16 {
	return 0
}

func (*OmsActiveOrder) OrderIdMetaAttribute(meta int) string {
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

func (*OmsActiveOrder) OrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsActiveOrder) OrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsActiveOrder) OrderIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsActiveOrder) ClusterSessionIdId() uint16 {
	return 2
}

func (*OmsActiveOrder) ClusterSessionIdSinceVersion() uint16 {
	return 0
}

func (o *OmsActiveOrder) ClusterSessionIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ClusterSessionIdSinceVersion()
}

func (*OmsActiveOrder) ClusterSessionIdDeprecated() uint16 {
	return 0
}

func (*OmsActiveOrder) ClusterSessionIdMetaAttribute(meta int) string {
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

func (*OmsActiveOrder) ClusterSessionIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsActiveOrder) ClusterSessionIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsActiveOrder) ClusterSessionIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsActiveOrder) LastReqReceiveTimeId() uint16 {
	return 3
}

func (*OmsActiveOrder) LastReqReceiveTimeSinceVersion() uint16 {
	return 0
}

func (o *OmsActiveOrder) LastReqReceiveTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LastReqReceiveTimeSinceVersion()
}

func (*OmsActiveOrder) LastReqReceiveTimeDeprecated() uint16 {
	return 0
}

func (*OmsActiveOrder) LastReqReceiveTimeMetaAttribute(meta int) string {
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

func (*OmsActiveOrder) LastReqReceiveTimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsActiveOrder) LastReqReceiveTimeMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsActiveOrder) LastReqReceiveTimeNullValue() int64 {
	return math.MinInt64
}

func (*OmsActiveOrder) LastUpdateTimeId() uint16 {
	return 4
}

func (*OmsActiveOrder) LastUpdateTimeSinceVersion() uint16 {
	return 0
}

func (o *OmsActiveOrder) LastUpdateTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LastUpdateTimeSinceVersion()
}

func (*OmsActiveOrder) LastUpdateTimeDeprecated() uint16 {
	return 0
}

func (*OmsActiveOrder) LastUpdateTimeMetaAttribute(meta int) string {
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

func (*OmsActiveOrder) LastUpdateTimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsActiveOrder) LastUpdateTimeMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsActiveOrder) LastUpdateTimeNullValue() int64 {
	return math.MinInt64
}

func (*OmsActiveOrder) LastEventIdId() uint16 {
	return 5
}

func (*OmsActiveOrder) LastEventIdSinceVersion() uint16 {
	return 0
}

func (o *OmsActiveOrder) LastEventIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LastEventIdSinceVersion()
}

func (*OmsActiveOrder) LastEventIdDeprecated() uint16 {
	return 0
}

func (*OmsActiveOrder) LastEventIdMetaAttribute(meta int) string {
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

func (*OmsActiveOrder) LastEventIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsActiveOrder) LastEventIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsActiveOrder) LastEventIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsActiveOrder) InstrumentIdId() uint16 {
	return 6
}

func (*OmsActiveOrder) InstrumentIdSinceVersion() uint16 {
	return 0
}

func (o *OmsActiveOrder) InstrumentIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.InstrumentIdSinceVersion()
}

func (*OmsActiveOrder) InstrumentIdDeprecated() uint16 {
	return 0
}

func (*OmsActiveOrder) InstrumentIdMetaAttribute(meta int) string {
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

func (*OmsActiveOrder) InstrumentIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsActiveOrder) InstrumentIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsActiveOrder) InstrumentIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsActiveOrder) UserIdId() uint16 {
	return 7
}

func (*OmsActiveOrder) UserIdSinceVersion() uint16 {
	return 0
}

func (o *OmsActiveOrder) UserIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.UserIdSinceVersion()
}

func (*OmsActiveOrder) UserIdDeprecated() uint16 {
	return 0
}

func (*OmsActiveOrder) UserIdMetaAttribute(meta int) string {
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

func (*OmsActiveOrder) UserIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsActiveOrder) UserIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsActiveOrder) UserIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsActiveOrder) PortfolioIdId() uint16 {
	return 8
}

func (*OmsActiveOrder) PortfolioIdSinceVersion() uint16 {
	return 0
}

func (o *OmsActiveOrder) PortfolioIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PortfolioIdSinceVersion()
}

func (*OmsActiveOrder) PortfolioIdDeprecated() uint16 {
	return 0
}

func (*OmsActiveOrder) PortfolioIdMetaAttribute(meta int) string {
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

func (*OmsActiveOrder) PortfolioIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsActiveOrder) PortfolioIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsActiveOrder) PortfolioIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsActiveOrder) ClientOrderIdId() uint16 {
	return 9
}

func (*OmsActiveOrder) ClientOrderIdSinceVersion() uint16 {
	return 0
}

func (o *OmsActiveOrder) ClientOrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ClientOrderIdSinceVersion()
}

func (*OmsActiveOrder) ClientOrderIdDeprecated() uint16 {
	return 0
}

func (*OmsActiveOrder) ClientOrderIdMetaAttribute(meta int) string {
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

func (*OmsActiveOrder) ClientOrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsActiveOrder) ClientOrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsActiveOrder) ClientOrderIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsActiveOrder) BrokerageClientIdId() uint16 {
	return 10
}

func (*OmsActiveOrder) BrokerageClientIdSinceVersion() uint16 {
	return 0
}

func (o *OmsActiveOrder) BrokerageClientIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.BrokerageClientIdSinceVersion()
}

func (*OmsActiveOrder) BrokerageClientIdDeprecated() uint16 {
	return 0
}

func (*OmsActiveOrder) BrokerageClientIdMetaAttribute(meta int) string {
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

func (*OmsActiveOrder) BrokerageClientIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsActiveOrder) BrokerageClientIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsActiveOrder) BrokerageClientIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsActiveOrder) QuantityId() uint16 {
	return 11
}

func (*OmsActiveOrder) QuantitySinceVersion() uint16 {
	return 0
}

func (o *OmsActiveOrder) QuantityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.QuantitySinceVersion()
}

func (*OmsActiveOrder) QuantityDeprecated() uint16 {
	return 0
}

func (*OmsActiveOrder) QuantityMetaAttribute(meta int) string {
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

func (*OmsActiveOrder) QuantityMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsActiveOrder) QuantityMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsActiveOrder) QuantityNullValue() int64 {
	return math.MinInt64
}

func (*OmsActiveOrder) LimitPriceId() uint16 {
	return 12
}

func (*OmsActiveOrder) LimitPriceSinceVersion() uint16 {
	return 0
}

func (o *OmsActiveOrder) LimitPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LimitPriceSinceVersion()
}

func (*OmsActiveOrder) LimitPriceDeprecated() uint16 {
	return 0
}

func (*OmsActiveOrder) LimitPriceMetaAttribute(meta int) string {
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

func (*OmsActiveOrder) LimitPriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsActiveOrder) LimitPriceMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsActiveOrder) LimitPriceNullValue() int64 {
	return math.MinInt64
}

func (*OmsActiveOrder) StopPriceId() uint16 {
	return 13
}

func (*OmsActiveOrder) StopPriceSinceVersion() uint16 {
	return 0
}

func (o *OmsActiveOrder) StopPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StopPriceSinceVersion()
}

func (*OmsActiveOrder) StopPriceDeprecated() uint16 {
	return 0
}

func (*OmsActiveOrder) StopPriceMetaAttribute(meta int) string {
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

func (*OmsActiveOrder) StopPriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsActiveOrder) StopPriceMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsActiveOrder) StopPriceNullValue() int64 {
	return math.MinInt64
}

func (*OmsActiveOrder) TriggerPriceId() uint16 {
	return 14
}

func (*OmsActiveOrder) TriggerPriceSinceVersion() uint16 {
	return 0
}

func (o *OmsActiveOrder) TriggerPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TriggerPriceSinceVersion()
}

func (*OmsActiveOrder) TriggerPriceDeprecated() uint16 {
	return 0
}

func (*OmsActiveOrder) TriggerPriceMetaAttribute(meta int) string {
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

func (*OmsActiveOrder) TriggerPriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsActiveOrder) TriggerPriceMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsActiveOrder) TriggerPriceNullValue() int64 {
	return math.MinInt64
}

func (*OmsActiveOrder) TotalFilledId() uint16 {
	return 15
}

func (*OmsActiveOrder) TotalFilledSinceVersion() uint16 {
	return 0
}

func (o *OmsActiveOrder) TotalFilledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TotalFilledSinceVersion()
}

func (*OmsActiveOrder) TotalFilledDeprecated() uint16 {
	return 0
}

func (*OmsActiveOrder) TotalFilledMetaAttribute(meta int) string {
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

func (*OmsActiveOrder) TotalFilledMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsActiveOrder) TotalFilledMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsActiveOrder) TotalFilledNullValue() int64 {
	return math.MinInt64
}

func (*OmsActiveOrder) FilledVwapId() uint16 {
	return 16
}

func (*OmsActiveOrder) FilledVwapSinceVersion() uint16 {
	return 0
}

func (o *OmsActiveOrder) FilledVwapInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.FilledVwapSinceVersion()
}

func (*OmsActiveOrder) FilledVwapDeprecated() uint16 {
	return 0
}

func (*OmsActiveOrder) FilledVwapMetaAttribute(meta int) string {
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

func (*OmsActiveOrder) FilledVwapMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsActiveOrder) FilledVwapMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsActiveOrder) FilledVwapNullValue() int64 {
	return math.MinInt64
}

func (*OmsActiveOrder) ExpireTimeId() uint16 {
	return 17
}

func (*OmsActiveOrder) ExpireTimeSinceVersion() uint16 {
	return 0
}

func (o *OmsActiveOrder) ExpireTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ExpireTimeSinceVersion()
}

func (*OmsActiveOrder) ExpireTimeDeprecated() uint16 {
	return 0
}

func (*OmsActiveOrder) ExpireTimeMetaAttribute(meta int) string {
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

func (*OmsActiveOrder) ExpireTimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsActiveOrder) ExpireTimeMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsActiveOrder) ExpireTimeNullValue() int64 {
	return math.MinInt64
}

func (*OmsActiveOrder) OrderFlagsId() uint16 {
	return 18
}

func (*OmsActiveOrder) OrderFlagsSinceVersion() uint16 {
	return 0
}

func (o *OmsActiveOrder) OrderFlagsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderFlagsSinceVersion()
}

func (*OmsActiveOrder) OrderFlagsDeprecated() uint16 {
	return 0
}

func (*OmsActiveOrder) OrderFlagsMetaAttribute(meta int) string {
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

func (*OmsActiveOrder) SideId() uint16 {
	return 19
}

func (*OmsActiveOrder) SideSinceVersion() uint16 {
	return 0
}

func (o *OmsActiveOrder) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SideSinceVersion()
}

func (*OmsActiveOrder) SideDeprecated() uint16 {
	return 0
}

func (*OmsActiveOrder) SideMetaAttribute(meta int) string {
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

func (*OmsActiveOrder) TimeInForceId() uint16 {
	return 20
}

func (*OmsActiveOrder) TimeInForceSinceVersion() uint16 {
	return 0
}

func (o *OmsActiveOrder) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TimeInForceSinceVersion()
}

func (*OmsActiveOrder) TimeInForceDeprecated() uint16 {
	return 0
}

func (*OmsActiveOrder) TimeInForceMetaAttribute(meta int) string {
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

func (*OmsActiveOrder) SelfMatchPreventionModeId() uint16 {
	return 21
}

func (*OmsActiveOrder) SelfMatchPreventionModeSinceVersion() uint16 {
	return 0
}

func (o *OmsActiveOrder) SelfMatchPreventionModeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SelfMatchPreventionModeSinceVersion()
}

func (*OmsActiveOrder) SelfMatchPreventionModeDeprecated() uint16 {
	return 0
}

func (*OmsActiveOrder) SelfMatchPreventionModeMetaAttribute(meta int) string {
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
