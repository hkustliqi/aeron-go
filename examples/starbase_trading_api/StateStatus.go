// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"reflect"
)

type StateStatusEnum int8
type StateStatusValues struct {
	ACTIVE    StateStatusEnum
	DELETED   StateStatusEnum
	DISABLED  StateStatusEnum
	NullValue StateStatusEnum
}

var StateStatus = StateStatusValues{0, -1, 1, -128}

func (s StateStatusEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt8(_w, int8(s)); err != nil {
		return err
	}
	return nil
}

func (s *StateStatusEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadInt8(_r, (*int8)(s)); err != nil {
		return err
	}
	return nil
}

func (s StateStatusEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(StateStatus)
	for idx := 0; idx < value.NumField(); idx++ {
		if s == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on StateStatus, unknown enumeration value %d", s)
}

func (*StateStatusEnum) EncodedLength() int64 {
	return 1
}

func (*StateStatusEnum) ACTIVESinceVersion() uint16 {
	return 0
}

func (s *StateStatusEnum) ACTIVEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.ACTIVESinceVersion()
}

func (*StateStatusEnum) ACTIVEDeprecated() uint16 {
	return 0
}

func (*StateStatusEnum) DELETEDSinceVersion() uint16 {
	return 0
}

func (s *StateStatusEnum) DELETEDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.DELETEDSinceVersion()
}

func (*StateStatusEnum) DELETEDDeprecated() uint16 {
	return 0
}

func (*StateStatusEnum) DISABLEDSinceVersion() uint16 {
	return 0
}

func (s *StateStatusEnum) DISABLEDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.DISABLEDSinceVersion()
}

func (*StateStatusEnum) DISABLEDDeprecated() uint16 {
	return 0
}
