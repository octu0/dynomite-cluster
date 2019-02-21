package main

import(
  "log"
  "gopkg.in/urfave/cli.v1"
)

func validate_action(c *cli.Context) error {
  _, err := prepareAction(c.Command.Name, c)
  if err != nil {
    return err
  }
  log.Printf("info: t.b.d.")
  return nil
}

func init() {
  AddCommand(cli.Command {
    Name: "validate",
    Usage: "validation for node setup before cluster",
    Flags: []cli.Flag{
      cli.StringFlag{
        Name: "peer-host, ph",
        Usage: "redis replica source hostname",
        Value: "127.0.0.1",
      },
      cli.IntFlag{
        Name: "peer-port, pp",
        Usage: "redis replica source port",
        Value: 6379,
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
    Action: validate_action,
  })
}
