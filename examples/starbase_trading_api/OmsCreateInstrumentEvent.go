// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OmsCreateInstrumentEvent struct {
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

func (o *OmsCreateInstrumentEvent) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, o.InstrumentId); err != nil {
		return err
	}
	for idx := 0; idx < 2; idx++ {
		if err := _m.WriteInt64(_w, o.InstrumentUuid[idx]); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, o.BaseAsset); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.QuoteAsset); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.UnderlyingSpot); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.SettlementInterval); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.MinQuantity); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.BaseIncrement); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.QuoteIncrement); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.InitialMarginFractionFactor); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.MaintenanceMarginFractionFactor); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.Symbol[:]); err != nil {
		return err
	}
	if err := o.InstrumentType.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (o *OmsCreateInstrumentEvent) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !o.InstrumentIdInActingVersion(actingVersion) {
		o.InstrumentId = o.InstrumentIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.InstrumentId); err != nil {
			return err
		}
	}
	if !o.InstrumentUuidInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			o.InstrumentUuid[idx] = o.InstrumentUuidNullValue()
		}
	} else {
		for idx := 0; idx < 2; idx++ {
			if err := _m.ReadInt64(_r, &o.InstrumentUuid[idx]); err != nil {
				return err
			}
		}
	}
	if !o.BaseAssetInActingVersion(actingVersion) {
		o.BaseAsset = o.BaseAssetNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.BaseAsset); err != nil {
			return err
		}
	}
	if !o.QuoteAssetInActingVersion(actingVersion) {
		o.QuoteAsset = o.QuoteAssetNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.QuoteAsset); err != nil {
			return err
		}
	}
	if !o.UnderlyingSpotInActingVersion(actingVersion) {
		o.UnderlyingSpot = o.UnderlyingSpotNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.UnderlyingSpot); err != nil {
			return err
		}
	}
	if !o.SettlementIntervalInActingVersion(actingVersion) {
		o.SettlementInterval = o.SettlementIntervalNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.SettlementInterval); err != nil {
			return err
		}
	}
	if !o.MinQuantityInActingVersion(actingVersion) {
		o.MinQuantity = o.MinQuantityNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.MinQuantity); err != nil {
			return err
		}
	}
	if !o.BaseIncrementInActingVersion(actingVersion) {
		o.BaseIncrement = o.BaseIncrementNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.BaseIncrement); err != nil {
			return err
		}
	}
	if !o.QuoteIncrementInActingVersion(actingVersion) {
		o.QuoteIncrement = o.QuoteIncrementNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.QuoteIncrement); err != nil {
			return err
		}
	}
	if !o.InitialMarginFractionFactorInActingVersion(actingVersion) {
		o.InitialMarginFractionFactor = o.InitialMarginFractionFactorNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.InitialMarginFractionFactor); err != nil {
			return err
		}
	}
	if !o.MaintenanceMarginFractionFactorInActingVersion(actingVersion) {
		o.MaintenanceMarginFractionFactor = o.MaintenanceMarginFractionFactorNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.MaintenanceMarginFractionFactor); err != nil {
			return err
		}
	}
	if !o.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			o.Symbol[idx] = o.SymbolNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.Symbol[:]); err != nil {
			return err
		}
	}
	if o.InstrumentTypeInActingVersion(actingVersion) {
		if err := o.InstrumentType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > o.SbeSchemaVersion() && blockLength > o.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-o.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := o.RangeCheck(actingVersion, o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (o *OmsCreateInstrumentEvent) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.InstrumentIdInActingVersion(actingVersion) {
		if o.InstrumentId < o.InstrumentIdMinValue() || o.InstrumentId > o.InstrumentIdMaxValue() {
			return fmt.Errorf("Range check failed on o.InstrumentId (%v < %v > %v)", o.InstrumentIdMinValue(), o.InstrumentId, o.InstrumentIdMaxValue())
		}
	}
	if o.InstrumentUuidInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if o.InstrumentUuid[idx] < o.InstrumentUuidMinValue() || o.InstrumentUuid[idx] > o.InstrumentUuidMaxValue() {
				return fmt.Errorf("Range check failed on o.InstrumentUuid[%d] (%v < %v > %v)", idx, o.InstrumentUuidMinValue(), o.InstrumentUuid[idx], o.InstrumentUuidMaxValue())
			}
		}
	}
	if o.BaseAssetInActingVersion(actingVersion) {
		if o.BaseAsset < o.BaseAssetMinValue() || o.BaseAsset > o.BaseAssetMaxValue() {
			return fmt.Errorf("Range check failed on o.BaseAsset (%v < %v > %v)", o.BaseAssetMinValue(), o.BaseAsset, o.BaseAssetMaxValue())
		}
	}
	if o.QuoteAssetInActingVersion(actingVersion) {
		if o.QuoteAsset < o.QuoteAssetMinValue() || o.QuoteAsset > o.QuoteAssetMaxValue() {
			return fmt.Errorf("Range check failed on o.QuoteAsset (%v < %v > %v)", o.QuoteAssetMinValue(), o.QuoteAsset, o.QuoteAssetMaxValue())
		}
	}
	if o.UnderlyingSpotInActingVersion(actingVersion) {
		if o.UnderlyingSpot != o.UnderlyingSpotNullValue() && (o.UnderlyingSpot < o.UnderlyingSpotMinValue() || o.UnderlyingSpot > o.UnderlyingSpotMaxValue()) {
			return fmt.Errorf("Range check failed on o.UnderlyingSpot (%v < %v > %v)", o.UnderlyingSpotMinValue(), o.UnderlyingSpot, o.UnderlyingSpotMaxValue())
		}
	}
	if o.SettlementIntervalInActingVersion(actingVersion) {
		if o.SettlementInterval != o.SettlementIntervalNullValue() && (o.SettlementInterval < o.SettlementIntervalMinValue() || o.SettlementInterval > o.SettlementIntervalMaxValue()) {
			return fmt.Errorf("Range check failed on o.SettlementInterval (%v < %v > %v)", o.SettlementIntervalMinValue(), o.SettlementInterval, o.SettlementIntervalMaxValue())
		}
	}
	if o.MinQuantityInActingVersion(actingVersion) {
		if o.MinQuantity < o.MinQuantityMinValue() || o.MinQuantity > o.MinQuantityMaxValue() {
			return fmt.Errorf("Range check failed on o.MinQuantity (%v < %v > %v)", o.MinQuantityMinValue(), o.MinQuantity, o.MinQuantityMaxValue())
		}
	}
	if o.BaseIncrementInActingVersion(actingVersion) {
		if o.BaseIncrement < o.BaseIncrementMinValue() || o.BaseIncrement > o.BaseIncrementMaxValue() {
			return fmt.Errorf("Range check failed on o.BaseIncrement (%v < %v > %v)", o.BaseIncrementMinValue(), o.BaseIncrement, o.BaseIncrementMaxValue())
		}
	}
	if o.QuoteIncrementInActingVersion(actingVersion) {
		if o.QuoteIncrement < o.QuoteIncrementMinValue() || o.QuoteIncrement > o.QuoteIncrementMaxValue() {
			return fmt.Errorf("Range check failed on o.QuoteIncrement (%v < %v > %v)", o.QuoteIncrementMinValue(), o.QuoteIncrement, o.QuoteIncrementMaxValue())
		}
	}
	if o.InitialMarginFractionFactorInActingVersion(actingVersion) {
		if o.InitialMarginFractionFactor != o.InitialMarginFractionFactorNullValue() && (o.InitialMarginFractionFactor < o.InitialMarginFractionFactorMinValue() || o.InitialMarginFractionFactor > o.InitialMarginFractionFactorMaxValue()) {
			return fmt.Errorf("Range check failed on o.InitialMarginFractionFactor (%v < %v > %v)", o.InitialMarginFractionFactorMinValue(), o.InitialMarginFractionFactor, o.InitialMarginFractionFactorMaxValue())
		}
	}
	if o.MaintenanceMarginFractionFactorInActingVersion(actingVersion) {
		if o.MaintenanceMarginFractionFactor != o.MaintenanceMarginFractionFactorNullValue() && (o.MaintenanceMarginFractionFactor < o.MaintenanceMarginFractionFactorMinValue() || o.MaintenanceMarginFractionFactor > o.MaintenanceMarginFractionFactorMaxValue()) {
			return fmt.Errorf("Range check failed on o.MaintenanceMarginFractionFactor (%v < %v > %v)", o.MaintenanceMarginFractionFactorMinValue(), o.MaintenanceMarginFractionFactor, o.MaintenanceMarginFractionFactorMaxValue())
		}
	}
	if o.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			if o.Symbol[idx] < o.SymbolMinValue() || o.Symbol[idx] > o.SymbolMaxValue() {
				return fmt.Errorf("Range check failed on o.Symbol[%d] (%v < %v > %v)", idx, o.SymbolMinValue(), o.Symbol[idx], o.SymbolMaxValue())
			}
		}
	}
	for idx, ch := range o.Symbol {
		if ch > 127 {
			return fmt.Errorf("o.Symbol[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	if err := o.InstrumentType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func OmsCreateInstrumentEventInit(o *OmsCreateInstrumentEvent) {
	o.UnderlyingSpot = math.MinInt64
	o.SettlementInterval = math.MinInt64
	o.InitialMarginFractionFactor = math.MinInt64
	o.MaintenanceMarginFractionFactor = math.MinInt64
	return
}

func (*OmsCreateInstrumentEvent) SbeBlockLength() (blockLength uint16) {
	return 177
}

func (*OmsCreateInstrumentEvent) SbeTemplateId() (templateId uint16) {
	return 1102
}

func (*OmsCreateInstrumentEvent) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*OmsCreateInstrumentEvent) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*OmsCreateInstrumentEvent) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*OmsCreateInstrumentEvent) InstrumentIdId() uint16 {
	return 1
}

