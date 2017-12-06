var Fabric_Client = require('fabric-client');
var path = require('path');
var util = require('util');
var os = require('os');

//
var fabric_client = new Fabric_Client();

// setup the fabric network
var channel = fabric_client.newChannel('mychannel');
var peer = fabric_client.newPeer('grpc://localhost:7051');
channel.addPeer(peer);
var order = fabric_client.newOrderer('grpc://localhost:7050')
channel.addOrderer(order);

//
var member_user = null;
var store_path = path.join(__dirname, 'hfc-key-store');
console.log('Store path:'+store_path);
var tx_id = null;

var prividex_transaction_id = "FF";
var consumer_id = "c1";
var consumer_public_key = "c_PK";
var user_id = "user1";
var user_public_key = "u_PK";
var provider_id = 'p1';
var provider_public_key = "p_PK";
var identity_asset_name = "kyc_compliance";
var signature_1 = 's1';
var signature_2 = 's2';

var user_defined_opts = {
    TransactionID: prividex_transaction_id,
    ConsumerID: consumer_id,
    ConsumerPublicKey: consumer_public_key,
    UserID: user_id,
    UserPublicKey: user_public_key,
    ProviderID: provider_id,
    ProviderPublicKey: provider_public_key,
    IdentityAssetName: identity_asset_name,
    Signature1: signature_1,
    Signature2: signature_2
};

var g_options = {
    fabric_client: fabric_client,
    channel: channel,
    peer: peer,
    order: order,
    member_user: "",
    store_path: store_path,
    tx_id: tx_id,
    user_id: 'user1',
    chaincode_id: 'prividex',
    fcn: 'iHandshake',
    channel_id: 'mychannel'
};

var HLClient = require('./index.js')(g_options);

HLClient.fwdinitHandshakeReqtoBC(user_defined_opts);

// HLClient.fwdHandshakeResptoBC(user_defined_opts);

//HLClient.fwdConfHandshaketoBC(user_defined_opts);

// var identity_asset = "ewogICJpZEFzc2V0IiA6IHsKICAgICJuYW1lT2ZJZEFzc2V0IiA6ICJLWUNfZm9yX0JhbmtpbmciLAogICAgIm5hbWVPZlVzZXIiICAgIDogIkNoZXJyeSBCZXJyeSIsCiAgICAicGFzc3BvcnQiIDogewogICAgICAibnVtYmVyIiA6ICJONTg3ODY1IiwKICAgICAgImV4cGlyYXRpb24tZGF0ZSIgOiAiMDYtMDctMjAyMiIsCiAgICAgICJpc3N1aW5nLWNvdW50cnkiIDogIlVESyIKICAgIH0sCiAgICAiU1NOIiA6IHsKICAgICAgIm51bWJlciIgOiAiNjc4MC05NS0zMjQ1IiwKICAgICAgImV4cGlyYXRpb24tZGF0ZSIgOiAiMjAtMDctMjAxMCIKICAgIH0KCiAgfQp9";

// var user_defined_opts = {
//     TransactionID: prividex_transaction_id,
//     ConsumerID: consumer_id,
//     ConsumerPublicKey: consumer_public_key,
//     UserID: user_id,
//     UserPublicKey: user_public_key,
//     ProviderID: provider_id,
//     ProviderPublicKey: provider_public_key,
//     IdentityAssetName: identity_asset_name,
//     IdAsset: identity_asset,
//     Signature1: signature_1,
//     Signature2: signature_2
// };

// HLClient.fwdinitAssetTransferReqtoBC(user_defined_opts);

// HLClient.fwdconfAssetTransferfromBC(user_defined_opts);


