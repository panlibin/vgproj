package rank

const (
	sqlLoadRank      = "select rank_type,obj_id,val,change_time,extra,name,server_id,lev,title_id from global_rank"
	sqlInsertRankObj = "insert into global_rank values(?,?,?,?,?,?,?,?,?)"
	sqlUpdateRankObj = "update global_rank set val=?,change_time=?,extra=?,name=?,server_id=?,lev=?,title_id=? where rank_type=? and obj_id=?"
	sqlDeleteRankObj = "delete from global_rank where rank_type=? and obj_id=?"
)
