package main

const apphelptmpl = `{{if .Name}}NAME:   {{ "\n    " }}{{.Name   }}{{ "\n" }}
{{end          }}{{if .Usage   }}USAGE:  {{ "\n    " }}{{.Usage  }}{{ "\n" }}
{{end          }}{{if .Version }}VERSION:{{ "\n    " }}{{.Version}}{{ "\n" }}
{{end          }}{{if .Author  }}AUTHOR: {{ "\n    " }}{{.Author }}{{ "\n" }}
{{end          }}{{if .Email   }}EMAIL:  {{ "\n    " }}{{.Email  }}{{ "\n" }}
{{end          }}{{if .Commands}}COMMANDS:
    {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t"}}{{.Description}}
    {{end}}{{end}}`

const cmdhelptmpl = `{{if .Name}}NAME:        {{ "\n    "}}{{.Name       }}{{ "\n" }}
{{end}}{{if              .Usage}}USAGE:       {{ "\n    "}}{{.Usage      }}{{ "\n" }}
{{end}}{{if        .Description}}DESCRIPTION: {{ "\n    "}}{{.Description}}{{ "\n" }}
{{end}}{{if              .Flags}}OPTIONS:
	{{range .Flags}}{{.}}
	{{end}}{{end}}`
