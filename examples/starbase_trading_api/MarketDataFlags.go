// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"io"
)

type MarketDataFlags [8]bool
type MarketDataFlagsChoiceValue uint8
type MarketDataFlagsChoiceValues struct {
	IsTransactionStart MarketDataFlagsChoiceValue
	IsTransactionEnd   MarketDataFlagsChoiceValue
	ClearBook          MarketDataFlagsChoiceValue
}

var MarketDataFlagsChoice = MarketDataFlagsChoiceValues{0, 1, 7}

func (m *MarketDataFlags) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	var wireval uint8 = 0
	for k, v := range m {
		if v {
			wireval |= (1 << uint(k))
		}
	}
	return _m.WriteUint8(_w, wireval)
}

func (m *MarketDataFlags) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	var wireval uint8

	if err := _m.ReadUint8(_r, &wireval); err != nil {
		return err
	}

	var idx uint
	for idx = 0; idx < 8; idx++ {
		m[idx] = (wireval & (1 << idx)) > 0
	}
	return nil
}

func (MarketDataFlags) EncodedLength() int64 {
	return 1
}

func (*MarketDataFlags) isTransactionStartSinceVersion() uint16 {
	return 0
}

func (m *MarketDataFlags) isTransactionStartInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.isTransactionStartSinceVersion()
}

func (*MarketDataFlags) isTransactionStartDeprecated() uint16 {
	return 0
}

func (*MarketDataFlags) isTransactionEndSinceVersion() uint16 {
	return 0
}

func (m *MarketDataFlags) isTransactionEndInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.isTransactionEndSinceVersion()
}

func (*MarketDataFlags) isTransactionEndDeprecated() uint16 {
	return 0
}

func (*MarketDataFlags) clearBookSinceVersion() uint16 {
	return 0
}

func (m *MarketDataFlags) clearBookInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.clearBookSinceVersion()
}

func (*MarketDataFlags) clearBookDeprecated() uint16 {
	return 0
}
