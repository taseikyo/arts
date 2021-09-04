/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2021-08-29 14:54:46
 * @link    github.com/taseikyo
 */

package main

// 按照来源分段import
import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"git.xiaorui.cc/ocean/jellyfish/internal/api"
	"git.xiaorui.cc/ocean/jellyfish/internal/config"
	"git.xiaorui.cc/ocean/jellyfish/internal/middleware"
	"git.xiaorui.cc/ocean/jellyfish/pkg/log"
)

// 字段对齐，注释对齐
type User struct {
	Username   string // ⽤户名
	Email      string // 邮箱
	URI        string // 后缀
	API        string // 地址
	IsOpen     bool   // 开放
	CreateTime bool   // 创建时间
}

// Bool使用”判断”语义的前缀
var (
	isExist      bool
	hasConflict  bool
	canManage    bool
	allowGitHook bool
)

type Scheme string

const (
	HTTP  Scheme = "http"
	HTTPS Scheme = "https"
)

// 自定义类型常量, iota从1开始
const (
	ModeAdd = iota + 1
	ModeDel
	ModeUpdate
	ModeUpsert
)

// 动态参数
type Option interface {
	// ...
}

func WithCache(c bool) Option {
	// ...
}
func WithLogger(log *zap.Logger) Option {
	// ...
}

// Open creates a connection.
func Open(opts ...Option) (*Connection, error) {
	// ...
}

// 默认对象
const (
	_defaultPort = 8080
	defaultUser  = "user"
)

var (
	ErrNotFound = errors.New("not found")
)

// 注意函数内做好语义拆分
// 不要try-catch那样使用panic
// 代码要减少 if for 嵌套

// 接⼝合理性检
type Handler struct {
	// ...
}

// ⽤于触发编译期的接⼝的合理性检查机制
// 如果Handler没有实现http.Handler,会在编译期报错
var _ http.Handler = (*Handler)(nil)

func (h *Handler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	// ...
}

var (
	defaultGitlabClient *Client
	once                = sync.Once{}
)

// NewGitlabClient create gitlab client
func NewGitlabClient(disfName string) (*Client, error) {
	once.Do(func() {
		defaultGitlabClient, err = newClient(disfName)
	})
	return defaultGitlabClient, err
}

// 安全的单例
// 函数注释

type T struct{ a int }

// value receiver
func (tv T) Mv(a int) { tv.a = 123 }

// pointer receiver
func (tp *T) Mp(f int) { tp.a = 123 }

// 值接收器
// 指针接收器

func main() {

}
