// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type EngBookState struct {
	InstrumentId                    int64
	InstrumentUuid                  [2]int64
	BaseAsset                       int64
	QuoteAsset                      int64
	UnderlyingSpot                  int64
	SettlementInterval              int64
	MinQuantity                     int64
	BaseIncrement                   int64
	QuoteIncrement                  int64
	InitialMarginFractionFactor     int64
	MaintenanceMarginFractionFactor int64
	LastTradePrice                  int64
	Symbol                          [80]byte
	OrderQueueId                    int64
	NextMdInstrSeqNo                int64
	InstrumentType                  InstrumentTypeEnum
}

func (e *EngBookState) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := e.RangeCheck(e.SbeSchemaVersion(), e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, e.InstrumentId); err != nil {
		return err
	}
	for idx := 0; idx < 2; idx++ {
		if err := _m.WriteInt64(_w, e.InstrumentUuid[idx]); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, e.BaseAsset); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.QuoteAsset); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.UnderlyingSpot); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.SettlementInterval); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.MinQuantity); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.BaseIncrement); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.QuoteIncrement); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.InitialMarginFractionFactor); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.MaintenanceMarginFractionFactor); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.LastTradePrice); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.Symbol[:]); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.OrderQueueId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.NextMdInstrSeqNo); err != nil {
		return err
	}
	if err := e.InstrumentType.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (e *EngBookState) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !e.InstrumentIdInActingVersion(actingVersion) {
		e.InstrumentId = e.InstrumentIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.InstrumentId); err != nil {
			return err
		}
	}
	if !e.InstrumentUuidInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			e.InstrumentUuid[idx] = e.InstrumentUuidNullValue()
		}
	} else {
		for idx := 0; idx < 2; idx++ {
			if err := _m.ReadInt64(_r, &e.InstrumentUuid[idx]); err != nil {
				return err
			}
		}
	}
	if !e.BaseAssetInActingVersion(actingVersion) {
		e.BaseAsset = e.BaseAssetNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.BaseAsset); err != nil {
			return err
		}
	}
	if !e.QuoteAssetInActingVersion(actingVersion) {
		e.QuoteAsset = e.QuoteAssetNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.QuoteAsset); err != nil {
			return err
		}
	}
	if !e.UnderlyingSpotInActingVersion(actingVersion) {
		e.UnderlyingSpot = e.UnderlyingSpotNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.UnderlyingSpot); err != nil {
			return err
		}
	}
	if !e.SettlementIntervalInActingVersion(actingVersion) {
		e.SettlementInterval = e.SettlementIntervalNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.SettlementInterval); err != nil {
			return err
		}
	}
	if !e.MinQuantityInActingVersion(actingVersion) {
		e.MinQuantity = e.MinQuantityNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.MinQuantity); err != nil {
			return err
		}
	}
	if !e.BaseIncrementInActingVersion(actingVersion) {
		e.BaseIncrement = e.BaseIncrementNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.BaseIncrement); err != nil {
			return err
		}
	}
	if !e.QuoteIncrementInActingVersion(actingVersion) {
		e.QuoteIncrement = e.QuoteIncrementNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.QuoteIncrement); err != nil {
			return err
		}
	}
	if !e.InitialMarginFractionFactorInActingVersion(actingVersion) {
		e.InitialMarginFractionFactor = e.InitialMarginFractionFactorNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.InitialMarginFractionFactor); err != nil {
			return err
		}
	}
	if !e.MaintenanceMarginFractionFactorInActingVersion(actingVersion) {
		e.MaintenanceMarginFractionFactor = e.MaintenanceMarginFractionFactorNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.MaintenanceMarginFractionFactor); err != nil {
			return err
		}
	}
	if !e.LastTradePriceInActingVersion(actingVersion) {
		e.LastTradePrice = e.LastTradePriceNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.LastTradePrice); err != nil {
			return err
		}
	}
	if !e.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			e.Symbol[idx] = e.SymbolNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, e.Symbol[:]); err != nil {
			return err
		}
	}
	if !e.OrderQueueIdInActingVersion(actingVersion) {
		e.OrderQueueId = e.OrderQueueIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.OrderQueueId); err != nil {
			return err
		}
	}
	if !e.NextMdInstrSeqNoInActingVersion(actingVersion) {
		e.NextMdInstrSeqNo = e.NextMdInstrSeqNoNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.NextMdInstrSeqNo); err != nil {
			return err
		}
	}
	if e.InstrumentTypeInActingVersion(actingVersion) {
		if err := e.InstrumentType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > e.SbeSchemaVersion() && blockLength > e.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-e.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := e.RangeCheck(actingVersion, e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (e *EngBookState) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.InstrumentIdInActingVersion(actingVersion) {
		if e.InstrumentId < e.InstrumentIdMinValue() || e.InstrumentId > e.InstrumentIdMaxValue() {
			return fmt.Errorf("Range check failed on e.InstrumentId (%v < %v > %v)", e.InstrumentIdMinValue(), e.InstrumentId, e.InstrumentIdMaxValue())
		}
	}
	if e.InstrumentUuidInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if e.InstrumentUuid[idx] < e.InstrumentUuidMinValue() || e.InstrumentUuid[idx] > e.InstrumentUuidMaxValue() {
				return fmt.Errorf("Range check failed on e.InstrumentUuid[%d] (%v < %v > %v)", idx, e.InstrumentUuidMinValue(), e.InstrumentUuid[idx], e.InstrumentUuidMaxValue())
			}
		}
	}
	if e.BaseAssetInActingVersion(actingVersion) {
		if e.BaseAsset < e.BaseAssetMinValue() || e.BaseAsset > e.BaseAssetMaxValue() {
			return fmt.Errorf("Range check failed on e.BaseAsset (%v < %v > %v)", e.BaseAssetMinValue(), e.BaseAsset, e.BaseAssetMaxValue())
		}
	}
	if e.QuoteAssetInActingVersion(actingVersion) {
		if e.QuoteAsset < e.QuoteAssetMinValue() || e.QuoteAsset > e.QuoteAssetMaxValue() {
			return fmt.Errorf("Range check failed on e.QuoteAsset (%v < %v > %v)", e.QuoteAssetMinValue(), e.QuoteAsset, e.QuoteAssetMaxValue())
		}
	}
	if e.UnderlyingSpotInActingVersion(actingVersion) {
		if e.UnderlyingSpot != e.UnderlyingSpotNullValue() && (e.UnderlyingSpot < e.UnderlyingSpotMinValue() || e.UnderlyingSpot > e.UnderlyingSpotMaxValue()) {
			return fmt.Errorf("Range check failed on e.UnderlyingSpot (%v < %v > %v)", e.UnderlyingSpotMinValue(), e.UnderlyingSpot, e.UnderlyingSpotMaxValue())
		}
	}
	if e.SettlementIntervalInActingVersion(actingVersion) {
		if e.SettlementInterval != e.SettlementIntervalNullValue() && (e.SettlementInterval < e.SettlementIntervalMinValue() || e.SettlementInterval > e.SettlementIntervalMaxValue()) {
			return fmt.Errorf("Range check failed on e.SettlementInterval (%v < %v > %v)", e.SettlementIntervalMinValue(), e.SettlementInterval, e.SettlementIntervalMaxValue())
		}
	}
	if e.MinQuantityInActingVersion(actingVersion) {
		if e.MinQuantity < e.MinQuantityMinValue() || e.MinQuantity > e.MinQuantityMaxValue() {
			return fmt.Errorf("Range check failed on e.MinQuantity (%v < %v > %v)", e.MinQuantityMinValue(), e.MinQuantity, e.MinQuantityMaxValue())
		}
	}
	if e.BaseIncrementInActingVersion(actingVersion) {
		if e.BaseIncrement < e.BaseIncrementMinValue() || e.BaseIncrement > e.BaseIncrementMaxValue() {
			return fmt.Errorf("Range check failed on e.BaseIncrement (%v < %v > %v)", e.BaseIncrementMinValue(), e.BaseIncrement, e.BaseIncrementMaxValue())
		}
	}
	if e.QuoteIncrementInActingVersion(actingVersion) {
		if e.QuoteIncrement < e.QuoteIncrementMinValue() || e.QuoteIncrement > e.QuoteIncrementMaxValue() {
			return fmt.Errorf("Range check failed on e.QuoteIncrement (%v < %v > %v)", e.QuoteIncrementMinValue(), e.QuoteIncrement, e.QuoteIncrementMaxValue())
		}
	}
	if e.InitialMarginFractionFactorInActingVersion(actingVersion) {
		if e.InitialMarginFractionFactor != e.InitialMarginFractionFactorNullValue() && (e.InitialMarginFractionFactor < e.InitialMarginFractionFactorMinValue() || e.InitialMarginFractionFactor > e.InitialMarginFractionFactorMaxValue()) {
			return fmt.Errorf("Range check failed on e.InitialMarginFractionFactor (%v < %v > %v)", e.InitialMarginFractionFactorMinValue(), e.InitialMarginFractionFactor, e.InitialMarginFractionFactorMaxValue())
		}
	}
	if e.MaintenanceMarginFractionFactorInActingVersion(actingVersion) {
		if e.MaintenanceMarginFractionFactor != e.MaintenanceMarginFractionFactorNullValue() && (e.MaintenanceMarginFractionFactor < e.MaintenanceMarginFractionFactorMinValue() || e.MaintenanceMarginFractionFactor > e.MaintenanceMarginFractionFactorMaxValue()) {
			return fmt.Errorf("Range check failed on e.MaintenanceMarginFractionFactor (%v < %v > %v)", e.MaintenanceMarginFractionFactorMinValue(), e.MaintenanceMarginFractionFactor, e.MaintenanceMarginFractionFactorMaxValue())
		}
	}
	if e.LastTradePriceInActingVersion(actingVersion) {
		if e.LastTradePrice < e.LastTradePriceMinValue() || e.LastTradePrice > e.LastTradePriceMaxValue() {
			return fmt.Errorf("Range check failed on e.LastTradePrice (%v < %v > %v)", e.LastTradePriceMinValue(), e.LastTradePrice, e.LastTradePriceMaxValue())
		}
	}
	if e.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			if e.Symbol[idx] < e.SymbolMinValue() || e.Symbol[idx] > e.SymbolMaxValue() {
				return fmt.Errorf("Range check failed on e.Symbol[%d] (%v < %v > %v)", idx, e.SymbolMinValue(), e.Symbol[idx], e.SymbolMaxValue())
			}
		}
	}
	for idx, ch := range e.Symbol {
		if ch > 127 {
			return fmt.Errorf("e.Symbol[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	if e.OrderQueueIdInActingVersion(actingVersion) {
		if e.OrderQueueId < e.OrderQueueIdMinValue() || e.OrderQueueId > e.OrderQueueIdMaxValue() {
			return fmt.Errorf("Range check failed on e.OrderQueueId (%v < %v > %v)", e.OrderQueueIdMinValue(), e.OrderQueueId, e.OrderQueueIdMaxValue())
		}
	}
	if e.NextMdInstrSeqNoInActingVersion(actingVersion) {
		if e.NextMdInstrSeqNo < e.NextMdInstrSeqNoMinValue() || e.NextMdInstrSeqNo > e.NextMdInstrSeqNoMaxValue() {
			return fmt.Errorf("Range check failed on e.NextMdInstrSeqNo (%v < %v > %v)", e.NextMdInstrSeqNoMinValue(), e.NextMdInstrSeqNo, e.NextMdInstrSeqNoMaxValue())
		}
	}
	if err := e.InstrumentType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func EngBookStateInit(e *EngBookState) {
	e.UnderlyingSpot = math.MinInt64
	e.SettlementInterval = math.MinInt64
	e.InitialMarginFractionFactor = math.MinInt64
	e.MaintenanceMarginFractionFactor = math.MinInt64
	return
}

func (*EngBookState) SbeBlockLength() (blockLength uint16) {
	return 201
}

func (*EngBookState) SbeTemplateId() (templateId uint16) {
	return 2102
}

func (*EngBookState) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*EngBookState) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*EngBookState) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*EngBookState) InstrumentIdId() uint16 {
	return 1
}

