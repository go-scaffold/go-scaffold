# go-scaffold

[![Go Report
Card](https://goreportcard.com/badge/github.com/go-scaffold/go-scaffold)](https://goreportcard.com/report/github.com/go-scaffold/go-scaffold)
[![CI
Status](https://github.com/go-scaffold/go-scaffold/workflows/Continuous%20integration/badge.svg)](https://github.com/go-scaffold/go-scaffold/actions)

Command line application that generates files/projects from a
[template](https://pkg.go.dev/text/template).

The app is heavily inspired by [Helm](https://helm.sh/), but with the intent to
be general purpose.

You can look at the [examples](./examples) folder to checkout some examples.

**Note**: if you are looking for the legacy version of the application, please
check the `legacy` branch.

## Build and install

To build the executable:

```sh
make go-build
```

To build the docker image:

```sh
make docker-build
```

To build both:

```sh
make build
```

To build and install it:

```sh
make install
```

## Install using Go

If you have Go installed, you can install the go-scaffold directly:

```sh
go install github.com/go-scaffold/go-scaffold/cmd/go-scaffold@master
```

## Usage

The go-scaffold CLI provides two main commands:

### Create a new template project

Use the `create` command to initialize a new template project with the required
file structure:

```sh
go-scaffold create [name]
```

This command creates:
- `Manifest.yaml` - Defines the template configuration
- `values.yaml` - Contains default values for the template
- `templates/` directory - Contains the actual template files

If no name is provided, the command will create the template structure in the
current directory (which must be empty).

### Template generation

In order to generate the output files from a template:

```sh
go-scaffold generate [<flags>] <template-dir> <output-dir>
```

Use the `-f` flag to specify overriding values files:

```sh
go-scaffold generate -f ./examples/hello-world-markdown/values-project1.yaml ./examples/hello-world-markdown build/
```

The template directory should contain:
- `Manifest.yaml` (optional, but recommended)
- `values.yaml` - Default values for the template
- `templates/` directory - Contains template files (all files will be processed
  as templates)

### Template Structure

Templates should be organized as follows:

```
my-template/
├── Manifest.yaml      # Template configuration (optional)
├── values.yaml        # Default values
└── templates/         # Template files
    ├── file1.txt      # Template file
    ├── file2.yaml     # Another template file
    └── ...
```

Values in templates can be accessed using `.Values.key` syntax (e.g., `{{
.Values.project }}`). In templates, use the `.Values` prefix to access values
from the values.yaml file.

### Template Functions

The go-scaffold template engine provides a rich set of functions to use in your
templates, including:

#### Built-in functions from [Sprig](https://github.com/Masterminds/sprig)

The template engine includes all functions from the Sprig library. These
include:

**String Functions:**
- `abbrev`, `abbrevboth`, `trunc` - String abbreviation functions
- `trim`, `upper`, `lower`, `title`, `untitle` - String case and trimming
  functions
- `substr` - Substring function
- `repeat` - Repeats a string n times
- `trimAll`, `trimSuffix`, `trimPrefix` - Advanced trimming
- `nospace`, `initials` - String manipulation
- `randAlphaNum`, `randAlpha`, `randAscii`, `randNumeric` - Random string
  generators
- `snakecase`, `camelcase`, `kebabcase` - Case conversion functions
- `wrap`, `wrapWith` - Text wrapping
- `contains`, `hasPrefix`, `hasSuffix` - String matching
- `quote`, `squote` - Quote functions
- `cat` - Concatenates strings
- `indent`, `nindent` - Indentation functions
- `replace` - Replaces occurrences of a string
- `plural` - Creates pluralized strings
- `regexMatch`, `regexFindAll`, `regexFind`, `regexReplaceAll`, `regexSplit` -
  Regex functions
- `regexQuoteMeta` - Quotes regex metacharacters

**Math Functions:**
- `add`, `add1`, `addf`, `add1f` - Addition functions
- `sub`, `subf` - Subtraction functions
- `mul`, `mulf` - Multiplication functions
- `div`, `divf` - Division functions
- `mod` - Modulo operation
- `max`, `maxf`, `min`, `minf` - Comparison functions
- `ceil`, `floor`, `round` - Rounding functions
- `randInt` - Generates random integer
- `seq` - Generates sequence of integers

**Date Functions:**
- `date`, `date_in_zone` - Date formatting
- `date_modify`, `ago` - Date manipulation
- `duration`, `durationRound` - Duration formatting
- `now` - Current time
- `unixEpoch` - Unix epoch time

**Collection/Array Functions:**
- `join` - Joins array elements with separator
- `split`, `splitList`, `splitn` - String splitting functions
- `sortAlpha` - Sorts alphabetically
- `uniq` - Removes duplicates
- `without` - Returns list without specified values
- `concat` - Concatenates arrays
- `slice` - Returns a slice of an array
- `first`, `last`, `rest`, `initial` - Array element access
- `reverse` - Reverses an array
- `append`, `push`, `prepend` - Array modification
- `chunk` - Chunks array into smaller arrays

**Dictionary Functions:**
- `dict` - Creates a dictionary
- `get`, `set`, `unset` - Dictionary access/modification
- `hasKey` - Checks if key exists
- `pluck`, `keys`, `pick`, `omit` - Dictionary manipulation
- `merge`, `mergeOverwrite` - Dictionary merging
- `values` - Returns all values
- `dig` - Navigates nested dictionary

**Type Conversion Functions:**
- `atoi` - String to integer conversion
- `int`, `int64`, `float64` - Type conversions
- `toString`, `toStrings` - String conversions
- `toJson`, `toPrettyJson`, `toRawJson` - JSON conversions
- `fromJson` - JSON parsing

**Reflection Functions:**
- `typeOf`, `typeIs`, `typeIsLike` - Type checking
- `kindOf`, `kindIs` - Kind checking
- `deepEqual` - Deep equality check

**Default/Conditional Functions:**
- `default` - Provides default value
- `empty` - Checks if value is empty
- `coalesce` - Returns first non-empty value
- `ternary` - Ternary operator
- `all`, `any` - Logical functions
- `fail` - Causes template to fail with error

**Path Functions:**
- `base`, `dir`, `clean`, `ext`, `isAbs` - Path manipulation
- OS-specific variants with `os` prefix

**Encoding Functions:**
- `b64enc`, `b64dec` - Base64 encoding/decoding
- `b32enc`, `b32dec` - Base32 encoding/decoding

**Crypto Functions:**
- `bcrypt`, `htpasswd` - Password hashing
- Certificate generation functions
- `encryptAES`, `decryptAES` - AES encryption/decryption

**Hash/Checksum Functions:**
- `sha1sum`, `sha256sum`, `sha512sum` - Hash functions
- `adler32sum` - Adler32 checksum

**Network Functions:**
- `getHostByName` - Gets host IP by name

**OS Environment Functions:**
- `env`, `expandenv` - Environment variables

**SemVer Functions:**
- `semver`, `semverCompare` - Semantic versioning

**Utility Functions:**
- `tuple`, `list` - Creates lists
- `until`, `untilStep` - Range functions
- `deepCopy` - Deep copies data structure

#### Custom functions provided by go-scaffold

In addition to the Sprig functions, go-scaffold provides these custom functions:

- `include` - Includes and renders a named template (useful as a replacement of
  the built in [template](https://pkg.go.dev/text/template#hdr-Actions), so it
  can be used in [pipelines](https://pkg.go.dev/text/template#hdr-Pipelines))
- `debug` - Prints debug information to console during template rendering, useful for troubleshooting variable values (e.g., `{{ debug .Values.myVariable }}` or `{{ debug "Debug info:" .Values.someValue }}`). Note: This function returns an empty string and doesn't affect template output.

#### Examples of custom functions usage

Here are some examples of how to use the custom functions in your templates:

**Using camelcase:**
```go
{{ "hello world" | camelcase }} → "HelloWorld"
{{ "my-package-name" | camelcase }} → "MyPackageName"
```

**Using replace:**
```go
{{ "hello world" | replace "world" "gopher" }} → "hello gopher"
{{ "foo-bar-baz" | replace "-" "_" }} → "foo_bar_baz"
```

**Using sequence:**
```go
{{ range $i := sequence 3 }}{{ $i }}{{ end }} → "012"
{{ range $i := sequence 5 }}Item {{ $i }}{{ end }} → "Item 0Item 1Item 2Item 3Item 4"
```

**Using include:**
```go
{{ include "partial_template_name" . }}
```

**Using debug:**
```go
{{ debug .Values.myVariable }}
{{ debug "Current value:" .Values.someValue }}
{{ debug "Multiple values:" .Values.val1 .Values.val2 }}
```
This function prints debug information to the console during template rendering and returns an empty string, so it doesn't affect the template output.
