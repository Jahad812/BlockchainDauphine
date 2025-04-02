// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

error SeedTooShort();

/// @title Coinflip 10 in a Row
/// @author Jahad Jafarov
/// @notice Contract used as part of the course Solidity and Smart Contract development
contract Coinflip is Ownable {
    string public seed;

    constructor() Ownable(msg.sender) {
        // Set the seed to "It is a good practice to rotate seeds often in gambling"
        seed = "It is a good practice to rotate seeds often in gambling";
    }

    /// @notice Checks user input against contract generated guesses
    /// @param Guesses is a fixed array of 10 elements which holds the user's guesses. The guesses are either 1 or 0 for heads or tails
    /// @return true if user correctly guesses each flip correctly or false otherwise
    function userInput(uint8[10] calldata Guesses) external view returns (bool) {
        // Get the contract generated flips by calling the helper function getFlips()
        uint8[10] memory flips = getFlips();

        // Compare each element of the user's guesses with the generated guesses
        for (uint8 i = 0; i < 10; i++) {
            if (Guesses[i] != flips[i]) {
                return false; // Return false if any guess does not match
            }
        }

        return true; // Return true if all guesses match
    }

    /// @notice Allows the owner of the contract to change the seed to a new one
    /// @param NewSeed is a string which represents the new seed
    function seedRotation(string memory NewSeed) public onlyOwner {
        // Cast the string into a bytes array so we may perform operations on it
        bytes memory newSeedBytes = bytes(NewSeed);

        // Get the length of the array (i.e., how many characters in this string)
        uint256 seedlength = newSeedBytes.length;

        // Check if the seed is less than 10 characters
        if (seedlength < 10) {
            revert SeedTooShort();
        }

        // Set the seed variable as the NewSeed
        seed = NewSeed;
    }

    // -------------------- helper functions -------------------- //
    /// @notice This function generates 10 random flips by hashing characters of the seed
    /// @return a fixed 10 element array of type uint8 with only 1 or 0 as its elements
    function getFlips() public view returns (uint8[10] memory) {
        // Cast the seed into a bytes array and get its length
        bytes memory stringInBytes = bytes(seed);
        uint256 seedlength = stringInBytes.length;

        // Initialize an empty fixed array with 10 uint8 elements
        uint8[10] memory results;

        // Setting the interval for grabbing characters
        uint256 interval = seedlength / 10;

        // Input the correct form for the for loop
        for (uint8 i = 0; i < 10; i++) {
            // Generating a pseudo-random number by hashing together the character and the block timestamp
            uint256 randomNum = uint256(keccak256(abi.encodePacked(stringInBytes[i * interval], block.timestamp)));

            // If the result is an even unsigned integer, record it as 1 in the results array, otherwise record it as zero
            results[i] = randomNum % 2 == 0 ? 1 : 0;
        }

        // Return the resulting fixed array
        return results;
    }
}