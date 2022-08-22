// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MarketDataSnapshotTrailer struct {
	Timestamp   int64
	LogPosition int64
}

func (m *MarketDataSnapshotTrailer) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	return nil
}

func (m *MarketDataSnapshotTrailer) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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

func (m *MarketDataSnapshotTrailer) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	return nil
}

func MarketDataSnapshotTrailerInit(m *MarketDataSnapshotTrailer) {
	return
}

func (*MarketDataSnapshotTrailer) SbeBlockLength() (blockLength uint16) {
	return 16
}

func (*MarketDataSnapshotTrailer) SbeTemplateId() (templateId uint16) {
	return 2112
}

func (*MarketDataSnapshotTrailer) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*MarketDataSnapshotTrailer) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*MarketDataSnapshotTrailer) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*MarketDataSnapshotTrailer) TimestampId() uint16 {
	return 1
}

func (*MarketDataSnapshotTrailer) TimestampSinceVersion() uint16 {
	return 0
}

func (m *MarketDataSnapshotTrailer) TimestampInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TimestampSinceVersion()
}

func (*MarketDataSnapshotTrailer) TimestampDeprecated() uint16 {
	return 0
}

func (*MarketDataSnapshotTrailer) TimestampMetaAttribute(meta int) string {
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

func (*MarketDataSnapshotTrailer) TimestampMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataSnapshotTrailer) TimestampMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataSnapshotTrailer) TimestampNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataSnapshotTrailer) LogPositionId() uint16 {
	return 2
}

func (*MarketDataSnapshotTrailer) LogPositionSinceVersion() uint16 {
	return 0
}

func (m *MarketDataSnapshotTrailer) LogPositionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LogPositionSinceVersion()
}

func (*MarketDataSnapshotTrailer) LogPositionDeprecated() uint16 {
	return 0
}

func (*MarketDataSnapshotTrailer) LogPositionMetaAttribute(meta int) string {
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

func (*MarketDataSnapshotTrailer) LogPositionMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataSnapshotTrailer) LogPositionMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataSnapshotTrailer) LogPositionNullValue() int64 {
	return math.MinInt64
}
