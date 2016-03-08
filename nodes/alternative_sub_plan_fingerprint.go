// Auto-generated - DO NOT EDIT

package pg_query

func (node AlternativeSubPlan) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlternativeSubPlan")

	if len(node.Subplans.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Subplans.Fingerprint(&subCtx, "Subplans")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("subplans")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.Xpr != nil {
		subCtx := FingerprintSubContext{}
		node.Xpr.Fingerprint(&subCtx, "Xpr")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("xpr")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}