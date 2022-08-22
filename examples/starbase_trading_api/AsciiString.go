// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"math"
)

type AsciiString struct {
	Length  uint16
	VarData byte
}

func (a *AsciiString) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint16(_w, a.Length); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, a.VarData); err != nil {
		return err
	}
	return nil
}

func (a *AsciiString) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !a.LengthInActingVersion(actingVersion) {
		a.Length = a.LengthNullValue()
	} else {
		if err := _m.ReadUint16(_r, &a.Length); err != nil {
			return err
		}
	}
	if !a.VarDataInActingVersion(actingVersion) {
		a.VarData = a.VarDataNullValue()
	} else {
		if err := _m.ReadUint8(_r, &a.VarData); err != nil {
			return err
		}
	}
	return nil
}

func (a *AsciiString) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if a.LengthInActingVersion(actingVersion) {
		if a.Length < a.LengthMinValue() || a.Length > a.LengthMaxValue() {
			return fmt.Errorf("Range check failed on a.Length (%v < %v > %v)", a.LengthMinValue(), a.Length, a.LengthMaxValue())
		}
	}
	if a.VarDataInActingVersion(actingVersion) {
		if a.VarData < a.VarDataMinValue() || a.VarData > a.VarDataMaxValue() {
			return fmt.Errorf("Range check failed on a.VarData (%v < %v > %v)", a.VarDataMinValue(), a.VarData, a.VarDataMaxValue())
		}
	}
	return nil
}

func AsciiStringInit(a *AsciiString) {
	return
}

func (*AsciiString) EncodedLength() int64 {
	return -1
}

func (*AsciiString) LengthMinValue() uint16 {
	return 0
}

func (*AsciiString) LengthMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*AsciiString) LengthNullValue() uint16 {
	return math.MaxUint16
}

func (*AsciiString) LengthSinceVersion() uint16 {
	return 0
}

func (a *AsciiString) LengthInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.LengthSinceVersion()
}

func (*AsciiString) LengthDeprecated() uint16 {
	return 0
}

func (*AsciiString) VarDataMinValue() byte {
	return byte(32)
}

func (*AsciiString) VarDataMaxValue() byte {
	return byte(126)
}

func (*AsciiString) VarDataNullValue() byte {
	return 0
}

func (*AsciiString) VarDataSinceVersion() uint16 {
	return 0
}

func (a *AsciiString) VarDataInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.VarDataSinceVersion()
}

func (*AsciiString) VarDataDeprecated() uint16 {
	return 0
}
