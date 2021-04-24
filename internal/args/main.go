package args

type FileUpload struct {
	Dest string `arg:"required"`
	Src  string `arg:"required"`
}

type RegisterArgs struct {
	Username string
	Password string
	Host     string
	Port     string
}

type Args struct {
	Register   *RegisterArgs `arg:"subcommand:register"`
	Upload     *FileUpload   `arg:"subommand:upload"`
	ConfigPath string        `default:"~/.restless/config" env:"CONF_PATH"`
}

func NewArgs() *Args {
	return &Args{}
}
