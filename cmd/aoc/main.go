package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/foax/aoc2025/internal/day01"
	"github.com/foax/aoc2025/internal/day02"
	"github.com/foax/aoc2025/internal/day03"
	"github.com/foax/aoc2025/internal/day04"
	"github.com/foax/aoc2025/internal/day05"
	"github.com/foax/aoc2025/internal/day06"
	"github.com/foax/aoc2025/internal/day07"
	"github.com/foax/aoc2025/internal/day08"
	"github.com/foax/aoc2025/internal/day09"
)

const (
	input_directory = "input"
	aoc_edition     = "aoc2025"
)

type PartSolver func([]string) (string, error)

var days = map[int][]PartSolver{
	1: {day01.Part1, day01.Part2},
	2: {day02.Part1, day02.Part2},
	3: {day03.Part1, day03.Part2},
	4: {day04.Part1, day04.Part2},
	5: {day05.Part1, day05.Part2},
	6: {day06.Part1, day06.Part2},
	7: {day07.Part1, day07.Part2},
	8: {day08.Part1, day08.Part2},
	9: {day09.Part1, day09.Part2},
}

func main() {
	var logLevel slog.Level
	logLevelFlag := flag.String("loglevel", "INFO", "Log level to use for output")
	flag.Parse()
	switch strings.ToLower(*logLevelFlag) {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		panic(fmt.Sprintf("Invalid log level provided: %v", *logLevelFlag))
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: logLevel,
	})).With("aoc", aoc_edition, "section", "main")

	logger.Info("Start", "loglevel", logLevel)
	defer logger.Info("End")
	logger.Debug("flag args", "args", flag.Args())

	args := flag.Args()
	if len(args) == 0 {
		logger.Error("No day argument passed on the command line")
		os.Exit(1)
	}
	dayNum, err := strconv.Atoi(args[0])
	if err != nil {
		logger.Error("Invalid day argument passed on the command line", "arg", args[0], "error", err)
		os.Exit(1)
	}
	day, ok := days[dayNum]
	if !ok {
		logger.Error("No solution implemented for requested day", "day", dayNum)
		os.Exit(1)
	}

	var reader io.Reader
	info, err := os.Stdin.Stat()
	if err != nil {
		logger.Error("Error checking stdin", "error", err)
		os.Exit(1)
	}
	if (info.Mode() & os.ModeCharDevice) == 0 {
		logger.Info("Reading input from stdin")
		reader = os.Stdin
	} else {
		logger.Debug("stdin is a terminal (no piped input)")
		filename := fmt.Sprintf("%s/day%02d.txt", input_directory, dayNum)
		reader, err = os.Open(filename)
		if err != nil {
			logger.Error("Error opening input file", "filename", filename, "error", err)
			os.Exit(1)
		}
		logger.Info("Reading input from input file", "filename", filename)
	}
	scanner := bufio.NewScanner(reader)

	var input []string
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	logger.Debug("Lines of input read", "lines", len(input))

	for partIdx, part := range day {
		partNum := partIdx + 1
		logger.Info("Solving", "day", dayNum, "part", partNum)
		result, err := part(input)
		if err != nil {
			logger.Error("Error solving", "day", dayNum, "part", partNum, "error", err)
			os.Exit(1)
		}
		fmt.Printf("%s/day%d/part%d: %s\n", aoc_edition, dayNum, partNum, result)
		logger.Info("Result", "part", partNum, "result", result)
	}
}
