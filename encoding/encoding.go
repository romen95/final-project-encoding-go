package encoding

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

type NotFoundJSONError struct {
	JSONData JSONData
}

func (err NotFoundJSONError) Error() string {
	return fmt.Sprintf("Данные в структуре %d не найдены", err.JSONData.DockerCompose)
}

type NotFoundYAMLrror struct {
	YAMLData YAMLData
}

func (err NotFoundYAMLrror) Error() string {
	return fmt.Sprintf("Данные в структуре %d не найдены", err.YAMLData.DockerCompose)
}

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {

	// создаем файл yamlOutput.yml
	f, err := os.Create(j.FileOutput)
	if err != nil {
		fmt.Printf("ошибка при создании файла: %s", err.Error())
		return err
	}

	// когда программа завершится, надо закрыть дескриптор файла
	defer f.Close()

	// читаем файл jsonInput.json
	jsonFile, err := os.ReadFile(j.FileInput)
	if err != nil {
		fmt.Printf("ошибка чтения файла: %s", err.Error())
		return err
	}

	// десериализуем
	if err = yaml.Unmarshal(jsonFile, &j.DockerCompose); err != nil {
		var err NotFoundJSONError
		err.JSONData.DockerCompose = j.DockerCompose
		return err
	}

	// сериализуем
	yamlBytes, err := yaml.Marshal(&j.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка при сериализации yaml: %s", err.Error())
		return err
	}

	// записываем слайс байт в файл
	_, err = f.Write(yamlBytes)
	if err != nil {
		fmt.Printf("ошибка при записи данных в файл: %s", err.Error())
		return err
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {

	// создаем файл jsonOutput.json
	f, err := os.Create(y.FileOutput)
	if err != nil {
		fmt.Printf("ошибка при создании файла: %s", err.Error())
		return err
	}

	// когда программа завершится, надо закрыть дескриптор файла
	defer f.Close()

	// читаем файл yamlInput.yaml
	yamlFile, err := os.ReadFile(y.FileInput)
	if err != nil {
		fmt.Printf("ошибка чтения файла: %s", err.Error())
		return err
	}

	// десериализуем
	if err = yaml.Unmarshal(yamlFile, &y.DockerCompose); err != nil {
		var err NotFoundYAMLrror
		err.YAMLData.DockerCompose = y.DockerCompose
		return err
	}

	// сериализуем
	jsonBytes, err := json.Marshal(&y.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка при сериализации yaml: %s", err.Error())
		return err
	}

	// записываем слайс байт в файл
	_, err = f.Write(jsonBytes)
	if err != nil {
		fmt.Printf("ошибка при записи данных в файл: %s", err.Error())
		return err
	}

	return nil
}
