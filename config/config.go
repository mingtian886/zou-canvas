package config

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

const legacyDefaultJWTSecret = "infinite-canvas"

type Config struct {
	Port           string `env:"PORT" envDefault:"8080"`
	AdminUsername  string `env:"ADMIN_USERNAME" envDefault:"admin"`
	AdminPassword  string `env:"ADMIN_PASSWORD"`
	JWTSecret      string `env:"JWT_SECRET"`
	JWTExpireHours int    `env:"JWT_EXPIRE_HOURS" envDefault:"168"`
	StorageDriver  string `env:"STORAGE_DRIVER" envDefault:"sqlite"`
	DatabaseDSN    string `env:"DATABASE_DSN" envDefault:"data/infinite-canvas.db"`
	AILogDir       string `env:"AI_LOG_DIR" envDefault:"data/logs/ai-calls"`
}

var Cfg Config

func Load() error {
	_ = godotenv.Load()
	if err := env.Parse(&Cfg); err != nil {
		return err
	}
	normalizeDockerSQLiteDSN("/app/data")
	return resolveJWTSecret()
}

func resolveJWTSecret() error {
	if secret := strings.TrimSpace(Cfg.JWTSecret); secret != "" && secret != legacyDefaultJWTSecret {
		return nil
	}
	if secret, err := ReadDataFile("jwt-secret.txt"); err != nil {
		return err
	} else if secret = strings.TrimSpace(secret); secret != "" && secret != legacyDefaultJWTSecret {
		Cfg.JWTSecret = secret
		return nil
	}
	secret, err := randomSecret()
	if err != nil {
		return err
	}
	if _, err := WriteDataFile("jwt-secret.txt", secret+"\n"); err != nil {
		return err
	}
	Cfg.JWTSecret = secret
	return nil
}

func DataFilePath(name string) string {
	dsn := strings.TrimSpace(Cfg.DatabaseDSN)
	if strings.EqualFold(strings.TrimSpace(Cfg.StorageDriver), "sqlite") && dsn != "" && dsn != ":memory:" {
		path, _, _ := strings.Cut(dsn, "?")
		if filepath.IsAbs(path) {
			return filepath.Join(filepath.Dir(path), name)
		}
	}
	return filepath.Join("data", name)
}

func ReadDataFile(name string) (string, error) {
	data, err := os.ReadFile(DataFilePath(name))
	if errors.Is(err, os.ErrNotExist) {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func WriteDataFile(name, content string) (string, error) {
	path := DataFilePath(name)
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return "", err
	}
	if err := os.WriteFile(path, []byte(content), 0600); err != nil {
		return "", err
	}
	if err := os.Chmod(path, 0600); err != nil {
		return "", err
	}
	return path, nil
}

func RemoveDataFile(name string) error {
	err := os.Remove(DataFilePath(name))
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return err
}

func RandomSecret() (string, error) {
	return randomSecret()
}

func normalizeDockerSQLiteDSN(appDataDir string) {
	driver := strings.ToLower(strings.TrimSpace(Cfg.StorageDriver))
	if driver != "" && driver != "sqlite" {
		return
	}
	dsn := strings.TrimSpace(Cfg.DatabaseDSN)
	if dsn == "" || dsn == ":memory:" || strings.HasPrefix(dsn, "file:") {
		return
	}
	pathPart, suffix := dsn, ""
	if index := strings.Index(dsn, "?"); index >= 0 {
		pathPart = dsn[:index]
		suffix = dsn[index:]
	}
	if filepath.IsAbs(pathPart) {
		return
	}
	slashPath := filepath.ToSlash(pathPart)
	if slashPath != "data" && !strings.HasPrefix(slashPath, "data/") {
		return
	}
	if _, err := os.Stat(appDataDir); err != nil {
		return
	}
	Cfg.DatabaseDSN = filepath.Join(filepath.Dir(appDataDir), filepath.FromSlash(slashPath)) + suffix
}

func randomSecret() (string, error) {
	buf := make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(buf), nil
}