func (*OmsCreateInstrumentEvent) InstrumentIdSinceVersion() uint16 {
	return 0
}

func (o *OmsCreateInstrumentEvent) InstrumentIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.InstrumentIdSinceVersion()
}

func (*OmsCreateInstrumentEvent) InstrumentIdDeprecated() uint16 {
	return 0
}

func (*OmsCreateInstrumentEvent) InstrumentIdMetaAttribute(meta int) string {
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

func (*OmsCreateInstrumentEvent) InstrumentIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsCreateInstrumentEvent) InstrumentIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsCreateInstrumentEvent) InstrumentIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsCreateInstrumentEvent) InstrumentUuidId() uint16 {
	return 2
}

func (*OmsCreateInstrumentEvent) InstrumentUuidSinceVersion() uint16 {
	return 0
}

func (o *OmsCreateInstrumentEvent) InstrumentUuidInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.InstrumentUuidSinceVersion()
}

func (*OmsCreateInstrumentEvent) InstrumentUuidDeprecated() uint16 {
	return 0
}

func (*OmsCreateInstrumentEvent) InstrumentUuidMetaAttribute(meta int) string {
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

func (*OmsCreateInstrumentEvent) InstrumentUuidMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsCreateInstrumentEvent) InstrumentUuidMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsCreateInstrumentEvent) InstrumentUuidNullValue() int64 {
	return math.MinInt64
}

