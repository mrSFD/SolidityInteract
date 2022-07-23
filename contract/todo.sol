pragma solidity >=0.5.0 <0.9.0;
contract Todo{
    address public owner;
    Task[] tasks;
    struct Task{
		string name;
		string message;
        uint id;
    }
    
    constructor(){
        owner = msg.sender;
    }
	
    modifier isOwner(){
		require(owner == msg.sender);
		_;
	}
	
    function add(string memory _name, string memory _message, uint _id) public isOwner {
		tasks.push(Task(_name,_message, _id));
    }

    function get(uint _id) public view returns (Task memory) {
        return tasks[_id];
    }
    
    function list() external view returns (Task[] memory){
        return tasks;
    }
}