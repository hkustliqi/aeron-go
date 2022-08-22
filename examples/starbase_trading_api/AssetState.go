// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type AssetState struct {
	AssetUuid              [2]int64
	MarginCollateralWeight int64
}

func (a *AssetState) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := a.RangeCheck(a.SbeSchemaVersion(), a.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	for idx := 0; idx < 2; idx++ {
		if err := _m.WriteInt64(_w, a.AssetUuid[idx]); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, a.MarginCollateralWeight); err != nil {
		return err
	}
	return nil
}

func (a *AssetState) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !a.AssetUuidInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			a.AssetUuid[idx] = a.AssetUuidNullValue()
		}
	} else {
		for idx := 0; idx < 2; idx++ {
			if err := _m.ReadInt64(_r, &a.AssetUuid[idx]); err != nil {
				return err
			}
		}
	}
	if !a.MarginCollateralWeightInActingVersion(actingVersion) {
		a.MarginCollateralWeight = a.MarginCollateralWeightNullValue()
	} else {
		if err := _m.ReadInt64(_r, &a.MarginCollateralWeight); err != nil {
			return err
		}
	}
	if actingVersion > a.SbeSchemaVersion() && blockLength > a.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-a.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := a.RangeCheck(actingVersion, a.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (a *AssetState) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if a.AssetUuidInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if a.AssetUuid[idx] < a.AssetUuidMinValue() || a.AssetUuid[idx] > a.AssetUuidMaxValue() {
				return fmt.Errorf("Range check failed on a.AssetUuid[%d] (%v < %v > %v)", idx, a.AssetUuidMinValue(), a.AssetUuid[idx], a.AssetUuidMaxValue())
			}
		}
	}
	if a.MarginCollateralWeightInActingVersion(actingVersion) {
		if a.MarginCollateralWeight < a.MarginCollateralWeightMinValue() || a.MarginCollateralWeight > a.MarginCollateralWeightMaxValue() {
			return fmt.Errorf("Range check failed on a.MarginCollateralWeight (%v < %v > %v)", a.MarginCollateralWeightMinValue(), a.MarginCollateralWeight, a.MarginCollateralWeightMaxValue())
		}
	}
	return nil
}

func AssetStateInit(a *AssetState) {
	return
}

func (*AssetState) SbeBlockLength() (blockLength uint16) {
	return 24
}

func (*AssetState) SbeTemplateId() (templateId uint16) {
	return 10
}

func (*AssetState) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*AssetState) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*AssetState) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*AssetState) AssetUuidId() uint16 {
	return 1
}

func (*AssetState) AssetUuidSinceVersion() uint16 {
	return 0
}

func (a *AssetState) AssetUuidInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.AssetUuidSinceVersion()
}

func (*AssetState) AssetUuidDeprecated() uint16 {
	return 0
}

func (*AssetState) AssetUuidMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*AssetState) AssetUuidMinValue() int64 {
	return math.MinInt64 + 1
}

func (*AssetState) AssetUuidMaxValue() int64 {
	return math.MaxInt64
}

func (*AssetState) AssetUuidNullValue() int64 {
	return math.MinInt64
}

func (*AssetState) MarginCollateralWeightId() uint16 {
	return 2
}

func (*AssetState) MarginCollateralWeightSinceVersion() uint16 {
	return 0
}

func (a *AssetState) MarginCollateralWeightInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.MarginCollateralWeightSinceVersion()
}

func (*AssetState) MarginCollateralWeightDeprecated() uint16 {
	return 0
}

func (*AssetState) MarginCollateralWeightMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*AssetState) MarginCollateralWeightMinValue() int64 {
	return math.MinInt64 + 1
}

func (*AssetState) MarginCollateralWeightMaxValue() int64 {
	return math.MaxInt64
}

func (*AssetState) MarginCollateralWeightNullValue() int64 {
	return math.MinInt64
}