func (*OmsCreateInstrumentEvent) BaseAssetId() uint16 {
	return 3
}

func (*OmsCreateInstrumentEvent) BaseAssetSinceVersion() uint16 {
	return 0
}

func (o *OmsCreateInstrumentEvent) BaseAssetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.BaseAssetSinceVersion()
}

func (*OmsCreateInstrumentEvent) BaseAssetDeprecated() uint16 {
	return 0
}

func (*OmsCreateInstrumentEvent) BaseAssetMetaAttribute(meta int) string {
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

func (*OmsCreateInstrumentEvent) BaseAssetMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsCreateInstrumentEvent) BaseAssetMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsCreateInstrumentEvent) BaseAssetNullValue() int64 {
	return math.MinInt64
}

func (*OmsCreateInstrumentEvent) QuoteAssetId() uint16 {
	return 4
}

func (*OmsCreateInstrumentEvent) QuoteAssetSinceVersion() uint16 {
	return 0
}

func (o *OmsCreateInstrumentEvent) QuoteAssetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.QuoteAssetSinceVersion()
}

func (*OmsCreateInstrumentEvent) QuoteAssetDeprecated() uint16 {
	return 0
}

func (*OmsCreateInstrumentEvent) QuoteAssetMetaAttribute(meta int) string {
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

func (*OmsCreateInstrumentEvent) QuoteAssetMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsCreateInstrumentEvent) QuoteAssetMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsCreateInstrumentEvent) QuoteAssetNullValue() int64 {
	return math.MinInt64
}

