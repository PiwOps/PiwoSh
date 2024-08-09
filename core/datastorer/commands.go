package datastorer

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

func serialize(value interface{}) (string, error) {
	switch v := value.(type) {
	case string:
		return v, nil
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", v), nil
	case uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v), nil
	case float32, float64:
		return fmt.Sprintf("%f", v), nil
	case bool:
		return strconv.FormatBool(v), nil
	default:
		return "", fmt.Errorf("unsupported type: %T", value)
	}
}

func deserialize(value string, result interface{}) error {
	switch v := result.(type) {
	case *string:
		*v = value
	case *int:
		i, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		*v = i
	case *bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		*v = b
	case *float64:
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		*v = f
	default:
		return fmt.Errorf("unsupported type: %T", result)
	}
	return nil
}

func GetCmdConfigValue(key string, result interface{}) error {
	var cmdConfig CommandsConfig
	dbResult := DB.Where("key = ?", key).First(&cmdConfig)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	return deserialize(cmdConfig.Value, result)
}

func SetCmdConfigValue(key string, value interface{}) error {
	var cmdConfig CommandsConfig
	result := DB.Where("key = ?", key).First(&cmdConfig)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return result.Error
	}

	serializedValue, err := serialize(value)
	if err != nil {
		return err
	}

	cmdConfig.Key = key
	cmdConfig.Value = serializedValue

	if result.Error == gorm.ErrRecordNotFound {
		return DB.Create(&cmdConfig).Error
	}

	return DB.Save(&cmdConfig).Error
}

func DeleteCmdConfigValue(key string) error {
	return DB.Where("key = ?", key).Delete(&CommandsConfig{}).Error
}

func GetAllCmdConfigValues() (map[string]string, error) {
	var configs []CommandsConfig
	if err := DB.Find(&configs).Error; err != nil {
		return nil, err
	}

	configMap := make(map[string]string)
	for _, config := range configs {
		configMap[config.Key] = config.Value
	}

	return configMap, nil
}