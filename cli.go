package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/Songmu/prompter"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const clockingTypeIn string = "syussya"
const clockingTypeOut string = "taisya"

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

func clockIn(clockingOut bool) error {
	cws_url := os.Getenv("CWS_URL")
	cws_userid := os.Getenv("CWS_USERID")
	cws_password := os.Getenv("CWS_PASSWORD")
	clocking_endpoint := cws_url + "/srwtimerec"

	var clockingType string
	if clockingOut {
		clockingType = clockingTypeOut
	} else {
		clockingType = clockingTypeIn
	}

	values := url.Values{}
	values.Add("user_id", cws_userid)
	values.Add("password", cws_password)
	values.Add("dakoku", clockingType)

	res, err := http.PostForm(clocking_endpoint, values)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if 200 != res.StatusCode {
		return fmt.Errorf(res.Status)
	}

	doc, _ := goquery.NewDocumentFromReader(res.Body)

	if mes := doc.Find("table > tbody > tr > td > strong").Text(); mes != "\u00a0" {
		errMes, _ := decodeShiftJIS(mes)
		return errors.New(errMes)
	}

	clockingTime, err := decodeShiftJIS(doc.Find("table > tbody > tr > td > font > u").Text())
	if err != nil {
		return err
	}

	reg := regexp.MustCompile(`.+(\d\d)時(\d\d)分.+`)
	clockingTime = reg.ReplaceAllString(clockingTime, "$1:$2")

	fmt.Println(clockingTime)

	return nil
}

func decodeShiftJIS(s string) (string, error) {
	r, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(s), japanese.ShiftJIS.NewDecoder()))
	if err != nil {
		return "", err
	}
	return string(r), nil
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		yes    bool
		out    bool
		status bool

		version bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.BoolVar(&yes, "yes", false, "Skip y/n prompt")
	flags.BoolVar(&yes, "y", false, "Skip y/n prompt (Short)")
	flags.BoolVar(&out, "out", false, "Clocking out")
	flags.BoolVar(&out, "o", false, "Clocking out (Short)")
	flags.BoolVar(&status, "status", false, "Just show clock in/out time and exit")
	flags.BoolVar(&status, "s", false, "Just show clock in/out time and exit (Short)")

	flags.BoolVar(&version, "version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	if !yes && !prompter.YN("OK?", true) {
		fmt.Println("Canceled")
		return ExitCodeError
	}

	err := clockIn(out)
	if err != nil {
		fmt.Println(err)
		return ExitCodeError
	}

	_ = status

	return ExitCodeOK
}
