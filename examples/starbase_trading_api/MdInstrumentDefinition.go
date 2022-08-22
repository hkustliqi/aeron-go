// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MdInstrumentDefinition struct {
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

func (m *MdInstrumentDefinition) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := m.RangeCheck(m.SbeSchemaVersion(), m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	for idx := 0; idx < 2; idx++ {
		if err := _m.WriteInt64(_w, m.InstrumentUuid[idx]); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, m.BaseAsset); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.QuoteAsset); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.UnderlyingSpot); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.SettlementInterval); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.MinQuantity); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.BaseIncrement); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.QuoteIncrement); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.InitialMarginFractionFactor); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.MaintenanceMarginFractionFactor); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.Symbol[:]); err != nil {
		return err
	}
	if err := m.InstrumentType.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (m *MdInstrumentDefinition) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !m.InstrumentUuidInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			m.InstrumentUuid[idx] = m.InstrumentUuidNullValue()
		}
	} else {
		for idx := 0; idx < 2; idx++ {
			if err := _m.ReadInt64(_r, &m.InstrumentUuid[idx]); err != nil {
				return err
			}
		}
	}
	if !m.BaseAssetInActingVersion(actingVersion) {
		m.BaseAsset = m.BaseAssetNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.BaseAsset); err != nil {
			return err
		}
	}
	if !m.QuoteAssetInActingVersion(actingVersion) {
		m.QuoteAsset = m.QuoteAssetNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.QuoteAsset); err != nil {
			return err
		}
	}
	if !m.UnderlyingSpotInActingVersion(actingVersion) {
		m.UnderlyingSpot = m.UnderlyingSpotNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.UnderlyingSpot); err != nil {
			return err
		}
	}
	if !m.SettlementIntervalInActingVersion(actingVersion) {
		m.SettlementInterval = m.SettlementIntervalNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.SettlementInterval); err != nil {
			return err
		}
	}
	if !m.MinQuantityInActingVersion(actingVersion) {
		m.MinQuantity = m.MinQuantityNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.MinQuantity); err != nil {
			return err
		}
	}
	if !m.BaseIncrementInActingVersion(actingVersion) {
		m.BaseIncrement = m.BaseIncrementNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.BaseIncrement); err != nil {
			return err
		}
	}
	if !m.QuoteIncrementInActingVersion(actingVersion) {
		m.QuoteIncrement = m.QuoteIncrementNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.QuoteIncrement); err != nil {
			return err
		}
	}
	if !m.InitialMarginFractionFactorInActingVersion(actingVersion) {
		m.InitialMarginFractionFactor = m.InitialMarginFractionFactorNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.InitialMarginFractionFactor); err != nil {
			return err
		}
	}
	if !m.MaintenanceMarginFractionFactorInActingVersion(actingVersion) {
		m.MaintenanceMarginFractionFactor = m.MaintenanceMarginFractionFactorNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.MaintenanceMarginFractionFactor); err != nil {
			return err
		}
	}
	if !m.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			m.Symbol[idx] = m.SymbolNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, m.Symbol[:]); err != nil {
			return err
		}
	}
	if m.InstrumentTypeInActingVersion(actingVersion) {
		if err := m.InstrumentType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := m.RangeCheck(actingVersion, m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (m *MdInstrumentDefinition) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.InstrumentUuidInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if m.InstrumentUuid[idx] < m.InstrumentUuidMinValue() || m.InstrumentUuid[idx] > m.InstrumentUuidMaxValue() {
				return fmt.Errorf("Range check failed on m.InstrumentUuid[%d] (%v < %v > %v)", idx, m.InstrumentUuidMinValue(), m.InstrumentUuid[idx], m.InstrumentUuidMaxValue())
			}
		}
	}
	if m.BaseAssetInActingVersion(actingVersion) {
		if m.BaseAsset < m.BaseAssetMinValue() || m.BaseAsset > m.BaseAssetMaxValue() {
			return fmt.Errorf("Range check failed on m.BaseAsset (%v < %v > %v)", m.BaseAssetMinValue(), m.BaseAsset, m.BaseAssetMaxValue())
		}
	}
	if m.QuoteAssetInActingVersion(actingVersion) {
		if m.QuoteAsset < m.QuoteAssetMinValue() || m.QuoteAsset > m.QuoteAssetMaxValue() {
			return fmt.Errorf("Range check failed on m.QuoteAsset (%v < %v > %v)", m.QuoteAssetMinValue(), m.QuoteAsset, m.QuoteAssetMaxValue())
		}
	}
	if m.UnderlyingSpotInActingVersion(actingVersion) {
		if m.UnderlyingSpot != m.UnderlyingSpotNullValue() && (m.UnderlyingSpot < m.UnderlyingSpotMinValue() || m.UnderlyingSpot > m.UnderlyingSpotMaxValue()) {
			return fmt.Errorf("Range check failed on m.UnderlyingSpot (%v < %v > %v)", m.UnderlyingSpotMinValue(), m.UnderlyingSpot, m.UnderlyingSpotMaxValue())
		}
	}
	if m.SettlementIntervalInActingVersion(actingVersion) {
		if m.SettlementInterval != m.SettlementIntervalNullValue() && (m.SettlementInterval < m.SettlementIntervalMinValue() || m.SettlementInterval > m.SettlementIntervalMaxValue()) {
			return fmt.Errorf("Range check failed on m.SettlementInterval (%v < %v > %v)", m.SettlementIntervalMinValue(), m.SettlementInterval, m.SettlementIntervalMaxValue())
		}
	}
	if m.MinQuantityInActingVersion(actingVersion) {
		if m.MinQuantity < m.MinQuantityMinValue() || m.MinQuantity > m.MinQuantityMaxValue() {
			return fmt.Errorf("Range check failed on m.MinQuantity (%v < %v > %v)", m.MinQuantityMinValue(), m.MinQuantity, m.MinQuantityMaxValue())
		}
	}
	if m.BaseIncrementInActingVersion(actingVersion) {
		if m.BaseIncrement < m.BaseIncrementMinValue() || m.BaseIncrement > m.BaseIncrementMaxValue() {
			return fmt.Errorf("Range check failed on m.BaseIncrement (%v < %v > %v)", m.BaseIncrementMinValue(), m.BaseIncrement, m.BaseIncrementMaxValue())
		}
	}
	if m.QuoteIncrementInActingVersion(actingVersion) {
		if m.QuoteIncrement < m.QuoteIncrementMinValue() || m.QuoteIncrement > m.QuoteIncrementMaxValue() {
			return fmt.Errorf("Range check failed on m.QuoteIncrement (%v < %v > %v)", m.QuoteIncrementMinValue(), m.QuoteIncrement, m.QuoteIncrementMaxValue())
		}
	}
	if m.InitialMarginFractionFactorInActingVersion(actingVersion) {
		if m.InitialMarginFractionFactor != m.InitialMarginFractionFactorNullValue() && (m.InitialMarginFractionFactor < m.InitialMarginFractionFactorMinValue() || m.InitialMarginFractionFactor > m.InitialMarginFractionFactorMaxValue()) {
			return fmt.Errorf("Range check failed on m.InitialMarginFractionFactor (%v < %v > %v)", m.InitialMarginFractionFactorMinValue(), m.InitialMarginFractionFactor, m.InitialMarginFractionFactorMaxValue())
		}
	}
	if m.MaintenanceMarginFractionFactorInActingVersion(actingVersion) {
		if m.MaintenanceMarginFractionFactor != m.MaintenanceMarginFractionFactorNullValue() && (m.MaintenanceMarginFractionFactor < m.MaintenanceMarginFractionFactorMinValue() || m.MaintenanceMarginFractionFactor > m.MaintenanceMarginFractionFactorMaxValue()) {
			return fmt.Errorf("Range check failed on m.MaintenanceMarginFractionFactor (%v < %v > %v)", m.MaintenanceMarginFractionFactorMinValue(), m.MaintenanceMarginFractionFactor, m.MaintenanceMarginFractionFactorMaxValue())
		}
	}
	if m.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 80; idx++ {
			if m.Symbol[idx] < m.SymbolMinValue() || m.Symbol[idx] > m.SymbolMaxValue() {
				return fmt.Errorf("Range check failed on m.Symbol[%d] (%v < %v > %v)", idx, m.SymbolMinValue(), m.Symbol[idx], m.SymbolMaxValue())
			}
		}
	}
	for idx, ch := range m.Symbol {
		if ch > 127 {
			return fmt.Errorf("m.Symbol[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	if err := m.InstrumentType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func MdInstrumentDefinitionInit(m *MdInstrumentDefinition) {
	m.UnderlyingSpot = math.MinInt64
	m.SettlementInterval = math.MinInt64
	m.InitialMarginFractionFactor = math.MinInt64
	m.MaintenanceMarginFractionFactor = math.MinInt64
	return
}

func (*MdInstrumentDefinition) SbeBlockLength() (blockLength uint16) {
	return 169
}

func (*MdInstrumentDefinition) SbeTemplateId() (templateId uint16) {
	return 522
}

func (*MdInstrumentDefinition) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*MdInstrumentDefinition) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*MdInstrumentDefinition) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*MdInstrumentDefinition) InstrumentUuidId() uint16 {
	return 1
}