func (*EngBookState) InstrumentIdSinceVersion() uint16 {
	return 0
}

func (e *EngBookState) InstrumentIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.InstrumentIdSinceVersion()
}

func (*EngBookState) InstrumentIdDeprecated() uint16 {
	return 0
}

func (*EngBookState) InstrumentIdMetaAttribute(meta int) string {
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

func (*EngBookState) InstrumentIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngBookState) InstrumentIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngBookState) InstrumentIdNullValue() int64 {
	return math.MinInt64
}

func (*EngBookState) InstrumentUuidId() uint16 {
	return 2
}

func (*EngBookState) InstrumentUuidSinceVersion() uint16 {
	return 0
}

func (e *EngBookState) InstrumentUuidInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.InstrumentUuidSinceVersion()
}

func (*EngBookState) InstrumentUuidDeprecated() uint16 {
	return 0
}

func (*EngBookState) InstrumentUuidMetaAttribute(meta int) string {
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

func (*EngBookState) InstrumentUuidMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngBookState) InstrumentUuidMaxValue() int64 {
	return math.MaxInt64
}

func (*EngBookState) InstrumentUuidNullValue() int64 {
	return math.MinInt64
}

func (*EngBookState) BaseAssetId() uint16 {
	return 3
}

func (*EngBookState) BaseAssetSinceVersion() uint16 {
	return 0
}

