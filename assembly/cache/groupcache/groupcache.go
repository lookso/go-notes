package main

func main() {
	////定义节点数量以及地址
	//peers_addrs := []string{"http://127.0.0.1:8001", "http://127.0.0.1:8002"}
	//db, _ := sql.Open("sqlite3", "./console.db")
	//
	//if len(os.Args) != 2 {
	//	fmt.Println("\r\n Usage local_addr \t\n local_addr must in(127.0.0.1:8001,127.0.0.1:8002)\r\n")
	//	os.Exit(1)
	//}
	//local_addr := os.Args[1]
	//peers := groupcache.NewHTTPPool("http://" + local_addr)
	//peers.Set(peers_addrs...)
	//
	//// 获取group对象
	//image_cache := groupcache.NewGroup("testGroup", 8<<30,
	//	// 自定义数据获取来源
	//	groupcache.GetterFunc(
	//		func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
	//			rows, _ := db.Query("SELECT key, value FROM tbl_cache_map where key = ?", key)
	//			for rows.Next() {
	//				p := new(TblCache)
	//				err := rows.Scan(&p.Key, &p.Value)
	//				if err != nil {
	//					fmt.Println(err)
	//				}
	//				fmt.Printf("get %s of value from tbl_cache_map\n", key)
	//				dest.SetString("tbl_cache_map.value : " + p.Value)
	//			}
	//			return nil
	//		}))
	//
	//// 定义返回方式
	//http.HandleFunc("/get", func(rw http.ResponseWriter, r *http.Request) {
	//	var data []byte
	//	k := r.URL.Query().Get("key")
	//	fmt.Printf("user get %s of value from groupcache\n", k)
	//	image_cache.Get(nil, k, groupcache.AllocatingByteSliceSink(&data))
	//	rw.Write([]byte(data))
	//})
	//
	//log.Fatal(http.ListenAndServe(local_addr, nil))
}
