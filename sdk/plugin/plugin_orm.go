/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package plugin

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"

	"go.osspkg.com/do"
	"go.osspkg.com/goppy/v2/orm"

	"go.arwos.org/arwos/sdk/env"
)

func createSchema(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	code := hex.EncodeToString(h.Sum(nil))
	s = strings.TrimSpace("plugin_" + code)
	return strings.ToLower(s)
}

func (a *_plugin) WithDatabase(tags []string) {
	tags = do.Unique[string](tags)
	envs := make([]env.Env, 0, 10)
	conf := `
pgsql:
`
	for _, tag := range tags {
		envSuffix := env.CanonicalName(tag)

		conf += fmt.Sprintf(`
    - tags: %[1]s
      host: @env(DB_%[2]s_HOST#127.0.0.1)
      port: @env(DB_%[2]s_PORT#5432)
      user: @env(DB_%[2]s_USER#postgres)
      password: @env(DB_%[2]s_PASSWRD#postgres)
      schema: @env(DB_%[2]s_SCHEMA#postgres)
      app_name: %[3]s
`, tag, envSuffix, strings.TrimSpace(a.info.Name))

		envs = append(envs,
			env.Env{
				Description: "PostgresSQL HOST for tag: " + tag,
				Key:         "DB_" + envSuffix + "_HOST", Default: "127.0.0.1",
			},
			env.Env{
				Description: "PostgresSQL PORT for tag: " + tag,
				Key:         "DB_" + envSuffix + "_PORT", Default: "5432",
			},
			env.Env{
				Description: "PostgresSQL USER for tag: " + tag,
				Key:         "DB_" + envSuffix + "_USER", Default: "postgres",
			},
			env.Env{
				Description: "PostgresSQL PASSWORD for tag: " + tag,
				Key:         "DB_" + envSuffix + "_PASSWRD", Default: "postgres",
			},
			env.Env{
				Description: "PostgresSQL SCHEMA for tag: " + tag,
				Key:         "DB_" + envSuffix + "_SCHEMA", Default: createSchema(a.info.Package),
			},
		)
	}

	a.WithDependencies(Dep{
		Inject: []any{
			orm.WithORM(),
			orm.WithPgsqlClient(),
		},
		Config: conf,
		Envs:   envs,
	})
}
