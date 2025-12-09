package main

import (
	"bufio"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/foax/aoc2025/internal/day01"
	"github.com/foax/aoc2025/internal/day02"
	"github.com/foax/aoc2025/internal/day03"
	"github.com/foax/aoc2025/internal/day04"
	"github.com/foax/aoc2025/internal/day05"
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

	slog.Info("Start", "loglevel", logLevel)
	defer slog.Info("End")
	slog.Debug("flag args", "args", flag.Args())

	args := flag.Args()
	if len(args) == 0 {
		slog.Error("No day argument passed on the command line")
		os.Exit(1)
	}

	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	slog.Debug("Lines of input read", "lines", len(input))

	dayNum, err := strconv.Atoi(args[0])
	if err != nil {
		slog.Error("Invalid day argument passed on the command line", "arg", args[0], "error", err)
		os.Exit(1)
	}

	day, ok := days[dayNum]
	if !ok {
		slog.Error("No solution implemented for requested day", "day", dayNum)
		os.Exit(1)
	}

	slog.Info("Solving", "day", dayNum, "part", 1)
	part1Result, err := day.Part1(input)
	if err != nil {
		slog.Error("Error solving Part 1", "day", dayNum, "error", err)
		os.Exit(1)
	}
	slog.Info("Result", "part", 1, "result", part1Result)

	slog.Info("Solving", "day", dayNum, "part", 2)
	part2Result, err := day.Part2(input)
	if err != nil {
		slog.Error("Error solving Part 2", "day", dayNum, "error", err)
		os.Exit(1)
	}
	slog.Info("Result", "part", 2, "result", part2Result)
}
