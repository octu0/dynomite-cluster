package main

import(
  "log"
  "gopkg.in/urfave/cli.v1"
)

func discovery_action(c *cli.Context) error {
  _, err := prepareAction(c.Command.Name, c)
  if err != nil {
    return err
  }
  log.Printf("info: t.b.d.")
  return nil
}

func init() {
  AddCommand(cli.Command {
    Name: "discovery",
    Usage: "for pickup token and find neighbors node",
    Flags: []cli.Flag{
      cli.StringFlag{
        Name: "cluster-seed, s",
        Usage: "dynomite seed ip:port",
        Value: "127.0.0.1:2101",
      },
      cli.IntFlag{
        Name: "token, N",
        Usage: "number of target token `N`",
        Value: 0,
      },
      cli.StringFlag{
        Name: "replica-host, rh",
        Usage: "redis replica destination hostname",
        Value: "127.0.0.2",
      },
      cli.IntFlag{
        Name: "replica-port, rp",
        Usage: "redis replica destination port",
        Value: 6379,
      },
    },
    Action: discovery_action,
  })
}
