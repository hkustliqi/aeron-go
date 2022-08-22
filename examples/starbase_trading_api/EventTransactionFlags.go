// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"io"
)

type EventTransactionFlags [16]bool
type EventTransactionFlagsChoiceValue uint8
type EventTransactionFlagsChoiceValues struct {
	IsTransactionStart EventTransactionFlagsChoiceValue
	IsTransactionEnd   EventTransactionFlagsChoiceValue
}

var EventTransactionFlagsChoice = EventTransactionFlagsChoiceValues{0, 1}

func (e *EventTransactionFlags) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	var wireval uint16 = 0
	for k, v := range e {
		if v {
			wireval |= (1 << uint(k))
		}
	}
	return _m.WriteUint16(_w, wireval)
}

func (e *EventTransactionFlags) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	var wireval uint16

	if err := _m.ReadUint16(_r, &wireval); err != nil {
		return err
	}

	var idx uint
	for idx = 0; idx < 16; idx++ {
		e[idx] = (wireval & (1 << idx)) > 0
	}
	return nil
}

func (EventTransactionFlags) EncodedLength() int64 {
	return 2
}

func (*EventTransactionFlags) isTransactionStartSinceVersion() uint16 {
	return 0
}

func (e *EventTransactionFlags) isTransactionStartInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.isTransactionStartSinceVersion()
}

func (*EventTransactionFlags) isTransactionStartDeprecated() uint16 {
	return 0
}

func (*EventTransactionFlags) isTransactionEndSinceVersion() uint16 {
	return 0
}

func (e *EventTransactionFlags) isTransactionEndInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.isTransactionEndSinceVersion()
}

func (*EventTransactionFlags) isTransactionEndDeprecated() uint16 {
	return 0
}
