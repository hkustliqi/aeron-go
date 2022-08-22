// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type InstrumentMarketDataSnapshotHeader struct {
	Timestamp       int64
	LogPosition     int64
	LastMdSeqNum    int64
	LastInstrSeqNum int64
	LastTradePrice  int64
	InstrumentId    int64
	InstrumentUuid  [2]int64
	BaseAsset       int64
	QuoteAsset      int64
	MinQuantity     int64
	BaseIncrement   int64
	QuoteIncrement  int64
	Symbol          [80]byte
}

func (i *InstrumentMarketDataSnapshotHeader) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := i.RangeCheck(i.SbeSchemaVersion(), i.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, i.Timestamp); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, i.LogPosition); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, i.LastMdSeqNum); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, i.LastInstrSeqNum); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, i.LastTradePrice); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, i.InstrumentId); err != nil {
		return err
	}
	for idx := 0; idx < 2; idx++ {
		if err := _m.WriteInt64(_w, i.InstrumentUuid[idx]); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, i.BaseAsset); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, i.QuoteAsset); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, i.MinQuantity); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, i.BaseIncrement); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, i.QuoteIncrement); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, i.Symbol[:]); err != nil {
		return err
	}
	return nil
}

func (i *InstrumentMarketDataSnapshotHeader) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !i.TimestampInActingVersion(actingVersion) {
		i.Timestamp = i.TimestampNullValue()
	} else {
		if err := _m.ReadInt64(_r, &i.Timestamp); err != nil {
			return err
		}
	}
	if !i.LogPositionInActingVersion(actingVersion) {
		i.LogPosition = i.LogPositionNullValue()
	} else {
		if err := _m.ReadInt64(_r, &i.LogPosition); err != nil {
			return err
		}
	}
	if !i.LastMdSeqNumInActingVersion(actingVersion) {
		i.LastMdSeqNum = i.LastMdSeqNumNullValue()
	} else {
		if err := _m.ReadInt64(_r, &i.LastMdSeqNum); err != nil {
			return err
		}
	}
	if !i.LastInstrSeqNumInActingVersion(actingVersion) {
		i.LastInstrSeqNum = i.LastInstrSeqNumNullValue()
	} else {
		if err := _m.ReadInt64(_r, &i.LastInstrSeqNum); err != nil {
			return err
		}
	}
	if !i.LastTradePriceInActingVersion(actingVersion) {
		i.LastTradePrice = i.LastTradePriceNullValue()
	} else {
		if err := _m.ReadInt64(_r, &i.LastTradePrice); err != nil {
			return err
		}
	}
	if !i.InstrumentIdInActingVersion(actingVersion) {
		i.InstrumentId = i.InstrumentIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &i.InstrumentId); err != nil {
			return err
		}
	}
	if !i.InstrumentUuidInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			i.InstrumentUuid[idx] = i.InstrumentUuidNullValue()
		}
	} else {
		for idx := 0; idx < 2; idx++ {
			if err := _m.ReadInt64(_r, &i.InstrumentUuid[idx]); err != nil {
				return err
			}
		}
	}
	if !i.BaseAssetInActingVersion(actingVersion) {
		i.BaseAsset = i.BaseAssetNullValue()
	} else {
		if err := _m.ReadInt64(_r, &i.BaseAsset); err != nil {
			return err
		}
	}
	if !i.QuoteAssetInActingVersion(actingVersion) {
		i.QuoteAsset = i.QuoteAssetNullValue()
	} else {
		if err := _m.ReadInt64(_r, &i.QuoteAsset); err != nil {
			return err
		}
	}
	if !i.MinQuantityInActingVersion(actingVersion) {
		i.MinQuantity = i.MinQuantityNullValue()
	} else {
		if err := _m.ReadInt64(_r, &i.MinQuantity); err != nil {
			return err
		}
	}
	if !i.BaseIncrementInActingVersion(actingVersion) {
		i.BaseIncrement = i.BaseIncrementNullValue()
	} else {
		if err := _m.ReadInt64(_r, &i.BaseIncrement); err != nil {
			return err
		}
	}
	if !i.QuoteIncrementInActingVersion(actingVersion) {
		i.QuoteIncrement = i.QuoteIncrementNullValue()
	} else {
		if err := _m.ReadInt64(_r, &i.QuoteIncrement); err != nil {
			return err
		}
	}
	if !i.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			i.Symbol[idx] = i.SymbolNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, i.Symbol[:]); err != nil {
			return err
		}
	}
	if actingVersion > i.SbeSchemaVersion() && blockLength > i.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-i.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := i.RangeCheck(actingVersion, i.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (i *InstrumentMarketDataSnapshotHeader) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if i.TimestampInActingVersion(actingVersion) {
		if i.Timestamp < i.TimestampMinValue() || i.Timestamp > i.TimestampMaxValue() {
			return fmt.Errorf("Range check failed on i.Timestamp (%v < %v > %v)", i.TimestampMinValue(), i.Timestamp, i.TimestampMaxValue())
		}
	}
	if i.LogPositionInActingVersion(actingVersion) {
		if i.LogPosition < i.LogPositionMinValue() || i.LogPosition > i.LogPositionMaxValue() {
			return fmt.Errorf("Range check failed on i.LogPosition (%v < %v > %v)", i.LogPositionMinValue(), i.LogPosition, i.LogPositionMaxValue())
		}
	}
	if i.LastMdSeqNumInActingVersion(actingVersion) {
		if i.LastMdSeqNum < i.LastMdSeqNumMinValue() || i.LastMdSeqNum > i.LastMdSeqNumMaxValue() {
			return fmt.Errorf("Range check failed on i.LastMdSeqNum (%v < %v > %v)", i.LastMdSeqNumMinValue(), i.LastMdSeqNum, i.LastMdSeqNumMaxValue())
		}
	}
	if i.LastInstrSeqNumInActingVersion(actingVersion) {
		if i.LastInstrSeqNum < i.LastInstrSeqNumMinValue() || i.LastInstrSeqNum > i.LastInstrSeqNumMaxValue() {
			return fmt.Errorf("Range check failed on i.LastInstrSeqNum (%v < %v > %v)", i.LastInstrSeqNumMinValue(), i.LastInstrSeqNum, i.LastInstrSeqNumMaxValue())
		}
	}
	if i.LastTradePriceInActingVersion(actingVersion) {
		if i.LastTradePrice < i.LastTradePriceMinValue() || i.LastTradePrice > i.LastTradePriceMaxValue() {
			return fmt.Errorf("Range check failed on i.LastTradePrice (%v < %v > %v)", i.LastTradePriceMinValue(), i.LastTradePrice, i.LastTradePriceMaxValue())
		}
	}
	if i.InstrumentIdInActingVersion(actingVersion) {
		if i.InstrumentId < i.InstrumentIdMinValue() || i.InstrumentId > i.InstrumentIdMaxValue() {
			return fmt.Errorf("Range check failed on i.InstrumentId (%v < %v > %v)", i.InstrumentIdMinValue(), i.InstrumentId, i.InstrumentIdMaxValue())
		}
	}
	if i.InstrumentUuidInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if i.InstrumentUuid[idx] < i.InstrumentUuidMinValue() || i.InstrumentUuid[idx] > i.InstrumentUuidMaxValue() {
				return fmt.Errorf("Range check failed on i.InstrumentUuid[%d] (%v < %v > %v)", idx, i.InstrumentUuidMinValue(), i.InstrumentUuid[idx], i.InstrumentUuidMaxValue())
			}
		}
	}
	if i.BaseAssetInActingVersion(actingVersion) {
		if i.BaseAsset < i.BaseAssetMinValue() || i.BaseAsset > i.BaseAssetMaxValue() {
			return fmt.Errorf("Range check failed on i.BaseAsset (%v < %v > %v)", i.BaseAssetMinValue(), i.BaseAsset, i.BaseAssetMaxValue())
		}
	}
	if i.QuoteAssetInActingVersion(actingVersion) {
		if i.QuoteAsset < i.QuoteAssetMinValue() || i.QuoteAsset > i.QuoteAssetMaxValue() {
			return fmt.Errorf("Range check failed on i.QuoteAsset (%v < %v > %v)", i.QuoteAssetMinValue(), i.QuoteAsset, i.QuoteAssetMaxValue())
		}
	}
	if i.MinQuantityInActingVersion(actingVersion) {
		if i.MinQuantity < i.MinQuantityMinValue() || i.MinQuantity > i.MinQuantityMaxValue() {
			return fmt.Errorf("Range check failed on i.MinQuantity (%v < %v > %v)", i.MinQuantityMinValue(), i.MinQuantity, i.MinQuantityMaxValue())
		}
	}
	if i.BaseIncrementInActingVersion(actingVersion) {
		if i.BaseIncrement < i.BaseIncrementMinValue() || i.BaseIncrement > i.BaseIncrementMaxValue() {
			return fmt.Errorf("Range check failed on i.BaseIncrement (%v < %v > %v)", i.BaseIncrementMinValue(), i.BaseIncrement, i.BaseIncrementMaxValue())
		}
	}
	if i.QuoteIncrementInActingVersion(actingVersion) {
		if i.QuoteIncrement < i.QuoteIncrementMinValue() || i.QuoteIncrement > i.QuoteIncrementMaxValue() {
			return fmt.Errorf("Range check failed on i.QuoteIncrement (%v < %v > %v)", i.QuoteIncrementMinValue(), i.QuoteIncrement, i.QuoteIncrementMaxValue())
		}
	}
	if i.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			if i.Symbol[idx] < i.SymbolMinValue() || i.Symbol[idx] > i.SymbolMaxValue() {
				return fmt.Errorf("Range check failed on i.Symbol[%d] (%v < %v > %v)", idx, i.SymbolMinValue(), i.Symbol[idx], i.SymbolMaxValue())
			}
		}
	}
	for idx, ch := range i.Symbol {
		if ch > 127 {
			return fmt.Errorf("i.Symbol[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	return nil
}

func InstrumentMarketDataSnapshotHeaderInit(i *InstrumentMarketDataSnapshotHeader) {
	return
}

func (*InstrumentMarketDataSnapshotHeader) SbeBlockLength() (blockLength uint16) {
	return 184
}

func (*InstrumentMarketDataSnapshotHeader) SbeTemplateId() (templateId uint16) {
	return 2113
}

func (*InstrumentMarketDataSnapshotHeader) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*InstrumentMarketDataSnapshotHeader) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*InstrumentMarketDataSnapshotHeader) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*InstrumentMarketDataSnapshotHeader) TimestampId() uint16 {
	return 1
}

