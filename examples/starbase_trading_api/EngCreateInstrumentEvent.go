// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type EngCreateInstrumentEvent struct {
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
	Symbol                          [80]byte
	InstrumentType                  InstrumentTypeEnum
}

func (e *EngCreateInstrumentEvent) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	if err := _m.WriteBytes(_w, e.Symbol[:]); err != nil {
		return err
	}
	if err := e.InstrumentType.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (e *EngCreateInstrumentEvent) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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

func (e *EngCreateInstrumentEvent) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func EngCreateInstrumentEventInit(e *EngCreateInstrumentEvent) {
	e.UnderlyingSpot = math.MinInt64
	e.SettlementInterval = math.MinInt64
	e.InitialMarginFractionFactor = math.MinInt64
	e.MaintenanceMarginFractionFactor = math.MinInt64
	return
}

func (*EngCreateInstrumentEvent) SbeBlockLength() (blockLength uint16) {
	return 177
}

func (*EngCreateInstrumentEvent) SbeTemplateId() (templateId uint16) {
	return 2012
}

func (*EngCreateInstrumentEvent) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*EngCreateInstrumentEvent) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*EngCreateInstrumentEvent) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*EngCreateInstrumentEvent) InstrumentIdId() uint16 {
	return 1
}

func (*EngCreateInstrumentEvent) InstrumentIdSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrumentEvent) InstrumentIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.InstrumentIdSinceVersion()
}

func (*EngCreateInstrumentEvent) InstrumentIdDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrumentEvent) InstrumentIdMetaAttribute(meta int) string {
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

func (*EngCreateInstrumentEvent) InstrumentIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCreateInstrumentEvent) InstrumentIdMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCreateInstrumentEvent) InstrumentIdNullValue() int64 {
	return math.MinInt64
}

func (*EngCreateInstrumentEvent) InstrumentUuidId() uint16 {
	return 2
}

func (*EngCreateInstrumentEvent) InstrumentUuidSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrumentEvent) InstrumentUuidInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.InstrumentUuidSinceVersion()
}

func (*EngCreateInstrumentEvent) InstrumentUuidDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrumentEvent) InstrumentUuidMetaAttribute(meta int) string {
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

func (*EngCreateInstrumentEvent) InstrumentUuidMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCreateInstrumentEvent) InstrumentUuidMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCreateInstrumentEvent) InstrumentUuidNullValue() int64 {
	return math.MinInt64
}

func (*EngCreateInstrumentEvent) BaseAssetId() uint16 {
	return 3
}

func (*EngCreateInstrumentEvent) BaseAssetSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrumentEvent) BaseAssetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.BaseAssetSinceVersion()
}

func (*EngCreateInstrumentEvent) BaseAssetDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrumentEvent) BaseAssetMetaAttribute(meta int) string {
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

func (*EngCreateInstrumentEvent) BaseAssetMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCreateInstrumentEvent) BaseAssetMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCreateInstrumentEvent) BaseAssetNullValue() int64 {
	return math.MinInt64
}

func (*EngCreateInstrumentEvent) QuoteAssetId() uint16 {
	return 4
}

func (*EngCreateInstrumentEvent) QuoteAssetSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrumentEvent) QuoteAssetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.QuoteAssetSinceVersion()
}

func (*EngCreateInstrumentEvent) QuoteAssetDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrumentEvent) QuoteAssetMetaAttribute(meta int) string {
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

func (*EngCreateInstrumentEvent) QuoteAssetMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCreateInstrumentEvent) QuoteAssetMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCreateInstrumentEvent) QuoteAssetNullValue() int64 {
	return math.MinInt64
}

func (*EngCreateInstrumentEvent) UnderlyingSpotId() uint16 {
	return 5
}

func (*EngCreateInstrumentEvent) UnderlyingSpotSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrumentEvent) UnderlyingSpotInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.UnderlyingSpotSinceVersion()
}

func (*EngCreateInstrumentEvent) UnderlyingSpotDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrumentEvent) UnderlyingSpotMetaAttribute(meta int) string {
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

func (*EngCreateInstrumentEvent) UnderlyingSpotMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCreateInstrumentEvent) UnderlyingSpotMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCreateInstrumentEvent) UnderlyingSpotNullValue() int64 {
	return math.MinInt64
}

func (*EngCreateInstrumentEvent) SettlementIntervalId() uint16 {
	return 6
}

func (*EngCreateInstrumentEvent) SettlementIntervalSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrumentEvent) SettlementIntervalInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SettlementIntervalSinceVersion()
}

func (*EngCreateInstrumentEvent) SettlementIntervalDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrumentEvent) SettlementIntervalMetaAttribute(meta int) string {
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

func (*EngCreateInstrumentEvent) SettlementIntervalMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCreateInstrumentEvent) SettlementIntervalMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCreateInstrumentEvent) SettlementIntervalNullValue() int64 {
	return math.MinInt64
}

