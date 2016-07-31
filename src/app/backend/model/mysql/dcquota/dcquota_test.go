package dcquota

import (
	"fmt"
	mysql "app/backend/common/util/mysql"
	"testing"
)

func Test_NewDcQuota(*testing.T) {
	dcQuota := NewDcQuota(1, 1, 1000, 10, 20, 1, 2, 100, 10, 0, 1, "1000", "add dcquota")
	fmt.Printf("%v\n", dcQuota)
}

func Test_QueryDcQuotaById(*testing.T) {
	mysql.NewMysqlClient(mysql.DB_HOST, mysql.DB_USER, mysql.DB_PASSWORD, mysql.DB_NAME, mysql.MAX_POOL_SIZE)
	mysql.MysqlInstance().Open()

	dc := new(DcQuota)
	dc.QueryDcQuotaById(1)
	fmt.Printf("%v\n", dc)
}

func Test_InsertDcQuota(*testing.T) {
	mysql.NewMysqlClient(mysql.DB_HOST, mysql.DB_USER, mysql.DB_PASSWORD, mysql.DB_NAME, mysql.MAX_POOL_SIZE)
	mysql.MysqlInstance().Open()


	dc := NewDcQuota(1, 1, 1000, 10, 20, 1, 2, 100, 10, 0, 1, "1000", "add dcquota")
	dc.InsertDcQuota(2)

}

func Test_UpdateDcQuota(*testing.T) {
	mysql.NewMysqlClient(mysql.DB_HOST, mysql.DB_USER, mysql.DB_PASSWORD, mysql.DB_NAME, mysql.MAX_POOL_SIZE)
	mysql.MysqlInstance().Open()

	dc := new(DcQuota)
	dc.QueryDcQuotaById(2)

	dc.PodNumLimit = 1100
	dc.UpdateDcQuota(2)

	dc.QueryDcQuotaById(2)

	fmt.Printf("%v\n", dc)
}

func Test_DeleteQuota(*testing.T) {

	mysql.NewMysqlClient(mysql.DB_HOST, mysql.DB_USER, mysql.DB_PASSWORD, mysql.DB_NAME, mysql.MAX_POOL_SIZE)
	mysql.MysqlInstance().Open()

	dc := new(DcQuota)
	dc.QueryDcQuotaById(2)
	dc.DeleteDcQuota(2)
}

func Test_EncodeJson_DecodeJson(*testing.T) {

	dc := NewDcQuota(1, 1, 1000, 10, 20, 1, 2, 100, 10, 0, 1, "1000", "add dcquota")
	fmt.Printf("%s\n", dc.EncodeJson())

	d := new(DcQuota)
	d.DecodeJson(dc.EncodeJson())
	fmt.Printf("%v\n", d)
}