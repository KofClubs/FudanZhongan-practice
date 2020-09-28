/**
 *Submitted for verification at Etherscan.io on 2020-09-28
 */

pragma solidity ^0.4.18;

contract Test {
    string internal str;

    function Test() public {
        str = "Hello!";
    }

    function getStr() public view returns (string) {
        return str;
    }
}
