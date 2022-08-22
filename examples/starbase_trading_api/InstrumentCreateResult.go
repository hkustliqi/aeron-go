// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"reflect"
)

type InstrumentCreateResultEnum int8
type InstrumentCreateResultValues struct {
	SUCCESS                            InstrumentCreateResultEnum
	REJECT_ALREADY_EXISTS              InstrumentCreateResultEnum
	REJECT_INVALID_INCREMENT           InstrumentCreateResultEnum
	REJECT_INVALID_MIN_QUANTITY        InstrumentCreateResultEnum
	REJECT_INVALID_SETTLEMENT_INTERVAL InstrumentCreateResultEnum
	REJECT_INVALID_UNDERLYING_SPOT     InstrumentCreateResultEnum
	REJECT_OTHER                       InstrumentCreateResultEnum
	NullValue                          InstrumentCreateResultEnum
}

var InstrumentCreateResult = InstrumentCreateResultValues{0, 1, 2, 3, 4, 5, 127, -128}

func (i InstrumentCreateResultEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt8(_w, int8(i)); err != nil {
		return err
	}
	return nil
}

func (i *InstrumentCreateResultEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadInt8(_r, (*int8)(i)); err != nil {
		return err
	}
	return nil
}

func (i InstrumentCreateResultEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(InstrumentCreateResult)
	for idx := 0; idx < value.NumField(); idx++ {
		if i == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on InstrumentCreateResult, unknown enumeration value %d", i)
}

func (*InstrumentCreateResultEnum) EncodedLength() int64 {
	return 1
}

func (*InstrumentCreateResultEnum) SUCCESSSinceVersion() uint16 {
	return 0
}

func (i *InstrumentCreateResultEnum) SUCCESSInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.SUCCESSSinceVersion()
}

func (*InstrumentCreateResultEnum) SUCCESSDeprecated() uint16 {
	return 0
}

func (*InstrumentCreateResultEnum) REJECT_ALREADY_EXISTSSinceVersion() uint16 {
	return 0
}

func (i *InstrumentCreateResultEnum) REJECT_ALREADY_EXISTSInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.REJECT_ALREADY_EXISTSSinceVersion()
}

func (*InstrumentCreateResultEnum) REJECT_ALREADY_EXISTSDeprecated() uint16 {
	return 0
}

func (*InstrumentCreateResultEnum) REJECT_INVALID_INCREMENTSinceVersion() uint16 {
	return 0
}

func (i *InstrumentCreateResultEnum) REJECT_INVALID_INCREMENTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.REJECT_INVALID_INCREMENTSinceVersion()
}

func (*InstrumentCreateResultEnum) REJECT_INVALID_INCREMENTDeprecated() uint16 {
	return 0
}

func (*InstrumentCreateResultEnum) REJECT_INVALID_MIN_QUANTITYSinceVersion() uint16 {
	return 0
}

func (i *InstrumentCreateResultEnum) REJECT_INVALID_MIN_QUANTITYInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.REJECT_INVALID_MIN_QUANTITYSinceVersion()
}

func (*InstrumentCreateResultEnum) REJECT_INVALID_MIN_QUANTITYDeprecated() uint16 {
	return 0
}

func (*InstrumentCreateResultEnum) REJECT_INVALID_SETTLEMENT_INTERVALSinceVersion() uint16 {
	return 0
}

func (i *InstrumentCreateResultEnum) REJECT_INVALID_SETTLEMENT_INTERVALInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.REJECT_INVALID_SETTLEMENT_INTERVALSinceVersion()
}

func (*InstrumentCreateResultEnum) REJECT_INVALID_SETTLEMENT_INTERVALDeprecated() uint16 {
	return 0
}

func (*InstrumentCreateResultEnum) REJECT_INVALID_UNDERLYING_SPOTSinceVersion() uint16 {
	return 0
}

func (i *InstrumentCreateResultEnum) REJECT_INVALID_UNDERLYING_SPOTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.REJECT_INVALID_UNDERLYING_SPOTSinceVersion()
}

func (*InstrumentCreateResultEnum) REJECT_INVALID_UNDERLYING_SPOTDeprecated() uint16 {
	return 0
}

func (*InstrumentCreateResultEnum) REJECT_OTHERSinceVersion() uint16 {
	return 0
}

func (i *InstrumentCreateResultEnum) REJECT_OTHERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.REJECT_OTHERSinceVersion()
}

func (*InstrumentCreateResultEnum) REJECT_OTHERDeprecated() uint16 {
	return 0
}
