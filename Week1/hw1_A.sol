// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract BinaryConverter {
    function binaryToDecimal(string memory binary) public pure returns (uint256) {
        bytes memory b = bytes(binary);
        uint256 decimal = 0;
        uint256 power = 1; // 2^0

        for (uint256 i = b.length; i > 0; i--) {
            require(b[i - 1] == '0' || b[i - 1] == '1', "Invalid binary character");
            if (b[i - 1] == '1') {
                decimal += power;
            }
            power *= 2;
        }

        return decimal;
    }
}
