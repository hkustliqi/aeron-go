// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MdOrderDelete struct {
	OrderId int64
}

func (m *MdOrderDelete) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := m.RangeCheck(m.SbeSchemaVersion(), m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, m.OrderId); err != nil {
		return err
	}
	return nil
}

func (m *MdOrderDelete) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !m.OrderIdInActingVersion(actingVersion) {
		m.OrderId = m.OrderIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.OrderId); err != nil {
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

func (m *MdOrderDelete) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.OrderIdInActingVersion(actingVersion) {
		if m.OrderId < m.OrderIdMinValue() || m.OrderId > m.OrderIdMaxValue() {
			return fmt.Errorf("Range check failed on m.OrderId (%v < %v > %v)", m.OrderIdMinValue(), m.OrderId, m.OrderIdMaxValue())
		}
	}
	return nil
}

func MdOrderDeleteInit(m *MdOrderDelete) {
	return
}

func (*MdOrderDelete) SbeBlockLength() (blockLength uint16) {
	return 8
}

func (*MdOrderDelete) SbeTemplateId() (templateId uint16) {
	return 521
}

func (*MdOrderDelete) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*MdOrderDelete) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*MdOrderDelete) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*MdOrderDelete) OrderIdId() uint16 {
	return 1
}

func (*MdOrderDelete) OrderIdSinceVersion() uint16 {
	return 0
}

func (m *MdOrderDelete) OrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OrderIdSinceVersion()
}

func (*MdOrderDelete) OrderIdDeprecated() uint16 {
	return 0
}

func (*MdOrderDelete) OrderIdMetaAttribute(meta int) string {
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

func (*MdOrderDelete) OrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MdOrderDelete) OrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*MdOrderDelete) OrderIdNullValue() int64 {
	return math.MinInt64
}