func (e *EngBookState) BaseAssetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.BaseAssetSinceVersion()
}

func (*EngBookState) BaseAssetDeprecated() uint16 {
	return 0
}

func (*EngBookState) BaseAssetMetaAttribute(meta int) string {
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

func (*EngBookState) BaseAssetMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngBookState) BaseAssetMaxValue() int64 {
	return math.MaxInt64
}

func (*EngBookState) BaseAssetNullValue() int64 {
	return math.MinInt64
}

func (*EngBookState) QuoteAssetId() uint16 {
	return 4
}

func (*EngBookState) QuoteAssetSinceVersion() uint16 {
	return 0
}

func (e *EngBookState) QuoteAssetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.QuoteAssetSinceVersion()
}

func (*EngBookState) QuoteAssetDeprecated() uint16 {
	return 0
}

func (*EngBookState) QuoteAssetMetaAttribute(meta int) string {
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

func (*EngBookState) QuoteAssetMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngBookState) QuoteAssetMaxValue() int64 {
	return math.MaxInt64
}

func (*EngBookState) QuoteAssetNullValue() int64 {
	return math.MinInt64
}

func (*EngBookState) UnderlyingSpotId() uint16 {
	return 5
}

func (*EngBookState) UnderlyingSpotSinceVersion() uint16 {
	return 0
}

