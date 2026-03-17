// ps-gen generates the file param-set-auto.go which contains all the
// Bind* methods on ParamSet[N]. To regenerate, run:
//
//	go run ./cmd/ps-gen/main.go
//
// The output is written to: assist/param-set-auto.go
package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"text/template"
)

// outputPath is the destination file relative to the working directory from
// which you invoke `go run`.  Adjust this to match your repository layout.
const outputPath = "assist/param-set-auto.go"

// typeSpec describes a single Bind* method (or Bind*Slice pair).
type typeSpec struct {
	// GoType is the Go type name used in the method signature, e.g. "bool",
	// "int32", "net.IPMask".
	GoType string

	// FlagType is the pflag method suffix, e.g. "Bool", "Int32", "IPMask".
	// For BindEnum this differs from the method name.
	FlagType string

	// MethodName overrides the generated Bind<MethodName> name when it differs
	// from FlagType (only needed for BindEnum).
	MethodName string

	// HasSlice indicates whether a Bind*Slice variant should also be generated.
	HasSlice bool

	// SliceGoType overrides the slice element type when it differs from GoType
	// (not currently needed, but kept for extensibility).
	SliceGoType string

	// Comment is the doc-comment body (the part after "binds ").
	Comment string

	// SliceComment is the doc-comment body for the slice variant.
	SliceComment string
}

func (s typeSpec) Resolve() typeSpec {
	if s.MethodName == "" {
		s.MethodName = s.FlagType
	}
	if s.SliceGoType == "" {
		s.SliceGoType = s.GoType
	}
	return s
}

var specs = []typeSpec{
	{
		GoType:       "bool",
		FlagType:     "Bool",
		HasSlice:     true,
		Comment:      "bool slice flag with a shorthand if\n// 'info.Short' has been set otherwise binds without a short name.",
		SliceComment: "[]bool slice flag with a shorthand if 'info.Short' has been set\n// otherwise binds without a short name.",
	},
	{
		GoType:       "time.Duration",
		FlagType:     "Duration",
		HasSlice:     true,
		Comment:      "time.Duration slice flag with a shorthand if\n// 'info.Short' has been set otherwise binds without a short name.",
		SliceComment: "[]time.Duration slice flag with a shorthand if 'info.Short' has been set\n// otherwise binds without a short name.",
	},
	{
		GoType:     "string",
		FlagType:   "String",
		MethodName: "Enum",
		HasSlice:   false,
		Comment:    "enum slice flag with a shorthand if\n// 'info.Short' has been set otherwise binds without a short name.",
	},
	{
		GoType:       "float32",
		FlagType:     "Float32",
		HasSlice:     true,
		Comment:      "float32 slice flag with a shorthand if\n// 'info.Short' has been set otherwise binds without a short name.",
		SliceComment: "[]float32 slice flag with a shorthand if 'info.Short' has been set\n// otherwise binds without a short name.",
	},
	{
		GoType:       "float64",
		FlagType:     "Float64",
		HasSlice:     true,
		Comment:      "float64 slice flag with a shorthand if\n// 'info.Short' has been set otherwise binds without a short name.",
		SliceComment: "[]float64 slice flag with a shorthand if 'info.Short' has been set\n// otherwise binds without a short name.",
	},
	{
		GoType:       "int",
		FlagType:     "Int",
		HasSlice:     true,
		Comment:      "int slice flag with a shorthand if\n// 'info.Short' has been set otherwise binds without a short name.",
		SliceComment: "[]int slice flag with a shorthand if 'info.Short' has been set\n// otherwise binds without a short name.",
	},
	{
		GoType:   "int16",
		FlagType: "Int16",
		HasSlice: false,
		Comment:  "int16 slice flag with a shorthand if\n// 'info.Short' has been set otherwise binds without a short name.",
	},
	{
		GoType:       "int32",
		FlagType:     "Int32",
		HasSlice:     true,
		Comment:      "int32 slice flag with a shorthand if\n// 'info.Short' has been set otherwise binds without a short name.",
		SliceComment: "[]int32 slice flag with a shorthand if 'info.Short' has been set\n// otherwise binds without a short name.",
	},
	{
		GoType:       "int64",
		FlagType:     "Int64",
		HasSlice:     true,
		Comment:      "int64 slice flag with a shorthand if\n// 'info.Short' has been set otherwise binds without a short name.",
		SliceComment: "[]int64 slice flag with a shorthand if 'info.Short' has been set\n// otherwise binds without a short name.",
	},
	{
		GoType:   "int8",
		FlagType: "Int8",
		HasSlice: false,
		Comment:  "int8 slice flag with a shorthand if\n// 'info.Short' has been set otherwise binds without a short name.",
	},
	{
		GoType:   "net.IPMask",
		FlagType: "IPMask",
		HasSlice: false,
		Comment:  "net.IPMask slice flag with a shorthand if\n// 'info.Short' has been set otherwise binds without a short name.",
	},
	{
		GoType:   "net.IPNet",
		FlagType: "IPNet",
		HasSlice: false,
		Comment:  "net.IPNet slice flag with a shorthand if\n// 'info.Short' has been set otherwise binds without a short name.",
	},
	{
		GoType:       "string",
		FlagType:     "String",
		HasSlice:     true,
		Comment:      "string slice flag with a shorthand if\n// 'info.Short' has been set otherwise binds without a short name.",
		SliceComment: "[]string slice flag with a shorthand if 'info.Short' has been set\n// otherwise binds without a short name.",
	},
	{
		GoType:   "uint16",
		FlagType: "Uint16",
		HasSlice: false,
		Comment:  "uint16 slice flag with a shorthand if\n// 'info.Short' has been set otherwise binds without a short name.",
	},
	{
		GoType:   "uint32",
		FlagType: "Uint32",
		HasSlice: false,
		Comment:  "uint32 slice flag with a shorthand if\n// 'info.Short' has been set otherwise binds without a short name.",
	},
	{
		GoType:   "uint64",
		FlagType: "Uint64",
		HasSlice: false,
		Comment:  "uint64 slice flag with a shorthand if\n// 'info.Short' has been set otherwise binds without a short name.",
	},
	{
		GoType:   "uint8",
		FlagType: "Uint8",
		HasSlice: false,
		Comment:  "uint8 slice flag with a shorthand if\n// 'info.Short' has been set otherwise binds without a short name.",
	},
	{
		GoType:       "uint",
		FlagType:     "Uint",
		HasSlice:     true,
		Comment:      "uint slice flag with a shorthand if\n// 'info.Short' has been set otherwise binds without a short name.",
		SliceComment: "[]uint slice flag with a shorthand if 'info.Short' has been set\n// otherwise binds without a short name.",
	},
}

