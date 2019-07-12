/******************************************************************************
 * Copyright © 2013-2019 The Komodo Platform Developers.                      *
 *                                                                            *
 * See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at                  *
 * the top-level directory of this distribution for the individual copyright  *
 * holder information and the developer policies on copyright and licensing.  *
 *                                                                            *
 * Unless otherwise agreed in a custom licensing agreement, no part of the    *
 * Komodo Platform software, including this file may be copied, modified,     *
 * propagated or distributed except according to the terms contained in the   *
 * LICENSE file                                                               *
 *                                                                            *
 * Removal or modification of this copyright notice is prohibited.            *
 *                                                                            *
 ******************************************************************************/

package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsKomodoAddress(t *testing.T) {
	assert.Truef(t, IsKomodoAddress("RM5wffThEVKQdG98uLa2gc8Nk4CzX9Fq4q"), "should be true")
	assert.Truef(t, IsKomodoAddress("RXrQPqU4SwARri1m2n7232TDECvjzXCJh4"), "should be true")
	assert.False(t, IsKomodoAddress("BXrQPqU4SwARri1m2n7232TDECvjzXCJh4"), "should be false")
}
