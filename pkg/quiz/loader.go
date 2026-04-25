package quiz

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

func LoadDir(dir string) ([]Quiz, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("não foi possível ler %s: %w", dir, err)
	}

	var paths []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(entry.Name()))
		if ext == ".yaml" || ext == ".yml" {
			paths = append(paths, filepath.Join(dir, entry.Name()))
		}
	}

	sort.Strings(paths)

	if len(paths) == 0 {
		return nil, fmt.Errorf("nenhum arquivo de quiz .yaml ou .yml encontrado em %s", dir)
	}

	quizzes := make([]Quiz, 0, len(paths))
	for _, path := range paths {
		quiz, err := loadFile(path)
		if err != nil {
			return nil, err
		}
		quizzes = append(quizzes, quiz)
	}

	return quizzes, nil
}

func loadFile(path string) (Quiz, error) {
	file, err := os.Open(path)
	if err != nil {
		return Quiz{}, fmt.Errorf("não foi possível abrir %s: %w", path, err)
	}
	defer file.Close()

	var quiz Quiz
	decoder := yaml.NewDecoder(file)
	decoder.KnownFields(true)
	if err := decoder.Decode(&quiz); err != nil {
		return Quiz{}, fmt.Errorf("arquivo %s inválido: %w", path, err)
	}

	if err := validate(quiz); err != nil {
		return Quiz{}, fmt.Errorf("arquivo %s inválido: %w", path, err)
	}

	return quiz, nil
}

func validate(quiz Quiz) error {
	if strings.TrimSpace(quiz.Name) == "" {
		return fmt.Errorf("o quiz precisa ter um nome")
	}
	if len(quiz.Questions) == 0 {
		return fmt.Errorf("o quiz precisa ter pelo menos uma pergunta")
	}

	for index, question := range quiz.Questions {
		if strings.TrimSpace(question.Topic) == "" {
			return fmt.Errorf("a pergunta %d está sem tópico", index+1)
		}
		if strings.TrimSpace(question.Prompt) == "" {
			return fmt.Errorf("a pergunta %d está sem enunciado", index+1)
		}
		if len(question.Options) < 2 {
			return fmt.Errorf("a pergunta %d precisa ter ao menos duas alternativas", index+1)
		}
		if question.Answer < 0 || question.Answer >= len(question.Options) {
			return fmt.Errorf("a pergunta %d tem índice de resposta inválido", index+1)
		}
	}

	return nil
}
