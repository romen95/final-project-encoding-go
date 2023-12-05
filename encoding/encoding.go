package encoding

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

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

	var yamlData YAMLData

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
	if err = yaml.Unmarshal(jsonFile, &yamlData.DockerCompose); err != nil {
		fmt.Printf("ошибка при десериализации: %s", err.Error())
		return err
	}

	// сериализуем
	yamlBytes, err := yaml.Marshal(&yamlData.DockerCompose)
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

	var jsonData JSONData

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
	if err = yaml.Unmarshal(yamlFile, &jsonData.DockerCompose); err != nil {
		fmt.Printf("ошибка при десериализации: %s", err.Error())
		return err
	}

	// сериализуем
	jsonBytes, err := json.Marshal(&jsonData.DockerCompose)
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
