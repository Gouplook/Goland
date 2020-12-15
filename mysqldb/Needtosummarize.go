/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/15 10:04
@Description:待总结

*********************************************/
package mysqldb

// 带分页查询
//maps := model.GetByPage(where, args.GetStart(), args.GetPageSize(), flag)

//
//func (b *BusFrontDisburseDetailModel) GetByPage(where []base.WhereItem,start,pageSize int, orderByType int,field...string) []map[string]interface{}{
//	if len(where) == 0 {
//		return []map[string]interface{}{}
//	}
//	where = append(where,base.WhereItem{b.Field.F_del_status,BUSDISBURSEMODEL_STATUS_NOTDELETE})
//	if len(field) > 0 {
//		b.Model.Field(field)
//	}
//	var orderBy string
//	if orderByType == 1 {
//		orderBy = b.Field.F_disburse_time+" asc"
//	}else if orderByType == 2 {
//		orderBy = b.Field.F_disburse_time+" desc"
//	}else {
//		orderBy = b.Field.F_disburse_id+" desc"
//	}
//	return b.Model.Where(where).Limit(start,pageSize).OrderBy(orderBy).Select()
//}