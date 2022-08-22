// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type LastOrderEventId struct {
	LastOrderEventId int64
	UserId           int64
}

func (l *LastOrderEventId) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := l.RangeCheck(l.SbeSchemaVersion(), l.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, l.LastOrderEventId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, l.UserId); err != nil {
		return err
	}
	return nil
}

func (l *LastOrderEventId) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !l.LastOrderEventIdInActingVersion(actingVersion) {
		l.LastOrderEventId = l.LastOrderEventIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &l.LastOrderEventId); err != nil {
			return err
		}
	}
	if !l.UserIdInActingVersion(actingVersion) {
		l.UserId = l.UserIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &l.UserId); err != nil {
			return err
		}
	}
	if actingVersion > l.SbeSchemaVersion() && blockLength > l.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-l.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := l.RangeCheck(actingVersion, l.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (l *LastOrderEventId) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if l.LastOrderEventIdInActingVersion(actingVersion) {
		if l.LastOrderEventId < l.LastOrderEventIdMinValue() || l.LastOrderEventId > l.LastOrderEventIdMaxValue() {
			return fmt.Errorf("Range check failed on l.LastOrderEventId (%v < %v > %v)", l.LastOrderEventIdMinValue(), l.LastOrderEventId, l.LastOrderEventIdMaxValue())
		}
	}
	if l.UserIdInActingVersion(actingVersion) {
		if l.UserId < l.UserIdMinValue() || l.UserId > l.UserIdMaxValue() {
			return fmt.Errorf("Range check failed on l.UserId (%v < %v > %v)", l.UserIdMinValue(), l.UserId, l.UserIdMaxValue())
		}
	}
	return nil
}

func LastOrderEventIdInit(l *LastOrderEventId) {
	return
}

func (*LastOrderEventId) SbeBlockLength() (blockLength uint16) {
	return 16
}

func (*LastOrderEventId) SbeTemplateId() (templateId uint16) {
	return 1103
}

func (*LastOrderEventId) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*LastOrderEventId) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*LastOrderEventId) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*LastOrderEventId) LastOrderEventIdId() uint16 {
	return 1
}

func (*LastOrderEventId) LastOrderEventIdSinceVersion() uint16 {
	return 0
}

func (l *LastOrderEventId) LastOrderEventIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.LastOrderEventIdSinceVersion()
}

func (*LastOrderEventId) LastOrderEventIdDeprecated() uint16 {
	return 0
}

func (*LastOrderEventId) LastOrderEventIdMetaAttribute(meta int) string {
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

func (*LastOrderEventId) LastOrderEventIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*LastOrderEventId) LastOrderEventIdMaxValue() int64 {
	return math.MaxInt64
}

func (*LastOrderEventId) LastOrderEventIdNullValue() int64 {
	return math.MinInt64
}

func (*LastOrderEventId) UserIdId() uint16 {
	return 2
}

func (*LastOrderEventId) UserIdSinceVersion() uint16 {
	return 0
}

func (l *LastOrderEventId) UserIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.UserIdSinceVersion()
}

func (*LastOrderEventId) UserIdDeprecated() uint16 {
	return 0
}

func (*LastOrderEventId) UserIdMetaAttribute(meta int) string {
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

func (*LastOrderEventId) UserIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*LastOrderEventId) UserIdMaxValue() int64 {
	return math.MaxInt64
}

func (*LastOrderEventId) UserIdNullValue() int64 {
	return math.MinInt64
}
