// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"reflect"
)

type CancelOrderRejectReasonEnum int8
type CancelOrderRejectReasonValues struct {
	ERROR                       CancelOrderRejectReasonEnum
	UNKNOWN_ORDER               CancelOrderRejectReasonEnum
	ORDER_FILLED                CancelOrderRejectReasonEnum
	NOT_ALLOWED_BY_MARKET_STATE CancelOrderRejectReasonEnum
	INVALID_USER                CancelOrderRejectReasonEnum
	USER_PERMISSION_ERROR       CancelOrderRejectReasonEnum
	INVALID_PORTFOLIO           CancelOrderRejectReasonEnum
	CANCEL_PENDING              CancelOrderRejectReasonEnum
	NullValue                   CancelOrderRejectReasonEnum
}

var CancelOrderRejectReason = CancelOrderRejectReasonValues{0, 1, 2, 3, 4, 5, 6, 7, -128}

func (c CancelOrderRejectReasonEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt8(_w, int8(c)); err != nil {
		return err
	}
	return nil
}

func (c *CancelOrderRejectReasonEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadInt8(_r, (*int8)(c)); err != nil {
		return err
	}
	return nil
}

func (c CancelOrderRejectReasonEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(CancelOrderRejectReason)
	for idx := 0; idx < value.NumField(); idx++ {
		if c == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on CancelOrderRejectReason, unknown enumeration value %d", c)
}

func (*CancelOrderRejectReasonEnum) EncodedLength() int64 {
	return 1
}

func (*CancelOrderRejectReasonEnum) ERRORSinceVersion() uint16 {
	return 0
}

func (c *CancelOrderRejectReasonEnum) ERRORInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.ERRORSinceVersion()
}

func (*CancelOrderRejectReasonEnum) ERRORDeprecated() uint16 {
	return 0
}

func (*CancelOrderRejectReasonEnum) UNKNOWN_ORDERSinceVersion() uint16 {
	return 0
}

func (c *CancelOrderRejectReasonEnum) UNKNOWN_ORDERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.UNKNOWN_ORDERSinceVersion()
}

func (*CancelOrderRejectReasonEnum) UNKNOWN_ORDERDeprecated() uint16 {
	return 0
}

func (*CancelOrderRejectReasonEnum) ORDER_FILLEDSinceVersion() uint16 {
	return 0
}

func (c *CancelOrderRejectReasonEnum) ORDER_FILLEDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.ORDER_FILLEDSinceVersion()
}

func (*CancelOrderRejectReasonEnum) ORDER_FILLEDDeprecated() uint16 {
	return 0
}

func (*CancelOrderRejectReasonEnum) NOT_ALLOWED_BY_MARKET_STATESinceVersion() uint16 {
	return 0
}

func (c *CancelOrderRejectReasonEnum) NOT_ALLOWED_BY_MARKET_STATEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.NOT_ALLOWED_BY_MARKET_STATESinceVersion()
}

func (*CancelOrderRejectReasonEnum) NOT_ALLOWED_BY_MARKET_STATEDeprecated() uint16 {
	return 0
}

func (*CancelOrderRejectReasonEnum) INVALID_USERSinceVersion() uint16 {
	return 0
}

func (c *CancelOrderRejectReasonEnum) INVALID_USERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.INVALID_USERSinceVersion()
}

func (*CancelOrderRejectReasonEnum) INVALID_USERDeprecated() uint16 {
	return 0
}

func (*CancelOrderRejectReasonEnum) USER_PERMISSION_ERRORSinceVersion() uint16 {
	return 0
}

func (c *CancelOrderRejectReasonEnum) USER_PERMISSION_ERRORInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.USER_PERMISSION_ERRORSinceVersion()
}

func (*CancelOrderRejectReasonEnum) USER_PERMISSION_ERRORDeprecated() uint16 {
	return 0
}

func (*CancelOrderRejectReasonEnum) INVALID_PORTFOLIOSinceVersion() uint16 {
	return 0
}

func (c *CancelOrderRejectReasonEnum) INVALID_PORTFOLIOInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.INVALID_PORTFOLIOSinceVersion()
}

func (*CancelOrderRejectReasonEnum) INVALID_PORTFOLIODeprecated() uint16 {
	return 0
}

func (*CancelOrderRejectReasonEnum) CANCEL_PENDINGSinceVersion() uint16 {
	return 0
}

func (c *CancelOrderRejectReasonEnum) CANCEL_PENDINGInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.CANCEL_PENDINGSinceVersion()
}

func (*CancelOrderRejectReasonEnum) CANCEL_PENDINGDeprecated() uint16 {
	return 0
}
