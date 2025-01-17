// Copyright 2018-2019 CERN
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
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

package datasvc

import (
	"net/http"
	"strings"

	"github.com/cs3org/reva/pkg/appctx"
)

func (s *svc) doPut(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := appctx.GetLogger(ctx)
	fn := r.URL.Path

	fsfn := strings.TrimPrefix(fn, s.conf.ProviderPath)
	err := s.storage.Upload(ctx, fsfn, r.Body)
	if err != nil {
		log.Error().Err(err).Msg("error uploading file")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.Body.Close()
	w.WriteHeader(http.StatusOK)
}
