// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type EngCreateInstrument struct {
	CorrelationId                   int64
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
	Symbol                          [80]byte
	InstrumentType                  InstrumentTypeEnum
}

func (e *EngCreateInstrument) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := e.RangeCheck(e.SbeSchemaVersion(), e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, e.CorrelationId); err != nil {
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
	if err := _m.WriteBytes(_w, e.Symbol[:]); err != nil {
		return err
	}
	if err := e.InstrumentType.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (e *EngCreateInstrument) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !e.CorrelationIdInActingVersion(actingVersion) {
		e.CorrelationId = e.CorrelationIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.CorrelationId); err != nil {
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
	if !e.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			e.Symbol[idx] = e.SymbolNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, e.Symbol[:]); err != nil {
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

func (e *EngCreateInstrument) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.CorrelationIdInActingVersion(actingVersion) {
		if e.CorrelationId < e.CorrelationIdMinValue() || e.CorrelationId > e.CorrelationIdMaxValue() {
			return fmt.Errorf("Range check failed on e.CorrelationId (%v < %v > %v)", e.CorrelationIdMinValue(), e.CorrelationId, e.CorrelationIdMaxValue())
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
	if err := e.InstrumentType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func EngCreateInstrumentInit(e *EngCreateInstrument) {
	e.UnderlyingSpot = math.MinInt64
	e.SettlementInterval = math.MinInt64
	e.InitialMarginFractionFactor = math.MinInt64
	e.MaintenanceMarginFractionFactor = math.MinInt64
	return
}

func (*EngCreateInstrument) SbeBlockLength() (blockLength uint16) {
	return 177
}

func (*EngCreateInstrument) SbeTemplateId() (templateId uint16) {
	return 2011
}

func (*EngCreateInstrument) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*EngCreateInstrument) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*EngCreateInstrument) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*EngCreateInstrument) CorrelationIdId() uint16 {
	return 1
}

func (*EngCreateInstrument) CorrelationIdSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrument) CorrelationIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CorrelationIdSinceVersion()
}

func (*EngCreateInstrument) CorrelationIdDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrument) CorrelationIdMetaAttribute(meta int) string {
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

func (*EngCreateInstrument) CorrelationIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCreateInstrument) CorrelationIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCreateInstrument) CorrelationIdNullValue() int64 {
	return math.MinInt64
}

func (*EngCreateInstrument) InstrumentUuidId() uint16 {
	return 2
}

func (*EngCreateInstrument) InstrumentUuidSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrument) InstrumentUuidInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.InstrumentUuidSinceVersion()
}

func (*EngCreateInstrument) InstrumentUuidDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrument) InstrumentUuidMetaAttribute(meta int) string {
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

func (*EngCreateInstrument) InstrumentUuidMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCreateInstrument) InstrumentUuidMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCreateInstrument) InstrumentUuidNullValue() int64 {
	return math.MinInt64
}

func (*EngCreateInstrument) BaseAssetId() uint16 {
	return 3
}

func (*EngCreateInstrument) BaseAssetSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrument) BaseAssetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.BaseAssetSinceVersion()
}

func (*EngCreateInstrument) BaseAssetDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrument) BaseAssetMetaAttribute(meta int) string {
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

func (*EngCreateInstrument) BaseAssetMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCreateInstrument) BaseAssetMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCreateInstrument) BaseAssetNullValue() int64 {
	return math.MinInt64
}

func (*EngCreateInstrument) QuoteAssetId() uint16 {
	return 4
}

func (*EngCreateInstrument) QuoteAssetSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrument) QuoteAssetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.QuoteAssetSinceVersion()
}

func (*EngCreateInstrument) QuoteAssetDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrument) QuoteAssetMetaAttribute(meta int) string {
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

func (*EngCreateInstrument) QuoteAssetMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCreateInstrument) QuoteAssetMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCreateInstrument) QuoteAssetNullValue() int64 {
	return math.MinInt64
}

func (*EngCreateInstrument) UnderlyingSpotId() uint16 {
	return 5
}

func (*EngCreateInstrument) UnderlyingSpotSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrument) UnderlyingSpotInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.UnderlyingSpotSinceVersion()
}

func (*EngCreateInstrument) UnderlyingSpotDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrument) UnderlyingSpotMetaAttribute(meta int) string {
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

func (*EngCreateInstrument) UnderlyingSpotMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCreateInstrument) UnderlyingSpotMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCreateInstrument) UnderlyingSpotNullValue() int64 {
	return math.MinInt64
}

