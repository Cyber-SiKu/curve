/*
 *  Copyright (c) 2023 NetEase Inc.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

/*
 * Project: CurveCli
 * Created Date: 2023-08-19
 * Author: chengyi (Cyber-SiKu)
 */

package main

import (
	"fmt"

	cmderror "github.com/opencurve/curve/tools-v2/internal/error"
	"github.com/opencurve/curve/tools-v2/pkg/cli/command/plugin"
	"github.com/opencurve/curve/tools-v2/pkg/config"
	"github.com/spf13/cobra"
)

type Hello struct {
	cmd *cobra.Command
}

var _ plugin.CurvePlugin = (*Hello)(nil)

func (h *Hello) Init(plugin *cobra.Command, args []string) error {
	h.cmd = plugin
	// add flag and re-parse
	config.AddBsMdsFlagOption(plugin)
	plugin.FParseErrWhitelist.UnknownFlags = true
	err := plugin.ParseFlags(args)
	if err != nil {
		return err
	}
	return nil
}

func (h *Hello) Run() error {
	fmt.Println("Hello, world!")
	mdsAddr, err := config.GetBsMdsAddrSlice(h.cmd)
	if err.TypeCode() != cmderror.CODE_SUCCESS {
		return err.ToError()
	}
	fmt.Println("mds addr:", mdsAddr)
	return nil
}

var Test Hello

func init() {
	plugin.InitCurvePlugin("Test")
}