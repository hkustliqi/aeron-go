// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type TradingEventSnapshotHeader struct {
	Timestamp                                       int64
	LogPosition                                     int64
	LastOrderEventId                                int64
	LastPortfolioBalanceUpdateRequestSequenceNumber int64
}

func (t *TradingEventSnapshotHeader) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := t.RangeCheck(t.SbeSchemaVersion(), t.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, t.Timestamp); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, t.LogPosition); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, t.LastOrderEventId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, t.LastPortfolioBalanceUpdateRequestSequenceNumber); err != nil {
		return err
	}
	return nil
}

func (t *TradingEventSnapshotHeader) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !t.TimestampInActingVersion(actingVersion) {
		t.Timestamp = t.TimestampNullValue()
	} else {
		if err := _m.ReadInt64(_r, &t.Timestamp); err != nil {
			return err
		}
	}
	if !t.LogPositionInActingVersion(actingVersion) {
		t.LogPosition = t.LogPositionNullValue()
	} else {
		if err := _m.ReadInt64(_r, &t.LogPosition); err != nil {
			return err
		}
	}
	if !t.LastOrderEventIdInActingVersion(actingVersion) {
		t.LastOrderEventId = t.LastOrderEventIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &t.LastOrderEventId); err != nil {
			return err
		}
	}
	if !t.LastPortfolioBalanceUpdateRequestSequenceNumberInActingVersion(actingVersion) {
		t.LastPortfolioBalanceUpdateRequestSequenceNumber = t.LastPortfolioBalanceUpdateRequestSequenceNumberNullValue()
	} else {
		if err := _m.ReadInt64(_r, &t.LastPortfolioBalanceUpdateRequestSequenceNumber); err != nil {
			return err
		}
	}
	if actingVersion > t.SbeSchemaVersion() && blockLength > t.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-t.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := t.RangeCheck(actingVersion, t.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (t *TradingEventSnapshotHeader) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if t.TimestampInActingVersion(actingVersion) {
		if t.Timestamp < t.TimestampMinValue() || t.Timestamp > t.TimestampMaxValue() {
			return fmt.Errorf("Range check failed on t.Timestamp (%v < %v > %v)", t.TimestampMinValue(), t.Timestamp, t.TimestampMaxValue())
		}
	}
	if t.LogPositionInActingVersion(actingVersion) {
		if t.LogPosition < t.LogPositionMinValue() || t.LogPosition > t.LogPositionMaxValue() {
			return fmt.Errorf("Range check failed on t.LogPosition (%v < %v > %v)", t.LogPositionMinValue(), t.LogPosition, t.LogPositionMaxValue())
		}
	}
	if t.LastOrderEventIdInActingVersion(actingVersion) {
		if t.LastOrderEventId < t.LastOrderEventIdMinValue() || t.LastOrderEventId > t.LastOrderEventIdMaxValue() {
			return fmt.Errorf("Range check failed on t.LastOrderEventId (%v < %v > %v)", t.LastOrderEventIdMinValue(), t.LastOrderEventId, t.LastOrderEventIdMaxValue())
		}
	}
	if t.LastPortfolioBalanceUpdateRequestSequenceNumberInActingVersion(actingVersion) {
		if t.LastPortfolioBalanceUpdateRequestSequenceNumber < t.LastPortfolioBalanceUpdateRequestSequenceNumberMinValue() || t.LastPortfolioBalanceUpdateRequestSequenceNumber > t.LastPortfolioBalanceUpdateRequestSequenceNumberMaxValue() {
			return fmt.Errorf("Range check failed on t.LastPortfolioBalanceUpdateRequestSequenceNumber (%v < %v > %v)", t.LastPortfolioBalanceUpdateRequestSequenceNumberMinValue(), t.LastPortfolioBalanceUpdateRequestSequenceNumber, t.LastPortfolioBalanceUpdateRequestSequenceNumberMaxValue())
		}
	}
	return nil
}

func TradingEventSnapshotHeaderInit(t *TradingEventSnapshotHeader) {
	return
}

func (*TradingEventSnapshotHeader) SbeBlockLength() (blockLength uint16) {
	return 32
}

