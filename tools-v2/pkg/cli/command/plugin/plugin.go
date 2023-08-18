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
 * Created Date: 2023-08-18
 * Author: chengyi (Cyber-SiKu)
 */

package plugin

import (
	goplugin "plugin"

	basecmd "github.com/opencurve/curve/tools-v2/pkg/cli/command"
	"github.com/opencurve/curve/tools-v2/pkg/config"
	"github.com/opencurve/curve/tools-v2/pkg/output"
	"github.com/spf13/cobra"
)

var (
	PluginName string
)

type PluginCommand struct {
	basecmd.FinalCurveCmd
}

var _ basecmd.FinalCurveCmdFunc = (*PluginCommand)(nil) // check interface

func NewPluginCurveCommand() *cobra.Command {
	return NewPluginCommand().Cmd
}

func NewPluginCommand() *PluginCommand {
	pluginCmd := &PluginCommand{
		FinalCurveCmd: basecmd.FinalCurveCmd{
			Use:     "plugin",
			Short:   "plugin for curve",
			// Example: "",
		},
	}

	basecmd.NewFinalCurveCli(&pluginCmd.FinalCurveCmd, pluginCmd)
	// delay parse flags
	pluginCmd.Cmd.DisableFlagParsing = true
	return pluginCmd
}

func (pCmd *PluginCommand) AddFlags() {
	config.AddSoRequiredFlag(pCmd.Cmd)
}

func (pCmd *PluginCommand) Init(cmd *cobra.Command, args []string) error {
	// enable flag parse and ignore unknown flags
	pCmd.Cmd.DisableFlagParsing = false
	pCmd.Cmd.FParseErrWhitelist.UnknownFlags = true
	err := pCmd.Cmd.ParseFlags(args)
	if err != nil {
		return err
	}
	soPath := config.GetFlagString(pCmd.Cmd, config.SO)
	p, err := goplugin.Open(soPath)
	if err != nil {
		return err
	}

	v, err := p.Lookup(PluginName)
	if err != nil {
		return err
	}
	curvePlugin := v.(CurvePlugin)
	curvePlugin.Init(cmd, args)
	return curvePlugin.Run()
}

func (pCmd *PluginCommand) Print(cmd *cobra.Command, args []string) error {
	return output.FinalCmdOutput(&pCmd.FinalCurveCmd, pCmd)
}

func (pCmd *PluginCommand) RunCommand(cmd *cobra.Command, args []string) error {
	return nil
}

func (pCmd *PluginCommand) ResultPlainOutput() error {
	return output.FinalCmdOutputPlain(&pCmd.FinalCurveCmd)
}
