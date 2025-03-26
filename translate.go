package main

import (
	"context"
	"errors"
	"strings"

	"github.com/lemon-mint/coord/llm"
	apiv1beta1 "github.com/lemon-mint/llm-translation/api/v1beta1"
	"gosuda.org/deeplingua/translation"
)

func getLanguageName(lang apiv1beta1.Language) string {
	switch lang {
	case apiv1beta1.Language_ENGLISH:
		return "English"
	case apiv1beta1.Language_SPANISH:
		return "Spanish"
	case apiv1beta1.Language_CHINESE:
		return "Chinese"
	case apiv1beta1.Language_KOREAN:
		return "Korean"
	case apiv1beta1.Language_JAPANESE:
		return "Japanese"
	case apiv1beta1.Language_GERMAN:
		return "German"
	case apiv1beta1.Language_RUSSIAN:
		return "Russian"
	case apiv1beta1.Language_FRENCH:
		return "French"
	case apiv1beta1.Language_DUTCH:
		return "Dutch"
	case apiv1beta1.Language_ITALIAN:
		return "Italian"
	case apiv1beta1.Language_INDONESIAN:
		return "Indonesian"
	case apiv1beta1.Language_PORTUGUESE:
		return "Portuguese"
	case apiv1beta1.Language_TAIWANESE:
		return "Taiwanese"
	}

	return ""
}

var (
	ErrNoResponse = errors.New("no response")
)

func translate(ctx context.Context, model llm.Model, text string, from, to apiv1beta1.Language) (string, error) {
	if text == "" {
		return "", nil
	}

	toName := getLanguageName(to)

	output, err := translation.TranslateText(ctx, model, text, toName)
	if err != nil {
		return "", err
	}

	output = strings.TrimSpace(output)
	if len(output) > 0 {
		return output, nil
	}

	return "", ErrNoResponse
}
