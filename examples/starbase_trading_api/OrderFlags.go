// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"io"
)

type OrderFlags [64]bool
type OrderFlagsChoiceValue uint8
type OrderFlagsChoiceValues struct {
	IsMakerOrder         OrderFlagsChoiceValue
	IsLeadMakerOrder     OrderFlagsChoiceValue
	IsMarketDataVisible  OrderFlagsChoiceValue
	IsWorking            OrderFlagsChoiceValue
	CancelOnDisconnect   OrderFlagsChoiceValue
	PostOnly             OrderFlagsChoiceValue
	IsAggressor          OrderFlagsChoiceValue
	IsStopTriggerPending OrderFlagsChoiceValue
	IsCancelPending      OrderFlagsChoiceValue
}

var OrderFlagsChoice = OrderFlagsChoiceValues{OrderFlagsChoiceValue(0), OrderFlagsChoiceValue(1), OrderFlagsChoiceValue(2), OrderFlagsChoiceValue(3), OrderFlagsChoiceValue(4), OrderFlagsChoiceValue(5), OrderFlagsChoiceValue(6), OrderFlagsChoiceValue(7), OrderFlagsChoiceValue(8)}

func (o *OrderFlags) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	var wireval uint64 = 0
	for k, v := range o {
		if v {
			wireval |= (1 << uint(k))
		}
	}
	return _m.WriteUint64(_w, wireval)
}

func (o *OrderFlags) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	var wireval uint64

	if err := _m.ReadUint64(_r, &wireval); err != nil {
		return err
	}

	var idx uint
	for idx = 0; idx < 64; idx++ {
		o[idx] = (wireval & (1 << idx)) > 0
	}
	return nil
}

func (OrderFlags) EncodedLength() int64 {
	return 8
}

func (*OrderFlags) isMakerOrderSinceVersion() uint16 {
	return 0
}

func (o *OrderFlags) isMakerOrderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.isMakerOrderSinceVersion()
}

func (*OrderFlags) isMakerOrderDeprecated() uint16 {
	return 0
}

func (*OrderFlags) isLeadMakerOrderSinceVersion() uint16 {
	return 0
}

func (o *OrderFlags) isLeadMakerOrderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.isLeadMakerOrderSinceVersion()
}

func (*OrderFlags) isLeadMakerOrderDeprecated() uint16 {
	return 0
}

func (*OrderFlags) isMarketDataVisibleSinceVersion() uint16 {
	return 0
}

func (o *OrderFlags) isMarketDataVisibleInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.isMarketDataVisibleSinceVersion()
}

func (*OrderFlags) isMarketDataVisibleDeprecated() uint16 {
	return 0
}

func (*OrderFlags) isWorkingSinceVersion() uint16 {
	return 0
}

func (o *OrderFlags) isWorkingInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.isWorkingSinceVersion()
}

func (*OrderFlags) isWorkingDeprecated() uint16 {
	return 0
}

func (*OrderFlags) cancelOnDisconnectSinceVersion() uint16 {
	return 0
}

func (o *OrderFlags) cancelOnDisconnectInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.cancelOnDisconnectSinceVersion()
}

func (*OrderFlags) cancelOnDisconnectDeprecated() uint16 {
	return 0
}

func (*OrderFlags) postOnlySinceVersion() uint16 {
	return 0
}

func (o *OrderFlags) postOnlyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.postOnlySinceVersion()
}

func (*OrderFlags) postOnlyDeprecated() uint16 {
	return 0
}

func (*OrderFlags) isAggressorSinceVersion() uint16 {
	return 0
}

func (o *OrderFlags) isAggressorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.isAggressorSinceVersion()
}

func (*OrderFlags) isAggressorDeprecated() uint16 {
	return 0
}

func (*OrderFlags) isStopTriggerPendingSinceVersion() uint16 {
	return 0
}

func (o *OrderFlags) isStopTriggerPendingInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.isStopTriggerPendingSinceVersion()
}

func (*OrderFlags) isStopTriggerPendingDeprecated() uint16 {
	return 0
}

func (*OrderFlags) isCancelPendingSinceVersion() uint16 {
	return 0
}

func (o *OrderFlags) isCancelPendingInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.isCancelPendingSinceVersion()
}

func (*OrderFlags) isCancelPendingDeprecated() uint16 {
	return 0
}
