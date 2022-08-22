// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"reflect"
)

type InstrumentTypeEnum int8
type InstrumentTypeValues struct {
	SPOT             InstrumentTypeEnum
	PERPETUAL_FUTURE InstrumentTypeEnum
	NullValue        InstrumentTypeEnum
}

var InstrumentType = InstrumentTypeValues{0, 1, -128}

func (i InstrumentTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt8(_w, int8(i)); err != nil {
		return err
	}
	return nil
}

func (i *InstrumentTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadInt8(_r, (*int8)(i)); err != nil {
		return err
	}
	return nil
}

func (i InstrumentTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(InstrumentType)
	for idx := 0; idx < value.NumField(); idx++ {
		if i == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on InstrumentType, unknown enumeration value %d", i)
}

func (*InstrumentTypeEnum) EncodedLength() int64 {
	return 1
}

func (*InstrumentTypeEnum) SPOTSinceVersion() uint16 {
	return 0
}

func (i *InstrumentTypeEnum) SPOTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.SPOTSinceVersion()
}

func (*InstrumentTypeEnum) SPOTDeprecated() uint16 {
	return 0
}

func (*InstrumentTypeEnum) PERPETUAL_FUTURESinceVersion() uint16 {
	return 0
}

func (i *InstrumentTypeEnum) PERPETUAL_FUTUREInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.PERPETUAL_FUTURESinceVersion()
}

func (*InstrumentTypeEnum) PERPETUAL_FUTUREDeprecated() uint16 {
	return 0
}
