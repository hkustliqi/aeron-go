// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MdTradeSummary struct {
	AggressorOrderId     int64
	AggressorReceiveTime int64
	VwapPrice            int64
	DeepestPrice         int64
	Quantity             int64
}

func (m *MdTradeSummary) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := m.RangeCheck(m.SbeSchemaVersion(), m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, m.AggressorOrderId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.AggressorReceiveTime); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.VwapPrice); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.DeepestPrice); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.Quantity); err != nil {
		return err
	}
	return nil
}

func (m *MdTradeSummary) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !m.AggressorOrderIdInActingVersion(actingVersion) {
		m.AggressorOrderId = m.AggressorOrderIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.AggressorOrderId); err != nil {
			return err
		}
	}
	if !m.AggressorReceiveTimeInActingVersion(actingVersion) {
		m.AggressorReceiveTime = m.AggressorReceiveTimeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.AggressorReceiveTime); err != nil {
			return err
		}
	}
	if !m.VwapPriceInActingVersion(actingVersion) {
		m.VwapPrice = m.VwapPriceNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.VwapPrice); err != nil {
			return err
		}
	}
	if !m.DeepestPriceInActingVersion(actingVersion) {
		m.DeepestPrice = m.DeepestPriceNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.DeepestPrice); err != nil {
			return err
		}
	}
	if !m.QuantityInActingVersion(actingVersion) {
		m.Quantity = m.QuantityNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.Quantity); err != nil {
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

func (m *MdTradeSummary) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.AggressorOrderIdInActingVersion(actingVersion) {
		if m.AggressorOrderId < m.AggressorOrderIdMinValue() || m.AggressorOrderId > m.AggressorOrderIdMaxValue() {
			return fmt.Errorf("Range check failed on m.AggressorOrderId (%v < %v > %v)", m.AggressorOrderIdMinValue(), m.AggressorOrderId, m.AggressorOrderIdMaxValue())
		}
	}
	if m.AggressorReceiveTimeInActingVersion(actingVersion) {
		if m.AggressorReceiveTime < m.AggressorReceiveTimeMinValue() || m.AggressorReceiveTime > m.AggressorReceiveTimeMaxValue() {
			return fmt.Errorf("Range check failed on m.AggressorReceiveTime (%v < %v > %v)", m.AggressorReceiveTimeMinValue(), m.AggressorReceiveTime, m.AggressorReceiveTimeMaxValue())
		}
	}
	if m.VwapPriceInActingVersion(actingVersion) {
		if m.VwapPrice < m.VwapPriceMinValue() || m.VwapPrice > m.VwapPriceMaxValue() {
			return fmt.Errorf("Range check failed on m.VwapPrice (%v < %v > %v)", m.VwapPriceMinValue(), m.VwapPrice, m.VwapPriceMaxValue())
		}
	}
	if m.DeepestPriceInActingVersion(actingVersion) {
		if m.DeepestPrice < m.DeepestPriceMinValue() || m.DeepestPrice > m.DeepestPriceMaxValue() {
			return fmt.Errorf("Range check failed on m.DeepestPrice (%v < %v > %v)", m.DeepestPriceMinValue(), m.DeepestPrice, m.DeepestPriceMaxValue())
		}
	}
	if m.QuantityInActingVersion(actingVersion) {
		if m.Quantity < m.QuantityMinValue() || m.Quantity > m.QuantityMaxValue() {
			return fmt.Errorf("Range check failed on m.Quantity (%v < %v > %v)", m.QuantityMinValue(), m.Quantity, m.QuantityMaxValue())
		}
	}
	return nil
}

func MdTradeSummaryInit(m *MdTradeSummary) {
	return
}

func (*MdTradeSummary) SbeBlockLength() (blockLength uint16) {
	return 40
}

func (*MdTradeSummary) SbeTemplateId() (templateId uint16) {
	return 533
}

func (*MdTradeSummary) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*MdTradeSummary) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*MdTradeSummary) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*MdTradeSummary) AggressorOrderIdId() uint16 {
	return 1
}

func (*MdTradeSummary) AggressorOrderIdSinceVersion() uint16 {
	return 0
}

func (m *MdTradeSummary) AggressorOrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.AggressorOrderIdSinceVersion()
}

func (*MdTradeSummary) AggressorOrderIdDeprecated() uint16 {
	return 0
}

func (*MdTradeSummary) AggressorOrderIdMetaAttribute(meta int) string {
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

func (*MdTradeSummary) AggressorOrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MdTradeSummary) AggressorOrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*MdTradeSummary) AggressorOrderIdNullValue() int64 {
	return math.MinInt64
}

func (*MdTradeSummary) AggressorReceiveTimeId() uint16 {
	return 2
}

func (*MdTradeSummary) AggressorReceiveTimeSinceVersion() uint16 {
	return 0
}

func (m *MdTradeSummary) AggressorReceiveTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.AggressorReceiveTimeSinceVersion()
}

func (*MdTradeSummary) AggressorReceiveTimeDeprecated() uint16 {
	return 0
}

func (*MdTradeSummary) AggressorReceiveTimeMetaAttribute(meta int) string {
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

func (*MdTradeSummary) AggressorReceiveTimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MdTradeSummary) AggressorReceiveTimeMaxValue() int64 {
	return math.MaxInt64
}

func (*MdTradeSummary) AggressorReceiveTimeNullValue() int64 {
	return math.MinInt64
}

func (*MdTradeSummary) VwapPriceId() uint16 {
	return 3
}

func (*MdTradeSummary) VwapPriceSinceVersion() uint16 {
	return 0
}

func (m *MdTradeSummary) VwapPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.VwapPriceSinceVersion()
}

func (*MdTradeSummary) VwapPriceDeprecated() uint16 {
	return 0
}

func (*MdTradeSummary) VwapPriceMetaAttribute(meta int) string {
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

func (*MdTradeSummary) VwapPriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MdTradeSummary) VwapPriceMaxValue() int64 {
	return math.MaxInt64
}

func (*MdTradeSummary) VwapPriceNullValue() int64 {
	return math.MinInt64
}

func (*MdTradeSummary) DeepestPriceId() uint16 {
	return 4
}

func (*MdTradeSummary) DeepestPriceSinceVersion() uint16 {
	return 0
}

func (m *MdTradeSummary) DeepestPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.DeepestPriceSinceVersion()
}

func (*MdTradeSummary) DeepestPriceDeprecated() uint16 {
	return 0
}

func (*MdTradeSummary) DeepestPriceMetaAttribute(meta int) string {
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

func (*MdTradeSummary) DeepestPriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MdTradeSummary) DeepestPriceMaxValue() int64 {
	return math.MaxInt64
}

func (*MdTradeSummary) DeepestPriceNullValue() int64 {
	return math.MinInt64
}

func (*MdTradeSummary) QuantityId() uint16 {
	return 5
}

func (*MdTradeSummary) QuantitySinceVersion() uint16 {
	return 0
}

func (m *MdTradeSummary) QuantityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QuantitySinceVersion()
}

func (*MdTradeSummary) QuantityDeprecated() uint16 {
	return 0
}

func (*MdTradeSummary) QuantityMetaAttribute(meta int) string {
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

func (*MdTradeSummary) QuantityMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MdTradeSummary) QuantityMaxValue() int64 {
	return math.MaxInt64
}

func (*MdTradeSummary) QuantityNullValue() int64 {
	return math.MinInt64
}
