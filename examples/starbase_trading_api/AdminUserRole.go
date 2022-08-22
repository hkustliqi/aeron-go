// Generated SBE (Simple Binary Encoding) message codec

package starbase_trading_api

import (
	"fmt"
	"io"
	"reflect"
)

type AdminUserRoleEnum int8
type AdminUserRoleValues struct {
	EXCHANGE      AdminUserRoleEnum
	EXCHANGE_READ AdminUserRoleEnum
	NullValue     AdminUserRoleEnum
}

var AdminUserRole = AdminUserRoleValues{1, 2, -128}

func (a AdminUserRoleEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt8(_w, int8(a)); err != nil {
		return err
	}
	return nil
}

func (a *AdminUserRoleEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadInt8(_r, (*int8)(a)); err != nil {
		return err
	}
	return nil
}

func (a AdminUserRoleEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(AdminUserRole)
	for idx := 0; idx < value.NumField(); idx++ {
		if a == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on AdminUserRole, unknown enumeration value %d", a)
}

func (*AdminUserRoleEnum) EncodedLength() int64 {
	return 1
}

func (*AdminUserRoleEnum) EXCHANGESinceVersion() uint16 {
	return 0
}

func (a *AdminUserRoleEnum) EXCHANGEInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.EXCHANGESinceVersion()
}

func (*AdminUserRoleEnum) EXCHANGEDeprecated() uint16 {
	return 0
}

func (*AdminUserRoleEnum) EXCHANGE_READSinceVersion() uint16 {
	return 0
}

func (a *AdminUserRoleEnum) EXCHANGE_READInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.EXCHANGE_READSinceVersion()
}

func (*AdminUserRoleEnum) EXCHANGE_READDeprecated() uint16 {
	return 0
}
