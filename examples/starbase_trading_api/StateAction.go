// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"reflect"
)

type StateActionEnum uint8
type StateActionValues struct {
	CREATE    StateActionEnum
	UPDATE    StateActionEnum
	SNAPSHOT  StateActionEnum
	DELETE    StateActionEnum
	NullValue StateActionEnum
}

var StateAction = StateActionValues{0, 1, 2, 3, 255}

func (s StateActionEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(s)); err != nil {
		return err
	}
	return nil
}

func (s *StateActionEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(s)); err != nil {
		return err
	}
	return nil
}

func (s StateActionEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(StateAction)
	for idx := 0; idx < value.NumField(); idx++ {
		if s == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on StateAction, unknown enumeration value %d", s)
}

func (*StateActionEnum) EncodedLength() int64 {
	return 1
}

func (*StateActionEnum) CREATESinceVersion() uint16 {
	return 0
}

func (s *StateActionEnum) CREATEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.CREATESinceVersion()
}

func (*StateActionEnum) CREATEDeprecated() uint16 {
	return 0
}

func (*StateActionEnum) UPDATESinceVersion() uint16 {
	return 0
}

func (s *StateActionEnum) UPDATEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.UPDATESinceVersion()
}

func (*StateActionEnum) UPDATEDeprecated() uint16 {
	return 0
}

func (*StateActionEnum) SNAPSHOTSinceVersion() uint16 {
	return 0
}

func (s *StateActionEnum) SNAPSHOTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SNAPSHOTSinceVersion()
}

func (*StateActionEnum) SNAPSHOTDeprecated() uint16 {
	return 0
}

func (*StateActionEnum) DELETESinceVersion() uint16 {
	return 0
}

func (s *StateActionEnum) DELETEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.DELETESinceVersion()
}

func (*StateActionEnum) DELETEDeprecated() uint16 {
	return 0
}
