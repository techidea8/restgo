package utils

import "fmt"

type XmlCtx struct {
	Data map[string]interface{}
	Root string
}

func (ctx *XmlCtx) CreateRoot(key string) *XmlCtx {
	ctx.Root = key
	return ctx
}
func (ctx *XmlCtx) AddXmlElement(key string, data interface{}) *XmlCtx {
	ctx.Data[key] = data
	return ctx
}

func (ctx *XmlCtx) BuildXml(key string, data interface{}) string {
	root := "root"
	if len(ctx.Root) > 0 {
		root = ctx.Root
	}
	str := ""
	for k, v := range ctx.Data {
		str += fmt.Sprintf("<%s>%s</%s>", k, v.(string), k)
	}
	return fmt.Sprintf("<%s>%s</%s>", root, str, root)

}
