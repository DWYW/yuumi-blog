package administrator

import (
	"context"
	"yuumi/app/service/administrator/internal/data/mysql/schema"
	"yuumi/pkg/crypto"
)

func (m *Model) UpdatePassword(ctx context.Context, id int64, password string, salt string) (*schema.Administrator, error) {
	return m.UpdateWithID(ctx, id, schema.Administrator{Password: crypto.HmacSha256(salt, password)})
}
