// Copyright 2017 gf Author(https://github.com/dekinsq/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/dekinsq/gf.

package gdb

import (
	"database/sql"
	"fmt"
	"github.com/dekinsq/gf/os/gtime"
)

// Delete does "DELETE FROM ... " statement for the model.
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
func (m *Model) Delete(where ...interface{}) (result sql.Result, err error) {
	if len(where) > 0 {
		return m.Where(where[0], where[1:]...).Delete()
	}
	defer func() {
		if err == nil {
			m.checkAndRemoveCache()
		}
	}()
	var (
		fieldNameDelete                               = m.getSoftFieldNameDeleted()
		conditionWhere, conditionExtra, conditionArgs = m.formatCondition(false)
	)
	// Soft deleting.
	if !m.unscoped && fieldNameDelete != "" {
		return m.db.DoUpdate(
			m.getLink(true),
			m.tables,
			fmt.Sprintf(`%s=?`, m.db.QuoteString(fieldNameDelete)),
			conditionWhere+conditionExtra,
			append([]interface{}{gtime.Now().String()}, conditionArgs...),
		)
	}
	return m.db.DoDelete(m.getLink(true), m.tables, conditionWhere+conditionExtra, conditionArgs...)
}
