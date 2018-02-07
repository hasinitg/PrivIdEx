#include <iostream>
#include <libsnark/gadgetlib2/gadget.hpp>
#include <libsnark/gadgetlib2/pp.hpp>
#include "simple_example.hpp"
#include "r1cs_examples.hpp"
#include "run_r1cs_ppzksnark.hpp"

int main() {
    gadgetlib2::initPublicParamsFromDefaultPp();
    // Create an example constraint system and translate to libsnark format
    const libsnark::r1cs_example<libff::Fr<libff::default_ec_pp> > example = libsnark::gen_r1cs_example_from_gadgetlib2_protoboard(100);
    const bool test_serialization = false;
    // Run ppzksnark. Jump into function for breakdown
    const bool bit = libsnark::run_r1cs_ppzksnark<libff::default_ec_pp>(example, test_serialization);
    std::cout << "Hello, World!" << std::endl;
    return 0;
}

