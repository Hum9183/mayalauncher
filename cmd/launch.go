package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"github.com/spf13/cobra"
)

const (
	MAYA_UI_LANGUAGE = "MAYA_UI_LANGUAGE"

	en_US = "en_US"
	ja_JP = "ja_JP"
	zh_CN = "zh_CN"
	en = "en"
	ja = "ja"
	zh = "zh"

	DEFAULT_LANGUAGE = "default"
	DEFAULT_VERSION = 2000

	MAYA = "Maya"

	LANGUAGE_WARNING_TEXT =
`Please select a language from the following.
- en_US
- en
- ja_JP
- ja
- zh_CN
- zh`
)

var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Command to start Maya",
	RunE: func(cmd *cobra.Command, args []string) error {
		language, err := cmd.Flags().GetString("language")
		if err != nil {
			return err
		}

		version, err := cmd.Flags().GetUint16("version")
		if err != nil {
			return err
		}

		err = setLanguage(language)
		if err != nil {
			return err
		}

		if version == DEFAULT_VERSION {
			latestVer, err := getLatestVer()
			if err != nil {
				return err
			}
			version = latestVer
		}

		launch(version)

		fmt.Printf("language: %s\n", language)
		fmt.Printf("version: %d\n", version)
		fmt.Printf("Maya is launching")
		return nil
	},
}

func setLanguage(language string) error {
	if language == DEFAULT_LANGUAGE {
		return nil
	}

	switch language {
		case en_US, ja_JP, zh_CN, en, ja, zh:
		default:
			return fmt.Errorf(LANGUAGE_WARNING_TEXT)
	}

	os.Setenv(MAYA_UI_LANGUAGE, language)
	return nil
}

func getLatestVer() (uint16, error) {
	dirs, err := os.ReadDir(`C:\Program Files\Autodesk`)
	if err != nil {
		return 0, err
	}

	pattern := fmt.Sprintf(`%s\d\d\d\d`, MAYA)
	re, err := regexp.Compile(pattern)
	if err != nil {
		return 0, err
	}

	latestVer := DEFAULT_VERSION
	for _, v := range dirs {
		matches := re.FindAllString(v.Name(), -1)
		if len(matches) == 0 {
			continue
		}

		verStr := strings.Replace(matches[0], MAYA, "", -1)
		verInt, _ := strconv.Atoi(verStr)

		if verInt > latestVer {
			latestVer = verInt
		}
	}

	return uint16(latestVer), err
}

func launch(version uint16) error {
	exePath := fmt.Sprintf(`C:\Program Files\Autodesk\Maya%d\bin\maya.exe`, version)
	err := exec.Command(exePath).Start()
	if err != nil {
		return err
	}
	return err
}

func init() {
	rootCmd.AddCommand(launchCmd)
	launchCmd.Flags().StringP("language", "l", DEFAULT_LANGUAGE, "Maya UI Language")
	launchCmd.Flags().Uint16P("version", "v", DEFAULT_VERSION, "Maya major version(e.g.2022)")
}
