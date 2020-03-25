package cuslan

import (
	"vgproj/vggame/public"
)

type CustomLanguageManager struct {
	mapLan map[string]map[int32]string
}

func NewCustomLanguageManager() *CustomLanguageManager {
	pObj := new(CustomLanguageManager)
	pObj.mapLan = make(map[string]map[int32]string)

	return pObj
}

func (cl *CustomLanguageManager) OnLoadData() error {
	rows, err := public.Server.GetGlobalDb().Query(0, sqlLoadCustomLanguage)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var tmpCustomKey string
		var tmpLan int32
		var tmpCustomVal string
		err = rows.Scan(&tmpCustomKey, &tmpLan, &tmpCustomVal)
		if err != nil {
			break
		}
		mapLanVal, exist := cl.mapLan[tmpCustomKey]
		if !exist {
			mapLanVal = make(map[int32]string)
			cl.mapLan[tmpCustomKey] = mapLanVal
		}
		mapLanVal[tmpLan] = tmpCustomVal
	}

	return err
}

func (cl *CustomLanguageManager) OnInit() error {
	return nil
}

func (cl *CustomLanguageManager) OnRelease() {

}

func (cl *CustomLanguageManager) GetLanguageValue(key string, lan int32) string {
	mapLanVal, exist := cl.mapLan[key]
	if !exist {
		return ""
	}
	val, exist := mapLanVal[lan]
	if !exist {
		return ""
	}
	return val
}

func (cl *CustomLanguageManager) AddLanguageValue(key string, mapLanVal map[int32]string) {
	cl.mapLan[key] = mapLanVal

	for k, v := range mapLanVal {
		public.Server.GetGlobalDb().AsyncExec(nil, nil, 0, sqlInsertCustomLanguage, key, k, v)
	}
}
