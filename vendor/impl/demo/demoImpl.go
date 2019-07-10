package demo
import(
	"github.com/TarsCloud/TarsGo/tars"
	tarsconf "github.com/Andrew-M-C/tarsgo-tools/config"
	"github.com/Andrew-M-C/tarsgo-tools/sesslog"
	test "proxy/Test"
	"encoding/base64"
	"fmt"
)

const (
	OBJ_NAME	= "Base64Obj"
)


func AddServant() {
	l := sesslog.New(OBJ_NAME)
	defer l.Close()

	cfg := tars.GetServerConfig()
	servant := cfg.App + "." + cfg.Server + "." + OBJ_NAME
	l.Debug("servant: %s", servant)

	tarscfg, _ := tarsconf.NewConfig()
	if tarscfg == nil {
		return
	}

	endpoint, exist := tarscfg.GetString(fmt.Sprintf("tars/application/server/%sAdapter", servant), "endpoint", "undefined")
	l.Debug("endpoint: %s", endpoint)
	if exist {
		intf := new(test.Base64)
		impl := new(Base64Impl)
		intf.AddServant(impl, servant)
		l.Info("add %s OK", servant)
	}
}


type Base64Impl struct {}

func (_ *Base64Impl) Encode(req *test.Base64Req, resp *test.Base64Resp) error {
	l := sesslog.New(OBJ_NAME + "_Encode")
	defer l.Close()

	resp.Output = base64.StdEncoding.EncodeToString([]byte(req.Input))
	resp.Error.Code = 0
	resp.Error.Message = "success"
	return nil
}
