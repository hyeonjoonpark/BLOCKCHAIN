/**
 * print Hello, World with solidity
 */

// SPDX-License-Identifier: MIT
pragma solidity >=0.4.17 <0.6.0;

contract HelloWorld {
    function hello() public pure returns (string memory) {
        return "Hello, World!";
    }
}