func (*OmsCreateInstrumentEvent) UnderlyingSpotId() uint16 {
	return 5
}

func (*OmsCreateInstrumentEvent) UnderlyingSpotSinceVersion() uint16 {
	return 0
}

func (o *OmsCreateInstrumentEvent) UnderlyingSpotInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.UnderlyingSpotSinceVersion()
}

func (*OmsCreateInstrumentEvent) UnderlyingSpotDeprecated() uint16 {
	return 0
}

func (*OmsCreateInstrumentEvent) UnderlyingSpotMetaAttribute(meta int) string {
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

func (*OmsCreateInstrumentEvent) UnderlyingSpotMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsCreateInstrumentEvent) UnderlyingSpotMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsCreateInstrumentEvent) UnderlyingSpotNullValue() int64 {
	return math.MinInt64
}

func (*OmsCreateInstrumentEvent) SettlementIntervalId() uint16 {
	return 6
}

func (*OmsCreateInstrumentEvent) SettlementIntervalSinceVersion() uint16 {
	return 0
}

func (o *OmsCreateInstrumentEvent) SettlementIntervalInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SettlementIntervalSinceVersion()
}

func (*OmsCreateInstrumentEvent) SettlementIntervalDeprecated() uint16 {
	return 0
}

func (*OmsCreateInstrumentEvent) SettlementIntervalMetaAttribute(meta int) string {
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

func (*OmsCreateInstrumentEvent) SettlementIntervalMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsCreateInstrumentEvent) SettlementIntervalMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsCreateInstrumentEvent) SettlementIntervalNullValue() int64 {
	return math.MinInt64
}

func (*OmsCreateInstrumentEvent) MinQuantityId() uint16 {
	return 7
}

func (*OmsCreateInstrumentEvent) MinQuantitySinceVersion() uint16 {
	return 0
}

func (o *OmsCreateInstrumentEvent) MinQuantityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MinQuantitySinceVersion()
}

func (*OmsCreateInstrumentEvent) MinQuantityDeprecated() uint16 {
	return 0
}

func (*OmsCreateInstrumentEvent) MinQuantityMetaAttribute(meta int) string {
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

func (*OmsCreateInstrumentEvent) MinQuantityMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsCreateInstrumentEvent) MinQuantityMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsCreateInstrumentEvent) MinQuantityNullValue() int64 {
	return math.MinInt64
}

func (*OmsCreateInstrumentEvent) BaseIncrementId() uint16 {
	return 8
}

func (*OmsCreateInstrumentEvent) BaseIncrementSinceVersion() uint16 {
	return 0
}

func (o *OmsCreateInstrumentEvent) BaseIncrementInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.BaseIncrementSinceVersion()
}

func (*OmsCreateInstrumentEvent) BaseIncrementDeprecated() uint16 {
	return 0
}

func (*OmsCreateInstrumentEvent) BaseIncrementMetaAttribute(meta int) string {
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

func (*OmsCreateInstrumentEvent) BaseIncrementMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsCreateInstrumentEvent) BaseIncrementMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsCreateInstrumentEvent) BaseIncrementNullValue() int64 {
	return math.MinInt64
}

func (*OmsCreateInstrumentEvent) QuoteIncrementId() uint16 {
	return 9
}

