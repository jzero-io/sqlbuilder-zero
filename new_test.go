package sqlbuilder_zero

import (
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	"testing"
)

func Test1(t *testing.T) {
	sqlbuilder.DefaultFlavor = sqlbuilder.PostgreSQL
	sb := sqlbuilder.Select("id", "uuid").From("T_system_dict_key")

	sb.Where(sb.E("id", 1))
	sql, args := sb.Build()
	fmt.Println(sql, args)

	sb = sqlbuilder.Select("*").From("T_system_dict_key")
	sb.Limit(10).Offset((1 - 1) * 10)
	sql, args = sb.Build()
	fmt.Println(sql, args)
}

func Test2(t *testing.T) {
	sb := sqlbuilder.Select("id", "uuid").From("T_system_dict_key")

	sb.Where(sb.E("id", 1))
	sql, args := sb.Build()
	fmt.Println(sql, args)

	sb = sqlbuilder.Select("*").From("T_system_dict_key")
	sb.Limit(10).Offset((1 - 1) * 10)
	sql, args = sb.BuildWithFlavor(sqlbuilder.MySQL)
	fmt.Println(sql, args)

}
