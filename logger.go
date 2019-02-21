package cluster

import(
  "log"
  "os"
  "io"
  "time"

  "github.com/comail/colog"
  "github.com/lestrrat-go/file-rotatelogs"
)

type GeneralLogger struct {
  std io.Writer
  sub io.Writer
  r   *rotatelogs.RotateLogs
  cl  *log.Logger
}
func NewGeneralLogger(config Config) *GeneralLogger {
  rotate, err := rotatelogs.New(
    config.LogDir + "/general_log.%Y%m%d",
    rotatelogs.WithRotationTime(1 * time.Minute),
    rotatelogs.WithMaxAge(-1),
  )
  if err != nil {
    log.Fatalf("error: file logger creation failed: %s", err.Error())
  }

  c := colog.NewCoLog(rotate, "cluster ", log.Ldate | log.Ltime | log.Lshortfile)
  c.SetDefaultLevel(colog.LDebug)
  c.SetMinLevel(colog.LInfo)
  if config.DebugMode {
    c.SetMinLevel(colog.LDebug)
    if config.VerboseMode {
      c.SetMinLevel(colog.LTrace)
    }
  }

  l    := new(GeneralLogger)
  l.std = os.Stdout
  l.sub = rotate
  l.r   = rotate
  l.cl  = c.NewLogger()
  return l
}
func (l *GeneralLogger) Write(p []byte) (int, error) {
  l.std.Write(p)
  return l.sub.Write(p)
}
func (l *GeneralLogger) Logger() *log.Logger {
  return l.cl
}
