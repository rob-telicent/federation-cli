# federation-cli

## Main command

```
federation-cli --help
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  federation-cli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  consume     A brief description of your command
  debug       A brief description of your command
  help        Help about any command
  topics      A brief description of your command

Flags:
      --apiKey string     help for apiKey
      --cacert string     help for cacert
      --cert string       help for cert
      --cfgFile string    help for cfgFile (default is $HOME/.federation-cli.yaml)
      --clientId string   help for clientId
  -h, --help              help for federation-cli
      --insecure          help for insecure
      --key string        help for key
      --mtls              help for mtls
      --port int          help for port (default 443)
      --server string     help for server (default "127.0.0.1")
```

## Topics


```
federation-cli topics --help
A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  federation-cli topics [flags]

Flags:
  -h, --help   help for topics

Global Flags:
      --apiKey string     help for apiKey
      --cacert string     help for cacert
      --cert string       help for cert
      --cfgFile string    help for cfgFile (default is $HOME/.federation-cli.yaml)
      --clientId string   help for clientId
      --insecure          help for insecure
      --key string        help for key
      --mtls              help for mtls
      --port int          help for port (default 443)
      --server string     help for server (default "127.0.0.1")
```

## Consume

```
federation-cli consume --help
A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  federation-cli consume [flags]

Flags:
  -h, --help           help for consume
      --offset int     the offset to consume from
      --topic string   the topic to consume

Global Flags:
      --apiKey string     help for apiKey
      --cacert string     help for cacert
      --cert string       help for cert
      --cfgFile string    help for cfgFile (default is $HOME/.federation-cli.yaml)
      --clientId string   help for clientId
      --insecure          help for insecure
      --key string        help for key
      --mtls              help for mtls
      --port int          help for port (default 443)
      --server string     help for server (default "127.0.0.1")
```
