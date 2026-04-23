package name

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/boundedinfinity/canonical_model/idiomatic/util/errorer"
	"github.com/boundedinfinity/canonical_model/idiomatic/util/stringer"
)

func Email1(address string, truncate int) *Formatter {
	return New(fmt.Sprintf(
		"{given[1]:normalize[%d],lower}.{family[-1]:normalize[%d],lower}@%s",
		truncate, truncate, address,
	))
}

func WesternFormatter() *Formatter {
	return New("{given} {middle} {family}")
}

func EasternFormatter() *Formatter {
	return New("{family} {given} {middle}")
}

func New(pattern string) *Formatter {
	return &Formatter{
		Pattern: pattern,
	}
}

type Formatter struct {
	Pattern  string
	contexts []formatterContext
	all      formatterContext
	parsed   bool
}

var (
	regex1             = regexp.MustCompile(`{(.+?)}`)
	ErrFormat          = errorer.New("format error")
	ErrFormatPattern   = errorer.New("invalid pattern")
	errFormatPatternFn = errorer.Func(ErrFormat, ErrFormatPattern)
)

type formatterContext struct {
	part    string
	replace string
	count   int
	fns     []func(string) string
}

type counted struct {
	name  string
	count int
}

func getCounted(s string) (counted, error) {
	var counted counted
	var err error
	start := stringer.Index(s, "[")
	end := stringer.Index(s, "]")

	if start > 0 && end > 0 {
		counted.name = s[:start]
		n := s[start+1 : end]

		if count, err := strconv.Atoi(n); err != nil {
			return counted, err
		} else {
			counted.count = count
		}
	} else if start > 0 && end < 0 || start < 0 && end > 0 {
		return counted, errorer.New("unmatched brackets")
	} else {
		counted.name = s
	}

	return counted, err
}

func (this *Formatter) Parse() error {
	if this.parsed {
		return nil
	}

	//  errFormatPatternFn("unmatched brackets in %s", context.replace)
	matches := regex1.FindAllStringSubmatch(this.Pattern, -1)

	for _, match := range matches {
		context := formatterContext{
			replace: match[0],
		}

		components := stringer.Split(match[1], ":")

		if counted, err := getCounted(components[0]); err != nil {
			return errFormatPatternFn("invalid count in %s: %v", context.replace, err)
		} else {
			context.part = counted.name
			context.count = counted.count
		}

		if len(components) > 1 {
			for _, control := range stringer.Split(components[1], ",") {
				counted, err := getCounted(control)

				if err != nil {
					return errFormatPatternFn("invalid count in %s: %v", context.replace, err)
				}

				switch counted.name {
				case "lower":
					context.fns = append(context.fns, stringer.ToLower[string])
				case "upper":
					context.fns = append(context.fns, stringer.ToUpper[string])
				case "title":
					context.fns = append(context.fns, stringer.ToTitle[string])
				case "truncate":
					if counted.count < 0 {
						return errFormatPatternFn("%s : %s count must be positive number", context.replace, counted)
					}

					context.fns = append(context.fns, func(s string) string {
						return stringer.Truncate(s, counted.count)
					})
				case "normalize":
					if counted.count < 0 {
						return errFormatPatternFn("%s : %s count must be positive number", context.replace, counted)
					} else if counted.count > 0 {
						context.fns = append(context.fns, func(s string) string {
							return stringer.Truncate(s, counted.count)
						})
					}

					context.fns = append(context.fns, func(s string) string {
						s = stringer.RemoveDiacritics(s)
						return s
					})
				default:
					return errFormatPatternFn("invalid control %s in %s", control, context.replace)
				}
			}
		}

		switch context.part {
		case "all":
			this.all = context
		default:
			this.contexts = append(this.contexts, context)
		}
	}

	return nil
}

func (this Formatter) Format(name Name) string {
	if !this.parsed {
		if err := this.Parse(); err != nil {
			panic(err)
		}
	}

	var output = this.Pattern

	process := func(names []string, context formatterContext, fns []func(string) string) {
		if context.count > 0 {
			names = names[:context.count]
		} else if context.count < 0 {
			names = names[len(names)+context.count:]
		}

		for _, fn := range fns {
			for i, name := range names {
				names[i] = fn(name)
			}
		}

		text := stringer.Join(names, " ")
		output = stringer.ReplaceAll(output, context.replace, text)
	}

	if this.all.part != "" {
		output = stringer.ReplaceAll(output, this.all.replace, "")
	}

	for _, context := range this.contexts {
		fns := append(this.all.fns, context.fns...)

		switch context.part {
		case "given":
			process(name.Given, context, fns)
		case "family":
			process(name.Family, context, fns)
		case "middle":
			process(name.Middle, context, fns)
		case "all":
			// skip
		default:
			panic(fmt.Sprintf("invalid part %s in %s", context.part, context.replace))
		}
	}

	output = stringer.CompactSpace(output)
	output = stringer.TrimSpace(output)
	return output
}
