package args

type RegisterArgs struct {
	Username string `arg:"required"`
	Password string `arg:"required"`
}

type LoginArgs struct {
	Username string `arg:"required"`
	Password string `arg:"required"`
}

type FileUpload struct {
	Dest string `arg:"required"`
	Src  string `arg:"required"`
}

type Args struct {
	Register *RegisterArgs `arg:"subcommand:register"`
	Login    *LoginArgs    `arg:"subcommand:login"`
	Upload   *FileUpload   `arg:"subcommand:upload"`
	Host     string        `default:"127.0.0.1"`
	Port     string        `default:"3000"`
}

func NewArgs() *Args {
	return &Args{}
}
