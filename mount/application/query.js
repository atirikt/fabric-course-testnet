const {Wallets, Gateway} = require('fabric-network');

const fs = require('fs');
const path = require('path');

async function main(){
  try{
    const ccpPath = path.resolve(__dirname,'..','network', 'organizations', 'peerOrganizations', 'org1.example.com', 'connection-org1.json');
    const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));
    
    const walletPath = path.join(process.cwd(), 'wallet');
    const wallet = await Wallets.newFileSystemWallet(walletPath);

    const gateway = new Gateway();
    await gateway.connect(ccp, {wallet, identity:'appUser2', discovery:{enabled:true, asLocalhost:true}});
    
    const network = await gateway.getNetwork('channel1');
    const contract = network.getContract('fabcar');

    const result = await contract.evaluateTransaction('queryAllCars');
    await gateway.disconnect();
    
    console.log('success');
    console.log(`result: ${result.toString()}`);

  }catch (error){
    console.log(error);
    process.exit(1);
  }
}

main();