// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MarketDataHeader struct {
	LogPosition   int64
	MdSeqNum      int64
	InstrSeqNum   int64
	Timestamp     int64
	InstrumentId  int64
	Flags         MarketDataFlags
	ShardId       int8
	SourceService SourceServiceEnum
	Side          SideEnum
	TemplateId    uint16
	BlockLength   uint16
}

func (m *MarketDataHeader) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := m.RangeCheck(m.SbeSchemaVersion(), m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, m.LogPosition); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.MdSeqNum); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.InstrSeqNum); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.Timestamp); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.InstrumentId); err != nil {
		return err
	}
	if err := m.Flags.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, m.ShardId); err != nil {
		return err
	}
	if err := m.SourceService.Encode(_m, _w); err != nil {
		return err
	}
	if err := m.Side.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, m.TemplateId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, m.BlockLength); err != nil {
		return err
	}
	return nil
}

func (m *MarketDataHeader) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !m.LogPositionInActingVersion(actingVersion) {
		m.LogPosition = m.LogPositionNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.LogPosition); err != nil {
			return err
		}
	}
	if !m.MdSeqNumInActingVersion(actingVersion) {
		m.MdSeqNum = m.MdSeqNumNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.MdSeqNum); err != nil {
			return err
		}
	}
	if !m.InstrSeqNumInActingVersion(actingVersion) {
		m.InstrSeqNum = m.InstrSeqNumNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.InstrSeqNum); err != nil {
			return err
		}
	}
	if !m.TimestampInActingVersion(actingVersion) {
		m.Timestamp = m.TimestampNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.Timestamp); err != nil {
			return err
		}
	}
	if !m.InstrumentIdInActingVersion(actingVersion) {
		m.InstrumentId = m.InstrumentIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.InstrumentId); err != nil {
			return err
		}
	}
	if m.FlagsInActingVersion(actingVersion) {
		if err := m.Flags.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.ShardIdInActingVersion(actingVersion) {
		m.ShardId = m.ShardIdNullValue()
	} else {
		if err := _m.ReadInt8(_r, &m.ShardId); err != nil {
			return err
		}
	}
	if m.SourceServiceInActingVersion(actingVersion) {
		if err := m.SourceService.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if m.SideInActingVersion(actingVersion) {
		if err := m.Side.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.TemplateIdInActingVersion(actingVersion) {
		m.TemplateId = m.TemplateIdNullValue()
	} else {
		if err := _m.ReadUint16(_r, &m.TemplateId); err != nil {
			return err
		}
	}
	if !m.BlockLengthInActingVersion(actingVersion) {
		m.BlockLength = m.BlockLengthNullValue()
	} else {
		if err := _m.ReadUint16(_r, &m.BlockLength); err != nil {
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

func (m *MarketDataHeader) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.LogPositionInActingVersion(actingVersion) {
		if m.LogPosition < m.LogPositionMinValue() || m.LogPosition > m.LogPositionMaxValue() {
			return fmt.Errorf("Range check failed on m.LogPosition (%v < %v > %v)", m.LogPositionMinValue(), m.LogPosition, m.LogPositionMaxValue())
		}
	}
	if m.MdSeqNumInActingVersion(actingVersion) {
		if m.MdSeqNum < m.MdSeqNumMinValue() || m.MdSeqNum > m.MdSeqNumMaxValue() {
			return fmt.Errorf("Range check failed on m.MdSeqNum (%v < %v > %v)", m.MdSeqNumMinValue(), m.MdSeqNum, m.MdSeqNumMaxValue())
		}
	}
	if m.InstrSeqNumInActingVersion(actingVersion) {
		if m.InstrSeqNum < m.InstrSeqNumMinValue() || m.InstrSeqNum > m.InstrSeqNumMaxValue() {
			return fmt.Errorf("Range check failed on m.InstrSeqNum (%v < %v > %v)", m.InstrSeqNumMinValue(), m.InstrSeqNum, m.InstrSeqNumMaxValue())
		}
	}
	if m.TimestampInActingVersion(actingVersion) {
		if m.Timestamp < m.TimestampMinValue() || m.Timestamp > m.TimestampMaxValue() {
			return fmt.Errorf("Range check failed on m.Timestamp (%v < %v > %v)", m.TimestampMinValue(), m.Timestamp, m.TimestampMaxValue())
		}
	}
	if m.InstrumentIdInActingVersion(actingVersion) {
		if m.InstrumentId < m.InstrumentIdMinValue() || m.InstrumentId > m.InstrumentIdMaxValue() {
			return fmt.Errorf("Range check failed on m.InstrumentId (%v < %v > %v)", m.InstrumentIdMinValue(), m.InstrumentId, m.InstrumentIdMaxValue())
		}
	}
	if m.ShardIdInActingVersion(actingVersion) {
		if m.ShardId < m.ShardIdMinValue() || m.ShardId > m.ShardIdMaxValue() {
			return fmt.Errorf("Range check failed on m.ShardId (%v < %v > %v)", m.ShardIdMinValue(), m.ShardId, m.ShardIdMaxValue())
		}
	}
	if err := m.SourceService.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := m.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if m.TemplateIdInActingVersion(actingVersion) {
		if m.TemplateId < m.TemplateIdMinValue() || m.TemplateId > m.TemplateIdMaxValue() {
			return fmt.Errorf("Range check failed on m.TemplateId (%v < %v > %v)", m.TemplateIdMinValue(), m.TemplateId, m.TemplateIdMaxValue())
		}
	}
	if m.BlockLengthInActingVersion(actingVersion) {
		if m.BlockLength < m.BlockLengthMinValue() || m.BlockLength > m.BlockLengthMaxValue() {
			return fmt.Errorf("Range check failed on m.BlockLength (%v < %v > %v)", m.BlockLengthMinValue(), m.BlockLength, m.BlockLengthMaxValue())
		}
	}
	return nil
}

func MarketDataHeaderInit(m *MarketDataHeader) {
	return
}

func (*MarketDataHeader) SbeBlockLength() (blockLength uint16) {
	return 48
}

func (*MarketDataHeader) SbeTemplateId() (templateId uint16) {
	return 101
}

func (*MarketDataHeader) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*MarketDataHeader) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*MarketDataHeader) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*MarketDataHeader) LogPositionId() uint16 {
	return 1
}

func (*MarketDataHeader) LogPositionSinceVersion() uint16 {
	return 0
}

func (m *MarketDataHeader) LogPositionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LogPositionSinceVersion()
}

