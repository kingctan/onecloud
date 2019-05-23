// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package driver

import (
	"context"

	"github.com/pkg/errors"

	"yunion.io/x/onecloud/pkg/keystone/models"
	"yunion.io/x/onecloud/pkg/mcclient"
)

type SSQLDriver struct {
	SBaseDomainDriver
}

func NewSQLDriver(domainId string, conf models.TDomainConfigs) (IIdentityBackend, error) {
	drv := SSQLDriver{
		NewBaseDomainDriver(domainId, conf),
	}
	drv.virtual = &drv
	return &drv, nil
}

func (sql *SSQLDriver) Authenticate(ctx context.Context, ident mcclient.SAuthenticationIdentity) (*models.SUserExtended, error) {
	usrExt, err := models.UserManager.FetchUserExtended(
		ident.Password.User.Id,
		ident.Password.User.Name,
		ident.Password.User.Domain.Id,
		ident.Password.User.Domain.Name,
	)
	if err != nil {
		return nil, errors.Wrap(err, "UserManager.FetchUserExtended")
	}
	err = usrExt.VerifyPassword(ident.Password.User.Password)
	if err != nil {
		return nil, errors.Wrap(err, "usrExt.VerifyPassword")
	}
	return usrExt, nil
}