func (*InstrumentMarketDataSnapshotHeader) TimestampSinceVersion() uint16 {
	return 0
}

func (i *InstrumentMarketDataSnapshotHeader) TimestampInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.TimestampSinceVersion()
}

func (*InstrumentMarketDataSnapshotHeader) TimestampDeprecated() uint16 {
	return 0
}

func (*InstrumentMarketDataSnapshotHeader) TimestampMetaAttribute(meta int) string {
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

func (*InstrumentMarketDataSnapshotHeader) TimestampMinValue() int64 {
	return math.MinInt64 + 1
}

func (*InstrumentMarketDataSnapshotHeader) TimestampMaxValue() int64 {
	return math.MaxInt64
}

func (*InstrumentMarketDataSnapshotHeader) TimestampNullValue() int64 {
	return math.MinInt64
}

func (*InstrumentMarketDataSnapshotHeader) LogPositionId() uint16 {
	return 2
}

func (*InstrumentMarketDataSnapshotHeader) LogPositionSinceVersion() uint16 {
	return 0
}

func (i *InstrumentMarketDataSnapshotHeader) LogPositionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.LogPositionSinceVersion()
}

func (*InstrumentMarketDataSnapshotHeader) LogPositionDeprecated() uint16 {
	return 0
}

func (*InstrumentMarketDataSnapshotHeader) LogPositionMetaAttribute(meta int) string {
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

func (*InstrumentMarketDataSnapshotHeader) LogPositionMinValue() int64 {
	return math.MinInt64 + 1
}

func (*InstrumentMarketDataSnapshotHeader) LogPositionMaxValue() int64 {
	return math.MaxInt64
}

func (*InstrumentMarketDataSnapshotHeader) LogPositionNullValue() int64 {
	return math.MinInt64
}

func (*InstrumentMarketDataSnapshotHeader) LastMdSeqNumId() uint16 {
	return 3
}

func (*InstrumentMarketDataSnapshotHeader) LastMdSeqNumSinceVersion() uint16 {
	return 0
}

func (i *InstrumentMarketDataSnapshotHeader) LastMdSeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.LastMdSeqNumSinceVersion()
}

