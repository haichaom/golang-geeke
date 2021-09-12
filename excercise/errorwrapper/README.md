## Question
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

## Answer

Yes, we should handle it, and wrapper this error/throw it for the caller. We should get backtrace of the error and find the root-cause quickly.

## Console output

PS C:\Users\Administrator\Documents\golang-geeke\excercise\errorwrapper> go run main.go\
error: *errors.errorString sql: no rows in result set\
Trackback: sql: no rows in result set\
No record found\
github.com/haichaom/golang-geeke/excercise/errorwrapper/dao.GetDBRecord\
        C:/Users/Administrator/Documents/golang-geeke/excercise/errorwrapper/dao/dao.go:17\
github.com/haichaom/golang-geeke/excercise/errorwrapper/service.GetRecordByID\
        C:/Users/Administrator/Documents/golang-geeke/excercise/errorwrapper/service/service.go:9\
main.main\
        C:/Users/Administrator/Documents/golang-geeke/excercise/errorwrapper/main.go:11\
runtime.main\
        C:/Program Files/Go/src/runtime/proc.go:255\
runtime.goexit\
        C:/Program Files/Go/src/runtime/asm_amd64.s:1581\
service::getRecordByID failed with error!\
github.com/haichaom/golang-geeke/excercise/errorwrapper/service.GetRecordByID\
        C:/Users/Administrator/Documents/golang-geeke/excercise/errorwrapper/service/service.go:13\
main.main\
        C:/Users/Administrator/Documents/golang-geeke/excercise/errorwrapper/main.go:11\
runtime.main\
        C:/Program Files/Go/src/runtime/proc.go:255\
runtime.goexit\
        C:/Program Files/Go/src/runtime/asm_amd64.s:1581\