func (*MdInstrumentDefinition) InstrumentUuidSinceVersion() uint16 {
	return 0
}

func (m *MdInstrumentDefinition) InstrumentUuidInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.InstrumentUuidSinceVersion()
}

func (*MdInstrumentDefinition) InstrumentUuidDeprecated() uint16 {
	return 0
}

func (*MdInstrumentDefinition) InstrumentUuidMetaAttribute(meta int) string {
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

func (*MdInstrumentDefinition) InstrumentUuidMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MdInstrumentDefinition) InstrumentUuidMaxValue() int64 {
	return math.MaxInt64
}

func (*MdInstrumentDefinition) InstrumentUuidNullValue() int64 {
	return math.MinInt64
}

func (*MdInstrumentDefinition) BaseAssetId() uint16 {
	return 2
}

func (*MdInstrumentDefinition) BaseAssetSinceVersion() uint16 {
	return 0
}

func (m *MdInstrumentDefinition) BaseAssetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.BaseAssetSinceVersion()
}

func (*MdInstrumentDefinition) BaseAssetDeprecated() uint16 {
	return 0
}

func (*MdInstrumentDefinition) BaseAssetMetaAttribute(meta int) string {
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

func (*MdInstrumentDefinition) BaseAssetMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MdInstrumentDefinition) BaseAssetMaxValue() int64 {
	return math.MaxInt64
}

