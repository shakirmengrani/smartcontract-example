const SHA256 = require("crypto-js/sha256");

const Block = function(index, data, timestamp = new Date(), prevHash = ""){

    Block.prototype.calculateHash = function(){
        return SHA256(this.index + this.timestamp + this.prevHash + JSON.stringify(this.data)).toString();
    };
    this.index = index;
    this.data = data;
    this.timestamp = timestamp;
    this.prevHash = prevHash;
    this.hash = this.calculateHash();
};


const Blockchain = function(){

    Blockchain.prototype.createGensisBLock = function(){
        return new Block(0, "Genesis block"); 
    };

    Blockchain.prototype.getLatestBlock = function(){
        return this.chain[this.chain.length - 1];
    };

    Blockchain.prototype.addBlock = function(newBlock) { 
        newBlock.prevHash = this.getLatestBlock().hash; 
        newBlock.hash = newBlock.calculateHash(); 
        this.chain.push(newBlock); 
    };

    Blockchain.prototype.isChainValid = function() { 
        for (let i = 1; i < this.chain.length; i++){ 
            const currentBlock = this.chain[i]; 
            const prevBlock = this.chain[i - 1]; 
            if (currentBlock.hash !== currentBlock.calculateHash()) { 
                return false; 
            } 

            if (currentBlock.prevHash !== prevBlock.hash) { 
                return false; 
            } 
        } 
        return true; 
    } 
    this.chain = [this.createGensisBLock()];
}

const coin = new Blockchain()
for(let i = 1; i <= 10; i++){
    coin.addBlock(new Block(i, {amount: 8 * i}));
}
console.log(coin.chain);