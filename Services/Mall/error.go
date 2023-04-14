package Mall
import "errors"

var (
	ErrorMallNotExit   = errors.New("商户不存在")
	ErrorQueryFailed   = errors.New("查询数据失败")
	ErrorInsertFailed  = errors.New("插入数据失败")
)