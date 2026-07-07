package validation

import (
	"testing"

	"github.com/punnch/ankiwords/internal/model"
)

func TestValidateGeneratedCards(t *testing.T) {
	formatter := model.SentenceFormatter{Color: "#00557f"}
	cards := []model.GeneratedCard{{
		Word: "collapse",
		Fields: map[string]string{
			"Sentence":      "The system will <span style=\"color:#00557f;\"><b>collapse</b></span> under pressure.",
			"Definition":    "fall apart",
			"Transcription": "/kəˈlæps/",
			"Word":          "collapse",
			"Translation":   "разрушаться",
		},
	}}

	fields := []string{"Sentence", "Definition", "Transcription", "Word", "Translation"}
	if err := ValidateGeneratedCards(cards, []string{"collapse"}, fields, formatter); err != nil {
		t.Fatal(err)
	}
}

func TestValidateGeneratedCardsAllowsFrontBackModel(t *testing.T) {
	formatter := model.SentenceFormatter{Color: "#00557f"}
	cards := []model.GeneratedCard{{
		Word: "collapse",
		Fields: map[string]string{
			"Front": "collapse",
			"Back":  "to fall apart; разрушаться",
		},
	}}

	if err := ValidateGeneratedCards(cards, []string{"collapse"}, []string{"Front", "Back"}, formatter); err != nil {
		t.Fatal(err)
	}
}

func TestValidateGeneratedCardsSplitsCommaSeparatedExpectedWords(t *testing.T) {
	formatter := model.SentenceFormatter{Color: "#00557f"}
	cards := []model.GeneratedCard{
		{
			Word: "lay out",
			Fields: map[string]string{
				"Sentence":      "I will <span style=\"color:#00557f;\"><b>lay out</b></span> the plan tonight.",
				"Definition":    "arrange or explain clearly",
				"Transcription": "/leI aUt/",
				"Word":          "lay out",
				"Translation":   "разложить",
			},
		},
		{
			Word: "home",
			Fields: map[string]string{
				"Sentence":      "She went <span style=\"color:#00557f;\"><b>home</b></span> after work.",
				"Definition":    "the place where someone lives",
				"Transcription": "/hoUm/",
				"Word":          "home",
				"Translation":   "дом",
			},
		},
	}

	fields := []string{"Sentence", "Definition", "Transcription", "Word", "Translation"}
	if err := ValidateGeneratedCards(cards, []string{"lay out, home"}, fields, formatter); err != nil {
		t.Fatal(err)
	}
}

func TestValidateGeneratedCardsAllowsHighlightedWordVariant(t *testing.T) {
	formatter := model.SentenceFormatter{Color: "#00557f"}
	cards := []model.GeneratedCard{{
		Word: "lay out",
		Fields: map[string]string{
			"Sentence":      "She <span style=\"color:#00557f;\"><b>laid out</b></span> the clean clothes on the bed.",
			"Definition":    "arrange or explain clearly",
			"Transcription": "/leI aUt/",
			"Word":          "lay out",
			"Translation":   "разложить",
		},
	}}

	fields := []string{"Sentence", "Definition", "Transcription", "Word", "Translation"}
	if err := ValidateGeneratedCards(cards, []string{"lay out"}, fields, formatter); err != nil {
		t.Fatal(err)
	}
}

func TestValidateGeneratedCardsRejectsSentenceWithoutHighlight(t *testing.T) {
	formatter := model.SentenceFormatter{Color: "#00557f"}
	cards := []model.GeneratedCard{{
		Word: "lay out",
		Fields: map[string]string{
			"Sentence":      "She laid out the clean clothes on the bed.",
			"Definition":    "arrange or explain clearly",
			"Transcription": "/leI aUt/",
			"Word":          "lay out",
			"Translation":   "разложить",
		},
	}}

	fields := []string{"Sentence", "Definition", "Transcription", "Word", "Translation"}
	if err := ValidateGeneratedCards(cards, []string{"lay out"}, fields, formatter); err == nil {
		t.Fatal("expected validation error")
	}
}

func TestValidateGeneratedCardsRejectsUnknownField(t *testing.T) {
	formatter := model.SentenceFormatter{Color: "#00557f"}
	cards := []model.GeneratedCard{{
		Word: "collapse",
		Fields: map[string]string{
			"Front": "collapse",
			"Back":  "to fall apart",
			"Extra": "unexpected",
		},
	}}

	err := ValidateGeneratedCards(cards, []string{"collapse"}, []string{"Front", "Back"}, formatter)
	if err == nil {
		t.Fatal("expected validation error")
	}
}
