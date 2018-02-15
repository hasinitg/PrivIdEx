#include <iostream>
#include <libsnark/gadgetlib2/pp.hpp>
#include <libsnark/gadgetlib2/gadget.hpp>
//#include "simple_example.hpp"
#include "run_r1cs_ppzksnark.hpp"
#include "innerproduct/generator.hpp"
#include "innerproduct/prover.hpp"
#include "innerproduct/verifier.hpp"

int main() {

    //get the generator to create a constraint system in order to generate public, private keys based on that.
    const libsnark::r1cs_example<libff::Fr<libff::default_ec_pp>> generatorCS = libsnark::generateInnerProductR1CS(100);

    libsnark::r1cs_ppzksnark_keypair<libff::default_ec_pp> keyPair = libsnark::generateKeyPairForR1CS<libff::default_ec_pp>(generatorCS);

    //get the prover to create a constraint system with the inputs, in order to generate proof based on that the proof key.
    const libsnark::r1cs_example<libff::Fr<libff::default_ec_pp>> proverCS = libsnark::generateInnerProductR1CSWithInputs(100);

    libsnark::r1cs_ppzksnark_proof<libff::default_ec_pp> proof = libsnark::generateProofForR1CS<libff::default_ec_pp>(
            keyPair.pk, proverCS.primary_input, proverCS.auxiliary_input);

    const bool ans = libsnark::verifyProofOnR1CS<libff::default_ec_pp>(keyPair.vk, proverCS.primary_input, proof);

//    gadgetlib2::initPublicParamsFromDefaultPp();
//    // Create an example constraint system and translate to libsnark format
//    const libsnark::r1cs_example<libff::Fr<libff::default_ec_pp> > example = libsnark::gen_r1cs_example_from_gadgetlib2_protoboard(100);
//    const bool test_serialization = false;
//    // Run ppzksnark. Jump into function for breakdown
//    const bool bit = libsnark::run_r1cs_ppzksnark<libff::default_ec_pp>(example, test_serialization);

    std::cout << "Hello, World!" << std::endl;
    std::cout << ans << std::endl;
    return 0;
}