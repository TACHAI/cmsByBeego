package sysinit


func init()  {
	//utils.Cache =cache.New(60*time.Minute,120*time.Minute)

	initDB()
}