func (*InstrumentMarketDataSnapshotHeader) LastMdSeqNumDeprecated() uint16 {
	return 0
}

func (*InstrumentMarketDataSnapshotHeader) LastMdSeqNumMetaAttribute(meta int) string {
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

func (*InstrumentMarketDataSnapshotHeader) LastMdSeqNumMinValue() int64 {
	return math.MinInt64 + 1
}

func (*InstrumentMarketDataSnapshotHeader) LastMdSeqNumMaxValue() int64 {
	return math.MaxInt64
}

func (*InstrumentMarketDataSnapshotHeader) LastMdSeqNumNullValue() int64 {
	return math.MinInt64
}

func (*InstrumentMarketDataSnapshotHeader) LastInstrSeqNumId() uint16 {
	return 4
}

func (*InstrumentMarketDataSnapshotHeader) LastInstrSeqNumSinceVersion() uint16 {
	return 0
}

func (i *InstrumentMarketDataSnapshotHeader) LastInstrSeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.LastInstrSeqNumSinceVersion()
}

func (*InstrumentMarketDataSnapshotHeader) LastInstrSeqNumDeprecated() uint16 {
	return 0
}

func (*InstrumentMarketDataSnapshotHeader) LastInstrSeqNumMetaAttribute(meta int) string {
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

func (*InstrumentMarketDataSnapshotHeader) LastInstrSeqNumMinValue() int64 {
	return math.MinInt64 + 1
}

func (*InstrumentMarketDataSnapshotHeader) LastInstrSeqNumMaxValue() int64 {
	return math.MaxInt64
}

func (*InstrumentMarketDataSnapshotHeader) LastInstrSeqNumNullValue() int64 {
	return math.MinInt64
}

func (*InstrumentMarketDataSnapshotHeader) LastTradePriceId() uint16 {
	return 5
}

func (*InstrumentMarketDataSnapshotHeader) LastTradePriceSinceVersion() uint16 {
	return 0
}

func (i *InstrumentMarketDataSnapshotHeader) LastTradePriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.LastTradePriceSinceVersion()
}

