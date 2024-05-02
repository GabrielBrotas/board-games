package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

const promptContext = `Voce é o Organizador do jogo 'Quem é o Impostor?', voce vai ter que selecionar uma palavra para ser distribuída para os jogadores e eles tem que adivinhar quem é o impostor entre eles.
Regras:
- Palavras podem ser simples ("Canario", "Hipopótamo") ou composta ("Beija flor", "Guarda Roupa", "Todo Mundo Odeia o Chris").
- Para nome de objetos, pessoa, animais ou cidades utilize o nome completo (por exemplo ao inves de Pitt utilize Brad Pitt, ao inves de Rio utilize Rio de Janeiro)
- As palavras geradas devem ser conhecidas no Brasil ou conhecidas mundialmente
- Caso uma categorias seja fornecida, pode ser qualquer coisa relacionada aquela categoria (por exemplo, na categoria 'esporte' pode ser nome de atletas famosos, nome de esportes, nome de instrumentos ou acessorios utilizados em um esporte)
- Caso não seja fornecida uma categoria, a palavra pode ser qualquer coisa
- A palavra não pode ser uma palavra que seja impossível de ser representada ou adivinhada
- Retorne apenas a palavra selecionada, nenhuma outra informação deve ser retornada

A rodada vai começar,`

const topicos = "Esporte, Comida, Profissão, Animal, Objeto, Lugar, Pessoa, Filme, Música, Série, Desenho, Marca, Jogo, Instrumento, Transporte, Fruta, Verdura, Doce, Bebida, País, Cidade, Estado, Planeta, Astro, Elemento Químico, Personagens Famosos."

var usedWords = make(map[string]bool)

func generateWordFromOpenAI(category string, difficulty string) (string, error) {
	openAIKey := os.Getenv("OPENAI_KEY")
	if openAIKey == "" {
		return "", fmt.Errorf("OPENAI_KEY environment variable not set")
	}

	client := openai.NewClient(openAIKey)

	prompt := ""
	if category == "" {
		prompt = promptContext + " por favor selecione uma palavra aleatória para ser distribuída para os jogadores relacionada a qualquer categoria dentre as seguintes: " + topicos
	} else {
		prompt = promptContext + " por favor selecione uma palavra relacionada a categoria '" + category + "' para ser distribuída para os jogadores."
	}

	if difficulty != "" {
		prompt = prompt + " A palavra deve ser de dificuldade " + difficulty + "."
	}

	// Attempt to generate a unique word
	var word string
	var err error
	for attempts := 0; attempts < 10; attempts++ { // Limit the number of attempts to avoid infinite loops
		word, err = attemptToGenerateUniqueWord(client, prompt)
		if err != nil {
			log.Printf("Error generating word from OpenAI: %v", err)
			return "", err
		}
		if _, exists := usedWords[word]; !exists {
			usedWords[word] = true
			return word, nil
		}
		// Log the event of a duplicate word
		log.Println("Duplicate word generated, attempting again...")
		time.Sleep(1) // Sleep for 1 second to avoid rate limiting
	}
	fmt.Println("failed to generate a unique word after several attempts")
	return word, nil
}

func attemptToGenerateUniqueWord(client *openai.Client, prompt string) (string, error) {
	completionRequest := openai.ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	}

	completion, err := client.CreateChatCompletion(context.Background(), completionRequest)
	if err != nil {
		return "", err
	}

	if len(completion.Choices) > 0 && completion.Choices[0].Message.Content != "" {
		return completion.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no valid word returned from OpenAI")
}
