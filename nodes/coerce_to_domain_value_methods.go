// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CoerceToDomainValue) MarshalJSON() ([]byte, error) {
	type CoerceToDomainValueMarshalAlias CoerceToDomainValue
	return json.Marshal(map[string]interface{}{
		"COERCETODOMAINVALUE": (*CoerceToDomainValueMarshalAlias)(&node),
	})
}

func (node *CoerceToDomainValue) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CoerceToDomainValue) Deparse() string {
	panic("Not Implemented")
}