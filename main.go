package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func main() {
	storageDir, _ := os.Getwd()

	for {
		now := time.Now()
		date := fmt.Sprintf("%d-%02d-%02d", now.Year(), now.Month(), now.Day())

		filePath := fmt.Sprintf("%s/%s.log", storageDir, date)

		out, err := exec.Command("bash", "-c", "ss -nt src :443 | tail -n +2 | awk '{print $5}' | cut -d: -f1 | sort | uniq").Output()
		// out, err := exec.Command("bash", "-c", "ss -nt src :80 | tail -n +2 | awk '{print $5}' | cut -d: -f4 | cut -d] -f1 | sort | uniq").Output()

		if err != nil {
			fmt.Println(err.Error())
		}

		output := string(out)

		ips := strings.Split(output, "\n")

		for _, ip := range ips {
			if ip == "" {
				continue
			}

			f, _ := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

			scanner := bufio.NewScanner(f)

			matched := false
			line := 1
			for scanner.Scan() {
				matched, _ = regexp.MatchString(fmt.Sprintf("^%s$", ip), scanner.Text())
				if matched {
					break
				}

				line++
			}

			if (!matched) {
				f.WriteString(fmt.Sprintf("%s\n", ip))
				fmt.Println(fmt.Sprintf("new visitor %s on %s", ip, date))
			}

			f.Close()
		}

		time.Sleep(1 * time.Second)
	}
}