func (*TradingEventSnapshotHeader) SbeTemplateId() (templateId uint16) {
	return 1111
}

func (*TradingEventSnapshotHeader) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*TradingEventSnapshotHeader) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*TradingEventSnapshotHeader) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*TradingEventSnapshotHeader) TimestampId() uint16 {
	return 1
}

func (*TradingEventSnapshotHeader) TimestampSinceVersion() uint16 {
	return 0
}

func (t *TradingEventSnapshotHeader) TimestampInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.TimestampSinceVersion()
}

func (*TradingEventSnapshotHeader) TimestampDeprecated() uint16 {
	return 0
}

func (*TradingEventSnapshotHeader) TimestampMetaAttribute(meta int) string {
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

func (*TradingEventSnapshotHeader) TimestampMinValue() int64 {
	return math.MinInt64 + 1
}

func (*TradingEventSnapshotHeader) TimestampMaxValue() int64 {
	return math.MaxInt64
}

func (*TradingEventSnapshotHeader) TimestampNullValue() int64 {
	return math.MinInt64
}

func (*TradingEventSnapshotHeader) LogPositionId() uint16 {
	return 2
}

func (*TradingEventSnapshotHeader) LogPositionSinceVersion() uint16 {
	return 0
}

func (t *TradingEventSnapshotHeader) LogPositionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.LogPositionSinceVersion()
}

func (*TradingEventSnapshotHeader) LogPositionDeprecated() uint16 {
	return 0
}

func (*TradingEventSnapshotHeader) LogPositionMetaAttribute(meta int) string {
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

func (*TradingEventSnapshotHeader) LogPositionMinValue() int64 {
	return math.MinInt64 + 1
}

func (*TradingEventSnapshotHeader) LogPositionMaxValue() int64 {
	return math.MaxInt64
}

func (*TradingEventSnapshotHeader) LogPositionNullValue() int64 {
	return math.MinInt64
}

func (*TradingEventSnapshotHeader) LastOrderEventIdId() uint16 {
	return 3
}

func (*TradingEventSnapshotHeader) LastOrderEventIdSinceVersion() uint16 {
	return 0
}

func (t *TradingEventSnapshotHeader) LastOrderEventIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.LastOrderEventIdSinceVersion()
}

func (*TradingEventSnapshotHeader) LastOrderEventIdDeprecated() uint16 {
	return 0
}

func (*TradingEventSnapshotHeader) LastOrderEventIdMetaAttribute(meta int) string {
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

func (*TradingEventSnapshotHeader) LastOrderEventIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*TradingEventSnapshotHeader) LastOrderEventIdMaxValue() int64 {
	return math.MaxInt64
}

func (*TradingEventSnapshotHeader) LastOrderEventIdNullValue() int64 {
	return math.MinInt64
}

func (*TradingEventSnapshotHeader) LastPortfolioBalanceUpdateRequestSequenceNumberId() uint16 {
	return 4
}

func (*TradingEventSnapshotHeader) LastPortfolioBalanceUpdateRequestSequenceNumberSinceVersion() uint16 {
	return 0
}

func (t *TradingEventSnapshotHeader) LastPortfolioBalanceUpdateRequestSequenceNumberInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.LastPortfolioBalanceUpdateRequestSequenceNumberSinceVersion()
}

func (*TradingEventSnapshotHeader) LastPortfolioBalanceUpdateRequestSequenceNumberDeprecated() uint16 {
	return 0
}

func (*TradingEventSnapshotHeader) LastPortfolioBalanceUpdateRequestSequenceNumberMetaAttribute(meta int) string {
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

func (*TradingEventSnapshotHeader) LastPortfolioBalanceUpdateRequestSequenceNumberMinValue() int64 {
	return math.MinInt64 + 1
}

func (*TradingEventSnapshotHeader) LastPortfolioBalanceUpdateRequestSequenceNumberMaxValue() int64 {
	return math.MaxInt64
}

func (*TradingEventSnapshotHeader) LastPortfolioBalanceUpdateRequestSequenceNumberNullValue() int64 {
	return math.MinInt64
}
