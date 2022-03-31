// Package api                _ooOoo_
//                           o8888888o
//                           88" . "88
//                           (| -_- |)
//                            O\ = /O
//                        ____/`---'\____
//                      .   ' \\| |// `.
//                       / \\||| : |||// \
//                     / _||||| -:- |||||- \
//                       | | \\\ - /// | |
//                     | \_| ''\---/'' | |
//                      \ .-\__ `-` ___/-. /
//                   ___`. .' /--.--\ `. . __
//                ."" '< `.___\_<|>_/___.' >'"".
//               | | : `- \`.;`\ _ /`;.`/ - ` : | |
//                 \ \ `-. \_ __\ /__ _/ .-` / /
//         ======`-.____`-.___\_____/___.-`____.-'======
//                            `=---='
//
//         .............................................
//                  佛祖镇楼                  BUG辟易
//          佛曰:
//                  写字楼里写字间，写字间里程序员；
//                  程序人员写程序，又拿程序换酒钱。
//                  酒醒只在网上坐，酒醉还来网下眠；
//                  酒醉酒醒日复日，网上网下年复年。
//                  但愿老死电脑间，不愿鞠躬老板前；
//                  奔驰宝马贵者趣，公交自行程序员。
//                  别人笑我忒疯癫，我笑自己命太贱；
//                  不见满街漂亮妹，哪个归得程序员？
package api

/*
前端选型:html+css+vue.js

后端:go+gin+mysql+redis+docker

可视化:python+框架待定+matprotlab

测试:jmeter+postman
*/

//接口:

// 注册:xxx/join

//数据类型:json
//方法:post

//参数
//

type Join struct {
	UserId string
	Name   string
	Id     string
	Email  string
	Tel    string
	Pass1  string //加密
	Pass2  string
}

type Verify struct {
	A int
	B int
}
type Verify_result struct {
	Num int
}
type Verify_OK struct {
	OK bool
}
type Join_OK struct {
	OK bool
}

// 登录: xxx/login
//数据类型:json
//方法:post

type Login struct {
	Tel    string
	Passwd string
	Verify int
}

//+verify

type Result struct {
	Result int //200: 登陆成功 500: 用户名或密码错误 501: 验证码错误
	//cookie http.Cookie
}

//准入初筛: xxx/screening
//数据类型:json
//方法:post

type Screening struct {
	Id int
}
type Return struct {
	Status int //200:ok 500:no 501:信息不足
	Num    int
}

//进入页面:xxx/seckill
//数据类型:json
//方法:post

type Time struct {
	Time string //"http://api.m.taobao.com/rest/api3.do?api=mtop.common.getTimestamp"
}

type Book struct {
	Url string //    xxx/seckill/book_string
}

//秒杀 xxx/seckilid
//数据类型:json
//方法:post

type Seckillid struct {
	Uuid int
}
type Act struct {
	Id int
}
type Success struct {
	Success int //200:成功  500:未抢到 501:被过滤
}

//  xxx/
type Ensure struct {
	Name   string
	Id     string
	Email  string
	Tel    int
	Statue bool //1成功 0 失败

}

//
