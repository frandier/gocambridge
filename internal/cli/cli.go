package cli

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/frandier/gocambridge/internal/cambridge"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/manifoldco/promptui"
)

func Run(url, cookie string) error {
	re := regexp.MustCompile(`([^/]+)/class/([^/]+)/product/([^/]+)`)
	matches := re.FindStringSubmatch(url)

	if len(matches) != 4 {
		return fmt.Errorf("invalid URL")
	}

	org := matches[1]
	idClass := matches[2]
	product := matches[3]
	cambridgeClient := cambridge.NewCambridge(org, product, cookie)
	resp, err := cambridgeClient.GetUnits(idClass)

	if err != nil {
		return err
	}

	for {
		err := clearTerminal()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		myFigure := figure.NewFigure("Gocambridge", "", true)
		myFigure.Print()

		var units []string
		for _, result := range resp.Toc.Result {
			units = append(units, result.Name)
		}

		unitsView := promptui.Select{
			Label: "Select Unit",
			Items: units,
		}

		unitSelectedIndex, _, _ := unitsView.Run()

		var sections []string
		for _, result := range resp.Toc.Result[unitSelectedIndex].Items {
			sections = append(sections, result.Name)
		}

		sectionsView := promptui.Select{
			Label: "Select Section",
			Items: sections,
		}

		sectionSelectedIndex, _, _ := sectionsView.Run()

		var lessons []string
		for _, result := range resp.Toc.Result[unitSelectedIndex].Items[sectionSelectedIndex].Items {
			lessons = append(lessons, result.Name)
		}

		lessonsView := promptui.Select{
			Label: "Select Lesson",
			Items: lessons,
		}

		lessonSelectedIndex, _, _ := lessonsView.Run()

		itemCode := resp.Toc.Result[unitSelectedIndex].Items[sectionSelectedIndex].Items[lessonSelectedIndex].ItemCode
		parts := strings.Split(itemCode, "/")
		lessonCode := parts[len(parts)-1]

		solution, err := cambridgeClient.GetLessonResponse(product, lessonCode)
		if err != nil {
			return err
		}

		tableRows := make([]table.Row, len(solution))
		for index, result := range solution {
			tableRows[index] = table.Row{index + 1, result}
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "Correct Answer"})
		t.AppendRows(tableRows)
		t.Render()

		prompt := promptui.Prompt{
			Label:     "Find another solution?",
			IsConfirm: true,
		}

		result, err := prompt.Run()
		if err != nil || strings.ToLower(result) != "y" {
			fmt.Println("Bye!")
			os.Exit(0)
		}
	}
}
