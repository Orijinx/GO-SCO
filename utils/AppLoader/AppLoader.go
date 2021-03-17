package AppLoader

import DB"../../DataBase"
func Load() bool  {
	result := true
	if(!DB.CheckConnect()){
		panic("DB is not load")
	}
	return result

}
