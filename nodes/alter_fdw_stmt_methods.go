// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node AlterFdwStmt) MarshalJSON() ([]byte, error) {
	type AlterFdwStmtMarshalAlias AlterFdwStmt
	return json.Marshal(map[string]interface{}{
		"ALTERFDWSTMT": (*AlterFdwStmtMarshalAlias)(&node),
	})
}

func (node *AlterFdwStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node AlterFdwStmt) Deparse() string {
	panic("Not Implemented")
}