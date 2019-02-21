package main

import (
  "log"
  "runtime"
  "context"
  "os"
  "os/signal"
  "syscall"

  "github.com/comail/colog"
  "gopkg.in/urfave/cli.v1"

  "github.com/octu0/dynomite-cluster"
)

var (
  Commands = make([]cli.Command, 0)
)
func AddCommand(cmd cli.Command){
  Commands = append(Commands, cmd)
}

func createConfig(c *cli.Context) cluster.Config {
  config := cluster.Config{
    DebugMode:   c.GlobalBool("debug"),
    VerboseMode: c.GlobalBool("verbose"),
    Procs:       c.GlobalInt("procs"),
    LogDir:      c.GlobalString("log-dir"),
  }
  if config.Procs < 1 {
    config.Procs = 1
  }

  if config.DebugMode {
    colog.SetMinLevel(colog.LDebug)
    if config.VerboseMode {
      colog.SetMinLevel(colog.LTrace)
    }
  }
  return config
}

func prepareAction(commandName string, c *cli.Context) (context.Context, error) {
  config := createConfig(c)
  logger := cluster.NewGeneralLogger(config)
  colog.SetOutput(logger)

  ctx := context.Background()
  ctx  = context.WithValue(ctx, "config", config)
  ctx  = context.WithValue(ctx, "logger.general", logger)

  log.Printf("info: starting %s.%s-%s", cluster.AppName, commandName, cluster.Version)

  return ctx, nil
}

func waitSignal(f func(os.Signal)) {
  signal_chan := make(chan os.Signal, 10)
  signal.Notify(signal_chan, syscall.SIGTERM)
  signal.Notify(signal_chan, syscall.SIGHUP)
  signal.Notify(signal_chan, syscall.SIGQUIT)
  signal.Notify(signal_chan, syscall.SIGINT)

  for {
    select {
    case sig := <-signal_chan:
      log.Printf("info: signal trap(%s)", sig.String())
      f(sig)
    }
  }
}

func main(){
  colog.SetDefaultLevel(colog.LDebug)
  colog.SetMinLevel(colog.LInfo)

  colog.SetFormatter(&colog.StdFormatter{
    Flag: log.Ldate | log.Ltime | log.Lshortfile,
  })
  colog.Register()

  app         := cli.NewApp()
  app.Version  = cluster.Version
  app.Name     = cluster.AppName
  app.Author   = ""
  app.Email    = ""
  app.Usage    = ""
  app.Commands = Commands
  app.Flags    = []cli.Flag{
    cli.IntFlag{
      Name: "procs, P",
      Usage: "attach cpu(s)",
      Value: runtime.NumCPU(),
    },
    cli.BoolFlag{
      Name: "debug, d",
      Usage: "debug mode",
    },
    cli.BoolFlag{
      Name: "verbose, V",
      Usage: "verbose. more message",
    },
    cli.StringFlag{
      Name: "log-dir, l",
      Usage: "logs out dir",
      Value: "/tmp",
    },
  }
  if err := app.Run(os.Args); err != nil {
    log.Printf("error: %s", err.Error())
    cli.OsExiter(1)
  }
}
