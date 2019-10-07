import Block from './Block';
import Transaction from './Transaction';

export default class Blockchain {
  constructor() {
    this.chain = [this.createGenesisBlock()];
    this.difficulty = 2;
    this.pendingTransactions = [];
    this.miningReward = 100;
  }

  createGenesisBlock() {
    return new Block('01/01/2017', 'Genesis block', '');
  }

  getLatestBlock() {
    return this.chain[this.chain.length - 1];
  }

  minePendingTransactions(miningRewardAddress) {
    const rewardTx = new Transaction(
      null,
      miningRewardAddress,
      this.miningReward,
    );
    this.pendingTransactions.push(rewardTx);

    let block = new Block(Date.now(), this.pendingTransactions);
    block.mineBlock(this.difficulty);

    console.log('Block successfully mined!');
    this.chain.push(block);

    this.pendingTransactions = [];
  }

  addTransaction(transaction) {
    if (!transaction.fromAddress || !transaction.toAddress) {
      throw new Error('Transaction must include from and to address');
    }

    if (!transaction.isValid()) {
      throw new Error('Cannot add invalid transaction to chain');
    }

    this.pendingTransactions.push(transaction);
  }

  getBalanceOfAddress(address) {
    let balance = 0;

    for (const block of this.chain) {
      for (const trans of block.transactions) {
        if (trans.fromAddress === address) {
          balance -= trans.amount;
        }
        if (trans.toAddress === address) {
          balance += trans.amount;
        }
      }
    }

    return balance;
  }

  addBlock(newBlock) {
    newBlock.previousHash = this.getLatestBlock().hash;
    newBlock.mineBlock(this.difficulty);
    this.chain.push(newBlock);
  }

  isChainValid() {
    for (let i = 1; i < this.chain.length; i++) {
      const currentBlock = this.chain[i];
      const previousBlock = this.chain[i - 1];

      if (!currentBlock.hasValidTransactions()) {
        return false;
      }

      if (currentBlock.hash !== currentBlock.calculateHash()) {
        return false;
      }

      if (currentBlock.previousHash !== previousBlock.hash) {
        return false;
      }

      return true;
    }
  }
}

let saveCoin = new Blockchain();

// saveCoin.addBlock(new Block("10/07/2017", { amount : 4 }));
// saveCoin.addBlock(new Block("12/07/2017", { amount : 10 }));

// console.log(JSON.stringify(saveCoin, null, 4));

// saveCoin.addTransaction(new Transaction('address1', 'address2', 100));
// saveCoin.addTransaction(new Transaction('address1', 'address2', 100));

// console.log('\n');

// saveCoin.minePendingTransactions('xavier-address');

// console.log(saveCoin.getBalanceOfAddress('xavier-address'));

// saveCoin.minePendingTransactions('xavier-address');

// console.log(saveCoin.getBalanceOfAddress('xavier-address'));

const mykey = ec.keyFromPrivate('.......');
const myWalletAddress = mykey.getPublic('hex');

const tx1 = new Transaction(myWalletAddress, 'public key goes here', 10);
tx1.signTransaction(mykey);
saveCoin.addTransaction(tx1);

console.log('\n');

saveCoin.minePendingTransactions(myWalletAddress);

console.log(saveCoin.getBalanceOfAddress(myWalletAddress));
