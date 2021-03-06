/*
 * Copyright 2014 The starfruit Authors. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package module

import (
	"github.com/flatpeach/starfruit/message"
	"github.com/flatpeach/starfruit/server"
	"github.com/flatpeach/starfruit/user"
	"strings"
)

type Ison struct{}

func (module *Ison) Handle(s *server.Server, u *user.User, m *message.Message) error {
	// ISON <nickname> *( SPACE <nickname> )

	if len(m.Params) == 0 {
		u.SendErrorNeedMoreParams("ISON")

		return nil
	}

	var existedNickNames []string

	for _, nickname := range m.Params {
		if s.IsNickNameRegistered(nickname) {
			existedNickNames = append(existedNickNames, nickname)
		}
	}

	u.SendMessage(message.New(
		s.Config.Server.Name,
		message.RPL_ISON,
		[]string{u.NickName},
		strings.Join(existedNickNames, " "),
	))

	return nil
}
