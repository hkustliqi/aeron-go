// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"math"
)

type OmsUserRequestHeader struct {
	GatewayLogPosition int64
	RequestTime        int64
	CorrelationId      int64
	UserId             int64
}

func (o *OmsUserRequestHeader) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, o.GatewayLogPosition); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.RequestTime); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.CorrelationId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.UserId); err != nil {
		return err
	}
	return nil
}

func (o *OmsUserRequestHeader) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !o.GatewayLogPositionInActingVersion(actingVersion) {
		o.GatewayLogPosition = o.GatewayLogPositionNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.GatewayLogPosition); err != nil {
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
	if !o.CorrelationIdInActingVersion(actingVersion) {
		o.CorrelationId = o.CorrelationIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.CorrelationId); err != nil {
			return err
		}
	}
	if !o.UserIdInActingVersion(actingVersion) {
		o.UserId = o.UserIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.UserId); err != nil {
			return err
		}
	}
	return nil
}

func (o *OmsUserRequestHeader) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.GatewayLogPositionInActingVersion(actingVersion) {
		if o.GatewayLogPosition < o.GatewayLogPositionMinValue() || o.GatewayLogPosition > o.GatewayLogPositionMaxValue() {
			return fmt.Errorf("Range check failed on o.GatewayLogPosition (%v < %v > %v)", o.GatewayLogPositionMinValue(), o.GatewayLogPosition, o.GatewayLogPositionMaxValue())
		}
	}
	if o.RequestTimeInActingVersion(actingVersion) {
		if o.RequestTime < o.RequestTimeMinValue() || o.RequestTime > o.RequestTimeMaxValue() {
			return fmt.Errorf("Range check failed on o.RequestTime (%v < %v > %v)", o.RequestTimeMinValue(), o.RequestTime, o.RequestTimeMaxValue())
		}
	}
	if o.CorrelationIdInActingVersion(actingVersion) {
		if o.CorrelationId != o.CorrelationIdNullValue() && (o.CorrelationId < o.CorrelationIdMinValue() || o.CorrelationId > o.CorrelationIdMaxValue()) {
			return fmt.Errorf("Range check failed on o.CorrelationId (%v < %v > %v)", o.CorrelationIdMinValue(), o.CorrelationId, o.CorrelationIdMaxValue())
		}
	}
	if o.UserIdInActingVersion(actingVersion) {
		if o.UserId < o.UserIdMinValue() || o.UserId > o.UserIdMaxValue() {
			return fmt.Errorf("Range check failed on o.UserId (%v < %v > %v)", o.UserIdMinValue(), o.UserId, o.UserIdMaxValue())
		}
	}
	return nil
}

func OmsUserRequestHeaderInit(o *OmsUserRequestHeader) {
	o.CorrelationId = math.MinInt64
	return
}

func (*OmsUserRequestHeader) EncodedLength() int64 {
	return 32
}

func (*OmsUserRequestHeader) GatewayLogPositionMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsUserRequestHeader) GatewayLogPositionMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsUserRequestHeader) GatewayLogPositionNullValue() int64 {
	return math.MinInt64
}

func (*OmsUserRequestHeader) GatewayLogPositionSinceVersion() uint16 {
	return 0
}

func (o *OmsUserRequestHeader) GatewayLogPositionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.GatewayLogPositionSinceVersion()
}

func (*OmsUserRequestHeader) GatewayLogPositionDeprecated() uint16 {
	return 0
}

func (*OmsUserRequestHeader) RequestTimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsUserRequestHeader) RequestTimeMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsUserRequestHeader) RequestTimeNullValue() int64 {
	return math.MinInt64
}

func (*OmsUserRequestHeader) RequestTimeSinceVersion() uint16 {
	return 0
}

func (o *OmsUserRequestHeader) RequestTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.RequestTimeSinceVersion()
}

func (*OmsUserRequestHeader) RequestTimeDeprecated() uint16 {
	return 0
}

func (*OmsUserRequestHeader) CorrelationIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsUserRequestHeader) CorrelationIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsUserRequestHeader) CorrelationIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsUserRequestHeader) CorrelationIdSinceVersion() uint16 {
	return 0
}

func (o *OmsUserRequestHeader) CorrelationIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CorrelationIdSinceVersion()
}

func (*OmsUserRequestHeader) CorrelationIdDeprecated() uint16 {
	return 0
}

func (*OmsUserRequestHeader) UserIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OmsUserRequestHeader) UserIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OmsUserRequestHeader) UserIdNullValue() int64 {
	return math.MinInt64
}

func (*OmsUserRequestHeader) UserIdSinceVersion() uint16 {
	return 0
}

func (o *OmsUserRequestHeader) UserIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.UserIdSinceVersion()
}

func (*OmsUserRequestHeader) UserIdDeprecated() uint16 {
	return 0
}
