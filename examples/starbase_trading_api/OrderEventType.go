// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"reflect"
)

type OrderEventTypeEnum int8
type OrderEventTypeValues struct {
	SUBMITTED          OrderEventTypeEnum
	REJECTED           OrderEventTypeEnum
	CANCELED_BY_USER   OrderEventTypeEnum
	CANCELED_BY_SYSTEM OrderEventTypeEnum
	REPLACED           OrderEventTypeEnum
	FILL               OrderEventTypeEnum
	STOP_TRIGGERED     OrderEventTypeEnum
	SNAPSHOT           OrderEventTypeEnum
	NullValue          OrderEventTypeEnum
}

var OrderEventType = OrderEventTypeValues{1, 2, 3, 4, 5, 6, 7, 8, -128}

func (o OrderEventTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt8(_w, int8(o)); err != nil {
		return err
	}
	return nil
}

func (o *OrderEventTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadInt8(_r, (*int8)(o)); err != nil {
		return err
	}
	return nil
}

func (o OrderEventTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(OrderEventType)
	for idx := 0; idx < value.NumField(); idx++ {
		if o == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on OrderEventType, unknown enumeration value %d", o)
}

func (*OrderEventTypeEnum) EncodedLength() int64 {
	return 1
}

func (*OrderEventTypeEnum) SUBMITTEDSinceVersion() uint16 {
	return 0
}

func (o *OrderEventTypeEnum) SUBMITTEDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SUBMITTEDSinceVersion()
}

func (*OrderEventTypeEnum) SUBMITTEDDeprecated() uint16 {
	return 0
}

func (*OrderEventTypeEnum) REJECTEDSinceVersion() uint16 {
	return 0
}

func (o *OrderEventTypeEnum) REJECTEDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.REJECTEDSinceVersion()
}

func (*OrderEventTypeEnum) REJECTEDDeprecated() uint16 {
	return 0
}

func (*OrderEventTypeEnum) CANCELED_BY_USERSinceVersion() uint16 {
	return 0
}

func (o *OrderEventTypeEnum) CANCELED_BY_USERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CANCELED_BY_USERSinceVersion()
}

func (*OrderEventTypeEnum) CANCELED_BY_USERDeprecated() uint16 {
	return 0
}

func (*OrderEventTypeEnum) CANCELED_BY_SYSTEMSinceVersion() uint16 {
	return 0
}

func (o *OrderEventTypeEnum) CANCELED_BY_SYSTEMInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CANCELED_BY_SYSTEMSinceVersion()
}

func (*OrderEventTypeEnum) CANCELED_BY_SYSTEMDeprecated() uint16 {
	return 0
}

func (*OrderEventTypeEnum) REPLACEDSinceVersion() uint16 {
	return 0
}

func (o *OrderEventTypeEnum) REPLACEDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.REPLACEDSinceVersion()
}

func (*OrderEventTypeEnum) REPLACEDDeprecated() uint16 {
	return 0
}

func (*OrderEventTypeEnum) FILLSinceVersion() uint16 {
	return 0
}

func (o *OrderEventTypeEnum) FILLInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.FILLSinceVersion()
}

func (*OrderEventTypeEnum) FILLDeprecated() uint16 {
	return 0
}

func (*OrderEventTypeEnum) STOP_TRIGGEREDSinceVersion() uint16 {
	return 0
}

func (o *OrderEventTypeEnum) STOP_TRIGGEREDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.STOP_TRIGGEREDSinceVersion()
}

func (*OrderEventTypeEnum) STOP_TRIGGEREDDeprecated() uint16 {
	return 0
}

func (*OrderEventTypeEnum) SNAPSHOTSinceVersion() uint16 {
	return 0
}

func (o *OrderEventTypeEnum) SNAPSHOTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SNAPSHOTSinceVersion()
}

func (*OrderEventTypeEnum) SNAPSHOTDeprecated() uint16 {
	return 0
}
