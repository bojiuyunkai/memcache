package main

import (
	"fmt"
	"github.com/pangudashu/memcache"
	"reflect"
)

type SonT struct {
	Sid int
}

type User struct {
	Id    int64
	Count float64
	Sex   bool
	Name  string
	List  map[int]interface{}
	Son   *SonT
}

func main() {

	maxCnt := 32 //最大连接数
	initCnt := 0 //初始化连接数
	//初始化连接池
	mc, err := memcache.NewMemcache("127.0.0.1:12000", maxCnt, initCnt)
	if err != nil {
		fmt.Println(err)
		return
	}

	cmd_set(mc)
	cmd_get(mc)

	//clear pool
	mc.Close()
}

func cmd_set(mc *memcache.Memcache) {
	/*
		+----------------------------------------------------------------------+
		| Commond: Set                                                         |
		+----------------------------------------------------------------------+
		| @param key string                                                    |
		| @param value interface{}                                             |
		| @param expiration uint32  (*可选,默认值:0) 过期时间                  |
		| @param cas uint64 (*可选,默认值:0) 数据版本号，用于cas原子操作       |
		+----------------------------------------------------------------------+
		| @return result bool 操作成功:true 失败:false                         |
		| @return err error 操作成功:nil                                       |
		+----------------------------------------------------------------------+
	*/
	fmt.Println("\n+----------------------------------[Set]------------------------------------+\n")

	fmt.Println(mc.Set("pangudashu_bool", true, 0))                             //bool
	fmt.Println(mc.Set("pangudashu_string", "[string]this is string~", 0))      //string
	fmt.Println(mc.Set("pangudashu_bytes", []byte("[byte]this is []byte~"), 0)) //[]byte
	fmt.Println(mc.Set("pangudashu_int", 123456445, 0))                         //int
	fmt.Println(mc.Set("pangudashu_int8", int8(-128), 0))                       //int8
	fmt.Println(mc.Set("pangudashu_int16", int16(-3400), 0))                    //int16
	fmt.Println(mc.Set("pangudashu_int32", int32(-429496729), 0))               //int32
	fmt.Println(mc.Set("pangudashu_int64", int64(-8589934591), 0))              //int64
	fmt.Println(mc.Set("pangudashu_uint8", uint8(130), 0))                      //uint8
	fmt.Println(mc.Set("pangudashu_uint16", uint16(1300), 0))                   //uint16
	fmt.Println(mc.Set("pangudashu_uint32", uint32(130000000), 0))              //uint32
	fmt.Println(mc.Set("pangudashu_uint64", uint64(1300000000000000), 0))       //uint64
	fmt.Println(mc.Set("pangudashu_float32", float32(3.14), 0))                 //float32
	fmt.Println(mc.Set("pangudashu_float64", float64(3.1415926), 0))            //float64

	user := User{
		Id:    7,
		Count: 10000.888,
		Sex:   true,
		Name:  "盘古大叔",
		List:  make(map[int]interface{}),
		Son: &SonT{
			Sid: 80009,
		},
	}
	user.List[1001] = "北京1"
	user.List[1002] = "北京2"

	fmt.Println(mc.Set("pangudashu_struct", user, 0)) //struct
}

func cmd_get(mc *memcache.Memcache) {
	/*
		+----------------------------------------------------------------------+
		| Commond: Get                                                         |
		+----------------------------------------------------------------------+
		| @param key string                                                    |
		| @param format_struct interface{}  (*类型为struct时必选)              |
		+----------------------------------------------------------------------+
		| @return value interface{} 操作失败:nil                               |
		| @return cas uint64 操作成功:>0 失败:0                                |
		| @return err error 操作成功:nil                                       |
		+----------------------------------------------------------------------+
	*/

	fmt.Println("\n+----------------------------------[Get]------------------------------------+\n")

	var val [15]interface{}
	var ver [15]uint64
	var err [15]error

	val[1], ver[1], err[1] = mc.Get("pangudashu_bool")       //bool
	val[2], ver[2], err[2] = mc.Get("pangudashu_string")     //string
	val[3], ver[3], err[3] = mc.Get("pangudashu_bytes")      //[]byte
	val[4], ver[4], err[4] = mc.Get("pangudashu_int")        //int
	val[5], ver[5], err[5] = mc.Get("pangudashu_int8")       //int8
	val[6], ver[6], err[6] = mc.Get("pangudashu_int16")      //int16
	val[7], ver[7], err[7] = mc.Get("pangudashu_int32")      //int32
	val[8], ver[8], err[8] = mc.Get("pangudashu_int64")      //int64
	val[9], ver[9], err[9] = mc.Get("pangudashu_uint8")      //uint8
	val[10], ver[10], err[10] = mc.Get("pangudashu_uint16")  //uint16
	val[11], ver[11], err[11] = mc.Get("pangudashu_uint32")  //uint32
	val[12], ver[12], err[12] = mc.Get("pangudashu_uint64")  //uint64
	val[13], ver[13], err[13] = mc.Get("pangudashu_float32") //float32
	val[14], ver[14], err[14] = mc.Get("pangudashu_float64") //float64

	for i := 1; i < len(val); i++ {
		fmt.Println("No.", i, "\n\t【value】", val[i], "\n\t【value type】", reflect.TypeOf(val[i]), "\n\t【cas】", ver[i], "\n\t【error】", err[i])
	}

	var user User
	if _, _, e := mc.Get("pangudashu_struct", &user); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(user)
	}
}