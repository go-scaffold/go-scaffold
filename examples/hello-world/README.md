# hello-world

Example template that simply replace the name in [HELLO.md.tpl](HELLO.md.tpl)
with the one provided in the command line.

```text
.
├── .go-scaffold
│   └── prompts.yaml
├── HELLO.md.tpl
└── README.md
```

The [prompt.yaml](.go-scaffold/prompts.yaml) contains the template variables to
ask from command line, in this case only the name. `README.md` is not processed
as it's not a template.

To execute the template, run the following in this folder:

```sh
go-scaffold -t ./ -o ../generated/hello-world
```

At this point if you open [../generated/hello-world/HELLO.md](../generated/hello-world/HELLO.md)
you'll find the string:

```md
# Hello Pippo
```
