// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package store

import (
	"crypto/sha1" //nolint:gosec
	"encoding/json"
	"fmt"
	"sync"

	"github.com/99designs/keyring"
	"github.com/streamnative/streamnative-mcp-server/pkg/auth"
	"k8s.io/utils/clock"
)

type KeyringStore struct {
	kr    keyring.Keyring
	clock clock.Clock
	lock  sync.Mutex
}

// storedItem represents an item stored in the keyring
type storedItem struct {
	Audience string
	UserName string
	Grant    auth.AuthorizationGrant
}

// NewKeyringStore creates a store based on a keyring.
func NewKeyringStore(kr keyring.Keyring) (*KeyringStore, error) {
	return &KeyringStore{
		kr:    kr,
		clock: clock.RealClock{},
	}, nil
}

var _ Store = &KeyringStore{}

func (f *KeyringStore) SaveGrant(audience string, grant auth.AuthorizationGrant) error {
	f.lock.Lock()
	defer f.lock.Unlock()

	var err error
	var userName string
	switch grant.Type {
	case auth.GrantTypeClientCredentials:
		if grant.ClientCredentials == nil {
			return ErrUnsupportedAuthData
		}
		userName = grant.ClientCredentials.ClientEmail
	default:
		return ErrUnsupportedAuthData
	}
	item := storedItem{
		Audience: audience,
		UserName: userName,
		Grant:    grant,
	}
	err = f.setItem(item)
	if err != nil {
		return err
	}
	return nil
}

func (f *KeyringStore) LoadGrant(audience string) (*auth.AuthorizationGrant, error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	item, err := f.getItem(audience)
	if err != nil {
		if err == keyring.ErrKeyNotFound {
			return nil, ErrNoAuthenticationData
		}
		return nil, err
	}
	switch item.Grant.Type {
	case auth.GrantTypeClientCredentials:
		if item.Grant.ClientCredentials == nil {
			return nil, ErrUnsupportedAuthData
		}
	default:
		return nil, ErrUnsupportedAuthData
	}
	return &item.Grant, nil
}

func (f *KeyringStore) WhoAmI(audience string) (string, error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	key := hashKeyringKey(audience)
	var label string
	var err error
	authItem, err := f.kr.GetMetadata(key)
	switch {
	case err == nil && authItem.Item != nil:
		label = authItem.Label
	case err == keyring.ErrMetadataNeedsCredentials || authItem.Item == nil:
		// in this case, the backend doesn't support looking up metadata
		// some backends properly return an error message, others just return nil
		// so we still need to grab the whole document to look up the label
		item, getErr := f.kr.Get(key)
		switch {
		case getErr == keyring.ErrKeyNotFound:
			err = ErrNoAuthenticationData
		case getErr != nil:
			err = fmt.Errorf("unable to get information from the keyring: %v", err)
		default:
			// clear the error as it worked
			err = nil
			label = item.Label
		}
	case err == keyring.ErrKeyNotFound:
		err = ErrNoAuthenticationData
	default:
		err = fmt.Errorf("unable to get information from the keyring: %v", err)
	}

	return label, err
}

func (f *KeyringStore) Logout() error {
	f.lock.Lock()
	defer f.lock.Unlock()

	var err error
	keys, err := f.kr.Keys()
	if err != nil {
		return fmt.Errorf("unable to get information from the keyring: %v", err)
	}
	for _, key := range keys {
		err = f.kr.Remove(key)
	}
	if err != nil {
		return fmt.Errorf("unable to update the keyring: %v", err)
	}
	return nil
}

func (f *KeyringStore) getItem(audience string) (storedItem, error) {
	key := hashKeyringKey(audience)
	i, err := f.kr.Get(key)
	if err != nil {
		return storedItem{}, err
	}
	var grant auth.AuthorizationGrant
	err = json.Unmarshal(i.Data, &grant)
	if err != nil {
		// the grant appears to be invalid
		return storedItem{}, ErrUnsupportedAuthData
	}
	return storedItem{
		Audience: audience,
		UserName: i.Label,
		Grant:    grant,
	}, nil
}

func (f *KeyringStore) setItem(item storedItem) error {
	key := hashKeyringKey(item.Audience)
	data, err := json.Marshal(item.Grant)
	if err != nil {
		return err
	}
	i := keyring.Item{
		Key:                         key,
		Data:                        data,
		Label:                       item.UserName,
		Description:                 "authorization grant",
		KeychainNotTrustApplication: false,
		KeychainNotSynchronizable:   false,
	}
	err = f.kr.Set(i)
	if err != nil {
		return fmt.Errorf("unable to update the keyring: %v", err)
	}
	return nil
}

// hashKeyringKey creates a safe key based on the given string
func hashKeyringKey(s string) string {
	h := sha1.New() //nolint:gosec
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
