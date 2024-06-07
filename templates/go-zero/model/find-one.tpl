func (m *default{{.upperStartCamelObject}}Model) FindOne(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error) {
	{{if .withCache}}{{.cacheKey}}
	var resp {{.upperStartCamelObject}}
	err := m.QueryRowCtx(ctx, &resp, {{.cacheKeyVariable}}, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
	    sb := sqlbuilder.Select({{.lowerStartCamelObject}}Rows).From(m.table)
	    sb.Where(sb.EQ("{{.originalPrimaryKey}}", {{.lowerStartCamelPrimaryKey}}))
        sql, args := sb.Build()
		return conn.QueryRowCtx(ctx, v, sql, args...)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}{{else}}sb := sqlbuilder.Select({{.lowerStartCamelObject}}Rows).From(m.table)
	sb.Where(sb.EQ("{{.originalPrimaryKey}}", {{.lowerStartCamelPrimaryKey}}))
	sb.Limit(1)
	sql, args := sb.Build()
	var resp {{.upperStartCamelObject}}
	err := m.conn.QueryRowCtx(ctx, &resp, sql, args...)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}{{end}}
}
