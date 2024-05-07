package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: program_name <path>")
		return
	}
	path, err := filepath.Abs(os.Args[1])
	if err != nil {
		fmt.Printf("Not a valid path: %s", os.Args[1])
		return
	}

	dirSizes := make(map[string]int64)
	dirmap := make(map[string]bool)

	err = filepath.WalkDir(path, func(filePath string, d fs.DirEntry, err error) error {
		if err != nil {
			//fmt.Printf("Error accessing path %s: %v\n", filePath, err)
			return nil
		}
		file_info, info_err := d.Info()
		if info_err != nil {
			//fmt.Printf("Error accessing path %s: %v\n", filePath, err)
			return nil
		}
		rootPath := splitBySlashOrBackslash(strings.TrimPrefix(filePath, path))
		if !d.IsDir() {
			dirSizes[rootPath] += file_info.Size()
		} else if rootPath != "" {
			if d.IsDir() {
				dirmap[rootPath] = true
			}
			dirSizes[rootPath] += 0
		}
		return nil
	})

	if err != nil {
		//fmt.Printf("Error walking the path %s: %v\n", path, err)
		return
	}

	pairs := make([]Pair, 0, len(dirSizes))
	for key, value := range dirSizes {
		pairs = append(pairs, Pair{key, value})
	}

	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	max := findMax(pairs)

	fmt.Printf("\nDirectory: %s \n", path)
	fmt.Printf("%s\n", strings.Repeat("-", len(path)+13))
	for _, pair := range pairs {
		if _, ok := dirmap[pair.Key]; ok {
			if len(pair.Key) > 15 {
				fmt.Printf("\x1b[94m%-15.15s...\x1b[0m |%s| %s\n", pair.Key[:15], generateProgressBar(pair.Value, max, 80), humanReadableBytes(pair.Value))
			} else {
				fmt.Printf("\x1b[94m%-18.18s\x1b[0m |%s| %s\n", pair.Key, generateProgressBar(pair.Value, max, 80), humanReadableBytes(pair.Value))
			}
		} else {
			if len(pair.Key) > 15 {
				fmt.Printf("%-15.15s... |%s| %s\n", pair.Key[:15], generateProgressBar(pair.Value, max, 80), humanReadableBytes(pair.Value))
			} else {
				fmt.Printf("%-18.18s |%s| %s\n", pair.Key, generateProgressBar(pair.Value, max, 80), humanReadableBytes(pair.Value))
			}
		}
	}
}

func splitBySlashOrBackslash(s string) string {
	str_arr := strings.FieldsFunc(s, func(r rune) bool {
		return r == '/' || r == '\\'
	})
	if len(str_arr) > 0 {
		return str_arr[0]
	}
	return ""
}

func humanReadableBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return strconv.Itoa(int(bytes)) + " B"
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

type Pair struct {
	Key   string
	Value int64
}

func findMax(data []Pair) int64 {
	var max int64
	for _, pair := range data {
		if pair.Value > max {
			max = pair.Value
		}
	}
	return max
}

func generateProgressBar(value, max int64, width int) string {
	percentage := float64(value) / float64(max)
	progressBarWidth := int(percentage * float64(width))

	progressBar := strings.Repeat("\u2588", progressBarWidth)
	spaces := strings.Repeat(" ", width-progressBarWidth)

	return progressBar + spaces
}