func (*MdInstrumentDefinition) BaseAssetNullValue() int64 {
	return math.MinInt64
}

func (*MdInstrumentDefinition) QuoteAssetId() uint16 {
	return 3
}

func (*MdInstrumentDefinition) QuoteAssetSinceVersion() uint16 {
	return 0
}

func (m *MdInstrumentDefinition) QuoteAssetInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QuoteAssetSinceVersion()
}

func (*MdInstrumentDefinition) QuoteAssetDeprecated() uint16 {
	return 0
}

func (*MdInstrumentDefinition) QuoteAssetMetaAttribute(meta int) string {
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

func (*MdInstrumentDefinition) QuoteAssetMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MdInstrumentDefinition) QuoteAssetMaxValue() int64 {
	return math.MaxInt64
}

func (*MdInstrumentDefinition) QuoteAssetNullValue() int64 {
	return math.MinInt64
}

func (*MdInstrumentDefinition) UnderlyingSpotId() uint16 {
	return 4
}

func (*MdInstrumentDefinition) UnderlyingSpotSinceVersion() uint16 {
	return 0
}

func (m *MdInstrumentDefinition) UnderlyingSpotInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.UnderlyingSpotSinceVersion()
}

func (*MdInstrumentDefinition) UnderlyingSpotDeprecated() uint16 {
	return 0
}

func (*MdInstrumentDefinition) UnderlyingSpotMetaAttribute(meta int) string {
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

func (*MdInstrumentDefinition) UnderlyingSpotMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MdInstrumentDefinition) UnderlyingSpotMaxValue() int64 {
	return math.MaxInt64
}

func (*MdInstrumentDefinition) UnderlyingSpotNullValue() int64 {
	return math.MinInt64
}

func (*MdInstrumentDefinition) SettlementIntervalId() uint16 {
	return 5
}

func (*MdInstrumentDefinition) SettlementIntervalSinceVersion() uint16 {
	return 0
}

func (m *MdInstrumentDefinition) SettlementIntervalInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SettlementIntervalSinceVersion()
}

func (*MdInstrumentDefinition) SettlementIntervalDeprecated() uint16 {
	return 0
}

func (*MdInstrumentDefinition) SettlementIntervalMetaAttribute(meta int) string {
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

func (*MdInstrumentDefinition) SettlementIntervalMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MdInstrumentDefinition) SettlementIntervalMaxValue() int64 {
	return math.MaxInt64
}

func (*MdInstrumentDefinition) SettlementIntervalNullValue() int64 {
	return math.MinInt64
}

func (*MdInstrumentDefinition) MinQuantityId() uint16 {
	return 6
}

func (*MdInstrumentDefinition) MinQuantitySinceVersion() uint16 {
	return 0
}

func (m *MdInstrumentDefinition) MinQuantityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MinQuantitySinceVersion()
}

func (*MdInstrumentDefinition) MinQuantityDeprecated() uint16 {
	return 0
}

func (*MdInstrumentDefinition) MinQuantityMetaAttribute(meta int) string {
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

func (*MdInstrumentDefinition) MinQuantityMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MdInstrumentDefinition) MinQuantityMaxValue() int64 {
	return math.MaxInt64
}

func (*MdInstrumentDefinition) MinQuantityNullValue() int64 {
	return math.MinInt64
}

func (*MdInstrumentDefinition) BaseIncrementId() uint16 {
	return 7
}

