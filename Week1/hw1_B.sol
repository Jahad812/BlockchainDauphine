// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract BitwiseOperations {

    function getBits(uint8 _number) public pure returns (uint256[] memory) {
        uint256 count = 0; // Count of set bits

        // Count the number of set bits
        for (uint8 i = 0; i < 8; i++) {
            if ((_number >> i) & 1 == 1) { 
                count++;
            }
        }

        // Create an array to store the powers of 2 for set bits
        uint256[] memory result = new uint256[](count); 

        // Populate the array with powers of 2 for set bits
        count = 0; 
        for (uint8 i = 0; i < 8; i++) {
            if ((_number >> i) & 1 == 1) { 
                result[count] = 1 << i; // Store 2^i for set bits
                count++;
            }
        }

        return result;
    }
}