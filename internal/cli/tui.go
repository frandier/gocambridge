package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/frandier/gocambridge/internal/cambridge"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/manifoldco/promptui"
)

func selectUnit(resp *cambridge.UnitsResult, cambridgeClient *cambridge.Cambridge, product string) error {
	clearAndPrintHeader()

	var units []string
	for _, result := range resp.Toc.Result {
		units = append(units, result.Name)
	}

	units, returnIndex := addReturnOption(units, true)

	unitsView := promptui.Select{
		Label: SelectUnitLabel,
		Items: units,
		Size:  SelectLimit,
	}

	unitSelectedIndex, _, err := unitsView.Run()
	if err != nil {
		return err
	}

	if unitSelectedIndex == returnIndex {
		fmt.Println(ExitMessage)
		os.Exit(0)
	}

	return selectSection(resp, cambridgeClient, product, unitSelectedIndex)
}

func selectSection(resp *cambridge.UnitsResult, cambridgeClient *cambridge.Cambridge, product string, unitSelectedIndex int) error {
	clearAndPrintHeader()

	var sections []string
	for _, result := range resp.Toc.Result[unitSelectedIndex].Items {
		sections = append(sections, result.Name)
	}

	sections, returnIndex := addReturnOption(sections, false)

	sectionsView := promptui.Select{
		Label: SelectSectionLabel,
		Items: sections,
		Size:  SelectLimit,
	}

	sectionSelectedIndex, _, err := sectionsView.Run()
	if err != nil {
		return err
	}

	if sectionSelectedIndex == returnIndex {
		return selectUnit(resp, cambridgeClient, product)
	}

	return selectLesson(resp, cambridgeClient, product, unitSelectedIndex, sectionSelectedIndex)
}

func selectLesson(resp *cambridge.UnitsResult, cambridgeClient *cambridge.Cambridge, product string, unitSelectedIndex, sectionSelectedIndex int) error {
	clearAndPrintHeader()

	var lessons []string
	for _, result := range resp.Toc.Result[unitSelectedIndex].Items[sectionSelectedIndex].Items {
		lessons = append(lessons, result.Name)
	}

	lessons, returnIndex := addReturnOption(lessons, false)

	lessonsView := promptui.Select{
		Label: SelectLessonLabel,
		Items: lessons,
		Size:  SelectLimit,
	}

	lessonSelectedIndex, _, err := lessonsView.Run()
	if err != nil {
		return err
	}

	if lessonSelectedIndex == returnIndex {
		return selectSection(resp, cambridgeClient, product, unitSelectedIndex)
	}

	itemCode := resp.Toc.Result[unitSelectedIndex].Items[sectionSelectedIndex].Items[lessonSelectedIndex].ItemCode
	parts := strings.Split(itemCode, "/")
	lessonCode := parts[len(parts)-1]

	return showSolution(cambridgeClient, product, lessonCode, resp, unitSelectedIndex, sectionSelectedIndex)
}

func showSolution(cambridgeClient *cambridge.Cambridge, product, lessonCode string, resp *cambridge.UnitsResult, unitSelectedIndex int, sectionSelectedIndex int) error {
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
	t.AppendHeader(table.Row{CorrectAnswersSymbol, CorrectAnswersLabel})
	t.AppendRows(tableRows)
	t.Render()

	return askForNextAction(resp, cambridgeClient, product, unitSelectedIndex, sectionSelectedIndex)
}

func askForNextAction(resp *cambridge.UnitsResult, cambridgeClient *cambridge.Cambridge, product string, unitSelectedIndex int, sectionSelectedIndex int) error {
	prompt := promptui.Prompt{
		Label: NextActionLabel,
	}

	result, err := prompt.Run()
	if err != nil {
		return err
	}

	switch strings.ToLower(result) {
	case InitialKey:
		return selectUnit(resp, cambridgeClient, product)
	case BackKey:
		return selectLesson(resp, cambridgeClient, product, unitSelectedIndex, sectionSelectedIndex)
	case ExitKey:
		fmt.Println(ExitMessage)
		os.Exit(0)
	default:
		fmt.Println(InvalidOptionMessage)
		return askForNextAction(resp, cambridgeClient, product, unitSelectedIndex, sectionSelectedIndex)
	}

	return nil
}