func (*MdInstrumentDefinition) BaseIncrementSinceVersion() uint16 {
	return 0
}

func (m *MdInstrumentDefinition) BaseIncrementInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.BaseIncrementSinceVersion()
}

func (*MdInstrumentDefinition) BaseIncrementDeprecated() uint16 {
	return 0
}

func (*MdInstrumentDefinition) BaseIncrementMetaAttribute(meta int) string {
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

func (*MdInstrumentDefinition) BaseIncrementMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MdInstrumentDefinition) BaseIncrementMaxValue() int64 {
	return math.MaxInt64
}

func (*MdInstrumentDefinition) BaseIncrementNullValue() int64 {
	return math.MinInt64
}

func (*MdInstrumentDefinition) QuoteIncrementId() uint16 {
	return 8
}

func (*MdInstrumentDefinition) QuoteIncrementSinceVersion() uint16 {
	return 0
}

func (m *MdInstrumentDefinition) QuoteIncrementInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QuoteIncrementSinceVersion()
}

func (*MdInstrumentDefinition) QuoteIncrementDeprecated() uint16 {
	return 0
}

func (*MdInstrumentDefinition) QuoteIncrementMetaAttribute(meta int) string {
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

func (*MdInstrumentDefinition) QuoteIncrementMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MdInstrumentDefinition) QuoteIncrementMaxValue() int64 {
	return math.MaxInt64
}

func (*MdInstrumentDefinition) QuoteIncrementNullValue() int64 {
	return math.MinInt64
}

func (*MdInstrumentDefinition) InitialMarginFractionFactorId() uint16 {
	return 9
}

func (*MdInstrumentDefinition) InitialMarginFractionFactorSinceVersion() uint16 {
	return 0
}

func (m *MdInstrumentDefinition) InitialMarginFractionFactorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.InitialMarginFractionFactorSinceVersion()
}

func (*MdInstrumentDefinition) InitialMarginFractionFactorDeprecated() uint16 {
	return 0
}

func (*MdInstrumentDefinition) InitialMarginFractionFactorMetaAttribute(meta int) string {
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

func (*MdInstrumentDefinition) InitialMarginFractionFactorMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MdInstrumentDefinition) InitialMarginFractionFactorMaxValue() int64 {
	return math.MaxInt64
}

func (*MdInstrumentDefinition) InitialMarginFractionFactorNullValue() int64 {
	return math.MinInt64
}

func (*MdInstrumentDefinition) MaintenanceMarginFractionFactorId() uint16 {
	return 10
}

func (*MdInstrumentDefinition) MaintenanceMarginFractionFactorSinceVersion() uint16 {
	return 0
}

func (m *MdInstrumentDefinition) MaintenanceMarginFractionFactorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MaintenanceMarginFractionFactorSinceVersion()
}

func (*MdInstrumentDefinition) MaintenanceMarginFractionFactorDeprecated() uint16 {
	return 0
}

func (*MdInstrumentDefinition) MaintenanceMarginFractionFactorMetaAttribute(meta int) string {
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

func (*MdInstrumentDefinition) MaintenanceMarginFractionFactorMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MdInstrumentDefinition) MaintenanceMarginFractionFactorMaxValue() int64 {
	return math.MaxInt64
}

func (*MdInstrumentDefinition) MaintenanceMarginFractionFactorNullValue() int64 {
	return math.MinInt64
}

func (*MdInstrumentDefinition) SymbolId() uint16 {
	return 11
}

func (*MdInstrumentDefinition) SymbolSinceVersion() uint16 {
	return 0
}

func (m *MdInstrumentDefinition) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SymbolSinceVersion()
}

func (*MdInstrumentDefinition) SymbolDeprecated() uint16 {
	return 0
}

func (*MdInstrumentDefinition) SymbolMetaAttribute(meta int) string {
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

func (*MdInstrumentDefinition) SymbolMinValue() byte {
	return byte(32)
}

func (*MdInstrumentDefinition) SymbolMaxValue() byte {
	return byte(126)
}

func (*MdInstrumentDefinition) SymbolNullValue() byte {
	return 0
}

func (m *MdInstrumentDefinition) SymbolCharacterEncoding() string {
	return "US-ASCII"
}

func (*MdInstrumentDefinition) InstrumentTypeId() uint16 {
	return 12
}

func (*MdInstrumentDefinition) InstrumentTypeSinceVersion() uint16 {
	return 0
}

func (m *MdInstrumentDefinition) InstrumentTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.InstrumentTypeSinceVersion()
}

func (*MdInstrumentDefinition) InstrumentTypeDeprecated() uint16 {
	return 0
}

func (*MdInstrumentDefinition) InstrumentTypeMetaAttribute(meta int) string {
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
