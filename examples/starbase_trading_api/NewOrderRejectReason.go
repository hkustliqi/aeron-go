// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"reflect"
)

type NewOrderRejectReasonEnum int16
type NewOrderRejectReasonValues struct {
	MISSING_FIELD             NewOrderRejectReasonEnum
	INVALID_FIELD             NewOrderRejectReasonEnum
	EXTRA_FIELD               NewOrderRejectReasonEnum
	FORBIDDEN                 NewOrderRejectReasonEnum
	UNKNOWN_INSTRUMENT        NewOrderRejectReasonEnum
	POST_ONLY                 NewOrderRejectReasonEnum
	INSUFFICIENT_FUNDS        NewOrderRejectReasonEnum
	USER_PERMISSION_ERROR     NewOrderRejectReasonEnum
	INVALID_USER              NewOrderRejectReasonEnum
	DUPLICATE_CLIENT_ORDER_ID NewOrderRejectReasonEnum
	INVALID_QUANTITY          NewOrderRejectReasonEnum
	INVALID_PRICE             NewOrderRejectReasonEnum
	STOP_PRICE_TYPE           NewOrderRejectReasonEnum
	STOP_POST_ONLY            NewOrderRejectReasonEnum
	STOP_TIF                  NewOrderRejectReasonEnum
	STOP_MARKET               NewOrderRejectReasonEnum
	LIMIT_EXPIRE_TIME         NewOrderRejectReasonEnum
	LIMIT_EXPIRE_GTT          NewOrderRejectReasonEnum
	MARKET_POST_ONLY          NewOrderRejectReasonEnum
	INSUFFICIENT_LIQUIDITY    NewOrderRejectReasonEnum
	CANCEL_ONLY               NewOrderRejectReasonEnum
	OFFLINE                   NewOrderRejectReasonEnum
	NOT_STARTED               NewOrderRejectReasonEnum
	TRADING_DISABLED          NewOrderRejectReasonEnum
	CROSSED_BOOK              NewOrderRejectReasonEnum
	NullValue                 NewOrderRejectReasonEnum
}

var NewOrderRejectReason = NewOrderRejectReasonValues{1, 2, 3, 4, 5, 7, 14, 15, 16, 17, 18, 19, 23, 24, 25, 26, 27, 28, 29, 35, 37, 38, 39, 41, 49, -32768}

func (n NewOrderRejectReasonEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt16(_w, int16(n)); err != nil {
		return err
	}
	return nil
}

func (n *NewOrderRejectReasonEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadInt16(_r, (*int16)(n)); err != nil {
		return err
	}
	return nil
}

func (n NewOrderRejectReasonEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(NewOrderRejectReason)
	for idx := 0; idx < value.NumField(); idx++ {
		if n == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on NewOrderRejectReason, unknown enumeration value %d", n)
}

func (*NewOrderRejectReasonEnum) EncodedLength() int64 {
	return 2
}

func (*NewOrderRejectReasonEnum) MISSING_FIELDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderRejectReasonEnum) MISSING_FIELDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.MISSING_FIELDSinceVersion()
}

func (*NewOrderRejectReasonEnum) MISSING_FIELDDeprecated() uint16 {
	return 0
}

func (*NewOrderRejectReasonEnum) INVALID_FIELDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderRejectReasonEnum) INVALID_FIELDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.INVALID_FIELDSinceVersion()
}

func (*NewOrderRejectReasonEnum) INVALID_FIELDDeprecated() uint16 {
	return 0
}

func (*NewOrderRejectReasonEnum) EXTRA_FIELDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderRejectReasonEnum) EXTRA_FIELDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.EXTRA_FIELDSinceVersion()
}

func (*NewOrderRejectReasonEnum) EXTRA_FIELDDeprecated() uint16 {
	return 0
}

func (*NewOrderRejectReasonEnum) FORBIDDENSinceVersion() uint16 {
	return 0
}

func (n *NewOrderRejectReasonEnum) FORBIDDENInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.FORBIDDENSinceVersion()
}