func (*InstrumentMarketDataSnapshotHeader) LastTradePriceDeprecated() uint16 {
	return 0
}

func (*InstrumentMarketDataSnapshotHeader) LastTradePriceMetaAttribute(meta int) string {
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

func (*InstrumentMarketDataSnapshotHeader) LastTradePriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*InstrumentMarketDataSnapshotHeader) LastTradePriceMaxValue() int64 {
	return math.MaxInt64
}

func (*InstrumentMarketDataSnapshotHeader) LastTradePriceNullValue() int64 {
	return math.MinInt64
}

func (*InstrumentMarketDataSnapshotHeader) InstrumentIdId() uint16 {
	return 6
}

func (*InstrumentMarketDataSnapshotHeader) InstrumentIdSinceVersion() uint16 {
	return 0
}

func (i *InstrumentMarketDataSnapshotHeader) InstrumentIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.InstrumentIdSinceVersion()
}

func (*InstrumentMarketDataSnapshotHeader) InstrumentIdDeprecated() uint16 {
	return 0
}

func (*InstrumentMarketDataSnapshotHeader) InstrumentIdMetaAttribute(meta int) string {
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

func (*InstrumentMarketDataSnapshotHeader) InstrumentIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*InstrumentMarketDataSnapshotHeader) InstrumentIdMaxValue() int64 {
	return math.MaxInt64
}

