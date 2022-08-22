// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MdOrderPut struct {
	OrderId  int64
	Price    int64
	Quantity int64
}

func (m *MdOrderPut) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := m.RangeCheck(m.SbeSchemaVersion(), m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, m.OrderId); err != nil {
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

func (m *MdOrderPut) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !m.OrderIdInActingVersion(actingVersion) {
		m.OrderId = m.OrderIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.OrderId); err != nil {
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

func (m *MdOrderPut) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.OrderIdInActingVersion(actingVersion) {
		if m.OrderId < m.OrderIdMinValue() || m.OrderId > m.OrderIdMaxValue() {
			return fmt.Errorf("Range check failed on m.OrderId (%v < %v > %v)", m.OrderIdMinValue(), m.OrderId, m.OrderIdMaxValue())
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

func MdOrderPutInit(m *MdOrderPut) {
	return
}

func (*MdOrderPut) SbeBlockLength() (blockLength uint16) {
	return 24
}

func (*MdOrderPut) SbeTemplateId() (templateId uint16) {
	return 520
}

func (*MdOrderPut) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*MdOrderPut) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*MdOrderPut) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*MdOrderPut) OrderIdId() uint16 {
	return 1
}

func (*MdOrderPut) OrderIdSinceVersion() uint16 {
	return 0
}

func (m *MdOrderPut) OrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OrderIdSinceVersion()
}

func (*MdOrderPut) OrderIdDeprecated() uint16 {
	return 0
}

func (*MdOrderPut) OrderIdMetaAttribute(meta int) string {
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

func (*MdOrderPut) OrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MdOrderPut) OrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*MdOrderPut) OrderIdNullValue() int64 {
	return math.MinInt64
}

func (*MdOrderPut) PriceId() uint16 {
	return 2
}

func (*MdOrderPut) PriceSinceVersion() uint16 {
	return 0
}

func (m *MdOrderPut) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.PriceSinceVersion()
}

func (*MdOrderPut) PriceDeprecated() uint16 {
	return 0
}

func (*MdOrderPut) PriceMetaAttribute(meta int) string {
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

func (*MdOrderPut) PriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MdOrderPut) PriceMaxValue() int64 {
	return math.MaxInt64
}

func (*MdOrderPut) PriceNullValue() int64 {
	return math.MinInt64
}

func (*MdOrderPut) QuantityId() uint16 {
	return 3
}

func (*MdOrderPut) QuantitySinceVersion() uint16 {
	return 0
}

func (m *MdOrderPut) QuantityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QuantitySinceVersion()
}

func (*MdOrderPut) QuantityDeprecated() uint16 {
	return 0
}

func (*MdOrderPut) QuantityMetaAttribute(meta int) string {
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

func (*MdOrderPut) QuantityMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MdOrderPut) QuantityMaxValue() int64 {
	return math.MaxInt64
}

func (*MdOrderPut) QuantityNullValue() int64 {
	return math.MinInt64
}
