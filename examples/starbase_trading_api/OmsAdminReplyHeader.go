// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"math"
)

type OmsAdminReplyHeader struct {
	CorrelationId int64
	RequestTime   int64
	LogPosition   int64
}

func (o *OmsAdminReplyHeader) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, o.CorrelationId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.RequestTime); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.LogPosition); err != nil {
		return err
	}
	return nil
}

func (o *OmsAdminReplyHeader) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !o.CorrelationIdInActingVersion(actingVersion) {
		o.CorrelationId = o.CorrelationIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.CorrelationId); err != nil {
			return err
		}
	}
	if !o.RequestTimeInActingVersion(actingVersion) {
		o.RequestTime = o.RequestTimeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.RequestTime); err != nil {
			return err
		}
	}
	if !o.LogPositionInActingVersion(actingVersion) {
		o.LogPosition = o.LogPositionNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.LogPosition); err != nil {
			return err
		}
	}
	return nil
}

func (o *OmsAdminReplyHeader) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.CorrelationIdInActingVersion(actingVersion) {
		if o.CorrelationId < o.CorrelationIdMinValue() || o.CorrelationId > o.CorrelationIdMaxValue() {
			return fmt.Errorf("Range check failed on o.CorrelationId (%v < %v > %v)", o.CorrelationIdMinValue(), o.CorrelationId, o.CorrelationIdMaxValue())
		}
	}
	if o.RequestTimeInActingVersion(actingVersion) {
		if o.RequestTime < o.RequestTimeMinValue() || o.RequestTime > o.RequestTimeMaxValue() {
			return fmt.Errorf("Range check failed on o.RequestTime (%v < %v > %v)", o.RequestTimeMinValue(), o.RequestTime, o.RequestTimeMaxValue())
		}
	}
	if o.LogPositionInActingVersion(actingVersion) {
		if o.LogPosition < o.LogPositionMinValue() || o.LogPosition > o.LogPositionMaxValue() {
			return fmt.Errorf("Range check failed on o.LogPosition (%v < %v > %v)", o.LogPositionMinValue(), o.LogPosition, o.LogPositionMaxValue())
		}
	}
	return nil
}

func OmsAdminReplyHeaderInit(o *OmsAdminReplyHeader) {
	return
}

func (*OmsAdminReplyHeader) EncodedLength() int64 {
	return 24
}

func (*OmsAdminReplyHeader) CorrelationIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsAdminReplyHeader) CorrelationIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsAdminReplyHeader) CorrelationIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsAdminReplyHeader) CorrelationIdSinceVersion() uint16 {
	return 0
}

func (o *OmsAdminReplyHeader) CorrelationIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CorrelationIdSinceVersion()
}

func (*OmsAdminReplyHeader) CorrelationIdDeprecated() uint16 {
	return 0
}

func (*OmsAdminReplyHeader) RequestTimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsAdminReplyHeader) RequestTimeMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsAdminReplyHeader) RequestTimeNullValue() int64 {
	return math.MinInt64
}

func (*OmsAdminReplyHeader) RequestTimeSinceVersion() uint16 {
	return 0
}

func (o *OmsAdminReplyHeader) RequestTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.RequestTimeSinceVersion()
}

func (*OmsAdminReplyHeader) RequestTimeDeprecated() uint16 {
	return 0
}

func (*OmsAdminReplyHeader) LogPositionMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsAdminReplyHeader) LogPositionMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsAdminReplyHeader) LogPositionNullValue() int64 {
	return math.MinInt64
}

func (*OmsAdminReplyHeader) LogPositionSinceVersion() uint16 {
	return 0
}

func (o *OmsAdminReplyHeader) LogPositionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LogPositionSinceVersion()
}

func (*OmsAdminReplyHeader) LogPositionDeprecated() uint16 {
	return 0
}
