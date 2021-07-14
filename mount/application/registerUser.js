const FabricCAServices = require('fabric-ca-client');
const {Wallets} = require('fabric-network');

const fs = require('fs');
const path = require('path');

async function main(){
  try{
    const ccpPath = path.resolve(__dirname,'..','network', 'organizations', 'peerOrganizations', 'org1.example.com', 'connection-org1.json');
    const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));
    
    const caInfo = ccp.certificateAuthorities['ca.org1.example.com'];
    const ca = new FabricCAServices(caInfo.url);
    
    const walletPath = path.join(process.cwd(), 'wallet');
    const wallet = await Wallets.newFileSystemWallet(walletPath);

    const adminIdentity = await wallet.get('admin');
    const provider = wallet.getProviderRegistry().getProvider(adminIdentity.type);
    const adminUser = await provider.getUserContext(adminIdentity, 'admin');

    const secret = await ca.register({affiliation:'org1.department1', enrollmentID:'appUser2', role:'client'}, adminUser);
    const enrollment = await ca.enroll({enrollmentID: 'appUser2', enrollmentSecret:secret});
    const x509Identity = {credentials:{certificate:enrollment.certificate, privateKey:enrollment.key.toBytes()}, mspId:'Org1MSP',type:'X.509'};

    await wallet.put('appUser2', x509Identity);

    console.log('success');


  }catch (error){
    console.log(error);
    process.exit(1);
  }
}

main();