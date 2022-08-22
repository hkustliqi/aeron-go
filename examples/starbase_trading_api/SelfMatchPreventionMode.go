// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"reflect"
)

type SelfMatchPreventionModeEnum int8
type SelfMatchPreventionModeValues struct {
	NONE         SelfMatchPreventionModeEnum
	CANCEL_BOTH  SelfMatchPreventionModeEnum
	CANCEL_OLDER SelfMatchPreventionModeEnum
	CANCEL_NEWER SelfMatchPreventionModeEnum
	NullValue    SelfMatchPreventionModeEnum
}

var SelfMatchPreventionMode = SelfMatchPreventionModeValues{1, 3, 4, 5, -128}

func (s SelfMatchPreventionModeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt8(_w, int8(s)); err != nil {
		return err
	}
	return nil
}

func (s *SelfMatchPreventionModeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadInt8(_r, (*int8)(s)); err != nil {
		return err
	}
	return nil
}

func (s SelfMatchPreventionModeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(SelfMatchPreventionMode)
	for idx := 0; idx < value.NumField(); idx++ {
		if s == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on SelfMatchPreventionMode, unknown enumeration value %d", s)
}

func (*SelfMatchPreventionModeEnum) EncodedLength() int64 {
	return 1
}

func (*SelfMatchPreventionModeEnum) NONESinceVersion() uint16 {
	return 0
}

func (s *SelfMatchPreventionModeEnum) NONEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.NONESinceVersion()
}

func (*SelfMatchPreventionModeEnum) NONEDeprecated() uint16 {
	return 0
}

func (*SelfMatchPreventionModeEnum) CANCEL_BOTHSinceVersion() uint16 {
	return 0
}

func (s *SelfMatchPreventionModeEnum) CANCEL_BOTHInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.CANCEL_BOTHSinceVersion()
}

func (*SelfMatchPreventionModeEnum) CANCEL_BOTHDeprecated() uint16 {
	return 0
}

func (*SelfMatchPreventionModeEnum) CANCEL_OLDERSinceVersion() uint16 {
	return 0
}

func (s *SelfMatchPreventionModeEnum) CANCEL_OLDERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.CANCEL_OLDERSinceVersion()
}

func (*SelfMatchPreventionModeEnum) CANCEL_OLDERDeprecated() uint16 {
	return 0
}

func (*SelfMatchPreventionModeEnum) CANCEL_NEWERSinceVersion() uint16 {
	return 0
}

func (s *SelfMatchPreventionModeEnum) CANCEL_NEWERInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.CANCEL_NEWERSinceVersion()
}

func (*SelfMatchPreventionModeEnum) CANCEL_NEWERDeprecated() uint16 {
	return 0
}
