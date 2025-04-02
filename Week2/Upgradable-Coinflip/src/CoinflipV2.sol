// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

error SeedTooShort();

/// @title Coinflip 10 in a Row (V2)
/// @author Jahad Jafarov
/// @notice Upgraded version of the Coinflip contract
contract CoinflipV2 is Initializable, OwnableUpgradeable, UUPSUpgradeable {
    string public seed;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /// @notice Initializes the contract (replaces the constructor)
    /// @param initialOwner The address of the initial owner
    function initialize(address initialOwner) public initializer {
        __Ownable_init(initialOwner);
        __UUPSUpgradeable_init();
        seed = "It is a good practice to rotate seeds often in gambling";
    }

    /// @notice Checks user input against contract generated guesses
    /// @param Guesses is a fixed array of 10 elements which holds the user's guesses. The guesses are either 1 or 0 for heads or tails
    /// @return true if user correctly guesses each flip correctly or false otherwise
    function UserInput(uint8[10] calldata Guesses) external view returns (bool) {
        uint8[10] memory flips = getFlips();

        for (uint8 i = 0; i < 10; i++) {
            if (Guesses[i] != flips[i]) {
                return false;
            }
        }

        return true;
    }

    /// @notice Allows the owner of the contract to change the seed to a new one
    /// @param NewSeed is a string which represents the new seed
    /// @param rotation is the number of characters to rotate
    function seedRotation(string memory NewSeed, uint256 rotation) public onlyOwner {
        bytes memory newSeedBytes = bytes(NewSeed);
        uint256 seedlength = newSeedBytes.length;

        if (seedlength < 10) {
            revert SeedTooShort();
        }

        // Rotate the seed
        bytes memory rotatedSeed = new bytes(seedlength);
        for (uint256 i = 0; i < seedlength; i++) {
            rotatedSeed[i] = newSeedBytes[(i + rotation) % seedlength];
        }

        // Set the rotated seed
        seed = string(rotatedSeed);
    }

    /// @notice This function generates 10 random flips by hashing characters of the seed
    /// @return a fixed 10 element array of type uint8 with only 1 or 0 as its elements
    function getFlips() public view returns (uint8[10] memory) {
        bytes memory stringInBytes = bytes(seed);
        uint256 seedlength = stringInBytes.length;
        uint8[10] memory results;
        uint256 interval = seedlength / 10;

        for (uint8 i = 0; i < 10; i++) {
            uint256 randomNum = uint256(keccak256(abi.encodePacked(stringInBytes[i * interval], block.timestamp)));
            results[i] = randomNum % 2 == 0 ? 1 : 0;
        }

        return results;
    }

    /// @notice Authorizes an upgrade to a new implementation
    /// @param newImplementation The address of the new implementation
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}
}