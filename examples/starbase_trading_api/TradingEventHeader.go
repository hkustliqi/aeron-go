// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"math"
)

type TradingEventHeader struct {
	ClusterSessionId   int64
	GatewayLogPosition int64
	RequestTime        int64
	CorrelationId      int64
}

func (t *TradingEventHeader) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, t.ClusterSessionId); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, t.GatewayLogPosition); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, t.RequestTime); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, t.CorrelationId); err != nil {
		return err
	}
	return nil
}

func (t *TradingEventHeader) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !t.ClusterSessionIdInActingVersion(actingVersion) {
		t.ClusterSessionId = t.ClusterSessionIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &t.ClusterSessionId); err != nil {
			return err
		}
	}
	if !t.GatewayLogPositionInActingVersion(actingVersion) {
		t.GatewayLogPosition = t.GatewayLogPositionNullValue()
	} else {
		if err := _m.ReadInt64(_r, &t.GatewayLogPosition); err != nil {
			return err
		}
	}
	if !t.RequestTimeInActingVersion(actingVersion) {
		t.RequestTime = t.RequestTimeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &t.RequestTime); err != nil {
			return err
		}
	}
	if !t.CorrelationIdInActingVersion(actingVersion) {
		t.CorrelationId = t.CorrelationIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &t.CorrelationId); err != nil {
			return err
		}
	}
	return nil
}

func (t *TradingEventHeader) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if t.ClusterSessionIdInActingVersion(actingVersion) {
		if t.ClusterSessionId < t.ClusterSessionIdMinValue() || t.ClusterSessionId > t.ClusterSessionIdMaxValue() {
			return fmt.Errorf("Range check failed on t.ClusterSessionId (%v < %v > %v)", t.ClusterSessionIdMinValue(), t.ClusterSessionId, t.ClusterSessionIdMaxValue())
		}
	}
	if t.GatewayLogPositionInActingVersion(actingVersion) {
		if t.GatewayLogPosition < t.GatewayLogPositionMinValue() || t.GatewayLogPosition > t.GatewayLogPositionMaxValue() {
			return fmt.Errorf("Range check failed on t.GatewayLogPosition (%v < %v > %v)", t.GatewayLogPositionMinValue(), t.GatewayLogPosition, t.GatewayLogPositionMaxValue())
		}
	}
	if t.RequestTimeInActingVersion(actingVersion) {
		if t.RequestTime != t.RequestTimeNullValue() && (t.RequestTime < t.RequestTimeMinValue() || t.RequestTime > t.RequestTimeMaxValue()) {
			return fmt.Errorf("Range check failed on t.RequestTime (%v < %v > %v)", t.RequestTimeMinValue(), t.RequestTime, t.RequestTimeMaxValue())
		}
	}
	if t.CorrelationIdInActingVersion(actingVersion) {
		if t.CorrelationId != t.CorrelationIdNullValue() && (t.CorrelationId < t.CorrelationIdMinValue() || t.CorrelationId > t.CorrelationIdMaxValue()) {
			return fmt.Errorf("Range check failed on t.CorrelationId (%v < %v > %v)", t.CorrelationIdMinValue(), t.CorrelationId, t.CorrelationIdMaxValue())
		}
	}
	return nil
}

func TradingEventHeaderInit(t *TradingEventHeader) {
	t.RequestTime = math.MinInt64
	t.CorrelationId = math.MinInt64
	return
}

func (*TradingEventHeader) EncodedLength() int64 {
	return 32
}

func (*TradingEventHeader) ClusterSessionIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*TradingEventHeader) ClusterSessionIdMaxValue() int64 {
	return math.MaxInt64
}

func (*TradingEventHeader) ClusterSessionIdNullValue() int64 {
	return math.MinInt64
}

func (*TradingEventHeader) ClusterSessionIdSinceVersion() uint16 {
	return 0
}

func (t *TradingEventHeader) ClusterSessionIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.ClusterSessionIdSinceVersion()
}

func (*TradingEventHeader) ClusterSessionIdDeprecated() uint16 {
	return 0
}

func (*TradingEventHeader) GatewayLogPositionMinValue() int64 {
	return math.MinInt64 + 1
}

func (*TradingEventHeader) GatewayLogPositionMaxValue() int64 {
	return math.MaxInt64
}

func (*TradingEventHeader) GatewayLogPositionNullValue() int64 {
	return math.MinInt64
}

func (*TradingEventHeader) GatewayLogPositionSinceVersion() uint16 {
	return 0
}

func (t *TradingEventHeader) GatewayLogPositionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.GatewayLogPositionSinceVersion()
}

func (*TradingEventHeader) GatewayLogPositionDeprecated() uint16 {
	return 0
}

func (*TradingEventHeader) RequestTimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*TradingEventHeader) RequestTimeMaxValue() int64 {
	return math.MaxInt64
}

func (*TradingEventHeader) RequestTimeNullValue() int64 {
	return math.MinInt64
}

func (*TradingEventHeader) RequestTimeSinceVersion() uint16 {
	return 0
}

func (t *TradingEventHeader) RequestTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.RequestTimeSinceVersion()
}

func (*TradingEventHeader) RequestTimeDeprecated() uint16 {
	return 0
}

func (*TradingEventHeader) CorrelationIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*TradingEventHeader) CorrelationIdMaxValue() int64 {
	return math.MaxInt64
}

func (*TradingEventHeader) CorrelationIdNullValue() int64 {
	return math.MinInt64
}

func (*TradingEventHeader) CorrelationIdSinceVersion() uint16 {
	return 0
}

func (t *TradingEventHeader) CorrelationIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.CorrelationIdSinceVersion()
}

func (*TradingEventHeader) CorrelationIdDeprecated() uint16 {
	return 0
}
