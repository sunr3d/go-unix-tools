package myFind

import (
	"flag"
	"fmt"
	"regexp"
)

type options struct {
	showFiles bool
	showDirs  bool
	showLinks bool
	extFilter string
	filepath  string
}

func newOptions(args []string) (*options, error) {
	fs := flag.NewFlagSet("myFind", flag.ContinueOnError)
	showFiles := fs.Bool("f", false, "Показать файлы")
	showDirs := fs.Bool("d", false, "Показать директории")
	showLinks := fs.Bool("sl", false, "Показать ссылки")
	extFilter := fs.String("ext", "", "Фильтр по расширению (требует флаг -f)")

	if err := fs.Parse(args); err != nil {
		return nil, fmt.Errorf(errFlagParse, err)
	}

	if fs.NArg() != 1 {
		return nil, fmt.Errorf(errSinglePath)
	}

	opts := &options{
		showFiles: *showFiles,
		showDirs:  *showDirs,
		showLinks: *showLinks,
		extFilter: *extFilter,
		filepath:  fs.Arg(0),
	}

	return opts, opts.validate()
}

func (o *options) validate() error {
	if o.extFilter != "" {
		if !o.showFiles {
			return fmt.Errorf(errExtWithoutFile)
		}
		if matches, _ := regexp.MatchString(`^[a-zA-Z0-9]+$`, o.extFilter); !matches {
			return fmt.Errorf(errInvalidExt, o.extFilter)
		}
	}

	if !o.showFiles && !o.showDirs && !o.showLinks {
		o.showFiles = true
		o.showDirs = true
		o.showLinks = true
	}

	return nil
}