func (*MarketDataHeader) LogPositionDeprecated() uint16 {
	return 0
}

func (*MarketDataHeader) LogPositionMetaAttribute(meta int) string {
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

func (*MarketDataHeader) LogPositionMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataHeader) LogPositionMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataHeader) LogPositionNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataHeader) MdSeqNumId() uint16 {
	return 2
}

func (*MarketDataHeader) MdSeqNumSinceVersion() uint16 {
	return 0
}

func (m *MarketDataHeader) MdSeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MdSeqNumSinceVersion()
}

func (*MarketDataHeader) MdSeqNumDeprecated() uint16 {
	return 0
}

func (*MarketDataHeader) MdSeqNumMetaAttribute(meta int) string {
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

func (*MarketDataHeader) MdSeqNumMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataHeader) MdSeqNumMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataHeader) MdSeqNumNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataHeader) InstrSeqNumId() uint16 {
	return 3
}

func (*MarketDataHeader) InstrSeqNumSinceVersion() uint16 {
	return 0
}

func (m *MarketDataHeader) InstrSeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.InstrSeqNumSinceVersion()
}

func (*MarketDataHeader) InstrSeqNumDeprecated() uint16 {
	return 0
}

func (*MarketDataHeader) InstrSeqNumMetaAttribute(meta int) string {
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

func (*MarketDataHeader) InstrSeqNumMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataHeader) InstrSeqNumMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataHeader) InstrSeqNumNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataHeader) TimestampId() uint16 {
	return 4
}

func (*MarketDataHeader) TimestampSinceVersion() uint16 {
	return 0
}

func (m *MarketDataHeader) TimestampInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TimestampSinceVersion()
}

