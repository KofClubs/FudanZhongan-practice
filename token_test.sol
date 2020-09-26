pragma solidity ^0.4.18;

interface tokenRecipient {
    function receiveApproval(
        address _from,
        uint256 _value,
        address _token,
        bytes _extraData
    ) public;
}

contract TokenTest {
    string public name;
    string public symbol;
    uint8 public decimals = 18;
    uint256 public totalSupply;

    mapping(address => uint256) public balanceOf;
    mapping(address => mapping(address => uint256)) public allowance;
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Burn(address indexed from, uint256 value);

    function TokenTest (
        string tokenName,
        string tokenSymbol,
        uint256 initialSupply,
    ) public {
        name = tokenName; // 名称
        symbol = tokenSymbol; // 符号
        totalSupply = initialSupply * 10**uint256(decimals); // 总供应量
        balanceOf[msg.sender] = totalSupply; // 智能合约创建者全得
    }

    function _transfer (
        address _from,
        address _to,
        uint256 _value
    ) internal {
        require(_to != 0x0);
        require(balanceOf[_from] >= _value);
        require(balanceOf[_to] + _value >= balanceOf[_to]); // 防止溢出
        balanceOf[_from] -= _value;
        balanceOf[_to] += _value;
        Transfer(_from, _to, _value);
    }

    function transfer (
        address _to,
        uint256 _value
    ) public {
        _transfer(msg.sender, _to, _value);
    }

    function transferFrom (
        address _from,
        address _to,
        uint256 _value
    ) public returns (bool success) {
        require(_value <= allowance[_from][msg.sender]);
        allowance[_from][msg.sender] -= _value;
        _transfer(_from, _to, _value);
        return true;
    }

    function approve (
        address _spender,
        uint256 _value
    ) public returns (bool success) {
        allowance[msg.sender][_spender] = _value;
        return true;
    }

    function burn (
        uint256 _value
    ) public returns(bool success) {
        require(balanceOf[msg.sender] >= _value);
        balanceOf[msg.sender] -= _value;
        totalSupply -= _value;
        Burn(msg.sender, _value);
        return true;
    }

    function burnFrom (
        address _from,
        uint256 _value
    ) public returns (bool success) {
        require(balanceOf[_from] >= _value);
        require(_value <= allowance[_from][msg.sender]);
        balanceOf[_from] -= _value;
        allowance[_from][msg.sender] -= _value;
        totalSupply -= _value;
        Burn(_from, _value);
        return true;
    }
}
