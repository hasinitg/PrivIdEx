var Fabric_Client = require('fabric-client');
var path = require('path');
var util = require('util');
var os = require('os');

// "tester" is the name of the docker container for Tomcat
var GLOBALID_SERVER_URL = "http://tester:8080/Tester/MiDocUpdaterServlet"
//var GLOBALID_SERVER_URL = "http://localhost:8080/Tester/MiDocUpdaterServlet"
// Replace chaincode ID accordingly
var CHAINCODEID = "prividex"

// var hfc = require('hfc');
// var util = require('util');

// // Create a client chain.	
// var chain = hfc.newChain('clientChain');
// //
// // set the port where the event service will listen
// chain.eventHubConnect("grpc://localhost:7053");
// //
// // // Get the eventHub service associated with the chain
// var eventHub = chain.getEventHub();


var fabric_client = new Fabric_Client();
// setup the fabric network
var channel = fabric_client.newChannel('mychannel');
var peer = fabric_client.newPeer('grpc://localhost:7051');
channel.addPeer(peer);
var order = fabric_client.newOrderer('grpc://localhost:7050')
channel.addOrderer(order);
var store_path = path.join(__dirname, 'hfc-key-store');


Fabric_Client.newDefaultKeyValueStore({ path: store_path
}).then((state_store) => {
  // assign the store to the fabric client
  fabric_client.setStateStore(state_store);
  var crypto_suite = Fabric_Client.newCryptoSuite();
  // use the same location for the state store (where the users' certificate are kept)
  // and the crypto store (where the users' keys are kept)
  var crypto_store = Fabric_Client.newCryptoKeyStore({path: store_path});
  crypto_suite.setCryptoKeyStore(crypto_store);
  fabric_client.setCryptoSuite(crypto_suite);

  // get the enrolled user from persistence, this user will sign all requests
  return fabric_client.getUserContext('user1', true);
}).then((user_from_store) => {

	let event_hub = fabric_client.newEventHub();
	event_hub.setPeerAddr('grpc://localhost:7053');
	event_hub.connect();

	console.log("here");
	var regEventId = event_hub.registerChaincodeEvent(CHAINCODEID, "InitHandShakeEvent", function(event) { 
    console.log("Entered");
	console.log(util.format("InitHandShakeEvent event received with payload: %j\n", event.payload.toString())); 
	// var params = event.payload.toString().split("***");
	// sendPost({"NRIC_Number":params[0], "UpdateAttributes":params[2]});
	});

}).catch((err) => {
  console.error('Failed to invoke successfully :: ' + err);
});


/*
let event_hub = fabric_client.newEventHub();
event_hub.setPeerAddr('grpc://localhost:7053');


// Register for Global ID registration event
var regEventId = eventHub.registerChaincodeEvent(CHAINCODEID, "InitHandShakeEvent", function(event) { 
	console.log(util.format("InitHandShakeEvent event received with payload: %j\n", event.payload.toString())); 
	// var params = event.payload.toString().split("***");
	// sendPost({"NRIC_Number":params[0], "UpdateAttributes":params[2]});
	});

// Register for Global ID update event
// var upEventId = eventHub.registerChaincodeEvent(CHAINCODEID, "updateEvent", function(event) { 
// 	console.log(util.format("Update ID event received, payload: %j\n", event.payload.toString()));
// 	var params = event.payload.toString().split("***");
// 	sendPost({"NRIC_Number":params[0], "UpdateAttributes":params[2]});
// 	});

function sendPost(payload) {
	var Client = require('node-rest-client').Client;
	console.log(payload);
 
	var client = new Client();
	var args = {
		parameters: payload,
		headers: { "Content-Type": "text/plain" }
	};
	client.post(GLOBALID_SERVER_URL, args, function (data, response) {
		console.log(data);
		//console.log(response);
	});
}



//Unregister events or a specific chaincode
//         eventHub.unregisterChaincodeEvent(registrationID);
//
//         // disconnect when done listening for events
//         process.on('exit', function() {
//             chain.eventHubDisconnect();
//             });
//

*/
