// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type CancelOrReplaceDetails struct {
	OriginalClientOrderId [2]int64
}

func (c *CancelOrReplaceDetails) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := c.RangeCheck(c.SbeSchemaVersion(), c.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	for idx := 0; idx < 2; idx++ {
		if err := _m.WriteInt64(_w, c.OriginalClientOrderId[idx]); err != nil {
			return err
		}
	}
	return nil
}

func (c *CancelOrReplaceDetails) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !c.OriginalClientOrderIdInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			c.OriginalClientOrderId[idx] = c.OriginalClientOrderIdNullValue()
		}
	} else {
		for idx := 0; idx < 2; idx++ {
			if err := _m.ReadInt64(_r, &c.OriginalClientOrderId[idx]); err != nil {
				return err
			}
		}
	}
	if actingVersion > c.SbeSchemaVersion() && blockLength > c.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-c.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := c.RangeCheck(actingVersion, c.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (c *CancelOrReplaceDetails) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if c.OriginalClientOrderIdInActingVersion(actingVersion) {
		for idx := 0; idx < 2; idx++ {
			if c.OriginalClientOrderId[idx] < c.OriginalClientOrderIdMinValue() || c.OriginalClientOrderId[idx] > c.OriginalClientOrderIdMaxValue() {
				return fmt.Errorf("Range check failed on c.OriginalClientOrderId[%d] (%v < %v > %v)", idx, c.OriginalClientOrderIdMinValue(), c.OriginalClientOrderId[idx], c.OriginalClientOrderIdMaxValue())
			}
		}
	}
	return nil
}

func CancelOrReplaceDetailsInit(c *CancelOrReplaceDetails) {
	return
}

func (*CancelOrReplaceDetails) SbeBlockLength() (blockLength uint16) {
	return 16
}

func (*CancelOrReplaceDetails) SbeTemplateId() (templateId uint16) {
	return 212
}

func (*CancelOrReplaceDetails) SbeSchemaId() (schemaId uint16) {
	return 1020
}

func (*CancelOrReplaceDetails) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*CancelOrReplaceDetails) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*CancelOrReplaceDetails) OriginalClientOrderIdId() uint16 {
	return 1
}

func (*CancelOrReplaceDetails) OriginalClientOrderIdSinceVersion() uint16 {
	return 0
}

func (c *CancelOrReplaceDetails) OriginalClientOrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.OriginalClientOrderIdSinceVersion()
}

func (*CancelOrReplaceDetails) OriginalClientOrderIdDeprecated() uint16 {
	return 0
}

func (*CancelOrReplaceDetails) OriginalClientOrderIdMetaAttribute(meta int) string {
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

func (*CancelOrReplaceDetails) OriginalClientOrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*CancelOrReplaceDetails) OriginalClientOrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*CancelOrReplaceDetails) OriginalClientOrderIdNullValue() int64 {
	return math.MinInt64
}
