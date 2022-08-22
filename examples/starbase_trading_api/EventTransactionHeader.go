// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type EventTransactionHeader struct {
	Timestamp     int64
	LogPosition   int64
	EventSeqNum   int64
	ShardId       int8
	SourceService SourceServiceEnum
	Flags         EventTransactionFlags
	TemplateId    uint16
	BlockLength   uint16
}

func (e *EventTransactionHeader) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := e.RangeCheck(e.SbeSchemaVersion(), e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, e.Timestamp); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.LogPosition); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.EventSeqNum); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, e.ShardId); err != nil {
		return err
	}
	if err := e.SourceService.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.Flags.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, e.TemplateId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, e.BlockLength); err != nil {
		return err
	}
	return nil
}

func (e *EventTransactionHeader) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !e.TimestampInActingVersion(actingVersion) {
		e.Timestamp = e.TimestampNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.Timestamp); err != nil {
			return err
		}
	}
	if !e.LogPositionInActingVersion(actingVersion) {
		e.LogPosition = e.LogPositionNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.LogPosition); err != nil {
			return err
		}
	}
	if !e.EventSeqNumInActingVersion(actingVersion) {
		e.EventSeqNum = e.EventSeqNumNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.EventSeqNum); err != nil {
			return err
		}
	}
	if !e.ShardIdInActingVersion(actingVersion) {
		e.ShardId = e.ShardIdNullValue()
	} else {
		if err := _m.ReadInt8(_r, &e.ShardId); err != nil {
			return err
		}
	}
	if e.SourceServiceInActingVersion(actingVersion) {
		if err := e.SourceService.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.FlagsInActingVersion(actingVersion) {
		if err := e.Flags.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.TemplateIdInActingVersion(actingVersion) {
		e.TemplateId = e.TemplateIdNullValue()
	} else {
		if err := _m.ReadUint16(_r, &e.TemplateId); err != nil {
			return err
		}
	}
	if !e.BlockLengthInActingVersion(actingVersion) {
		e.BlockLength = e.BlockLengthNullValue()
	} else {
		if err := _m.ReadUint16(_r, &e.BlockLength); err != nil {
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

func (e *EventTransactionHeader) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.TimestampInActingVersion(actingVersion) {
		if e.Timestamp < e.TimestampMinValue() || e.Timestamp > e.TimestampMaxValue() {
			return fmt.Errorf("Range check failed on e.Timestamp (%v < %v > %v)", e.TimestampMinValue(), e.Timestamp, e.TimestampMaxValue())
		}
	}
	if e.LogPositionInActingVersion(actingVersion) {
		if e.LogPosition < e.LogPositionMinValue() || e.LogPosition > e.LogPositionMaxValue() {
			return fmt.Errorf("Range check failed on e.LogPosition (%v < %v > %v)", e.LogPositionMinValue(), e.LogPosition, e.LogPositionMaxValue())
		}
	}
	if e.EventSeqNumInActingVersion(actingVersion) {
		if e.EventSeqNum < e.EventSeqNumMinValue() || e.EventSeqNum > e.EventSeqNumMaxValue() {
			return fmt.Errorf("Range check failed on e.EventSeqNum (%v < %v > %v)", e.EventSeqNumMinValue(), e.EventSeqNum, e.EventSeqNumMaxValue())
		}
	}
	if e.ShardIdInActingVersion(actingVersion) {
		if e.ShardId < e.ShardIdMinValue() || e.ShardId > e.ShardIdMaxValue() {
			return fmt.Errorf("Range check failed on e.ShardId (%v < %v > %v)", e.ShardIdMinValue(), e.ShardId, e.ShardIdMaxValue())
		}
	}
	if err := e.SourceService.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if e.TemplateIdInActingVersion(actingVersion) {
		if e.TemplateId < e.TemplateIdMinValue() || e.TemplateId > e.TemplateIdMaxValue() {
			return fmt.Errorf("Range check failed on e.TemplateId (%v < %v > %v)", e.TemplateIdMinValue(), e.TemplateId, e.TemplateIdMaxValue())
		}
	}
	if e.BlockLengthInActingVersion(actingVersion) {
		if e.BlockLength < e.BlockLengthMinValue() || e.BlockLength > e.BlockLengthMaxValue() {
			return fmt.Errorf("Range check failed on e.BlockLength (%v < %v > %v)", e.BlockLengthMinValue(), e.BlockLength, e.BlockLengthMaxValue())
		}
	}
	return nil
}

func EventTransactionHeaderInit(e *EventTransactionHeader) {
	return
}

func (*EventTransactionHeader) SbeBlockLength() (blockLength uint16) {
	return 32
}

func (*EventTransactionHeader) SbeTemplateId() (templateId uint16) {
	return 100
}

func (*EventTransactionHeader) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*EventTransactionHeader) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*EventTransactionHeader) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*EventTransactionHeader) TimestampId() uint16 {
	return 1
}

func (*EventTransactionHeader) TimestampSinceVersion() uint16 {
	return 0
}

