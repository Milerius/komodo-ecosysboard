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
	assert.Truef(t, IsLookLikeAKomodoAddress("RM5wffThEVKQdG98uLa2gc8Nk4CzX9Fq4q"), "should be true")
	assert.Truef(t, IsLookLikeAKomodoAddress("RXrQPqU4SwARri1m2n7232TDECvjzXCJh4"), "should be true")
	assert.Falsef(t, IsLookLikeAKomodoAddress("BXrQPqU4SwARri1m2n7232TDECvjzXCJh4"), "should be false")
}

func TestIsBlock(t *testing.T) {
	assert.Truef(t, IsLookLikeABlock("123456"), "should be true")
	assert.Truef(t, IsLookLikeABlock("123456"), "should be true")
	assert.Falsef(t, IsLookLikeABlock("RXrQPqU4SwARri1m2n7232TDECvjzXCJh4"), "should be false")
}

func TestIsLookLikeABlockHashOrTransactionId(t *testing.T) {
	assert.Truef(t, IsLookLikeABlockHashOrTransactionId("000000014c0797b609abef168e8df13c03b92415f3a9b00c9f583013b5824b06"), "should be true")
	assert.Truef(t, IsLookLikeABlockHashOrTransactionId("2909f0b98ca4c3812bb500bc79fa83d3b8c8159c8f4328ac80777968f59400c7"), "should be true")
	assert.Falsef(t, IsLookLikeABlockHashOrTransactionId("RXrQPqU4SwARri1m2n7232TDECvjzXCJh4"), "should be false")
}

func TestIsPathExist(t *testing.T) {
	assert.False(t, IsPathExist("/nonexistent"), "should be false")
	assert.True(t, IsPathExist("/tmp"), "should be true")
}
