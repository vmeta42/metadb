/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"configcenter/src/source_controller/coreservice/bussiness/transactionBus"
	"context"
	"fmt"
	"os"
	"runtime"

	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/types"
	"configcenter/src/common/util"
	"configcenter/src/source_controller/coreservice/app"
	"configcenter/src/source_controller/coreservice/app/options"

	"github.com/spf13/pflag"
)

func main() {
	common.SetIdentification(types.CC_MODULE_CORESERVICE)
	runtime.GOMAXPROCS(runtime.NumCPU())

	blog.InitLogs()
	defer blog.CloseLogs()

	op := options.NewServerOption()
	op.AddFlags(pflag.CommandLine)

	util.InitFlags()
	transactionBus.PrintTxnMap()
	transactionBus.CleanTimeOutTxn()
	ctx, cancel := context.WithCancel(context.Background())
	if err := app.Run(ctx, cancel, op); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		blog.Errorf("process stopped by %v", err)
		blog.CloseLogs()
		os.Exit(1)
	}
}
