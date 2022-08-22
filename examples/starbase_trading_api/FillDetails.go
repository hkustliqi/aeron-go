// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type FillDetails struct {
	MatchId    int64
	FillPrice  int64
	FillQty    int64
	NewTradeId int64
}

func (f *FillDetails) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := f.RangeCheck(f.SbeSchemaVersion(), f.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, f.MatchId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, f.FillPrice); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, f.FillQty); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, f.NewTradeId); err != nil {
		return err
	}
	return nil
}

func (f *FillDetails) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !f.MatchIdInActingVersion(actingVersion) {
		f.MatchId = f.MatchIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &f.MatchId); err != nil {
			return err
		}
	}
	if !f.FillPriceInActingVersion(actingVersion) {
		f.FillPrice = f.FillPriceNullValue()
	} else {
		if err := _m.ReadInt64(_r, &f.FillPrice); err != nil {
			return err
		}
	}
	if !f.FillQtyInActingVersion(actingVersion) {
		f.FillQty = f.FillQtyNullValue()
	} else {
		if err := _m.ReadInt64(_r, &f.FillQty); err != nil {
			return err
		}
	}
	if !f.NewTradeIdInActingVersion(actingVersion) {
		f.NewTradeId = f.NewTradeIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &f.NewTradeId); err != nil {
			return err
		}
	}
	if actingVersion > f.SbeSchemaVersion() && blockLength > f.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-f.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := f.RangeCheck(actingVersion, f.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (f *FillDetails) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if f.MatchIdInActingVersion(actingVersion) {
		if f.MatchId < f.MatchIdMinValue() || f.MatchId > f.MatchIdMaxValue() {
			return fmt.Errorf("Range check failed on f.MatchId (%v < %v > %v)", f.MatchIdMinValue(), f.MatchId, f.MatchIdMaxValue())
		}
	}
	if f.FillPriceInActingVersion(actingVersion) {
		if f.FillPrice < f.FillPriceMinValue() || f.FillPrice > f.FillPriceMaxValue() {
			return fmt.Errorf("Range check failed on f.FillPrice (%v < %v > %v)", f.FillPriceMinValue(), f.FillPrice, f.FillPriceMaxValue())
		}
	}
	if f.FillQtyInActingVersion(actingVersion) {
		if f.FillQty < f.FillQtyMinValue() || f.FillQty > f.FillQtyMaxValue() {
			return fmt.Errorf("Range check failed on f.FillQty (%v < %v > %v)", f.FillQtyMinValue(), f.FillQty, f.FillQtyMaxValue())
		}
	}
	if f.NewTradeIdInActingVersion(actingVersion) {
		if f.NewTradeId < f.NewTradeIdMinValue() || f.NewTradeId > f.NewTradeIdMaxValue() {
			return fmt.Errorf("Range check failed on f.NewTradeId (%v < %v > %v)", f.NewTradeIdMinValue(), f.NewTradeId, f.NewTradeIdMaxValue())
		}
	}
	return nil
}

func FillDetailsInit(f *FillDetails) {
	return
}

func (*FillDetails) SbeBlockLength() (blockLength uint16) {
	return 32
}

func (*FillDetails) SbeTemplateId() (templateId uint16) {
	return 213
}

func (*FillDetails) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*FillDetails) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*FillDetails) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*FillDetails) MatchIdId() uint16 {
	return 1
}

func (*FillDetails) MatchIdSinceVersion() uint16 {
	return 0
}

func (f *FillDetails) MatchIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= f.MatchIdSinceVersion()
}

func (*FillDetails) MatchIdDeprecated() uint16 {
	return 0
}

func (*FillDetails) MatchIdMetaAttribute(meta int) string {
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

func (*FillDetails) MatchIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*FillDetails) MatchIdMaxValue() int64 {
	return math.MaxInt64
}

func (*FillDetails) MatchIdNullValue() int64 {
	return math.MinInt64
}

func (*FillDetails) FillPriceId() uint16 {
	return 2
}

func (*FillDetails) FillPriceSinceVersion() uint16 {
	return 0
}

func (f *FillDetails) FillPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= f.FillPriceSinceVersion()
}

func (*FillDetails) FillPriceDeprecated() uint16 {
	return 0
}

func (*FillDetails) FillPriceMetaAttribute(meta int) string {
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

func (*FillDetails) FillPriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*FillDetails) FillPriceMaxValue() int64 {
	return math.MaxInt64
}

func (*FillDetails) FillPriceNullValue() int64 {
	return math.MinInt64
}

func (*FillDetails) FillQtyId() uint16 {
	return 3
}

func (*FillDetails) FillQtySinceVersion() uint16 {
	return 0
}

func (f *FillDetails) FillQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= f.FillQtySinceVersion()
}

func (*FillDetails) FillQtyDeprecated() uint16 {
	return 0
}

func (*FillDetails) FillQtyMetaAttribute(meta int) string {
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

func (*FillDetails) FillQtyMinValue() int64 {
	return math.MinInt64 + 1
}

func (*FillDetails) FillQtyMaxValue() int64 {
	return math.MaxInt64
}

func (*FillDetails) FillQtyNullValue() int64 {
	return math.MinInt64
}

func (*FillDetails) NewTradeIdId() uint16 {
	return 4
}

func (*FillDetails) NewTradeIdSinceVersion() uint16 {
	return 0
}

func (f *FillDetails) NewTradeIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= f.NewTradeIdSinceVersion()
}

func (*FillDetails) NewTradeIdDeprecated() uint16 {
	return 0
}

func (*FillDetails) NewTradeIdMetaAttribute(meta int) string {
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

func (*FillDetails) NewTradeIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*FillDetails) NewTradeIdMaxValue() int64 {
	return math.MaxInt64
}

func (*FillDetails) NewTradeIdNullValue() int64 {
	return math.MinInt64
}
