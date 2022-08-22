// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"reflect"
)

type CancelReasonEnum int8
type CancelReasonValues struct {
	UNKNOWN                CancelReasonEnum
	SELF_TRADE_PREVENTION  CancelReasonEnum
	TIME_IN_FORCE          CancelReasonEnum
	USER                   CancelReasonEnum
	ADMIN                  CancelReasonEnum
	PRICE_BOUND            CancelReasonEnum
	INSUFFICIENT_FUNDS     CancelReasonEnum
	INSUFFICIENT_LIQUIDITY CancelReasonEnum
	NullValue              CancelReasonEnum
}

var CancelReason = CancelReasonValues{0, 1, 2, 3, 4, 5, 6, 7, -128}

func (c CancelReasonEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt8(_w, int8(c)); err != nil {
		return err
	}
	return nil
}

func (c *CancelReasonEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadInt8(_r, (*int8)(c)); err != nil {
		return err
	}
	return nil
}

func (c CancelReasonEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(CancelReason)
	for idx := 0; idx < value.NumField(); idx++ {
		if c == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on CancelReason, unknown enumeration value %d", c)
}

func (*CancelReasonEnum) EncodedLength() int64 {
	return 1
}

func (*CancelReasonEnum) UNKNOWNSinceVersion() uint16 {
	return 0
}

func (c *CancelReasonEnum) UNKNOWNInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.UNKNOWNSinceVersion()
}

func (*CancelReasonEnum) UNKNOWNDeprecated() uint16 {
	return 0
}

func (*CancelReasonEnum) SELF_TRADE_PREVENTIONSinceVersion() uint16 {
	return 0
}

func (c *CancelReasonEnum) SELF_TRADE_PREVENTIONInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.SELF_TRADE_PREVENTIONSinceVersion()
}

func (*CancelReasonEnum) SELF_TRADE_PREVENTIONDeprecated() uint16 {
	return 0
}

func (*CancelReasonEnum) TIME_IN_FORCESinceVersion() uint16 {
	return 0
}

func (c *CancelReasonEnum) TIME_IN_FORCEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.TIME_IN_FORCESinceVersion()
}

func (*CancelReasonEnum) TIME_IN_FORCEDeprecated() uint16 {
	return 0
}

func (*CancelReasonEnum) USERSinceVersion() uint16 {
	return 0
}

func (c *CancelReasonEnum) USERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.USERSinceVersion()
}

func (*CancelReasonEnum) USERDeprecated() uint16 {
	return 0
}

func (*CancelReasonEnum) ADMINSinceVersion() uint16 {
	return 0
}

func (c *CancelReasonEnum) ADMINInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.ADMINSinceVersion()
}

func (*CancelReasonEnum) ADMINDeprecated() uint16 {
	return 0
}

func (*CancelReasonEnum) PRICE_BOUNDSinceVersion() uint16 {
	return 0
}

func (c *CancelReasonEnum) PRICE_BOUNDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.PRICE_BOUNDSinceVersion()
}

func (*CancelReasonEnum) PRICE_BOUNDDeprecated() uint16 {
	return 0
}

func (*CancelReasonEnum) INSUFFICIENT_FUNDSSinceVersion() uint16 {
	return 0
}

func (c *CancelReasonEnum) INSUFFICIENT_FUNDSInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.INSUFFICIENT_FUNDSSinceVersion()
}

func (*CancelReasonEnum) INSUFFICIENT_FUNDSDeprecated() uint16 {
	return 0
}

func (*CancelReasonEnum) INSUFFICIENT_LIQUIDITYSinceVersion() uint16 {
	return 0
}

func (c *CancelReasonEnum) INSUFFICIENT_LIQUIDITYInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.INSUFFICIENT_LIQUIDITYSinceVersion()
}

func (*CancelReasonEnum) INSUFFICIENT_LIQUIDITYDeprecated() uint16 {
	return 0
}
