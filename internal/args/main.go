package args

type FileUpload struct {
	Dest string `arg:"required"`
	Src  string `arg:"required"`
}

type RegisterArgs struct {
	Username string `arg:"required"`
	Password string `arg:"required"`
	Host     string `arg:"required"`
	Port     string `arg:"required"`
}

type Args struct {
	Register   *RegisterArgs `arg:"subcommand:register"`
	Upload     *FileUpload   `arg:"subcommand:upload"`
	ConfigPath string        `default:"~/.restless/config" env:"CONF_PATH"`
}

func NewArgs() *Args {
	return &Args{}
}