func (*EngCreateInstrument) SettlementIntervalId() uint16 {
	return 6
}

func (*EngCreateInstrument) SettlementIntervalSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrument) SettlementIntervalInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SettlementIntervalSinceVersion()
}

func (*EngCreateInstrument) SettlementIntervalDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrument) SettlementIntervalMetaAttribute(meta int) string {
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

func (*EngCreateInstrument) SettlementIntervalMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCreateInstrument) SettlementIntervalMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCreateInstrument) SettlementIntervalNullValue() int64 {
	return math.MinInt64
}

func (*EngCreateInstrument) MinQuantityId() uint16 {
	return 7
}

func (*EngCreateInstrument) MinQuantitySinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrument) MinQuantityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MinQuantitySinceVersion()
}

func (*EngCreateInstrument) MinQuantityDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrument) MinQuantityMetaAttribute(meta int) string {
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

func (*EngCreateInstrument) MinQuantityMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCreateInstrument) MinQuantityMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCreateInstrument) MinQuantityNullValue() int64 {
	return math.MinInt64
}

func (*EngCreateInstrument) BaseIncrementId() uint16 {
	return 8
}

func (*EngCreateInstrument) BaseIncrementSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrument) BaseIncrementInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.BaseIncrementSinceVersion()
}

func (*EngCreateInstrument) BaseIncrementDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrument) BaseIncrementMetaAttribute(meta int) string {
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

func (*EngCreateInstrument) BaseIncrementMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCreateInstrument) BaseIncrementMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCreateInstrument) BaseIncrementNullValue() int64 {
	return math.MinInt64
}

func (*EngCreateInstrument) QuoteIncrementId() uint16 {
	return 9
}

func (*EngCreateInstrument) QuoteIncrementSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrument) QuoteIncrementInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.QuoteIncrementSinceVersion()
}

func (*EngCreateInstrument) QuoteIncrementDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrument) QuoteIncrementMetaAttribute(meta int) string {
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

func (*EngCreateInstrument) QuoteIncrementMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCreateInstrument) QuoteIncrementMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCreateInstrument) QuoteIncrementNullValue() int64 {
	return math.MinInt64
}

func (*EngCreateInstrument) InitialMarginFractionFactorId() uint16 {
	return 10
}

func (*EngCreateInstrument) InitialMarginFractionFactorSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrument) InitialMarginFractionFactorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.InitialMarginFractionFactorSinceVersion()
}

func (*EngCreateInstrument) InitialMarginFractionFactorDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrument) InitialMarginFractionFactorMetaAttribute(meta int) string {
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

func (*EngCreateInstrument) InitialMarginFractionFactorMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCreateInstrument) InitialMarginFractionFactorMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCreateInstrument) InitialMarginFractionFactorNullValue() int64 {
	return math.MinInt64
}

func (*EngCreateInstrument) MaintenanceMarginFractionFactorId() uint16 {
	return 11
}

func (*EngCreateInstrument) MaintenanceMarginFractionFactorSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrument) MaintenanceMarginFractionFactorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MaintenanceMarginFractionFactorSinceVersion()
}

func (*EngCreateInstrument) MaintenanceMarginFractionFactorDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrument) MaintenanceMarginFractionFactorMetaAttribute(meta int) string {
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

func (*EngCreateInstrument) MaintenanceMarginFractionFactorMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCreateInstrument) MaintenanceMarginFractionFactorMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCreateInstrument) MaintenanceMarginFractionFactorNullValue() int64 {
	return math.MinInt64
}

func (*EngCreateInstrument) SymbolId() uint16 {
	return 12
}

func (*EngCreateInstrument) SymbolSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrument) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SymbolSinceVersion()
}

func (*EngCreateInstrument) SymbolDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrument) SymbolMetaAttribute(meta int) string {
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

func (*EngCreateInstrument) SymbolMinValue() byte {
	return byte(32)
}

func (*EngCreateInstrument) SymbolMaxValue() byte {
	return byte(126)
}

func (*EngCreateInstrument) SymbolNullValue() byte {
	return 0
}

func (e *EngCreateInstrument) SymbolCharacterEncoding() string {
	return "US-ASCII"
}

func (*EngCreateInstrument) InstrumentTypeId() uint16 {
	return 13
}

func (*EngCreateInstrument) InstrumentTypeSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrument) InstrumentTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.InstrumentTypeSinceVersion()
}

func (*EngCreateInstrument) InstrumentTypeDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrument) InstrumentTypeMetaAttribute(meta int) string {
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
