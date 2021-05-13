package main

//select COLUMN_NAME,DATA_TYPE,COLUMN_TYPE  from information_schema.COLUMNS where  table_schema = 'dbiot' and  table_name = 'mqtt_acl';
//COLUMN_NAME,DATA_TYPE,COLUMN_TYPE
import (
	"database/sql"
	"embed"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"html/template"
	"io/fs"
	"os"
	"path"
	"strings"
	"unicode"
)



type Column struct {
	ColumnName      string        `json:"colname"`
	DataType        string        `json:"datatype"`
	ColumnType      string        `json:"coltype"`
	Nump            int           `json:"nump"`
	Nums            int           `json:"nums"`
	Comment         string        `json:"comment"`
	ColumnKey       string        `json:"columnkey"`
	Extra           string        `json:"extra"`
	OrdinalPosition string        `json:"position"`
	SqlStr          template.HTML `json:"sqlstr"`
}

func (col *Column) IsKey() bool {
	return col.ColumnKey == "PRI"
}

func (col *Column) AutoIncrement() bool {
	return strings.Index(col.Extra, "auto_increment") > -1
}

func (col *Column) Build() string {
	return col.ColumnName + " " + col.DataType
}

type DstData struct {
	Package string   `json:"package'`
	Model   string   `json:"model'`
	ModelL  string   `json:"modell"`
	Columns []Column `json:"columns"`
}

func ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}
func lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

//abc_def_ghi=> AbcDefGhi
func transfer(in string) string {
	dstdata := make([]string, 0)
	inarr := strings.Split(in, "_")
	for _, v := range inarr {
		dstdata = append(dstdata, ucfirst(v))
	}
	return strings.Join(dstdata, "")
}

var datatypemap map[string]string = map[string]string{
	"int":      "int",
	"bigint":   "uint",
	"datetime": "restgo.DateTime",
	"date":     "restgo.Date",
	"varchar":  "string",
	"bit":      "int",
	"decimal":  "float64",
}

//Col int
func datatype(col Column) string {
	t := col.DataType
	r, ok := datatypemap[t]
	if ok {
		return r
	} else {
		return t
	}
}
func buildsql(col Column) template.HTML {
	uname := transfer(col.ColumnName)
	lname := lcfirst(uname)
	if col.ColumnName == "id" {
		return `restgo.BaseModel`
	}
	ret := uname + " " + datatype(col) + " " + " ` " + "json:\"" + lname + "\" form:\"" + lname + "\""
	if col.DataType == "date" || col.DataType == "datetime" {
		ret = ret + ` time_format:"2006-01-02 15:04:05" time_utc:"1"`
	}
	ret = ret + "`"

	return template.HTML(ret)
}

//静态资源处理

//go:embed tmpl/*
var Tmpls embed.FS

type fsFunc func(name string) (fs.File, error)

func (f fsFunc) Open(name string) (fs.File, error) {
	return f(name)
}

func TempleteFs(assets embed.FS, root string) fs.FS {
	handler := fsFunc(func(name string) (fs.File, error) {
		assetPath := path.Join(root, name)
		// If we can't find the asset, fs can handle the error
		file, err := assets.Open(assetPath)
		if err != nil {
			return nil, err
		}
		return file, err
	})
	return handler
}