func (*NewOrderRejectReasonEnum) FORBIDDENDeprecated() uint16 {
	return 0
}

func (*NewOrderRejectReasonEnum) UNKNOWN_INSTRUMENTSinceVersion() uint16 {
	return 0
}

func (n *NewOrderRejectReasonEnum) UNKNOWN_INSTRUMENTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.UNKNOWN_INSTRUMENTSinceVersion()
}

func (*NewOrderRejectReasonEnum) UNKNOWN_INSTRUMENTDeprecated() uint16 {
	return 0
}

func (*NewOrderRejectReasonEnum) POST_ONLYSinceVersion() uint16 {
	return 0
}

func (n *NewOrderRejectReasonEnum) POST_ONLYInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.POST_ONLYSinceVersion()
}

func (*NewOrderRejectReasonEnum) POST_ONLYDeprecated() uint16 {
	return 0
}

func (*NewOrderRejectReasonEnum) INSUFFICIENT_FUNDSSinceVersion() uint16 {
	return 0
}

func (n *NewOrderRejectReasonEnum) INSUFFICIENT_FUNDSInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.INSUFFICIENT_FUNDSSinceVersion()
}

func (*NewOrderRejectReasonEnum) INSUFFICIENT_FUNDSDeprecated() uint16 {
	return 0
}

func (*NewOrderRejectReasonEnum) USER_PERMISSION_ERRORSinceVersion() uint16 {
	return 0
}

func (n *NewOrderRejectReasonEnum) USER_PERMISSION_ERRORInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.USER_PERMISSION_ERRORSinceVersion()
}

func (*NewOrderRejectReasonEnum) USER_PERMISSION_ERRORDeprecated() uint16 {
	return 0
}

func (*NewOrderRejectReasonEnum) INVALID_USERSinceVersion() uint16 {
	return 0
}

func (n *NewOrderRejectReasonEnum) INVALID_USERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.INVALID_USERSinceVersion()
}

func (*NewOrderRejectReasonEnum) INVALID_USERDeprecated() uint16 {
	return 0
}

func (*NewOrderRejectReasonEnum) DUPLICATE_CLIENT_ORDER_IDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderRejectReasonEnum) DUPLICATE_CLIENT_ORDER_IDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.DUPLICATE_CLIENT_ORDER_IDSinceVersion()
}

func (*NewOrderRejectReasonEnum) DUPLICATE_CLIENT_ORDER_IDDeprecated() uint16 {
	return 0
}

func (*NewOrderRejectReasonEnum) INVALID_QUANTITYSinceVersion() uint16 {
	return 0
}

func (n *NewOrderRejectReasonEnum) INVALID_QUANTITYInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.INVALID_QUANTITYSinceVersion()
}

func (*NewOrderRejectReasonEnum) INVALID_QUANTITYDeprecated() uint16 {
	return 0
}

func (*NewOrderRejectReasonEnum) INVALID_PRICESinceVersion() uint16 {
	return 0
}

func (n *NewOrderRejectReasonEnum) INVALID_PRICEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.INVALID_PRICESinceVersion()
}

func (*NewOrderRejectReasonEnum) INVALID_PRICEDeprecated() uint16 {
	return 0
}

func (*NewOrderRejectReasonEnum) STOP_PRICE_TYPESinceVersion() uint16 {
	return 0
}

func (n *NewOrderRejectReasonEnum) STOP_PRICE_TYPEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.STOP_PRICE_TYPESinceVersion()
}

func (*NewOrderRejectReasonEnum) STOP_PRICE_TYPEDeprecated() uint16 {
	return 0
}

func (*NewOrderRejectReasonEnum) STOP_POST_ONLYSinceVersion() uint16 {
	return 0
}

func (n *NewOrderRejectReasonEnum) STOP_POST_ONLYInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.STOP_POST_ONLYSinceVersion()
}

func (*NewOrderRejectReasonEnum) STOP_POST_ONLYDeprecated() uint16 {
	return 0
}

