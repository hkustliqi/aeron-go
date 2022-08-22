// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type ActiveMarketDataSnapshot struct {
	InstrumentId  int64
	Timestamp     int64
	OrderId       int64
	Price         int64
	Quantity      int64
	SnapshotIndex int32
	Side          SideEnum
}

func (a *ActiveMarketDataSnapshot) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := a.RangeCheck(a.SbeSchemaVersion(), a.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, a.InstrumentId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, a.Timestamp); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, a.OrderId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, a.Price); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, a.Quantity); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, a.SnapshotIndex); err != nil {
		return err
	}
	if err := a.Side.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (a *ActiveMarketDataSnapshot) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !a.InstrumentIdInActingVersion(actingVersion) {
		a.InstrumentId = a.InstrumentIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &a.InstrumentId); err != nil {
			return err
		}
	}
	if !a.TimestampInActingVersion(actingVersion) {
		a.Timestamp = a.TimestampNullValue()
	} else {
		if err := _m.ReadInt64(_r, &a.Timestamp); err != nil {
			return err
		}
	}
	if !a.OrderIdInActingVersion(actingVersion) {
		a.OrderId = a.OrderIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &a.OrderId); err != nil {
			return err
		}
	}
	if !a.PriceInActingVersion(actingVersion) {
		a.Price = a.PriceNullValue()
	} else {
		if err := _m.ReadInt64(_r, &a.Price); err != nil {
			return err
		}
	}
	if !a.QuantityInActingVersion(actingVersion) {
		a.Quantity = a.QuantityNullValue()
	} else {
		if err := _m.ReadInt64(_r, &a.Quantity); err != nil {
			return err
		}
	}
	if !a.SnapshotIndexInActingVersion(actingVersion) {
		a.SnapshotIndex = a.SnapshotIndexNullValue()
	} else {
		if err := _m.ReadInt32(_r, &a.SnapshotIndex); err != nil {
			return err
		}
	}
	if a.SideInActingVersion(actingVersion) {
		if err := a.Side.Decode(_m, _r, actingVersion); err != nil {
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

func (a *ActiveMarketDataSnapshot) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if a.InstrumentIdInActingVersion(actingVersion) {
		if a.InstrumentId < a.InstrumentIdMinValue() || a.InstrumentId > a.InstrumentIdMaxValue() {
			return fmt.Errorf("Range check failed on a.InstrumentId (%v < %v > %v)", a.InstrumentIdMinValue(), a.InstrumentId, a.InstrumentIdMaxValue())
		}
	}
	if a.TimestampInActingVersion(actingVersion) {
		if a.Timestamp < a.TimestampMinValue() || a.Timestamp > a.TimestampMaxValue() {
			return fmt.Errorf("Range check failed on a.Timestamp (%v < %v > %v)", a.TimestampMinValue(), a.Timestamp, a.TimestampMaxValue())
		}
	}
	if a.OrderIdInActingVersion(actingVersion) {
		if a.OrderId < a.OrderIdMinValue() || a.OrderId > a.OrderIdMaxValue() {
			return fmt.Errorf("Range check failed on a.OrderId (%v < %v > %v)", a.OrderIdMinValue(), a.OrderId, a.OrderIdMaxValue())
		}
	}
	if a.PriceInActingVersion(actingVersion) {
		if a.Price < a.PriceMinValue() || a.Price > a.PriceMaxValue() {
			return fmt.Errorf("Range check failed on a.Price (%v < %v > %v)", a.PriceMinValue(), a.Price, a.PriceMaxValue())
		}
	}
	if a.QuantityInActingVersion(actingVersion) {
		if a.Quantity < a.QuantityMinValue() || a.Quantity > a.QuantityMaxValue() {
			return fmt.Errorf("Range check failed on a.Quantity (%v < %v > %v)", a.QuantityMinValue(), a.Quantity, a.QuantityMaxValue())
		}
	}
	if a.SnapshotIndexInActingVersion(actingVersion) {
		if a.SnapshotIndex < a.SnapshotIndexMinValue() || a.SnapshotIndex > a.SnapshotIndexMaxValue() {
			return fmt.Errorf("Range check failed on a.SnapshotIndex (%v < %v > %v)", a.SnapshotIndexMinValue(), a.SnapshotIndex, a.SnapshotIndexMaxValue())
		}
	}
	if err := a.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func ActiveMarketDataSnapshotInit(a *ActiveMarketDataSnapshot) {
	return
}

func (*ActiveMarketDataSnapshot) SbeBlockLength() (blockLength uint16) {
	return 45
}

func (*ActiveMarketDataSnapshot) SbeTemplateId() (templateId uint16) {
	return 2114
}

func (*ActiveMarketDataSnapshot) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*ActiveMarketDataSnapshot) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*ActiveMarketDataSnapshot) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*ActiveMarketDataSnapshot) InstrumentIdId() uint16 {
	return 1
}

func (*ActiveMarketDataSnapshot) InstrumentIdSinceVersion() uint16 {
	return 0
}