func (*OmsCreateInstrumentEvent) QuoteIncrementSinceVersion() uint16 {
	return 0
}

func (o *OmsCreateInstrumentEvent) QuoteIncrementInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.QuoteIncrementSinceVersion()
}

func (*OmsCreateInstrumentEvent) QuoteIncrementDeprecated() uint16 {
	return 0
}

func (*OmsCreateInstrumentEvent) QuoteIncrementMetaAttribute(meta int) string {
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

func (*OmsCreateInstrumentEvent) QuoteIncrementMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsCreateInstrumentEvent) QuoteIncrementMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsCreateInstrumentEvent) QuoteIncrementNullValue() int64 {
	return math.MinInt64
}

func (*OmsCreateInstrumentEvent) InitialMarginFractionFactorId() uint16 {
	return 10
}

func (*OmsCreateInstrumentEvent) InitialMarginFractionFactorSinceVersion() uint16 {
	return 0
}

func (o *OmsCreateInstrumentEvent) InitialMarginFractionFactorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.InitialMarginFractionFactorSinceVersion()
}

func (*OmsCreateInstrumentEvent) InitialMarginFractionFactorDeprecated() uint16 {
	return 0
}

func (*OmsCreateInstrumentEvent) InitialMarginFractionFactorMetaAttribute(meta int) string {
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

func (*OmsCreateInstrumentEvent) InitialMarginFractionFactorMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsCreateInstrumentEvent) InitialMarginFractionFactorMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsCreateInstrumentEvent) InitialMarginFractionFactorNullValue() int64 {
	return math.MinInt64
}

func (*OmsCreateInstrumentEvent) MaintenanceMarginFractionFactorId() uint16 {
	return 11
}

func (*OmsCreateInstrumentEvent) MaintenanceMarginFractionFactorSinceVersion() uint16 {
	return 0
}

func (o *OmsCreateInstrumentEvent) MaintenanceMarginFractionFactorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MaintenanceMarginFractionFactorSinceVersion()
}

func (*OmsCreateInstrumentEvent) MaintenanceMarginFractionFactorDeprecated() uint16 {
	return 0
}

func (*OmsCreateInstrumentEvent) MaintenanceMarginFractionFactorMetaAttribute(meta int) string {
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

func (*OmsCreateInstrumentEvent) MaintenanceMarginFractionFactorMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsCreateInstrumentEvent) MaintenanceMarginFractionFactorMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsCreateInstrumentEvent) MaintenanceMarginFractionFactorNullValue() int64 {
	return math.MinInt64
}

func (*OmsCreateInstrumentEvent) SymbolId() uint16 {
	return 12
}

func (*OmsCreateInstrumentEvent) SymbolSinceVersion() uint16 {
	return 0
}

func (o *OmsCreateInstrumentEvent) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SymbolSinceVersion()
}

func (*OmsCreateInstrumentEvent) SymbolDeprecated() uint16 {
	return 0
}

func (*OmsCreateInstrumentEvent) SymbolMetaAttribute(meta int) string {
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

func (*OmsCreateInstrumentEvent) SymbolMinValue() byte {
	return byte(32)
}

func (*OmsCreateInstrumentEvent) SymbolMaxValue() byte {
	return byte(126)
}

func (*OmsCreateInstrumentEvent) SymbolNullValue() byte {
	return 0
}

func (o *OmsCreateInstrumentEvent) SymbolCharacterEncoding() string {
	return "US-ASCII"
}

func (*OmsCreateInstrumentEvent) InstrumentTypeId() uint16 {
	return 13
}

func (*OmsCreateInstrumentEvent) InstrumentTypeSinceVersion() uint16 {
	return 0
}

func (o *OmsCreateInstrumentEvent) InstrumentTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.InstrumentTypeSinceVersion()
}

func (*OmsCreateInstrumentEvent) InstrumentTypeDeprecated() uint16 {
	return 0
}

func (*OmsCreateInstrumentEvent) InstrumentTypeMetaAttribute(meta int) string {
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
