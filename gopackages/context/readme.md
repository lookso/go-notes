1.使用原因
context是Go中广泛使用的程序包，由Google官方开发，在1.7版本引入。
它用来简化在多个go routine传递上下文数据、(手动/超时)中止routine树等操作，
比如，官方http包使用context传递请求的上下文数据，gRpc使用context来终止某个请求产生的routine树。
由于它使用简单，现在基本成了编写go基础库的通用规范。


context.Background 只应用在最高等级，作为所有派生 context 的根。
context.TODO 应用在不确定要使用什么的地方，或者当前函数以后会更新以便使用 context。
context 取消是建议性的，这些函数可能需要一些时间来清理和退出。
context.Value 应该很少使用，它不应该被用来传递可选参数。这使得 API 隐式的并且可以引起错误。取而代之的是，这些值应该作为参数传递。
不要将 context 存储在结构中，在函数中显式传递它们，最好是作为第一个参数。
永远不要传递不存在的 context 。相反，如果您不确定使用什么，使用一个 ToDo context。
Context 结构没有取消方法，因为只有派生 context 的函数才应该取消 context。


要创建context树，第一步是要有一个根结点。context.Background函数的返回值是一个空的context，
经常作为树的根结点，它一般由接收请求的第一个routine创建，不能被取消、没有值、也没有过期时间。

func Background() Context

之后该怎么创建其它的子孙节点呢？context包为我们提供了以下函数：
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithValue(parent Context, key interface{}, val interface{}) Context
这四个函数的第一个参数都是父context，返回一个Context类型的值，这样就层层创建出不同的节点。
子节点是从复制父节点得到的，并且根据接收的函数参数保存子节点的一些状态值，然后就可以将它传递给下层的routine了。

