// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"reflect"
)

type SourceServiceEnum int8
type SourceServiceValues struct {
	Oms       SourceServiceEnum
	Engine    SourceServiceEnum
	NullValue SourceServiceEnum
}

var SourceService = SourceServiceValues{0, 1, -128}

func (s SourceServiceEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt8(_w, int8(s)); err != nil {
		return err
	}
	return nil
}

func (s *SourceServiceEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadInt8(_r, (*int8)(s)); err != nil {
		return err
	}
	return nil
}

func (s SourceServiceEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(SourceService)
	for idx := 0; idx < value.NumField(); idx++ {
		if s == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on SourceService, unknown enumeration value %d", s)
}

func (*SourceServiceEnum) EncodedLength() int64 {
	return 1
}

func (*SourceServiceEnum) OmsSinceVersion() uint16 {
	return 0
}

func (s *SourceServiceEnum) OmsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.OmsSinceVersion()
}

func (*SourceServiceEnum) OmsDeprecated() uint16 {
	return 0
}

func (*SourceServiceEnum) EngineSinceVersion() uint16 {
	return 0
}

func (s *SourceServiceEnum) EngineInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.EngineSinceVersion()
}

func (*SourceServiceEnum) EngineDeprecated() uint16 {
	return 0
}