func (*MarketDataHeader) TimestampDeprecated() uint16 {
	return 0
}

func (*MarketDataHeader) TimestampMetaAttribute(meta int) string {
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

func (*MarketDataHeader) TimestampMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataHeader) TimestampMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataHeader) TimestampNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataHeader) InstrumentIdId() uint16 {
	return 5
}

func (*MarketDataHeader) InstrumentIdSinceVersion() uint16 {
	return 0
}

func (m *MarketDataHeader) InstrumentIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.InstrumentIdSinceVersion()
}

func (*MarketDataHeader) InstrumentIdDeprecated() uint16 {
	return 0
}

func (*MarketDataHeader) InstrumentIdMetaAttribute(meta int) string {
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

func (*MarketDataHeader) InstrumentIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataHeader) InstrumentIdMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataHeader) InstrumentIdNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataHeader) FlagsId() uint16 {
	return 6
}

func (*MarketDataHeader) FlagsSinceVersion() uint16 {
	return 0
}

func (m *MarketDataHeader) FlagsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.FlagsSinceVersion()
}

func (*MarketDataHeader) FlagsDeprecated() uint16 {
	return 0
}

func (*MarketDataHeader) FlagsMetaAttribute(meta int) string {
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

func (*MarketDataHeader) ShardIdId() uint16 {
	return 7
}

func (*MarketDataHeader) ShardIdSinceVersion() uint16 {
	return 0
}

func (m *MarketDataHeader) ShardIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ShardIdSinceVersion()
}

func (*MarketDataHeader) ShardIdDeprecated() uint16 {
	return 0
}

func (*MarketDataHeader) ShardIdMetaAttribute(meta int) string {
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

func (*MarketDataHeader) ShardIdMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MarketDataHeader) ShardIdMaxValue() int8 {
	return math.MaxInt8
}

func (*MarketDataHeader) ShardIdNullValue() int8 {
	return math.MinInt8
}

func (*MarketDataHeader) SourceServiceId() uint16 {
	return 8
}

func (*MarketDataHeader) SourceServiceSinceVersion() uint16 {
	return 0
}

func (m *MarketDataHeader) SourceServiceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SourceServiceSinceVersion()
}

func (*MarketDataHeader) SourceServiceDeprecated() uint16 {
	return 0
}

func (*MarketDataHeader) SourceServiceMetaAttribute(meta int) string {
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

func (*MarketDataHeader) SideId() uint16 {
	return 9
}

func (*MarketDataHeader) SideSinceVersion() uint16 {
	return 0
}

func (m *MarketDataHeader) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SideSinceVersion()
}

func (*MarketDataHeader) SideDeprecated() uint16 {
	return 0
}

func (*MarketDataHeader) SideMetaAttribute(meta int) string {
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

func (*MarketDataHeader) TemplateIdId() uint16 {
	return 10
}

func (*MarketDataHeader) TemplateIdSinceVersion() uint16 {
	return 0
}

func (m *MarketDataHeader) TemplateIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TemplateIdSinceVersion()
}

func (*MarketDataHeader) TemplateIdDeprecated() uint16 {
	return 0
}

func (*MarketDataHeader) TemplateIdMetaAttribute(meta int) string {
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

func (*MarketDataHeader) TemplateIdMinValue() uint16 {
	return 0
}

func (*MarketDataHeader) TemplateIdMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MarketDataHeader) TemplateIdNullValue() uint16 {
	return math.MaxUint16
}

func (*MarketDataHeader) BlockLengthId() uint16 {
	return 11
}

func (*MarketDataHeader) BlockLengthSinceVersion() uint16 {
	return 0
}

func (m *MarketDataHeader) BlockLengthInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.BlockLengthSinceVersion()
}

func (*MarketDataHeader) BlockLengthDeprecated() uint16 {
	return 0
}

func (*MarketDataHeader) BlockLengthMetaAttribute(meta int) string {
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

func (*MarketDataHeader) BlockLengthMinValue() uint16 {
	return 0
}

func (*MarketDataHeader) BlockLengthMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MarketDataHeader) BlockLengthNullValue() uint16 {
	return math.MaxUint16
}