func (*InstrumentMarketDataSnapshotHeader) InstrumentIdNullValue() int64 {
	return math.MinInt64
}

func (*InstrumentMarketDataSnapshotHeader) InstrumentUuidId() uint16 {
	return 7
}

func (*InstrumentMarketDataSnapshotHeader) InstrumentUuidSinceVersion() uint16 {
	return 0
}

func (i *InstrumentMarketDataSnapshotHeader) InstrumentUuidInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.InstrumentUuidSinceVersion()
}

func (*InstrumentMarketDataSnapshotHeader) InstrumentUuidDeprecated() uint16 {
	return 0
}

func (*InstrumentMarketDataSnapshotHeader) InstrumentUuidMetaAttribute(meta int) string {
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

func (*InstrumentMarketDataSnapshotHeader) InstrumentUuidMinValue() int64 {
	return math.MinInt64 + 1
}

func (*InstrumentMarketDataSnapshotHeader) InstrumentUuidMaxValue() int64 {
	return math.MaxInt64
}

func (*InstrumentMarketDataSnapshotHeader) InstrumentUuidNullValue() int64 {
	return math.MinInt64
}

func (*InstrumentMarketDataSnapshotHeader) BaseAssetId() uint16 {
	return 8
}

func (*InstrumentMarketDataSnapshotHeader) BaseAssetSinceVersion() uint16 {
	return 0
}

func (i *InstrumentMarketDataSnapshotHeader) BaseAssetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.BaseAssetSinceVersion()
}

func (*InstrumentMarketDataSnapshotHeader) BaseAssetDeprecated() uint16 {
	return 0
}

func (*InstrumentMarketDataSnapshotHeader) BaseAssetMetaAttribute(meta int) string {
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

func (*InstrumentMarketDataSnapshotHeader) BaseAssetMinValue() int64 {
	return math.MinInt64 + 1
}

func (*InstrumentMarketDataSnapshotHeader) BaseAssetMaxValue() int64 {
	return math.MaxInt64
}

func (*InstrumentMarketDataSnapshotHeader) BaseAssetNullValue() int64 {
	return math.MinInt64
}

func (*InstrumentMarketDataSnapshotHeader) QuoteAssetId() uint16 {
	return 9
}

func (*InstrumentMarketDataSnapshotHeader) QuoteAssetSinceVersion() uint16 {
	return 0
}

func (i *InstrumentMarketDataSnapshotHeader) QuoteAssetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.QuoteAssetSinceVersion()
}

func (*InstrumentMarketDataSnapshotHeader) QuoteAssetDeprecated() uint16 {
	return 0
}

func (*InstrumentMarketDataSnapshotHeader) QuoteAssetMetaAttribute(meta int) string {
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

func (*InstrumentMarketDataSnapshotHeader) QuoteAssetMinValue() int64 {
	return math.MinInt64 + 1
}

func (*InstrumentMarketDataSnapshotHeader) QuoteAssetMaxValue() int64 {
	return math.MaxInt64
}

func (*InstrumentMarketDataSnapshotHeader) QuoteAssetNullValue() int64 {
	return math.MinInt64
}

func (*InstrumentMarketDataSnapshotHeader) MinQuantityId() uint16 {
	return 10
}

func (*InstrumentMarketDataSnapshotHeader) MinQuantitySinceVersion() uint16 {
	return 0
}

func (i *InstrumentMarketDataSnapshotHeader) MinQuantityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.MinQuantitySinceVersion()
}

