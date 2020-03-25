package player

// data
const sqlUpdateLogin = "update player_data set last_login_ip=?,last_login_ts=?,last_logout_ts=? where player_id=?"
const sqlUpdateLogout = "update player_data set last_logout_ts=? where player_id=?"
const sqlUpdateDailyRefresh = "update player_data set last_daily_refresh_ts=? where player_id=?"
const sqlUpdateName = "update player_data set `name`=? where player_id=?"
const sqlUpdateHead = "update player_data set `head`=? where player_id=?"
const sqlUpdateLevExp = "update player_data set lev=?,exp=? where player_id=?"

// hero
const sqlInsertHero = "insert into player_hero(player_id,hero_id,star,lev) values(?,?,?,?)"
const sqlUpdateHero = "update player_hero set star=?,lev=? where player_id=? and hero_id=?"

// item
const sqlUpdateItem = "update player_item set item_num=? where player_id=? and prop_id=?"
const sqlInsertItem = "insert into player_item(player_id,item_id,item_num) values(?,?,?)"

// property
const sqlUpdateProp = "update player_property set prop_value=?,update_ts=? where player_id=? and prop_id=?"
const sqlInsertProp = "insert into player_property(player_id,prop_id,prop_value,update_ts) values(?,?,?,?)"

// vip
const sqlInsert = "insert into player_vip values(?,?,?)"
const sqlUpdate = "update player_vip set vip_lev=?,vip_exp=? where player_id=?"

const sqlInsertVipGift = "insert into player_vip_gift values(?,?)"

// settings
const sqlSettingsInsert = "insert into player_settings values(?,?)"
const sqlSettingsUpdate = "update player_settings set language=? where player_id=?"

// mail
const sqlInsertMailCtrl = "insert into player_mail_ctrl values(?,?)"
const sqlUpdateMailCtrl = "update player_mail_ctrl set global_mail_id=? where player_id=?"

const sqlInsertMail = "insert into player_mail_list values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
const sqlUpdateMailStatus = "update player_mail_list set is_new=?,is_read=?,is_got=? where player_id=? and mail_id=?"
const sqlDeleteMail = "delete from player_mail_list where player_id=? and mail_id=?"
