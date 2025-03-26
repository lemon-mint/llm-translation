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

var targetLanguageCustomPrompt = map[apiv1beta1.Language]string{
	apiv1beta1.Language_KOREAN: "Korean으로 번역할때 \"반드시\" 다음과 같은 말투를 사용하세요:\n\n- 말투: 모든 말의 끝을 \"~요\"로 사용하세요, 격식과 예의를 지켜서 친절하게 답변해주세요.\n- 예시:\n-     50%가 증가했어요.\n-     성능 개선을 목표로 해요.\n-     다음과 같은 거래를 '외상거래'라고 해요.\n-     김민수님 에게 1,0000원을 보낼게요.\n-     사자는 육식 동물이에요.\n-     2 × 2 = 4이므로 4 + 3y = 6이 돼요.\n-     오늘 날씨는 맑아요.\n-     가장 큰 7-10 double은 무엇인가요?\n-     파이썬으로 피보나치 수열을 구현해주세요.\n-     당신은 친절한 어시스턴트이에요.\n-     Gemma는 Google의 차세대 언어 모델이에요.\n-     여기서 좌측 항을 인수분해해요.\n",
}

func translate(ctx context.Context, model llm.Model, text string, from, to apiv1beta1.Language) (string, error) {
	if text == "" {
		return "", nil
	}

	toName := getLanguageName(to)
	custom := targetLanguageCustomPrompt[to]

	output, err := translation.TranslateTextCustomPrompt(ctx, model, text, toName, custom)
	if err != nil {
		return "", err
	}

	output = strings.TrimSpace(output)
	if len(output) > 0 {
		return output, nil
	}

	return "", ErrNoResponse
}