type Config struct{
	Database string `mapstructure:"database" json:"database"`
	Table string `mapstructure:"table" json:"table"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
	Host string `mapstructure:"host" json:"host"`
	Port string `mapstructure:"port" json:"port"`
	Model string `mapstructure:"model" json:"model"`
	Package string `mapstructure:"package" json:"package"`
	Dstdir string  `mapstructure:"dstdir" json:"dstdir"`
}

var db = flag.String("db", "test", "database name")
var table = flag.String("t", "test", "table name")
var modelin = flag.String("m", "", "out model")
var dstdir = flag.String("o", "./", "dist dir")
var user = flag.String("u", "root", "user name")
var passwd = flag.String("p", "", "password")
var host = flag.String("h", "127.0.0.1", "database host")
var port = flag.String("a", "3306", "mysql port")
var pkg = flag.String("pkg", "turinapp", "application package")
var cfgpath = flag.String("c", "./restgo.yaml", "config file path")

var model = ""
var config Config
func main() {
	flag.Parse()

	v := viper.New()

	v.SetConfigFile(*cfgpath)
	err := v.ReadInConfig()
	if err != nil {
		//panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.Unmarshal(&config)
	if *db != "test"{
		config.Database = *db
	}
	if *table != "test"{
		config.Table = *table
	}
	if *modelin != ""{
		config.Model = *modelin
	}
	if *dstdir != "./"{
		config.Dstdir = *dstdir
	}
	if *user != "root"{
		config.Username = *user
	}
	if *passwd != ""{
		config.Password = *passwd
	}
	if *host != "127.0.0.1"{
		config.Host = *host
	}
	if *port != "3306"{
		config.Port = *port
	}

	model = config.Model
	if model == "" {
		model = config.Table
	}
	model = strings.ToLower(model)
	// Open方法第二个参数:  用户名:密码@协议(ip:端口)/数据库
	dnsstr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Username,config.Password,config.Host,config.Port,config.Database)
	//fmt.Println(dnsstr)
	MtsqlDb, err := sql.Open("mysql", dnsstr)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer MtsqlDb.Close()
	columns := make([]Column, 0)
	rows, err := MtsqlDb.Query(`select COLUMN_NAME ,DATA_TYPE,COLUMN_TYPE,NUMERIC_PRECISION,NUMERIC_SCALE,COLUMN_COMMENT,column_key,extra,ORDINAL_POSITION  from information_schema.COLUMNS where  table_schema = ? and  table_name = ?`, *db, *table)
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var col Column
		rows.Scan(&col.ColumnName, &col.DataType, &col.ColumnType, &col.Nump, &col.Nums, &col.Comment, &col.ColumnKey, &col.Extra, &col.OrdinalPosition)
		col.SqlStr = buildsql(col)
		columns = append(columns, col)
	}

	tplfs := TempleteFs(Tmpls, "")
	tmpl, err := template.ParseFS(tplfs, "tmpl/*")

	if err != nil {
		fmt.Println(err)
		return
	}
	tpls := []string{
		"server/args", "server/model", "server/ctrl", "server/service",
	}
	for _, tpl := range tpls {
		os.MkdirAll(*dstdir+"/"+tpl, fs.FileMode(os.O_CREATE))
		f, err := os.OpenFile(*dstdir+"/"+tpl+"/"+model+".go", os.O_WRONLY|os.O_CREATE, 0766)
		if err != nil {
			fmt.Println(err)
			return
		}

		tmpl.ExecuteTemplate(f, tpl, DstData{
			Package: *pkg,
			Model:   ucfirst(transfer(model)),
			ModelL:  lcfirst(transfer(model)),
			Columns: columns,
		})
	}

	os.MkdirAll(*dstdir+"/ui/view/"+model, fs.FileMode(os.O_CREATE))
	f, err := os.OpenFile(*dstdir+"/ui/view/"+model+"/list.vue", os.O_WRONLY|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpl.ExecuteTemplate(f, "view/list", DstData{
		Package: *pkg,
		Model:   ucfirst(transfer(model)),
		ModelL:  lcfirst(transfer(model)),
		Columns: columns,
	})

	os.MkdirAll(*dstdir+"/ui/api/"+model, fs.FileMode(os.O_CREATE))
	f, err = os.OpenFile(*dstdir+"/ui/api/"+model+".js", os.O_WRONLY|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpl.ExecuteTemplate(f, "view/api", DstData{
		Package: *pkg,
		Model:   ucfirst(transfer(model)),
		ModelL:  lcfirst(transfer(model)),
		Columns: columns,
	})
	v.WriteConfig()
}
