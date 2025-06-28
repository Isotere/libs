package dotenv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	stringEnvName = "ENV_STRING"
	intEnvName    = "ENV_INT"
)

func TestLoad(t *testing.T) {
	t.Run("load_envs", func(t *testing.T) {
		err := os.Unsetenv(stringEnvName)
		if err != nil {
			t.FailNow()
		}

		err = Load(".env.example")
		assert.NoError(t, err)

		env, ok := os.LookupEnv(stringEnvName)
		if !ok {
			t.Error("test env does not exist")
		}

		if env != "some_value" {
			t.Error("test env has incorrect value")
		}
	})

	t.Run("get_string_val", func(t *testing.T) {
		err := os.Unsetenv(stringEnvName)
		if err != nil {
			t.FailNow()
		}

		err = Load(".env.example")
		assert.NoError(t, err)

		strVal := GetString(stringEnvName, "default value")
		assert.Equal(t, "some_value", strVal)

		nonExistsValue := GetString("NO_ENV_NAME", "default value")
		assert.Equal(t, "default value", nonExistsValue)
	})

	t.Run("get_int_val", func(t *testing.T) {
		err := os.Unsetenv(stringEnvName)
		if err != nil {
			t.FailNow()
		}

		err = Load(".env.example")
		assert.NoError(t, err)

		intVal := GetInt(intEnvName, 333)
		assert.Equal(t, 123, intVal)

		nonExistsValue := GetInt("NO_ENV_NAME", 333)
		assert.Equal(t, 333, nonExistsValue)
	})
}