func (e *EventTransactionHeader) TimestampInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TimestampSinceVersion()
}

func (*EventTransactionHeader) TimestampDeprecated() uint16 {
	return 0
}

func (*EventTransactionHeader) TimestampMetaAttribute(meta int) string {
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

func (*EventTransactionHeader) TimestampMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EventTransactionHeader) TimestampMaxValue() int64 {
	return math.MaxInt64
}

func (*EventTransactionHeader) TimestampNullValue() int64 {
	return math.MinInt64
}

func (*EventTransactionHeader) LogPositionId() uint16 {
	return 2
}

func (*EventTransactionHeader) LogPositionSinceVersion() uint16 {
	return 0
}

func (e *EventTransactionHeader) LogPositionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LogPositionSinceVersion()
}

func (*EventTransactionHeader) LogPositionDeprecated() uint16 {
	return 0
}

func (*EventTransactionHeader) LogPositionMetaAttribute(meta int) string {
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

func (*EventTransactionHeader) LogPositionMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EventTransactionHeader) LogPositionMaxValue() int64 {
	return math.MaxInt64
}

func (*EventTransactionHeader) LogPositionNullValue() int64 {
	return math.MinInt64
}

func (*EventTransactionHeader) EventSeqNumId() uint16 {
	return 3
}

func (*EventTransactionHeader) EventSeqNumSinceVersion() uint16 {
	return 0
}

func (e *EventTransactionHeader) EventSeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.EventSeqNumSinceVersion()
}

func (*EventTransactionHeader) EventSeqNumDeprecated() uint16 {
	return 0
}

func (*EventTransactionHeader) EventSeqNumMetaAttribute(meta int) string {
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

func (*EventTransactionHeader) EventSeqNumMinValue() int64 {
	return math.MinInt64 + 1
}

func (*EventTransactionHeader) EventSeqNumMaxValue() int64 {
	return math.MaxInt64
}

func (*EventTransactionHeader) EventSeqNumNullValue() int64 {
	return math.MinInt64
}

func (*EventTransactionHeader) ShardIdId() uint16 {
	return 4
}

func (*EventTransactionHeader) ShardIdSinceVersion() uint16 {
	return 0
}

func (e *EventTransactionHeader) ShardIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ShardIdSinceVersion()
}

func (*EventTransactionHeader) ShardIdDeprecated() uint16 {
	return 0
}

func (*EventTransactionHeader) ShardIdMetaAttribute(meta int) string {
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

func (*EventTransactionHeader) ShardIdMinValue() int8 {
	return math.MinInt8 + 1
}

func (*EventTransactionHeader) ShardIdMaxValue() int8 {
	return math.MaxInt8
}

func (*EventTransactionHeader) ShardIdNullValue() int8 {
	return math.MinInt8
}

func (*EventTransactionHeader) SourceServiceId() uint16 {
	return 5
}

func (*EventTransactionHeader) SourceServiceSinceVersion() uint16 {
	return 0
}

func (e *EventTransactionHeader) SourceServiceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SourceServiceSinceVersion()
}

func (*EventTransactionHeader) SourceServiceDeprecated() uint16 {
	return 0
}

func (*EventTransactionHeader) SourceServiceMetaAttribute(meta int) string {
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

func (*EventTransactionHeader) FlagsId() uint16 {
	return 6
}

func (*EventTransactionHeader) FlagsSinceVersion() uint16 {
	return 0
}

func (e *EventTransactionHeader) FlagsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FlagsSinceVersion()
}

func (*EventTransactionHeader) FlagsDeprecated() uint16 {
	return 0
}

func (*EventTransactionHeader) FlagsMetaAttribute(meta int) string {
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

func (*EventTransactionHeader) TemplateIdId() uint16 {
	return 7
}

func (*EventTransactionHeader) TemplateIdSinceVersion() uint16 {
	return 0
}

func (e *EventTransactionHeader) TemplateIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TemplateIdSinceVersion()
}

func (*EventTransactionHeader) TemplateIdDeprecated() uint16 {
	return 0
}

func (*EventTransactionHeader) TemplateIdMetaAttribute(meta int) string {
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

func (*EventTransactionHeader) TemplateIdMinValue() uint16 {
	return 0
}

func (*EventTransactionHeader) TemplateIdMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*EventTransactionHeader) TemplateIdNullValue() uint16 {
	return math.MaxUint16
}

func (*EventTransactionHeader) BlockLengthId() uint16 {
	return 8
}

func (*EventTransactionHeader) BlockLengthSinceVersion() uint16 {
	return 0
}

func (e *EventTransactionHeader) BlockLengthInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.BlockLengthSinceVersion()
}

func (*EventTransactionHeader) BlockLengthDeprecated() uint16 {
	return 0
}

func (*EventTransactionHeader) BlockLengthMetaAttribute(meta int) string {
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

func (*EventTransactionHeader) BlockLengthMinValue() uint16 {
	return 0
}

func (*EventTransactionHeader) BlockLengthMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*EventTransactionHeader) BlockLengthNullValue() uint16 {
	return math.MaxUint16
}