func (*NewOrderRejectReasonEnum) STOP_TIFSinceVersion() uint16 {
	return 0
}

func (n *NewOrderRejectReasonEnum) STOP_TIFInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.STOP_TIFSinceVersion()
}

func (*NewOrderRejectReasonEnum) STOP_TIFDeprecated() uint16 {
	return 0
}

func (*NewOrderRejectReasonEnum) STOP_MARKETSinceVersion() uint16 {
	return 0
}

func (n *NewOrderRejectReasonEnum) STOP_MARKETInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.STOP_MARKETSinceVersion()
}

func (*NewOrderRejectReasonEnum) STOP_MARKETDeprecated() uint16 {
	return 0
}

func (*NewOrderRejectReasonEnum) LIMIT_EXPIRE_TIMESinceVersion() uint16 {
	return 0
}

func (n *NewOrderRejectReasonEnum) LIMIT_EXPIRE_TIMEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.LIMIT_EXPIRE_TIMESinceVersion()
}

func (*NewOrderRejectReasonEnum) LIMIT_EXPIRE_TIMEDeprecated() uint16 {
	return 0
}

func (*NewOrderRejectReasonEnum) LIMIT_EXPIRE_GTTSinceVersion() uint16 {
	return 0
}

func (n *NewOrderRejectReasonEnum) LIMIT_EXPIRE_GTTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.LIMIT_EXPIRE_GTTSinceVersion()
}

func (*NewOrderRejectReasonEnum) LIMIT_EXPIRE_GTTDeprecated() uint16 {
	return 0
}

func (*NewOrderRejectReasonEnum) MARKET_POST_ONLYSinceVersion() uint16 {
	return 0
}

func (n *NewOrderRejectReasonEnum) MARKET_POST_ONLYInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.MARKET_POST_ONLYSinceVersion()
}

func (*NewOrderRejectReasonEnum) MARKET_POST_ONLYDeprecated() uint16 {
	return 0
}

func (*NewOrderRejectReasonEnum) INSUFFICIENT_LIQUIDITYSinceVersion() uint16 {
	return 0
}

func (n *NewOrderRejectReasonEnum) INSUFFICIENT_LIQUIDITYInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.INSUFFICIENT_LIQUIDITYSinceVersion()
}

func (*NewOrderRejectReasonEnum) INSUFFICIENT_LIQUIDITYDeprecated() uint16 {
	return 0
}

func (*NewOrderRejectReasonEnum) CANCEL_ONLYSinceVersion() uint16 {
	return 0
}

func (n *NewOrderRejectReasonEnum) CANCEL_ONLYInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.CANCEL_ONLYSinceVersion()
}

func (*NewOrderRejectReasonEnum) CANCEL_ONLYDeprecated() uint16 {
	return 0
}

func (*NewOrderRejectReasonEnum) OFFLINESinceVersion() uint16 {
	return 0
}

func (n *NewOrderRejectReasonEnum) OFFLINEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.OFFLINESinceVersion()
}

func (*NewOrderRejectReasonEnum) OFFLINEDeprecated() uint16 {
	return 0
}

func (*NewOrderRejectReasonEnum) NOT_STARTEDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderRejectReasonEnum) NOT_STARTEDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.NOT_STARTEDSinceVersion()
}

func (*NewOrderRejectReasonEnum) NOT_STARTEDDeprecated() uint16 {
	return 0
}

func (*NewOrderRejectReasonEnum) TRADING_DISABLEDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderRejectReasonEnum) TRADING_DISABLEDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.TRADING_DISABLEDSinceVersion()
}

func (*NewOrderRejectReasonEnum) TRADING_DISABLEDDeprecated() uint16 {
	return 0
}

func (*NewOrderRejectReasonEnum) CROSSED_BOOKSinceVersion() uint16 {
	return 0
}

func (n *NewOrderRejectReasonEnum) CROSSED_BOOKInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.CROSSED_BOOKSinceVersion()
}

func (*NewOrderRejectReasonEnum) CROSSED_BOOKDeprecated() uint16 {
	return 0
}
