// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"reflect"
)

type TimeInForceEnum int8
type TimeInForceValues struct {
	GOOD_TIL_CANCELED   TimeInForceEnum
	GOOD_TIL_TIME       TimeInForceEnum
	FILL_OR_KILL        TimeInForceEnum
	IMMEDIATE_OR_CANCEL TimeInForceEnum
	NullValue           TimeInForceEnum
}

var TimeInForce = TimeInForceValues{1, 2, 3, 4, -128}

func (t TimeInForceEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt8(_w, int8(t)); err != nil {
		return err
	}
	return nil
}

func (t *TimeInForceEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadInt8(_r, (*int8)(t)); err != nil {
		return err
	}
	return nil
}

func (t TimeInForceEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(TimeInForce)
	for idx := 0; idx < value.NumField(); idx++ {
		if t == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on TimeInForce, unknown enumeration value %d", t)
}

func (*TimeInForceEnum) EncodedLength() int64 {
	return 1
}

func (*TimeInForceEnum) GOOD_TIL_CANCELEDSinceVersion() uint16 {
	return 0
}

func (t *TimeInForceEnum) GOOD_TIL_CANCELEDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.GOOD_TIL_CANCELEDSinceVersion()
}

func (*TimeInForceEnum) GOOD_TIL_CANCELEDDeprecated() uint16 {
	return 0
}

func (*TimeInForceEnum) GOOD_TIL_TIMESinceVersion() uint16 {
	return 0
}

func (t *TimeInForceEnum) GOOD_TIL_TIMEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.GOOD_TIL_TIMESinceVersion()
}

func (*TimeInForceEnum) GOOD_TIL_TIMEDeprecated() uint16 {
	return 0
}

func (*TimeInForceEnum) FILL_OR_KILLSinceVersion() uint16 {
	return 0
}

func (t *TimeInForceEnum) FILL_OR_KILLInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.FILL_OR_KILLSinceVersion()
}

func (*TimeInForceEnum) FILL_OR_KILLDeprecated() uint16 {
	return 0
}

func (*TimeInForceEnum) IMMEDIATE_OR_CANCELSinceVersion() uint16 {
	return 0
}

func (t *TimeInForceEnum) IMMEDIATE_OR_CANCELInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.IMMEDIATE_OR_CANCELSinceVersion()
}

func (*TimeInForceEnum) IMMEDIATE_OR_CANCELDeprecated() uint16 {
	return 0
}
