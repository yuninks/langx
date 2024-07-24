package langx

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type langx struct {
	ops      *options
	codeMap  map[string]int
	transMap map[string]map[string]string
	mut      sync.Mutex
}

var l *langx = &langx{}

func init() {

	l = &langx{
		ops:      defaultOptions(),
		codeMap:  make(map[string]int),
		transMap: make(map[string]map[string]string),
		mut:      sync.Mutex{},
	}
}

// 设置
func InitLangx(ops ...Option) {
	for _, opt := range ops {
		opt(l.ops)
	}
}

// 这是语言的Code
func RegisterCode(datas map[string]int) {
	l.mut.Lock()
	defer l.mut.Unlock()
	l.codeMap = datas
}

// 追加覆盖Code
func AppendCode(datas map[string]int) {
	l.mut.Lock()
	defer l.mut.Unlock()
	for k, v := range datas {
		l.codeMap[k] = v
	}
}

// 注册语言翻译
func RegisterTrans(langName string, trans map[string]string) {
	l.mut.Lock()
	defer l.mut.Unlock()
	l.transMap[langName] = trans
}

// 追加覆盖翻译
func AppendTrans(langName string, trans map[string]string) {
	l.mut.Lock()
	defer l.mut.Unlock()
	for k, v := range trans {
		if l.transMap[langName] == nil {
			l.transMap[langName] = map[string]string{}
		}
		l.transMap[langName][k] = v
	}
}

// 直接读取文件夹获取配置
// 要求：
// 1.json格式文件
// 2.code.json为自定义响应码 格式map[string]int{}
// 3.其他的json文件为对应语音 格式map[string]string{}
// 4.如果json解析错误将会panic
func RegisterDir(dir string) error {
	// 遍历dir获取.json的文件
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(info.Name(), ".json") {
			// 读取文件
			fileName := strings.Replace(info.Name(), ".json", "", 1)
			by, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			if fileName == "code" {
				data := map[string]int{}
				err = json.Unmarshal(by, &data)
				if err != nil {
					return err
				}
				RegisterCode(data)
			} else {
				data := map[string]string{}
				err = json.Unmarshal(by, &data)
				if err != nil {
					return err
				}
				RegisterTrans(fileName, data)
			}
		}

		return nil
	})

	// code文件为状态码,其他为对应的语言文件，文件名为语言名
	return err
}

func RegisterEmbed(asset embed.FS) error {
	paths, err := readEmbedAllPath(asset, "")
	if err != nil {
		return err
	}

	for _, ff := range paths {
		// 获取文件的后缀

		fileName := strings.Replace(ff.path, ".json", "", 1)
		if fileName == "code" {
			data := map[string]int{}
			err = json.Unmarshal(ff.datas, &data)
			if err != nil {
				return err
			}
			RegisterCode(data)
		} else {
			data := map[string]string{}
			err = json.Unmarshal(ff.datas, &data)
			if err != nil {
				return err
			}
			RegisterTrans(fileName, data)
		}
	}

	return nil
}

type fileData struct {
	path  string
	datas []byte
}

func readEmbedAllPath(asset embed.FS, path string) ([]fileData, error) {
	newRoot := ""
	if path == "" {
		newRoot = "."
	} else {
		newRoot = path
	}

	fs, err := asset.ReadDir(newRoot)
	if err != nil {
		return nil, err
	}
	var files []fileData
	for _, f := range fs {
		if f.IsDir() {
			chaild := ""
			if newRoot == "." {
				chaild = f.Name()
			} else {
				chaild = newRoot + "/" + f.Name()
			}

			df, err := readEmbedAllPath(asset, chaild)
			if err != nil {
				return nil, err
			}
			files = append(files, df...)
			continue
		}

		by, err := asset.ReadFile(newRoot + "/" + f.Name())
		if err != nil {
			return nil, err
		}

		ff := fileData{
			path:  f.Name(),
			datas: by,
		}

		files = append(files, ff)
	}
	return files, nil
}

// 获取翻译
// 包含Code和Message
func GetTransFormat(lang string, key string, format map[string]string) (code int, str string) {
	code = GetCode(key)
	str = GetFormat(lang, key, format)
	return
}

func GetTrans(lang string, key string) (code int, str string) {
	return GetTransFormat(lang, key, map[string]string{})
}

func GetTransCtx(ctx context.Context, key string) (code int, str string) {
	return GetTrans(getLangFromCtx(ctx), key)
}

func GetTransFormatCtx(ctx context.Context, key string, format map[string]string) (code int, str string) {
	return GetTransFormat(getLangFromCtx(ctx), key, format)
}

// 根据Key获取code
func GetCode(key string) int {
	l.mut.Lock()
	defer l.mut.Unlock()
	code, ok := l.codeMap[key]
	if !ok {
		return l.ops.defaultCode
	}
	return code
}

// 获取翻译
func GetMsg(lang string, key string) string {
	l.mut.Lock()
	defer l.mut.Unlock()
	// 找指定语言
	str, ok := l.transMap[lang]
	if ok {
		val, ok := str[key]
		if ok {
			return val
		}
	}
	// 找默认语言
	str, ok = l.transMap[l.ops.defaultLang]
	if ok {
		val, ok := str[key]
		if ok {
			return val
		}
	}

	return key
}

// 从ctx里面获取语言
func GetMsgCtx(ctx context.Context, key string) string {
	return GetMsg(getLangFromCtx(ctx), key)
}

// 拼接回复
func GetFormat(lang string, key string, arr map[string]string) string {
	str := GetMsg(lang, key)
	for k, v := range arr {
		str = strings.ReplaceAll(str, fmt.Sprintf(l.ops.replaceKey, k), v)
	}
	return str
}

func GetFormatCtx(ctx context.Context, key string, arr map[string]string) string {
	return GetFormat(getLangFromCtx(ctx), key, arr)
}

// 获取默认Code
func GetDefaultCode() int {
	return l.ops.defaultCode
}

func getLangFromCtx(ctx context.Context) string {
	ctxVal := ctx.Value(l.ops.ctxLangKey)
	lang := l.ops.defaultLang
	if ctxVal != nil {
		lang = ctxVal.(string)
	}
	return lang
}
