var Fabric_Client = require('fabric-client');
var path = require('path');
var util = require('util');
var os = require('os');

// //
// var fabric_client = new Fabric_Client();

// // setup the fabric network
// var channel = fabric_client.newChannel('mychannel');
// var peer = fabric_client.newPeer('grpc://localhost:7051');
// channel.addPeer(peer);
// var order = fabric_client.newOrderer('grpc://localhost:7050')
// channel.addOrderer(order);

// var store_path = path.join(__dirname, 'hfc-key-store');

//
// var member_user = null;
// console.log('Store path:'+store_path);
// var tx_id = null;

//-------------------------------------------------------------------
// Fabric Client Wrangler - Wrapper library for the Hyperledger Fabric Client SDK
//-------------------------------------------------------------------


module.exports = function (g_options) {
	var invoke_cc = require('./invoke.js')();
	var query_cc = require('./query.js')();
	var HLClient = {};

	HLClient.fwdinitHandshakeReqtoBC = function (user_defined_options) {
		console.log("calling fwdinitHandshakeReqtoBC");
		var initHandShake_opts = {
		    HandshakeRecordType: "initHandshake",
		    TransactionID: user_defined_options['TransactionID'],
		    ConsumerID: user_defined_options['ConsumerID'],
		    ConsumerPublicKey: user_defined_options['ConsumerPublicKey'],
		    UserID: user_defined_options['UserID'],
		    UserPublicKey: user_defined_options['UserPublicKey'],
		    ProviderID: user_defined_options['ProviderID'],
		    ProviderPublicKey: user_defined_options['ProviderPublicKey'],
		    IdentityAssetName: user_defined_options['IdentityAssetName'],
		    Signature1: user_defined_options['Signature1'],
		    Signature2: user_defined_options['Signature2']
		};
		g_options['fcn'] = 'initHandshake';
		invoke_cc.invoke_chaincode(g_options, initHandShake_opts);
	};

	HLClient.fwdHandshakeResptoBC = function (user_defined_options) {
		console.log('fwdHandshakeResptoBC');
		var respHandshake_opts = {
		    HandshakeRecordType: "respHandshake",
		    TransactionID: user_defined_options['TransactionID'],
		    ConsumerID: user_defined_options['ConsumerID'],
		    ConsumerPublicKey: user_defined_options['ConsumerPublicKey'],
		    UserID: user_defined_options['UserID'],
		    UserPublicKey: user_defined_options['UserPublicKey'],
		    ProviderID: user_defined_options['ProviderID'],
		    ProviderPublicKey: user_defined_options['ProviderPublicKey'],
		    IdentityAssetName: user_defined_options['IdentityAssetName'],
		    Signature1: user_defined_options['Signature1'],
		};
		g_options['fcn'] = 'respHandshake';
		invoke_cc.invoke_chaincode(g_options, respHandshake_opts);
	};

	HLClient.fwdConfHandshaketoBC = function (user_defined_options) {
		console.log('fwdConfHandshaketoBC');
		var confHandshake_opts = {
		    HandshakeRecordType: "confHandshake",
		    TransactionID: user_defined_options['TransactionID'],
		    ConsumerID: user_defined_options['ConsumerID'],
		    ConsumerPublicKey: user_defined_options['ConsumerPublicKey'],
		    UserID: user_defined_options['UserID'],
		    UserPublicKey: user_defined_options['UserPublicKey'],
		    ProviderID: user_defined_options['ProviderID'],
		    ProviderPublicKey: user_defined_options['ProviderPublicKey'],
		    IdentityAssetName: user_defined_options['IdentityAssetName'],
		    Signature1: user_defined_options['Signature1'],
		};

		g_options['fcn'] = 'confHandshake';
		invoke_cc.invoke_chaincode(g_options, confHandshake_opts);
	};

	HLClient.fwdinitAssetTransferReqtoBC = function (user_defined_options) {
		var assetTransfer_opts = {
		    TransactionID: user_defined_options['TransactionID'],
		    ConsumerID: user_defined_options['ConsumerID'],
		    ConsumerPublicKey: user_defined_options['ConsumerPublicKey'],
		    UserID: user_defined_options['UserID'],
		    UserPublicKey: user_defined_options['UserPublicKey'],
		    ProviderID: user_defined_options['ProviderID'],
		    ProviderPublicKey: user_defined_options['ProviderPublicKey'],
		    IdentityAssetName: user_defined_options['IdentityAssetName'],
		    IdAsset: user_defined_options['IdAsset'],
		    Signature1: user_defined_options['Signature1'],
		};
		g_options['fcn'] = 'transferAsset';
		invoke_cc.invoke_chaincode(g_options, assetTransfer_opts);
	};

	HLClient.fwdconfAssetTransferfromBC = function (user_defined_options) {
		var confAssetTransfer_opts = {
		    TransactionID: user_defined_options['TransactionID'],
		    ConsumerID: user_defined_options['ConsumerID'],
		    ConsumerPublicKey: user_defined_options['ConsumerPublicKey'],
		    UserID: user_defined_options['UserID'],
		    UserPublicKey: user_defined_options['UserPublicKey'],
		    ProviderID: user_defined_options['ProviderID'],
		    ProviderPublicKey: user_defined_options['ProviderPublicKey'],
		    IdentityAssetName: user_defined_options['IdentityAssetName'],
		    Signature1: user_defined_options['Signature1'],
		};
		g_options['fcn'] = 'confirmReceiptOfAsset';
		invoke_cc.invoke_chaincode(g_options, confAssetTransfer_opts);
	};

	HLClient.getData = function (user_defined_options) {
		query_cc.invoke_chaincode(g_options, user_defined_options);
	};

	return HLClient;
};
