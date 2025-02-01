// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

contract ABIEncodingSize {

    function computeHelloWorldABI() public pure returns (uint256) {
        // Define the input values
        uint[] memory array = new uint[](2);
        array[0] = 1993;
        array[1] = 1994;
        bool flag = true;

        // ABI encodePacked the function selector + inputs
        bytes memory encodedData = abi.encodePacked(
            bytes4(keccak256("HelloWorld(uint[],bool)")), // Function selector
            array,                                         // Array input
            flag                                           // Boolean input
        );

        // Return the total size in bytes
        return encodedData.length;
    }
}