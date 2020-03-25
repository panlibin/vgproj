package mail

const sqlLoadGlobalMail = "select global_mail_id,source,source_ext,ts,first_type,second_type,title,title_params,content,content_params,attachments," +
	"vip_lev_limit from global_mail_list"
const sqlInsertGlobalMail = "insert into global_mail_list values(?,?,?,?,?,?,?,?,?,?,?,?)"
const sqlDeleteGlobalMail = "delete from global_mail_list where global_mail_id=?"