func (e *EngBookState) UnderlyingSpotInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.UnderlyingSpotSinceVersion()
}

func (*EngBookState) UnderlyingSpotDeprecated() uint16 {
	return 0
}

func (*EngBookState) UnderlyingSpotMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*EngBookState) UnderlyingSpotMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngBookState) UnderlyingSpotMaxValue() int64 {
	return math.MaxInt64
}

func (*EngBookState) UnderlyingSpotNullValue() int64 {
	return math.MinInt64
}

func (*EngBookState) SettlementIntervalId() uint16 {
	return 6
}

func (*EngBookState) SettlementIntervalSinceVersion() uint16 {
	return 0
}

func (e *EngBookState) SettlementIntervalInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SettlementIntervalSinceVersion()
}

func (*EngBookState) SettlementIntervalDeprecated() uint16 {
	return 0
}

func (*EngBookState) SettlementIntervalMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*EngBookState) SettlementIntervalMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngBookState) SettlementIntervalMaxValue() int64 {
	return math.MaxInt64
}

func (*EngBookState) SettlementIntervalNullValue() int64 {
	return math.MinInt64
}

func (*EngBookState) MinQuantityId() uint16 {
	return 7
}

func (*EngBookState) MinQuantitySinceVersion() uint16 {
	return 0
}

func (e *EngBookState) MinQuantityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MinQuantitySinceVersion()
}

func (*EngBookState) MinQuantityDeprecated() uint16 {
	return 0
}

func (*EngBookState) MinQuantityMetaAttribute(meta int) string {
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

func (*EngBookState) MinQuantityMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngBookState) MinQuantityMaxValue() int64 {
	return math.MaxInt64
}

func (*EngBookState) MinQuantityNullValue() int64 {
	return math.MinInt64
}

func (*EngBookState) BaseIncrementId() uint16 {
	return 8
}

func (*EngBookState) BaseIncrementSinceVersion() uint16 {
	return 0
}

func (e *EngBookState) BaseIncrementInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.BaseIncrementSinceVersion()
}

func (*EngBookState) BaseIncrementDeprecated() uint16 {
	return 0
}

func (*EngBookState) BaseIncrementMetaAttribute(meta int) string {
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

func (*EngBookState) BaseIncrementMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngBookState) BaseIncrementMaxValue() int64 {
	return math.MaxInt64
}

func (*EngBookState) BaseIncrementNullValue() int64 {
	return math.MinInt64
}

func (*EngBookState) QuoteIncrementId() uint16 {
	return 9
}

func (*EngBookState) QuoteIncrementSinceVersion() uint16 {
	return 0
}

func (e *EngBookState) QuoteIncrementInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.QuoteIncrementSinceVersion()
}

func (*EngBookState) QuoteIncrementDeprecated() uint16 {
	return 0
}

func (*EngBookState) QuoteIncrementMetaAttribute(meta int) string {
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

func (*EngBookState) QuoteIncrementMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngBookState) QuoteIncrementMaxValue() int64 {
	return math.MaxInt64
}

func (*EngBookState) QuoteIncrementNullValue() int64 {
	return math.MinInt64
}

func (*EngBookState) InitialMarginFractionFactorId() uint16 {
	return 10
}

