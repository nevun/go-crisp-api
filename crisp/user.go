// Copyright 2016 Crisp IM, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


// UserService service
type UserService service


// UserFilter mapping
type UserFilter struct {
  Model      *string         `json:"model,omitempty"`
  Criterion  *string         `json:"criterion,omitempty"`
  Operator   *string         `json:"operator,omitempty"`
  Query      *[]interface{}  `json:"query,omitempty"`
}
