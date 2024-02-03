/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package shuffle

import (
	"crypto/rand"
	"math/big"
)

func Byte(set *[]byte) {

	arrSize := len(*set)

	for i := arrSize - 1; i > 0; i-- {
		jBigInt, err := rand.Int(rand.Reader, big.NewInt(int64(arrSize)))
		if err != nil {
			panic(err)
		}
		j := int(jBigInt.Int64())
		(*set)[i], (*set)[j] = (*set)[j], (*set)[i]
	}

}

func String(pwd *string) {

	pwdBytes := []byte(*pwd)
	strSize := len(pwdBytes)

	for i := strSize - 1; i > 0; i-- {
		jBigInt, err := rand.Int(rand.Reader, big.NewInt(int64(strSize)))
		if err != nil {
			panic(err)
		}
		j := int(jBigInt.Int64())
		pwdBytes[i], pwdBytes[j] = pwdBytes[j], pwdBytes[i]
	}

	*pwd = string(pwdBytes)

}

func FisherYates(set []byte, length int) []byte {
	return fisherYatesSelector(set, length)
}