func (*EngBookState) InitialMarginFractionFactorSinceVersion() uint16 {
	return 0
}

func (e *EngBookState) InitialMarginFractionFactorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.InitialMarginFractionFactorSinceVersion()
}

func (*EngBookState) InitialMarginFractionFactorDeprecated() uint16 {
	return 0
}

func (*EngBookState) InitialMarginFractionFactorMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*EngBookState) InitialMarginFractionFactorMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngBookState) InitialMarginFractionFactorMaxValue() int64 {
	return math.MaxInt64
}

func (*EngBookState) InitialMarginFractionFactorNullValue() int64 {
	return math.MinInt64
}

func (*EngBookState) MaintenanceMarginFractionFactorId() uint16 {
	return 11
}

func (*EngBookState) MaintenanceMarginFractionFactorSinceVersion() uint16 {
	return 0
}

func (e *EngBookState) MaintenanceMarginFractionFactorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MaintenanceMarginFractionFactorSinceVersion()
}

func (*EngBookState) MaintenanceMarginFractionFactorDeprecated() uint16 {
	return 0
}

func (*EngBookState) MaintenanceMarginFractionFactorMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*EngBookState) MaintenanceMarginFractionFactorMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngBookState) MaintenanceMarginFractionFactorMaxValue() int64 {
	return math.MaxInt64
}

func (*EngBookState) MaintenanceMarginFractionFactorNullValue() int64 {
	return math.MinInt64
}

func (*EngBookState) LastTradePriceId() uint16 {
	return 12
}

func (*EngBookState) LastTradePriceSinceVersion() uint16 {
	return 0
}

func (e *EngBookState) LastTradePriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LastTradePriceSinceVersion()
}

func (*EngBookState) LastTradePriceDeprecated() uint16 {
	return 0
}

func (*EngBookState) LastTradePriceMetaAttribute(meta int) string {
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

func (*EngBookState) LastTradePriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngBookState) LastTradePriceMaxValue() int64 {
	return math.MaxInt64
}

func (*EngBookState) LastTradePriceNullValue() int64 {
	return math.MinInt64
}

func (*EngBookState) SymbolId() uint16 {
	return 13
}

func (*EngBookState) SymbolSinceVersion() uint16 {
	return 0
}

func (e *EngBookState) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SymbolSinceVersion()
}

func (*EngBookState) SymbolDeprecated() uint16 {
	return 0
}

func (*EngBookState) SymbolMetaAttribute(meta int) string {
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

func (*EngBookState) SymbolMinValue() byte {
	return byte(32)
}

func (*EngBookState) SymbolMaxValue() byte {
	return byte(126)
}

func (*EngBookState) SymbolNullValue() byte {
	return 0
}

func (e *EngBookState) SymbolCharacterEncoding() string {
	return "US-ASCII"
}

func (*EngBookState) OrderQueueIdId() uint16 {
	return 14
}

func (*EngBookState) OrderQueueIdSinceVersion() uint16 {
	return 0
}

func (e *EngBookState) OrderQueueIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderQueueIdSinceVersion()
}

func (*EngBookState) OrderQueueIdDeprecated() uint16 {
	return 0
}

func (*EngBookState) OrderQueueIdMetaAttribute(meta int) string {
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

func (*EngBookState) OrderQueueIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngBookState) OrderQueueIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngBookState) OrderQueueIdNullValue() int64 {
	return math.MinInt64
}

func (*EngBookState) NextMdInstrSeqNoId() uint16 {
	return 15
}

func (*EngBookState) NextMdInstrSeqNoSinceVersion() uint16 {
	return 0
}

func (e *EngBookState) NextMdInstrSeqNoInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NextMdInstrSeqNoSinceVersion()
}

func (*EngBookState) NextMdInstrSeqNoDeprecated() uint16 {
	return 0
}

func (*EngBookState) NextMdInstrSeqNoMetaAttribute(meta int) string {
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

func (*EngBookState) NextMdInstrSeqNoMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngBookState) NextMdInstrSeqNoMaxValue() int64 {
	return math.MaxInt64
}

func (*EngBookState) NextMdInstrSeqNoNullValue() int64 {
	return math.MinInt64
}

func (*EngBookState) InstrumentTypeId() uint16 {
	return 16
}

func (*EngBookState) InstrumentTypeSinceVersion() uint16 {
	return 0
}

func (e *EngBookState) InstrumentTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.InstrumentTypeSinceVersion()
}

func (*EngBookState) InstrumentTypeDeprecated() uint16 {
	return 0
}

func (*EngBookState) InstrumentTypeMetaAttribute(meta int) string {
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