func (a *ActiveMarketDataSnapshot) InstrumentIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.InstrumentIdSinceVersion()
}

func (*ActiveMarketDataSnapshot) InstrumentIdDeprecated() uint16 {
	return 0
}

func (*ActiveMarketDataSnapshot) InstrumentIdMetaAttribute(meta int) string {
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

func (*ActiveMarketDataSnapshot) InstrumentIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ActiveMarketDataSnapshot) InstrumentIdMaxValue() int64 {
	return math.MaxInt64
}

func (*ActiveMarketDataSnapshot) InstrumentIdNullValue() int64 {
	return math.MinInt64
}

func (*ActiveMarketDataSnapshot) TimestampId() uint16 {
	return 2
}

func (*ActiveMarketDataSnapshot) TimestampSinceVersion() uint16 {
	return 0
}

func (a *ActiveMarketDataSnapshot) TimestampInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.TimestampSinceVersion()
}

func (*ActiveMarketDataSnapshot) TimestampDeprecated() uint16 {
	return 0
}

func (*ActiveMarketDataSnapshot) TimestampMetaAttribute(meta int) string {
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

func (*ActiveMarketDataSnapshot) TimestampMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ActiveMarketDataSnapshot) TimestampMaxValue() int64 {
	return math.MaxInt64
}

func (*ActiveMarketDataSnapshot) TimestampNullValue() int64 {
	return math.MinInt64
}

func (*ActiveMarketDataSnapshot) OrderIdId() uint16 {
	return 3
}

func (*ActiveMarketDataSnapshot) OrderIdSinceVersion() uint16 {
	return 0
}

func (a *ActiveMarketDataSnapshot) OrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.OrderIdSinceVersion()
}

func (*ActiveMarketDataSnapshot) OrderIdDeprecated() uint16 {
	return 0
}

func (*ActiveMarketDataSnapshot) OrderIdMetaAttribute(meta int) string {
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

func (*ActiveMarketDataSnapshot) OrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ActiveMarketDataSnapshot) OrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*ActiveMarketDataSnapshot) OrderIdNullValue() int64 {
	return math.MinInt64
}

func (*ActiveMarketDataSnapshot) PriceId() uint16 {
	return 4
}

func (*ActiveMarketDataSnapshot) PriceSinceVersion() uint16 {
	return 0
}

func (a *ActiveMarketDataSnapshot) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.PriceSinceVersion()
}

func (*ActiveMarketDataSnapshot) PriceDeprecated() uint16 {
	return 0
}

func (*ActiveMarketDataSnapshot) PriceMetaAttribute(meta int) string {
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

func (*ActiveMarketDataSnapshot) PriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ActiveMarketDataSnapshot) PriceMaxValue() int64 {
	return math.MaxInt64
}

func (*ActiveMarketDataSnapshot) PriceNullValue() int64 {
	return math.MinInt64
}

func (*ActiveMarketDataSnapshot) QuantityId() uint16 {
	return 5
}

func (*ActiveMarketDataSnapshot) QuantitySinceVersion() uint16 {
	return 0
}

func (a *ActiveMarketDataSnapshot) QuantityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.QuantitySinceVersion()
}

func (*ActiveMarketDataSnapshot) QuantityDeprecated() uint16 {
	return 0
}

func (*ActiveMarketDataSnapshot) QuantityMetaAttribute(meta int) string {
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

func (*ActiveMarketDataSnapshot) QuantityMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ActiveMarketDataSnapshot) QuantityMaxValue() int64 {
	return math.MaxInt64
}

func (*ActiveMarketDataSnapshot) QuantityNullValue() int64 {
	return math.MinInt64
}

func (*ActiveMarketDataSnapshot) SnapshotIndexId() uint16 {
	return 6
}

func (*ActiveMarketDataSnapshot) SnapshotIndexSinceVersion() uint16 {
	return 0
}

func (a *ActiveMarketDataSnapshot) SnapshotIndexInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.SnapshotIndexSinceVersion()
}

func (*ActiveMarketDataSnapshot) SnapshotIndexDeprecated() uint16 {
	return 0
}

func (*ActiveMarketDataSnapshot) SnapshotIndexMetaAttribute(meta int) string {
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

func (*ActiveMarketDataSnapshot) SnapshotIndexMinValue() int32 {
	return math.MinInt32 + 1
}

func (*ActiveMarketDataSnapshot) SnapshotIndexMaxValue() int32 {
	return math.MaxInt32
}

func (*ActiveMarketDataSnapshot) SnapshotIndexNullValue() int32 {
	return math.MinInt32
}

func (*ActiveMarketDataSnapshot) SideId() uint16 {
	return 7
}

func (*ActiveMarketDataSnapshot) SideSinceVersion() uint16 {
	return 0
}

func (a *ActiveMarketDataSnapshot) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.SideSinceVersion()
}

func (*ActiveMarketDataSnapshot) SideDeprecated() uint16 {
	return 0
}

func (*ActiveMarketDataSnapshot) SideMetaAttribute(meta int) string {
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
