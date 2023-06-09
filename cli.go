package percentile

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type options struct {
	file    string
	pValues []int
	rValue  bool
}

// CLI is the command line interface object
type CLI struct {
	InStream  io.Reader
	OutStream io.Writer
	ErrStream io.Writer
}

// Run executes the CLI
func (c *CLI) Run(args []string) error {
	opt, err := parseFlag()
	if err != nil {
		return fmt.Errorf("failed parse flag: %w", err)
	}

	f, err := openFile(opt.file)
	if err != nil {
		return fmt.Errorf("failed open file: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var numbers []float64
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return fmt.Errorf("failed parse float: %w", err)
		}
		numbers = append(numbers, num)
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed scan file: %w", err)
	}

	sort.Float64s(numbers)

	for _, v := range opt.pValues {
		p := Get(numbers, v)
		if opt.rValue {
			_, _ = fmt.Fprintf(c.OutStream, "p%d: %.1f \n", v, p)
		} else {
			_, _ = fmt.Fprintf(c.OutStream, "p%d: %d \n", v, int(p))
		}
	}

	return nil
}

func parseFlag() (*options, error) {
	pOption := flag.String("p", "25,50,75,80,90,95,99", "Specify percentiles (comma-separated list of integers)")
	rOption := flag.Bool("r", false, "Don't Round percentile values")
	flag.Parse()

	file := "-"
	if len(flag.Args()) > 0 {
		file = flag.Args()[0]
	}

	// -p オプションをパースする
	var pValues []int
	for _, v := range strings.Split(*pOption, ",") {
		p, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("percentile must only contain integers: %w", err)
		}
		if p < 0 || p > 100 {
			return nil, fmt.Errorf("percentile must be between 0 and 100: %w", err)
		}
		pValues = append(pValues, p)
	}

	opt := &options{
		file:    file,
		pValues: pValues,
		rValue:  *rOption,
	}
	return opt, nil
}

func openFile(file string) (io.ReadCloser, error) {
	if file == "-" {
		return io.NopCloser(os.Stdin), nil
	}
	return os.Open(file)
}
