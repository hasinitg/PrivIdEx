//
// Created by Hasini Gunasinghe on 2/8/18.
//

#ifndef ZKSNARK_GENERATOR_HPP
#define ZKSNARK_GENERATOR_HPP

#include <libff/common/default_types/ec_pp.hpp>
#include "r1cs_examples.hpp"

namespace libsnark {
    r1cs_example<libff::Fr<libff::default_ec_pp>> generateInnerProductR1CS(const size_t size);

    template<typename ppT>
    r1cs_ppzksnark_keypair<ppT> generateKeyPairForR1CS(const r1cs_example<libff::Fr<ppT>> &example);

}

#include "generator.cpp"
#endif //ZKSNARK_GENERATOR_HPP
