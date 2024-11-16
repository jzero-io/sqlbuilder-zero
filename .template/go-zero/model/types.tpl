type (
	{{.lowerStartCamelObject}}Model interface{
		{{.method}}

		// custom interface generated by jzero
	    BulkInsert(ctx context.Context, session sqlx.Session, datas []*{{.upperStartCamelObject}}) error
        FindByCondition(ctx context.Context, session sqlx.Session, conds ...condition.Condition) ([]*{{.upperStartCamelObject}}, error)
        FindOneByCondition(ctx context.Context, session sqlx.Session, conds ...condition.Condition) (*{{.upperStartCamelObject}}, error)
        PageByCondition(ctx context.Context, session sqlx.Session, conds ...condition.Condition) ([]*{{.upperStartCamelObject}}, int64 ,error)
        UpdateFieldsByCondition(ctx context.Context, session sqlx.Session, field map[string]any, conds ...condition.Condition) error
        DeleteByCondition(ctx context.Context, session sqlx.Session, conds ...condition.Condition) error
	}

	default{{.upperStartCamelObject}}Model struct {
		{{if .withCache}}sqlc.CachedConn{{else}}conn sqlx.SqlConn{{end}}
		table string
	}

	{{.upperStartCamelObject}} struct {
		{{.fields}}
	}
)
