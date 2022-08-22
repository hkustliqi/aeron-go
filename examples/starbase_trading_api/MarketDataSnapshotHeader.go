// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MarketDataSnapshotHeader struct {
	Timestamp    int64
	LogPosition  int64
	LastMdSeqNum int64
}

func (m *MarketDataSnapshotHeader) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := m.RangeCheck(m.SbeSchemaVersion(), m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, m.Timestamp); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.LogPosition); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.LastMdSeqNum); err != nil {
		return err
	}
	return nil
}

func (m *MarketDataSnapshotHeader) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !m.TimestampInActingVersion(actingVersion) {
		m.Timestamp = m.TimestampNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.Timestamp); err != nil {
			return err
		}
	}
	if !m.LogPositionInActingVersion(actingVersion) {
		m.LogPosition = m.LogPositionNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.LogPosition); err != nil {
			return err
		}
	}
	if !m.LastMdSeqNumInActingVersion(actingVersion) {
		m.LastMdSeqNum = m.LastMdSeqNumNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.LastMdSeqNum); err != nil {
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

func (m *MarketDataSnapshotHeader) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.TimestampInActingVersion(actingVersion) {
		if m.Timestamp < m.TimestampMinValue() || m.Timestamp > m.TimestampMaxValue() {
			return fmt.Errorf("Range check failed on m.Timestamp (%v < %v > %v)", m.TimestampMinValue(), m.Timestamp, m.TimestampMaxValue())
		}
	}
	if m.LogPositionInActingVersion(actingVersion) {
		if m.LogPosition < m.LogPositionMinValue() || m.LogPosition > m.LogPositionMaxValue() {
			return fmt.Errorf("Range check failed on m.LogPosition (%v < %v > %v)", m.LogPositionMinValue(), m.LogPosition, m.LogPositionMaxValue())
		}
	}
	if m.LastMdSeqNumInActingVersion(actingVersion) {
		if m.LastMdSeqNum < m.LastMdSeqNumMinValue() || m.LastMdSeqNum > m.LastMdSeqNumMaxValue() {
			return fmt.Errorf("Range check failed on m.LastMdSeqNum (%v < %v > %v)", m.LastMdSeqNumMinValue(), m.LastMdSeqNum, m.LastMdSeqNumMaxValue())
		}
	}
	return nil
}

func MarketDataSnapshotHeaderInit(m *MarketDataSnapshotHeader) {
	return
}

func (*MarketDataSnapshotHeader) SbeBlockLength() (blockLength uint16) {
	return 24
}

func (*MarketDataSnapshotHeader) SbeTemplateId() (templateId uint16) {
	return 2111
}

func (*MarketDataSnapshotHeader) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*MarketDataSnapshotHeader) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*MarketDataSnapshotHeader) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*MarketDataSnapshotHeader) TimestampId() uint16 {
	return 1
}

func (*MarketDataSnapshotHeader) TimestampSinceVersion() uint16 {
	return 0
}

func (m *MarketDataSnapshotHeader) TimestampInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TimestampSinceVersion()
}

func (*MarketDataSnapshotHeader) TimestampDeprecated() uint16 {
	return 0
}

func (*MarketDataSnapshotHeader) TimestampMetaAttribute(meta int) string {
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

func (*MarketDataSnapshotHeader) TimestampMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataSnapshotHeader) TimestampMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataSnapshotHeader) TimestampNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataSnapshotHeader) LogPositionId() uint16 {
	return 2
}

func (*MarketDataSnapshotHeader) LogPositionSinceVersion() uint16 {
	return 0
}

func (m *MarketDataSnapshotHeader) LogPositionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LogPositionSinceVersion()
}

func (*MarketDataSnapshotHeader) LogPositionDeprecated() uint16 {
	return 0
}

func (*MarketDataSnapshotHeader) LogPositionMetaAttribute(meta int) string {
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

func (*MarketDataSnapshotHeader) LogPositionMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataSnapshotHeader) LogPositionMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataSnapshotHeader) LogPositionNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataSnapshotHeader) LastMdSeqNumId() uint16 {
	return 3
}

func (*MarketDataSnapshotHeader) LastMdSeqNumSinceVersion() uint16 {
	return 0
}

func (m *MarketDataSnapshotHeader) LastMdSeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LastMdSeqNumSinceVersion()
}

func (*MarketDataSnapshotHeader) LastMdSeqNumDeprecated() uint16 {
	return 0
}

func (*MarketDataSnapshotHeader) LastMdSeqNumMetaAttribute(meta int) string {
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

func (*MarketDataSnapshotHeader) LastMdSeqNumMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataSnapshotHeader) LastMdSeqNumMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataSnapshotHeader) LastMdSeqNumNullValue() int64 {
	return math.MinInt64
}
