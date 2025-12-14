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

var days = map[int]struct {
	Part1 func([]string) (string, error)
	Part2 func([]string) (string, error)
}{
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

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))
	slog.SetDefault(logger)

	slog.Info("Start", "aoc", aoc_edition, "section", "main", "loglevel", logLevel)
	defer slog.Info("End")
	slog.Debug("flag args", "aoc", aoc_edition, "section", "main", "args", flag.Args())

	args := flag.Args()
	if len(args) == 0 {
		slog.Error("No day argument passed on the command line", "aoc", aoc_edition, "section", "main")
		os.Exit(1)
	}
	dayNum, err := strconv.Atoi(args[0])
	if err != nil {
		slog.Error("Invalid day argument passed on the command line", "aoc", aoc_edition, "section", "main", "arg", args[0], "error", err)
		os.Exit(1)
	}
	day, ok := days[dayNum]
	if !ok {
		slog.Error("No solution implemented for requested day", "aoc", aoc_edition, "section", "main", "day", dayNum)
		os.Exit(1)
	}

	var reader io.Reader
	info, err := os.Stdin.Stat()
	if err != nil {
		slog.Error("Error checking stdin", "aoc", aoc_edition, "section", "main", "error", err)
		os.Exit(1)
	}
	if (info.Mode() & os.ModeCharDevice) == 0 {
		slog.Info("Reading input from stdin", "aoc", aoc_edition, "section", "main")
		reader = os.Stdin
	} else {
		slog.Debug("stdin is a terminal (no piped input)", "aoc", aoc_edition, "section", "main")
		filename := fmt.Sprintf("%s/day%02d.txt", input_directory, dayNum)
		reader, err = os.Open(filename)
		if err != nil {
			slog.Error("Error opening input file", "aoc", aoc_edition, "section", "main", "filename", filename, "error", err)
			os.Exit(1)
		}
		slog.Info("Reading input from input file", "aoc", aoc_edition, "section", "main", "filename", filename)
	}
	scanner := bufio.NewScanner(reader)

	var input []string
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	slog.Debug("Lines of input read", "aoc", aoc_edition, "section", "main", "lines", len(input))

	slog.Info("Solving", "aoc", aoc_edition, "section", "main", "day", dayNum, "part", 1)
	part1Result, err := day.Part1(input)
	if err != nil {
		slog.Error("Error solving", "aoc", aoc_edition, "section", "main", "day", dayNum, "part", 1, "error", err)
		os.Exit(1)
	}
	slog.Info("Result", "aoc", aoc_edition, "section", "main", "part", 1, "result", part1Result)

	slog.Info("Solving", "aoc", aoc_edition, "section", "main", "day", dayNum, "part", 2)
	part2Result, err := day.Part2(input)
	if err != nil {
		slog.Error("Error solving Part 2", "aoc", aoc_edition, "section", "main", "day", dayNum, "error", err)
		os.Exit(1)
	}
	slog.Info("Result", "aoc", aoc_edition, "section", "main", "part", 2, "result", part2Result)
}
