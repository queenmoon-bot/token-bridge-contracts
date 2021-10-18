// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2020, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

pragma solidity ^0.6.11;

import "arb-bridge-eth/contracts/libraries/BytesLib.sol";
import "arb-bridge-eth/contracts/libraries/DebugPrint.sol";

library BytesParser {
    using BytesLib for bytes;

    function toUint8(bytes memory input) internal pure returns (bool success, uint8 res) {
        if (input.length == 0) {
            success = false;
            // return the default value of uint8
        } else {
            // TODO: try catch to handle error
            success = true;
            res = abi.decode(input, (uint8));
        }
    }

    function toString(bytes memory input) internal pure returns (bool success, string memory res) {
        if (input.length == 0) {
            success = false;
            // return default value of string
        } else if (input.length == 32) {
            success = true;
            res = DebugPrint.bytes32string(input.toBytes32(0));
        } else {
            // TODO: try catch to handle error
            success = true;
            res = abi.decode(input, (string));
        }
    }
}
