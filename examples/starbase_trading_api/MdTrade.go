// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MdTrade struct {
	MatchId     int64
	BuyOrderId  int64
	SellOrderId int64
	Price       int64
	Quantity    int64
}

func (m *MdTrade) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := m.RangeCheck(m.SbeSchemaVersion(), m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, m.MatchId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.BuyOrderId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.SellOrderId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.Price); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.Quantity); err != nil {
		return err
	}
	return nil
}

func (m *MdTrade) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !m.MatchIdInActingVersion(actingVersion) {
		m.MatchId = m.MatchIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.MatchId); err != nil {
			return err
		}
	}
	if !m.BuyOrderIdInActingVersion(actingVersion) {
		m.BuyOrderId = m.BuyOrderIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.BuyOrderId); err != nil {
			return err
		}
	}
	if !m.SellOrderIdInActingVersion(actingVersion) {
		m.SellOrderId = m.SellOrderIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.SellOrderId); err != nil {
			return err
		}
	}
	if !m.PriceInActingVersion(actingVersion) {
		m.Price = m.PriceNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.Price); err != nil {
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

func (m *MdTrade) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.MatchIdInActingVersion(actingVersion) {
		if m.MatchId < m.MatchIdMinValue() || m.MatchId > m.MatchIdMaxValue() {
			return fmt.Errorf("Range check failed on m.MatchId (%v < %v > %v)", m.MatchIdMinValue(), m.MatchId, m.MatchIdMaxValue())
		}
	}
	if m.BuyOrderIdInActingVersion(actingVersion) {
		if m.BuyOrderId != m.BuyOrderIdNullValue() && (m.BuyOrderId < m.BuyOrderIdMinValue() || m.BuyOrderId > m.BuyOrderIdMaxValue()) {
			return fmt.Errorf("Range check failed on m.BuyOrderId (%v < %v > %v)", m.BuyOrderIdMinValue(), m.BuyOrderId, m.BuyOrderIdMaxValue())
		}
	}
	if m.SellOrderIdInActingVersion(actingVersion) {
		if m.SellOrderId != m.SellOrderIdNullValue() && (m.SellOrderId < m.SellOrderIdMinValue() || m.SellOrderId > m.SellOrderIdMaxValue()) {
			return fmt.Errorf("Range check failed on m.SellOrderId (%v < %v > %v)", m.SellOrderIdMinValue(), m.SellOrderId, m.SellOrderIdMaxValue())
		}
	}
	if m.PriceInActingVersion(actingVersion) {
		if m.Price < m.PriceMinValue() || m.Price > m.PriceMaxValue() {
			return fmt.Errorf("Range check failed on m.Price (%v < %v > %v)", m.PriceMinValue(), m.Price, m.PriceMaxValue())
		}
	}
	if m.QuantityInActingVersion(actingVersion) {
		if m.Quantity < m.QuantityMinValue() || m.Quantity > m.QuantityMaxValue() {
			return fmt.Errorf("Range check failed on m.Quantity (%v < %v > %v)", m.QuantityMinValue(), m.Quantity, m.QuantityMaxValue())
		}
	}
	return nil
}

func MdTradeInit(m *MdTrade) {
	m.BuyOrderId = math.MinInt64
	m.SellOrderId = math.MinInt64
	return
}

func (*MdTrade) SbeBlockLength() (blockLength uint16) {
	return 40
}

func (*MdTrade) SbeTemplateId() (templateId uint16) {
	return 530
}

func (*MdTrade) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*MdTrade) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*MdTrade) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*MdTrade) MatchIdId() uint16 {
	return 1
}

func (*MdTrade) MatchIdSinceVersion() uint16 {
	return 0
}

func (m *MdTrade) MatchIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MatchIdSinceVersion()
}

func (*MdTrade) MatchIdDeprecated() uint16 {
	return 0
}

func (*MdTrade) MatchIdMetaAttribute(meta int) string {
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

func (*MdTrade) MatchIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MdTrade) MatchIdMaxValue() int64 {
	return math.MaxInt64
}

func (*MdTrade) MatchIdNullValue() int64 {
	return math.MinInt64
}

func (*MdTrade) BuyOrderIdId() uint16 {
	return 2
}

func (*MdTrade) BuyOrderIdSinceVersion() uint16 {
	return 0
}

func (m *MdTrade) BuyOrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.BuyOrderIdSinceVersion()
}

func (*MdTrade) BuyOrderIdDeprecated() uint16 {
	return 0
}

func (*MdTrade) BuyOrderIdMetaAttribute(meta int) string {
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

func (*MdTrade) BuyOrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MdTrade) BuyOrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*MdTrade) BuyOrderIdNullValue() int64 {
	return math.MinInt64
}

func (*MdTrade) SellOrderIdId() uint16 {
	return 3
}

func (*MdTrade) SellOrderIdSinceVersion() uint16 {
	return 0
}

func (m *MdTrade) SellOrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SellOrderIdSinceVersion()
}

func (*MdTrade) SellOrderIdDeprecated() uint16 {
	return 0
}

func (*MdTrade) SellOrderIdMetaAttribute(meta int) string {
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

func (*MdTrade) SellOrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MdTrade) SellOrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*MdTrade) SellOrderIdNullValue() int64 {
	return math.MinInt64
}

func (*MdTrade) PriceId() uint16 {
	return 4
}

func (*MdTrade) PriceSinceVersion() uint16 {
	return 0
}

func (m *MdTrade) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.PriceSinceVersion()
}

func (*MdTrade) PriceDeprecated() uint16 {
	return 0
}

func (*MdTrade) PriceMetaAttribute(meta int) string {
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

func (*MdTrade) PriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MdTrade) PriceMaxValue() int64 {
	return math.MaxInt64
}

func (*MdTrade) PriceNullValue() int64 {
	return math.MinInt64
}

func (*MdTrade) QuantityId() uint16 {
	return 5
}

func (*MdTrade) QuantitySinceVersion() uint16 {
	return 0
}

func (m *MdTrade) QuantityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QuantitySinceVersion()
}

func (*MdTrade) QuantityDeprecated() uint16 {
	return 0
}

func (*MdTrade) QuantityMetaAttribute(meta int) string {
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

func (*MdTrade) QuantityMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MdTrade) QuantityMaxValue() int64 {
	return math.MaxInt64
}

func (*MdTrade) QuantityNullValue() int64 {
	return math.MinInt64
}
