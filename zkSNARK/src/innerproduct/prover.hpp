//
// Created by Hasini Gunasinghe on 2/8/18.
//

#ifndef ZKSNARK_PROVER_HPP
#define ZKSNARK_PROVER_HPP

#include <libff/common/default_types/ec_pp.hpp>
#include "r1cs_examples.hpp"

namespace libsnark{
    r1cs_example<libff::Fr<libff::default_ec_pp>> generateInnerProductR1CSWithInputs(const size_t size);

    template <typename ppT>
    r1cs_ppzksnark_proof<ppT> generateProofForR1CS(const r1cs_ppzksnark_proving_key<ppT> &pk,
                                                  const r1cs_ppzksnark_primary_input<ppT> &primary_input,
                                                  const r1cs_ppzksnark_auxiliary_input<ppT> &auxiliary_input);
}

#include "prover.cpp"
#endif //ZKSNARK_PROVER_HPP
