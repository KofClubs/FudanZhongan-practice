pragma solidity ^0.4.22;

/// @title 报名
contract SignUpSys {
    address public creator; // 智能合约创造者
    uint256 public count; // 报名总数
    mapping(address => bool) public signUpList; // 报名表

    constructor() public {
        creator = msg.sender;
        count = 1;
        signUpList[creator] = true;
    }

    function SignUp() public {
        require(creator != msg.sender);
        require(!signUpList[msg.sender], "This address is signed up.");
        signUpList[msg.sender] = true;
        count++;
    }
}
