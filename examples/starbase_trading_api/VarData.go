// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"math"
)

type VarData struct {
	Length  uint16
	VarData int8
}

func (v *VarData) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint16(_w, v.Length); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, v.VarData); err != nil {
		return err
	}
	return nil
}

func (v *VarData) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !v.LengthInActingVersion(actingVersion) {
		v.Length = v.LengthNullValue()
	} else {
		if err := _m.ReadUint16(_r, &v.Length); err != nil {
			return err
		}
	}
	if !v.VarDataInActingVersion(actingVersion) {
		v.VarData = v.VarDataNullValue()
	} else {
		if err := _m.ReadInt8(_r, &v.VarData); err != nil {
			return err
		}
	}
	return nil
}

func (v *VarData) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if v.LengthInActingVersion(actingVersion) {
		if v.Length < v.LengthMinValue() || v.Length > v.LengthMaxValue() {
			return fmt.Errorf("Range check failed on v.Length (%v < %v > %v)", v.LengthMinValue(), v.Length, v.LengthMaxValue())
		}
	}
	if v.VarDataInActingVersion(actingVersion) {
		if v.VarData < v.VarDataMinValue() || v.VarData > v.VarDataMaxValue() {
			return fmt.Errorf("Range check failed on v.VarData (%v < %v > %v)", v.VarDataMinValue(), v.VarData, v.VarDataMaxValue())
		}
	}
	return nil
}

func VarDataInit(v *VarData) {
	return
}

func (*VarData) EncodedLength() int64 {
	return -1
}

func (*VarData) LengthMinValue() uint16 {
	return 0
}

func (*VarData) LengthMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*VarData) LengthNullValue() uint16 {
	return math.MaxUint16
}

func (*VarData) LengthSinceVersion() uint16 {
	return 0
}

func (v *VarData) LengthInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.LengthSinceVersion()
}

func (*VarData) LengthDeprecated() uint16 {
	return 0
}

func (*VarData) VarDataMinValue() int8 {
	return math.MinInt8 + 1
}

func (*VarData) VarDataMaxValue() int8 {
	return math.MaxInt8
}

func (*VarData) VarDataNullValue() int8 {
	return math.MinInt8
}

func (*VarData) VarDataSinceVersion() uint16 {
	return 0
}

func (v *VarData) VarDataInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.VarDataSinceVersion()
}

func (*VarData) VarDataDeprecated() uint16 {
	return 0
}