func (*InstrumentMarketDataSnapshotHeader) MinQuantityDeprecated() uint16 {
	return 0
}

func (*InstrumentMarketDataSnapshotHeader) MinQuantityMetaAttribute(meta int) string {
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

func (*InstrumentMarketDataSnapshotHeader) MinQuantityMinValue() int64 {
	return math.MinInt64 + 1
}

func (*InstrumentMarketDataSnapshotHeader) MinQuantityMaxValue() int64 {
	return math.MaxInt64
}

func (*InstrumentMarketDataSnapshotHeader) MinQuantityNullValue() int64 {
	return math.MinInt64
}

func (*InstrumentMarketDataSnapshotHeader) BaseIncrementId() uint16 {
	return 11
}

func (*InstrumentMarketDataSnapshotHeader) BaseIncrementSinceVersion() uint16 {
	return 0
}

func (i *InstrumentMarketDataSnapshotHeader) BaseIncrementInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.BaseIncrementSinceVersion()
}

func (*InstrumentMarketDataSnapshotHeader) BaseIncrementDeprecated() uint16 {
	return 0
}

func (*InstrumentMarketDataSnapshotHeader) BaseIncrementMetaAttribute(meta int) string {
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

func (*InstrumentMarketDataSnapshotHeader) BaseIncrementMinValue() int64 {
	return math.MinInt64 + 1
}

func (*InstrumentMarketDataSnapshotHeader) BaseIncrementMaxValue() int64 {
	return math.MaxInt64
}

func (*InstrumentMarketDataSnapshotHeader) BaseIncrementNullValue() int64 {
	return math.MinInt64
}

func (*InstrumentMarketDataSnapshotHeader) QuoteIncrementId() uint16 {
	return 12
}

func (*InstrumentMarketDataSnapshotHeader) QuoteIncrementSinceVersion() uint16 {
	return 0
}

func (i *InstrumentMarketDataSnapshotHeader) QuoteIncrementInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.QuoteIncrementSinceVersion()
}

func (*InstrumentMarketDataSnapshotHeader) QuoteIncrementDeprecated() uint16 {
	return 0
}

func (*InstrumentMarketDataSnapshotHeader) QuoteIncrementMetaAttribute(meta int) string {
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

func (*InstrumentMarketDataSnapshotHeader) QuoteIncrementMinValue() int64 {
	return math.MinInt64 + 1
}

func (*InstrumentMarketDataSnapshotHeader) QuoteIncrementMaxValue() int64 {
	return math.MaxInt64
}

func (*InstrumentMarketDataSnapshotHeader) QuoteIncrementNullValue() int64 {
	return math.MinInt64
}

func (*InstrumentMarketDataSnapshotHeader) SymbolId() uint16 {
	return 13
}

func (*InstrumentMarketDataSnapshotHeader) SymbolSinceVersion() uint16 {
	return 0
}

func (i *InstrumentMarketDataSnapshotHeader) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.SymbolSinceVersion()
}

func (*InstrumentMarketDataSnapshotHeader) SymbolDeprecated() uint16 {
	return 0
}

func (*InstrumentMarketDataSnapshotHeader) SymbolMetaAttribute(meta int) string {
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

func (*InstrumentMarketDataSnapshotHeader) SymbolMinValue() byte {
	return byte(32)
}

func (*InstrumentMarketDataSnapshotHeader) SymbolMaxValue() byte {
	return byte(126)
}

func (*InstrumentMarketDataSnapshotHeader) SymbolNullValue() byte {
	return 0
}

func (i *InstrumentMarketDataSnapshotHeader) SymbolCharacterEncoding() string {
	return "US-ASCII"
}
