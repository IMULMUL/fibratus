/*
 * Copyright 2019-2020 by Nedim Sabic Sabic
 * https://www.fibratus.io
 * All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package app

import (
	"fmt"
	"github.com/rabbitstack/fibratus/cmd/fibratus/common"
	"github.com/rabbitstack/fibratus/pkg/config"
	kerrors "github.com/rabbitstack/fibratus/pkg/errors"
	"github.com/rabbitstack/fibratus/pkg/util/rest"
	"github.com/spf13/cobra"
	"os"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show runtime config",
	RunE:  printConfig,
}

var (
	// config command options
	c = config.NewWithOpts(config.WithStats())
)

func init() {
	c.MustViperize(configCmd)
}

func printConfig(cmd *cobra.Command, args []string) error {
	if err := common.Init(c, false); err != nil {
		return err
	}
	body, err := rest.Get(rest.WithTransport(c.API.Transport), rest.WithURI("config"))
	if err != nil {
		return kerrors.ErrHTTPServerUnavailable(c.API.Transport, err)
	}
	_, err = fmt.Fprintln(os.Stdout, string(body))
	if err != nil {
		return err
	}
	return nil
}