const tmplSrc = `// Code generated by gen-param-set; DO NOT EDIT.

package assist

import (
	"net"
	"time"
)
{{range .}}{{$s := .}}
// Bind{{$s.MethodName}} binds {{$s.Comment}}
func (params *ParamSet[N]) Bind{{$s.MethodName}}(info *FlagInfo, to *{{$s.GoType}}) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.{{$s.FlagType}}Var(to, info.FlagName(), info.Default.({{$s.GoType}}), info.Usage)
	} else {
		flagSet.{{$s.FlagType}}VarP(to, info.FlagName(), info.Short, info.Default.({{$s.GoType}}), info.Usage)
	}

	return params
}
{{if $s.HasSlice}}
// Bind{{$s.MethodName}}Slice binds {{$s.SliceComment}}
func (params *ParamSet[N]) Bind{{$s.MethodName}}Slice(info *FlagInfo, to *[]{{$s.SliceGoType}}) *ParamSet[N] {
	flagSet := params.ResolveFlagSet(info)
	if info.Short == "" {
		flagSet.{{$s.FlagType}}SliceVar(to, info.FlagName(), info.Default.([]{{$s.SliceGoType}}), info.Usage)
	} else {
		flagSet.{{$s.FlagType}}SliceVarP(to, info.FlagName(), info.Short, info.Default.([]{{$s.SliceGoType}}), info.Usage)
	}

	return params
}
{{end}}{{end}}`

func main() {
	// Resolve all specs (fill in derived fields).
	resolved := make([]typeSpec, len(specs))
	for i, s := range specs {
		resolved[i] = s.Resolve()
	}

	tmpl, err := template.New("param-set-auto").Parse(tmplSrc)
	if err != nil {
		log.Fatalf("parse template: %v", err)
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, resolved); err != nil {
		log.Fatalf("execute template: %v", err)
	}

	// Run through gofmt so the output is canonical.
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		// Emit the raw output to make debugging easier.
		fmt.Fprintf(os.Stderr, "--- raw output ---\n%s\n--- end ---\n", buf.String())
		log.Fatalf("gofmt: %v", err)
	}

	if err := os.WriteFile(outputPath, formatted, 0o644); err != nil { //nolint:gosec // ok
		log.Fatalf("write %s: %v", outputPath, err)
	}

	fmt.Printf("wrote %s\n", outputPath)
}
