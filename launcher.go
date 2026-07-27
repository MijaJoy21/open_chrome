package main

import (
	"fmt"
	"os/exec"
	"time"
)

func LaunchChrome() {

	chromePath := `C:\Program Files\Google\Chrome\Application\chrome.exe`

	for _, profile := range profiles {

		args := []string{
			"--new-window",
			"--profile-directory=" + profile.Name,
		}

		args = append(args, profile.URLs...)

		cmd := exec.Command(chromePath, args...)

		err := cmd.Start()

		if err != nil {
			fmt.Printf("Gagal membuka %s\n", profile.Name)
		} else {
			fmt.Printf("Berhasil membuka %s\n", profile.Name)
		}

		time.Sleep(2 * time.Second)
	}

}