func (*EngCreateInstrumentEvent) MinQuantityId() uint16 {
	return 7
}

func (*EngCreateInstrumentEvent) MinQuantitySinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrumentEvent) MinQuantityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MinQuantitySinceVersion()
}

func (*EngCreateInstrumentEvent) MinQuantityDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrumentEvent) MinQuantityMetaAttribute(meta int) string {
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

func (*EngCreateInstrumentEvent) MinQuantityMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCreateInstrumentEvent) MinQuantityMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCreateInstrumentEvent) MinQuantityNullValue() int64 {
	return math.MinInt64
}

func (*EngCreateInstrumentEvent) BaseIncrementId() uint16 {
	return 8
}

func (*EngCreateInstrumentEvent) BaseIncrementSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrumentEvent) BaseIncrementInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.BaseIncrementSinceVersion()
}

func (*EngCreateInstrumentEvent) BaseIncrementDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrumentEvent) BaseIncrementMetaAttribute(meta int) string {
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

func (*EngCreateInstrumentEvent) BaseIncrementMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCreateInstrumentEvent) BaseIncrementMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCreateInstrumentEvent) BaseIncrementNullValue() int64 {
	return math.MinInt64
}

func (*EngCreateInstrumentEvent) QuoteIncrementId() uint16 {
	return 9
}

func (*EngCreateInstrumentEvent) QuoteIncrementSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrumentEvent) QuoteIncrementInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.QuoteIncrementSinceVersion()
}

func (*EngCreateInstrumentEvent) QuoteIncrementDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrumentEvent) QuoteIncrementMetaAttribute(meta int) string {
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

func (*EngCreateInstrumentEvent) QuoteIncrementMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCreateInstrumentEvent) QuoteIncrementMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCreateInstrumentEvent) QuoteIncrementNullValue() int64 {
	return math.MinInt64
}

func (*EngCreateInstrumentEvent) InitialMarginFractionFactorId() uint16 {
	return 10
}

func (*EngCreateInstrumentEvent) InitialMarginFractionFactorSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrumentEvent) InitialMarginFractionFactorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.InitialMarginFractionFactorSinceVersion()
}

func (*EngCreateInstrumentEvent) InitialMarginFractionFactorDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrumentEvent) InitialMarginFractionFactorMetaAttribute(meta int) string {
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

func (*EngCreateInstrumentEvent) InitialMarginFractionFactorMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCreateInstrumentEvent) InitialMarginFractionFactorMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCreateInstrumentEvent) InitialMarginFractionFactorNullValue() int64 {
	return math.MinInt64
}

func (*EngCreateInstrumentEvent) MaintenanceMarginFractionFactorId() uint16 {
	return 11
}

func (*EngCreateInstrumentEvent) MaintenanceMarginFractionFactorSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrumentEvent) MaintenanceMarginFractionFactorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MaintenanceMarginFractionFactorSinceVersion()
}

func (*EngCreateInstrumentEvent) MaintenanceMarginFractionFactorDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrumentEvent) MaintenanceMarginFractionFactorMetaAttribute(meta int) string {
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

func (*EngCreateInstrumentEvent) MaintenanceMarginFractionFactorMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EngCreateInstrumentEvent) MaintenanceMarginFractionFactorMaxValue() int64 {
	return math.MaxInt64
}

func (*EngCreateInstrumentEvent) MaintenanceMarginFractionFactorNullValue() int64 {
	return math.MinInt64
}

func (*EngCreateInstrumentEvent) SymbolId() uint16 {
	return 12
}

func (*EngCreateInstrumentEvent) SymbolSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrumentEvent) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SymbolSinceVersion()
}

func (*EngCreateInstrumentEvent) SymbolDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrumentEvent) SymbolMetaAttribute(meta int) string {
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

func (*EngCreateInstrumentEvent) SymbolMinValue() byte {
	return byte(32)
}

func (*EngCreateInstrumentEvent) SymbolMaxValue() byte {
	return byte(126)
}

func (*EngCreateInstrumentEvent) SymbolNullValue() byte {
	return 0
}

func (e *EngCreateInstrumentEvent) SymbolCharacterEncoding() string {
	return "US-ASCII"
}

func (*EngCreateInstrumentEvent) InstrumentTypeId() uint16 {
	return 13
}

func (*EngCreateInstrumentEvent) InstrumentTypeSinceVersion() uint16 {
	return 0
}

func (e *EngCreateInstrumentEvent) InstrumentTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.InstrumentTypeSinceVersion()
}

func (*EngCreateInstrumentEvent) InstrumentTypeDeprecated() uint16 {
	return 0
}

func (*EngCreateInstrumentEvent) InstrumentTypeMetaAttribute(meta int) string {
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
